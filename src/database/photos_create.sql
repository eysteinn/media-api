CREATE TABLE IF NOT EXISTS files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    filename TEXT NOT NULL,
	filepath TEXT NOT NULL,
    unixtime INTEGER,
    sha256 TEXT NOT NULL,
	uuid TEXT NOT NULL
);
CREATE UNIQUE INDEX idx_sha256 ON files (sha256);
CREATE UNIQUE INDEX idx_uuid ON files (uuid);
CREATE  INDEX idx_unixtime ON files (unixtime);

