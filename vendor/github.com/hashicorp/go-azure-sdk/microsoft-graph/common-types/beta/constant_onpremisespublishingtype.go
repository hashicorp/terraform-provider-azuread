package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingType string

const (
	OnPremisesPublishingType_ApplicationProxy OnPremisesPublishingType = "applicationProxy"
	OnPremisesPublishingType_Authentication   OnPremisesPublishingType = "authentication"
	OnPremisesPublishingType_ExchangeOnline   OnPremisesPublishingType = "exchangeOnline"
	OnPremisesPublishingType_IntunePfx        OnPremisesPublishingType = "intunePfx"
	OnPremisesPublishingType_OflineDomainJoin OnPremisesPublishingType = "oflineDomainJoin"
	OnPremisesPublishingType_Provisioning     OnPremisesPublishingType = "provisioning"
)

func PossibleValuesForOnPremisesPublishingType() []string {
	return []string{
		string(OnPremisesPublishingType_ApplicationProxy),
		string(OnPremisesPublishingType_Authentication),
		string(OnPremisesPublishingType_ExchangeOnline),
		string(OnPremisesPublishingType_IntunePfx),
		string(OnPremisesPublishingType_OflineDomainJoin),
		string(OnPremisesPublishingType_Provisioning),
	}
}

func (s *OnPremisesPublishingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesPublishingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesPublishingType(input string) (*OnPremisesPublishingType, error) {
	vals := map[string]OnPremisesPublishingType{
		"applicationproxy": OnPremisesPublishingType_ApplicationProxy,
		"authentication":   OnPremisesPublishingType_Authentication,
		"exchangeonline":   OnPremisesPublishingType_ExchangeOnline,
		"intunepfx":        OnPremisesPublishingType_IntunePfx,
		"oflinedomainjoin": OnPremisesPublishingType_OflineDomainJoin,
		"provisioning":     OnPremisesPublishingType_Provisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesPublishingType(input)
	return &out, nil
}
