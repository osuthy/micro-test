package test

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/osuthy/micro-test/db/infra"

	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db"
)

func TestDBはカラムの値を正しいと判定する(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func(c TC) {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")
			DB(c, "conName").ShouldHaveTable(
				Table("test").
					Columns("column1", "column2").
					Record("A1", "A2").
					Record("B1", "B2"),
			)
		})
	})
	Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestDBはカラムの値の誤りを検出する(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func(c TC) {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")

			DB(c, "conName").ShouldHaveTable(
				Table("test").
					Columns("column1", "column2").
					Record("A1", "A2").
					Record("BUG", "B2"),
			)
		})
	})
	Run()
	assert.Equal(t, "A B", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}

func TestDBはカラム順序は無視して検証する(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("test", func() {
		It("test", func(c TC) {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A", "B")

			DB(c, "conName").ShouldHaveTable(
				Table("test").
					Columns("column2", "column1").
					Record("B", "A"),
			)
		})
	})

	Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestDBは行の順序が期待値と異なる場合はテストを失敗させる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func(c TC) {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")

			DB(c, "conName").ShouldHaveTable(
				Table("test").
					Columns("column1", "column2").
					Record("B1", "B2").
					Record("A1", "A2"),
			)
		})
	})

	Run()
	assert.Equal(t, "A B", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}

func Test失敗したテストに省略記法を使った場合(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It(func(c TC) {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")

			DB(c, "conName").ShouldHaveTable(
				Table("test").
					Columns("column1", "column2").
					Record("B1", "B2").
					Record("A1", "A2"),
			)
		})
	})

	Run()
	assert.Equal(t, "A", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}

func TestDBはデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	TruncateTable("mysql", "root:@/test_micro_test", "test")

	Describe("A", func() {
		It(func(tc TC) {
			DB(tc, "conName").HasRecords(
				Table("test").
					Columns("column1", "column2").
					Record("A1", "A2").
					Record("B1", "B2"),
			)
		})
	})
	Run()

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertNextRow(t, rows, "A1", "A2")
	AssertNextRow(t, rows, "B1", "B2")
	AssertNextIsNone(t, rows)
}

func Test事前条件のデータの補完(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "record_completion_all_type")
	defer resetSuites()
	resetSuites()
	TruncateTable("mysql", "root:@/test_micro_test", "record_completion_all_type")

	Describe("A", func() {
		It(func(tc TC) {
			DB(tc, "conName").HasRecords(
				Table("record_completion_all_type").
					Columns("dummy").
					Record("dummy1").
					Record("dummy2"),
			)
		})
	})
	Run()

	rows := Select("mysql", "root:@/test_micro_test", "record_completion_all_type")
	defer rows.Close()
	var (
		dummy string

		tinyintc   int
		smallintc  int
		mediumintc int
		intc       int
		bigintc    int

		floatc   float64
		doublec  float64
		decimalc float64

		charc       string
		varcharc    string
		tinytextc   string
		textc       string
		mediumtextc string
		longtextc   string

		datec      string
		datetimec  string
		timestampc string
		timec      string
		yearc      string

		bitc        uint8
		binaryc     uint8
		varbinaryc  uint8
		tinyblobc   uint8
		blobc       uint8
		mediumblobc uint8
		longblobc   uint8
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

	assert.Equal(t, uint8(0x0), bitc)
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

	assert.Equal(t, uint8(0x0), bitc)
	assert.Equal(t, uint8(0x0), binaryc)
	assert.Equal(t, uint8(0x0), varbinaryc)
	assert.Equal(t, uint8(0x0), tinyblobc)
	assert.Equal(t, uint8(0x0), blobc)
	assert.Equal(t, uint8(0x0), mediumblobc)
	assert.Equal(t, uint8(0x0), longblobc)

	AssertNextIsNone(t, rows)
}

func TestTruncate(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	Describe("A", func() {
		It(func(tc TC) {
			DB(tc, "conName").HasRecords(Table("test").
				Columns("column1", "column2").
				Record("A1", "A2").
				Record("B1", "B2"),
			)
			DB(tc, "conName").Truncate("test")
		})
	})
	Run()

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	assertTableIsEmpty(t, rows)
}

func assertTableIsEmpty(t *testing.T, rows *sql.Rows) {
	assert.False(t, rows.Next())
}
