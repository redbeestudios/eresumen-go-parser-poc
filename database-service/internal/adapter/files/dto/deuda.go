package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type Deuda struct {
	DenMes    string `json:"den_mes,omitempty"`
	DenMoneda string `json:"den_moneda,omitempty"`
	DnvMes    string `json:"dnv_mes,omitempty"`
}

func (h Deuda) toDomain() *contenido_resumen.Deuda {
	d := contenido_resumen.NewDeuda()

	d.SetDenMes(h.DenMes)
	d.SetDenMoneda(h.DenMoneda)
	d.SetDnvMes(h.DnvMes)

	return d
}

func deudaFromDomain(h *contenido_resumen.Deuda) *Deuda {
	if h == nil {
		return nil
	}
	return &Deuda{
		DenMes:    h.DenMes(),
		DenMoneda: h.DenMoneda(),
		DnvMes:    h.DnvMes(),
	}
}
