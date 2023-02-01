package serialization

import (
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	mastercard "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/controller/mastercard/dto"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"
	sf "github.com/sa-/slicefunk"
)

func SerializeUsuario(h *domain.Usuario) *mastercard.Usuario {

	return &mastercard.Usuario{
		Nombre:          h.Nombre(),
		Direccion:       h.Direccion(),
		Localidad:       h.Localidad(),
		CodigoPostal:    h.CodigoPostal(),
		Email:           h.Email(),
		NumeroCuenta:    h.NumeroCuenta(),
		NumeroDocumento: h.NumeroDocumento(),
	}

}

func SerializeIndice(h *domain.Indice) *mastercard.IndiceResumen {
	return &mastercard.IndiceResumen{
		FechaCierre:   h.FechaCierre(),
		Marca:         h.Marca(),
		Path:          h.Path(),
		Month:         h.Month(),
		Year:          h.Year(),
		ModeloResumen: h.ModeloResumen(),
	}
}

func SerializeResumen(h *domain.Resumen) *dto.Resumen {
	return &dto.Resumen{
		Clave:                 claveToDto(h.Clave()),
		DatosFiliatorios:      datosFiliatoriosToDto(h.DatosFiliatorios()),
		DatosLiquidacionSocio: datosLiquidacionSocioToDto(h.DatosLiquidacionSocio()),
		DatosDeAutorizado: sf.Map[*contenido_resumen.DatosDeAutorizado, *dto.DatosDeAutorizado](h.DatosDeAutorizado(),
			func(item *contenido_resumen.DatosDeAutorizado) *dto.DatosDeAutorizado {
				return datosDeAutorizadoToDto(item)
			}),
		DatosMesAnterior: datosMesAnteriorToDto(h.DatosMesAnterior()),
		RegistroTotales: sf.Map[*contenido_resumen.RegistroTotales, *dto.RegistroTotales](h.RegistroTotales(),
			func(item *contenido_resumen.RegistroTotales) *dto.RegistroTotales {
				return registroTotalesToDto(item)
			}),
		DatosMesActual: datosMesActualToDto(h.DatosMesActual()),
		DetalleCargosYAjustes: sf.Map[*contenido_resumen.DetalleCargosYAjustes, *dto.DetalleCargosYAjustes](h.DetalleCargosYAjustes(),
			func(item *contenido_resumen.DetalleCargosYAjustes) *dto.DetalleCargosYAjustes {
				return detalleCargosYAjustestoDto(item)
			}),
		RegistroCargos: sf.Map[*contenido_resumen.RegistroCargos, *dto.RegistroCargos](h.RegistroCargos(),
			func(item *contenido_resumen.RegistroCargos) *dto.RegistroCargos {
				return registroCargosToDto(item)
			}),
		Leyendas: sf.Map[*contenido_resumen.Leyenda, *dto.Leyenda](h.Leyendas(),
			func(item *contenido_resumen.Leyenda) *dto.Leyenda { return leyendaToDto(item) }),
		Limites: sf.Map[*contenido_resumen.Limite, *dto.Limite](h.Limites(),
			func(item *contenido_resumen.Limite) *dto.Limite { return limiteToDto(item) }),
		Tasas: sf.Map[*contenido_resumen.Tasa, *dto.Tasa](h.Tasas(),
			func(item *contenido_resumen.Tasa) *dto.Tasa { return tasaToDto(item) }),
		DeudasNoVencidas: sf.Map[*contenido_resumen.Deuda, *dto.Deuda](h.DeudasNoVencidas(),
			func(item *contenido_resumen.Deuda) *dto.Deuda { return deudaToDto(item) }),
	}
}

func claveToDto(h *contenido_resumen.Clave) *dto.Clave {
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

func datosDeAutorizadoToDto(h *contenido_resumen.DatosDeAutorizado) *dto.DatosDeAutorizado {
	movimientos := sf.Map[*contenido_resumen.DetalleCargosYAjustes, *dto.DetalleCargosYAjustes](
		h.Movimientos(), func(item *contenido_resumen.DetalleCargosYAjustes) *dto.DetalleCargosYAjustes {
			return detalleCargosYAjustestoDto(item)
		})

	return &dto.DatosDeAutorizado{

		Texto:          h.Texto(),
		Nombre:         h.Nombre(),
		Saldo:          h.Saldo(),
		SaldoEnDolares: h.SaldoEnDolares(),
		Movimientos:    movimientos,
		Numero:         h.Numero(),
	}

}

func datosFiliatoriosToDto(h *contenido_resumen.DatosFiliatorios) *dto.DatosFiliatorios {
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

func datosLiquidacionSocioToDto(h *contenido_resumen.DatosLiquidacionSocio) *dto.DatosLiquidacionSocio {
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

func datosMesActualToDto(h *contenido_resumen.DatosMesActual) *dto.DatosMesActual {
	return &dto.DatosMesActual{
		FechaCierre:        h.FechaCierre(),
		Saldo:              h.Saldo(),
		SaldoEnDolares:     h.SaldoEnDolares(),
		PagoMinimo:         h.PagoMinimo(),
		ProximoVencimiento: h.ProximoVencimiento(),
		ProximaFechaCierre: h.ProximaFechaCierre(),
	}
}

func datosMesAnteriorToDto(h *contenido_resumen.DatosMesAnterior) *dto.DatosMesAnterior {
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

func detalleCargosYAjustestoDto(h *contenido_resumen.DetalleCargosYAjustes) *dto.DetalleCargosYAjustes {

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

func deudaToDto(h *contenido_resumen.Deuda) *dto.Deuda {
	return &dto.Deuda{
		DenMes:    h.DenMes(),
		DenMoneda: h.DenMoneda(),
		DnvMes:    h.DnvMes(),
	}
}

func leyendaToDto(h *contenido_resumen.Leyenda) *dto.Leyenda {
	return &dto.Leyenda{
		RubroLeyendaInf: h.RubroLeyendaInf(),
		LeyendaInf:      h.LeyendaInf(),
		McaFacturacion:  h.McaFacturacion(),
	}
}

func limiteToDto(h *contenido_resumen.Limite) *dto.Limite {
	return &dto.Limite{
		Descripcion: h.Descripcion(),
		Limite:      h.Limite(),
	}
}

func registroCargosToDto(h *contenido_resumen.RegistroCargos) *dto.RegistroCargos {
	return &dto.RegistroCargos{
		Descripcion:    h.Descripcion(),
		Importe:        h.Importe(),
		ImporteDolares: h.ImporteDolares(),
	}
}

func registroTotalesToDto(h *contenido_resumen.RegistroTotales) *dto.RegistroTotales {
	return &dto.RegistroTotales{
		FechaPagoAju:          h.FechaPagoAju(),
		Descripcion:           h.Descripcion(),
		ImporteTotal:          h.ImporteTotal(),
		ImporteTotalEnDolares: h.ImporteTotalEnDolares(),
		MontoImputOtor:        h.MontoImputOtor(),
		MontoImputDol:         h.MontoImputDol(),
	}
}

func tasaToDto(item *contenido_resumen.Tasa) *dto.Tasa {
	return &dto.Tasa{
		Descripcion:   item.Descripcion(),
		Tasa:          item.Tasa(),
		TasaEnDolares: item.TasaEnDolares(),
	}
}
