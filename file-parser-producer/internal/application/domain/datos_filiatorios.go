package domain

type DatosFiliatorios struct {
	nombreSocio      string
	nroCtaEmpresa    string
	nombreCtaEmpresa string
	domicilioCodPcia string
	domicilioAdic1   string
	domicilioAdic2   string
	localidad        string
	paisOtorg        string
	tipoTarjeta      string
	denomTipoTarj    string
}

func (h *DatosFiliatorios) SetNombreSocio(nombreSocio string) {
	h.nombreSocio = nombreSocio
}

func (h *DatosFiliatorios) SetNroCtaEmpresa(nroCtaEmpresa string) {
	h.nroCtaEmpresa = nroCtaEmpresa
}

func (h *DatosFiliatorios) SetNombreCtaEmpresa(nombreCtaEmpresa string) {
	h.nombreCtaEmpresa = nombreCtaEmpresa
}

func (h *DatosFiliatorios) SetDomicilioCodPcia(domicilioCodPcia string) {
	h.domicilioCodPcia = domicilioCodPcia
}

func (h *DatosFiliatorios) SetDomicilioAdic1(domicilioAdic1 string) {
	h.domicilioAdic1 = domicilioAdic1
}

func (h *DatosFiliatorios) SetDomicilioAdic2(domicilioAdic2 string) {
	h.domicilioAdic2 = domicilioAdic2
}

func (h *DatosFiliatorios) SetLocalidad(localidad string) {
	h.localidad = localidad
}

func (h *DatosFiliatorios) SetPaisOtorg(paisOtorg string) {
	h.paisOtorg = paisOtorg
}

func (h *DatosFiliatorios) SetTipoTarjeta(tipoTarjeta string) {
	h.tipoTarjeta = tipoTarjeta
}

func (h *DatosFiliatorios) SetDenomTipoTarj(denomTipoTarj string) {
	h.denomTipoTarj = denomTipoTarj
}

func NewDatosFiliatorios() *DatosFiliatorios {
	return &DatosFiliatorios{}
}

func (h *DatosFiliatorios) NombreSocio() string {
	return h.nombreSocio
}

func (h *DatosFiliatorios) NroCtaEmpresa() string {
	return h.nroCtaEmpresa
}

func (h *DatosFiliatorios) NombreCtaEmpresa() string {
	return h.nombreCtaEmpresa
}

func (h *DatosFiliatorios) DomicilioCodPcia() string {
	return h.domicilioCodPcia
}

func (h *DatosFiliatorios) DomicilioAdic1() string {
	return h.domicilioAdic1
}

func (h *DatosFiliatorios) DomicilioAdic2() string {
	return h.domicilioAdic2
}

func (h *DatosFiliatorios) Localidad() string {
	return h.localidad
}

func (h *DatosFiliatorios) PaisOtorg() string {
	return h.paisOtorg
}

func (h *DatosFiliatorios) TipoTarjeta() string {
	return h.tipoTarjeta
}

func (h *DatosFiliatorios) DenomTipoTarj() string {
	return h.denomTipoTarj
}
