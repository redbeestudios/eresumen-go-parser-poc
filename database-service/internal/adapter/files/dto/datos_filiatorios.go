package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type DatosFiliatorios struct {
	NombreSocio      string `json:"nombre_socio,omitempty"`
	NroCtaEmpresa    string `json:"nro_cta_empresa,omitempty"`
	NombreCtaEmpresa string `json:"nombre_cta_empresa,omitempty"`
	DomicilioCodPcia string `json:"domicilio_cod_pcia,omitempty"`
	DomicilioAdic1   string `json:"domicilio_adic_1,omitempty"`
	DomicilioAdic2   string `json:"domicilio_adic_2,omitempty"`
	Localidad        string `json:"localidad,omitempty"`
	PaisOtorg        string `json:"pais_otorg,omitempty"`
	TipoTarjeta      string `json:"tipo_tarjeta,omitempty"`
	DenomTipoTarj    string `json:"denom_tipo_tarj,omitempty"`
}

func (h DatosFiliatorios) toDomain() *contenido_resumen.DatosFiliatorios {
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

func datosFiliatoriosFromDomain(h *contenido_resumen.DatosFiliatorios) *DatosFiliatorios {
	if h == nil {
		return nil
	}
	return &DatosFiliatorios{
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
