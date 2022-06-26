package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat, 2)
	scanner := bufio.NewScanner(r)

	var jsonParser fastjson.Parser

	for scanner.Scan() {
		val, err := jsonParser.ParseBytes(scanner.Bytes())
		if err != nil {
			return nil, err
		}

		email := string(val.GetStringBytes("Email"))
		if strings.Contains(email, domain) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}

	return result, nil
}
