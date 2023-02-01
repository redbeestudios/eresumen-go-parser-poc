package contenido_resumen

type Leyenda struct {
	rubroLeyendaInf string
	leyendaInf      string
	mcaFacturacion  string
}

func (l *Leyenda) SetRubroLeyendaInf(rubroLeyendaInf string) {
	l.rubroLeyendaInf = rubroLeyendaInf
}

func (l *Leyenda) SetLeyendaInf(leyendaInf string) {
	l.leyendaInf = leyendaInf
}

func (l *Leyenda) SetMcaFacturacion(mcaFacturacion string) {
	l.mcaFacturacion = mcaFacturacion
}

func NewLeyenda() *Leyenda {
	return &Leyenda{}
}

func (l Leyenda) RubroLeyendaInf() string {
	return l.rubroLeyendaInf
}

func (l Leyenda) LeyendaInf() string {
	return l.leyendaInf
}

func (l Leyenda) McaFacturacion() string {
	return l.mcaFacturacion
}
