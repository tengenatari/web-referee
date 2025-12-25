DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
            CREATE TABLE IF NOT EXISTS schema_%s.tournaments (
                id UUID UNIQUE NOT NULL,
                name VARCHAR(100) NOT NULL,
                date DATE,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            )', LPAD(i::text, 3, '0'));
            END LOOP;
    END $$;