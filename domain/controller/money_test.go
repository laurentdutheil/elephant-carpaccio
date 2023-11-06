package controller_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestString(t *testing.T) {
	dollar := NewDollar(103)
	assert.Equal(t, "$1.03", fmt.Sprint(dollar))
}
