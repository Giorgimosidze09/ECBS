# sync.py

import requests
import db
import config
import time

def sync_logs():
    logs = db.get_unsynced_logs()
    for log in logs:
        try:
            response = requests.post(f"{config.API_BASE_URL}/api/sync-log", json=log, timeout=5)
            if response.status_code == 200:
                db.mark_log_synced(log["id"])
                print(f"[SYNC] Log {log['id']} synced.")
            else:
                print(f"[ERROR] Sync failed for log {log['id']}: {response.status_code}")
        except Exception as e:
            print(f"[ERROR] Exception syncing log {log['id']}: {e}")

def sync_authorized_access():
    try:
        res = requests.get(f"{config.API_BASE_URL}/api/device-sync", params={"device_id": config.DEVICE_ID}, timeout=5)
        if res.status_code == 200:
            data = res.json()
            db.replace_authorized_list(data)
            print(f"[SYNC] Authorized list updated: {len(data)} records.")
        else:
            print(f"[ERROR] Fetching authorized access failed: {res.status_code}")
    except Exception as e:
        print(f"[ERROR] Exception while syncing authorized access: {e}")

def run_sync_loop():
    while True:
        print("[SYNC] Starting sync...")
        sync_logs()
        sync_authorized_access()
        print(f"[SYNC] Done. Waiting {config.SYNC_INTERVAL_SECONDS}s.\n")
        time.sleep(config.SYNC_INTERVAL_SECONDS)

# Optional: for standalone test
if __name__ == "__main__":
    run_sync_loop()
