package lib

import "strings"

func ParseStringToArray(text string, parse string) []string {
	attr := strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	parts := strings.Split(attr, parse)
	attributes := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			attributes = append(attributes, p)
		}
	}
	return attributes
}

func SafeFirst(arr []string) string {
	if arr == nil || len(arr) == 0 {
		return ""
	}
	return arr[0]
}
