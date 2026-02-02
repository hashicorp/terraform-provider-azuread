package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAdditionalDataOptions string

const (
	SecurityAdditionalDataOptions_AdvancedIndexing             SecurityAdditionalDataOptions = "advancedIndexing"
	SecurityAdditionalDataOptions_AllItemsInFolder             SecurityAdditionalDataOptions = "allItemsInFolder"
	SecurityAdditionalDataOptions_AllVersions                  SecurityAdditionalDataOptions = "allVersions"
	SecurityAdditionalDataOptions_HtmlTranscripts              SecurityAdditionalDataOptions = "htmlTranscripts"
	SecurityAdditionalDataOptions_LinkedFiles                  SecurityAdditionalDataOptions = "linkedFiles"
	SecurityAdditionalDataOptions_ListAttachments              SecurityAdditionalDataOptions = "listAttachments"
	SecurityAdditionalDataOptions_LocationsWithoutHits         SecurityAdditionalDataOptions = "locationsWithoutHits"
	SecurityAdditionalDataOptions_MessageConversationExpansion SecurityAdditionalDataOptions = "messageConversationExpansion"
)

func PossibleValuesForSecurityAdditionalDataOptions() []string {
	return []string{
		string(SecurityAdditionalDataOptions_AdvancedIndexing),
		string(SecurityAdditionalDataOptions_AllItemsInFolder),
		string(SecurityAdditionalDataOptions_AllVersions),
		string(SecurityAdditionalDataOptions_HtmlTranscripts),
		string(SecurityAdditionalDataOptions_LinkedFiles),
		string(SecurityAdditionalDataOptions_ListAttachments),
		string(SecurityAdditionalDataOptions_LocationsWithoutHits),
		string(SecurityAdditionalDataOptions_MessageConversationExpansion),
	}
}

func (s *SecurityAdditionalDataOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAdditionalDataOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAdditionalDataOptions(input string) (*SecurityAdditionalDataOptions, error) {
	vals := map[string]SecurityAdditionalDataOptions{
		"advancedindexing":             SecurityAdditionalDataOptions_AdvancedIndexing,
		"allitemsinfolder":             SecurityAdditionalDataOptions_AllItemsInFolder,
		"allversions":                  SecurityAdditionalDataOptions_AllVersions,
		"htmltranscripts":              SecurityAdditionalDataOptions_HtmlTranscripts,
		"linkedfiles":                  SecurityAdditionalDataOptions_LinkedFiles,
		"listattachments":              SecurityAdditionalDataOptions_ListAttachments,
		"locationswithouthits":         SecurityAdditionalDataOptions_LocationsWithoutHits,
		"messageconversationexpansion": SecurityAdditionalDataOptions_MessageConversationExpansion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAdditionalDataOptions(input)
	return &out, nil
}
