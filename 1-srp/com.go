package srp

type UserSRP struct {
	FirstName string
	LastName  string
}

func (u *UserSRP) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

type UserRepository struct {
	// Conecta no banco de dados
}

func (r *UserRepository) Save(u *UserSRP) error {
	// Salva o usuario no banco de dados
	return nil
}
