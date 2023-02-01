package domain

type DatosDeAutorizado struct {
	texto       string
	nombre      string
	sldo        string
	sldoDol     string
	movimientos []*DetalleCargosYAjustes
	numero      string
}

func (d *DatosDeAutorizado) SetTexto(texto string) {
	d.texto = texto
}

func (d *DatosDeAutorizado) SetNombre(nombre string) {
	d.nombre = nombre
}

func (d *DatosDeAutorizado) SetSldo(sldo string) {
	d.sldo = sldo
}

func (d *DatosDeAutorizado) SetSldoDol(sldoDol string) {
	d.sldoDol = sldoDol
}

func (d *DatosDeAutorizado) SetNumero(numero string) {
	d.numero = numero
}

func NewDatosDeAutorizado() *DatosDeAutorizado {
	return &DatosDeAutorizado{}
}

func (d DatosDeAutorizado) Texto() string {
	return d.texto
}

func (d DatosDeAutorizado) Nombre() string {
	return d.nombre
}

func (d DatosDeAutorizado) Saldo() string {
	return d.sldo
}

func (d DatosDeAutorizado) SaldoEnDolares() string {
	return d.sldoDol
}

func (d DatosDeAutorizado) Movimientos() []*DetalleCargosYAjustes {
	return d.movimientos
}

func (d DatosDeAutorizado) SetMovimientos(movimientos []*DetalleCargosYAjustes) {
	d.movimientos = movimientos
}

func (d DatosDeAutorizado) AddMovimientos(movimiento *DetalleCargosYAjustes) {
	d.movimientos = append(d.movimientos, movimiento)
}

func (d DatosDeAutorizado) Numero() string {
	return d.numero
}
