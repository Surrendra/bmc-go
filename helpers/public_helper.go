package helpers

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"regexp"
	"strconv"
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

func (h publicHelper) generatePaginationUrl(c *gin.Context, pageSize int, pageIndex int) map[string]string {
	u := url.URL{
		Scheme: "https",
		Host:   c.Request.Host,
		Path:   c.Request.URL.Path,
	}
	query := u.Query()
	query.Set("page_size", strconv.Itoa(pageSize))
	query.Set("page_index", strconv.Itoa(pageIndex))
	u.RawQuery = query.Encode()
	return map[string]string{
		"current_page":  u.String(),
		"next_page":     u.String(),
		"previous_page": u.String(),
	}
}
