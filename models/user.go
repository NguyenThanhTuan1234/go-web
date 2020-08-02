package models

type User struct {
	ID       int
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

func (u *User) GetId() int {
	return u.ID
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPassWord() []byte {
	return u.Password
}

func (u *User) GetFirst() string {
	return u.First
}

func (u *User) GetLast() string {
	return u.Last
}

func (u *User) GetRole() string {
	return u.Role
}

func NewUser(id int, userName string, passWord []byte, first string, last string, role string) User {
	return User{
		ID:       id,
		UserName: userName,
		Password: passWord,
		First:    first,
		Last:     last,
		Role:     role,
	}
}
