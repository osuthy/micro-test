package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/osuthy/micro-test/db"
	. "github.com/osuthy/micro-test/db/infra"
)

func TestDBはデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	TruncateTable("mysql", "root:@/test_micro_test", "test")

	db.DB("conName").HasRecords(
		db.Table("test").
			Columns("column1", "column2").
			Record("A1", "A2").
			Record("B1", "B2"),
	)

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertNextRow(t, rows, "A1", "A2")
	AssertNextRow(t, rows, "B1", "B2")
	AssertNextIsNone(t, rows)
}

func Test事前条件のデータの補完(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "record_completion_all_type")
	TruncateTable("mysql", "root:@/test_micro_test", "record_completion_all_type")

	db.DB("conName").HasRecords(
		db.Table("record_completion_all_type").
			Columns("dummy").
			Record("dummy1").
			Record("dummy2"),
	)

	rows := Select("mysql", "root:@/test_micro_test", "record_completion_all_type")
	defer rows.Close()
	var (
		dummy string

		tinyintc int
		smallintc int
		mediumintc int
		intc int
		bigintc int

		floatc float64
		doublec float64
		decimalc float64

		charc string
		varcharc string
		tinytextc string
		textc string
		mediumtextc string
		longtextc string

		datec string
		datetimec string
		timestampc string
		timec string
		yearc string

		bitc uint8
		binaryc uint8
		varbinaryc uint8
		tinyblobc uint8
		blobc uint8
		mediumblobc uint8
		longblobc uint8
	)
	rows.Next()
	rows.Scan(
		&dummy,
		&tinyintc,
		&smallintc,
		&mediumintc,
		&intc,
		&bigintc,

		&floatc,
		&doublec,
		&decimalc,

		&charc,
		&varcharc,
		&tinytextc,
		&textc,
		&mediumtextc,
		&longtextc,

		&datec,
		&datetimec,
		&timestampc,
		&timec,
		&yearc,

		&bitc,
		&binaryc,
		&varbinaryc,
		&tinyblobc,
		&blobc,
		&mediumblobc,
		&longblobc,
	)
	assert.Equal(t, "dummy1", dummy)
	assert.Equal(t, 0, tinyintc)
	assert.Equal(t, 0, smallintc)
	assert.Equal(t, 0, mediumintc)
	assert.Equal(t, 0, intc)
	assert.Equal(t, 0, bigintc)

	assert.Equal(t, 0.0, floatc)
	assert.Equal(t, 0.0, doublec)
	assert.Equal(t, 0.0, decimalc)

	assert.Equal(t, "", charc)
	assert.Equal(t, "", varcharc)
	assert.Equal(t, "", tinytextc)
	assert.Equal(t, "", textc)
	assert.Equal(t, "", mediumtextc)
	assert.Equal(t, "", longtextc)

	assert.Equal(t, "1970-01-01", datec)
	assert.Equal(t, "1970-01-01 10:00:00", datetimec)
	assert.Equal(t, "1970-01-01 10:00:00.0", timestampc)
	assert.Equal(t, "10:00:00.000000", timec)
	assert.Equal(t, "1901", yearc)

	assert.Equal(t,	uint8(0x0), bitc)
	assert.Equal(t, uint8(0x0), binaryc)
	assert.Equal(t, uint8(0x0), varbinaryc)
	assert.Equal(t, uint8(0x0), tinyblobc)
	assert.Equal(t, uint8(0x0), blobc)
	assert.Equal(t, uint8(0x0), mediumblobc)
	assert.Equal(t, uint8(0x0), longblobc)

	rows.Next()
	rows.Scan(
		&dummy,
		&tinyintc,
		&smallintc,
		&mediumintc,
		&intc,
		&bigintc,

		&floatc,
		&doublec,
		&decimalc,

		&charc,
		&varcharc,
		&tinytextc,
		&textc,
		&mediumtextc,
		&longtextc,

		&datec,
		&datetimec,
		&timestampc,
		&timec,
		&yearc,

		&bitc,
		&binaryc,
		&varbinaryc,
		&tinyblobc,
		&blobc,
		&mediumblobc,
		&longblobc,
	)
	assert.Equal(t, "dummy2", dummy)
	assert.Equal(t, 0, tinyintc)
	assert.Equal(t, 0, smallintc)
	assert.Equal(t, 0, mediumintc)
	assert.Equal(t, 0, intc)
	assert.Equal(t, 0, bigintc)

	assert.Equal(t, 0.0, floatc)
	assert.Equal(t, 0.0, doublec)
	assert.Equal(t, 0.0, decimalc)

	assert.Equal(t, "", charc)
	assert.Equal(t, "", varcharc)
	assert.Equal(t, "", tinytextc)
	assert.Equal(t, "", textc)
	assert.Equal(t, "", mediumtextc)
	assert.Equal(t, "", longtextc)

	assert.Equal(t, "1970-01-01", datec)
	assert.Equal(t, "1970-01-01 10:00:00", datetimec)
	assert.Equal(t, "1970-01-01 10:00:00.0", timestampc)
	assert.Equal(t, "10:00:00.000000", timec)
	assert.Equal(t, "1901", yearc)

	assert.Equal(t,	uint8(0x0), bitc)
	assert.Equal(t, uint8(0x0), binaryc)
	assert.Equal(t, uint8(0x0), varbinaryc)
	assert.Equal(t, uint8(0x0), tinyblobc)
	assert.Equal(t, uint8(0x0), blobc)
	assert.Equal(t, uint8(0x0), mediumblobc)
	assert.Equal(t, uint8(0x0), longblobc)

	AssertNextIsNone(t, rows)
}

