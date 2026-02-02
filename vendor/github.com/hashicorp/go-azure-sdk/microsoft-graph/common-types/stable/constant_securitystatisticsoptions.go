package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityStatisticsOptions string

const (
	SecurityStatisticsOptions_AdvancedIndexing      SecurityStatisticsOptions = "advancedIndexing"
	SecurityStatisticsOptions_IncludeQueryStats     SecurityStatisticsOptions = "includeQueryStats"
	SecurityStatisticsOptions_IncludeRefiners       SecurityStatisticsOptions = "includeRefiners"
	SecurityStatisticsOptions_IncludeUnindexedStats SecurityStatisticsOptions = "includeUnindexedStats"
	SecurityStatisticsOptions_LocationsWithoutHits  SecurityStatisticsOptions = "locationsWithoutHits"
)

func PossibleValuesForSecurityStatisticsOptions() []string {
	return []string{
		string(SecurityStatisticsOptions_AdvancedIndexing),
		string(SecurityStatisticsOptions_IncludeQueryStats),
		string(SecurityStatisticsOptions_IncludeRefiners),
		string(SecurityStatisticsOptions_IncludeUnindexedStats),
		string(SecurityStatisticsOptions_LocationsWithoutHits),
	}
}

func (s *SecurityStatisticsOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityStatisticsOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityStatisticsOptions(input string) (*SecurityStatisticsOptions, error) {
	vals := map[string]SecurityStatisticsOptions{
		"advancedindexing":      SecurityStatisticsOptions_AdvancedIndexing,
		"includequerystats":     SecurityStatisticsOptions_IncludeQueryStats,
		"includerefiners":       SecurityStatisticsOptions_IncludeRefiners,
		"includeunindexedstats": SecurityStatisticsOptions_IncludeUnindexedStats,
		"locationswithouthits":  SecurityStatisticsOptions_LocationsWithoutHits,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityStatisticsOptions(input)
	return &out, nil
}
