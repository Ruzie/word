package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

const API_URL = "https://api.dictionaryapi.dev/api/v2/entries/en/"

type ReponseType struct {
	Word     string `json:"word"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func log(msg string) {
	fmt.Print(msg)
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--h" || os.Args[1] == "--help" {
		log("Word v0.1\n\nUsage: word [query]\n")
		os.Exit(1)
	}

	query := os.Args[1]
	data := ""

	var resp []ReponseType
	/**
	Filter "%" as it can be abuse by trying to passing a "%[format_type] (ex. %d)" which will refer a non-existance memory address
	leades the program to cause a SIGSEGV (Segmentation violation) error.
	*/
	body, err := http.Get(API_URL + strings.ReplaceAll(query, "%", ""))
	if body.StatusCode == 404 {
		log("No results found matching your query\n")
		return
	}

	buf, err := ioutil.ReadAll(body.Body)
	err = json.Unmarshal([]byte(string(buf)), &resp)

	if err != nil {
		log("Failed to get the result\n\nDebug: " + err.Error())
	}

	defer body.Body.Close()

	data += strings.Title(strings.ToLower("# " + resp[0].Word + "\n\n"))
	for i := 0; i < len(resp[0].Meanings[0].Definitions); i++ {
		data += "* " + resp[0].Meanings[0].Definitions[i].Definition + "\n"
	}
	pretty, _ := glamour.Render(data, "dark")
	log(pretty)
}
