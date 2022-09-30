package wallets

type Configs struct {
	TransactionServiceAddress string `yaml:"transaction_service_address" json:"transaction_service_address" env:"TRANSACTION_SERVICE_ADDRESS"`
}
