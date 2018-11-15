package main

import "testing"

func TestExampleFromReadme(t *testing.T) {
	var f = NewEditor("A large span of text")

	f = f.Insert(16, "English ")
	compare(t, "A large span of English text", f.String())
	println(9)

	f = f.Delete(2, 6)
	compare(t, "A span of English text", f.String())
	println(13)


	f = f.Delete(10, 8)
	compare(t, "A span of text", f.String())
	println(18)

}

func TestOutOfBound(t *testing.T) {
	var f = NewEditor("A span of text")

	compare(t, f.String()+"!", f.Insert(150, "!").String())
	println(26)
	compare(t, f.String(), f.Delete(150, 20).String())
	println(28)
	compare(t, "A span of", f.Delete(9, 200).String())
	println(30)

}

func TestUndo(t *testing.T) {
	var f = NewEditor("A large span of text")
	compare(t, f.String(), f.Undo().String())
	println()

	f = f.Insert(16, "English ")
	compare(t, f.String(), f.Delete(2, 6).Undo().String())
}

func TestSeveralUndos(t *testing.T) {
	var f = NewEditor("A large span of text").
		Insert(16, "English ").
		Delete(2, 6).Delete(10, 8)

	compare(t, "A span of text", f.String())
	compare(t, "A large span of text", f.Undo().Undo().Undo().String())
}

func TestRedo(t *testing.T) {
	var f = NewEditor("A large span of text").
		Insert(16, "English ").Delete(2, 6)

	compare(t, f.String(), f.Undo().Redo().String())
}

func TestSeveralRedos(t *testing.T) {
	var f = NewEditor("A large span of text").
		Insert(16, "English ").Delete(2, 6).Delete(10, 8).
		Undo().Undo().Undo()

	compare(t, "A large span of text", f.String())
	compare(t, "A span of text", f.Redo().Redo().Redo().String())
}

func TestOpAfterUndoInvalidatesRedo(t *testing.T) {
	var f = NewEditor("A large span of text").
		Insert(16, "English ").Undo().Delete(0, 2)

	compare(t, f.String(), f.Redo().String())
}

func TestUnicode(t *testing.T) {
	var f = NewEditor("Жълтата дюля беше щастлива и замръзна като гьон.").
		Delete(49, 3).Insert(49, ", че пухът, който цъфна,")

	compare(t, "Жълтата дюля беше щастлива, че пухът, който цъфна, замръзна като гьон.", f.String())
}

func compare(t *testing.T, exp, got string) {
	if got != exp {
		t.Errorf("Expect: %q; got %q", exp, got)
	}
}