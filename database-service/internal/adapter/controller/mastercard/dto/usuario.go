package mastercard

type Usuario struct {
	Nombre          string `json:"nombre,omitempty"`
	Direccion       string `json:"direccion,omitempty"`
	Localidad       string `json:"localidad,omitempty"`
	CodigoPostal    string `json:"codigo_postal,omitempty"`
	Email           string `json:"email,omitempty"`
	NumeroCuenta    string `json:"numero_cuenta,omitempty"`
	NumeroDocumento string `json:"numero_documento,omitempty"`
}
