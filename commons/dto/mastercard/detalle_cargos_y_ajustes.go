package mastercard

type DetalleCargosYAjustes struct {
	CodigoMov       string `json:"codigo_mov,omitempty"`
	SubcodMov       string `json:"subcod_mov,omitempty"`
	CodFinanc       string `json:"cod_financ,omitempty"`
	FecOper         string `json:"fec_oper,omitempty"`
	Cupon           string `json:"cupon,omitempty"`
	NroComercio     string `json:"nro_comercio,omitempty"`
	Detalle         string `json:"detalle,omitempty"`
	Importe         string `json:"importe,omitempty"`
	ImporteDol      string `json:"importe_dol,omitempty"`
	NroAutoriz      string `json:"nro_autoriz,omitempty"`
	Concepto        string `json:"concepto,omitempty"`
	FecPresen       string `json:"fec_presen,omitempty"`
	MarcaCF         string `json:"marca_cf,omitempty"`
	CodMoneda       string `json:"cod_moneda,omitempty"`
	CapitalAdeudado string `json:"capital_adeudado,omitempty"`
	InteresCuota    string `json:"interes_cuota,omitempty"`
	TnaPr           string `json:"tna_pr,omitempty"`
	Vigencia        string `json:"vigencia,omitempty"`
}
