BEGIN;
CREATE TYPE kind as ENUM ('expenses', 'in_come', '');
ALTER TABLE items DROP COLUMN kind;
ALTER TABLE items DROP COLUMN kind kind NOT NULL DEFAULT 'expenses';
ALTER TABLE tags DROP COLUMN kind;
ALTER TABLE tags DROP COLUMN kind kind NOT NULL DEFAULT 'expenses';
COMMIT;