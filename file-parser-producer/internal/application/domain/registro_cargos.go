package domain

type RegistroCargos struct {
	descripcion    string
	importe        string
	importeDolares string
}

func (r *RegistroCargos) SetDescripcion(descripcion string) {
	r.descripcion = descripcion
}

func (r *RegistroCargos) SetImporte(importe string) {
	r.importe = importe
}

func (r *RegistroCargos) SetImporteDolares(importeDolares string) {
	r.importeDolares = importeDolares
}

func NewRegistroCargos() *RegistroCargos {
	return &RegistroCargos{}
}

func (r RegistroCargos) Descripcion() string {
	return r.descripcion
}

func (r RegistroCargos) Importe() string {
	return r.importe
}

func (r RegistroCargos) ImporteDolares() string {
	return r.importeDolares
}
