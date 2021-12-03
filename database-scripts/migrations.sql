CREATE TABLE IF NOT EXISTS player (
  player_id integer PRIMARY KEY,
  name VARCHAR (255) NOT NULL,
  avatar VARCHAR (255) NOT NULL,
  discord_id integer
);

CREATE TABLE IF NOT EXISTS game (
  game_id INTeger primary key,
  started_at date,
  finished_at date
);

CREATE TABLE IF NOT EXISTS round (
   game_id INTeger,
   player_id INTeger,
   reason VARCHAR (255),
   word VARCHAR (255) NOT NULL,
   submitted_at date,
   foreign KEY (game_id) REFERENCES game(game_id),
   FOREIGN KEY (player_id) REFERENCES player(player_id)
);