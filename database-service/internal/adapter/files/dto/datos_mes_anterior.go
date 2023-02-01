package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type DatosMesAnterior struct {
	FechaCierre      string `json:"fecha_cierre,omitempty"`
	Vencimiento      string `json:"vencimiento,omitempty"`
	SaldoEnDolares   string `json:"saldo_en_dolares,omitempty"`
	Saldo            string `json:"saldo,omitempty"`
	SaldoTotal       string `json:"saldo_total,omitempty"`
	PagoMinimo       string `json:"pago_minimo,omitempty"`
	Compras          string `json:"compras,omitempty"`
	ComprasEnDolares string `json:"compras_en_dolares,omitempty"`
}

func (h DatosMesAnterior) toDomain() *contenido_resumen.DatosMesAnterior {
	m := contenido_resumen.NewDatosMesAnterior()

	m.SetFechaCierre(h.FechaCierre)
	m.SetVencimiento(h.Vencimiento)
	m.SetSaldoEnDolares(h.SaldoEnDolares)
	m.SetSaldo(h.Saldo)
	m.SetSaldoTotal(h.SaldoTotal)
	m.SetPagoMinimo(h.PagoMinimo)
	m.SetCompras(h.Compras)
	m.SetComprasEnDolares(h.ComprasEnDolares)

	return m
}

func datosMesAnteriorFromDomain(h *contenido_resumen.DatosMesAnterior) *DatosMesAnterior {
	if h == nil {
		return nil
	}
	return &DatosMesAnterior{
		FechaCierre:      h.FechaCierre(),
		Vencimiento:      h.Vencimiento(),
		SaldoEnDolares:   h.SaldoEnDolares(),
		Saldo:            h.Saldo(),
		SaldoTotal:       h.SaldoTotal(),
		PagoMinimo:       h.PagoMinimo(),
		Compras:          h.Compras(),
		ComprasEnDolares: h.ComprasEnDolares(),
	}
}
