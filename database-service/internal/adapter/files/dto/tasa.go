package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type Tasa struct {
	Descripcion   string `json:"descripcion,omitempty"`
	Tasa          string `json:"tasa,omitempty"`
	TasaEnDolares string `json:"tasa_en_dolares,omitempty"`
}

func (h Tasa) toDomain() *contenido_resumen.Tasa {
	t := contenido_resumen.NewTasa()
	t.SetDescripcion(h.Descripcion)
	t.SetTasa(h.Tasa)
	t.SetTasaEnDolares(h.TasaEnDolares)
	return t
}

func tasaFromDomain(h *contenido_resumen.Tasa) *Tasa {
	if h == nil {
		return nil
	}
	return &Tasa{
		Descripcion:   h.Descripcion(),
		Tasa:          h.Tasa(),
		TasaEnDolares: h.TasaEnDolares(),
	}
}
