CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  phone VARCHAR(50) NOT NULL UNIQUE,
  address VARCHAR(255) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE validation_codes (
  id SERIAL PRIMARY KEY,
  code VARCHAR(50) NOT NULL,
  email VARCHAR(255) NOT NULL,
  used_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TYPE kind as ENUM ('expenses', 'in_come');

CREATE TABLE items (
    id BIGSERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    amount INTEGER NOT NULL,
    tag_ids INTEGER[] NOT NULL,
    kind kind NOT NULL,
    happened_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

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