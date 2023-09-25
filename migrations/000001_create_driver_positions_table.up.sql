
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS "driver_positions" (
    id SERIAL PRIMARY KEY,
    "driver_id" integer NOT NULL,
    "location" geometry(Point,4326) NOT NULL,
    "timestamp" timestamp DEFAULT CURRENT_TIMESTAMP
);

-- Create a spatial index for efficient proximity queries
CREATE INDEX IF NOT EXISTS idx_driver_positions_location ON driver_positions USING GIST (location);
