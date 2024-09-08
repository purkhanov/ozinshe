CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    is_admin BOOLEAN DEFAULT FALSE,
    phone_number VARCHAR(50),
    year_of_birth DATE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    director VARCHAR(255),
    producer VARCHAR(255),
    runtime SMALLINT NOT NULL,
    year SMALLINT NOT NULL,
    stars SMALLINT,
    series SMALLINT,
    seasons SMALLINT, 
    image VARCHAR(255),
    video_url VARCHAR(255)
);

CREATE TABLE screenshots (
    id SERIAL PRIMARY KEY,
    link VARCHAR(255) NOT NULL UNIQUE,
    movie_id INTEGER NOT NULL,
    CONSTRAINT fk_movie
        FOREIGN KEY (movie_id)
        REFERENCES movies(id)
        ON DELETE CASCADE
);

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    genre VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE movies_genres_assoc (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE NOT NULL,
    genre_id INTEGER REFERENCES genres(id) ON DELETE CASCADE NOT NULL,
    UNIQUE(movie_id, genre_id)
);

CREATE TABLE favorite_movies (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE NOT NULL,
    UNIQUE(user_id, movie_id)
);

CREATE TABLE watched_movies (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE NOT NULL,
    watched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, movie_id)
);
