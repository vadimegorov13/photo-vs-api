package models

import (
	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Title          string `json:"title"`
	Description    string `json:"description"`
	MaxPlayers     int    `json:"max_players"`
	MaxSubmissions int    `json:"max_submissions"`
	Status         string `json:"status"` // e.g. "upcoming", "in progress", "finished"
	OrganizerID    uint   `json:"organizer_id"`
}

// Organizer       User         `gorm:"foreignKey:OrganizerID" json:"organizer"`
// 	RegisteredUsers []User       `json:"registered_users"`
// 	Rounds          []Round      `json:"rounds"`
// 	Submissions     []Submission `json:"submissions"`

type Round struct {
	gorm.Model
	Order   int     `json:"order"`
	Matches []Match `json:"matches"`
}

type Match struct {
	gorm.Model
	TournamentID uint   `json:"tournament_id"`
	Player1      uint   `json:"player1"`
	Player2      uint   `json:"player2"`
	Submission1  uint   `json:"submission1"`
	Submission2  uint   `json:"submission2"`
	Winner       uint   `json:"winner"`
	Status       string `json:"status"` // e.g. "upcoming", "in progress", "finished"
}

type Submission struct {
	gorm.Model
	Title        string `json:"title"`
	TournamentID uint   `json:"tournament_id"`
	UserID       uint   `json:"user_id"`
	ImageURL     string `json:"image_url"`
	Votes        int    `json:"votes"`
	UserVotes    []Vote `json:"user_votes"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Vote struct {
	gorm.Model
	UserID       uint `json:"user_id"`
	TournamentID uint `json:"tournament_id"`
	SubmissionID uint `json:"submission_id"`
}
