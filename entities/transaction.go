package entities

type PayLoad struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type ServerRespons struct {
	Tx_Hash string `json:"tx_hash"`
}

type ServerResponsStatus struct {
	Tx_Status string `json:"tx_status"`
}
