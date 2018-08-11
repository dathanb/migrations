package error

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestNotImplemented(t *testing.T) {
	t.Run("Returns an error with msg 'Not implemented'", func(t *testing.T) {
		assert.Equal(t, 	"Not implemented", NotImplemented().Error())
	})
}
