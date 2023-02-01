package mastercard

type Tasa struct {
	Descripcion   string `json:"descripcion,omitempty"`
	Tasa          string `json:"tasa,omitempty"`
	TasaEnDolares string `json:"tasa_en_dolares,omitempty"`
}
