CREATE TABLE IF NOT EXISTS students (
    user_id TEXT NOT NULL PRIMARY KEY,
    university_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reviews (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES students(user_id) ON DELETE CASCADE,
    university_id TEXT NOT NULL,
    body TEXT NOT NULL,
    rating INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
)