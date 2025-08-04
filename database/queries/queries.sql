-- name: FindAllCustomers :many
SELECT * FROM customers;

-- name: FindCustomerById :one
SELECT * FROM customers WHERE id = $1;

-- name: InsertCustomer :exec
INSERT INTO customers (id, name, email) values ($1, $2, $3);

-- name: UpdateCustomer :exec
UPDATE customers SET name = $1, email = $2, updated_at = NOW() WHERE id = $3;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1;

-- name: FindAllFavoriteProdutsFromCustomer :many
SELECT * FROM favorites WHERE customer_id = $1;

-- name: InsertFavoriteCustomerProduct :exec
INSERT INTO favorites (customer_id, product_id) VALUES ($1, $2);

-- name: DeleteFavoriteCustomerProduct :exec
DELETE FROM favorites WHERE customer_id = $1 AND product_id = $2;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: FindUserById :one
SELECT * FROM users WHERE id = $1;