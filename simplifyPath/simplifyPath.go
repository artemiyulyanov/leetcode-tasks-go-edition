package simplifyPath

import (
	"strings"
)

func SimplifyPath(path string) string {
	splitted := filterArray(strings.Split(path, "/"), func(value string) bool {
		return len(value) > 0
	})

	new_path := ""

	for _, value := range splitted {
		if value == ".." {
			splitted_new_path := strings.Split(new_path, "/")
			new_path = strings.Join(splitted_new_path[:len(splitted_new_path)-1], "/")
		} else {
			if value == "." {
				continue
			}

			new_path += "/" + value
		}
	}

	if len(new_path) == 0 {
		return "/"
	}

	return new_path
}

func filterArray[T any](array []T, condition func(T) bool) []T {
	filtered := []T{}

	for _, value := range array {
		if condition(value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}
