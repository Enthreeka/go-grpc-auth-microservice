package entity

type User struct {
	ID       string `json:"id",`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
