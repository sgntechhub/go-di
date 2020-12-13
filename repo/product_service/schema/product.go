package schema

type Product struct {
	Id         uint64 `json:"product_id" bson:"product_id" redis:"product_id"`
	Name       string `json:"name" bson:"name" redis:"name"`
	MerchantId uint64 `json:"merchant_id" bson:"merchant_id" redis:"merchant_id"`
}
