package cmdcommons

import (
	"github.com/puzpuzpuz/xsync"
	"time"
)

type Process struct {
	StartTime time.Time
	Count     *xsync.Counter
}
