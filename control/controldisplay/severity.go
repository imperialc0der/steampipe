package controldisplay

import (
	"fmt"
	"log"
	"strings"

	"github.com/turbot/go-kit/helpers"
)

const severityMaxLen = len("CRITICAL")

type SeverityRenderer struct {
	severity string
}

func NewSeverityRenderer(severity string) *SeverityRenderer {
	return &SeverityRenderer{
		strings.ToUpper(severity),
	}
}

// Render returns ther severity oin upper case, got 'critical' and 'high' severities
// for all other values an empty string is returned
// NOTE: adds a trailing space
func (r SeverityRenderer) Render() string {
	log.Println("[TRACE] begin severity render")
	defer log.Println("[TRACE] end severity render")

	if helpers.StringSliceContains([]string{"CRITICAL", "HIGH"}, r.severity) {
		return fmt.Sprintf("%s ", ControlColors.Severity(r.severity))
	}
	return ""
}
