package contenido_resumen

type DatosMesAnterior struct {
	fechaCierre      string
	vencimiento      string
	saldoEnDolares   string
	saldo            string
	saldoTotal       string
	pagoMinimo       string
	compras          string
	comprasEnDolares string
}

func (d *DatosMesAnterior) SetFechaCierre(fechaCierre string) {
	d.fechaCierre = fechaCierre
}

func (d *DatosMesAnterior) SetVencimiento(vencimiento string) {
	d.vencimiento = vencimiento
}

func (d *DatosMesAnterior) SetSaldoEnDolares(saldoEnDolares string) {
	d.saldoEnDolares = saldoEnDolares
}

func (d *DatosMesAnterior) SetSaldo(saldo string) {
	d.saldo = saldo
}

func (d *DatosMesAnterior) SetSaldoTotal(saldoTotal string) {
	d.saldoTotal = saldoTotal
}

func (d *DatosMesAnterior) SetPagoMinimo(pagoMinimo string) {
	d.pagoMinimo = pagoMinimo
}

func (d *DatosMesAnterior) SetCompras(compras string) {
	d.compras = compras
}

func (d *DatosMesAnterior) SetComprasEnDolares(comprasEnDolares string) {
	d.comprasEnDolares = comprasEnDolares
}

func NewDatosMesAnterior() *DatosMesAnterior {
	return &DatosMesAnterior{}
}

func (d DatosMesAnterior) FechaCierre() string {
	return d.fechaCierre
}

func (d DatosMesAnterior) Vencimiento() string {
	return d.vencimiento
}

func (d DatosMesAnterior) SaldoEnDolares() string {
	return d.saldoEnDolares
}

func (d DatosMesAnterior) Saldo() string {
	return d.saldo
}

func (d DatosMesAnterior) SaldoTotal() string {
	return d.saldoTotal
}

func (d DatosMesAnterior) PagoMinimo() string {
	return d.pagoMinimo
}

func (d DatosMesAnterior) Compras() string {
	return d.compras
}

func (d DatosMesAnterior) ComprasEnDolares() string {
	return d.comprasEnDolares
}
