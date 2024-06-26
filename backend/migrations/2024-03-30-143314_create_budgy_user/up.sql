-- Your SQL goes here

CREATE TABLE budgy_user
(
    budgy_user_id SERIAL PRIMARY KEY,
    username      VARCHAR UNIQUE           NOT NULL,
    password      VARCHAR                  NOT NULL,
    email         VARCHAR UNIQUE           NOT NULL,
    deleted       BOOLEAN                  NOT NULL DEFAULT FALSE,
    inserted_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
)
