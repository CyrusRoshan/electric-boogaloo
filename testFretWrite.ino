int delayTime = 100;
int minPin = 2;
int maxPin = 13;

void setup() {
        for (int i = minPin; i <= maxPin; i++) {
                pinMode(i, OUTPUT);
        }

        Serial.begin(9600);
}



void loop() {
        for (int i = minPin; i <= maxPin; i++) {
                digitalWrite(i, HIGH);
                delay(delayTime);
                digitalWrite(i, LOW);
        }
}
