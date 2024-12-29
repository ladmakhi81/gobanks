package types

type CreateAccountReqBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TransferMoneyReqBody struct {
	ToAccount int     `json:"toAccount"`
	Amount    float64 `json:"amount"`
}
