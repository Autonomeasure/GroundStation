#include <SoftwareSerial.h>

SoftwareSerial radio(2, 3);

void setup() {
  Serial.begin(9600);
  radio.begin(9600);
}

void loop() {
  // if (radio.available() > 0) {
    String input = radio.readString();
    Serial.println(input);
  // }
  delay(20);
}
