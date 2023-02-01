package domain

import (
	"github.com/redbeestudios/go-parser-poc/commons/utils"
	"strings"
)

type Clave struct {
	nroRegistro          string
	nroSecuencia         string
	nroOrden             string
	entidad              string
	sucursal             string
	grupoCuentaCorriente string
	calleYNumero         string
	piso                 string
	departamento         string
	codigoPostal         string
	retencion            string
	vtoActual            string
	grupoAfinidad        string
	codProductoMarca     string
	modeloLiquidacion    string
	nroCuenta            string
	nroAutorizado        string
	eResumen             string
	cantidadResumen      string
}

func (h *Clave) SetNroRegistro(nroRegistro string) {
	h.nroRegistro = nroRegistro
}

func (h *Clave) SetNroSecuencia(nroSecuencia string) {
	h.nroSecuencia = nroSecuencia
}

func (h *Clave) SetNroOrden(nroOrden string) {
	h.nroOrden = nroOrden
}

func (h *Clave) SetEntidad(entidad string) {
	h.entidad = entidad
}

func (h *Clave) SetSucursal(sucursal string) {
	h.sucursal = sucursal
}

func (h *Clave) SetGrupoCuentaCorriente(grupoCuentaCorriente string) {
	h.grupoCuentaCorriente = grupoCuentaCorriente
}

func (h *Clave) SetCalleYNumero(calleYNumero string) {
	h.calleYNumero = calleYNumero
}

func (h *Clave) SetPiso(piso string) {
	h.piso = piso
}

func (h *Clave) SetDepartamento(departamento string) {
	h.departamento = departamento
}

func (h *Clave) SetCodigoPostal(codigoPostal string) {
	h.codigoPostal = codigoPostal
}

func (h *Clave) SetRetencion(retencion string) {
	h.retencion = retencion
}

func (h *Clave) SetVtoActual(vtoActual string) {
	h.vtoActual = vtoActual
}

func (h *Clave) SetGrupoAfinidad(grupoAfinidad string) {
	h.grupoAfinidad = grupoAfinidad
}

func (h *Clave) SetCodProductoMarca(codProductoMarca string) {
	h.codProductoMarca = codProductoMarca
}

func (h *Clave) SetModeloLiquidacion(modeloLiquidacion string) {
	h.modeloLiquidacion = modeloLiquidacion
}

func (h *Clave) SetNroCuenta(nroCuenta string) {
	h.nroCuenta = nroCuenta
}

func (h *Clave) SetNroAutorizado(nroAutorizado string) {
	h.nroAutorizado = nroAutorizado
}

func (h *Clave) SetEResumen(eResumen string) {
	h.eResumen = eResumen
}

func (h *Clave) SetCantidadResumen(cantidadResumen string) {
	h.cantidadResumen = cantidadResumen
}

func NewHeader() *Clave {
	return &Clave{}
}

func (h *Clave) NroRegistro() string {
	return h.nroRegistro
}

func (h *Clave) NroSecuencia() string {
	return h.nroSecuencia
}

func (h *Clave) NroOrden() string {
	return h.nroOrden
}

func (h *Clave) Entidad() string {
	return h.entidad
}

func (h *Clave) Sucursal() string {
	return h.sucursal
}

func (h *Clave) GrupoCtaCte() string {
	return h.grupoCuentaCorriente
}

func (h *Clave) CalleYNro() string {
	return h.calleYNumero
}

func (h *Clave) Piso() string {
	return h.piso
}

func (h *Clave) Departamento() string {
	return h.departamento
}

func (h *Clave) CodigoPostal() string {
	return h.codigoPostal
}

func (h *Clave) Retencion() string {
	return h.retencion
}

func (h *Clave) VtoActual() string {
	return h.vtoActual
}

func (h *Clave) CodProdMarca() string {
	return h.codProductoMarca
}

func (h *Clave) GrupoAfinidad() string {
	return h.grupoAfinidad
}

func (h *Clave) ModeloLiquidacion() string {
	return h.modeloLiquidacion
}

func (h *Clave) NroCuenta() string {
	return h.nroCuenta
}

func (h *Clave) NroAutorizado() string {
	return h.nroAutorizado
}

func (h *Clave) EResumen() string {
	return h.eResumen
}

func (h *Clave) CantResumen() string {
	return h.cantidadResumen
}

func ClaveFromString(line string) *Clave {

	r := utils.NewLineReader(line)
	h := NewHeader()
	h.SetNroRegistro(r.Read(5))
	h.SetNroSecuencia(r.Read(5, 1))
	h.SetNroOrden(r.Read(6, 1))
	h.SetEntidad(r.Read(3, 1))
	h.SetSucursal(r.Read(3, 1))
	h.SetGrupoCuentaCorriente(r.Read(2, 1))
	h.SetCalleYNumero(r.Read(30, 1))
	h.SetPiso(r.Read(2, 1))
	h.SetDepartamento(r.Read(3, 1))
	h.SetCodigoPostal(r.Read(8, 1))
	h.SetRetencion(r.Read(1, 1))
	h.SetVtoActual(r.Read(9, 1))
	h.SetGrupoAfinidad(r.Read(5, 1))
	h.SetCodProductoMarca(r.Read(5, 1))
	h.SetModeloLiquidacion(r.Read(5, 1))
	h.SetNroCuenta(strings.TrimLeft(r.Read(9, 1), "0"))
	h.SetNroAutorizado(r.Read(1, 1))
	h.SetEResumen(r.Read(1, 1))
	h.SetCantidadResumen(r.Read(5, 26))

	return h
}
