-- DROP TABLE reservations;
-- DROP TABLE halls;
-- DROP TABLE timings;
-- DROP TABLE movies;


 CREATE TABLE IF NOT EXISTS halls (
    id NUMERIC PRIMARY KEY,
    seats NUMERIC,
    emptyseats VARCHAR ,
    reservedseats VARCHAR,
    movie VARCHAR,
    FOREIGN KEY (movie) REFERENCES movies(movie) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS movies(
    id Serial PRIMARY KEY,
    title VARCHAR NOT NULL UNIQUE
    --poster 
);
CREATE TABLE IF NOT EXISTS seats(
    seat_status boolean,
    seat_title VARCHAR,
    hall NUMERIC,
    movie_period timestamp,
    PRIMARY KEY(hall,seat_title,movie_period),
    FOREIGN KEY (hall) REFERENCES halls(id),
    FOREIGN KEY (movie_period) REFERENCES timings(movie_period)

)
CREATE TABLE IF NOT EXISTS timings(
    id Serial PRIMARY KEY,
    movie_period timestamp UNIQUE,
    movie_id VARCHAR,
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);

CREATE TABLE IF NOT EXISTS reservations(
    id Serial PRIMARY KEY,
    hall integer,
    seat VARCHAR,
    movie varchar,
    useremail varchar,
    timing timestamp NOT NULL UNIQUE,
    FOREIGN KEY (hall) REFERENCES halls(id),
    FOREIGN KEY (movie) REFERENCES movies(movie),
    FOREIGN KEY (timing) REFERENCES timings(movie_period)

);




-- Insert INTO users("username") VALUES 
-- ('hossam');