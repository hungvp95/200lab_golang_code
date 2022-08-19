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