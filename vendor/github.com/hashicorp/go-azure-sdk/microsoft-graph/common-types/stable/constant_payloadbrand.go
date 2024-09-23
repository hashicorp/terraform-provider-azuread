package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadBrand string

const (
	PayloadBrand_Adobe           PayloadBrand = "adobe"
	PayloadBrand_AmericanExpress PayloadBrand = "americanExpress"
	PayloadBrand_CapitalOne      PayloadBrand = "capitalOne"
	PayloadBrand_Dhl             PayloadBrand = "dhl"
	PayloadBrand_DocuSign        PayloadBrand = "docuSign"
	PayloadBrand_Dropbox         PayloadBrand = "dropbox"
	PayloadBrand_Facebook        PayloadBrand = "facebook"
	PayloadBrand_FirstAmerican   PayloadBrand = "firstAmerican"
	PayloadBrand_Microsoft       PayloadBrand = "microsoft"
	PayloadBrand_Netflix         PayloadBrand = "netflix"
	PayloadBrand_Other           PayloadBrand = "other"
	PayloadBrand_Scotiabank      PayloadBrand = "scotiabank"
	PayloadBrand_SendGrid        PayloadBrand = "sendGrid"
	PayloadBrand_StewartTitle    PayloadBrand = "stewartTitle"
	PayloadBrand_SyrinxCloud     PayloadBrand = "syrinxCloud"
	PayloadBrand_Teams           PayloadBrand = "teams"
	PayloadBrand_Tesco           PayloadBrand = "tesco"
	PayloadBrand_Unknown         PayloadBrand = "unknown"
	PayloadBrand_WellsFargo      PayloadBrand = "wellsFargo"
	PayloadBrand_Zoom            PayloadBrand = "zoom"
)

func PossibleValuesForPayloadBrand() []string {
	return []string{
		string(PayloadBrand_Adobe),
		string(PayloadBrand_AmericanExpress),
		string(PayloadBrand_CapitalOne),
		string(PayloadBrand_Dhl),
		string(PayloadBrand_DocuSign),
		string(PayloadBrand_Dropbox),
		string(PayloadBrand_Facebook),
		string(PayloadBrand_FirstAmerican),
		string(PayloadBrand_Microsoft),
		string(PayloadBrand_Netflix),
		string(PayloadBrand_Other),
		string(PayloadBrand_Scotiabank),
		string(PayloadBrand_SendGrid),
		string(PayloadBrand_StewartTitle),
		string(PayloadBrand_SyrinxCloud),
		string(PayloadBrand_Teams),
		string(PayloadBrand_Tesco),
		string(PayloadBrand_Unknown),
		string(PayloadBrand_WellsFargo),
		string(PayloadBrand_Zoom),
	}
}

func (s *PayloadBrand) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadBrand(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadBrand(input string) (*PayloadBrand, error) {
	vals := map[string]PayloadBrand{
		"adobe":           PayloadBrand_Adobe,
		"americanexpress": PayloadBrand_AmericanExpress,
		"capitalone":      PayloadBrand_CapitalOne,
		"dhl":             PayloadBrand_Dhl,
		"docusign":        PayloadBrand_DocuSign,
		"dropbox":         PayloadBrand_Dropbox,
		"facebook":        PayloadBrand_Facebook,
		"firstamerican":   PayloadBrand_FirstAmerican,
		"microsoft":       PayloadBrand_Microsoft,
		"netflix":         PayloadBrand_Netflix,
		"other":           PayloadBrand_Other,
		"scotiabank":      PayloadBrand_Scotiabank,
		"sendgrid":        PayloadBrand_SendGrid,
		"stewarttitle":    PayloadBrand_StewartTitle,
		"syrinxcloud":     PayloadBrand_SyrinxCloud,
		"teams":           PayloadBrand_Teams,
		"tesco":           PayloadBrand_Tesco,
		"unknown":         PayloadBrand_Unknown,
		"wellsfargo":      PayloadBrand_WellsFargo,
		"zoom":            PayloadBrand_Zoom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadBrand(input)
	return &out, nil
}
