package category

type CategoryCreate struct {
	Name string `json:"nama_kategori" binding:"required"`
}

type CategoryUpdate struct {
	Name string `json:"nama_kategori"`
}
