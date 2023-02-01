package mastercard

type IndiceResumen struct {
	FechaCierre   string `json:"fecha_cierre,omitempty"`
	Marca         string `json:"marca,omitempty"`
	Path          string `json:"path,omitempty"`
	Month         string `json:"month,omitempty"`
	Year          string `json:"year,omitempty"`
	ModeloResumen string `json:"modelo_resumen,omitempty"`
}
