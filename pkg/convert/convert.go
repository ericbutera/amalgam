package convert

import "strconv"

func ParseUInt(s string) (uint, error) {
	// TODO: this is probably an anti-pattern, but i am not typing this out everywhere
	// TODO: research generics
	number, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(number), nil
}
