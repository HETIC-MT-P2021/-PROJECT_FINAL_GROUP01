drop table round;
drop table game;
drop table player;

create table if not exists player (
    player_id serial primary key,
    name varchar(255) not null,
    avatar varchar(255) not null,
    discord_id varchar(255) not null unique
);

create table if not exists game (
    game_id serial primary key,
    started_at integer,
    finished_at integer
);

create table if not exists round (
    game_id integer references game,
    player_id varchar(255) references player(discord_id),
    reason varchar(255),
    word varchar(255) NOT NULL,
    submitted_at integer
);