package tipo_registro

type TipoRegistro int

const (
	HeaderControl TipoRegistro = iota + 1
	TrailerControl
	DatosFiliatorios
	DatosLiquidacionSocio
	DatosDeAutorizado
	DatosMesAnterior
	RegistroTotales
	DatosMesActual
	DetalleCargosComprasYAjustes
	RegistroCargos
	Leyendas
	Limites
	Tasas
	DeudasNoVencidasSemestre
	Invalid
)

func NewTipoRegistro(tipoRegistro string) TipoRegistro {
	switch tipoRegistro {
	case "00000":
		return HeaderControl
	case "99999":
		return TrailerControl
	case "00010":
		return DatosFiliatorios
	case "00020":
		return DatosLiquidacionSocio
	case "00030":
		return DatosDeAutorizado
	case "00040":
		return DatosMesAnterior
	case "00050":
		return RegistroTotales
	case "00060":
		return DatosMesActual
	case "00080":
		return DetalleCargosComprasYAjustes
	case "00090":
		return RegistroCargos
	case "00100":
		return Leyendas
	case "00110":
		return Limites
	case "00120":
		return Tasas
	case "00130":
		return DeudasNoVencidasSemestre
	default:
		return Invalid
	}
}

func (t TipoRegistro) String() string {
	switch t {
	case HeaderControl:
		return "00000"
	case TrailerControl:
		return "99999"
	case DatosFiliatorios:
		return "00010"
	case DatosLiquidacionSocio:
		return "00020"
	case DatosDeAutorizado:
		return "00030"
	case DatosMesAnterior:
		return "00040"
	case RegistroTotales:
		return "00050"
	case DatosMesActual:
		return "00060"
	case DetalleCargosComprasYAjustes:
		return "00080"
	case RegistroCargos:
		return "00090"
	case Leyendas:
		return "00100"
	case Limites:
		return "00110"
	case Tasas:
		return "00120"
	case DeudasNoVencidasSemestre:
		return "00130"
	default:
		return ""
	}
}
