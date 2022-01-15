package entity

type User struct {
	ID int
	Age int
	Name string
	Gender int
}

func (u *User) IsAdult() bool {
	return u.Age >= 20
}