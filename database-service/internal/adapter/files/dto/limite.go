package files

import "github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain/contenido_resumen"

type Limite struct {
	Descripcion string `json:"descripcion,omitempty"`
	Limite      string `json:"limite,omitempty"`
}

func (h Limite) toDomain() *contenido_resumen.Limite {
	l := contenido_resumen.NewLimite()
	l.SetLimite(h.Limite)
	l.SetDescripcion(h.Descripcion)

	return l
}

func limiteFromDomain(h *contenido_resumen.Limite) *Limite {
	if h == nil {
		return nil
	}
	return &Limite{
		Descripcion: h.Descripcion(),
		Limite:      h.Limite(),
	}
}
