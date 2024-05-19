package models

type MoneyTransfer struct {
	ID         int64 `json:"id"`
	RequestID  int64 `json:"request_id"`
	FromUserID int64 `json:"from_user_id"`
	ToUserID   int64 `json:"to_user_id"`
	Sum        int64 `json:"sum"`
}
