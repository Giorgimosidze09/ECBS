import RPi.GPIO as GPIO
import time

RELAY_PIN = 17  # Adjust as needed
GPIO.setmode(GPIO.BCM)
GPIO.setup(RELAY_PIN, GPIO.OUT)
GPIO.output(RELAY_PIN, GPIO.LOW)

def trigger_relay(duration=2):
    GPIO.output(RELAY_PIN, GPIO.HIGH)
    time.sleep(duration)
    GPIO.output(RELAY_PIN, GPIO.LOW) 