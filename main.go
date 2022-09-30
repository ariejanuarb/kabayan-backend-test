package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//1. prepare the empty struct
type Passenger struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 2. parse json data into slice []Passenger
func GetPassengerList() []Passenger {
	// 2.1 load the json file
	content, err := os.ReadFile("backend-titanic-test-1.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 2.2 unmarshal json file to slice
	var passengers []Passenger
	err2 := json.Unmarshal(content, &passengers)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	return passengers
}

// 3. group the name by its age
func GroupByAge(passengers []Passenger) map[int][]int {
	group := make(map[int][]int, 0)
	// 3.1 iterate trough the passengers slice for every same age
	for index, passenger := range passengers {
		ageGroup := group[passenger.Age]
		ageGroup = append(ageGroup, index)
		group[passenger.Age] = ageGroup
	}
	return group
}

// 4. find the most frequent age numbers
func FindMostCommonAge(ageGroups map[int][]int) []int {
	mostFrequentAges := make([]int, 0)
	biggestGroupSize := 0

	//  4.1 comparing the group size based on it's common age
	for age, ageGroup := range ageGroups {
		// is most frequent age
		if biggestGroupSize < len(ageGroup) {
			biggestGroupSize = len(ageGroup)
			mostFrequentAges = []int{age}
			// 4.2 split the array if there is multiple group with the same maximum fequency of the age
		} else if biggestGroupSize == len(ageGroup) { // is one of the most frequent age
			mostFrequentAges = append(mostFrequentAges, age)
		}
		// we assume that there always be group with most frequent age, so we do nothing if there is none
	}
	return mostFrequentAges
}

func main() {
	// load the slice
	passengers := GetPassengerList()

	// sort the slice by its name
	sort.Slice(passengers, func(i, j int) bool {
		if passengers[i].Age == passengers[j].Age {
			return passengers[i].Name < passengers[j].Name
		}
		return passengers[i].Age < passengers[j].Age
	})

	// age => []position
	// Length of the array count as the number of occurences
	ageGrouper := GroupByAge(passengers)

	// find most frequent age between group
	mostFrequentAges := FindMostCommonAge(ageGrouper)

	// print the passenger with most frequent-common age
	for _, age := range mostFrequentAges {
		fmt.Print("[")
		for _, passengerIndex := range ageGrouper[age] {
			fmt.Print(" ", passengers[passengerIndex].Name, " ")
		}
		fmt.Print("]")
	}
}
