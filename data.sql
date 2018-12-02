-- DROP TABLE reservations;
-- DROP TABLE halls;
-- DROP TABLE timings;
-- DROP TABLE movies;
CREATE TABLE IF NOT EXISTS movies(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL ,
    release_date TEXT,
    poster_path TEXT,
    vote_average FLOAT,
    overview Text,
    isAvialabe Boolean
);

CREATE TABLE IF NOT EXISTS timings(
    id  Serial PRIMARY Key  ,
    movie_period VARCHAR UNIQUE,
    movie_id integer,
    FOREIGN KEY (movie_id) REFERENCES movies(id),
);

CREATE TABLE IF NOT EXISTS halls (
    id NUMERIC PRIMARY KEY,
    seats NUMERIC,
    movie integer,
    FOREIGN KEY (movie) REFERENCES movies(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS reservations(
    id Serial PRIMARY KEY,
    hall integer,
    seat VARCHAR,
    movie integer,
    useremail varchar,
    day DATE,
    timing integer,
    FOREIGN KEY (hall) REFERENCES halls(id),
    FOREIGN KEY (movie) REFERENCES movies(id)
    -- FOREIGN KEY (timing) REFERENCES timings(id)

);
