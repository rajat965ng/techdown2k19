package domains

type User struct {
	Id       uint64 `json:"id"`
	FistName string `json:"fist_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
