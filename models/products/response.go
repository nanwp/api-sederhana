package products

import "github.com/nanwp/api-sederhana/models/category"

type ProductResponse struct {
	SKU      string                    `json:"sku"`
	Name     string                    `json:"nama_produk"`
	Stock    int                       `json:"stock"`
	Price    int                       `json:"price"`
	Image    string                    `json:"image"`
	Category category.CategoryResponse `json:"kategori"`
}

type ProductCreateResponse struct {
	SKU        string `json:"sku"`
	Name       string `json:"nama_produk"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	CategoryId int    `json:"kategori_id"`
}
