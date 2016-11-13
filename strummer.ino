int string1 = 0;
int string2 = 0;
int string3 = 0;
int string4 = 0;

int triggerVal = 1000;

int string1Pin = 3;
int string2Pin = 2;
int string3Pin = 1;
int string4Pin = 0;

int delayTime = 5;

void setup()

{
  pinMode(string1Pin, INPUT);
  pinMode(string2Pin, INPUT);
  pinMode(string3Pin, INPUT);
  pinMode(string4Pin, INPUT);
  Serial.begin(9600);
}



void loop()

{
  string1 = analogRead(string1Pin);
  delay(delayTime);
  string2 = analogRead(string2Pin);
  delay(delayTime);
  string3 = analogRead(string3Pin);
  delay(delayTime);
  string4 = analogRead(string4Pin);
  delay(delayTime);

  Serial.print("[");
  Serial.print(string1 < triggerVal);
  Serial.print(", ");
  Serial.print(string2 < triggerVal);
  Serial.print(", ");
  Serial.print(string3 < triggerVal);
  Serial.print(", ");
  Serial.print(string4 < triggerVal);
  Serial.print("]");
  Serial.println();

  delay(20);
}
