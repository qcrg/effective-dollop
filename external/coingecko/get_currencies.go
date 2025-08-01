package coingecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type nested struct {
	Usd float64 `json:"usd"`
}

type crr_resp = map[string]nested
type Currencies map[string]nested

func GetCurrencies(symbols []string) (Currencies, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/price?vs_currencies=%s&symbols=%s&precision=4",
		"usd",
		strings.Join(symbols, ","),
	)
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf(
			"Request to coingecko with error: %s (%s)",
			string(data),
			res.Status,
		)
	}
	crrsp := make(crr_resp)
	dec := json.NewDecoder(res.Body)
	if !dec.More() {
		return nil, errors.New("Reponse is empty or is invalid")
	}
	err = dec.Decode(&crrsp)
	if err != nil {
		return nil, err
	}
	return crrsp, nil
}
