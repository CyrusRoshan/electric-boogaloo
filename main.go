package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tarm/serial"
)

func main() {

	// 1411 is left, 1412 is right

	// fretduino := &serial.Config{Name: "/dev/cu.usbmodem1411", Baud: 9600}
	// s, err := serial.OpenPort(fretduino)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	strumduinoPort := &serial.Config{Name: "/dev/cu.usbmodem1411", Baud: 9600}
	strumduino, err := serial.OpenPort(strumduinoPort)
	if err != nil {
		log.Fatal(err)
	}

	var strumduinoString string
	strumduinoOutput := make([]int, 4)

	// server stuff here
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run()

	// read in data here
	for {
		buf := make([]byte, 128)
		n, err := strumduino.Read(buf)
		panicif(err)

		strumduinoString += string(buf[:n])

		if strings.Contains(strumduinoString, "\n") {
			err = json.Unmarshal([]byte(strumduinoString), &strumduinoOutput)
			if err != nil { //needed because of first read in being halfway through a json output and hackathon code so no singleton
				fmt.Println(err)
				strumduinoString = ""
				continue
			}

			fmt.Println(strumduinoOutput)

			strumduinoString = ""
		}

	}
}

func panicif(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MustMarshal(data interface{}) []byte {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return out
}
