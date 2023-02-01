package mail

import (
	"bytes"
	"fmt"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/domain"
	"html/template"
)

func BuildMailBody(enviar *domain.Mail) string {
	htmlfile, err := template.ParseFiles("/home/rodrigo/dev/prisma/eresumen/go-parser/go-parser-poc/email-consumer/internal/application/resources/Template.html")

	if err != nil {
		return ""
	}

	var templateBytes bytes.Buffer

	mailParams := mailParams{
		Nombre: enviar.Nombre(),
		Fecha:  enviar.Fecha(),
		Url:    fmt.Sprintf("http://localhost:8085/eResumen/%s", enviar.CodigoVinculacion()),
	}

	err = htmlfile.Execute(&templateBytes, mailParams)
	if err != nil {
		return ""
	}
	return templateBytes.String()
}

type mailParams struct {
	Nombre string `json:"nombre,omitempty"`
	Fecha  string `json:"fecha,omitempty"`
	Url    string `json:"url,omitempty"`
}
