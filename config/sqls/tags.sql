-- name: CreateTag :one
INSERT INTO tags (
  user_id, 
  name,
  sign,
  kind
) VALUES (
  $1, 
  $2, 
  $3, 
  $4
) RETURNING *;

-- name: UpdateTag :one
UPDATE tags
SET
  user_id = @user_id,
  name = CASE WHEN @name::varchar = '' THEN name ELSE @name END,
  sign = CASE WHEN @sign::varchar = '' THEN sign ELSE @sign END,
  kind = CASE WHEN @kind::varchar = '' THEN kind ELSE @kind END
WHERE id = @id
RETURNING *;