CREATE TYPE backlog_type AS ENUM ('STORY', 'BUG');

CREATE TABLE backlog_items
(
    id          uuid         NOT NULL,
    type        backlog_type NOT NULL,
    summary     text         NOT NULL DEFAULT '',
    story_point integer
);
