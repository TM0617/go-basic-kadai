package main

import (
	"fmt"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *Todo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.ID)
	t.Completed = true
}

func printTodos(todos []Todo) {
	for _, todo := range todos {
		status := "未完了"
		if todo.Completed {
			status = "完 了"
		}
		fmt.Printf("[%s] (ID: %d) %s \n", status, todo.ID, todo.Title)
	}
}

func main() {

	todo1 := Todo{
		ID:        1,
		Title:     "学習計画",
		Completed: false,
	}
	todo2 := Todo{
		ID:        2,
		Title:     "環境構築",
		Completed: false,
	}
	todo3 := Todo{
		ID:        3,
		Title:     "基礎文法",
		Completed: false,
	}
	todos := []Todo{todo1, todo2, todo3}

	fmt.Println("---初期状態---")
	printTodos(todos)

	for i := range todos {
		if todos[i].ID < 3 {
			todos[i].Complete()
		}
	}

	fmt.Println("---最終状態---")
	printTodos(todos)

}
