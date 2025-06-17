package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBase interface {
	Entity
	RecommendationBase() BaseRecommendationBaseImpl
}

var _ RecommendationBase = BaseRecommendationBaseImpl{}

type BaseRecommendationBaseImpl struct {
	// List of actions to take to complete a recommendation.
	ActionSteps *[]ActionStep `json:"actionSteps,omitempty"`

	// An explanation of why completing the recommendation will benefit you. Corresponds to the Value section of a
	// recommendation shown in the Microsoft Entra admin center.
	Benefits nullable.Type[string] `json:"benefits,omitempty"`

	Category *RecommendationCategory `json:"category,omitempty"`

	// The date and time when the recommendation was detected as applicable to your directory.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The title of the recommendation.
	DisplayName *string `json:"displayName,omitempty"`

	// The directory feature that the recommendation is related to.
	FeatureAreas *[]RecommendationFeatureAreas `json:"featureAreas,omitempty"`

	// The future date and time when a recommendation should be completed.
	ImpactStartDateTime nullable.Type[string] `json:"impactStartDateTime,omitempty"`

	// Indicates the scope of impact of a recommendation. tenantLevel indicates that the recommendation impacts the whole
	// tenant. Other possible values include users, apps.
	ImpactType nullable.Type[string] `json:"impactType,omitempty"`

	// The list of directory objects associated with the recommendation.
	ImpactedResources *[]ImpactedResource `json:"impactedResources,omitempty"`

	// Describes why a recommendation uniquely applies to your directory. Corresponds to the Description section of a
	// recommendation shown in the Microsoft Entra admin center.
	Insights nullable.Type[string] `json:"insights,omitempty"`

	// The most recent date and time a recommendation was deemed applicable to your directory.
	LastCheckedDateTime nullable.Type[string] `json:"lastCheckedDateTime,omitempty"`

	// Name of the user who last updated the status of the recommendation.
	LastModifiedBy nullable.Type[string] `json:"lastModifiedBy,omitempty"`

	// The date and time the status of a recommendation was last updated.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The future date and time when the status of a postponed recommendation will be active again.
	PostponeUntilDateTime nullable.Type[string] `json:"postponeUntilDateTime,omitempty"`

	Priority *RecommendationPriority `json:"priority,omitempty"`

	// Friendly shortname to identify the recommendation. The possible values are: adfsAppsMigration, enableDesktopSSO,
	// enablePHS, enableProvisioning, switchFromPerUserMFA, tenantMFA, thirdPartyApps, turnOffPerUserMFA,
	// useAuthenticatorApp, useMyApps, staleApps, staleAppCreds, applicationCredentialExpiry, servicePrincipalKeyExpiry,
	// adminMFAV2, blockLegacyAuthentication, integratedApps, mfaRegistrationV2, pwagePolicyNew, passwordHashSync, oneAdmin,
	// roleOverlap, selfServicePasswordReset, signinRiskPolicy, userRiskPolicy, verifyAppPublisher, privateLinkForAAD,
	// appRoleAssignmentsGroups, appRoleAssignmentsUsers, managedIdentity, overprivilegedApps, unknownFutureValue,
	// longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests,
	// aadGraphDeprecationApplication, aadGraphDeprecationServicePrincipal, mfaServerDeprecation. Use the Prefer:
	// include-unknown-enum-members request header to get the following value(s) in this evolvable enum:
	// longLivedCredentials , aadConnectDeprecated , adalToMsalMigration , ownerlessApps , inactiveGuests ,
	// aadGraphDeprecationApplication , aadGraphDeprecationServicePrincipal , mfaServerDeprecation.
	RecommendationType *RecommendationType `json:"recommendationType,omitempty"`

	// The current release type of the recommendation. The possible values are: preview, generallyAvailable,
	// unknownFutureValue.
	ReleaseType nullable.Type[string] `json:"releaseType,omitempty"`

	// Description of the impact on users of the remediation. Only applies to recommendations with category set to
	// identitySecureScore.
	RemediationImpact nullable.Type[string] `json:"remediationImpact,omitempty"`

	// The required licenses to view the recommendation. The possible values are: notApplicable, microsoftEntraIdFree,
	// microsoftEntraIdP1, microsoftEntraIdP2, microsoftEntraIdGovernance, microsoftEntraWorkloadId, unknownFutureValue.
	RequiredLicenses *RequiredLicenses `json:"requiredLicenses,omitempty"`

	Status *RecommendationStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseRecommendationBaseImpl) RecommendationBase() BaseRecommendationBaseImpl {
	return s
}

func (s BaseRecommendationBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RecommendationBase = RawRecommendationBaseImpl{}

// RawRecommendationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRecommendationBaseImpl struct {
	recommendationBase BaseRecommendationBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawRecommendationBaseImpl) RecommendationBase() BaseRecommendationBaseImpl {
	return s.recommendationBase
}

func (s RawRecommendationBaseImpl) Entity() BaseEntityImpl {
	return s.recommendationBase.Entity()
}

var _ json.Marshaler = BaseRecommendationBaseImpl{}

func (s BaseRecommendationBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRecommendationBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRecommendationBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRecommendationBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.recommendationBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRecommendationBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalRecommendationBaseImplementation(input []byte) (RecommendationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecommendationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.recommendation") {
		var out Recommendation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Recommendation: %+v", err)
		}
		return out, nil
	}

	var parent BaseRecommendationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRecommendationBaseImpl: %+v", err)
	}

	return RawRecommendationBaseImpl{
		recommendationBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
