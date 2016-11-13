int delayTime = 100;

void setup() {
        pinMode(A0, INPUT);
        pinMode(A1, INPUT);
        pinMode(A2, INPUT);

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
