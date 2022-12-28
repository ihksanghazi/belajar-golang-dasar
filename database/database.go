package database

var connection string

func init() {
	connection = "MYSQL"
}

func Database() string {
	return connection
}
