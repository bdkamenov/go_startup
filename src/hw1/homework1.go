package main

func Repeater(s, sep string) func(int) string {

	return func(cnt int) string {
		var result string = ""

		for i := 0; i < cnt; i++ {

			result += s

			if i != cnt-1 {
				result += sep
			}
		}
		return result
	}
}

func Generator(gen func(int) int, initial int) func() int {

	next := initial

	return func() int {
		prev := next
		next = gen(next)
		return prev
	}
}

func MapReducer1(mapper func(int) int, reducer func(int, int) int, initial int) func(...int) int {

	return func(args ...int) int {

		cntOfArgs := len(args)

		if cntOfArgs == 0 {
			return initial
		} else {
			return reducer(mapper(args[0]), MapReducer1(mapper, reducer, initial)(args[1:]...))
		}

	}

}

func MapReducer2(mapper func(int) int, reducer func(int, int) int, initial int) func(...int) int {

	return func(args ...int) int {

		cntOfArgs := len(args)

		if cntOfArgs == 0 {
			return initial
		} else {
			return reducer(MapReducer(mapper, reducer, initial)(args[1:]...), mapper(args[0]))
		}

	}

}

func MapReducer(mapper func(int) int, reducer func(int, int) int, initial int) func(...int) int {

	return func(args ...int) int {

		if len(args) == 0 {

			return initial

		} else {

			result := reducer(initial, mapper(args[0]))

			for _, v := range args[1:] {
				result = reducer(result, mapper(v))
			}

			return result
		}
	}
}

func main() {

	powerSum := MapReducer(
		func(v int) int { return v * v },
		func(a, v int) int { return a + v },
		0,
	)

	println(powerSum(3))
}
