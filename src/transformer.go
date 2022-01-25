package main

import "errors"

type Transformer struct {
	transform func(tc TransformationConfig, entry *Entry) error
}

func transformData(server *NewsServer, transformationConfig TransformationConfig, entries []Entry) error {
	ttype := transformationConfig.TransformationType
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
		err := transformer.transform(transformationConfig, &entries[i])
		if err != nil {
			return err
		}
	}

	return nil
}
