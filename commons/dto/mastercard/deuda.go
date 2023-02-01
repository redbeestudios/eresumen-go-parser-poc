package mastercard

type Deuda struct {
	DenMes    string `json:"den_mes,omitempty"`
	DenMoneda string `json:"den_moneda,omitempty"`
	DnvMes    string `json:"dnv_mes,omitempty"`
}
