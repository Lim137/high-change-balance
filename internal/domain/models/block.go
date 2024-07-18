package models

type Block struct {
	Number       string `json:"number"`
	Transactions []struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
	} `json:"transactions"`
}
