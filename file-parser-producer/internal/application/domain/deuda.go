package domain

type Deuda struct {
	denMes    string
	denMoneda string
	dnvMes    string
}

func (d *Deuda) SetDenMes(denMes string) {
	d.denMes = denMes
}

func (d *Deuda) SetDenMoneda(denMoneda string) {
	d.denMoneda = denMoneda
}

func (d *Deuda) SetDnvMes(dnvMes string) {
	d.dnvMes = dnvMes
}

func NewDeuda() *Deuda {
	return &Deuda{}
}

func (d Deuda) DenMes() string {
	return d.denMes
}

func (d Deuda) DenMoneda() string {
	return d.denMoneda
}

func (d Deuda) DnvMes() string {
	return d.dnvMes
}
