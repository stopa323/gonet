package language

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// ConfigFromFile loads config from given path
func ConfigFromFile(path string) (*Config, hcl.Diagnostics) {
	log.Debugf("parsing HCL file: %s", path)
	parser := hclparse.NewParser()
	f, parserDiags := parser.ParseHCLFile(path)
	if parserDiags.HasErrors() {
		return nil, parserDiags
	}

	rootSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "connection", LabelNames: []string{"kind", "name"}},
		},
	}

	content, contentDiags := f.Body.Content(rootSchema)
	if contentDiags.HasErrors() {
		return nil, contentDiags
	}

	log.Debug("no syntax errors found. Beginning content unmarshaling...")
	cfg := NewConfig()
	var diags hcl.Diagnostics
	for _, block := range content.Blocks {
		switch block.Type {
		case "connection":
			connectionKind := block.Labels[0]
			connectionName := block.Labels[1]

			if !isValidConnectionKind(connectionKind) {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "invalid connection kind",
					Detail:   fmt.Sprintf("%s is invalid connection kind", connectionKind),
				})
				return nil, diags
			}
			diags = cfg.AddConnection(connectionKind, connectionName,
				block.Body)
			if diags.HasErrors() {
				return nil, diags
			}
		default:
			// This should not be reached
			panic(fmt.Sprintf("unknown block type: %s", block.Type))
		}
	}
	log.Infof("loaded network config from: %s", path)
	return cfg, diags
}
