DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
            DROP TABLE IF NOT EXISTS schema_%s.games', LPAD(i::text, 3, '0'));
            END LOOP;
    END $$;