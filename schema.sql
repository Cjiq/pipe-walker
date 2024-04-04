CREATE TABLE IF NOT EXISTS nodes (
    id INTEGER PRIMARY KEY,
    name TEXT,
    node_id INTEGER,
    n_id INTEGER
);

CREATE TABLE IF NOT EXISTS lines (
    id INTEGER PRIMARY KEY,
    name TEXT,
    type TEXT,
    n_id INTEGER,
    FOREIGN KEY (n_id) REFERENCES nodes (n_id)
);
