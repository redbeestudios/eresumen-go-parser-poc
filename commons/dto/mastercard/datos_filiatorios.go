package mastercard

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
