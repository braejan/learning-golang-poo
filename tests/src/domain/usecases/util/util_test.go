package usecases_test

import (
	"testing"

	uc "github.com/braejan/learning-golang-poo/tests/src/domain/usecases/util"
	"github.com/stretchr/testify/assert"
)

var usecases uc.UtilUsecases

func initUsecase() {
	usecases = uc.NewUtilUsecase()
}

func TestReverseStringShouldBeOk(t *testing.T) {
	initUsecase()
	text := "Hello Tests"
	expected := "stseT olleH"
	result, err := usecases.ReverseString(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestReverseStringWithError(t *testing.T) {
	initUsecase()
	_, err := usecases.ReverseString("")
	assert.NotNil(t, err)
}
