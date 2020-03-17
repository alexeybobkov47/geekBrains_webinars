package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func CheckoutDummy(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	switch key {
	case "42":
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"status": 200, "balance": 100500}`)
	case "100500":
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"status": 400, "err": "bad_balance"}`)
	case "__broken_json":
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"status": 400`) // broken json
	case "__internal_error":
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type CheckoutResult struct {
	Status  int
	Balance int
	Err     string
}

type Cart struct {
	PaymentApiURL string
}

func (c *Cart) Checkout(id string) (*CheckoutResult, error) {
	url := c.PaymentApiURL + "?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := new(CheckoutResult)
	if err = json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}
