from pad4pi import rpi_gpio
import time

KEYPAD = [
    ["1","2","3","A"],
    ["4","5","6","B"],
    ["7","8","9","C"],
    ["*","0","#","D"]
]
ROW_PINS = [5,6,13,19] # BCM numbering (adjust as needed)
COL_PINS = [12,16,20,21]

factory = rpi_gpio.KeypadFactory()
keypad = factory.create_keypad(keypad=KEYPAD, row_pins=ROW_PINS, col_pins=COL_PINS)

_pin_code = []
_done = False

def _handle_key(key):
    global _done
    if key == "#":
        _done = True
    else:
        _pin_code.append(key)

def get_pin_code():
    global _pin_code, _done
    _pin_code = []
    _done = False
    keypad.registerKeyPressHandler(_handle_key)
    print("Enter PIN, finish with #:")
    while not _done:
        time.sleep(0.1)
    keypad.unregisterKeyPressHandler(_handle_key)
    code = ''.join(_pin_code)
    _pin_code = []
    _done = False
    return code 