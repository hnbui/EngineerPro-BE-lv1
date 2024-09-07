/// Cho 1 file a.txt, trả về slice các struct là người, với các thông tin lấy được từ filer đó, tên cần được in hoa tất cả, nghề nghiệp cần được viết thường tất cả (tham khâo đọc file ỏ https://zetcode.com/golang/readfile/)

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name        string
	Job         string
	YearOfBirth int
}

func processFile(filename string) []*Person {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	people := make([]*Person, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Split(line, "|")
		if len(fields) < 3 {
			log.Printf("Skipping line with insufficient fields: %s\n", line)
			continue
		}
		year, err := strconv.Atoi(strings.TrimSpace(fields[2]))
		if err != nil {
			log.Printf("Invalid year format for line: %s\n", line)
			continue
		}
		people = append(people, &Person{
			Name:        strings.ToUpper(strings.TrimSpace(fields[0])),
			Job:         strings.ToLower(strings.TrimSpace(fields[1])),
			YearOfBirth: year,
		})
	}
	return people
}

func main() {
	people := processFile("a.txt")
	for _, person := range people {
		fmt.Printf("%s (%s) - %d\n", person.Name, person.Job, person.YearOfBirth)
	}
}
