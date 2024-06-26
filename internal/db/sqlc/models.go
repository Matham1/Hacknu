// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Contest struct {
	ID         int64
	ReadingID  pgtype.Int8
	DiktantID  pgtype.Int8
	SpeechesID pgtype.Int8
	StartTime  pgtype.Timestamp
	EndTime    pgtype.Timestamp
}

type Diktant struct {
	ID   int64
	Text pgtype.Text
}

type QuestionVariant struct {
	ID          int64
	QuestionID  pgtype.Int8
	Option      string
	IsCorrect   bool
	Explanation string
	CreatedAt   pgtype.Timestamptz
}

type Reading struct {
	ID        int64
	Text      pgtype.Text
	CreatedAt pgtype.Timestamptz
}

type ReadingQuestion struct {
	ID        int64
	Text      string
	Question  string
	ReadingID pgtype.Int8
	CreatedAt pgtype.Timestamptz
}

type Speech struct {
	ID   int64
	Text pgtype.Text
}

type User struct {
	ID        int64
	FullName  pgtype.Text
	Username  string
	Password  string
	CreatedAt pgtype.Timestamptz
}
