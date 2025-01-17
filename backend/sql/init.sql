CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(355) NOT NULL,
    password varchar(450) NOT NULL,
    username varchar(355) NOT NULL UNIQUE,
    isAdmin BOOLEAN DEFAULT false
)
;

CREATE TABLE IF NOT EXISTS reviews (
    id serial PRIMARY KEY,
    userid serial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    isActive BOOLEAN DEFAULT true
)
;

CREATE TABLE IF NOT EXISTS reviews_feedback (
    id serial PRIMARY KEY,
    reviewid serial NOT NULL REFERENCES reviews(id) ON DELETE CASCADE,
    reviewerid serial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT
)
;