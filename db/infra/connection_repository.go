package infra

import (
	"database/sql"
)

// 自動テスト用のヘルパー

type ConnectionMap struct {
	driverName            string
	connectionInformation string
	connection            *Connection
}

var maps []ConnectionMap = []ConnectionMap{}

func FindDBConnection(driverName, connectionInformation string) *Connection {
	for _, connection := range maps {
		if connection.driverName == driverName && connection.connectionInformation == connectionInformation {
			return connection.connection
		}
	}
	driver, _ := sql.Open(driverName, connectionInformation)
	newConnection := NewConnection(driver, driverName)
	maps = append(maps, ConnectionMap{driverName, connectionInformation, newConnection})
	return newConnection
}
