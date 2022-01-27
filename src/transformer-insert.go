package main

import "os"
import "bufio"
import "fmt"
import "strings"
import "math/rand"

// private
var _transformInsert_nounRegister map[string]bool

func loadWords(path string) (map[string]bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res = map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

var InsertTransformer = Transformer{
	transform: func(tc transformationConfig, entry *Entry) error {
		if _transformInsert_nounRegister == nil {
			var err error
			_transformInsert_nounRegister, err = loadWords(tc.Params["anchorDataPath"])
			if err != nil {
				return fmt.Errorf("cannot load nouns: %w", err)
			}
		}

		words := strings.Fields(entry.Headline)
		indices := make([]int, 0)
		for i := 0; i < len(words); i++ {
			if _transformInsert_nounRegister[words[i]] {
				indices = append(indices, i)
			}
		}

		if len(indices) == 0 {
			return nil
		}

		x := rand.Intn(len(indices))
		index := indices[x]
		words = append(words[:index+1], words[index:]...)
		words[index] = tc.Params["text"]
		entry.Headline = strings.Join(words, " ")
		return nil
	},
}
