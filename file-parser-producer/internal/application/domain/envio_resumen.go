package domain

import (
	"strings"
)

type ResumenAEnviar struct {
	resumen           *Resumen
	modeloResumen     string
	destinatario      *Destinatario
	fechaEnvio        string
	marca             string
	periodo           string
	codigoVinculacion string
}

func (r *ResumenAEnviar) CodigoVinculacion() string {
	return r.codigoVinculacion
}

func NewResumenAEnviar(modeloResumen string, destinatario *Destinatario, fechaEnvio string, marca string, periodo string, codigoVinculacion string) *ResumenAEnviar {
	return &ResumenAEnviar{modeloResumen: modeloResumen, destinatario: destinatario, fechaEnvio: fechaEnvio, marca: marca, periodo: periodo, codigoVinculacion: codigoVinculacion}
}

func (r *ResumenAEnviar) Resumen() *Resumen {
	return r.resumen
}

func (r *ResumenAEnviar) SetResumen(resumen *Resumen) {
	r.resumen = resumen
}

func (r *ResumenAEnviar) ModeloResumen() string {
	return r.modeloResumen
}

func (r *ResumenAEnviar) Destinatario() *Destinatario {
	return r.destinatario
}

func (r *ResumenAEnviar) FechaEnvio() string {
	return r.fechaEnvio
}

func (r *ResumenAEnviar) Marca() string {
	return r.marca
}

func (r *ResumenAEnviar) Periodo() string {
	return r.periodo
}

func FromLineaDestinatarios(line string) *ResumenAEnviar {
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
