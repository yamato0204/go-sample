package model

type Report struct {
	ID           string `db:"id" json:"id"`
	Comment      string `db:"comment" json:"comment"`
	ThumbnailUrl string `db:"thumbnail_url" json:"thumbnail_url"`
	AuthorID     string `db:"author_id" json:"author_id"`
	Recipe       Recipe `db:"recipe" json:"recipe_id"`
	CreatedAt    string `db:"created_at" json:"created_at"`
	UpdatedAt    string `db:"updated_at" json:"updated_at"`
}
