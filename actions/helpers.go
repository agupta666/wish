package actions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/agupta666/hash/utils"
)

func writeFileName(name string, ext string, w http.ResponseWriter) {
	fileName := fmt.Sprintf("%s%s", name, ext)
	value := fmt.Sprintf("attachment; filename=%s", fileName)
	w.Header().Set("Content-Disposition", value)
}

func writeMimeType(ext string, w http.ResponseWriter) {
	mimeType := utils.TypeByExtension(ext)
	w.Header().Set("Content-Type", mimeType)
}

// ParseComplete handlers a called after the expression parsing is complete
type ParseComplete func(args []string) (Action, error)

func parseExpr(expr, name string, handler ParseComplete) (Action, error) {
	expr = strings.TrimPrefix(expr, name+"[")
	expr = strings.TrimSuffix(expr, "]")
	args := utils.SplitAndTrim(expr, ",")

	return handler(args)
}
