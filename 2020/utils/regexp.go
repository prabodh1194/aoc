package utils

import (
	"regexp"
)

func RegexpToMap(r *regexp.Regexp, txt string) map[string]string {
	matches := r.FindStringSubmatch(txt) // matches is [][]string
	names := r.SubexpNames()

	var res map[string]string
	res = make(map[string]string)

	for i, match := range matches {
		if names[i] == "" {
			continue
		}
		res[names[i]] = match
	}

	return res
}

func MultiRegexToMap(r *regexp.Regexp, txt string) []map[string]string {
	matches := r.FindAllStringSubmatch(txt, -1)
	names := r.SubexpNames()

	var res []map[string]string

	for _, match := range matches {
		var ressubmatch map[string]string
		ressubmatch = make(map[string]string)

		for i, submatch := range match {
			if names[i] == "" || submatch == "" {
				continue
			}
			ressubmatch[names[i]] = submatch
		}

		if len(ressubmatch) > 0 {
			res = append(res, ressubmatch)
		}
	}

	return res
}
