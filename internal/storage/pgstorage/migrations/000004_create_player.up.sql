DO $$
    DECLARE
        i INTEGER;
    BEGIN
        FOR i IN 1..512 LOOP
                EXECUTE format('
            CREATE TABLE IF NOT EXISTS schema_%s.players (
                id SERIAL PRIMARY KEY,
                mac_mahon INT NOT NULL,
                tournament_id INT,
                user_id INT,
                FOREIGN KEY (tournament_id) REFERENCES schema_%s.tournaments(id),
                FOREIGN KEY (user_id) REFERENCES schema_%s.users(id)
             )', LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'), LPAD(i::text, 3, '0'));
        END LOOP;
    END $$;
