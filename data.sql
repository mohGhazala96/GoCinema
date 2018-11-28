-- DROP TABLE cinemas;
-- DROP TABLE movies;
-- DROP TABLE halls;
-- DROP TABLE users;
-- DROP TABLE reservations;

CREATE TABLE cinemas(
    id Serial PRIMARY KEY,
    cinema VARCHAR,
    halls NUMERIC
);

CREATE TABLE movies(
    id Serial PRIMARY KEY,
    movie VARCHAR,
    cinema VARCHAR ARRAY[100]
    -- FOREIGN KEY (cinema) REFERENCES cinemas(cinema)
);
CREATE TABLE halls(
    id Serial PRIMARY KEY,
    cinema VARCHAR,
    seats NUMERIC,
    emptyseats VARCHAR ARRAY[100],
    reservedseats VARCHAR ARRAY[100],
    movie VARCHAR
    -- FOREIGN KEY (cinema) REFERENCES cinemas(cinema),
    -- FOREIGN KEY (movie) REFERENCES movies(movie)
);

CREATE TABLE users(
    id Serial PRIMARY KEY,
    username VARCHAR,
    email VARCHAR,
    phone NUMERIC
    -- reservations varchar(100) ARRAY[100],
    -- FOREIGN KEY (reservations) REFERENCES reservations(id)
);
CREATE TABLE reservations(
    id Serial PRIMARY KEY,
    username VARCHAR,
    hall NUMERIC,
    cinema VARCHAR,
    seat VARCHAR,
    movie varchar
    -- FOREIGN KEY (hall) REFERENCES halls(id),
    -- FOREIGN KEY (cinema) REFERENCES cinemas(cinema),
    -- FOREIGN KEY (username) REFERENCES users(username)
);



Insert INTO users("username") VALUES 
('hossam');