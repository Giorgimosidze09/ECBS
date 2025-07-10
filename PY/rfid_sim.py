import json
from datetime import datetime

def read_rfid_card():
    # Simulated UID bytes from MFRC522
    uid_bytes = [0xAB, 0x12, 0x34, 0x56]
    return uid_bytes

def uid_bytes_to_str(uid_bytes):
    return ''.join(f'{byte:02X}' for byte in uid_bytes)

def read_keypad_pin():
    # Simulate user entering PIN on keypad
    pin = input("Enter PIN code (or leave empty to skip): ").strip()
    return pin if pin else None

def build_payload_card(card_id, device_id):
    return {
        "card_id": card_id,
        "pin_code": None,
        "device_id": device_id,
        "method": "card",
        "timestamp": datetime.utcnow().isoformat() + "Z"
    }

def build_payload_pin(pin_code, device_id):
    return {
        "card_id": None,
        "pin_code": pin_code,
        "device_id": device_id,
        "method": "pin",
        "timestamp": datetime.utcnow().isoformat() + "Z"
    }

def main():
    device_id = "elevator-01"

    # First, try keypad PIN input
    pin_code = read_keypad_pin()
    if pin_code:
        payload = build_payload_pin(pin_code, device_id)
    else:
        # If no PIN entered, simulate RFID card read
        uid_bytes = read_rfid_card()
        card_id = uid_bytes_to_str(uid_bytes)
        payload = build_payload_card(card_id, device_id)

    json_payload = json.dumps(payload, indent=2)
    print("Payload to send to backend:")
    print(json_payload)

if __name__ == "__main__":
    main()
