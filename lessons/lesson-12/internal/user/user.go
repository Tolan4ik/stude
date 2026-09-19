package user

// TODO 1: Объяви структуру User с двумя полями:
// - Name string (экспортируемое поле, с большой буквы)
// - id   int    (НЕэкспортируемое поле, с маленькой буквы)
type User struct {
	Name string
	id   int
}

// TODO 2: Напиши функцию-конструктор New(name string, id int) User
// Она должна создавать и возвращать структуру User с переданными name и id.
func New(name string, id int) User {

	return User{
		Name: name,
		id:   id,
	}
}

// TODO 3: Напиши экспортируемый метод GetID() int для User
// Он нужен, чтобы внешний мир мог прочитать id через метод,
// хотя само поле id остаётся приватным!
func (u User) GetID() int {
	return u.id
}
