package entity

type AccountInfos struct {
	Accounts  map[string]*AccountInfo `json:"accounts"`
}

type AccountInfo struct {
	AccountID string `json:"accountID"`
	Tokens map[string]*TokenInfo `json:"tokens"`
}

type TokenInfo struct {
	Balance string `json:"balance"`
}

func NewAccounts() *AccountInfos {
	return &AccountInfos{
		Accounts: make(map[string]*AccountInfo),
	}
}

func NewAccount() *AccountInfo {
	return &AccountInfo{
		Tokens: make(map[string]*TokenInfo),
	}
}