CREATE DATABASE IF NOT EXISTS go_todo;

USE go_todo;

CREATE TABLE IF NOT EXISTS todo
(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(1000) NOT NULL
);

INSERT INTO todo (name, description)
VALUES ('Supermarket', 'milk, bread, egg, pasta, meat, cookies and fruit'),
    ('Meeting - work team', 'today, at 5 pm'),
    ('Football game', 'Friday, at 7 pm');
