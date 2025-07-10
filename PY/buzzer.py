import RPi.GPIO as GPIO
import time

BUZZER_PIN = 27  # Adjust as needed
GPIO.setmode(GPIO.BCM)
GPIO.setup(BUZZER_PIN, GPIO.OUT)
GPIO.output(BUZZER_PIN, GPIO.LOW)

def beep(duration=0.2):
    GPIO.output(BUZZER_PIN, GPIO.HIGH)
    time.sleep(duration)
    GPIO.output(BUZZER_PIN, GPIO.LOW)

def beep_granted():
    beep(0.5)  # Long beep for granted

def beep_denied():
    for _ in range(2):
        beep(0.1)
        time.sleep(0.1) 