package line_parser

import (
	"github.com/redbeestudios/go-parser-poc/commons/utils"
	. "github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain/tipo_registro"
	"strings"
)

func DatosAutorizadoFromString(line string) *DatosDeAutorizado {
	reader := utils.NewLineReader(line)
	datosDeAutorizado := NewDatosDeAutorizado()
	datosDeAutorizado.SetTexto(reader.Read(15, 1))
	datosDeAutorizado.SetNombre(reader.Read(26, 1))
	datosDeAutorizado.SetSldo(reader.Read(16, 1))
	datosDeAutorizado.SetSldoDol(reader.Read(13))
	datosDeAutorizado.SetNumero("0")
	return datosDeAutorizado
}

func DatosFiliatoriosFromString(line string) *DatosFiliatorios {
	reader := utils.NewLineReader(line)
	d := NewDatosFiliatorios()
	d.SetNombreSocio(reader.Read(30, 1))
	d.SetNroCtaEmpresa(reader.Read(7, 1))
	d.SetNombreCtaEmpresa(reader.Read(20, 1))
	d.SetDomicilioCodPcia(reader.Read(1, 1))
	d.SetDomicilioAdic1(reader.Read(17, 1))
	d.SetDomicilioAdic2(reader.Read(43, 1))
	d.SetLocalidad(reader.Read(20, 1))
	d.SetPaisOtorg(reader.Read(1, 1))
	d.SetTipoTarjeta(reader.Read(2, 1))
	d.SetDenomTipoTarj(reader.Read(24))

	return d
}

func DatosLiquidacionSocioFromString(line string) *DatosLiquidacionSocio {
	r := utils.NewLineReader(line)
	d := NewDatosLiquidacionSocio()
	d.SetCuit(r.Read(13, 1))
	d.SetCategoriaIva(r.Read(1, 1))
	d.SetDescripcionCategoriaIva(r.Read(25, 1))
	d.SetSbmTasas(r.Read(1, 1))
	d.SetSbmGtosSV(r.Read(1, 1))
	d.SetSbmParam(r.Read(1, 1))
	d.SetDvSocio(r.Read(1, 1))
	d.SetMarcaCBU(r.Read(1, 1))
	d.SetCbu(r.Read(22, 1))
	d.SetDvUnifSuc(r.Read(1, 1))
	d.SetLeyDebit(r.Read(1, 1))
	d.SetTipCta(r.Read(1, 1))
	d.SetNroCtaDebito(r.Read(10, 1))
	d.SetTipPago(r.Read(1, 1))
	d.SetTipoCtaCte(r.Read(1, 1))
	d.SetTarjetaTitular(r.Read(16, 1))
	d.SetRegPagYAjSant(r.Read(5, 1))
	d.SetRegDetalleYMsg(r.Read(5, 1))
	d.SetCantidadHojas(r.Read(3, 11))
	d.SetTipoProducto(r.Read(3, 1))
	d.SetComprobante(r.Read(12))

	return d
}

func DatosMesActualFromString(line string) *DatosMesActual {
	r := utils.NewLineReader(line)
	d := NewDatosMesActual()

	d.SetFechaCierre(r.Read(9, 1))
	d.SetSaldo(r.Read(16, 1))
	d.SetSaldoEnDolares(r.Read(13, 1))
	d.SetPagoMinimo(r.Read(16, 1))
	d.SetProximoVencimiento(r.Read(9, 1))
	d.SetProximaFechaCierre(r.Read(9))

	return d
}

func DatosMesAnteriorFromString(line string) *DatosMesAnterior {
	r := utils.NewLineReader(line)

	d := NewDatosMesAnterior()

	d.SetFechaCierre(r.Read(9, 1))
	d.SetVencimiento(r.Read(9, 1))
	d.SetSaldoEnDolares(r.Read(12, 1))
	d.SetSaldo(r.Read(16, 1))
	d.SetSaldoTotal(r.Read(16, 1))
	d.SetPagoMinimo(r.Read(16, 1))
	d.SetCompras(r.Read(16, 1))
	d.SetComprasEnDolares(r.Read(10))

	return d
}

func DetalleCargosYAjustesFromString(line string) *DetalleCargosYAjustes {
	r := utils.NewLineReader(line)
	d := NewDetalleCargosYAjustes()

	d.SetCodigoMov(r.Read(1, 1))
	d.SetSubcodMov(r.Read(1, 1))
	d.SetCodFinanc(r.Read(1, 1))
	d.SetFecOper(r.Read(9, 1))
	d.SetCupon(r.Read(5, 1))
	d.SetNroComercio(r.Read(9, 1))
	d.SetDetalle(r.Read(39, 1))
	d.SetImporte(r.Read(16, 1))
	d.SetImporteDol(r.Read(12, 1))
	d.SetNroAutoriz(r.Read(1, 1))
	d.SetConcepto(r.Read(3, 1))
	d.SetFecPresen(r.Read(7, 1))
	d.SetMarcaCF(r.Read(1, 1))
	d.SetCodMoneda(r.Read(3, 1))
	d.SetCapitalAdeudado(r.Read(16, 1))
	d.SetInteresCuota(r.Read(16, 1))
	d.SetTnaPr(r.Read(6, 1))
	d.SetVigencia(r.Read(1))

	return d
}

