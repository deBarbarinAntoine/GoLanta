package GoLanta

import (
	"encoding/json"
	"html/template"
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
		for _, char := range chars {
			if char.Id == id {
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

//	statToHTML
//
// converts the Character 's stats into StatHTML
// for displaying purposes in the templates.
func statToHTML(chars []Character) []StatsHTML {
	var statsHTML []StatsHTML
	for _, char := range chars {
		var statHTML StatsHTML

		// Strength
		if char.Strength > 5 {
			statHTML.Strength.Line = " style=\"background-color: limegreen\""
			statHTML.Strength.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Strength-5))
		} else {
			statHTML.Strength.Line = " style=\"background-color: orange\""
			statHTML.Strength.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Strength))
		}

		// Agility
		if char.Agility > 5 {
			statHTML.Agility.Line = " style=\"background-color: limegreen\""
			statHTML.Agility.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Agility-5))
		} else {
			statHTML.Agility.Line = " style=\"background-color: orange\""
			statHTML.Agility.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Agility))
		}

		// Stamina
		if char.Stamina > 5 {
			statHTML.Stamina.Line = " style=\"background-color: limegreen\""
			statHTML.Stamina.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Stamina-5))
		} else {
			statHTML.Stamina.Line = " style=\"background-color: orange\""
			statHTML.Stamina.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Stamina))
		}

		// Vitality
		if char.Vitality > 5 {
			statHTML.Vitality.Line = " style=\"background-color: limegreen\""
			statHTML.Vitality.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Vitality-5))
		} else {
			statHTML.Vitality.Line = " style=\"background-color: orange\""
			statHTML.Vitality.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Vitality))
		}

		// Initiative
		if char.Initiative > 5 {
			statHTML.Initiative.Line = " style=\"background-color: limegreen\""
			statHTML.Initiative.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Initiative-5))
		} else {
			statHTML.Initiative.Line = " style=\"background-color: orange\""
			statHTML.Initiative.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Initiative))
		}

		// Intelligence
		if char.Intelligence > 5 {
			statHTML.Intelligence.Line = " style=\"background-color: limegreen\""
			statHTML.Intelligence.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Intelligence-5))
		} else {
			statHTML.Intelligence.Line = " style=\"background-color: orange\""
			statHTML.Intelligence.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Intelligence))
		}

		// Knowledge
		if char.Knowledge > 5 {
			statHTML.Knowledge.Line = " style=\"background-color: limegreen\""
			statHTML.Knowledge.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Knowledge-5))
		} else {
			statHTML.Knowledge.Line = " style=\"background-color: orange\""
			statHTML.Knowledge.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Knowledge))
		}

		// Fame
		if char.Fame > 5 {
			statHTML.Fame.Line = " style=\"background-color: limegreen\""
			statHTML.Fame.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: limegreen\"></div>", char.Fame-5))
		} else {
			statHTML.Fame.Line = " style=\"background-color: orange\""
			statHTML.Fame.StatBar = template.HTML(strings.Repeat("<div class=\"stat-bar\" style=\"background-color: orange\"></div>", char.Fame))
		}

		statsHTML = append(statsHTML, statHTML)
	}
	return statsHTML
}

//	getTotalStat
//
// sums all Character 's stats to give a total.
func getTotalStat(chars []Character) []int {
	var total []int
	for _, char := range chars {
		total = append(total, char.Strength+char.Agility+char.Stamina+char.Vitality+char.Initiative+char.Intelligence+char.Knowledge+char.Fame)
	}
	return total
}
