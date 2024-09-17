package srp

type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Save() error {
	// Salva usuario no banco de dados

	return nil
}
