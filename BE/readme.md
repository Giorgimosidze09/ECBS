-- Create User
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Giorgi", "email":"giorgi@example.com"}'

-- Create / Assign Card
curl -X POST http://localhost:8080/cards/assign \
  -H "Content-Type: application/json" \
  -d '{"card_id":"ABC123XYZ", "user_id":1}'

-- Top Up Balance
curl -X POST http://localhost:8080/balances/topup \
  -H "Content-Type: application/json" \
  -d '{"user_id":1, "card_id":1, "balance":20.00}'

-- Validate Card
curl -X POST http://localhost:8080/cards/validate \
  -H "Content-Type: application/json" \
  -d '{"card_id":"ABC123XYZ"}'