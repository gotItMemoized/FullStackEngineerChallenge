CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(355) NOT NULL,
    password varchar(450) NOT NULL,
    username varchar(355) NOT NULL,
    isAdmin BOOLEAN DEFAULT false
)
;