package mastercard

type DatosMesAnterior struct {
	FechaCierre      string `json:"fecha_cierre,omitempty"`
	Vencimiento      string `json:"vencimiento,omitempty"`
	SaldoEnDolares   string `json:"saldo_en_dolares,omitempty"`
	Saldo            string `json:"saldo,omitempty"`
	SaldoTotal       string `json:"saldo_total,omitempty"`
	PagoMinimo       string `json:"pago_minimo,omitempty"`
	Compras          string `json:"compras,omitempty"`
	ComprasEnDolares string `json:"compras_en_dolares,omitempty"`
}
