package mastercard

type Mensaje struct {
	Resumen *ResumenAEnviar `json:"resumen,omitempty"`
	OpCount int             `json:"op_count,omitempty"`
	Id      string          `json:"id,omitempty"`
}
