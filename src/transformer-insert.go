package main

import "os"
import "bufio"
import "fmt"

var nouns []string

func loadWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

var InsertTransformer = Transformer{
	transform: func(tc TransformationConfig, entry *Entry) error {
		if nouns == nil {
			nouns, err := loadWords(tc.Params["anchorDataPath"])
			if err != nil {
				// XXX this is an ugly way to handle the error
				fmt.Println("cannot load nouns:", err)
				os.Exit(4)
			}
			fmt.Println("Got nouns", nouns)
		}
		entry.Headline = "xxxzy"
		return nil
	},
}
