package route

import (
	"strings"
)

func DoReply(request string) string {
	request = strings.TrimSpace(request)
	query := strings.Fields(request)

	return query[0]
}
