CREATE TABLE IF NOT EXISTS todos (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at DATE,
    completed_at DATE
);

CREATE USER IF NOT EXISTS 'todoUser'@'localhost' IDENTIFIED BY 'SecretPassword';
GRANT ALL PRIVILEGES ON todos.* TO 'todoUser'@'localhost';
