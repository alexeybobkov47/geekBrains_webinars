package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestCase struct {
	ID      string
	Result  *CheckoutResult
	IsError bool
}

func TestCartCheckout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CheckoutDummy))

	for caseNum, item := range createCases() {
		c := &Cart{
			PaymentApiURL: ts.URL,
		}

		result, err := c.Checkout(item.ID)
		if err != nil && !item.IsError {
			t.Errorf("[%d] unexpected error: %#v", caseNum, err)
		}

		if err == nil && item.IsError {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}

		if !reflect.DeepEqual(item.Result, result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", caseNum, item.Result, result)
		}
	}
	ts.Close()
}

func createCases() []TestCase {
	return []TestCase{
		{
			ID: "42",
			Result: &CheckoutResult{
				Status:  200,
				Balance: 100500,
			},
			IsError: false,
		},
		{
			ID: "100500",
			Result: &CheckoutResult{
				Status: 400,
				Err:    "bad_balance",
			},
			IsError: false,
		},
		{
			ID:      "__broken_json",
			IsError: true,
		},
		{
			ID:      "__internal_error",
			IsError: true,
		},
	}
}
