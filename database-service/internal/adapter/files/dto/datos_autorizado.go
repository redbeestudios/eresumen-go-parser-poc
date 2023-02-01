package files

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"
	sf "github.com/sa-/slicefunk"
)

type DatosDeAutorizado struct {
	Texto          string                   `json:"texto,omitempty"`
	Nombre         string                   `json:"nombre,omitempty"`
	Saldo          string                   `json:"saldo,omitempty"`
	SaldoEnDolares string                   `json:"saldo_en_dolares,omitempty"`
	Movimientos    []*DetalleCargosYAjustes `json:"movimientos,omitempty"`
	Numero         string                   `json:"numero,omitempty"`
}

func (h DatosDeAutorizado) toDomain() *contenido_resumen.DatosDeAutorizado {
	movimientos := sf.Map[*DetalleCargosYAjustes, *contenido_resumen.DetalleCargosYAjustes](
		h.Movimientos, func(item *DetalleCargosYAjustes) *contenido_resumen.DetalleCargosYAjustes {
			return item.toDomain()
		})

	datos := contenido_resumen.NewDatosDeAutorizado()

	datos.SetTexto(h.Texto)
	datos.SetNombre(h.Nombre)
	datos.SetSldo(h.Saldo)
	datos.SetSldoDol(h.SaldoEnDolares)
	datos.SetNumero(h.Numero)

	datos.SetMovimientos(movimientos)

	return datos
}

func datosDeAutorizadoFromDomain(h *contenido_resumen.DatosDeAutorizado) *DatosDeAutorizado {
	if h == nil {
		return nil
	}
	return &DatosDeAutorizado{
		Texto:          h.Texto(),
		Nombre:         h.Nombre(),
		Saldo:          h.Saldo(),
		SaldoEnDolares: h.SaldoEnDolares(),
		Movimientos: sf.Map[*contenido_resumen.DetalleCargosYAjustes, *DetalleCargosYAjustes](h.Movimientos(), func(item *contenido_resumen.DetalleCargosYAjustes) *DetalleCargosYAjustes {
			return detalleCargosYAjustesFromDomain(item)
		}),
		Numero: h.Numero(),
	}
}
