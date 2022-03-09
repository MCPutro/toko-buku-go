package helper

type BookRequest struct {
	Title    string
	Author   string
	Stock    uint8
	Price    float32
	Discount float32
}

type BookResponse struct {
	ID       uint8
	Title    string
	Author   string
	Stock    uint8
	Price    float32
	Discount float32
}

//type UpdateBookRequest struct {
//}
//
//type UpdateBookResponse struct {
//}
