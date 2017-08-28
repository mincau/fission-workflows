package typedvalues

import (
	"strings"

	"github.com/fission/fission-workflow/pkg/types"
)

type Parser interface {
	Parse(i interface{}, allowedTypes ...string) (*types.TypedValue, error)
}

type Formatter interface {
	Format(v *types.TypedValue) (interface{}, error)
}

type ParserFormatter interface {
	Parser
	Formatter
}

// Splits valueTypes of format '<language>/<type>' into (format, type)
func ParseType(valueType string) (format string, subType string) {
	parts := strings.SplitN(valueType, "/", 2)

	if len(parts) == 0 {
		return "", ""
	}

	if len(parts) == 1 {
		switch parts[0] {
		case TYPE_EXPRESSION:
			fallthrough
		case TYPE_RAW:
			return FORMAT_RESERVED, parts[0]
		default:
			return parts[0], ""
		}
	}

	return parts[0], parts[1]
}

func FormatType(parts ...string) string {
	return strings.Join(parts, "/")
}

func IsFormat(targetValueType string, format string) bool {
	f, _ := ParseType(targetValueType)
	return strings.EqualFold(f, format)
}

var DefaultParserFormatter = newDefaultParserFormatter()

func Parse(i interface{}, allowedTypes ...string) (*types.TypedValue, error) {
	return DefaultParserFormatter.Parse(i, allowedTypes...)
}

func Format(v *types.TypedValue) (interface{}, error) {
	return DefaultParserFormatter.Format(v)
}

func newDefaultParserFormatter() ParserFormatter {
	// TODO Less verbose
	jsPf := &JsonParserFormatter{}
	return NewComposedParserFormatter(map[string]ParserFormatter{
		FormatType(FORMAT_JSON, TYPE_BOOL):   jsPf,
		FormatType(FORMAT_JSON, TYPE_STRING): jsPf,
		FormatType(FORMAT_JSON, TYPE_ARRAY):  jsPf,
		FormatType(FORMAT_JSON, TYPE_OBJECT): jsPf,
		FormatType(TYPE_FLOW):                &ControlFlowParserFormatter{},
		FormatType(TYPE_EXPRESSION):          &ExprParserFormatter{},
		FormatType(TYPE_RAW):                 &RawParserFormatter{},
	}, []string{
		FormatType(FORMAT_JSON, TYPE_BOOL),
		FormatType(FORMAT_JSON, TYPE_STRING),
		FormatType(FORMAT_JSON, TYPE_ARRAY),
		FormatType(FORMAT_JSON, TYPE_OBJECT),
		FormatType(TYPE_FLOW),
		FormatType(TYPE_EXPRESSION),
		FormatType(TYPE_RAW),
	}...)
}
