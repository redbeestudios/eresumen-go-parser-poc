package mastercard

type DatosDeAutorizado struct {
	Texto          string                   `json:"texto,omitempty"`
	Nombre         string                   `json:"nombre,omitempty"`
	Saldo          string                   `json:"saldo,omitempty"`
	SaldoEnDolares string                   `json:"saldo_en_dolares,omitempty"`
	Movimientos    []*DetalleCargosYAjustes `json:"movimientos,omitempty"`
	Numero         string                   `json:"numero,omitempty"`
}
