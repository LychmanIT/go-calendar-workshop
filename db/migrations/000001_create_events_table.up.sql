CREATE TABLE events (
    id VARCHAR ( 100 ) primary key,
    title VARCHAR ( 100 ) NOT NULL,
    description VARCHAR ( 255 ) NOT NULL,
    time VARCHAR ( 50 ) NOT NULL,
    timezone VARCHAR ( 50 ) NOT NULL,
    duration INTEGER
);

CREATE TABLE event_notes (
    id VARCHAR ( 100 ) primary key,
    event_id VARCHAR ( 100 ),
    text VARCHAR ( 255 ) not null,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

