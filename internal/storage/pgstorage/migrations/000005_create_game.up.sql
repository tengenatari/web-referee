DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
        CREATE TABLE IF NOT EXISTS schema_%s.games  (
             id SERIAL PRIMARY KEY,
             game_url TEXT NOT NULL,
             tour_num INT NOT NULL ,
             result_black INT,
             result_white INT,
             white INT,
             black INT
        )', LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'));
            END LOOP;
    END $$;
