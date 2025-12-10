CREATE TABLE IF NOT EXISTS public.players (
    id SERIAL PRIMARY KEY,
    mac_mahon INT NOT NULL,
    tournament_id INT,
    user_id INT,
    FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id),
    FOREIGN KEY (user_id) REFERENCES public.users(id)
 )