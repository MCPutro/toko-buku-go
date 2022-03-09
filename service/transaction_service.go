package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
)

type TransactionService interface {
	BuyBook(ctx context.Context, request helper.TransactionRequest) (*entity.Transaction, error)
}
