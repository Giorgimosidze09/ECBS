# main.py

import db
import time
import platform

if platform.system() == "Darwin":  # MacOS
    from rfid_sim import read_rfid_card, uid_bytes_to_str
    def read_rfid_binary():
        uid_bytes = read_rfid_card()
        return uid_bytes_to_str(uid_bytes)
else:
    from rfid_reader import read_rfid_binary
    import keypad_reader
    import display
    import relay
    import buzzer

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
            card_id = read_rfid_binary()
            print(f"[RFID] Read card_id: {card_id}")
            pin_code = None
        else:
            if platform.system() == "Darwin":
                pin_code = input("Enter PIN code: ").strip()
            else:
                display.show_message("Enter PIN:")
                pin_code = keypad_reader.get_pin_code()
                display.show_message(f"PIN Entered: {pin_code}")
            card_id = None

        if method == "card":
            authorized = db.is_authorized_card(card_id if card_id is not None else "", device_id)
        else:
            authorized = db.is_authorized_pin(pin_code if pin_code is not None else "", device_id)

        if authorized:
            print("Access GRANTED.")
            if platform.system() != "Darwin":
                if method == "card":
                    balance = db.get_card_balance(card_id if card_id is not None else "")
                    display.show_message(f"Access GRANTED\nBalance: {balance:.2f}")
                else:
                    display.show_message("Access GRANTED")
                relay.trigger_relay()
                buzzer.beep_granted()
            result = "granted"
        else:
            print("Access DENIED.")
            if platform.system() != "Darwin":
                display.show_message("Access DENIED")
                buzzer.beep_denied()
            result = "denied"

        db.log_access(method, card_id if card_id is not None else "", pin_code if pin_code is not None else "", device_id, result)
        time.sleep(1)

if __name__ == "__main__":
    simulate_access()
