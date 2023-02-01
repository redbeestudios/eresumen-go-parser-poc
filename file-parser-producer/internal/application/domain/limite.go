package domain

type Limite struct {
	descripcion string
	limite      string
}

func (l *Limite) SetDescripcion(descripcion string) {
	l.descripcion = descripcion
}

func (l *Limite) SetLimite(limite string) {
	l.limite = limite
}

func NewLimite() *Limite {
	return &Limite{}
}

func (l Limite) Descripcion() string {
	return l.descripcion
}

func (l Limite) Limite() string {
	return l.limite
}
