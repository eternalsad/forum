CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author_id INTEGER,
    content TEXT,
    title TEXT,
    creation_date DATETIME
);