func DeudaFromString(line string) *Deuda {
	reader := utils.NewLineReader(line)
	d := NewDeuda()

	d.SetDenMes(reader.Read(13, 1))
	d.SetDenMoneda(reader.Read(3, 1))
	d.SetDnvMes(reader.Read(10))

	return d
}

func LeyendaFromString(line string) *Leyenda {
	lineReader := utils.NewLineReader(line)
	l := NewLeyenda()
	l.SetRubroLeyendaInf(lineReader.Read(1, 1))
	l.SetLeyendaInf(lineReader.Read(100, 1))
	l.SetMcaFacturacion(lineReader.Read(1))
	return l
}

func LimiteFromString(line string) *Limite {
	r := utils.NewLineReader(line)
	l := NewLimite()

	l.SetDescripcion(r.Read(24, 1))
	l.SetLimite(r.Read(15))

	return l
}

func RegistroCargosFromString(line string) *RegistroCargos {
	r := utils.NewLineReader(line)
	c := NewRegistroCargos()

	c.SetDescripcion(r.Read(33, 1))
	c.SetImporte(r.Read(16, 1))
	c.SetImporteDolares(r.Read(10))

	return c
}

func RegistroTotalesFromString(line string) *RegistroTotales {
	r := utils.NewLineReader(line)
	t := NewRegistroTotales()

	t.SetFechaPagoAju(r.Read(9, 1))
	t.SetDescripcion(r.Read(33, 1))
	t.SetImporteTotal(r.Read(16, 1))
	t.SetImporteTotalEnDolares(r.Read(10, 1))
	t.SetMontoImputOtor(r.Read(16, 1))
	t.SetMontoImputDol(r.Read(10))

	return t
}

func TasaFromString(line string) *Tasa {
	reader := utils.NewLineReader(line)
	tasa := NewTasa()
	tasa.SetDescripcion(reader.Read(3, 1))
	tasa.SetTasa(reader.Read(9, 1))
	tasa.SetTasaEnDolares(reader.Read(7))
	return tasa
}

func AddFieldFromLine(nroRegistro string, lineToParse string, resumenParcial *Resumen) *Resumen {
	switch nroRegistro {
	case tipo_registro.DatosFiliatorios.String():
		datosFiliatorios := DatosFiliatoriosFromString(lineToParse)
		resumenParcial.SetDatosFiliatorios(datosFiliatorios)
		break
	case tipo_registro.DatosLiquidacionSocio.String():
		datosLiquidacionSocio := DatosLiquidacionSocioFromString(lineToParse)
		resumenParcial.SetDatosLiquidacionSocio(datosLiquidacionSocio)
		break
	case tipo_registro.DatosDeAutorizado.String():
		datosDeAutorizado := DatosAutorizadoFromString(lineToParse)
		resumenParcial.AddDatosDeAutorizado(datosDeAutorizado)
		break
	case tipo_registro.DatosMesAnterior.String():
		datosMesAnterior := DatosMesAnteriorFromString(lineToParse)
		resumenParcial.SetDatosMesAnterior(datosMesAnterior)
		break
	case tipo_registro.RegistroTotales.String():
		registroTotales := RegistroTotalesFromString(lineToParse)
		resumenParcial.AddRegistroTotales(registroTotales)
		break
	case tipo_registro.DatosMesActual.String():
		datosMesActual := DatosMesActualFromString(lineToParse)
		resumenParcial.SetDatosMesActual(datosMesActual)
		break
	case tipo_registro.DetalleCargosComprasYAjustes.String():
		detalleCargosYAjustes := DetalleCargosYAjustesFromString(lineToParse)
		resumenParcial.AddDetalleCargosYAjustes(detalleCargosYAjustes)
		break
	case tipo_registro.RegistroCargos.String():
		registroCargos := RegistroCargosFromString(lineToParse)
		resumenParcial.AddRegistroCargos(registroCargos)
		break
	case tipo_registro.Leyendas.String():
		leyenda := LeyendaFromString(lineToParse)
		resumenParcial.AddLeyenda(leyenda)
		break
	case tipo_registro.Limites.String():
		limite := LimiteFromString(lineToParse)
		resumenParcial.AddLimite(limite)
		break
	case tipo_registro.Tasas.String():
		tasa := TasaFromString(lineToParse)
		resumenParcial.AddTasa(tasa)
		break
	case tipo_registro.DeudasNoVencidasSemestre.String():
		deuda := DeudaFromString(lineToParse)
		resumenParcial.AddDeudaNoVencida(deuda)
		break
	}
	return resumenParcial

}

func ResumenAEnviarFromMailing(line string) *ResumenAEnviar {
	lineArray := strings.Split(line, "|")
	modeloResumen := lineArray[0]
	nombre := lineArray[1]
	direccion := lineArray[2]
	localidad := lineArray[3]
	codigoPostal := lineArray[4]
	email := lineArray[5]
	codigoVinculacion := lineArray[6]
	marca := lineArray[7]
	periodo := lineArray[8]
	nroDocumento := lineArray[9]

	codigoVinculacionDividido := strings.Split(codigoVinculacion, "-")

	numeroCuenta := codigoVinculacionDividido[0]
	fechaEnvio := codigoVinculacionDividido[1]

	destinatario := NewDestinatario(nombre, direccion, localidad, codigoPostal, email, numeroCuenta, nroDocumento)

	resumenAEnviar := NewResumenAEnviar(
		modeloResumen, destinatario, fechaEnvio, marca, periodo, codigoVinculacion)

	return resumenAEnviar

}
