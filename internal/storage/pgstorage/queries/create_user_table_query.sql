CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    passwd_hashed_salt VARCHAR(1024) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tigr_id VARCHAR(20) NOT NULL,
    rating INT NOT NULL
)