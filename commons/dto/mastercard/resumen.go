package mastercard

type Resumen struct {
	Clave                 *Clave                   `json:"header,omitempty"`
	DatosFiliatorios      *DatosFiliatorios        `json:"datos_filiatorios,omitempty"`
	DatosLiquidacionSocio *DatosLiquidacionSocio   `json:"datos_liquidacion_socio,omitempty"`
	DatosDeAutorizado     []*DatosDeAutorizado     `json:"datos_de_autorizado,omitempty"`
	DatosMesAnterior      *DatosMesAnterior        `json:"datos_mes_anterior,omitempty"`
	RegistroTotales       []*RegistroTotales       `json:"registro_totales,omitempty"`
	DatosMesActual        *DatosMesActual          `json:"datos_mes_actual,omitempty"`
	DetalleCargosYAjustes []*DetalleCargosYAjustes `json:"detalle_cargos_y_ajustes,omitempty"`
	RegistroCargos        []*RegistroCargos        `json:"registro_cargos,omitempty"`
	Leyendas              []*Leyenda               `json:"leyendas,omitempty"`
	Limites               []*Limite                `json:"limites,omitempty"`
	Tasas                 []*Tasa                  `json:"tasas,omitempty"`
	DeudasNoVencidas      []*Deuda                 `json:"deudas_no_vencidas,omitempty"`
}
