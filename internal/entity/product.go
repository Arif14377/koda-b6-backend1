package entity

type Products struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
}

type RequestProducts struct {
	Name        string `json:"name" example:"Bakwan"`
	Description string `json:"description" example:"Bakwan panas berlimpah sayur dengan kenikmatan gigitan cabe rawit pedas"`
	Qty         int    `json:"qty" example:"100"`
	Price       int    `json:"price" example:"3000"`
}

type ProductVariant struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	AddPrice int    `json:"addPrice"`
}

type ProductSize struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	AddPrice int    `json:"addPrice"`
}

type ProductImage struct {
	Id        int    `json:"id"`
	ProductId int    `json:"productId"`
	Path      string `json:"path"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductCategory struct {
	Id         int `json:"id"`
	ProductId  int `json:"productId"`
	CategoryId int `json:"categoryId"`
}
