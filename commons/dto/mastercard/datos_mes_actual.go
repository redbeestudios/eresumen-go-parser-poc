package mastercard

type DatosMesActual struct {
	FechaCierre        string `json:"fecha_cierre,omitempty"`
	Saldo              string `json:"saldo,omitempty"`
	SaldoEnDolares     string `json:"saldo_en_dolares,omitempty"`
	PagoMinimo         string `json:"pago_minimo,omitempty"`
	ProximoVencimiento string `json:"proximo_vencimiento,omitempty"`
	ProximaFechaCierre string `json:"proxima_fecha_cierre,omitempty"`
}
