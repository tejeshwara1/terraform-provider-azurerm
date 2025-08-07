package pluginsdk

import (
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// PopulatedResourceIDFromString takes a string and returns a ResourceId with all values populated.
// If any errors are encountered, this returns nil.
func PopulatedResourceIDFromString(input string, insensitively bool) resourceids.ResourceId {
	id := recaser.ResourceIdTypeFromResourceId(input)
	if id == nil {
		return nil
	}

	parser := resourceids.NewParserFromResourceIdType(id)
	parseResult, err := parser.Parse(input, insensitively)
	if err != nil {
		return nil
	}

	if err := id.FromParseResult(*parseResult); err != nil {
		return nil
	}

	// Sanity check
	// If for whatever reason the input string and the formatted ID don't match case-insensitively, return nil
	if !strings.EqualFold(input, id.ID()) {
		return nil
	}

	return id
}
