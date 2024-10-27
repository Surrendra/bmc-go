package helpers

import (
	"regexp"
	"strings"
)

type publicHelper struct{}

func NewPublicHelper() *publicHelper {
	return &publicHelper{}
}

type PublicHelper interface {
	MakeSlugFromString(string) string
}

func (h publicHelper) MakeSlugFromString(text string) string {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^a-z0-9\s-]`)
	text = re.ReplaceAllString(text, "")

	re = regexp.MustCompile(`[\s-]+`)
	text = re.ReplaceAllString(text, "-")

	text = strings.Trim(text, "-")
	return text
}
