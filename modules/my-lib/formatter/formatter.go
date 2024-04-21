package formatter

import (
	"strings"
	// line 6 generate a compile error: use of invalid internal package
	"github.com/johnnydacosta/go-labs/modules/my-lib/math/internal/randutils"
)

func Upper(s string) string {
	randutils.Rand()
	return strings.ToUpper(s)
}
