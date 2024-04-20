CREATE TYPE status AS ENUM ('pending','approved','denied');

CREATE TABLE IF NOT EXISTS requests
(
    id           TEXT PRIMARY KEY,
    user_id      TEXT      NOT NULL,
    status       status    NOT NULL DEFAULT 'pending',
    doc_image_id TEXT      NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reasons
(
    request_id TEXT REFERENCES requests (id),
    reason     TEXT NOT NULL
);

