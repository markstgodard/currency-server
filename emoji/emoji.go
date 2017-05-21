package emoji

import "fmt"

var currencies = map[string]string{
	"USD": "💵",
	"GBP": "💷",
	"CAD": "💩",
}

func Convert(amt float64, code string) (string, error) {
	v, ok := currencies[code]
	if !ok {
		return "🤷", fmt.Errorf("Unknown currency code: %s", code)
	}

	return v, nil
}
