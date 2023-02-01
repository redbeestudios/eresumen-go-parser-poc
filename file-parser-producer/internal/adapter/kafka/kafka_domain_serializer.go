package kafka

import (
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	sf "github.com/sa-/slicefunk"
)

func Serialize(resumen *domain.ResumenAEnviar, opCount int, id string) *dto.Mensaje {
	resumenDto := resumenFromDomain(resumen.Resumen())
	destinatario := destinatarioFromDomain(resumen.Destinatario())

	resumenAEnviar := &dto.ResumenAEnviar{
		resumenDto,
		resumen.ModeloResumen(),
		destinatario,
		resumen.FechaEnvio(),
		resumen.Marca(),
		resumen.Periodo(),
		resumen.CodigoVinculacion(),
	}
	return &dto.Mensaje{
		Resumen: resumenAEnviar,
		OpCount: opCount,
		Id:      id,
	}
}

func claveFromDomain(h *domain.Clave) *dto.Clave {
	if h == nil {
		return nil
	}
	return &dto.Clave{
		NroRegistro:          h.NroRegistro(),
		NroSecuencia:         h.NroSecuencia(),
		NroOrden:             h.NroOrden(),
		Entidad:              h.Entidad(),
		Sucursal:             h.Sucursal(),
		GrupoCuentaCorriente: h.GrupoCtaCte(),
		CalleYNumero:         h.CalleYNro(),
		Piso:                 h.Piso(),
		Departamento:         h.Departamento(),
		CodigoPostal:         h.CodigoPostal(),
		Retencion:            h.Retencion(),
		VtoActual:            h.VtoActual(),
		GrupoAfinidad:        h.GrupoAfinidad(),
		CodProductoMarca:     h.CodProdMarca(),
		ModeloLiquidacion:    h.ModeloLiquidacion(),
		NroCuenta:            h.NroCuenta(),
		NroAutorizado:        h.NroAutorizado(),
		EResumen:             h.EResumen(),
		CantidadResumen:      h.CantResumen(),
	}
}

func resumenFromDomain(h *domain.Resumen) *dto.Resumen {
	return &dto.Resumen{
		Clave:                 claveFromDomain(h.Clave()),
		DatosFiliatorios:      datosFiliatoriosFromDomain(h.DatosFiliatorios()),
		DatosLiquidacionSocio: datosLiquidacionSocioFromDomain(h.DatosLiquidacionSocio()),
		DatosDeAutorizado: sf.Map[*domain.DatosDeAutorizado, *dto.DatosDeAutorizado](h.DatosDeAutorizado(),
			func(item *domain.DatosDeAutorizado) *dto.DatosDeAutorizado {
				return datosDeAutorizadoFromDomain(item)
			}),
		DatosMesAnterior: datosMesAnteriorFromDomain(h.DatosMesAnterior()),
		RegistroTotales: sf.Map[*domain.RegistroTotales, *dto.RegistroTotales](h.RegistroTotales(),
			func(item *domain.RegistroTotales) *dto.RegistroTotales {
				return registroTotalesFromDomain(item)
			}),
		DatosMesActual: datosMesActualFromDomain(h.DatosMesActual()),
		DetalleCargosYAjustes: sf.Map[*domain.DetalleCargosYAjustes, *dto.DetalleCargosYAjustes](h.DetalleCargosYAjustes(),
			func(item *domain.DetalleCargosYAjustes) *dto.DetalleCargosYAjustes {
				return detalleCargosYAjustesFromDomain(item)
			}),
		RegistroCargos: sf.Map[*domain.RegistroCargos, *dto.RegistroCargos](h.RegistroCargos(),
			func(item *domain.RegistroCargos) *dto.RegistroCargos { return registroCargosFromDomain(item) }),
		Leyendas: sf.Map[*domain.Leyenda, *dto.Leyenda](h.Leyendas(),
			func(item *domain.Leyenda) *dto.Leyenda { return leyendaFromDomain(item) }),
		Limites: sf.Map[*domain.Limite, *dto.Limite](h.Limites(),
			func(item *domain.Limite) *dto.Limite { return limiteFromDomain(item) }),
		Tasas: sf.Map[*domain.Tasa, *dto.Tasa](h.Tasas(),
			func(item *domain.Tasa) *dto.Tasa { return tasaFromDomain(item) }),
		DeudasNoVencidas: sf.Map[*domain.Deuda, *dto.Deuda](h.DeudasNoVencidas(),
			func(item *domain.Deuda) *dto.Deuda { return deudaFromDomain(item) }),
	}
}

