package model

type Order struct {
	UserUUID  string         `json:"user_uuid"`
	Salads    []OrderProduct `json:"salads"`
	Garnishes []OrderProduct `json:"garnishes"`
	Meats     []OrderProduct `json:"meats"`
	Soups     []OrderProduct `json:"soups"`
	Drinks    []OrderProduct `json:"drinks"`
	Desserts  []OrderProduct `json:"desserts"`
}

type OrderProduct struct {
	Count       int    `json:"count"`
	ProductUUID string `json:"product_uuid"`
}
