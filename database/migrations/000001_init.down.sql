-- Drop foreign key constraint before dropping the table to avoid reference errors
ALTER TABLE IF EXISTS queue_position
    DROP CONSTRAINT IF EXISTS fk_queue;

-- Drop the tables in reverse order
DROP TABLE IF EXISTS queue_position;
DROP TABLE IF EXISTS queue;
DROP TABLE IF EXISTS subjects;
