package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkApplicationIdentityType string

const (
	TeamworkApplicationIdentityType_AadApplication     TeamworkApplicationIdentityType = "aadApplication"
	TeamworkApplicationIdentityType_Bot                TeamworkApplicationIdentityType = "bot"
	TeamworkApplicationIdentityType_Office365Connector TeamworkApplicationIdentityType = "office365Connector"
	TeamworkApplicationIdentityType_OutgoingWebhook    TeamworkApplicationIdentityType = "outgoingWebhook"
	TeamworkApplicationIdentityType_TenantBot          TeamworkApplicationIdentityType = "tenantBot"
)

func PossibleValuesForTeamworkApplicationIdentityType() []string {
	return []string{
		string(TeamworkApplicationIdentityType_AadApplication),
		string(TeamworkApplicationIdentityType_Bot),
		string(TeamworkApplicationIdentityType_Office365Connector),
		string(TeamworkApplicationIdentityType_OutgoingWebhook),
		string(TeamworkApplicationIdentityType_TenantBot),
	}
}

func (s *TeamworkApplicationIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkApplicationIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkApplicationIdentityType(input string) (*TeamworkApplicationIdentityType, error) {
	vals := map[string]TeamworkApplicationIdentityType{
		"aadapplication":     TeamworkApplicationIdentityType_AadApplication,
		"bot":                TeamworkApplicationIdentityType_Bot,
		"office365connector": TeamworkApplicationIdentityType_Office365Connector,
		"outgoingwebhook":    TeamworkApplicationIdentityType_OutgoingWebhook,
		"tenantbot":          TeamworkApplicationIdentityType_TenantBot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkApplicationIdentityType(input)
	return &out, nil
}
