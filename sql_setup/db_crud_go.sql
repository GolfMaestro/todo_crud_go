CREATE TABLE persons (
    id        SERIAL PRIMARY KEY,
    name      VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL
);

CREATE TABLE tasks (
    id          SERIAL PRIMARY KEY,
    person_id   INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    title       VARCHAR(255) NOT NULL,
    is_complete BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
