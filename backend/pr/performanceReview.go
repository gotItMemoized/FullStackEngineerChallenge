package pr

import (
	"database/sql"

	"github.com/gotItMemoized/FullStackEngineerChallenge/backend/user"
)

type Review struct {
	ID       string     `json:"id" db:"id"`
	UserID   string     `json:"-" db:"userid"`
	User     user.User  `json:"user" db:"user"`
	IsActive bool       `json:"isActive" db:"isactive"`
	Feedback []Feedback `json:"feedback" db:"feedback"`
}

type Feedback struct {
	ID         string         `json:"id" db:"id"`
	ReviewID   string         `json:"review" db:"reviewid"`
	ReviewerID string         `json:"-" db:"reviewerid"`
	Reviewer   user.User      `json:"reviewer" db:"reviewer"`
	Message    sql.NullString `json:"message" db:"message"`
}

type FlatFeedback struct {
	ID         string         `json:"id" db:"id"`
	ReviewID   string         `json:"reviewId" db:"reviewid"`
	ReviewerID string         `json:"reviewerId" db:"reviewerid"`
	UserID     string         `json:"userId" db:"userid"`
	Name       string         `json:"name" db:"name"`
	Username   string         `json:"username" db:"username"`
	Message    sql.NullString `json:"message" db:"message"`
}
