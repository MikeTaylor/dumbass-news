package main

import "regexp"

// private
var _transformDisemvowel_re *regexp.Regexp

var DisemvowelTransformer = Transformer{
	transform: func(server *NewsServer, tc transformationConfig, entry *Entry) error {
		if _transformDisemvowel_re == nil {
			_transformDisemvowel_re = regexp.MustCompile("[aeiou]")
		}

		entry.Headline = _transformDisemvowel_re.ReplaceAllString(entry.Headline, "")
		return nil
	},
}
