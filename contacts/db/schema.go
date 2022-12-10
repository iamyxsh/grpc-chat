package db

var Schema = `
CREATE TABLE IF NOT EXISTS users (
    number VARCHAR(10) PRIMARY KEY NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS contacts (
    number VARCHAR(10) NOT NULL,
    contact VARCHAR(10) NOT NULL,
    UNIQUE (number, contact)
);
`
