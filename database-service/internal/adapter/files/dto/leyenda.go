package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type Leyenda struct {
	RubroLeyendaInf string `json:"rubro_leyenda_inf,omitempty"`
	LeyendaInf      string `json:"leyenda_inf,omitempty"`
	McaFacturacion  string `json:"mca_facturacion,omitempty"`
}

func (h Leyenda) toDomain() *contenido_resumen.Leyenda {
	l := contenido_resumen.NewLeyenda()
	l.SetRubroLeyendaInf(h.RubroLeyendaInf)
	l.SetLeyendaInf(h.LeyendaInf)
	l.SetMcaFacturacion(h.McaFacturacion)
	return l
}

func leyendaFromDomain(h *contenido_resumen.Leyenda) *Leyenda {
	if h == nil {
		return nil
	}
	return &Leyenda{
		RubroLeyendaInf: h.RubroLeyendaInf(),
		LeyendaInf:      h.LeyendaInf(),
		McaFacturacion:  h.McaFacturacion(),
	}
}
