package db

type Connection struct {
}

type R map[string]interface{}

func (this Connection) ShouldHaveTable(records...R) {
}

func DefineConnection(connectionName, rdbms, information string) {
}

func DB(connectionName, tableName string) *Connection {
	return new(Connection)
}
