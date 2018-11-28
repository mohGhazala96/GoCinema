-- DROP TABLE users;
-- DROP TABLE reservations;
-- DROP TABLE movies;
-- DROP TABLE halls;
-- DROP TABLE cinemas;

CREATE TABLE cinemas(
    id Serial PRIMARY KEY,
    cinema VARCHAR  NOT NULL UNIQUE,
    halls NUMERIC
);
CREATE TABLE halls(
    id Serial PRIMARY KEY,
    cinema VARCHAR,
    seats NUMERIC,
    emptyseats VARCHAR ,
    reservedseats VARCHAR,
    FOREIGN KEY (cinema) REFERENCES cinemas(cinema) ON DELETE CASCADE
);
CREATE TABLE movies(
    id Serial PRIMARY KEY,
    movie VARCHAR NOT NULL UNIQUE,
    cinema VARCHAR,
    hall integer,
    timing timestamp NOT NULL UNIQUE,
    FOREIGN KEY (cinema) REFERENCES cinemas(cinema) ON DELETE CASCADE,
    FOREIGN KEY (hall) REFERENCES halls(id) ON DELETE CASCADE

);

CREATE TABLE reservations(
    id Serial PRIMARY KEY,
    username VARCHAR,
    hall integer,
    cinema VARCHAR,
    seat VARCHAR,
    movie varchar,
    timing timestamp NOT NULL UNIQUE,
    FOREIGN KEY (hall) REFERENCES halls(id),
    FOREIGN KEY (cinema) REFERENCES cinemas(cinema),
    FOREIGN KEY (movie) REFERENCES movies(movie),
    FOREIGN KEY (timing) REFERENCES movies(timing)

);

CREATE TABLE users(
    id Serial PRIMARY KEY,
    username VARCHAR,
    email VARCHAR,
    phone NUMERIC,
    reservations integer,
    FOREIGN KEY (reservations) REFERENCES reservations(id)
);



Insert INTO users("username") VALUES 
('hossam');