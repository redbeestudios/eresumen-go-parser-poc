package domain

type Mail struct {
	nombre string
	fecha  string
	url    string
	email  string
}

func (m *Mail) Url() string {
	return m.url
}

func (m *Mail) Email() string {
	return m.email
}

func (m *Mail) SetEmail(email string) {
	m.email = email
}

func (m *Mail) Nombre() string {
	return m.nombre
}

func (m *Mail) SetNombre(nombre string) {
	m.nombre = nombre
}

func (m *Mail) Fecha() string {
	return m.fecha
}

func (m *Mail) SetFecha(fecha string) {
	m.fecha = fecha
}

func (m *Mail) CodigoVinculacion() string {
	return m.url
}

func (m *Mail) SetUrl(url string) {
	m.url = url
}

func NewMail(nombre string, fecha string, url string, email string) *Mail {
	return &Mail{nombre: nombre, fecha: fecha, url: url, email: email}
}
