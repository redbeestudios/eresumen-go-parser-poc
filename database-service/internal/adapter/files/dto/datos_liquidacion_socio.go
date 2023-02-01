package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type DatosLiquidacionSocio struct {
	Cuit                    string `json:"cuit,omitempty"`
	CategoriaIva            string `json:"categoria_iva,omitempty"`
	DescripcionCategoriaIva string `json:"descripcion_categoria_iva,omitempty"`
	SbmTasas                string `json:"sbm_tasas,omitempty"`
	SbmGtosSV               string `json:"sbm_gtos_sv,omitempty"`
	SbmParam                string `json:"sbm_param,omitempty"`
	DvSocio                 string `json:"dv_socio,omitempty"`
	MarcaCBU                string `json:"marca_cbu,omitempty"`
	Cbu                     string `json:"cbu,omitempty"`
	DvUnifSuc               string `json:"dv_unif_suc,omitempty"`
	LeyDebit                string `json:"ley_debit,omitempty"`
	TipCta                  string `json:"tip_cta,omitempty"`
	NroCtaDebito            string `json:"nro_cta_debito,omitempty"`
	TipPago                 string `json:"tip_pago,omitempty"`
	TipoCtaCte              string `json:"tipo_cta_cte,omitempty"`
	TarjetaTitular          string `json:"tarjeta_titular,omitempty"`
	RegPagYAjSant           string `json:"reg_pag_y_aj_sant,omitempty"`
	RegDetalleYMsg          string `json:"reg_detalle_y_msg,omitempty"`
	CantidadHojas           string `json:"cantidad_hojas,omitempty"`
	TipoProducto            string `json:"tipo_producto,omitempty"`
	Comprobante             string `json:"comprobante,omitempty"`
}

func (h DatosLiquidacionSocio) toDomain() *contenido_resumen.DatosLiquidacionSocio {
	d := contenido_resumen.NewDatosLiquidacionSocio()

	d.SetCuit(h.Cuit)
	d.SetCategoriaIva(h.CategoriaIva)
	d.SetDescripcionCategoriaIva(h.DescripcionCategoriaIva)
	d.SetSbmTasas(h.SbmTasas)
	d.SetSbmGtosSV(h.SbmGtosSV)
	d.SetSbmParam(h.SbmParam)
	d.SetDvSocio(h.DvSocio)
	d.SetMarcaCBU(h.MarcaCBU)
	d.SetCbu(h.Cbu)
	d.SetDvUnifSuc(h.DvUnifSuc)
	d.SetLeyDebit(h.LeyDebit)
	d.SetTipCta(h.TipCta)
	d.SetNroCtaDebito(h.NroCtaDebito)
	d.SetTipPago(h.TipPago)
	d.SetTipoCtaCte(h.TipoCtaCte)
	d.SetTarjetaTitular(h.TarjetaTitular)
	d.SetRegPagYAjSant(h.RegPagYAjSant)
	d.SetRegDetalleYMsg(h.RegDetalleYMsg)
	d.SetCantidadHojas(h.CantidadHojas)
	d.SetTipoProducto(h.TipoProducto)
	d.SetComprobante(h.Comprobante)

	return d
}

func datosLiquidacionSocioFromDomain(h *contenido_resumen.DatosLiquidacionSocio) *DatosLiquidacionSocio {
	if h == nil {
		return nil
	}
	return &DatosLiquidacionSocio{
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
