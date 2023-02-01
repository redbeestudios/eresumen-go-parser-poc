package domain

type Indice struct {
	fechaCierre   string
	marca         string
	path          string
	month         string
	year          string
	modeloResumen string
}

func (r *Indice) FechaCierre() string {
	return r.fechaCierre
}

func (r *Indice) SetFechaCierre(fechaCierre string) {
	r.fechaCierre = fechaCierre
}

func (r *Indice) Marca() string {
	return r.marca
}

func (r *Indice) SetMarca(marca string) {
	r.marca = marca
}

func (r *Indice) Path() string {
	return r.path
}

func (r *Indice) SetPath(path string) {
	r.path = path
}

func (r *Indice) Month() string {
	return r.month
}

func (r *Indice) SetMonth(month string) {
	r.month = month
}

func (r *Indice) Year() string {
	return r.year
}

func (r *Indice) SetYear(year string) {
	r.year = year
}

func (r *Indice) ModeloResumen() string {
	return r.modeloResumen
}

func (r *Indice) SetModeloResumen(modeloResumen string) {
	r.modeloResumen = modeloResumen
}

func NewIndice() *Indice {
	return &Indice{}
}
