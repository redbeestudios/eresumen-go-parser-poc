package mastercard

type ResumenAEnviar struct {
	Resumen           *Resumen      `json:"resumen,omitempty"`
	ModeloResumen     string        `json:"modelo_resumen,omitempty"`
	Destinatario      *Destinatario `json:"destinatario,omitempty"`
	FechaCierre       string        `json:"fecha_cierre,omitempty"`
	Marca             string        `json:"marca,omitempty"`
	Periodo           string        `json:"periodo,omitempty"`
	CodigoVinculacion string        `json:"codigo_vinculacion,omitempty"`
}
