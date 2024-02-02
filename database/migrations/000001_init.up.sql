CREATE TABLE IF NOT EXISTS subjects
(
    id   SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS queue
(
    id          SERIAL PRIMARY KEY,
    title       varchar(255) NOT NULL,
    subject_id  INTEGER,
    is_open     BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS queue_position
(
    id          SERIAL PRIMARY KEY,
    queue_id    INTEGER,
    entered_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_queue
        FOREIGN KEY(queue_id) REFERENCES queue(id)
        ON DELETE CASCADE
);
