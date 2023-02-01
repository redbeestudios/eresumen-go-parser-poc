package domain

import (
	"github.com/puzpuzpuz/xsync"
)

type Resultado struct {
	ResumenesAEnviar *xsync.MapOf[string, *ResumenAEnviar]
	DestSinLiq       []string
	LiqSinDest       []string
	OpCount          int
	ProcessId        string
}
