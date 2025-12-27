DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
            CREATE TABLE IF NOT EXISTS schema_%s.players (
                id UUID UNIQUE NOT NULL,
                mac_mahon INT NOT NULL,
                tournament_id UUID,
                user_id UUID
             )', LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'));
        END LOOP;
    END $$;
