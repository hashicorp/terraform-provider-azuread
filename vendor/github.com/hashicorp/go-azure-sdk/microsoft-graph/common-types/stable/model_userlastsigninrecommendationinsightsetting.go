package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewRecommendationInsightSetting = UserLastSignInRecommendationInsightSetting{}

type UserLastSignInRecommendationInsightSetting struct {
	// Optional. Indicates the time period of inactivity (with respect to the start date of the review instance) that
	// recommendations will be configured from. The recommendation will be to deny if the user is inactive during the
	// look-back duration. For reviews of groups and Microsoft Entra roles, any duration is accepted. For reviews of
	// applications, 30 days is the maximum duration. If not specified, the duration is 30 days.
	RecommendationLookBackDuration nullable.Type[string] `json:"recommendationLookBackDuration,omitempty"`

	// Indicates whether inactivity is calculated based on the user's inactivity in the tenant or in the application. The
	// possible values are tenant, application, unknownFutureValue. application is only relevant when the access review is a
	// review of an assignment to an application.
	SignInScope *UserSignInRecommendationScope `json:"signInScope,omitempty"`

	// Fields inherited from AccessReviewRecommendationInsightSetting

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserLastSignInRecommendationInsightSetting) AccessReviewRecommendationInsightSetting() BaseAccessReviewRecommendationInsightSettingImpl {
	return BaseAccessReviewRecommendationInsightSettingImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserLastSignInRecommendationInsightSetting{}

func (s UserLastSignInRecommendationInsightSetting) MarshalJSON() ([]byte, error) {
	type wrapper UserLastSignInRecommendationInsightSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserLastSignInRecommendationInsightSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserLastSignInRecommendationInsightSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userLastSignInRecommendationInsightSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserLastSignInRecommendationInsightSetting: %+v", err)
	}

	return encoded, nil
}
