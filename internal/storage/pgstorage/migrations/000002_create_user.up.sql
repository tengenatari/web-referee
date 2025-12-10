DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
            CREATE TABLE IF NOT EXISTS schema_%s.users (
                id SERIAL PRIMARY KEY,
                passwd_hashed_salt VARCHAR(1024) NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                tigr_id VARCHAR(20) NOT NULL,
                rating INT NOT NULL
            )', LPAD(i::text, 3, '0'));
            END LOOP;
    END $$;