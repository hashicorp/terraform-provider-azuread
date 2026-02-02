package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRegion string

const (
	NetworkaccessRegion_AustraliaEast      NetworkaccessRegion = "australiaEast"
	NetworkaccessRegion_AustraliaSouthEast NetworkaccessRegion = "australiaSouthEast"
	NetworkaccessRegion_BrazilSouth        NetworkaccessRegion = "brazilSouth"
	NetworkaccessRegion_CanadaCentral      NetworkaccessRegion = "canadaCentral"
	NetworkaccessRegion_CanadaEast         NetworkaccessRegion = "canadaEast"
	NetworkaccessRegion_CentralIndia       NetworkaccessRegion = "centralIndia"
	NetworkaccessRegion_CentralUS          NetworkaccessRegion = "centralUS"
	NetworkaccessRegion_EastUS             NetworkaccessRegion = "eastUS"
	NetworkaccessRegion_EastUS2            NetworkaccessRegion = "eastUS2"
	NetworkaccessRegion_FranceCentral      NetworkaccessRegion = "franceCentral"
	NetworkaccessRegion_FranceSouth        NetworkaccessRegion = "franceSouth"
	NetworkaccessRegion_GermanyWestCentral NetworkaccessRegion = "germanyWestCentral"
	NetworkaccessRegion_IsraelCentral      NetworkaccessRegion = "israelCentral"
	NetworkaccessRegion_ItalyNorth         NetworkaccessRegion = "italyNorth"
	NetworkaccessRegion_JapanEast          NetworkaccessRegion = "japanEast"
	NetworkaccessRegion_JapanWest          NetworkaccessRegion = "japanWest"
	NetworkaccessRegion_KoreaCentral       NetworkaccessRegion = "koreaCentral"
	NetworkaccessRegion_KoreaSouth         NetworkaccessRegion = "koreaSouth"
	NetworkaccessRegion_NorthCentralUS     NetworkaccessRegion = "northCentralUS"
	NetworkaccessRegion_NorthEurope        NetworkaccessRegion = "northEurope"
	NetworkaccessRegion_PolandCentral      NetworkaccessRegion = "polandCentral"
	NetworkaccessRegion_SouthAfricaNorth   NetworkaccessRegion = "southAfricaNorth"
	NetworkaccessRegion_SouthAfricaWest    NetworkaccessRegion = "southAfricaWest"
	NetworkaccessRegion_SouthCentralUS     NetworkaccessRegion = "southCentralUS"
	NetworkaccessRegion_SouthEastAsia      NetworkaccessRegion = "southEastAsia"
	NetworkaccessRegion_SouthIndia         NetworkaccessRegion = "southIndia"
	NetworkaccessRegion_SwedenCentral      NetworkaccessRegion = "swedenCentral"
	NetworkaccessRegion_SwitzerlandNorth   NetworkaccessRegion = "switzerlandNorth"
	NetworkaccessRegion_UaeNorth           NetworkaccessRegion = "uaeNorth"
	NetworkaccessRegion_UkSouth            NetworkaccessRegion = "ukSouth"
	NetworkaccessRegion_WestCentralUS      NetworkaccessRegion = "westCentralUS"
	NetworkaccessRegion_WestEurope         NetworkaccessRegion = "westEurope"
	NetworkaccessRegion_WestUS             NetworkaccessRegion = "westUS"
	NetworkaccessRegion_WestUS2            NetworkaccessRegion = "westUS2"
	NetworkaccessRegion_WestUS3            NetworkaccessRegion = "westUS3"
)

