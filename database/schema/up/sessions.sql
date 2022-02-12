CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid VARCHAR(128),
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES (user_id) users
);
