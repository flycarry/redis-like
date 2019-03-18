package route

import (
	"github.com/flycarry/redis-like/storage"
	"strings"
)

var data *storage.Data

func init() {
	data = storage.NewData()
}

func DoReply(request string) string {
	request = strings.TrimSpace(request)
	query := strings.Fields(request)
	if len(query) <= 1 {
		return "not a query"
	}
	method := query[0]
	switch method {
	case "get":
		response, err := data.GetString(query[1])
		if err != 0 {
			return "err"
		}
		if len(query) != 2 {
			return "parm not good"
		}
		return response
	case "set":
		err := data.SetString(query[1], query[2])
		if err != 0 {
			return "err"
		}
		if len(query) != 3 {
			return "parm not good"
		}
		return "good"
	}
	return ""
}
