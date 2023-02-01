package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type Header struct {
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

func (h Header) toDomain() *contenido_resumen.Clave {
	d := contenido_resumen.NewClave()
	d.SetNroRegistro(h.NroRegistro)
	d.SetNroSecuencia(h.NroSecuencia)
	d.SetNroOrden(h.NroOrden)
	d.SetEntidad(h.Entidad)
	d.SetSucursal(h.Sucursal)
	d.SetGrupoCuentaCorriente(h.GrupoCuentaCorriente)
	d.SetCalleYNumero(h.CalleYNumero)
	d.SetPiso(h.Piso)
	d.SetDepartamento(h.Departamento)
	d.SetCodigoPostal(h.CodigoPostal)
	d.SetRetencion(h.Retencion)
	d.SetVtoActual(h.VtoActual)
	d.SetGrupoAfinidad(h.GrupoAfinidad)
	d.SetCodProductoMarca(h.CodProductoMarca)
	d.SetModeloLiquidacion(h.ModeloLiquidacion)
	d.SetNroCuenta(h.NroCuenta)
	d.SetNroAutorizado(h.NroAutorizado)
	d.SetEResumen(h.EResumen)
	d.SetCantidadResumen(h.CantidadResumen)

	return d

}

func claveFromDomain(h *contenido_resumen.Clave) *Header {
	if h == nil {
		return nil
	}
	return &Header{
		NroRegistro:          h.NroRegistro(),
		NroSecuencia:         h.NroSecuencia(),
		NroOrden:             h.NroOrden(),
		Entidad:              h.Entidad(),
		Sucursal:             h.Sucursal(),
		GrupoCuentaCorriente: h.GrupoCtaCte(),
		CalleYNumero:         h.CalleYNro(),
		Piso:                 h.Piso(),
		Departamento:         h.Departamento(),
		CodigoPostal:         h.CodigoPostal(),
		Retencion:            h.Retencion(),
		VtoActual:            h.VtoActual(),
		GrupoAfinidad:        h.GrupoAfinidad(),
		CodProductoMarca:     h.CodProdMarca(),
		ModeloLiquidacion:    h.ModeloLiquidacion(),
		NroCuenta:            h.NroCuenta(),
		NroAutorizado:        h.NroAutorizado(),
		EResumen:             h.EResumen(),
		CantidadResumen:      h.CantResumen(),
	}
}
