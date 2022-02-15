CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50),
    email VARCHAR(100) UNIQUE,
    register_date DATETIME,
    password VARCHAR(128)
);