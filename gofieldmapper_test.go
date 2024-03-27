package gofieldmapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID       string   `fmap:"-"`
	Username string   `fmap:"my_username,omitempty,mask=uname"`
	Name     string   `fmap:"my_name,omitempty,mask=name"`
	Age      int      `fmap:"age,omitempty,mask=age"`
	Meta     []string `fmap:"meta,omitempty"`
}

var user = &User{
	ID:       "123",
	Username: "foo_bar",
	Name:     "Foo Bar",
	Age:      30,
	Meta:     []string{"other"},
}

func TestFullMap(t *testing.T) {
	mp, err := Make(user)
	expected := map[string]any{
		"my_username": "foo_bar",
		"my_name":     "Foo Bar",
		"age":         30,
		"meta":        []string{"other"},
	}
	t.Log(mp)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, mp)
	}
}

func TestWithOmit(t *testing.T) {
	u := &User{
		ID:       "123",
		Username: "foo_bar",
		Name:     "",
		Age:      30,
		Meta:     []string{},
	}
	mp, err := Make(u, WithOmit())
	expected := map[string]any{
		"my_username": "foo_bar",
		"age":         30,
	}
	t.Log(mp)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, mp)
	}
}

func TestWithMask(t *testing.T) {
	mp, err := Make(user, WithMask([]string{"uname", "name"}))
	expected := map[string]any{
		"my_username": "foo_bar",
		"my_name":     "Foo Bar",
	}
	if assert.NoError(t, err) {
		assert.Equal(t, expected, mp)
	}
}
