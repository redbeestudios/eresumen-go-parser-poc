package mastercard

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
