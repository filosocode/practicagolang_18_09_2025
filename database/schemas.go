package database

const CreateTableReadings = `
CREATE TABLE IF NOT EXISTS readings (
    id SERIAL PRIMARY KEY,
    device_id TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    consumption DOUBLE PRECISION NOT NULL
);
`
