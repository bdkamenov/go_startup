package main

import (
	"go_startup/src/hw2"
)

func main() {

	var f = hw2.NewEditor("Тука пиша някъв стринг")
	println("new editor:", f.String())

	f = f.Insert(29, "тъп ")
	println(f.String())

	f = f.Insert(12314, ". ")
	println(f.String())

	f = f.Delete(36, 195959)
	println(f.String())

	f = f.Undo()
	println(f.String())

	f = f.Undo()
	println(f.String())

	f = f.Insert(29, "hubav ")
	println(f.String())

	f = f.Undo()
	println(f.String())

	f = f.Redo()
	println(f.String())

}