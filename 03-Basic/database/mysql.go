package database

var connection string

/*
*
Function init akan selalu dieksekusi pertama kali ketika sebuah package digunakan
*/
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
