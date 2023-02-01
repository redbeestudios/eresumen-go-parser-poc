package serialization

import (
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"
	sf "github.com/sa-/slicefunk"
)

func Deserialize(message *dto.ResumenAEnviar) (*domain.Usuario, *domain.Indice, *domain.Resumen) {
	contenidoResumen := resumenToDomain(message.Resumen)

	indice := indiceToDomain(message)

	usuario := usuarioToDomain(message)

	return usuario, indice, contenidoResumen
}

func usuarioToDomain(h *dto.ResumenAEnviar) *domain.Usuario {
	usuario := domain.NewUsuario()
	d := h.Destinatario
	usuario.SetNombre(d.Nombre)
	usuario.SetLocalidad(d.Localidad)
	usuario.SetCodigoPostal(d.CodigoPostal)
	usuario.SetNumeroCuenta(d.NumeroCuenta)
	usuario.SetDireccion(d.Direccion)
	usuario.SetEmail(d.Email)
	usuario.SetNumeroDocumento(d.NumeroDocumento)

	return usuario

}

func indiceToDomain(h *dto.ResumenAEnviar) *domain.Indice {
	resumen := domain.NewIndice()
	resumen.SetFechaCierre(h.FechaCierre)
	resumen.SetModeloResumen(h.ModeloResumen)
	resumen.SetMarca(h.Marca)
	resumen.SetMonth(h.Periodo)
	resumen.SetYear(h.FechaCierre)

	return resumen
}

func resumenToDomain(h *dto.Resumen) *domain.Resumen {
	resumenToReturn := domain.NewContenidoResumen()
	resumenToReturn.SetClave(claveToDomain(h.Clave))
	if h.DatosFiliatorios != nil {
		resumenToReturn.SetDatosFiliatorios(datosFiliatoriosToDomain(h.DatosFiliatorios))
	}
	if h.DatosLiquidacionSocio != nil {
		resumenToReturn.SetDatosLiquidacionSocio(datosLiquidacionSocioToDomain(h.DatosLiquidacionSocio))
	}
	if h.DatosMesAnterior != nil {
		resumenToReturn.SetDatosMesAnterior(datosMesAnteriorToDomain(h.DatosMesAnterior))
	}
	if h.DatosMesActual != nil {
		resumenToReturn.SetDatosMesActual(datosMesActualToDomain(h.DatosMesActual))
	}
	resumenToReturn.SetDatosDeAutorizado(sf.Map[*dto.DatosDeAutorizado, *contenido_resumen.DatosDeAutorizado](h.DatosDeAutorizado,
		func(item *dto.DatosDeAutorizado) *contenido_resumen.DatosDeAutorizado {
			return datosDeAutorizadoToDomain(item)
		}))
	resumenToReturn.SetRegistroTotales(sf.Map[*dto.RegistroTotales, *contenido_resumen.RegistroTotales](h.RegistroTotales,
		func(item *dto.RegistroTotales) *contenido_resumen.RegistroTotales {
			return registroTotalesToDomain(item)
		}))
	resumenToReturn.SetDetalleCargosYAjustes(sf.Map[*dto.DetalleCargosYAjustes, *contenido_resumen.DetalleCargosYAjustes](h.DetalleCargosYAjustes,
		func(item *dto.DetalleCargosYAjustes) *contenido_resumen.DetalleCargosYAjustes {
			return detalleCargosYAjustestoDomain(item)
		}))
	resumenToReturn.SetRegistroCargos(sf.Map[*dto.RegistroCargos, *contenido_resumen.RegistroCargos](h.RegistroCargos,
		func(item *dto.RegistroCargos) *contenido_resumen.RegistroCargos {
			return registroCargosToDomain(item)
		}))
	resumenToReturn.SetLeyendas(sf.Map[*dto.Leyenda, *contenido_resumen.Leyenda](h.Leyendas,
		func(item *dto.Leyenda) *contenido_resumen.Leyenda { return leyendaToDomain(item) }))
	resumenToReturn.SetLimites(sf.Map[*dto.Limite, *contenido_resumen.Limite](h.Limites,
		func(item *dto.Limite) *contenido_resumen.Limite { return limiteToDomain(item) }))
	resumenToReturn.SetTasas(sf.Map[*dto.Tasa, *contenido_resumen.Tasa](h.Tasas,
		func(item *dto.Tasa) *contenido_resumen.Tasa { return tasasToDomain(item) }))
	resumenToReturn.SetDeudasNoVencidas(sf.Map[*dto.Deuda, *contenido_resumen.Deuda](h.DeudasNoVencidas,
		func(item *dto.Deuda) *contenido_resumen.Deuda { return deudaToDomain(item) }))
	return resumenToReturn

}

