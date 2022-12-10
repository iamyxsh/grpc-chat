package db

var Schema = `
CREATE TABLE IF NOT EXISTS users (
    number VARCHAR(10) PRIMARY KEY NOT NULL UNIQUE
);`
