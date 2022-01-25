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
			var err error
			nouns, err = loadWords(tc.Params["anchorDataPath"])
			if err != nil {
				return fmt.Errorf("cannot load nouns: %w", err)
			}
			fmt.Println("Got nouns", nouns)
		}
		entry.Headline = "xxxzy"
		return nil
	},
}
