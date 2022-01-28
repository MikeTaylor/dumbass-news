package main

import "errors"

type Transformer struct {
	transform func(server *NewsServer, tc transformationConfig, entry *Entry) error
}

func transformData(server *NewsServer, tc transformationConfig, entries []Entry) error {
	ttype := tc.TransformationType
	var transformer Transformer
	switch ttype {
	case "disemvowel":
		transformer = DisemvowelTransformer
	case "insert":
		transformer = InsertTransformer
	// more cases here
	default:
		return errors.New("unsupported transformer-type: " + ttype)
	}

	for i := 0; i < len(entries); i++ {
		err := transformer.transform(server, tc, &entries[i])
		if err != nil {
			return err
		}
	}

	return nil
}
