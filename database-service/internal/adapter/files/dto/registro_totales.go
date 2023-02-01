package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type RegistroTotales struct {
	FechaPagoAju          string `json:"fecha_pago_aju,omitempty"`
	Descripcion           string `json:"descripcion,omitempty"`
	ImporteTotal          string `json:"importe_total,omitempty"`
	ImporteTotalEnDolares string `json:"importe_total_en_dolares,omitempty"`
	MontoImputOtor        string `json:"monto_imput_otor,omitempty"`
	MontoImputDol         string `json:"monto_imput_dol,omitempty"`
}

func (h RegistroTotales) toDomain() *contenido_resumen.RegistroTotales {
	r := contenido_resumen.NewRegistroTotales()
	r.SetFechaPagoAju(h.FechaPagoAju)
	r.SetDescripcion(h.Descripcion)
	r.SetImporteTotal(h.ImporteTotal)
	r.SetImporteTotalEnDolares(h.ImporteTotalEnDolares)
	r.SetMontoImputOtor(h.MontoImputOtor)
	r.SetMontoImputDol(h.MontoImputDol)
	return r
}

func registroTotalesFromDomain(h *contenido_resumen.RegistroTotales) *RegistroTotales {
	if h == nil {
		return nil
	}
	return &RegistroTotales{
		FechaPagoAju:          h.FechaPagoAju(),
		Descripcion:           h.Descripcion(),
		ImporteTotal:          h.ImporteTotal(),
		ImporteTotalEnDolares: h.ImporteTotalEnDolares(),
		MontoImputOtor:        h.MontoImputOtor(),
		MontoImputDol:         h.MontoImputDol(),
	}
}
