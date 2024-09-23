package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorGroupRegion string

const (
	ConnectorGroupRegion_Asia ConnectorGroupRegion = "asia"
	ConnectorGroupRegion_Aus  ConnectorGroupRegion = "aus"
	ConnectorGroupRegion_Eur  ConnectorGroupRegion = "eur"
	ConnectorGroupRegion_Ind  ConnectorGroupRegion = "ind"
	ConnectorGroupRegion_Nam  ConnectorGroupRegion = "nam"
)

func PossibleValuesForConnectorGroupRegion() []string {
	return []string{
		string(ConnectorGroupRegion_Asia),
		string(ConnectorGroupRegion_Aus),
		string(ConnectorGroupRegion_Eur),
		string(ConnectorGroupRegion_Ind),
		string(ConnectorGroupRegion_Nam),
	}
}

func (s *ConnectorGroupRegion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorGroupRegion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorGroupRegion(input string) (*ConnectorGroupRegion, error) {
	vals := map[string]ConnectorGroupRegion{
		"asia": ConnectorGroupRegion_Asia,
		"aus":  ConnectorGroupRegion_Aus,
		"eur":  ConnectorGroupRegion_Eur,
		"ind":  ConnectorGroupRegion_Ind,
		"nam":  ConnectorGroupRegion_Nam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorGroupRegion(input)
	return &out, nil
}
