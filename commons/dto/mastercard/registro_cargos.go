package mastercard

type RegistroCargos struct {
	Descripcion    string `json:"descripcion,omitempty"`
	Importe        string `json:"importe,omitempty"`
	ImporteDolares string `json:"importe_dolares,omitempty"`
}
