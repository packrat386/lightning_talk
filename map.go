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
	Payments []Payment `json"payments"`
}

type Payment struct {
	ID             string                 `json:"id"`
	Amount         string                 `json:"amount"`
	Type           string                 `json:"type"`
	PaymentDetails map[string]interface{} `json:"payment_details"`
}

// END STRUCTS OMIT

// MAIN OMIT
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

// END MAIN OMIT

// USAGE OMIT

func (p Payment) Execute() {
	if p.Type == "ach" {
		routingNumber, ok := p.PaymentDetails["routing_number"].(string)
		if !ok {
			panic("routing number not a string")
		}

		accountNumber, ok := p.PaymentDetails["account_number"].(string)
		if !ok {
			panic("account number not a string")
		}

		fmt.Printf(
			"Executing ach payment\nrouting_number: %s\naccount_number: %s\n",
			routingNumber,
			accountNumber,
		)
	} else if p.Type == "credit_card" { // ...
		// END USAGE OMIT
		cardNumber, ok := p.PaymentDetails["card_number"].(string)
		if !ok {
			panic("routing number not a string")
		}

		expiration, ok := p.PaymentDetails["expiration"].(string)
		if !ok {
			panic("account number not a string")
		}

		fmt.Printf(
			"Executing credit card payment\ncard_payment: %s\nexpiration: %s\n",
			cardNumber,
			expiration,
		)
	} else {
		panic("unrecognized type")
	}
}
