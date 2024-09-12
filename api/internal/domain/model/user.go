package model

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// func (u *User) CalculateLevel() int {
// 	level := int(math.Floor(float64(u.Experience) / 100.0))
// 	return level
// }
