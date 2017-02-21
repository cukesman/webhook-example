package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHelloGopher(t *testing.T) {

	assert := assert.New(t)

	const Event1Json = `{
	  "event": "123"
	}`

	t.Run("Parse event from json.", func(t *testing.T) {
		var event1 Event
		err := json.Unmarshal([]byte(Event1Json), &event1)

		assert.Nil(err)

		assert.Equal("123", event1.Type)
	})

}
