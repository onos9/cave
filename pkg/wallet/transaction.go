package wallet

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func getTransaction(text string) (map[string]string, error) {
	var temp string
	var istd bool
	var isVal bool
	payment := make(map[string]string)
	tkn := html.NewTokenizer(strings.NewReader(text))

	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			if err := tkn.Err(); err == io.EOF {
				return payment, nil
			}
			return nil, tkn.Err()
		case tt == html.StartTagToken:
			istd = tkn.Token().Data == "td"
		case tt == html.TextToken:
			d := strings.TrimSpace(tkn.Token().Data)
			d = strings.Replace(d, "\n", "", -1)

			if isVal && d != ":" && d != "" {
				field := strings.Replace(temp, "Transaction ", "", -1)
				payment[field] = d
			}
			if istd && d != ":" && d != "" {
				temp = d
				isVal = !isVal
			}
			istd = false
		}
	}
}