func datosDeAutorizadoFromDomain(h *domain.DatosDeAutorizado) *dto.DatosDeAutorizado {
	if h == nil {
		return nil
	}
	return &dto.DatosDeAutorizado{
		Texto:          h.Texto(),
		Nombre:         h.Nombre(),
		Saldo:          h.Saldo(),
		SaldoEnDolares: h.SaldoEnDolares(),
		Movimientos: sf.Map[*domain.DetalleCargosYAjustes, *dto.DetalleCargosYAjustes](h.Movimientos(), func(item *domain.DetalleCargosYAjustes) *dto.DetalleCargosYAjustes {
			return detalleCargosYAjustesFromDomain(item)
		}),
		Numero: h.Numero(),
	}
}

func datosFiliatoriosFromDomain(h *domain.DatosFiliatorios) *dto.DatosFiliatorios {
	if h == nil {
		return nil
	}
	return &dto.DatosFiliatorios{
		NombreSocio:      h.NombreSocio(),
		NroCtaEmpresa:    h.NroCtaEmpresa(),
		NombreCtaEmpresa: h.NombreCtaEmpresa(),
		DomicilioCodPcia: h.DomicilioCodPcia(),
		DomicilioAdic1:   h.DomicilioAdic1(),
		DomicilioAdic2:   h.DomicilioAdic2(),
		Localidad:        h.Localidad(),
		PaisOtorg:        h.PaisOtorg(),
		TipoTarjeta:      h.TipoTarjeta(),
		DenomTipoTarj:    h.DenomTipoTarj(),
	}

}

func datosLiquidacionSocioFromDomain(h *domain.DatosLiquidacionSocio) *dto.DatosLiquidacionSocio {
	if h == nil {
		return nil
	}
	return &dto.DatosLiquidacionSocio{
		Cuit:                    h.Cuit(),
		CategoriaIva:            h.CategoriaIva(),
		DescripcionCategoriaIva: h.DescripcionCategoriaIva(),
		SbmTasas:                h.SbmTasas(),
		SbmGtosSV:               h.SbmGtosSV(),
		SbmParam:                h.SbmParam(),
		DvSocio:                 h.DvSocio(),
		MarcaCBU:                h.MarcaCBU(),
		Cbu:                     h.Cbu(),
		DvUnifSuc:               h.DvUnifSuc(),
		LeyDebit:                h.LeyDebit(),
		TipCta:                  h.TipCta(),
		NroCtaDebito:            h.NroCtaDebito(),
		TipPago:                 h.TipPago(),
		TipoCtaCte:              h.TipoCtaCte(),
		TarjetaTitular:          h.TarjetaTitular(),
		RegPagYAjSant:           h.RegPagYAjSant(),
		RegDetalleYMsg:          h.RegDetalleYMsg(),
		CantidadHojas:           h.CantidadHojas(),
		TipoProducto:            h.TipoProducto(),
		Comprobante:             h.Comprobante(),
	}

}

func datosMesActualFromDomain(h *domain.DatosMesActual) *dto.DatosMesActual {
	if h == nil {
		return nil
	}
	return &dto.DatosMesActual{
		FechaCierre:        h.FechaCierre(),
		Saldo:              h.Saldo(),
		SaldoEnDolares:     h.SaldoEnDolares(),
		PagoMinimo:         h.PagoMinimo(),
		ProximoVencimiento: h.ProximoVencimiento(),
		ProximaFechaCierre: h.ProximaFechaCierre(),
	}
}

