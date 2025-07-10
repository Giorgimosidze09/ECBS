# main.py

import db
import time
from rfid_reader import read_rfid_binary  # Use real RFID reader

def simulate_access():
    device_id = "elevator-01"

    print("Elevator access simulation. Type 'exit' to quit.")
    while True:
        method = input("Enter input method (card/pin): ").strip().lower()
        if method == "exit":
            break
        if method not in ("card", "pin"):
            print("Invalid method. Use 'card' or 'pin'.")
            continue

        if method == "card":
            card_id = read_rfid_binary()  # Read from real RFID hardware
            print(f"[RFID] Read card_id: {card_id}")
            pin_code = None
            authorized = db.is_authorized_card(card_id, device_id)
        else:
            pin_code = input("Enter PIN code: ").strip()
            card_id = None
            authorized = db.is_authorized_pin(pin_code, device_id)

        if authorized:
            print("Access GRANTED.")
            result = "granted"
        else:
            print("Access DENIED.")
            result = "denied"

        # Log the access attempt
        db.log_access(method, card_id, pin_code, device_id, result)

        time.sleep(1)  # simulate delay

if __name__ == "__main__":
    simulate_access()
