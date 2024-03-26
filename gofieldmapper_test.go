package gofieldmapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID       string   `fmap:"-"`
	Username string   `fmap:"username,omitempty"`
	Name     string   `fmap:"name,omitempty"`
	Age      int      `fmap:"age,omitempty"`
	Meta     []string `fmap:"metadata,omitempty"`
}

func TestGet(t *testing.T) {
	user := &User{
		ID:       "123",
		Username: "foobar",
		Name:     "Foo Bar",
		Age:      50,
		Meta:     []string{},
	}
	mp, err := Get(user)
	t.Log(mp)
	if assert.NoError(t, err) {
		assert.Equal(t, 3, len(mp))
	}

	user = &User{
		ID:       "123",
		Username: "foobar1",
	}
	mp, err = Get(user)
	t.Log(mp)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(mp))
	}
}
