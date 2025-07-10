from pirc522 import RFID

def read_rfid_binary():
    rdr = RFID()
    print("Place your card near the reader...")
    while True:
        (error, data) = rdr.request()
        if not error:
            (error, uid) = rdr.anticoll()
            if not error:
                card_id = ''.join(f'{x:02X}' for x in uid)
                print(f"Card detected! UID: {card_id}")
                return card_id 