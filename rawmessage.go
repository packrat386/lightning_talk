package main

import (
	"encoding/json"
	"fmt"
)

var data = `{
    "payments": [
        {
            "id": "12345",
            "amount": "3.50",
            "type": "ach",
            "payment_details": {
                "routing_number": "123456789",
                "account_number": "987654321"
            }
        },
        {
            "id": "12345",
            "amount": "3.50",
            "type": "credit_card",
            "payment_details": {
                "card_number": "1111222233334444",
                "expiration": "0618"
            }
        }
    ]
}`

// STRUCTS OMIT
type PaymentCollection struct {
	Payments []*Payment `json:"payments"`
}

type Payment struct {
	ID     string          `json:"id"`
	Amount string          `json:"amount"`
	Type   string          `json:"type"`
	Detail json.RawMessage `json:"payment_details"`
}

// END STRUCTS OMIT

// INTERNALS OMIT

type achDetail struct {
	RoutingNumber string `json:"routing_number"`
	AccountNumber string `json:"account_number"`
}

func (a *achDetail) Execute() {
	fmt.Printf("Executing ach payment\nrouting_number: %s\naccount_number: %s\n", a.RoutingNumber, a.AccountNumber)
}

type cardDetail struct {
	CardNumber string `json:"card_number"`
	Expiration string `json:"expiration"`
}

func (c *cardDetail) Execute() {
	fmt.Printf("Executing cc payment\ncard_number: %s\nexpiration: %s\n", c.CardNumber, c.Expiration)
}

// END INTERNALS OMIT

func main() {
	pmts := new(PaymentCollection)
	err := json.Unmarshal([]byte(data), pmts)
	if err != nil {
		panic(err)
	}

	for _, p := range pmts.Payments {
		p.Execute()
	}
}

// USAGE OMIT

type Executor interface {
	Execute()
}

func (p Payment) Execute() {
	var internal Executor
	if p.Type == "ach" {
		internal = new(achDetail)
	} else if p.Type == "credit_card" {
		internal = new(cardDetail)
	} else {
		panic("unrecognized type")
	}

	err := json.Unmarshal(p.Detail, internal)
	if err != nil {
		panic(err)
	}
	internal.Execute()
}

// END USAGE OMIT
