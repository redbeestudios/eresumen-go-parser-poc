package mastercard

type Clave struct {
	NroRegistro          string `json:"nro_registro,omitempty"`
	NroSecuencia         string `json:"nro_secuencia,omitempty"`
	NroOrden             string `json:"nro_orden,omitempty"`
	Entidad              string `json:"entidad,omitempty"`
	Sucursal             string `json:"sucursal,omitempty"`
	GrupoCuentaCorriente string `json:"grupo_cuenta_corriente,omitempty"`
	CalleYNumero         string `json:"calle_y_numero,omitempty"`
	Piso                 string `json:"piso,omitempty"`
	Departamento         string `json:"departamento,omitempty"`
	CodigoPostal         string `json:"codigo_postal,omitempty"`
	Retencion            string `json:"retencion,omitempty"`
	VtoActual            string `json:"vto_actual,omitempty"`
	GrupoAfinidad        string `json:"grupo_afinidad,omitempty"`
	CodProductoMarca     string `json:"cod_producto_marca,omitempty"`
	ModeloLiquidacion    string `json:"modelo_liquidacion,omitempty"`
	NroCuenta            string `json:"nro_cuenta,omitempty"`
	NroAutorizado        string `json:"nro_autorizado,omitempty"`
	EResumen             string `json:"e_resumen,omitempty"`
	CantidadResumen      string `json:"cantidad_resumen,omitempty"`
}
