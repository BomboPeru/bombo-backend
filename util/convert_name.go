package util

import (
	"regexp"
	"strings"
	"errors"
	"io/ioutil"
)

var tableNames []string


func init() {
	tableNamesFile := "./util/names_eq.csv"
	if err := getTableNames(tableNamesFile); err != nil {
		panic(err)
	}
}


func getTableNames(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	rawTable := string(data)

	tableNames = strings.Split(rawTable, "\n")
	return nil
}

func GetMiniName(name string) (string, error) {

	name = strings.Replace(name, "(C)", "", -1)
	if strings.HasSuffix(name, " ") {
		name = name[0:len(name)-1]
	}
	reg, err := regexp.Compile(`([A-Z])\w+ ([A-Z])\w+`)
	if err != nil {
		return name, err
	}

	if reg.Match([]byte(name)) {
		chunks := strings.Split(name, " ")
		if !(len(chunks)>1) {
			return name, errors.New("invalid name")
		}

		finalName := chunks[0] +
			" " +
			strings.ToUpper(string(chunks[1][0])) +
			"."
		return finalName, nil
	}

	return name, errors.New("not match with that name")
}


func MatchNames(name1, name2 string) bool {

	indexName1 := -1
	indexName2 := -2

	mini1, _ := GetMiniName(name1)
	mini2, _ := GetMiniName(name2)

	if name1 == mini2 || name2 == mini1 {
		return true
	}

	for i, row :=  range tableNames {
		if strings.Contains(row, name1) {
			indexName1 = i
			if indexName2 != -2 {
				break
			}
		}

		if strings.Contains(row, name2) {
			indexName2 = i
			if indexName1 != -1 {
				break
			}
		}
	}
	return indexName1 == indexName2

}