package api

import (
	"context"

	desc "github.com/volnistii11/accounting_service/balance/pkg/api/transfer/v1"
	servicepb "github.com/volnistii11/accounting_service/balance/pkg/api/transfer/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	servicepb.UnimplementedTransferServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateMoneyTransfer(ctx context.Context, info *desc.CreateMoneyTransferRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
