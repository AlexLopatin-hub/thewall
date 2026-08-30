CREATE SCHEMA thewall;

CREATE TABLE thewall.posts (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    text VARCHAR(1000) NOT NULL CHECK(
            char_length(text) BETWEEN 1 AND 1000
    ),
    created_at TIMESTAMPTZ NOT NULL
);
