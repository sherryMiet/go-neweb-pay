package neweb_pay

import (
	"github.com/sony/sonyflake"
	"strconv"
)

func PtrNilString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func GenSonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return ""
	}
	return strconv.FormatUint(id, 16)
}
