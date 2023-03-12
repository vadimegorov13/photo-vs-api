package models

import (
	"time"
)

type Basics struct {
	Id        string    `json:"id" gorm:"primarykey;uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Tournament struct {
	Basics
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	MaxPlayers      int           `json:"max_players"`
	MaxSubmissions  int           `json:"max_submissions"`
	OrganizerId     string        `json:"organizer_id"`
	RegisteredUsers []*User       `json:"registered_users" gorm:"many2many:tournament_users;"`
	Rounds          []*Round      `json:"rounds" gorm:"foreignKey:TournamentId;"`
	Submissions     []*Submission `json:"submissions" gorm:"foreignKey:TournamentId;"`
	Status          string        `json:"status"` // e.g. "pending", "in progress", "finished"
}

type Round struct {
	Basics
	Order        int         `json:"order"`
	TournamentId string      `json:"tournament_id"`
	Tournament   *Tournament `json:"tournament" gorm:"foreignKey:TournamentId"`
	Matches      []*Match    `json:"matches" gorm:"foreignKey:RoundId"`
	Status       string      `json:"status"` // e.g. "pending", "in progress", "finished"
}

type Match struct {
	Basics
	RoundId       string      `json:"round_id"`
	Round         *Round      `json:"round" gorm:"foreignKey:RoundId"`
	UserId1       string      `json:"user_id1"`
	User1         *User       `json:"user1" gorm:"foreignKey:UserId1"`
	UserId2       string      `json:"user_id2"`
	User2         *User       `json:"user2" gorm:"foreignKey:UserId2"`
	SubmissionId1 string      `json:"submission_id1"`
	Submission1   *Submission `json:"submission1" gorm:"foreignKey:SubmissionId1"`
	SubmissionId2 string      `json:"submission_id2"`
	Submission2   *Submission `json:"submission2" gorm:"foreignKey:SubmissionId2"`
	WinnerId      string      `json:"winner_id"`
	Winner        User        `json:"winner" gorm:"foreignKey:WinnerId"`
	Status        string      `json:"status"` // e.g. "pending", "in progress", "finished"
}

type Submission struct {
	Basics
	Title        string      `json:"title"`
	TournamentId string      `json:"tournament_id"`
	Tournament   *Tournament `json:"tournament" gorm:"foreignKey:TournamentId"`
	UserId       string      `json:"user_id"`
	User         *User       `json:"user" gorm:"foreignKey:UserId"`
	ImageURL     string      `json:"image_url"`
	Votes        int         `json:"votes"`
	UserVotes    []*Vote     `json:"user_votes" gorm:"foreignKey:SubmissionId"`
}

type User struct {
	Basics
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Submissions []*Submission `json:"submissions" gorm:"foreignKey:UserId"`
	Votes       []*Vote       `json:"votes" gorm:"foreignKey:UserId"`
	Tournaments []*Tournament `json:"tournaments" gorm:"many2many:tournament_users"`
}

type Vote struct {
	Basics
	UserId       string      `json:"user_id"`
	TournamentId string      `json:"tournament_id"`
	SubmissionId string      `json:"submission_id"`
	User         *User       `json:"user" gorm:"foreignKey:UserId"`
	Tournament   *Tournament `json:"tournament" gorm:"foreignKey:TournamentId"`
	Submission   *Submission `json:"submission" gorm:"foreignKey:SubmissionId"`
}
