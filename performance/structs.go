package performance

type Request struct {
	TransactionID string `json:"transaction_id"`
	PlayLoad      []int  `json:"playload"`
}

type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression    string `json:"exp"`
}
