package domain

import . "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type Resumen struct {
	header                *Clave
	datosFiliatorios      *DatosFiliatorios
	datosLiquidacionSocio *DatosLiquidacionSocio
	datosDeAutorizado     []*DatosDeAutorizado
	datosMesAnterior      *DatosMesAnterior
	registroTotales       []*RegistroTotales
	datosMesActual        *DatosMesActual
	detalleCargosYAjustes []*DetalleCargosYAjustes
	registroCargos        []*RegistroCargos
	leyendas              []*Leyenda
	limites               []*Limite
	tasas                 []*Tasa
	deudasNoVencidas      []*Deuda
}

func NewContenidoResumen() *Resumen {
	return &Resumen{}
}

func (r *Resumen) Clave() *Clave {
	return r.header
}

func (r *Resumen) SetClave(header *Clave) *Resumen {
	r.header = header
	return r
}

func (r *Resumen) DatosFiliatorios() *DatosFiliatorios {
	return r.datosFiliatorios
}

func (r *Resumen) SetDatosFiliatorios(datosFiliatorios *DatosFiliatorios) *Resumen {
	r.datosFiliatorios = datosFiliatorios
	return r
}

func (r *Resumen) DatosLiquidacionSocio() *DatosLiquidacionSocio {
	return r.datosLiquidacionSocio
}

func (r *Resumen) SetDatosLiquidacionSocio(datosLiquidacionSocio *DatosLiquidacionSocio) *Resumen {
	r.datosLiquidacionSocio = datosLiquidacionSocio
	return r
}

func (r *Resumen) DatosDeAutorizado() []*DatosDeAutorizado {
	return r.datosDeAutorizado
}

func (r *Resumen) SetDatosDeAutorizado(datosDeAutorizado []*DatosDeAutorizado) *Resumen {
	r.datosDeAutorizado = datosDeAutorizado
	return r
}

func (r *Resumen) AddDatosDeAutorizado(datosDeAutorizado *DatosDeAutorizado) *Resumen {
	r.datosDeAutorizado = append(r.datosDeAutorizado, datosDeAutorizado)
	return r
}

func (r *Resumen) DatosMesAnterior() *DatosMesAnterior {
	return r.datosMesAnterior
}

func (r *Resumen) SetDatosMesAnterior(datosMesAnterior *DatosMesAnterior) *Resumen {
	r.datosMesAnterior = datosMesAnterior
	return r
}

func (r *Resumen) RegistroTotales() []*RegistroTotales {
	return r.registroTotales
}

func (r *Resumen) SetRegistroTotales(registroTotales []*RegistroTotales) *Resumen {
	r.registroTotales = registroTotales
	return r
}

func (r *Resumen) AddRegistroTotales(registroTotales *RegistroTotales) *Resumen {
	r.registroTotales = append(r.registroTotales, registroTotales)
	return r
}

func (r *Resumen) DatosMesActual() *DatosMesActual {
	return r.datosMesActual
}

func (r *Resumen) SetDatosMesActual(datosMesActual *DatosMesActual) *Resumen {
	r.datosMesActual = datosMesActual
	return r
}

func (r *Resumen) DetalleCargosYAjustes() []*DetalleCargosYAjustes {
	return r.detalleCargosYAjustes
}

func (r *Resumen) SetDetalleCargosYAjustes(detalleCargosYAjustes []*DetalleCargosYAjustes) *Resumen {
	r.detalleCargosYAjustes = detalleCargosYAjustes
	return r
}

func (r *Resumen) AddDetalleCargosYAjustes(detalleCargosYAjustes *DetalleCargosYAjustes) *Resumen {
	r.detalleCargosYAjustes = append(r.detalleCargosYAjustes, detalleCargosYAjustes)
	return r
}

func (r *Resumen) RegistroCargos() []*RegistroCargos {
	return r.registroCargos
}

func (r *Resumen) SetRegistroCargos(registroCargos []*RegistroCargos) *Resumen {
	r.registroCargos = registroCargos
	return r
}

func (r *Resumen) AddRegistroCargos(registroCargos *RegistroCargos) *Resumen {
	r.registroCargos = append(r.registroCargos, registroCargos)
	return r
}

func (r *Resumen) Leyendas() []*Leyenda {
	return r.leyendas
}

func (r *Resumen) SetLeyendas(leyendas []*Leyenda) *Resumen {
	r.leyendas = leyendas
	return r
}

func (r *Resumen) AddLeyenda(leyenda *Leyenda) *Resumen {
	r.leyendas = append(r.leyendas, leyenda)
	return r
}

func (r *Resumen) Limites() []*Limite {
	return r.limites
}

func (r *Resumen) SetLimites(limites []*Limite) *Resumen {
	r.limites = limites
	return r
}

func (r *Resumen) AddLimite(limite *Limite) *Resumen {
	r.limites = append(r.limites, limite)
	return r
}

func (r *Resumen) Tasas() []*Tasa {
	return r.tasas
}

func (r *Resumen) SetTasas(tasas []*Tasa) *Resumen {
	r.tasas = tasas
	return r
}

func (r *Resumen) AddTasa(tasa *Tasa) *Resumen {
	r.tasas = append(r.tasas, tasa)
	return r
}

func (r *Resumen) DeudasNoVencidas() []*Deuda {
	return r.deudasNoVencidas
}

func (r *Resumen) SetDeudasNoVencidas(deudasNoVencidas []*Deuda) *Resumen {
	r.deudasNoVencidas = deudasNoVencidas
	return r
}

func (r *Resumen) AddDeudaNoVencida(deudaNoVencida *Deuda) *Resumen {
	r.deudasNoVencidas = append(r.deudasNoVencidas, deudaNoVencida)
	return r
}
