-- DROP TABLE reservations;
-- DROP TABLE halls;
-- DROP TABLE timings;
-- DROP TABLE movies;


CREATE TABLE halls(
    id Serial PRIMARY KEY,
    seats NUMERIC,
    emptyseats VARCHAR ,
    reservedseats VARCHAR,
    movie VARCHAR,
    FOREIGN KEY (movie) REFERENCES movies(movie) ON DELETE CASCADE
);
CREATE TABLE movies(
    id Serial PRIMARY KEY,
    movie VARCHAR NOT NULL UNIQUE,
    timing timestamp NOT NULL UNIQUE
    --poster 
);
CREATE TABLE timings(
    id Serial PRIMARY KEY,
    movie_period VARCHAR UNIQUE,
    movie VARCHAR,
    FOREIGN KEY (movie) REFERENCES movies(movie)
);

CREATE TABLE reservations(
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