package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// 1. prepare empty struct to store the json data
type Person struct {
	Age  int
	Name string
}

// 2. parse json data into slice []Person
func GetPassengerList() []Person {
	// 2.1 read the json file
	content, err := os.ReadFile("backend-titanic-test.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 2.2 unmarshal to decode json file
	var persons []Person
	err2 := json.Unmarshal(content, &persons)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	return persons
}

// 3. find most frequent age numbers
func MostCommonAge(persons []Person) (mostCommonAge int, names []string) {
	namesByAge := map[int][]string{}
	// 3.1 group the name by it's common age
	for _, p := range persons {
		value, found := namesByAge[p.Age]
		if !found {
			value = []string{}
		}
		namesByAge[p.Age] = append(value, p.Name)
	}
	// 3.2 compare the group size based on how many name it has
	for age, nameList := range namesByAge {
		if len(nameList) > len(names) {
			mostCommonAge, names = age, nameList
		}
	}
	return mostCommonAge, names
}

func main() {
	// call the slice from parsed json file
	passengers := GetPassengerList()

	// sort it's name apabethically
	sort.Slice(passengers, func(i, j int) bool {
		if passengers[i].Age == passengers[j].Age {
			return passengers[i].Name < passengers[j].Name
		}
		return passengers[i].Age < passengers[j].Age
	})

	// find the most common age from sorted slice
	age, names := MostCommonAge(passengers)

	// print most common age occured from total 93 passanger
	fmt.Printf("most common age : %d\n", age)

	// print the names based on its common age
	fmt.Println(names)
}
