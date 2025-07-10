from luma.core.interface.serial import i2c
from luma.oled.device import ssd1306
from luma.core.render import canvas

serial = i2c(port=1, address=0x3C)
device = ssd1306(serial)

def show_message(msg):
    with canvas(device) as draw:
        draw.text((0, 0), msg, fill="white") 