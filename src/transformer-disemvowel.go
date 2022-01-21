package main

import "regexp"

var re *regexp.Regexp

func disemvowelTransform(entry *Entry) error {
	if re == nil {
		re = regexp.MustCompile("[aeiou]")
	}

	entry.Headline = re.ReplaceAllString(entry.Headline, "")
	return nil
}

var DisemvowelTransformer = Transformer{
	transform: disemvowelTransform,
}
