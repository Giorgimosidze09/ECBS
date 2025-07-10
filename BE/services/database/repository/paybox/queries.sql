-- name: CheckPayboxTransactionExists :one
SELECT EXISTS (
  SELECT 1 FROM paybox_transactions WHERE transaction_id = $1
);

-- name: CreatePayboxTransaction :exec
INSERT INTO paybox_transactions (transaction_id, card_id, amount, source, created_at)
VALUES ($1, $2, $3, $4, $5); 