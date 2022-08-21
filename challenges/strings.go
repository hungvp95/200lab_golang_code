package challenges

import (
	"fmt"
	"strings"
)

func StringConcat() {
	s1 := "Go"
	s2 := "lang"
	s3 := "Programming"
	s4 := "Language"

	s := s1 + " " + s2 + " " + s3 + " " + s4
	fmt.Println(s)
}

func StringJoin() {
	q := []string{"github.com", "definev", "200lab_golang"}
	r := strings.Join(q, "/")
	fmt.Println(r)
}

func StringSprint() {
	name := "daiduong"
	domain := "workmail"
	service := "gmail"

	email := fmt.Sprintf("%s.%s@%s.com", name, domain, service)
	fmt.Println(email)
}

func StringBuilder() {
	c := [4]string{"D", "a", "r", "t"}
	var builder strings.Builder
	for _, item := range c {
		builder.WriteString(item)
	}
	fmt.Printf("Builder: %s\n", builder.String())

	b := [4]byte{'l', 'a', 'n', 'g'}
	for _, item := range b {
		builder.WriteByte(item)
	}
	fmt.Printf("Builder: %s\n", builder.String())

	builder.WriteRune(' ')

	r := [7]rune{'F', 'l', 'u', 't', 't', 'e', 'r'}
	for _, item := range r {
		builder.WriteRune(item)
	}

	fmt.Printf("Builder: %s\n", builder.String())
	fmt.Printf("Builder raw: %s\n", builder.String())
}

func StringCompare() {
	s1 := "gopher"
	s2 := "Gopher"
	s3 := "gopher"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	isEqual := s1 == s2

	//"gopher" is not same as "Gopher" and hence `false`
	fmt.Printf("S1 and S2 are Equal? %t \n", isEqual)
	fmt.Printf("s1 == s2 : %v\n", s1 == s2)

	// "gopher" is not equal to "Gopher" and hence `true`
	fmt.Printf("s1 != s2 : %v\n", s1 != s2)

	// "Gopher" comes first lexicographically than "gopher" so return true
	// G -> 71 in ASCII and g -> 103
	fmt.Printf("s2 < s3 : %v\n", s2 < s3)
	fmt.Printf("s2 <= s3 : %v\n", s2 <= s3)

	// "Gopher" is not greater than "gopher" as `G` comes first in ASCII table
	// So G value is less than g i.e. 71 > 103 which is false
	fmt.Printf("s2 > s3 : %v\n", s2 > s3)
	fmt.Printf("s2 >= s3 : %v\n", s2 >= s3)
}

func StringManipulation() {
	s := "golang"

	// ToUpper
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(s))

	// ToLower
	fmt.Printf("ToLower: %s\n", strings.ToLower(s))
	
	// Replace
	fmt.Printf("Replace: %s\n", strings.Repeat(s, 3))
}
