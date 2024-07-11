package model

import "time"

type User struct {
	UserID    string    `db:"user_id"`
	UserName  string    `db:"username"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Video struct {
	VideoID     string    `db:"video_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Genre struct {
	GenreID   string `db:"genre_id"`
	GenreName string `db:"genre_name"`
}
