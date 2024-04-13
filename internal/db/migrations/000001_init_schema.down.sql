ALTER TABLE "reading_questions" DROP CONSTRAINT IF EXISTS reading_questions_reading_id_fkey;
ALTER TABLE "question_variants" DROP CONSTRAINT IF EXISTS question_variants_question_id_fkey;
ALTER TABLE "contests" DROP CONSTRAINT IF EXISTS contests_reading_id_fkey;
ALTER TABLE "contests" DROP CONSTRAINT IF EXISTS contests_diktant_id_fkey;
ALTER TABLE "contests" DROP CONSTRAINT IF EXISTS contests_speeches_id_fkey;

-- Drop all tables, starting with the ones that depend on others
DROP TABLE IF EXISTS "contests";
DROP TABLE IF EXISTS "question_variants";
DROP TABLE IF EXISTS "reading_questions";
DROP TABLE IF EXISTS "speeches";
DROP TABLE IF EXISTS "diktants";
DROP TABLE IF EXISTS "readings";
DROP TABLE IF EXISTS "users";