func PossibleValuesForNetworkaccessRegion() []string {
	return []string{
		string(NetworkaccessRegion_AustraliaEast),
		string(NetworkaccessRegion_AustraliaSouthEast),
		string(NetworkaccessRegion_BrazilSouth),
		string(NetworkaccessRegion_CanadaCentral),
		string(NetworkaccessRegion_CanadaEast),
		string(NetworkaccessRegion_CentralIndia),
		string(NetworkaccessRegion_CentralUS),
		string(NetworkaccessRegion_EastUS),
		string(NetworkaccessRegion_EastUS2),
		string(NetworkaccessRegion_FranceCentral),
		string(NetworkaccessRegion_FranceSouth),
		string(NetworkaccessRegion_GermanyWestCentral),
		string(NetworkaccessRegion_IsraelCentral),
		string(NetworkaccessRegion_ItalyNorth),
		string(NetworkaccessRegion_JapanEast),
		string(NetworkaccessRegion_JapanWest),
		string(NetworkaccessRegion_KoreaCentral),
		string(NetworkaccessRegion_KoreaSouth),
		string(NetworkaccessRegion_NorthCentralUS),
		string(NetworkaccessRegion_NorthEurope),
		string(NetworkaccessRegion_PolandCentral),
		string(NetworkaccessRegion_SouthAfricaNorth),
		string(NetworkaccessRegion_SouthAfricaWest),
		string(NetworkaccessRegion_SouthCentralUS),
		string(NetworkaccessRegion_SouthEastAsia),
		string(NetworkaccessRegion_SouthIndia),
		string(NetworkaccessRegion_SwedenCentral),
		string(NetworkaccessRegion_SwitzerlandNorth),
		string(NetworkaccessRegion_UaeNorth),
		string(NetworkaccessRegion_UkSouth),
		string(NetworkaccessRegion_WestCentralUS),
		string(NetworkaccessRegion_WestEurope),
		string(NetworkaccessRegion_WestUS),
		string(NetworkaccessRegion_WestUS2),
		string(NetworkaccessRegion_WestUS3),
	}
}

func (s *NetworkaccessRegion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessRegion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessRegion(input string) (*NetworkaccessRegion, error) {
	vals := map[string]NetworkaccessRegion{
		"australiaeast":      NetworkaccessRegion_AustraliaEast,
		"australiasoutheast": NetworkaccessRegion_AustraliaSouthEast,
		"brazilsouth":        NetworkaccessRegion_BrazilSouth,
		"canadacentral":      NetworkaccessRegion_CanadaCentral,
		"canadaeast":         NetworkaccessRegion_CanadaEast,
		"centralindia":       NetworkaccessRegion_CentralIndia,
		"centralus":          NetworkaccessRegion_CentralUS,
		"eastus":             NetworkaccessRegion_EastUS,
		"eastus2":            NetworkaccessRegion_EastUS2,
		"francecentral":      NetworkaccessRegion_FranceCentral,
		"francesouth":        NetworkaccessRegion_FranceSouth,
		"germanywestcentral": NetworkaccessRegion_GermanyWestCentral,
		"israelcentral":      NetworkaccessRegion_IsraelCentral,
		"italynorth":         NetworkaccessRegion_ItalyNorth,
		"japaneast":          NetworkaccessRegion_JapanEast,
		"japanwest":          NetworkaccessRegion_JapanWest,
		"koreacentral":       NetworkaccessRegion_KoreaCentral,
		"koreasouth":         NetworkaccessRegion_KoreaSouth,
		"northcentralus":     NetworkaccessRegion_NorthCentralUS,
		"northeurope":        NetworkaccessRegion_NorthEurope,
		"polandcentral":      NetworkaccessRegion_PolandCentral,
		"southafricanorth":   NetworkaccessRegion_SouthAfricaNorth,
		"southafricawest":    NetworkaccessRegion_SouthAfricaWest,
		"southcentralus":     NetworkaccessRegion_SouthCentralUS,
		"southeastasia":      NetworkaccessRegion_SouthEastAsia,
		"southindia":         NetworkaccessRegion_SouthIndia,
		"swedencentral":      NetworkaccessRegion_SwedenCentral,
		"switzerlandnorth":   NetworkaccessRegion_SwitzerlandNorth,
		"uaenorth":           NetworkaccessRegion_UaeNorth,
		"uksouth":            NetworkaccessRegion_UkSouth,
		"westcentralus":      NetworkaccessRegion_WestCentralUS,
		"westeurope":         NetworkaccessRegion_WestEurope,
		"westus":             NetworkaccessRegion_WestUS,
		"westus2":            NetworkaccessRegion_WestUS2,
		"westus3":            NetworkaccessRegion_WestUS3,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessRegion(input)
	return &out, nil
}
