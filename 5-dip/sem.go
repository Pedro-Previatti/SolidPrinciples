package semdip

type MySQLConnection struct{}

func NewMySQLConnection() *MySQLConnection {
	return &MySQLConnection{}
}

func (conn *MySQLConnection) Connect() {
	/* ... */
}

type PasswordReminder struct {
	dbConnection *MySQLConnection
}

func NewPasswordReminder() *PasswordReminder {
	return &PasswordReminder{
		dbConnection: NewMySQLConnection(),
	}
}
