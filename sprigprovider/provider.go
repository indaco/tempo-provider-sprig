package sprigprovider

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/indaco/tempo-api/templatefuncs"
)

// SprigProvider implements TemplateFuncProvider using Masterminds/sprig.
type SprigProvider struct{}

// GetFunctions returns all Sprig template functions.
func (p *SprigProvider) GetFunctions() template.FuncMap {
	return sprig.FuncMap()
}

// Provider is the global instance of the SprigProvider.
var Provider templatefuncs.TemplateFuncProvider = &SprigProvider{}
