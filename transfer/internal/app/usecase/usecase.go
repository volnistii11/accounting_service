package usecase

import (
	"context"

	"github.com/volnistii11/accounting_service/transfer/internal/app/models"
)

type UseCase struct {
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) CreateMoneyTransfer(ctx context.Context, info models.MoneyTransfer) error {
	return nil
}
