package base62

import "strings"

const DICTIONARY = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(id int64) string {
	dictionary := []byte(DICTIONARY)
	idBytes := []byte{}
	for {
		if id <= 0 {
			break
		}
		i := id % 62
		v := []byte{dictionary[i]}
		idBytes = append(v, idBytes...)
		id = (id - i) / 62
	}
	return string(idBytes)
}

func Decode(id string) int64 {
	idBytes := []byte(id)
	idInt := strings.IndexByte(DICTIONARY, idBytes[0])
	if idInt == -1 {
		//return 0, errors.New("Base62Decode error, id:" + id)
	}
	for i := 1; i < len(idBytes); i++ {
		i := strings.IndexByte(DICTIONARY, idBytes[i])
		idInt = idInt*62 + i
	}
	return int64(idInt)
}
