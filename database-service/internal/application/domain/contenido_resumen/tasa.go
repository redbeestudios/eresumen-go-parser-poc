package contenido_resumen

type Tasa struct {
	descripcion   string
	tasa          string
	tasaEnDolares string
}

func (t *Tasa) SetDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *Tasa) SetTasa(tasa string) {
	t.tasa = tasa
}

func (t *Tasa) SetTasaEnDolares(tasaEnDolares string) {
	t.tasaEnDolares = tasaEnDolares
}

func NewTasa() *Tasa {
	return &Tasa{}
}

func (t Tasa) Descripcion() string {
	return t.descripcion
}

func (t Tasa) Tasa() string {
	return t.tasa
}

func (t Tasa) TasaEnDolares() string {
	return t.tasaEnDolares
}
