CREATE TABLE sub_tasks
(
    id              uuid PRIMARY KEY,
    summary         text NOT NULL DEFAULT '',
    story_point     integer,
    backlog_item_id uuid NOT NULL,
    CONSTRAINT sub_tasks_backlog_item_id_fk FOREIGN KEY (backlog_item_id) REFERENCES backlog_items (id)
);
