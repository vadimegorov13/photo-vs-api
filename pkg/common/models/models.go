package models

import (
	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	MaxPlayers      int           `json:"max_players"`
	MaxSubmissions  int           `json:"max_submissions"`
	Status          string        `json:"status"` // e.g. "upcoming", "in progress", "finished"
	OrganizerID     uint          `json:"organizer_id"`
	RegisteredUsers []*User       `gorm:"many2many:tournament_user" json:"registered_users"`
	Rounds          []*Round      `json:"rounds"`
	Submissions     []*Submission `json:"submissions"`
}

type Round struct {
	gorm.Model
	Order        int      `json:"order"`
	TournamentID uint     `json:"tournament_id"`
	Matches      []*Match `gorm:"foreignKey:TournamentID" json:"matches"`
}

type Match struct {
	gorm.Model
	TournamentID  uint   `json:"tournament_id"`
	UserID1       uint   `json:"user_id1"`
	UserID2       uint   `json:"user_id2"`
	SubmissionID1 uint   `json:"submission_id1"`
	SubmissionID2 uint   `json:"submission_id2"`
	WinnerID      uint   `json:"winner_id"`
	Status        string `json:"status"` // e.g. "upcoming", "in progress", "finished"
}

type Submission struct {
	gorm.Model
	Title        string  `json:"title"`
	TournamentID uint    `json:"tournament_id"`
	UserID       uint    `json:"user_id"`
	ImageURL     string  `json:"image_url"`
	Votes        int     `json:"votes"`
	UserVotes    []*Vote `json:"user_votes"`
}

type User struct {
	gorm.Model
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Submissions []*Submission `json:"submissions"`
	Votes       []*Vote       `json:"votes"`
	Tournaments []*Tournament `gorm:"many2many:user_votes" json:"tournamnets"`
}

type Vote struct {
	gorm.Model
	UserID       uint `json:"user_id"`
	TournamentID uint `json:"tournament_id"`
	SubmissionID uint `json:"submission_id"`
}
