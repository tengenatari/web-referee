DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
        CREATE TABLE IF NOT EXISTS schema_%s.games  (
             id UUID UNIQUE NOT NULL,
             game_url TEXT,
             tour_num INT NOT NULL,
             result_black INT,
             result_white INT,
             white UUID,
             black UUID
        )', LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'));
            END LOOP;
    END $$;
