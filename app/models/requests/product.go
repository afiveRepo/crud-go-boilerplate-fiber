package requests

type InputProduct struct {
	ProductName        string  `json:"product_name"`
	PoductPrice        float64 `json:"product_price"`
	ProductStok        float64 `json:"product_stok"`
	ProductInformation string  `json:"product_information"`
}
type UpdateProduct struct {
	ProductUID         string  `json:"product_uid"`
	ProductName        string  `json:"product_name"`
	PoductPrice        float64 `json:"product_price"`
	ProductStok        float64 `json:"product_stok"`
	ProductInformation string  `json:"product_information"`
}
