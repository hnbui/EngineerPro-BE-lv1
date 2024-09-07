/// Viết hàm tạo ra 1 struct về 1 người (gồm tên, nghề nghiệp, năm sinh), và struct có method tính tuổi, method kiểm tra người có hợp với nghề của mình không (năm sinh chia hết cho số chữ trong tên).

package main

import "fmt"

type Person struct {
	Name        string
	Job         string
	YearOfBirth int
}

func (p *Person) calculateAge() int {
	return 2024 - p.YearOfBirth
}

func (p *Person) isCompatibleToJob() bool {
	return p.YearOfBirth%len(p.Name) == 0
}

func main() {
	person := Person{
		Name:        "Nam",
		Job:         "Developer",
		YearOfBirth: 1997,
	}
	fmt.Println(person.calculateAge())
	fmt.Println(person.isCompatibleToJob())
}
