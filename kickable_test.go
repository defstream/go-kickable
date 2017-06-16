package kickable_test

import (
	"strings"
	"testing"

	"github.com/defstream/go-kickable"
	"github.com/stretchr/testify/assert"
)

func TestNo(t *testing.T) {
	assert := assert.New(t)
	assert.True(strings.HasPrefix(kickable.No, "No"))
}

func TestYes(t *testing.T) {
	assert := assert.New(t)
	assert.True(strings.HasPrefix(kickable.Yes, "Yes"))
}

func TestKickable(t *testing.T) {
	assert := assert.New(t)
	assertKickableSuccess(assert)
	assertKickableFailure(assert)
}

func TestCanIKick(t *testing.T) {
	assert := assert.New(t)
	assertCanIKickSuccess(assert)
	assertCanIKickFailure(assert)
}

func assertKickableSuccess(assert *assert.Assertions) {
	v := kickable.Kickable("IT")
	assert.True(v, "IT should be kickable")

	v = kickable.Kickable("It")
	assert.True(v, "It should be kickable")

	v = kickable.Kickable("iT")
	assert.True(v, "iT should be kickable")

	v = kickable.Kickable("it")
	assert.True(v, "it should be kickable")
}

func assertKickableFailure(assert *assert.Assertions) {
	v := kickable.Kickable("")
	assert.False(v, "'' should not be kickable")

	v = kickable.Kickable("not it")
	assert.False(v, "'not it' should not be kickable")

	v = kickable.Kickable("blah")
	assert.False(v, "'blah' should not be kickable")

	v = kickable.Kickable("0")
	assert.False(v, "'0' should not be kickable")

	v = kickable.Kickable("1")
	assert.False(v, "'1' should not be kickable")
}

func assertCanIKickSuccess(assert *assert.Assertions) {
	v := kickable.CanIKick("IT")
	assert.True(strings.HasPrefix(v, "Yes"))

	v = kickable.CanIKick("It")
	assert.True(strings.HasPrefix(v, "Yes"))

	v = kickable.CanIKick("iT")
	assert.True(strings.HasPrefix(v, "Yes"))

	v = kickable.CanIKick("it")
	assert.True(strings.HasPrefix(v, "Yes"))
}

func assertCanIKickFailure(assert *assert.Assertions) {
	v := kickable.CanIKick("")
	assert.True(strings.HasPrefix(v, "No"))

	v = kickable.CanIKick("not it")
	assert.True(strings.HasPrefix(v, "No"))

	v = kickable.CanIKick("blah")
	assert.True(strings.HasPrefix(v, "No"))

	v = kickable.CanIKick("0")
	assert.True(strings.HasPrefix(v, "No"))

	v = kickable.CanIKick("1")
	assert.True(strings.HasPrefix(v, "No"))

	v = kickable.CanIKick("true")
	assert.True(strings.HasPrefix(v, "No"))
}
