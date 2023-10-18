package main

import (
	"errors"
	"fmt"
)

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (userMap map[string]user, err error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap = make(map[string]user)
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}
	return userMap, nil
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {

	if _, ok := users[name]; !ok {
		return false, errors.New("user not in map")
	}
	delete(users, name)
	return true, nil
}

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

func main() {
	// ages := make(map[string]int)
	// ages["matt"] = 22
	// ages["alyssia"] = 23

	ages := map[string]int{
		"matt":    22,
		"alyssia": 23,
	}

	fmt.Println(ages)
	fmt.Println(len(ages))

	fmt.Println(getNameCounts([]string{"matt", "matthew", "mathieu", "alyssia", "timmy", "elizabeth", "timmeh"}))
}
