package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/helper"
)

type TransactionService interface {
	BuyBook(ctx context.Context, request helper.TransactionRequest) (*helper.TransactionResponse, error)
	FindByCustomerEmail(ctx context.Context, email string) (*[]helper.TransactionResponse, error)
}
