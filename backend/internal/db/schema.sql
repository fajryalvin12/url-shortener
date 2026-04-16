CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);