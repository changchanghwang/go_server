package presentation

type CreateDto struct {
	ProductId string `json:"productId"`
	Stock     int    `json:"stock"`
}
