package dto

type NewTransactionResponse struct {
	TransactionID string  `json:"transaction-id" xml:"transaction-id"`
	LegalBalance  float64 `json:"legal-balance" xml:"legal-balance"`
	AvailBalance  float64 `json:"available-balance" xml:"available-balance"`
}
