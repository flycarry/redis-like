package route

import (
	"github.com/flycarry/redis-like/storage"
	"strings"
)

func DoReply(request string) string {
	request = strings.TrimSpace(request)
	query := strings.Fields(request)
	result, err := storage.GetResult(query...)
	if err != nil {
		return err.Error()
	} else {
		return result
	}
}
