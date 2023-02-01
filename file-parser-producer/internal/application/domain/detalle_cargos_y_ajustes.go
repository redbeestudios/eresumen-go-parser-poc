package domain

type DetalleCargosYAjustes struct {
	codigoMov       string
	subcodMov       string
	codFinanc       string
	fecOper         string
	cupon           string
	nroComercio     string
	detalle         string
	importe         string
	importeDol      string
	nroAutoriz      string
	concepto        string
	fecPresen       string
	marcaCF         string
	codMoneda       string
	capitalAdeudado string
	interesCuota    string
	tnaPr           string
	vigencia        string
}

func (d *DetalleCargosYAjustes) SetCodigoMov(codigoMov string) {
	d.codigoMov = codigoMov
}

func (d *DetalleCargosYAjustes) SetSubcodMov(subcodMov string) {
	d.subcodMov = subcodMov
}

func (d *DetalleCargosYAjustes) SetCodFinanc(codFinanc string) {
	d.codFinanc = codFinanc
}

func (d *DetalleCargosYAjustes) SetFecOper(fecOper string) {
	d.fecOper = fecOper
}

func (d *DetalleCargosYAjustes) SetCupon(cupon string) {
	d.cupon = cupon
}

func (d *DetalleCargosYAjustes) SetNroComercio(nroComercio string) {
	d.nroComercio = nroComercio
}

func (d *DetalleCargosYAjustes) SetDetalle(detalle string) {
	d.detalle = detalle
}

func (d *DetalleCargosYAjustes) SetImporte(importe string) {
	d.importe = importe
}

func (d *DetalleCargosYAjustes) SetImporteDol(importeDol string) {
	d.importeDol = importeDol
}

func (d *DetalleCargosYAjustes) SetNroAutoriz(nroAutoriz string) {
	d.nroAutoriz = nroAutoriz
}

func (d *DetalleCargosYAjustes) SetConcepto(concepto string) {
	d.concepto = concepto
}

func (d *DetalleCargosYAjustes) SetFecPresen(fecPresen string) {
	d.fecPresen = fecPresen
}

func (d *DetalleCargosYAjustes) SetMarcaCF(marcaCF string) {
	d.marcaCF = marcaCF
}

func (d *DetalleCargosYAjustes) SetCodMoneda(codMoneda string) {
	d.codMoneda = codMoneda
}

func (d *DetalleCargosYAjustes) SetCapitalAdeudado(capitalAdeudado string) {
	d.capitalAdeudado = capitalAdeudado
}

func (d *DetalleCargosYAjustes) SetInteresCuota(interesCuota string) {
	d.interesCuota = interesCuota
}

func (d *DetalleCargosYAjustes) SetTnaPr(tnaPr string) {
	d.tnaPr = tnaPr
}

func (d *DetalleCargosYAjustes) SetVigencia(vigencia string) {
	d.vigencia = vigencia
}

func (d DetalleCargosYAjustes) CodigoMov() string {
	return d.codigoMov
}

func (d DetalleCargosYAjustes) SubcodMov() string {
	return d.subcodMov
}

func (d DetalleCargosYAjustes) CodFinanc() string {
	return d.codFinanc
}

func (d DetalleCargosYAjustes) FecOper() string {
	return d.fecOper
}

func (d DetalleCargosYAjustes) Cupon() string {
	return d.cupon
}

func (d DetalleCargosYAjustes) NroComercio() string {
	return d.nroComercio
}

func (d DetalleCargosYAjustes) Detalle() string {
	return d.detalle
}

func (d DetalleCargosYAjustes) Importe() string {
	return d.importe
}

func (d DetalleCargosYAjustes) ImporteDol() string {
	return d.importeDol
}

func (d DetalleCargosYAjustes) NroAutoriz() string {
	return d.nroAutoriz
}

func (d DetalleCargosYAjustes) Concepto() string {
	return d.concepto
}

func (d DetalleCargosYAjustes) FecPresen() string {
	return d.fecPresen
}

func (d DetalleCargosYAjustes) MarcaCF() string {
	return d.marcaCF
}

func (d DetalleCargosYAjustes) CodMoneda() string {
	return d.codMoneda
}

func (d DetalleCargosYAjustes) CapitalAdeudado() string {
	return d.capitalAdeudado
}

func (d DetalleCargosYAjustes) InteresCuota() string {
	return d.interesCuota
}

func (d DetalleCargosYAjustes) TnaPr() string {
	return d.tnaPr
}

func (d DetalleCargosYAjustes) Vigencia() string {
	return d.vigencia
}

func NewDetalleCargosYAjustes() *DetalleCargosYAjustes {
	return &DetalleCargosYAjustes{}
}
