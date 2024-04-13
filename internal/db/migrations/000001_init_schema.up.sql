CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "full_name" VARCHAR(32),
  "username" VARCHAR(32) UNIQUE NOT NULL,
  "password" VARCHAR(32) NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "readings" (
  "id" BIGSERIAL PRIMARY KEY,
  "text" TEXT,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "reading_questions" (
  "id" BIGSERIAL PRIMARY KEY,
  "text" TEXT NOT NULL,
  "question" TEXT NOT NULL,
  "reading_id" bigint,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "question_variants" (
  "id" BIGSERIAL PRIMARY KEY,
  "question_id" bigint,
  "option" TEXT NOT NULL,
  "is_correct" bool NOT NULL,
  "explanation" TEXT NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "diktants" (
  "id" BIGSERIAL PRIMARY KEY,
  "text" TEXT
);

CREATE TABLE "speeches" (
  "id" BIGSERIAL PRIMARY KEY,
  "text" TEXT
);

CREATE TABLE "contests" (
  "id" BIGSERIAL PRIMARY KEY,
  "reading_id" bigint,
  "diktant_id" bigint,
  "speeches_id" bigint,
  "start_time" timestamp,
  "end_time" timestamp
);

ALTER TABLE "reading_questions" ADD FOREIGN KEY ("reading_id") REFERENCES "readings" ("id");

ALTER TABLE "question_variants" ADD FOREIGN KEY ("question_id") REFERENCES "reading_questions" ("id");

ALTER TABLE "contests" ADD FOREIGN KEY ("reading_id") REFERENCES "readings" ("id");

ALTER TABLE "contests" ADD FOREIGN KEY ("diktant_id") REFERENCES "diktants" ("id");

ALTER TABLE "contests" ADD FOREIGN KEY ("speeches_id") REFERENCES "speeches" ("id");
