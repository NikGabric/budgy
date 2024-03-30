-- Your SQL goes here

CREATE TABLE transaction_type (
  transaction_type_id SERIAL PRIMARY KEY,
  title VARCHAR NOT NULL,
  description VARCHAR,
  deleted BOOLEAN NOT NULL DEFAULT FALSE,
  inserted_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  budgy_user_id INT NOT NULL REFERENCES budgy_user(budgy_user_id)
)
