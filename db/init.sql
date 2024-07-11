CREATE TABLE user (
    user_id VARCHAR(36) BINARY NOT NULL PRIMARY KEY,
    username VARCHAR(100) BINARY NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE video (
    video_id VARCHAR(36) BINARY NOT NULL PRIMARY KEY,
    title VARCHAR(255) BINARY NOT NULL,
    description TEXT,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE genre (
    genre_id VARCHAR(36) BINARY NOT NULL,
    genre_name VARCHAR(20) BINARY NOT NULL
);

CREATE TABLE video_genre (
    video_id VARCHAR(36) BINARY NOT NULL,
    genre_id VARCHAR(36) BINARY NOT NULL
);

CREATE TABLE user_interest (
   user_id VARCHAR(36) BINARY NOT NULL,
   genre_id VARCHAR(36) BINARY NOT NULL,
   weight INT NOT NULL
);

CREATE TABLE user_no_interest (
    user_id VARCHAR(36) BINARY NOT NULL,
    genre_id VARCHAR(36) BINARY NOT NULL,
    weight INT NOT NULL
);

CREATE TABLE watch_history (
    user_id VARCHAR(36) BINARY NOT NULL,
    video_id VARCHAR(36) BINARY NOT NULL,
    watched_at DATETIME(3) NOT NULL
);

