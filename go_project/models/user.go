package models
type User struct {
	Id       int
	UserName string
	Age      int
	Email    string
	AddTime  int
}

func (User) TabelUser() string {
	return "user"
}