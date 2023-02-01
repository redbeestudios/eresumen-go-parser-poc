package usecase

import (
	"bufio"
	"fmt"
	"github.com/puzpuzpuz/xsync"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"golang.org/x/text/encoding/charmap"
	"log"
	"os"
	"time"
)

type SummaryFileReader struct{}

func (c *SummaryFileReader) readSummaryLines(path string) (*xsync.MapOf[string, []string], error) {
	var accountNumbers = xsync.NewMapOf[[]string]()

	operationStart := time.Now()

	file, err := os.Open(fmt.Sprintf("%s/%s", path, "ERESUMEN"))
	if err != nil {
		log.Printf("Error opening file path %s: %s", path, err.Error())
		return nil, err
	}

	decoder := charmap.ISO8859_1.NewDecoder()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		lineaResumen, err := decoder.String(scanner.Text())

		if err != nil {
			log.Printf("Error converting string to utf-8: %s", err)
			continue
		}

		nroRegistro := lineaResumen[0:5]
		nroSecuencia := lineaResumen[5:10]

		campoControl := nroRegistro + nroSecuencia

		if campoControl != "0000000000" && campoControl != "9999999999" {
			clave := domain.ClaveFromString(lineaResumen)
			lineas, _ := accountNumbers.Load(clave.NroCuenta())
			accountNumbers.Store(clave.NroCuenta(), append(lineas, lineaResumen))

		}

	}

	file.Close()

	log.Printf("Operation Time: %s", time.Since(operationStart))
	return accountNumbers, nil
}

func NewSummaryFileReader() *SummaryFileReader {
	return &SummaryFileReader{}
}
