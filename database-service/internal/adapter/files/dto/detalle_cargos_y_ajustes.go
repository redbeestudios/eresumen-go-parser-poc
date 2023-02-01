package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type DetalleCargosYAjustes struct {
	CodigoMov       string `json:"codigo_mov,omitempty"`
	SubcodMov       string `json:"subcod_mov,omitempty"`
	CodFinanc       string `json:"cod_financ,omitempty"`
	FecOper         string `json:"fec_oper,omitempty"`
	Cupon           string `json:"cupon,omitempty"`
	NroComercio     string `json:"nro_comercio,omitempty"`
	Detalle         string `json:"detalle,omitempty"`
	Importe         string `json:"importe,omitempty"`
	ImporteDol      string `json:"importe_dol,omitempty"`
	NroAutoriz      string `json:"nro_autoriz,omitempty"`
	Concepto        string `json:"concepto,omitempty"`
	FecPresen       string `json:"fec_presen,omitempty"`
	MarcaCF         string `json:"marca_cf,omitempty"`
	CodMoneda       string `json:"cod_moneda,omitempty"`
	CapitalAdeudado string `json:"capital_adeudado,omitempty"`
	InteresCuota    string `json:"interes_cuota,omitempty"`
	TnaPr           string `json:"tna_pr,omitempty"`
	Vigencia        string `json:"vigencia,omitempty"`
}

func (h DetalleCargosYAjustes) toDomain() *contenido_resumen.DetalleCargosYAjustes {

	d := contenido_resumen.NewDetalleCargosYAjustes()
	d.SetCodigoMov(h.CodigoMov)
	d.SetSubcodMov(h.SubcodMov)
	d.SetCodFinanc(h.CodFinanc)
	d.SetFecOper(h.FecOper)
	d.SetCupon(h.Cupon)
	d.SetNroComercio(h.NroComercio)
	d.SetDetalle(h.Detalle)
	d.SetImporte(h.Importe)
	d.SetImporteDol(h.ImporteDol)
	d.SetNroAutoriz(h.NroAutoriz)
	d.SetConcepto(h.Concepto)
	d.SetFecPresen(h.FecPresen)
	d.SetMarcaCF(h.MarcaCF)
	d.SetCodMoneda(h.CodMoneda)
	d.SetCapitalAdeudado(h.CapitalAdeudado)
	d.SetInteresCuota(h.InteresCuota)
	d.SetTnaPr(h.TnaPr)
	d.SetVigencia(h.Vigencia)

	return d
}

func detalleCargosYAjustesFromDomain(h *contenido_resumen.DetalleCargosYAjustes) *DetalleCargosYAjustes {
	if h == nil {
		return nil
	}
	return &DetalleCargosYAjustes{
		CodigoMov:       h.CodigoMov(),
		SubcodMov:       h.SubcodMov(),
		CodFinanc:       h.CodFinanc(),
		FecOper:         h.FecOper(),
		Cupon:           h.Cupon(),
		NroComercio:     h.NroComercio(),
		Detalle:         h.Detalle(),
		Importe:         h.Importe(),
		ImporteDol:      h.ImporteDol(),
		NroAutoriz:      h.NroAutoriz(),
		Concepto:        h.Concepto(),
		FecPresen:       h.FecPresen(),
		MarcaCF:         h.MarcaCF(),
		CodMoneda:       h.CodMoneda(),
		CapitalAdeudado: h.CapitalAdeudado(),
		InteresCuota:    h.InteresCuota(),
		TnaPr:           h.TnaPr(),
		Vigencia:        h.Vigencia(),
	}
}
