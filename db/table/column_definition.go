package table

type ColumnDefinition struct {
	rdbms string
	tableName string
	columnMetaInformations *ColumnMetaInformation
}

func NewColumnDefinition(rdbms, tableName string, infos []*ColumnMetaInformation) *ColumnDefinition {
	return &ColumnDefinition{rdbms: rdbms, tableName: tableName, columnMetaInformations: infos}
}

