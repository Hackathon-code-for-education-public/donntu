CREATE TYPE roles AS ENUM ('applicant', 'student', 'manager');

CREATE TABLE IF NOT EXISTS credentials
(
    id                TEXT PRIMARY KEY,
    role              roles           NOT NULL DEFAULT 'applicant',
    email             TEXT UNIQUE,
    password          TEXT            NOT NULL,
    created_at        TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP
);