func datosMesAnteriorFromDomain(h *domain.DatosMesAnterior) *dto.DatosMesAnterior {
	if h == nil {
		return nil
	}
	return &dto.DatosMesAnterior{
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

func destinatarioFromDomain(destinatario *domain.Destinatario) *dto.Destinatario {
	return &dto.Destinatario{
		Nombre:          destinatario.Nombre(),
		Direccion:       destinatario.Direccion(),
		Localidad:       destinatario.Localidad(),
		CodigoPostal:    destinatario.CodigoPostal(),
		Email:           destinatario.Email(),
		NumeroCuenta:    destinatario.NumeroCuenta(),
		NumeroDocumento: destinatario.NumeroDocumento(),
	}
}

func detalleCargosYAjustesFromDomain(h *domain.DetalleCargosYAjustes) *dto.DetalleCargosYAjustes {
	if h == nil {
		return nil
	}
	return &dto.DetalleCargosYAjustes{
		CodigoMov:       h.CodigoMov(),
		SubcodMov:       h.SubcodMov(),
		CodFinanc:       h.CodFinanc(),
		FecOper:         h.FecOper(),
		Cupon:           h.Cupon(),
		NroComercio:     h.NroComercio(),
		Detalle:         h.Detalle(),
		Importe:         h.Importe(),
		ImporteDol:      h.ImporteDol(),
		NroAutoriz:      h.NroAutoriz(),
		Concepto:        h.Concepto(),
		FecPresen:       h.FecPresen(),
		MarcaCF:         h.MarcaCF(),
		CodMoneda:       h.CodMoneda(),
		CapitalAdeudado: h.CapitalAdeudado(),
		InteresCuota:    h.InteresCuota(),
		TnaPr:           h.TnaPr(),
		Vigencia:        h.Vigencia(),
	}
}

func deudaFromDomain(h *domain.Deuda) *dto.Deuda {
	if h == nil {
		return nil
	}
	return &dto.Deuda{
		DenMes:    h.DenMes(),
		DenMoneda: h.DenMoneda(),
		DnvMes:    h.DnvMes(),
	}
}

func leyendaFromDomain(h *domain.Leyenda) *dto.Leyenda {
	if h == nil {
		return nil
	}
	return &dto.Leyenda{
		RubroLeyendaInf: h.RubroLeyendaInf(),
		LeyendaInf:      h.LeyendaInf(),
		McaFacturacion:  h.McaFacturacion(),
	}
}

func registroCargosFromDomain(h *domain.RegistroCargos) *dto.RegistroCargos {
	if h == nil {
		return nil
	}
	return &dto.RegistroCargos{
		Descripcion:    h.Descripcion(),
		Importe:        h.Importe(),
		ImporteDolares: h.ImporteDolares(),
	}
}

func registroTotalesFromDomain(h *domain.RegistroTotales) *dto.RegistroTotales {
	if h == nil {
		return nil
	}
	return &dto.RegistroTotales{
		FechaPagoAju:          h.FechaPagoAju(),
		Descripcion:           h.Descripcion(),
		ImporteTotal:          h.ImporteTotal(),
		ImporteTotalEnDolares: h.ImporteTotalEnDolares(),
		MontoImputOtor:        h.MontoImputOtor(),
		MontoImputDol:         h.MontoImputDol(),
	}
}

func tasaFromDomain(h *domain.Tasa) *dto.Tasa {
	if h == nil {
		return nil
	}
	return &dto.Tasa{
		Descripcion:   h.Descripcion(),
		Tasa:          h.Tasa(),
		TasaEnDolares: h.TasaEnDolares(),
	}
}

func limiteFromDomain(h *domain.Limite) *dto.Limite {
	if h == nil {
		return nil
	}
	return &dto.Limite{
		Descripcion: h.Descripcion(),
		Limite:      h.Limite(),
	}
}
