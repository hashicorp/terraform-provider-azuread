package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRegionGroup string

const (
	CloudPCRegionGroup_Asia            CloudPCRegionGroup = "asia"
	CloudPCRegionGroup_Australia       CloudPCRegionGroup = "australia"
	CloudPCRegionGroup_Canada          CloudPCRegionGroup = "canada"
	CloudPCRegionGroup_Default         CloudPCRegionGroup = "default"
	CloudPCRegionGroup_Euap            CloudPCRegionGroup = "euap"
	CloudPCRegionGroup_EuropeUnion     CloudPCRegionGroup = "europeUnion"
	CloudPCRegionGroup_France          CloudPCRegionGroup = "france"
	CloudPCRegionGroup_Germany         CloudPCRegionGroup = "germany"
	CloudPCRegionGroup_India           CloudPCRegionGroup = "india"
	CloudPCRegionGroup_Japan           CloudPCRegionGroup = "japan"
	CloudPCRegionGroup_Norway          CloudPCRegionGroup = "norway"
	CloudPCRegionGroup_SouthAmerica    CloudPCRegionGroup = "southAmerica"
	CloudPCRegionGroup_SouthKorea      CloudPCRegionGroup = "southKorea"
	CloudPCRegionGroup_Switzerland     CloudPCRegionGroup = "switzerland"
	CloudPCRegionGroup_UnitedKingdom   CloudPCRegionGroup = "unitedKingdom"
	CloudPCRegionGroup_UsCentral       CloudPCRegionGroup = "usCentral"
	CloudPCRegionGroup_UsEast          CloudPCRegionGroup = "usEast"
	CloudPCRegionGroup_UsGovernment    CloudPCRegionGroup = "usGovernment"
	CloudPCRegionGroup_UsGovernmentDOD CloudPCRegionGroup = "usGovernmentDOD"
	CloudPCRegionGroup_UsWest          CloudPCRegionGroup = "usWest"
)

func PossibleValuesForCloudPCRegionGroup() []string {
	return []string{
		string(CloudPCRegionGroup_Asia),
		string(CloudPCRegionGroup_Australia),
		string(CloudPCRegionGroup_Canada),
		string(CloudPCRegionGroup_Default),
		string(CloudPCRegionGroup_Euap),
		string(CloudPCRegionGroup_EuropeUnion),
		string(CloudPCRegionGroup_France),
		string(CloudPCRegionGroup_Germany),
		string(CloudPCRegionGroup_India),
		string(CloudPCRegionGroup_Japan),
		string(CloudPCRegionGroup_Norway),
		string(CloudPCRegionGroup_SouthAmerica),
		string(CloudPCRegionGroup_SouthKorea),
		string(CloudPCRegionGroup_Switzerland),
		string(CloudPCRegionGroup_UnitedKingdom),
		string(CloudPCRegionGroup_UsCentral),
		string(CloudPCRegionGroup_UsEast),
		string(CloudPCRegionGroup_UsGovernment),
		string(CloudPCRegionGroup_UsGovernmentDOD),
		string(CloudPCRegionGroup_UsWest),
	}
}

func (s *CloudPCRegionGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRegionGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRegionGroup(input string) (*CloudPCRegionGroup, error) {
	vals := map[string]CloudPCRegionGroup{
		"asia":            CloudPCRegionGroup_Asia,
		"australia":       CloudPCRegionGroup_Australia,
		"canada":          CloudPCRegionGroup_Canada,
		"default":         CloudPCRegionGroup_Default,
		"euap":            CloudPCRegionGroup_Euap,
		"europeunion":     CloudPCRegionGroup_EuropeUnion,
		"france":          CloudPCRegionGroup_France,
		"germany":         CloudPCRegionGroup_Germany,
		"india":           CloudPCRegionGroup_India,
		"japan":           CloudPCRegionGroup_Japan,
		"norway":          CloudPCRegionGroup_Norway,
		"southamerica":    CloudPCRegionGroup_SouthAmerica,
		"southkorea":      CloudPCRegionGroup_SouthKorea,
		"switzerland":     CloudPCRegionGroup_Switzerland,
		"unitedkingdom":   CloudPCRegionGroup_UnitedKingdom,
		"uscentral":       CloudPCRegionGroup_UsCentral,
		"useast":          CloudPCRegionGroup_UsEast,
		"usgovernment":    CloudPCRegionGroup_UsGovernment,
		"usgovernmentdod": CloudPCRegionGroup_UsGovernmentDOD,
		"uswest":          CloudPCRegionGroup_UsWest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRegionGroup(input)
	return &out, nil
}
