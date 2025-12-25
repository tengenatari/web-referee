CREATE TABLE IF NOT EXISTS public.users (
    id UUID UNIQUE  NOT NULL,
    name VARCHAR(100) NOT NULL,
    passwd_hashed_salt VARCHAR(1024),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tigr_id VARCHAR(20) NOT NULL,
    rating INT NOT NULL
)