package model

type Recipe struct {
	ID           string   `db:"id"`
	Title        string   `db:"comment"`
	ThumbnailUrl string   `db:"thumbnail_url"`
	Recipe       string   `db:"recipe"`
	Category     Category `db:"category"`
	CreatedAt    string   `db:"created_at"`
	UpdatedAt    string   `db:"updated_at"`
}
