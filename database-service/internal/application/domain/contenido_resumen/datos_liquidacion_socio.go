package contenido_resumen

type DatosLiquidacionSocio struct {
	cuit                    string
	categoriaIva            string
	descripcionCategoriaIva string
	sbmTasas                string
	sbmGtosSV               string
	sbmParam                string
	dvSocio                 string
	marcaCBU                string
	cbu                     string
	dvUnifSuc               string
	leyDebit                string
	tipCta                  string
	nroCtaDebito            string
	tipPago                 string
	tipoCtaCte              string
	tarjetaTitular          string
	regPagYAjSant           string
	regDetalleYMsg          string
	cantidadHojas           string
	tipoProducto            string
	comprobante             string
}

func (d *DatosLiquidacionSocio) SetCuit(cuit string) {
	d.cuit = cuit
}

func (d *DatosLiquidacionSocio) SetCategoriaIva(categoriaIva string) {
	d.categoriaIva = categoriaIva
}

func (d *DatosLiquidacionSocio) SetDescripcionCategoriaIva(descripcionCategoriaIva string) {
	d.descripcionCategoriaIva = descripcionCategoriaIva
}

func (d *DatosLiquidacionSocio) SetSbmTasas(sbmTasas string) {
	d.sbmTasas = sbmTasas
}

func (d *DatosLiquidacionSocio) SetSbmGtosSV(sbmGtosSV string) {
	d.sbmGtosSV = sbmGtosSV
}

func (d *DatosLiquidacionSocio) SetSbmParam(sbmParam string) {
	d.sbmParam = sbmParam
}

func (d *DatosLiquidacionSocio) SetDvSocio(dvSocio string) {
	d.dvSocio = dvSocio
}

func (d *DatosLiquidacionSocio) SetMarcaCBU(marcaCBU string) {
	d.marcaCBU = marcaCBU
}

func (d *DatosLiquidacionSocio) SetCbu(cbu string) {
	d.cbu = cbu
}

func (d *DatosLiquidacionSocio) SetDvUnifSuc(dvUnifSuc string) {
	d.dvUnifSuc = dvUnifSuc
}

func (d *DatosLiquidacionSocio) SetLeyDebit(leyDebit string) {
	d.leyDebit = leyDebit
}

func (d *DatosLiquidacionSocio) SetTipCta(tipCta string) {
	d.tipCta = tipCta
}

func (d *DatosLiquidacionSocio) SetNroCtaDebito(nroCtaDebito string) {
	d.nroCtaDebito = nroCtaDebito
}

func (d *DatosLiquidacionSocio) SetTipPago(tipPago string) {
	d.tipPago = tipPago
}

func (d *DatosLiquidacionSocio) SetTipoCtaCte(tipoCtaCte string) {
	d.tipoCtaCte = tipoCtaCte
}

func (d *DatosLiquidacionSocio) SetTarjetaTitular(tarjetaTitular string) {
	d.tarjetaTitular = tarjetaTitular
}

func (d *DatosLiquidacionSocio) SetRegPagYAjSant(regPagYAjSant string) {
	d.regPagYAjSant = regPagYAjSant
}

func (d *DatosLiquidacionSocio) SetRegDetalleYMsg(regDetalleYMsg string) {
	d.regDetalleYMsg = regDetalleYMsg
}

func (d *DatosLiquidacionSocio) SetCantidadHojas(cantidadHojas string) {
	d.cantidadHojas = cantidadHojas
}

func (d *DatosLiquidacionSocio) SetTipoProducto(tipoProducto string) {
	d.tipoProducto = tipoProducto
}

func (d *DatosLiquidacionSocio) SetComprobante(comprobante string) {
	d.comprobante = comprobante
}

func (d DatosLiquidacionSocio) Cuit() string {
	return d.cuit
}

func (d DatosLiquidacionSocio) CategoriaIva() string {
	return d.categoriaIva
}

func (d DatosLiquidacionSocio) DescripcionCategoriaIva() string {
	return d.descripcionCategoriaIva
}

func (d DatosLiquidacionSocio) SbmTasas() string {
	return d.sbmTasas
}

func (d DatosLiquidacionSocio) SbmGtosSV() string {
	return d.sbmGtosSV
}

func (d DatosLiquidacionSocio) SbmParam() string {
	return d.sbmParam
}

func (d DatosLiquidacionSocio) DvSocio() string {
	return d.dvSocio
}

func (d DatosLiquidacionSocio) MarcaCBU() string {
	return d.marcaCBU
}

func (d DatosLiquidacionSocio) Cbu() string {
	return d.cbu
}

func (d DatosLiquidacionSocio) DvUnifSuc() string {
	return d.dvUnifSuc
}

func (d DatosLiquidacionSocio) LeyDebit() string {
	return d.leyDebit
}

func (d DatosLiquidacionSocio) TipCta() string {
	return d.tipCta
}

func (d DatosLiquidacionSocio) NroCtaDebito() string {
	return d.nroCtaDebito
}

func (d DatosLiquidacionSocio) TipPago() string {
	return d.tipPago
}

func (d DatosLiquidacionSocio) TipoCtaCte() string {
	return d.tipoCtaCte
}

func (d DatosLiquidacionSocio) TarjetaTitular() string {
	return d.tarjetaTitular
}

func (d DatosLiquidacionSocio) RegPagYAjSant() string {
	return d.regPagYAjSant
}

func (d DatosLiquidacionSocio) RegDetalleYMsg() string {
	return d.regDetalleYMsg
}

func (d DatosLiquidacionSocio) CantidadHojas() string {
	return d.cantidadHojas
}

func (d DatosLiquidacionSocio) TipoProducto() string {
	return d.tipoProducto
}

func (d DatosLiquidacionSocio) Comprobante() string {
	return d.comprobante
}

func NewDatosLiquidacionSocio() *DatosLiquidacionSocio {
	return &DatosLiquidacionSocio{}
}
