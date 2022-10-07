package youtube

import "strings"

func parseToTsQuery(str string) string {
	strs := strings.Split(str, " ")

	return strings.Join(strs, " & ")
}
