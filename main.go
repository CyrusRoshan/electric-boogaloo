package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/tarm/serial"
)

func main() {

	// 1411 is left, 1412 is right

	fretduinoPort := &serial.Config{Name: "/dev/cu.usbmodem1421", Baud: 9600}
	fretduino, err := serial.OpenPort(fretduinoPort)
	if err != nil {
		log.Fatal(err)
	}

	strumduinoPort := &serial.Config{Name: "/dev/cu.usbmodem1411", Baud: 9600}
	strumduino, err := serial.OpenPort(strumduinoPort)
	if err != nil {
		log.Fatal(err)
	}

	var strumduinoString string
	strumduinoOutput := make([]int, 4)

	var fretduinoString string
	rawFretduinoOutput := make([]int, 4)
	fretduinoOutput := make([]int, 4)

	// server stuff here
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"neck": fretduinoOutput,
			"body": strumduinoOutput,
		})
	})

	type Asdf struct {
		Strings []int `form:"strings" json:"strings" binding:"required"`
	}

	router.POST("/", func(c *gin.Context) {
		var json Asdf
		c.BindJSON(&json)

		fmt.Println(json.Strings)

		fretduino.Write([]byte("z"))
		fretduino.Write([]byte("Z"))
		switch json.Strings[0] {
		case 1:
			fretduino.Write([]byte("0"))
		case 2:
			fretduino.Write([]byte("1"))
		case 3:
			fretduino.Write([]byte("2"))
		}

		switch json.Strings[1] {
		case 1:
			fretduino.Write([]byte("3"))
		case 2:
			fretduino.Write([]byte("4"))
		case 3:
			fretduino.Write([]byte("5"))
		}

		switch json.Strings[2] {
		case 1:
			fretduino.Write([]byte("6"))
		case 2:
			fretduino.Write([]byte("7"))
		case 3:
			fretduino.Write([]byte("8"))
		}

		switch json.Strings[3] {
		case 1:
			fretduino.Write([]byte("9"))
		case 2:
			fretduino.Write([]byte("a"))
		case 3:
			fretduino.Write([]byte("b"))
		}

		fmt.Println(json.Strings)
	})
	go router.Run()

	// read in data here
	for {
		buf := make([]byte, 128)
		n, err := strumduino.Read(buf)
		panicif(err)

		strumduinoString += string(buf[:n])

		if strings.Contains(strumduinoString, "\n") {
			start := strings.Index(strumduinoString, "[")
			if start == -1 {
				strumduinoString = ""
				continue
			}

			err = json.Unmarshal([]byte(strumduinoString)[:strings.Index(strumduinoString, "\n")], &strumduinoOutput)
			if err != nil { //needed because of first read in being halfway through a json output and hackathon code so no singleton
				fmt.Println(err)
				strumduinoString = ""
				continue
			}

			strumduinoString = ""
		}

		buf = make([]byte, 128)
		n, err = fretduino.Read(buf)
		panicif(err)

		fretduinoString += string(buf[:n])

		if strings.Contains(fretduinoString, "\n") {
			start := strings.Index(fretduinoString, "[")
			if start == -1 {
				fretduinoString = ""
				continue
			}

			err = json.Unmarshal([]byte(fretduinoString)[:strings.Index(fretduinoString, "\n")], &rawFretduinoOutput)
			if err != nil { //needed because of first read in being halfway through a json output and hackathon code so no singleton
				fmt.Println(err)
				fretduinoString = ""
				continue
			}

			fretduinoString = ""
		}

		for i := 0; i <= 3; i++ {
			fretduinoOutput[i] = 0
		}

		for i := 0; i <= 2; i++ {
			switch {
			case nearTen(rawFretduinoOutput[i], 510):
				fretduinoOutput[0] = i + 1
			case nearTen(rawFretduinoOutput[i], 340):
				fretduinoOutput[1] = i + 1
			case nearTen(rawFretduinoOutput[i], 257):
				fretduinoOutput[2] = i + 1
			case nearTen(rawFretduinoOutput[i], 197):
				fretduinoOutput[3] = i + 1
				///
			case nearTen(rawFretduinoOutput[i], 610):
				fretduinoOutput[0] = i + 1
				fretduinoOutput[1] = i + 1
			case nearTen(rawFretduinoOutput[i], 465):
				fretduinoOutput[1] = i + 1
				fretduinoOutput[2] = i + 1
			case nearTen(rawFretduinoOutput[i], 380):
				fretduinoOutput[2] = i + 1
				fretduinoOutput[3] = i + 1
				///
			case nearTen(rawFretduinoOutput[i], 660):
				fretduinoOutput[0] = i + 1
				fretduinoOutput[1] = i + 1
				fretduinoOutput[2] = i + 1
			case nearTen(rawFretduinoOutput[i], 530):
				fretduinoOutput[1] = i + 1
				fretduinoOutput[2] = i + 1
				fretduinoOutput[3] = i + 1
				//
			case nearTen(rawFretduinoOutput[i], 685):
				fretduinoOutput[0] = i + 1
				fretduinoOutput[1] = i + 1
				fretduinoOutput[2] = i + 1
				fretduinoOutput[3] = i + 1
			default:
			}
		}

		// 0: 500 - 520 : 510
		// 1: 335 - 345 : 340
		// 2: 250 - 265 : 257
		// 3: 190 - 205 : 197
		//
		// 0, 1: 605 - 615 : 610
		// 1, 2: 465
		// 2, 3: 380
		//
		// 0, 1, 2, 3: 685
		//
		// 0, 1, 2: 660
		// 1, 2, 3: 530

		fmt.Println(strumduinoOutput, rawFretduinoOutput, fretduinoOutput)

		for i := 0; i < len(strumduinoOutput); i++ {
			if strumduinoOutput[i] == 0 {
				continue
			}
		}
	}
}

func nearTen(input int, val int) bool {
	return (val-10 <= input) && (val+10 >= input)
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
