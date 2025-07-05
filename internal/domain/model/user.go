package model

type User struct {
	ID    string
	Name  string
	Email string
}

func (u User) IsValidName() bool {
	return len(u.Name) > 0 && len(u.Name) <= 10
}

func (u User) IsValidEmail() bool {
	return len(u.Email) > 0 && len(u.Email) <= 25
}
