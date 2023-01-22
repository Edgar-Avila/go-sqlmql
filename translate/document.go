package translate

import (
	"fmt"
	"strings"
)

const indent = 2

func makePair[T any](key string, val T, level int, f func(T) (string, error)) (string, error) {
	v, err := f(val)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s: %s", strings.Repeat(" ", level*indent), key, v), nil
}

func makeDoc[T any](keys []string, vals []T, level int, f func(T) (string, error)) (string, error) {
	lk, lv := len(keys), len(vals)
	if lv != lk {
		return "", fmt.Errorf("%v values expected, %v given", lk, lv)
	}
	docPairs := make([]string, len(vals))
	for i, val := range vals {
		pair, err := makePair(keys[i], val, level+1, f)
		if err != nil {
			return "", err
		}
		docPairs[i] = pair
	}
	return fmt.Sprintf("{\n%s\n}", strings.Join(docPairs, ",\n")), nil
}
