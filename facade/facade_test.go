package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	assert.Equal(t, expect, ret)
}
