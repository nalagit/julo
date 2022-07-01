package db

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersToken struct {
	Username    string `json:"username"`
	Token       string `json:"token"`
	GeneratedAt string `json:"generated_at"`
	ExpiredAt   string `json:"expired_at"`
}

type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type WalletStatus struct {
	Id int `json:"id"`
}

type Wallet struct {
	Id        int    `json:"id"`
	OwnedBy   string `json:"owned_by"`
	Status    string `json:"status"`
	EnabledAt string `json:"enabled_at"`
	Balance   int    `json:"balance"`
}

type WalletDisable struct {
	Id         string `json:"id"`
	OwnedBy    string `json:"owned_by"`
	Status     string `json:"status"`
	DisabledAt string `json:"disabled_at"`
	Balance    int    `json:"balance"`
}

type Deposit struct {
	Id          int    `json:"id"`
	DepositBy   string `json:"deposited_by"`
	Status      string `json:"status"`
	DepositAt   string `json:"deposit_at"`
	Amount      int    `json:"amount"`
	ReferenceId string `json:"reference_id"`
}

type DepositRequest struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}
