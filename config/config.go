package config

type AliConfig struct {
	AppID        string `mapstructure:"app_id"`
	PrivateKey   string `mapstructure:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url"`
	ReturnURL    string `mapstructure:"return_url"`
	ProductCode  string `mapstructure:"product_code"`
	Subject      string `mapstructure:"subject"`
}
