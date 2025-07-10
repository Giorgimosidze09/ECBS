import sqlite3
from datetime import datetime

DB_PATH = "elevator.db"

def get_connection():
    conn = sqlite3.connect(DB_PATH)
    conn.row_factory = sqlite3.Row  # So fetchall returns dict-like rows
    return conn

def is_authorized_card(card_id: str, device_id: str) -> bool:
    """Check if a card_id is authorized for the device and not expired."""
    with get_connection() as conn:
        cur = conn.cursor()
        cur.execute("""
            SELECT id FROM authorized_access
            WHERE card_id = ? AND device_id = ? AND active = 1
              AND datetime(expires_at) > datetime('now')
        """, (card_id, device_id))
        return cur.fetchone() is not None

def is_authorized_pin(pin_code: str, device_id: str) -> bool:
    """Check if a PIN code is authorized for the device and not expired."""
    with get_connection() as conn:
        cur = conn.cursor()
        cur.execute("""
            SELECT id FROM authorized_access
            WHERE pin_code = ? AND device_id = ? AND active = 1
              AND datetime(expires_at) > datetime('now')
        """, (pin_code, device_id))
        return cur.fetchone() is not None

def log_access(method: str, card_id: str | None, pin_code: str | None, device_id: str, result: str) -> None:
    """Log access attempts with their result."""
    with get_connection() as conn:
        cur = conn.cursor()
        cur.execute("""
            INSERT INTO access_logs (method, card_id, pin_code, device_id, timestamp, result, synced)
            VALUES (?, ?, ?, ?, ?, ?, 0)
        """, (method, card_id, pin_code, device_id, datetime.utcnow().isoformat(), result))
        conn.commit()

def get_unsynced_logs() -> list[dict]:
    """Return list of access logs which have not been synced yet."""
    with get_connection() as conn:
        cur = conn.cursor()
        cur.execute("""
            SELECT * FROM access_logs WHERE synced = 0
        """)
        return [dict(row) for row in cur.fetchall()]

def mark_log_synced(log_id: int) -> None:
    """Mark a given access log entry as synced."""
    with get_connection() as conn:
        cur = conn.cursor()
        cur.execute("UPDATE access_logs SET synced = 1 WHERE id = ?", (log_id,))
        conn.commit()

def replace_authorized_list(new_list: list[dict]) -> None:
    """Replace the entire authorized_access table with new data."""
    with get_connection() as conn:
        cur = conn.cursor()
        cur.execute("DELETE FROM authorized_access")
        for entry in new_list:
            cur.execute("""
                INSERT INTO authorized_access (card_id, pin_code, user_id, device_id, expires_at, type, active)
                VALUES (?, ?, ?, ?, ?, ?, ?)
            """, (
                entry.get("card_id"),
                entry.get("pin_code"),
                entry.get("user_id"),
                entry.get("device_id"),
                entry.get("expires_at"),
                entry.get("type"),
                entry.get("active", 1),
            ))
        conn.commit()
