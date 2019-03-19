package route

import (
	"github.com/flycarry/redis-like/storage"
	"strings"
)

func DoReply(request string) string {
	request = strings.TrimSpace(request)
	query := strings.Fields(request)
	f := storage.GetFunc(query...)
	return f(query[1:]...)
}
