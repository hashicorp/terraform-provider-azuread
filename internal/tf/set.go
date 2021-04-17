package tf

import (
	"hash/crc32"
)

func SetFuncHashId(val interface{}) int {
	m := val.(map[string]interface{})
	i, ok := m["id"]
	if !ok {
		return 0
	}
	id, ok := i.(string)
	if !ok {
		return 0
	}
	v := int(crc32.ChecksumIEEE([]byte(id)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
