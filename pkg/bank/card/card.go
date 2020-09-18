package card

import (
	"bank/pkg/bank/types"
)

//PaymentSource fff
func PaymentSources(cards []types.Card) []types.PaymentSource {

	var data []types.PaymentSource

	for _, card := range cards {

		if card.Active == false {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		data = append(data, types.PaymentSource{Type: "cards", Number: string(card.PAN), Balance: card.Balance})
	}
	return data
}
