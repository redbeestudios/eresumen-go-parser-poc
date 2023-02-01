package domain

type RegistroTotales struct {
	fechaPagoAju          string
	descripcion           string
	importeTotal          string
	importeTotalEnDolares string
	montoImputOtor        string
	montoImputDol         string
}

func (r *RegistroTotales) SetFechaPagoAju(fechaPagoAju string) {
	r.fechaPagoAju = fechaPagoAju
}

func (r *RegistroTotales) SetDescripcion(descripcion string) {
	r.descripcion = descripcion
}

func (r *RegistroTotales) SetImporteTotal(importeTotal string) {
	r.importeTotal = importeTotal
}

func (r *RegistroTotales) SetImporteTotalEnDolares(importeTotalEnDolares string) {
	r.importeTotalEnDolares = importeTotalEnDolares
}

func (r *RegistroTotales) SetMontoImputOtor(montoImputOtor string) {
	r.montoImputOtor = montoImputOtor
}

func (r *RegistroTotales) SetMontoImputDol(montoImputDol string) {
	r.montoImputDol = montoImputDol
}

func NewRegistroTotales() *RegistroTotales {
	return &RegistroTotales{}
}

func (r RegistroTotales) FechaPagoAju() string {
	return r.fechaPagoAju
}

func (r RegistroTotales) Descripcion() string {
	return r.descripcion
}

func (r RegistroTotales) ImporteTotal() string {
	return r.importeTotal
}

func (r RegistroTotales) ImporteTotalEnDolares() string {
	return r.importeTotalEnDolares
}

func (r RegistroTotales) MontoImputOtor() string {
	return r.montoImputOtor
}

func (r RegistroTotales) MontoImputDol() string {
	return r.montoImputDol
}
