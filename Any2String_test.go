package go_transform

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolToStringTrueCase(t *testing.T) {
	var input bool
	input = true
	result := BoolToString(input)
	assert.Equal(t, "true", result)
}
func TestBoolToStringFalseCase(t *testing.T) {
	var input bool
	input = false
	result := BoolToString(input)
	assert.Equal(t, "false", result)
}
func TestAnyIntToString(t *testing.T) {
	t.Log(math.MaxInt64)
	assert.Equal(t, "127", IntToString(math.MaxInt8))
	assert.Equal(t, "32767", IntToString(math.MaxInt16))
	assert.Equal(t, "2147483647", IntToString(math.MaxInt32))
	assert.Equal(t, "9223372036854775807", IntToString(math.MaxInt64))
	assert.Equal(t, "-128", IntToString(math.MinInt8))
	assert.Equal(t, "-32768", IntToString(math.MinInt16))
	assert.Equal(t, "-2147483648", IntToString(math.MinInt32))
	assert.Equal(t, "-9223372036854775808", IntToString(math.MinInt64))
	assert.Equal(t, "0", IntToString(0))
}
func TestAnyUintToString(t *testing.T) {
	assert.Equal(t, "18446744073709551615", UintToString(math.MaxUint))
	assert.Equal(t, "255", UintToString(math.MaxUint8))
	assert.Equal(t, "65535", UintToString(math.MaxUint16))
	assert.Equal(t, "4294967295", UintToString(math.MaxUint32))
	assert.Equal(t, "18446744073709551615", UintToString(math.MaxUint64))
	assert.Equal(t, "0", UintToString(0))
}