func claveToDomain(h *dto.Clave) *contenido_resumen.Clave {
	d := contenido_resumen.NewClave()
	d.SetNroRegistro(h.NroRegistro)
	d.SetNroSecuencia(h.NroSecuencia)
	d.SetNroOrden(h.NroOrden)
	d.SetEntidad(h.Entidad)
	d.SetSucursal(h.Sucursal)
	d.SetGrupoCuentaCorriente(h.GrupoCuentaCorriente)
	d.SetCalleYNumero(h.CalleYNumero)
	d.SetPiso(h.Piso)
	d.SetDepartamento(h.Departamento)
	d.SetCodigoPostal(h.CodigoPostal)
	d.SetRetencion(h.Retencion)
	d.SetVtoActual(h.VtoActual)
	d.SetGrupoAfinidad(h.GrupoAfinidad)
	d.SetCodProductoMarca(h.CodProductoMarca)
	d.SetModeloLiquidacion(h.ModeloLiquidacion)
	d.SetNroCuenta(h.NroCuenta)
	d.SetNroAutorizado(h.NroAutorizado)
	d.SetEResumen(h.EResumen)
	d.SetCantidadResumen(h.CantidadResumen)

	return d

}

func datosDeAutorizadoToDomain(h *dto.DatosDeAutorizado) *contenido_resumen.DatosDeAutorizado {
	movimientos := sf.Map[*dto.DetalleCargosYAjustes, *contenido_resumen.DetalleCargosYAjustes](
		h.Movimientos, func(item *dto.DetalleCargosYAjustes) *contenido_resumen.DetalleCargosYAjustes {
			return detalleCargosYAjustestoDomain(item)
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

func datosFiliatoriosToDomain(h *dto.DatosFiliatorios) *contenido_resumen.DatosFiliatorios {
	d := contenido_resumen.NewDatosFiliatorios()

	d.SetNombreSocio(h.NombreSocio)
	d.SetNroCtaEmpresa(h.NroCtaEmpresa)
	d.SetNombreCtaEmpresa(h.NombreCtaEmpresa)
	d.SetDomicilioCodPcia(h.DomicilioCodPcia)
	d.SetDomicilioAdic1(h.DomicilioAdic1)
	d.SetDomicilioAdic2(h.DomicilioAdic2)
	d.SetLocalidad(h.Localidad)
	d.SetPaisOtorg(h.PaisOtorg)
	d.SetTipoTarjeta(h.TipoTarjeta)
	d.SetDenomTipoTarj(h.DenomTipoTarj)

	return d

}

func datosLiquidacionSocioToDomain(h *dto.DatosLiquidacionSocio) *contenido_resumen.DatosLiquidacionSocio {
	d := contenido_resumen.NewDatosLiquidacionSocio()

	d.SetCuit(h.Cuit)
	d.SetCategoriaIva(h.CategoriaIva)
	d.SetDescripcionCategoriaIva(h.DescripcionCategoriaIva)
	d.SetSbmTasas(h.SbmTasas)
	d.SetSbmGtosSV(h.SbmGtosSV)
	d.SetSbmParam(h.SbmParam)
	d.SetDvSocio(h.DvSocio)
	d.SetMarcaCBU(h.MarcaCBU)
	d.SetCbu(h.Cbu)
	d.SetDvUnifSuc(h.DvUnifSuc)
	d.SetLeyDebit(h.LeyDebit)
	d.SetTipCta(h.TipCta)
	d.SetNroCtaDebito(h.NroCtaDebito)
	d.SetTipPago(h.TipPago)
	d.SetTipoCtaCte(h.TipoCtaCte)
	d.SetTarjetaTitular(h.TarjetaTitular)
	d.SetRegPagYAjSant(h.RegPagYAjSant)
	d.SetRegDetalleYMsg(h.RegDetalleYMsg)
	d.SetCantidadHojas(h.CantidadHojas)
	d.SetTipoProducto(h.TipoProducto)
	d.SetComprobante(h.Comprobante)

	return d
}

func datosMesActualToDomain(h *dto.DatosMesActual) *contenido_resumen.DatosMesActual {
	m := contenido_resumen.NewDatosMesActual()

	m.SetFechaCierre(h.FechaCierre)
	m.SetSaldo(h.Saldo)
	m.SetSaldoEnDolares(h.SaldoEnDolares)
	m.SetPagoMinimo(h.PagoMinimo)
	m.SetProximoVencimiento(h.ProximoVencimiento)
	m.SetProximaFechaCierre(h.ProximaFechaCierre)

	return m
}

func datosMesAnteriorToDomain(h *dto.DatosMesAnterior) *contenido_resumen.DatosMesAnterior {
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

func detalleCargosYAjustestoDomain(h *dto.DetalleCargosYAjustes) *contenido_resumen.DetalleCargosYAjustes {

	d := contenido_resumen.NewDetalleCargosYAjustes()
	d.SetCodigoMov(h.CodigoMov)
	d.SetSubcodMov(h.SubcodMov)
	d.SetCodFinanc(h.CodFinanc)
	d.SetFecOper(h.FecOper)
	d.SetCupon(h.Cupon)
	d.SetNroComercio(h.NroComercio)
	d.SetDetalle(h.Detalle)
	d.SetImporte(h.Importe)
	d.SetImporteDol(h.ImporteDol)
	d.SetNroAutoriz(h.NroAutoriz)
	d.SetConcepto(h.Concepto)
	d.SetFecPresen(h.FecPresen)
	d.SetMarcaCF(h.MarcaCF)
	d.SetCodMoneda(h.CodMoneda)
	d.SetCapitalAdeudado(h.CapitalAdeudado)
	d.SetInteresCuota(h.InteresCuota)
	d.SetTnaPr(h.TnaPr)
	d.SetVigencia(h.Vigencia)

	return d
}

func deudaToDomain(h *dto.Deuda) *contenido_resumen.Deuda {
	d := contenido_resumen.NewDeuda()

	d.SetDenMes(h.DenMes)
	d.SetDenMoneda(h.DenMoneda)
	d.SetDnvMes(h.DnvMes)

	return d
}

func leyendaToDomain(h *dto.Leyenda) *contenido_resumen.Leyenda {
	l := contenido_resumen.NewLeyenda()
	l.SetRubroLeyendaInf(h.RubroLeyendaInf)
	l.SetLeyendaInf(h.LeyendaInf)
	l.SetMcaFacturacion(h.McaFacturacion)
	return l
}

func limiteToDomain(h *dto.Limite) *contenido_resumen.Limite {
	l := contenido_resumen.NewLimite()
	l.SetLimite(h.Limite)
	l.SetDescripcion(h.Descripcion)

	return l
}

func registroCargosToDomain(h *dto.RegistroCargos) *contenido_resumen.RegistroCargos {
	r := contenido_resumen.NewRegistroCargos()
	r.SetDescripcion(h.Descripcion)
	r.SetImporte(h.Importe)
	r.SetImporteDolares(h.ImporteDolares)
	return r
}

func registroTotalesToDomain(h *dto.RegistroTotales) *contenido_resumen.RegistroTotales {
	r := contenido_resumen.NewRegistroTotales()
	r.SetImporteTotal(h.ImporteTotal)
	r.SetMontoImputDol(h.MontoImputDol)
	r.SetDescripcion(h.Descripcion)
	r.SetImporteTotalEnDolares(h.ImporteTotalEnDolares)
	r.SetMontoImputOtor(h.MontoImputOtor)
	r.SetFechaPagoAju(h.FechaPagoAju)

	return r
}

func tasasToDomain(item *dto.Tasa) *contenido_resumen.Tasa {
	t := contenido_resumen.NewTasa()
	t.SetDescripcion(item.Descripcion)
	t.SetTasa(item.Tasa)
	t.SetTasaEnDolares(item.TasaEnDolares)

	return t
}
