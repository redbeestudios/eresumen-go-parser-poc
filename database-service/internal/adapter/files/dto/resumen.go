package files

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"
	sf "github.com/sa-/slicefunk"
)

type Resumen struct {
	Header                *Header                  `json:"header,omitempty"`
	DatosFiliatorios      *DatosFiliatorios        `json:"datos_filiatorios,omitempty"`
	DatosLiquidacionSocio *DatosLiquidacionSocio   `json:"datos_liquidacion_socio,omitempty"`
	DatosDeAutorizado     []*DatosDeAutorizado     `json:"datos_de_autorizado,omitempty"`
	DatosMesAnterior      *DatosMesAnterior        `json:"datos_mes_anterior,omitempty"`
	RegistroTotales       []*RegistroTotales       `json:"registro_totales,omitempty"`
	DatosMesActual        *DatosMesActual          `json:"datos_mes_actual,omitempty"`
	DetalleCargosYAjustes []*DetalleCargosYAjustes `json:"detalle_cargos_y_ajustes,omitempty"`
	RegistroCargos        []*RegistroCargos        `json:"registro_cargos,omitempty"`
	Leyendas              []*Leyenda               `json:"leyendas,omitempty"`
	Limites               []*Limite                `json:"limites,omitempty"`
	Tasas                 []*Tasa                  `json:"tasas,omitempty"`
	DeudasNoVencidas      []*Deuda                 `json:"deudas_no_vencidas,omitempty"`
}

func (h Resumen) ToDomain() *domain.Resumen {
	resumen := domain.NewContenidoResumen()
	resumen.SetClave(h.Header.toDomain())
	if h.DatosFiliatorios != nil {
		resumen.SetDatosFiliatorios(h.DatosFiliatorios.toDomain())
	}
	if h.DatosLiquidacionSocio != nil {
		resumen.SetDatosLiquidacionSocio(h.DatosLiquidacionSocio.toDomain())
	}
	if h.DatosMesAnterior != nil {
		resumen.SetDatosMesAnterior(h.DatosMesAnterior.toDomain())
	}
	if h.DatosMesActual != nil {
		resumen.SetDatosMesActual(h.DatosMesActual.toDomain())
	}
	resumen.SetDatosDeAutorizado(sf.Map[*DatosDeAutorizado, *contenido_resumen.DatosDeAutorizado](h.DatosDeAutorizado,
		func(item *DatosDeAutorizado) *contenido_resumen.DatosDeAutorizado { return item.toDomain() }))
	resumen.SetRegistroTotales(sf.Map[*RegistroTotales, *contenido_resumen.RegistroTotales](h.RegistroTotales,
		func(item *RegistroTotales) *contenido_resumen.RegistroTotales { return item.toDomain() }))
	resumen.SetDetalleCargosYAjustes(sf.Map[*DetalleCargosYAjustes, *contenido_resumen.DetalleCargosYAjustes](h.DetalleCargosYAjustes,
		func(item *DetalleCargosYAjustes) *contenido_resumen.DetalleCargosYAjustes { return item.toDomain() }))
	resumen.SetRegistroCargos(sf.Map[*RegistroCargos, *contenido_resumen.RegistroCargos](h.RegistroCargos,
		func(item *RegistroCargos) *contenido_resumen.RegistroCargos { return item.toDomain() }))
	resumen.SetLeyendas(sf.Map[*Leyenda, *contenido_resumen.Leyenda](h.Leyendas,
		func(item *Leyenda) *contenido_resumen.Leyenda { return item.toDomain() }))
	resumen.SetLimites(sf.Map[*Limite, *contenido_resumen.Limite](h.Limites,
		func(item *Limite) *contenido_resumen.Limite { return item.toDomain() }))
	resumen.SetTasas(sf.Map[*Tasa, *contenido_resumen.Tasa](h.Tasas,
		func(item *Tasa) *contenido_resumen.Tasa { return item.toDomain() }))
	resumen.SetDeudasNoVencidas(sf.Map[*Deuda, *contenido_resumen.Deuda](h.DeudasNoVencidas,
		func(item *Deuda) *contenido_resumen.Deuda { return item.toDomain() }))
	return resumen

}

func DtoFromDomain(h *domain.Resumen) *Resumen {
	return &Resumen{
		Header:                claveFromDomain(h.Clave()),
		DatosFiliatorios:      datosFiliatoriosFromDomain(h.DatosFiliatorios()),
		DatosLiquidacionSocio: datosLiquidacionSocioFromDomain(h.DatosLiquidacionSocio()),
		DatosDeAutorizado: sf.Map[*contenido_resumen.DatosDeAutorizado, *DatosDeAutorizado](h.DatosDeAutorizado(),
			func(item *contenido_resumen.DatosDeAutorizado) *DatosDeAutorizado {
				return datosDeAutorizadoFromDomain(item)
			}),
		DatosMesAnterior: datosMesAnteriorFromDomain(h.DatosMesAnterior()),
		RegistroTotales: sf.Map[*contenido_resumen.RegistroTotales, *RegistroTotales](h.RegistroTotales(),
			func(item *contenido_resumen.RegistroTotales) *RegistroTotales { return registroTotalesFromDomain(item) }),
		DatosMesActual: datosMesActualFromDomain(h.DatosMesActual()),
		DetalleCargosYAjustes: sf.Map[*contenido_resumen.DetalleCargosYAjustes, *DetalleCargosYAjustes](h.DetalleCargosYAjustes(),
			func(item *contenido_resumen.DetalleCargosYAjustes) *DetalleCargosYAjustes {
				return detalleCargosYAjustesFromDomain(item)
			}),
		RegistroCargos: sf.Map[*contenido_resumen.RegistroCargos, *RegistroCargos](h.RegistroCargos(),
			func(item *contenido_resumen.RegistroCargos) *RegistroCargos { return registroCargosFromDomain(item) }),
		Leyendas: sf.Map[*contenido_resumen.Leyenda, *Leyenda](h.Leyendas(),
			func(item *contenido_resumen.Leyenda) *Leyenda { return leyendaFromDomain(item) }),
		Limites: sf.Map[*contenido_resumen.Limite, *Limite](h.Limites(),
			func(item *contenido_resumen.Limite) *Limite { return limiteFromDomain(item) }),
		Tasas: sf.Map[*contenido_resumen.Tasa, *Tasa](h.Tasas(),
			func(item *contenido_resumen.Tasa) *Tasa { return tasaFromDomain(item) }),
		DeudasNoVencidas: sf.Map[*contenido_resumen.Deuda, *Deuda](h.DeudasNoVencidas(),
			func(item *contenido_resumen.Deuda) *Deuda { return deudaFromDomain(item) }),
	}
}
