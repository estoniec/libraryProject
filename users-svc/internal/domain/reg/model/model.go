package model

type User struct {
	ID       int64
	Phone    string
	Username string
	Class    string
	Status   int
}

func CreateUser(
	id int64,
	phone string,
	username string,
	class string) User {
	return User{
		ID:       id,
		Phone:    phone,
		Username: username,
		Class:    class,
		Status:   0,
	}
}
