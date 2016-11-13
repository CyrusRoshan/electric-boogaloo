int minPin = 2;
int maxPin = 13;
int delayTime = 100;

void setup() {
        pinMode(A0, INPUT);
        pinMode(A1, INPUT);
        pinMode(A2, INPUT);

        for (int i = minPin; i <= maxPin; i++) {
                pinMode(i, OUTPUT);
        }

        Serial.begin(9600);
}



void loop() {
        Serial.print("[");
        Serial.print(analogRead(A0));
        Serial.print(", ");
        Serial.print(analogRead(A1));
        Serial.print(", ");
        Serial.print(analogRead(A2));
        Serial.print("]");
        Serial.println();

        delay(delayTime);
}


void serialEvent() {
        if (Serial.available()){
                char inChar = (char)Serial.read();
                if (inChar == 'z'){
                        for (int i = minPin; i <= maxPin; i++) {
                                digitalWrite(i, LOW);
                        }
                } else if (inChar == '0'){
                        digitalWrite(minPin, HIGH);
                } else if (inChar == '1'){
                        digitalWrite(minPin + 1, HIGH);
                } else if (inChar == '2'){
                        digitalWrite(minPin + 2, HIGH);
                } else if (inChar == '3'){
                        digitalWrite(minPin + 3, HIGH);
                } else if (inChar == '4'){
                        digitalWrite(minPin + 4, HIGH);
                } else if (inChar == '5'){
                        digitalWrite(minPin + 5, HIGH);
                } else if (inChar == '6'){
                        digitalWrite(minPin + 6, HIGH);
                } else if (inChar == '7'){
                        digitalWrite(minPin + 7, HIGH);
                } else if (inChar == '8'){
                        digitalWrite(minPin + 8, HIGH);
                } else if (inChar == '9'){
                        digitalWrite(minPin + 9, HIGH);
                } else if (inChar == 'a'){
                        digitalWrite(minPin + 10, HIGH);
                } else if (inChar == 'b'){
                        digitalWrite(minPin + 11, HIGH);
                }
        }
}
