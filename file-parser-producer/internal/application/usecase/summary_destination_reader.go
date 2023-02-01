package usecase

import (
	"bufio"
	"fmt"
	"github.com/puzpuzpuz/xsync"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/usecase/line_parser"
	"golang.org/x/text/encoding/charmap"
	"log"
	"os"
	"strings"
	"time"
)

type SummaryDestinationReader struct {
}

func (c *SummaryDestinationReader) getDestinations(path string) (*xsync.MapOf[string, *domain.ResumenAEnviar], error) {

	var summariesWithDestination = xsync.NewMapOf[*domain.ResumenAEnviar]()

	operationStart := time.Now()

	file, err := os.Open(fmt.Sprintf("%s/%s", path, "TJERESUOCM.txt"))
	if err != nil {
		log.Printf("Error opening file path %s: %s", path, err.Error())
		return nil, err
	}

	defer file.Close()

	decoder := charmap.ISO8859_1.NewDecoder()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	for scanner.Scan() {

		lineaDestinatarios, err := decoder.String(scanner.Text())

		if err != nil {
			log.Printf("Error converting string to utf-8: %s", err)
			continue
		}

		resumenAEnviar := line_parser.ResumenAEnviarFromMailing(lineaDestinatarios)

		key := strings.TrimLeft(resumenAEnviar.Destinatario().NumeroCuenta(), "0")

		summariesWithDestination.Store(key, resumenAEnviar)

	}

	log.Printf("Operation Time: %s", time.Since(operationStart))
	return summariesWithDestination, nil
}

func NewSummaryDestinationReader() *SummaryDestinationReader {
	return &SummaryDestinationReader{}
}
