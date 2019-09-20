package table

type ColumnDefinition struct {
	columnMetaInformations []*ColumnMetaInformation
}

func NewColumnDefinition(infos []*ColumnMetaInformation) *ColumnDefinition {
	return &ColumnDefinition{columnMetaInformations: infos}
}

func (this *ColumnDefinition) FillTableWithDefaultValue(table *Table) *Table {
	rows := []*Row{}
	for _, row := range table.rows {
		columns := []*Column{}
		for _, info := range this.columnMetaInformations {
			defaultColumn := info.defaultColumn()
			if !row.hasSameName(defaultColumn) {
				columns = append(columns, defaultColumn)
			}
		}
		columns = append(columns, row.columns...)
		rows = append(rows, NewRow(columns).Sorted())
	}
	return NewTable(table.name, rows)
}

