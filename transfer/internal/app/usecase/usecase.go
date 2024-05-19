package usecase

import (
	"context"

	"github.com/volnistii11/accounting_service/transfer/internal/app/models"
	desc "github.com/volnistii11/accounting_service/transfer/pkg/api/transfer/v1"
)

type UseCase struct {
	client desc.TransferClient
}

func NewUseCase(client desc.TransferClient) *UseCase {
	return &UseCase{
		client: client,
	}
}

func (u *UseCase) CreateMoneyTransfer(ctx context.Context, info models.MoneyTransfer) error {
	_, err := u.client.CreateMoneyTransfer(ctx,
		&desc.CreateMoneyTransferRequest{
			Info: convertMsg(info),
		})
	if err != nil {
		return err
	}

	return nil
}

func convertMsg(msg models.MoneyTransfer) *desc.MoneyTransferInfo {
	return &desc.MoneyTransferInfo{
		RequestId:  uint32(msg.RequestID),
		FromUserId: uint32(msg.FromUserID),
		ToUserId:   uint32(msg.ToUserID),
		Sum:        uint32(msg.Sum),
	}
}
