package domain

type Destinatario struct {
	nombre          string
	direccion       string
	localidad       string
	codigoPostal    string
	email           string
	numeroCuenta    string
	numeroDocumento string
}

func NewDestinatario(nombre string, direccion string, localidad string, codigoPostal string, email string, numeroCuenta string, numeroDocumento string) *Destinatario {
	return &Destinatario{nombre: nombre, direccion: direccion, localidad: localidad, codigoPostal: codigoPostal, email: email, numeroCuenta: numeroCuenta, numeroDocumento: numeroDocumento}
}

func (d *Destinatario) Email() string {
	return d.email
}

func (d *Destinatario) SetEmail(email string) {
	d.email = email
}

func (d *Destinatario) Nombre() string {
	return d.nombre
}

func (d *Destinatario) Direccion() string {
	return d.direccion
}

func (d *Destinatario) Localidad() string {
	return d.localidad
}

func (d *Destinatario) CodigoPostal() string {
	return d.codigoPostal
}

func (d *Destinatario) NumeroCuenta() string {
	return d.numeroCuenta
}

func (d *Destinatario) NumeroDocumento() string {
	return d.numeroDocumento
}
