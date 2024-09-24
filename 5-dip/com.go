package comdip

type DBConnectionInterface interface {
	Connect()
}

type MySQLConnection struct{}

func (conn *MySQLConnection) Connect() {
	/* ... */
}

type OracleConnection struct{}

func (conn *OracleConnection) Connect() {
	/* ... */
}

type SQLServerConnection struct{}

func (conn *SQLServerConnection) Connect() {
	/* ... */
}

type PasswordReminder struct {
	dbConnection DBConnectionInterface
}

func NewPasswordReminder(dbConnection DBConnectionInterface) *PasswordReminder {
	return &PasswordReminder{
		dbConnection: dbConnection,
	}
}
