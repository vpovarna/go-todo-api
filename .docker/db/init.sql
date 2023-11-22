CREATE TABLE IF NOT EXISTS todos (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at DATETIME,
    completed_at DATETIME
);

CREATE USER IF NOT EXISTS 'todoUser'@'%' IDENTIFIED BY 'SecretPassword';
GRANT ALL PRIVILEGES ON todosDB.* TO 'todoUser'@'%';
FLUSH PRIVILEGES;
