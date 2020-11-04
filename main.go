package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/urfave/cli/v2"
)

// OutputKitinyi ...
type OutputKitinyi struct {
	Tiren   string `json:"tiren"`
	Alay    string `json:"alay"`
	Nyinyir string `json:"nyinyir"`
	Kipitil string `json:"kipitil"`
}

func createText(text string) (OutputKitinyi, error) {
	var vowNum = map[string]int{
		"a": 4, "A": 4, "e": 3, "E": 3, "i": 1, "I": 1, "o": 0, "O": 0, "u": 7, "U": 7,
	}

	var vowI = map[string]string{
		"a": "i", "A": "I", "e": "i", "E": "I", "i": "i", "I": "I", "o": "i", "O": "I", "u": "i", "U": "I",
	}

	var generated string = ""
	var alayed = []string{}
	var nyinyir = []string{}
	var kipitil = []string{}
	var result OutputKitinyi

	for idx, char := range text {
		if idx%2 == 0 && !unicode.IsSpace(char) {
			alayed = append(alayed, strings.ToUpper(string(char)))
		} else {
			alayed = append(alayed, string(char))
		}
	}

	generated = strings.Join(alayed, "")

	// replace char to num
	for idx, char := range generated {
		if _, ok := vowNum[string(char)]; ok {
			alayed[idx] = fmt.Sprint(vowNum[string(char)])
		}
	}

	// replace vocal with letter i
	for _, char := range generated {
		if _, ok := vowI[string(char)]; ok {
			nyinyir = append(nyinyir, fmt.Sprint(vowI[string(char)]))
			// fmt.Println(vowI[string(char)])
		} else {
			nyinyir = append(nyinyir, string(char))
		}
	}

	// convert nyinyir to kipitil
	for _, char := range nyinyir {
		if _, ok := vowNum[char]; ok {
			kipitil = append(kipitil, fmt.Sprint(vowNum[char]))
		} else {
			kipitil = append(kipitil, char)
		}
	}

	result = OutputKitinyi{
		Tiren:   generated,
		Alay:    strings.Join(alayed, ""),
		Nyinyir: strings.ToLower(strings.Join(nyinyir, "")),
		Kipitil: strings.Join(kipitil, ""),
	}

	return result, nil
}

func main() {
	var words string

	app := &cli.App{
		Name:  "kitinyi-go",
		Usage: "A golang version of kitinyi. Mingibih kiti-kiti ying kiti inggip minyibilkin igir sipiyi kiti bisi biginikin.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "text",
				Value:       "",
				Usage:       "Text you want to `change`",
				Aliases:     []string{"t"},
				Destination: &words,
			},
		},
		Action: func(c *cli.Context) error {

			if words != "" {
				// fmt.Println("text", words)
				result, _ := createText(words)
				output, err := json.MarshalIndent(&result, "", " ")
				if err != nil {
					fmt.Println("error marshalling to JSON:", err)
				}

				fmt.Println(string(output))
			} else {
				fmt.Println(`--t "your text"`)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
