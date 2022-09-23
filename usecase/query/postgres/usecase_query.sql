-- name: GetBacklogItems :many
SELECT *
FROM backlog_items bi
         LEFT JOIN sub_tasks st
                   ON st.backlog_item_id = bi.id LIMIT 100;
