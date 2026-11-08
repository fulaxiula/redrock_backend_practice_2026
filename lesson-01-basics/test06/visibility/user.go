package visibility

type User struct {
	Name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		age:  age,
	}
}

func (u *User) Age() int {
	return u.age
}

func (u *User) Grow() {
	u.grow()
}

func (u *User) grow() {
	u.age++
}

func (u *User) Stage() string {
	return stage(u.age)
}