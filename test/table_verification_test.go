package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/ShoichiroKitano/micro_test/db/infra"
	. "github.com/ShoichiroKitano/micro_test"

	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/runner"
)

type PrinterSpy struct {
	results []string
}

func NewPrinterSpy() *PrinterSpy {
	return new(PrinterSpy)
}

func (this *PrinterSpy) Println(str string) {
	this.results = append(this.results, str)
}

func setUpSpy() *PrinterSpy {
	spy := NewPrinterSpy()
	runner.SetPrinter(spy)
	return spy
}

func TestDBはカラムの値を正しいと判定する(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")
			db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
			db.DB("conName").ShouldHaveTable(
				db.Table("test").
					Columns("column1", "column2").
					Record("A1", "A2").
					Record("B1", "B2"),
			)
		})
	})
	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestDBはカラムの値の誤りを検出する(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")

			db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
			db.DB("conName").ShouldHaveTable(
				db.Table("test").
					Columns("column1", "column2").
					Record("A1", "A2").
					Record("BUG", "B2"),
			)
		})
	})
	runner.Run()
	assert.Equal(t, "A B", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}

func TestDBはカラム順序は無視して検証する(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("test", func() {
		It("test", func() {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A", "B")

			db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
			db.DB("conName").ShouldHaveTable(
				db.Table("test").
					Columns("column2", "column1").
					Record("B", "A"),
			)
		})
	})

	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestDBは行の順序が期待値と異なる場合はテストを失敗させる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")

			db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
			db.DB("conName").ShouldHaveTable(
				db.Table("test").
					Columns("column1", "column2").
					Record("B1", "B2").
					Record("A1", "A2"),
			)
		})
	})

	runner.Run()
	assert.Equal(t, "A B", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}

func Test失敗したテストに省略記法を使った場合(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	defer resetSuites()
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It(func() {
			InsertIntoTest("mysql", "root:@/test_micro_test", "A1", "A2")
			InsertIntoTest("mysql", "root:@/test_micro_test", "B1", "B2")

			db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
			db.DB("conName").ShouldHaveTable(
				db.Table("test").
					Columns("column1", "column2").
					Record("B1", "B2").
					Record("A1", "A2"),
			)
		})
	})

	runner.Run()
	assert.Equal(t, "A", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}
