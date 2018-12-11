package test

import (
	"github.com/ShoichiroKitano/micro_test/mock_time"

	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestMockTimeはローカルマシンMacの日付を任意の日付に変更できる(t *testing.T) {
	expected := time.Date(2018,12,12,0,0,0,0,time.UTC)
	exYear, exMonth, exDay := expected.Date()

	mock_time.SetTime(2018,12,12)

	assert.Equal(t, time.Now().Year(), exYear)
	assert.Equal(t, time.Now().Month(), exMonth)
	assert.Equal(t, time.Now().Day(), exDay)
}
