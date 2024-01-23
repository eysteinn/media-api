CREATE TABLE IF NOT EXISTS faces (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
	photo_uuid TEXT NOT NULL,
    person_uuid TEXT,
    descriptor BLOB
);
CREATE UNIQUE INDEX idx_uuid ON files (uuid);
CREATE INDEX person_uuid ON files (person_uuid);


