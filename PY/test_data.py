# test_data.py

import sqlite3
from datetime import datetime, timedelta

DB_PATH = "elevator.db"

def insert_sample_data():
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()

    # Clear existing data
    cur.execute("DELETE FROM authorized_access")

    # Insert sample authorized cards and pins
    now = datetime.utcnow()
    future = now + timedelta(days=30)

    samples = [
        # Authorized card
        ("AB123456", None, 1, "elevator-01", future.isoformat(), "card", 1),
        # Authorized PIN
        (None, "1234", 2, "elevator-01", future.isoformat(), "pin", 1),
    ]

    for card_id, pin_code, user_id, device_id, expires_at, _type, active in samples:
        cur.execute("""
            INSERT INTO authorized_access 
            (card_id, pin_code, user_id, device_id, expires_at, type, active)
            VALUES (?, ?, ?, ?, ?, ?, ?)
        """, (card_id, pin_code, user_id, device_id, expires_at, _type, active))

    conn.commit()
    conn.close()
    print("Sample data inserted.")

if __name__ == "__main__":
    insert_sample_data()
