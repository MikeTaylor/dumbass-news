package main

import "regexp"

var re *regexp.Regexp

var DisemvowelTransformer = Transformer{
	transform: func(tc TransformationConfig, entry *Entry) error {
		if re == nil {
			re = regexp.MustCompile("[aeiou]")
		}

		entry.Headline = re.ReplaceAllString(entry.Headline, "")
		return nil
	},
}
