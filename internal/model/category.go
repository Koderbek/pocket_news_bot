package model

type Category struct {
	Id   int8   `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
}

func AllCategory() *Category {
	return &Category{
		Id:   0,
		Code: "all",
		Name: "🧐 Выбрать все",
	}
}
