import requests

def topup(card_id, amount):
    url = "http://your-backend-url/api/paybox/topup"
    data = {"card_id": card_id, "amount": amount}
    response = requests.post(url, json=data)
    if response.status_code == 200:
        print("Top-up successful!")
        print(response.json())
    else:
        print("Top-up failed:", response.text)

if __name__ == "__main__":
    card_id = input("Enter your card ID: ").strip()
    amount = float(input("Enter amount to top up: "))
    topup(card_id, amount)