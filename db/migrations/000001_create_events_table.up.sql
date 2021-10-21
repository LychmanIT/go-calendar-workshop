CREATE TABLE events (
    id SERIAL primary key,
    title VARCHAR ( 100 ) NOT NULL,
    description VARCHAR ( 255 ) NOT NULL,
    time VARCHAR ( 50 ) NOT NULL,
    timezone VARCHAR ( 50 ) NOT NULL,
    duration INTEGER
);

CREATE TABLE event_notes (
    id SERIAL primary key,
    event_id INTEGER,
    text TEXT not null,
    FOREIGN KEY (event_id) REFERENCES events(id)
);

