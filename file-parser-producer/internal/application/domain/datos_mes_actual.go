package domain

type DatosMesActual struct {
	fechaCierre        string
	saldo              string
	saldoEnDolares     string
	pagoMinimo         string
	proximoVencimiento string
	proximaFechaCierre string
}

func (d *DatosMesActual) SetFechaCierre(fechaCierre string) {
	d.fechaCierre = fechaCierre
}

func (d *DatosMesActual) SetSaldo(saldo string) {
	d.saldo = saldo
}

func (d *DatosMesActual) SetSaldoEnDolares(saldoEnDolares string) {
	d.saldoEnDolares = saldoEnDolares
}

func (d *DatosMesActual) SetPagoMinimo(pagoMinimo string) {
	d.pagoMinimo = pagoMinimo
}

func (d *DatosMesActual) SetProximoVencimiento(proximoVencimiento string) {
	d.proximoVencimiento = proximoVencimiento
}

func (d *DatosMesActual) SetProximaFechaCierre(proximaFechaCierre string) {
	d.proximaFechaCierre = proximaFechaCierre
}

func NewDatosMesActual() *DatosMesActual {
	return &DatosMesActual{}
}

func (d DatosMesActual) FechaCierre() string {
	return d.fechaCierre
}

func (d DatosMesActual) Saldo() string {
	return d.saldo
}

func (d DatosMesActual) SaldoEnDolares() string {
	return d.saldoEnDolares
}

func (d DatosMesActual) PagoMinimo() string {
	return d.pagoMinimo
}

func (d DatosMesActual) ProximoVencimiento() string {
	return d.proximoVencimiento
}

func (d DatosMesActual) ProximaFechaCierre() string {
	return d.proximaFechaCierre
}
