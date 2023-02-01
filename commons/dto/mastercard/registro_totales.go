package mastercard

type RegistroTotales struct {
	FechaPagoAju          string `json:"fecha_pago_aju,omitempty"`
	Descripcion           string `json:"descripcion,omitempty"`
	ImporteTotal          string `json:"importe_total,omitempty"`
	ImporteTotalEnDolares string `json:"importe_total_en_dolares,omitempty"`
	MontoImputOtor        string `json:"monto_imput_otor,omitempty"`
	MontoImputDol         string `json:"monto_imput_dol,omitempty"`
}
