CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id         INT auto_increment PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    nick       VARCHAR(50) NOT NULL UNIQUE,
    email      VARCHAR(50) NOT NULL UNIQUE,
    pass       VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
)
    engine=innodb;

CREATE TABLE followers
(
    user_id     INT NOT NULL,
    follower_id INT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, follower_id)
)
    engine=innodb;

CREATE TABLE posts
(
    id         INT auto_increment PRIMARY KEY,
    title      VARCHAR(50) NOT NULL,
    content    VARCHAR(300) NOT NULL,
    autor_id   INT NOT NULL,
    likes      INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (autor_id) REFERENCES users (id) ON DELETE CASCADE
)
    engine=innodb;