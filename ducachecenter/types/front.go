package types

type FrontCustomerActivateCacheCenter struct {
	CustomerId   int    `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Status       string `json:"status"`
	UsageMode    string `json:"usage_mode"`
	TimesLimit   int    `json:"times_limit"`
}
