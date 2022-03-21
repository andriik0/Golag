package main

import "fmt"

func main() {
	var a interface{}
	a = "jelly"
	fmt.Printf("(%v, %T)\n", a, a)
	a = float32(42)
	fmt.Printf("(%v, %T)\n", a, a)

	aa, ok := a.(string) //проверяем возможность приведения типа
	if ok {
		fmt.Println(aa)
	}

	switch a.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	}

	fmt.Println(customstructure{"Lopes", "Da Vegas"})

	err := m()
	if err != nil {
		fmt.Println(err.Error())
	}
}

//реализация стрингера
type customstructure struct {
	name, lastname string
}

func (m customstructure) String() string {
	return m.name + " " + m.lastname
}

//реализация кастомной ошибки
func m() error {
	return &AppError{
		Err:    fmt.Errorf("my error"),
		Custom: "value here",
		Field:  89,
	}
}

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (e *AppError) Error() string {
	fmt.Println(e.Custom)
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}
