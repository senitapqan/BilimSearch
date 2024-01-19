SELECT 
    li.id AS lesson_item_id,
    li.date AS lesson_date,
    t.id As Task
FROM t_lesson_items li JOIN t_task t ON li.id = t.lesson_item_id
WHERE li.lesson_id = 6
ORDER BY li.id;
