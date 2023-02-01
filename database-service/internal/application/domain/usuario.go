package domain

type Usuario struct {
	nombre          string
	direccion       string
	localidad       string
	codigoPostal    string
	email           string
	numeroCuenta    string
	numeroDocumento string
}

func (u *Usuario) Nombre() string {
	return u.nombre
}

func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}

func (u *Usuario) Direccion() string {
	return u.direccion
}

func (u *Usuario) SetDireccion(direccion string) {
	u.direccion = direccion
}

func (u *Usuario) Localidad() string {
	return u.localidad
}

func (u *Usuario) SetLocalidad(localidad string) {
	u.localidad = localidad
}

func (u *Usuario) CodigoPostal() string {
	return u.codigoPostal
}

func (u *Usuario) SetCodigoPostal(codigoPostal string) {
	u.codigoPostal = codigoPostal
}

func (u *Usuario) Email() string {
	return u.email
}

func (u *Usuario) SetEmail(email string) {
	u.email = email
}

func (u *Usuario) NumeroCuenta() string {
	return u.numeroCuenta
}

func (u *Usuario) SetNumeroCuenta(numeroCuenta string) {
	u.numeroCuenta = numeroCuenta
}

func (u *Usuario) NumeroDocumento() string {
	return u.numeroDocumento
}

func (u *Usuario) SetNumeroDocumento(numeroDocumento string) {
	u.numeroDocumento = numeroDocumento
}

func NewUsuario() *Usuario {
	return &Usuario{}
}
