DROP TABLE IF EXISTS attendances;

ALTER TABLE IF EXISTS attendances 
    DROP CONSTRAINT IF EXISTS attendances_fk;