

-- name: CreateAccount :one
INSERT INTO accounts (
  owner, 
  balance, 
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY name
LIMIT $1  
OFFSET $2;

-- name: UpodateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE from accounts WHERE id = $1;
