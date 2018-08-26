package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	t.Run("Can unmarshal from StackExchange datetime format", func(t *testing.T) {
		var parsedTime Time
		datetime := "\"2013-12-18T19:59:50.907\""

		err := json.Unmarshal([]byte(datetime), &parsedTime)
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2013, 12, 18, 19, 59, 50, 907000000, time.UTC), parsedTime.Time)
	})
}

