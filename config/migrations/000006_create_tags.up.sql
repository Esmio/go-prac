CREATE TABLE IF NOT EXISTS tags (
  id BIGSERIAL PRIMARY KEY,
  user_id SERIAL NOT NULL,
  name varchar(50) NOT NULL,
  sign varchar(10) NOT NULL,
  kind kind NOT NULL,
  deleteed_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
)