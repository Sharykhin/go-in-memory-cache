package cache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	err := NewError(UnsupportedTypeCode, "an error")

	assert.Equal(
		t,
		fmt.Sprintf("Error code: %d. Error message: %s", UnsupportedTypeCode, "an error"),
		err.Error(),
	)
}
