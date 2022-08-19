package challenges

import "fmt"

type Data struct {
	name string
	age  int
}

func MutationSlice() {
	// Slice
	s := []int{1, 2, 3}
	fmt.Printf("S[1] -> %p\n", &s[1])
	p := s                            // Same as s, pointer to s array
	p[1] = 4                          // Shadowing s[1]
	fmt.Printf("P[1] -> %p\n", &p[1]) // Same as s[1]

	fmt.Println("s =", s)
	fmt.Println("p =", p)
}

func MutationArray() {
	a := [3]int{1, 2, 3} // Array is fixed-size
	b := a               // Copy all value in a to b

	// Different pointer
	fmt.Printf("A[1] -> %p\n", &a[1])
	fmt.Printf("B[1] -> %p\n", &b[1])

	// Mutate element 1 in b
	b[1] = 4
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func MutationMap() {
	// Map
	m := map[string]int{"level": 5, "health": 9}
	fmt.Println("m =", m)
	n := m
	n["food"] = 12

	fmt.Println("m =", m)
	fmt.Println("n =", n)
}

func MutationStruct() {
	t1 := Data{
		name: "Foo",
		age: 21,
	}

	t2 := t1

	t2.age = 19

	fmt.Printf("T1 pointer: %p\n", &t1)
	fmt.Printf(" | T1.age pointer: %p\n", &t1.age)

	fmt.Printf("T2 pointer: %p\n", &t2)
	fmt.Printf(" | T2.age pointer: %p\n", &t2.age)

	fmt.Printf("T1 = %v\n", t1)
	fmt.Printf("T2 = %v\n", t2)
}