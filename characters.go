package GoLanta

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var jsonfile = path + "content/characters.json"

//	retrieveChars
//
// retrieves all Character present in characters.json and stores them in a slice of Character.
// It returns the slice of Character and an error.
func retrieveChars() ([]Character, error) {
	var chars []Character

	data, err := os.ReadFile(jsonfile)

	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, &chars)
	if err != nil {
		return nil, err
	}

	return chars, nil
}

//	searchChar
//
// retrieves all Character in which Character.Name contains `search` and returns them in a slice.
func searchChar(search string) []Character {
	var result []Character

	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}

	for _, char := range chars {
		if strings.Contains(strings.ToLower(char.Name), strings.ToLower(search)) {
			result = append(result, char)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

//	changeChars
//
// overwrites characters.json with `chars` in json format.
func changeChars(chars []Character) {
	data, errJSON := json.Marshal(chars)
	if errJSON != nil {
		log.Fatal("log: createChar()\t JSON Marshall error!\n", errJSON)
	}
	errWrite := os.WriteFile(jsonfile, data, 0666)
	if errWrite != nil {
		log.Fatal("log: createChar()\t WriteFile error!\n", errWrite)
	}
}

//	getIdNewChar
//
// returns first unused id in characters.json.
func getIdNewChar() int {
	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}
	var id int
	var idFound bool
	for id = 1; !idFound; id++ {
		idFound = true
		for _, article := range chars {
			if article.Id == id {
				idFound = false
			}
		}
	}
	id--
	return id
}

//	createChar
//
// adds the Character `newChar` to characters.json.
func createChar(newChar Character) {
	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}
	chars = append(chars, newChar)
	changeChars(chars)
}

//	removeChar
//
// remove the Character which Character.Id is sent in argument from characters.json.
func removeChar(id int) {
	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}
	for i, char := range chars {
		if char.Id == id {
			chars = append(chars[:i], chars[i+1:]...)
		}
	}
	changeChars(chars)
}

// selectChar
// returns the Character which Character.Id matches the `id` argument.
func selectChar(id int) (Character, bool) {
	var char Character
	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}
	var ok bool
	for _, singleChar := range chars {
		if singleChar.Id == id {
			ok = true
			char = singleChar
		}
	}
	return char, ok
}

//	updateChar
//
// modifies the Character in characters.json that matches
// `updatedChar`'s Id with `updatedChar`'s content.
func updateChar(updatedChar Character) {
	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}
	for i, char := range chars {
		if char.Id == updatedChar.Id {
			chars[i] = updatedChar
		}
	}
	changeChars(chars)
}
