package utils

import "regexp"

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
