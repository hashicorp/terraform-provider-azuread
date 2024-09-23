package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportCriteria string

const (
	SecurityExportCriteria_PartiallyIndexed SecurityExportCriteria = "partiallyIndexed"
	SecurityExportCriteria_SearchHits       SecurityExportCriteria = "searchHits"
)

func PossibleValuesForSecurityExportCriteria() []string {
	return []string{
		string(SecurityExportCriteria_PartiallyIndexed),
		string(SecurityExportCriteria_SearchHits),
	}
}

func (s *SecurityExportCriteria) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityExportCriteria(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityExportCriteria(input string) (*SecurityExportCriteria, error) {
	vals := map[string]SecurityExportCriteria{
		"partiallyindexed": SecurityExportCriteria_PartiallyIndexed,
		"searchhits":       SecurityExportCriteria_SearchHits,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityExportCriteria(input)
	return &out, nil
}
