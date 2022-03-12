package helper

type BookRequest struct {
	Title    string
	Author   string
	Stock    uint8
	Price    float32
	Discount uint8
}

type BookResponse struct {
	Id       string
	Title    string
	Author   string
	Stock    uint8
	Price    float32
	Discount uint8
}

//type UpdateBookRequest struct {
//}
//
//type UpdateBookResponse struct {
//}
