package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type DatosMesActual struct {
	FechaCierre        string `json:"fecha_cierre,omitempty"`
	Saldo              string `json:"saldo,omitempty"`
	SaldoEnDolares     string `json:"saldo_en_dolares,omitempty"`
	PagoMinimo         string `json:"pago_minimo,omitempty"`
	ProximoVencimiento string `json:"proximo_vencimiento,omitempty"`
	ProximaFechaCierre string `json:"proxima_fecha_cierre,omitempty"`
}

func (h DatosMesActual) toDomain() *contenido_resumen.DatosMesActual {
	m := contenido_resumen.NewDatosMesActual()

	m.SetFechaCierre(h.FechaCierre)
	m.SetSaldo(h.Saldo)
	m.SetSaldoEnDolares(h.SaldoEnDolares)
	m.SetPagoMinimo(h.PagoMinimo)
	m.SetProximoVencimiento(h.ProximoVencimiento)
	m.SetProximaFechaCierre(h.ProximaFechaCierre)

	return m
}

func datosMesActualFromDomain(h *contenido_resumen.DatosMesActual) *DatosMesActual {
	if h == nil {
		return nil
	}
	return &DatosMesActual{
		FechaCierre:        h.FechaCierre(),
		Saldo:              h.Saldo(),
		SaldoEnDolares:     h.SaldoEnDolares(),
		PagoMinimo:         h.PagoMinimo(),
		ProximoVencimiento: h.ProximoVencimiento(),
		ProximaFechaCierre: h.ProximaFechaCierre(),
	}
}
