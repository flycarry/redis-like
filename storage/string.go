package storage

type str struct {
	data string
}

func NewStr(s string) *str {
	return &str{s}
}
func (s *str) Value() string {
	return s.data
}
