package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadIndustry string

const (
	PayloadIndustry_Banking           PayloadIndustry = "banking"
	PayloadIndustry_BusinessServices  PayloadIndustry = "businessServices"
	PayloadIndustry_Construction      PayloadIndustry = "construction"
	PayloadIndustry_Consulting        PayloadIndustry = "consulting"
	PayloadIndustry_ConsumerServices  PayloadIndustry = "consumerServices"
	PayloadIndustry_CourierServices   PayloadIndustry = "courierServices"
	PayloadIndustry_Education         PayloadIndustry = "education"
	PayloadIndustry_Energy            PayloadIndustry = "energy"
	PayloadIndustry_FinancialServices PayloadIndustry = "financialServices"
	PayloadIndustry_Government        PayloadIndustry = "government"
	PayloadIndustry_Healthcare        PayloadIndustry = "healthcare"
	PayloadIndustry_Hospitality       PayloadIndustry = "hospitality"
	PayloadIndustry_IT                PayloadIndustry = "IT"
	PayloadIndustry_Insurance         PayloadIndustry = "insurance"
	PayloadIndustry_Legal             PayloadIndustry = "legal"
	PayloadIndustry_Manufacturing     PayloadIndustry = "manufacturing"
	PayloadIndustry_Other             PayloadIndustry = "other"
	PayloadIndustry_RealEstate        PayloadIndustry = "realEstate"
	PayloadIndustry_Retail            PayloadIndustry = "retail"
	PayloadIndustry_Telecom           PayloadIndustry = "telecom"
	PayloadIndustry_Unknown           PayloadIndustry = "unknown"
)

func PossibleValuesForPayloadIndustry() []string {
	return []string{
		string(PayloadIndustry_Banking),
		string(PayloadIndustry_BusinessServices),
		string(PayloadIndustry_Construction),
		string(PayloadIndustry_Consulting),
		string(PayloadIndustry_ConsumerServices),
		string(PayloadIndustry_CourierServices),
		string(PayloadIndustry_Education),
		string(PayloadIndustry_Energy),
		string(PayloadIndustry_FinancialServices),
		string(PayloadIndustry_Government),
		string(PayloadIndustry_Healthcare),
		string(PayloadIndustry_Hospitality),
		string(PayloadIndustry_IT),
		string(PayloadIndustry_Insurance),
		string(PayloadIndustry_Legal),
		string(PayloadIndustry_Manufacturing),
		string(PayloadIndustry_Other),
		string(PayloadIndustry_RealEstate),
		string(PayloadIndustry_Retail),
		string(PayloadIndustry_Telecom),
		string(PayloadIndustry_Unknown),
	}
}

func (s *PayloadIndustry) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadIndustry(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadIndustry(input string) (*PayloadIndustry, error) {
	vals := map[string]PayloadIndustry{
		"banking":           PayloadIndustry_Banking,
		"businessservices":  PayloadIndustry_BusinessServices,
		"construction":      PayloadIndustry_Construction,
		"consulting":        PayloadIndustry_Consulting,
		"consumerservices":  PayloadIndustry_ConsumerServices,
		"courierservices":   PayloadIndustry_CourierServices,
		"education":         PayloadIndustry_Education,
		"energy":            PayloadIndustry_Energy,
		"financialservices": PayloadIndustry_FinancialServices,
		"government":        PayloadIndustry_Government,
		"healthcare":        PayloadIndustry_Healthcare,
		"hospitality":       PayloadIndustry_Hospitality,
		"it":                PayloadIndustry_IT,
		"insurance":         PayloadIndustry_Insurance,
		"legal":             PayloadIndustry_Legal,
		"manufacturing":     PayloadIndustry_Manufacturing,
		"other":             PayloadIndustry_Other,
		"realestate":        PayloadIndustry_RealEstate,
		"retail":            PayloadIndustry_Retail,
		"telecom":           PayloadIndustry_Telecom,
		"unknown":           PayloadIndustry_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadIndustry(input)
	return &out, nil
}
