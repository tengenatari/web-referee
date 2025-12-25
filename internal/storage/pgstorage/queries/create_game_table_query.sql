CREATE TABLE IF NOT EXISTS public.games (
     id UUID UNIQUE NOT NULL,
     game_url TEXT NOT NULL,
     tour_num INT NOT NULL ,
     result_black INT,
     result_white INT,
     white INT,
     black INT,
     FOREIGN KEY (white) REFERENCES public.player(id),
     FOREIGN KEY (black) REFERENCES public.player(id)
)