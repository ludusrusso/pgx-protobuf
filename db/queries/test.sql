-- name: CreateTest :exec
INSERT INTO test_cols (data) VALUES ($1);

-- name: GetTests :many
SELECT * FROM test_cols;