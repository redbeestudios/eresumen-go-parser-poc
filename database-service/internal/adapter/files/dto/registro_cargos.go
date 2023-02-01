package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type RegistroCargos struct {
	Descripcion    string `json:"descripcion,omitempty"`
	Importe        string `json:"importe,omitempty"`
	ImporteDolares string `json:"importe_dolares,omitempty"`
}

func (h RegistroCargos) toDomain() *contenido_resumen.RegistroCargos {
	r := contenido_resumen.NewRegistroCargos()
	r.SetDescripcion(h.Descripcion)
	r.SetImporte(h.Importe)
	r.SetImporteDolares(h.ImporteDolares)
	return r
}

func registroCargosFromDomain(h *contenido_resumen.RegistroCargos) *RegistroCargos {
	if h == nil {
		return nil
	}
	return &RegistroCargos{
		Descripcion:    h.Descripcion(),
		Importe:        h.Importe(),
		ImporteDolares: h.ImporteDolares(),
	}
}
