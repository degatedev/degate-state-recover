package entity

type AccountUpdate struct {
	AccountID string
	AccountValue *AccountValue
}

type AccountUpdatePack struct {
	AccountUpdateList []*AccountUpdate
}

type TransferToAccount struct {
	AccountID string
	Owner string
}