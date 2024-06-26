// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: urls.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createContest = `-- name: CreateContest :one
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
) RETURNING id, reading_id, diktant_id, speeches_id, start_time, end_time
`

type CreateContestParams struct {
	ReadingID  pgtype.Int8
	DiktantID  pgtype.Int8
	SpeechesID pgtype.Int8
	StartTime  pgtype.Timestamp
	EndTime    pgtype.Timestamp
}

func (q *Queries) CreateContest(ctx context.Context, arg CreateContestParams) (Contest, error) {
	row := q.db.QueryRow(ctx, createContest,
		arg.ReadingID,
		arg.DiktantID,
		arg.SpeechesID,
		arg.StartTime,
		arg.EndTime,
	)
	var i Contest
	err := row.Scan(
		&i.ID,
		&i.ReadingID,
		&i.DiktantID,
		&i.SpeechesID,
		&i.StartTime,
		&i.EndTime,
	)
	return i, err
}

const createDiktant = `-- name: CreateDiktant :one
INSERT INTO diktants (
  text
) VALUES (
  $1
) RETURNING id, text
`

func (q *Queries) CreateDiktant(ctx context.Context, text pgtype.Text) (Diktant, error) {
	row := q.db.QueryRow(ctx, createDiktant, text)
	var i Diktant
	err := row.Scan(&i.ID, &i.Text)
	return i, err
}

const createQuestionVariant = `-- name: CreateQuestionVariant :one
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
) RETURNING id, question_id, option, is_correct, explanation, created_at
`

type CreateQuestionVariantParams struct {
	QuestionID  pgtype.Int8
	Option      string
	IsCorrect   bool
	Explanation string
}

func (q *Queries) CreateQuestionVariant(ctx context.Context, arg CreateQuestionVariantParams) (QuestionVariant, error) {
	row := q.db.QueryRow(ctx, createQuestionVariant,
		arg.QuestionID,
		arg.Option,
		arg.IsCorrect,
		arg.Explanation,
	)
	var i QuestionVariant
	err := row.Scan(
		&i.ID,
		&i.QuestionID,
		&i.Option,
		&i.IsCorrect,
		&i.Explanation,
		&i.CreatedAt,
	)
	return i, err
}

const createReading = `-- name: CreateReading :one
INSERT INTO readings (
  text
) VALUES (
  $1
) RETURNING id, text, created_at
`

func (q *Queries) CreateReading(ctx context.Context, text pgtype.Text) (Reading, error) {
	row := q.db.QueryRow(ctx, createReading, text)
	var i Reading
	err := row.Scan(&i.ID, &i.Text, &i.CreatedAt)
	return i, err
}

const createReadingQuestion = `-- name: CreateReadingQuestion :one
INSERT INTO reading_questions (
  text,
  question,
  reading_id
) VALUES (
  $1,
  $2,
  $3
) RETURNING id, text, question, reading_id, created_at
`

type CreateReadingQuestionParams struct {
	Text      string
	Question  string
	ReadingID pgtype.Int8
}

func (q *Queries) CreateReadingQuestion(ctx context.Context, arg CreateReadingQuestionParams) (ReadingQuestion, error) {
	row := q.db.QueryRow(ctx, createReadingQuestion, arg.Text, arg.Question, arg.ReadingID)
	var i ReadingQuestion
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Question,
		&i.ReadingID,
		&i.CreatedAt,
	)
	return i, err
}

const createSpeech = `-- name: CreateSpeech :one
INSERT INTO speeches (
  text
) VALUES (
  $1
) RETURNING id, text
`

func (q *Queries) CreateSpeech(ctx context.Context, text pgtype.Text) (Speech, error) {
	row := q.db.QueryRow(ctx, createSpeech, text)
	var i Speech
	err := row.Scan(&i.ID, &i.Text)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  full_name,
  username,
  password
) VALUES (
  $1,
  $2,
  $3
) RETURNING id, full_name, username, password, created_at
`

type CreateUserParams struct {
	FullName pgtype.Text
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.FullName, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getContest = `-- name: GetContest :one
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
GROUP BY c.id, c.start_time, c.end_time, r.id, r.text, d.id, d.text, s.id, s.text
`

type GetContestRow struct {
	ContestID   int64
	StartTime   pgtype.Timestamp
	EndTime     pgtype.Timestamp
	ReadingID   int64
	ReadingText pgtype.Text
	DiktantID   int64
	DiktantText pgtype.Text
	SpeechID    int64
	SpeechText  pgtype.Text
	Questions   []byte
}

func (q *Queries) GetContest(ctx context.Context, id int64) (GetContestRow, error) {
	row := q.db.QueryRow(ctx, getContest, id)
	var i GetContestRow
	err := row.Scan(
		&i.ContestID,
		&i.StartTime,
		&i.EndTime,
		&i.ReadingID,
		&i.ReadingText,
		&i.DiktantID,
		&i.DiktantText,
		&i.SpeechID,
		&i.SpeechText,
		&i.Questions,
	)
	return i, err
}

const getLastTwoContests = `-- name: GetLastTwoContests :many
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
LIMIT 2
`

type GetLastTwoContestsRow struct {
	ContestID   int64
	StartTime   pgtype.Timestamp
	EndTime     pgtype.Timestamp
	ReadingID   int64
	ReadingText pgtype.Text
	DiktantID   int64
	DiktantText pgtype.Text
	SpeechID    int64
	SpeechText  pgtype.Text
	Questions   []byte
}

func (q *Queries) GetLastTwoContests(ctx context.Context) ([]GetLastTwoContestsRow, error) {
	rows, err := q.db.Query(ctx, getLastTwoContests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLastTwoContestsRow
	for rows.Next() {
		var i GetLastTwoContestsRow
		if err := rows.Scan(
			&i.ContestID,
			&i.StartTime,
			&i.EndTime,
			&i.ReadingID,
			&i.ReadingText,
			&i.DiktantID,
			&i.DiktantText,
			&i.SpeechID,
			&i.SpeechText,
			&i.Questions,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, full_name, username, password, created_at FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
