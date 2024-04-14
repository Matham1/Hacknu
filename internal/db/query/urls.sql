-- name: CreateUser :one
INSERT INTO users (
  full_name,
  username,
  password
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: CreateReading :one
INSERT INTO readings (
  text
) VALUES (
  $1
) RETURNING *;

-- name: CreateReadingQuestion :one
INSERT INTO reading_questions (
  text,
  question,
  reading_id
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: CreateQuestionVariant :one
INSERT INTO question_variants (
  question_id,
  option,
  is_correct,
  explanation
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

-- name: CreateDiktant :one
INSERT INTO diktants (
  text
) VALUES (
  $1
) RETURNING *;

-- name: CreateSpeech :one
INSERT INTO speeches (
  text
) VALUES (
  $1
) RETURNING *;

-- name: CreateContest :one
INSERT INTO contests (
  reading_id,
  diktant_id,
  speeches_id,
  start_time,
  end_time
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: GetContest :one
SELECT
  c.id AS contest_id,
  c.start_time,
  c.end_time,
  r.id AS reading_id,
  r.text AS reading_text,
  d.id AS diktant_id,
  d.text AS diktant_text,
  s.id AS speech_id,
  s.text AS speech_text,
  json_agg(
    json_build_object(
      'text', rq.text,
      'question', rq.question,
      'variants', (
        SELECT json_agg(
          json_build_object(
            'option', qv.option,
            'is_correct', qv.is_correct,
            'explanation', qv.explanation
          )
        )
        FROM question_variants qv
        WHERE qv.question_id = rq.id
      )
    )
  ) AS questions
FROM contests c
JOIN readings r ON c.reading_id = r.id
JOIN diktants d ON c.diktant_id = d.id
JOIN speeches s ON c.speeches_id = s.id
JOIN reading_questions rq ON rq.reading_id = r.id
WHERE c.id = $1  -- Replace $1 with the specific contest ID you're querying for
GROUP BY c.id, c.start_time, c.end_time, r.id, r.text, d.id, d.text, s.id, s.text;

-- name: GetLastTwoContests :many
SELECT
  c.id AS contest_id,
  c.start_time,
  c.end_time,
  r.id AS reading_id,
  r.text AS reading_text,
  d.id AS diktant_id,
  d.text AS diktant_text,
  s.id AS speech_id,
  s.text AS speech_text,
  json_agg(
    json_build_object(
      'text', rq.text,
      'question', rq.question,
      'variants', (
        SELECT json_agg(
          json_build_object(
            'option', qv.option,
            'is_correct', qv.is_correct,
            'explanation', qv.explanation
          )
        )
        FROM question_variants qv
        WHERE qv.question_id = rq.id
      )
    )
  ) AS questions
FROM contests c
JOIN readings r ON c.reading_id = r.id
JOIN diktants d ON c.diktant_id = d.id
JOIN speeches s ON c.speeches_id = s.id
JOIN reading_questions rq ON rq.reading_id = r.id
GROUP BY c.id, c.start_time, c.end_time, r.id, r.text, d.id, d.text, s.id, s.text
ORDER BY c.start_time DESC
LIMIT 2;






-- CREATE TABLE "users" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "full_name" varchar,
--   "username" varchar UNIQUE NOT NULL,
--   "password" varchar NOT NULL,
--   "created_at" timestamptz DEFAULT (now())
-- );

-- CREATE TABLE "readings" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "text" varchar,
--   "created_at" timestamptz DEFAULT (now())
-- );

-- CREATE TABLE "reading_questions" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "text" varchar NOT NULL,
--   "question" varchar NOT NULL,
--   "reading_id" bigint,
--   "created_at" timestamptz DEFAULT (now())
-- );

-- CREATE TABLE "question_variants" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "question_id" bigint,
--   "option" varchar NOT NULL,
--   "is_correct" bool NOT NULL,
--   "explanation" varchar NOT NULL,
--   "created_at" timestamptz DEFAULT (now())
-- );

-- CREATE TABLE "diktants" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "text" varchar
-- );

-- CREATE TABLE "speeches" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "text" varchar
-- );

-- CREATE TABLE "contests" (
--   "reading_id" bigint,
--   "diktant_id" bigint,
--   "speeches_id" bigint,
--   "start_time" timestamp NOT NULL,
--   "end_time" timestamp NOT NULL
-- );

-- ALTER TABLE "reading_questions" ADD FOREIGN KEY ("reading_id") REFERENCES "readings" ("id");

-- ALTER TABLE "question_variants" ADD FOREIGN KEY ("question_id") REFERENCES "reading_questions" ("id");

-- ALTER TABLE "contests" ADD FOREIGN KEY ("reading_id") REFERENCES "readings" ("id");

-- ALTER TABLE "contests" ADD FOREIGN KEY ("diktant_id") REFERENCES "diktants" ("id");

-- ALTER TABLE "contests" ADD FOREIGN KEY ("speeches_id") REFERENCES "speeches" ("id");
