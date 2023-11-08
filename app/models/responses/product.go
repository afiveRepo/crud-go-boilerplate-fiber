package responses

type ProductResponse struct {
	ProductID          uint64  `json:"product_id"`
	ProductName        string  `json:"product_name"`
	PoductPrice        float64 `json:"product_price"`
	ProductStok        float64 `json:"product_stok"`
	ProductInformation string  `json:"product_information"`
}
