package db

var Schema = `
CREATE TABLE IF NOT EXISTS users (
    number VARCHAR(10) PRIMARY KEY NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS messages (
    sender VARCHAR(10) NOT NULL,
    receiver VARCHAR(10) NOT NULL,
    body TEXT,
    delivered BOOLEAN,
    timestamp TIMESTAMPTZ DEFAULT Now() 
);
`
