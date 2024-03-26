## How to use
`Get` accepts any struct or pointer to a struct and returns a map of values based on tags. 

### Example

```
type User struct {
	ID       string   `fmap:"-"`
	Username string   `fmap:"username,omitempty"`
	Name     string   `fmap:"name,omitempty"`
	Age      int      `fmap:"age,omitempty"`
	Meta     []string `fmap:"metadata,omitempty"`
}

user := &User{
    ID:       "123",
    Username: "foobar",
    Name:     "Foo Bar",
    Age:      50,
    Meta:     []string{},
}
mp, err := Get(user)
map[age:50 name:Foo Bar username:foobar]

user = &User{
    ID:       "123",
    Username: "foobar1",
}
mp, err = Get(user)
map[username:foobar1]

```