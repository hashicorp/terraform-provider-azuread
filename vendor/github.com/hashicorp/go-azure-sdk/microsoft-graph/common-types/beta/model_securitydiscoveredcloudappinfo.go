package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityDiscoveredCloudAppInfo{}

type SecurityDiscoveredCloudAppInfo struct {
	CsaStarLevel               *SecurityAppInfoCsaStarLevel               `json:"csaStarLevel,omitempty"`
	DataAtRestEncryptionMethod *SecurityAppInfoDataAtRestEncryptionMethod `json:"dataAtRestEncryptionMethod,omitempty"`

	// Indicates the countries or regions in which your data center resides.
	DataCenter *string `json:"dataCenter,omitempty"`

	DataRetentionPolicy *SecurityAppInfoDataRetentionPolicy `json:"dataRetentionPolicy,omitempty"`

	// Indicates the data types that an end user can upload to the app. The possible values are: documents, mediaFiles,
	// codingFiles, creditCards, databaseFiles, none, unknown, unknownFutureValue.
	DataTypes *[]SecurityAppInfoUploadedDataTypes `json:"dataTypes,omitempty"`

	// Indicates the date when the app domain was registered.
	DomainRegistrationDateTime nullable.Type[string] `json:"domainRegistrationDateTime,omitempty"`

	EncryptionProtocol *SecurityAppInfoEncryptionProtocol `json:"encryptionProtocol,omitempty"`
	FedRampLevel       *SecurityAppInfoFedRampLevel       `json:"fedRampLevel,omitempty"`

	// Indicates the year that the specific app vendor was established.
	Founded *int64 `json:"founded,omitempty"`

	// Indicates the GDPR readiness of the app in relation to policies app provides to safeguard personal user data.
	GdprReadinessStatement *string `json:"gdprReadinessStatement,omitempty"`

	// Indicates the location of the headquarters of the app.
	Headquarters *string `json:"headquarters,omitempty"`

	Holding *SecurityAppInfoHolding `json:"holding,omitempty"`

	// Indicates the company name that provides hosting services for the app.
	HostingCompany *string `json:"hostingCompany,omitempty"`

	IsAdminAuditTrail                                 *SecurityCloudAppInfoState `json:"isAdminAuditTrail,omitempty"`
	IsCobitCompliant                                  *SecurityCloudAppInfoState `json:"isCobitCompliant,omitempty"`
	IsCoppaCompliant                                  *SecurityCloudAppInfoState `json:"isCoppaCompliant,omitempty"`
	IsDataAuditTrail                                  *SecurityCloudAppInfoState `json:"isDataAuditTrail,omitempty"`
	IsDataClassification                              *SecurityCloudAppInfoState `json:"isDataClassification,omitempty"`
	IsDataOwnership                                   *SecurityCloudAppInfoState `json:"isDataOwnership,omitempty"`
	IsDisasterRecoveryPlan                            *SecurityCloudAppInfoState `json:"isDisasterRecoveryPlan,omitempty"`
	IsDmca                                            *SecurityCloudAppInfoState `json:"isDmca,omitempty"`
	IsFerpaCompliant                                  *SecurityCloudAppInfoState `json:"isFerpaCompliant,omitempty"`
	IsFfiecCompliant                                  *SecurityCloudAppInfoState `json:"isFfiecCompliant,omitempty"`
	IsFileSharing                                     *SecurityCloudAppInfoState `json:"isFileSharing,omitempty"`
	IsFinraCompliant                                  *SecurityCloudAppInfoState `json:"isFinraCompliant,omitempty"`
	IsFismaCompliant                                  *SecurityCloudAppInfoState `json:"isFismaCompliant,omitempty"`
	IsGaapCompliant                                   *SecurityCloudAppInfoState `json:"isGaapCompliant,omitempty"`
	IsGdprDataProtectionImpactAssessment              *SecurityCloudAppInfoState `json:"isGdprDataProtectionImpactAssessment,omitempty"`
	IsGdprDataProtectionOfficer                       *SecurityCloudAppInfoState `json:"isGdprDataProtectionOfficer,omitempty"`
	IsGdprDataProtectionSecureCrossBorderDataTransfer *SecurityCloudAppInfoState `json:"isGdprDataProtectionSecureCrossBorderDataTransfer,omitempty"`
	IsGdprImpactAssessment                            *SecurityCloudAppInfoState `json:"isGdprImpactAssessment,omitempty"`
	IsGdprLawfulBasisForProcessing                    *SecurityCloudAppInfoState `json:"isGdprLawfulBasisForProcessing,omitempty"`
	IsGdprReportDataBreaches                          *SecurityCloudAppInfoState `json:"isGdprReportDataBreaches,omitempty"`
	IsGdprRightToAccess                               *SecurityCloudAppInfoState `json:"isGdprRightToAccess,omitempty"`
	IsGdprRightToBeInformed                           *SecurityCloudAppInfoState `json:"isGdprRightToBeInformed,omitempty"`
	IsGdprRightToDataPortablility                     *SecurityCloudAppInfoState `json:"isGdprRightToDataPortablility,omitempty"`
	IsGdprRightToErasure                              *SecurityCloudAppInfoState `json:"isGdprRightToErasure,omitempty"`
	IsGdprRightToObject                               *SecurityCloudAppInfoState `json:"isGdprRightToObject,omitempty"`
	IsGdprRightToRectification                        *SecurityCloudAppInfoState `json:"isGdprRightToRectification,omitempty"`
	IsGdprRightToRestrictionOfProcessing              *SecurityCloudAppInfoState `json:"isGdprRightToRestrictionOfProcessing,omitempty"`
	IsGdprRightsRelatedToAutomatedDecisionMaking      *SecurityCloudAppInfoState `json:"isGdprRightsRelatedToAutomatedDecisionMaking,omitempty"`
	IsGdprSecureCrossBorderDataControl                *SecurityCloudAppInfoState `json:"isGdprSecureCrossBorderDataControl,omitempty"`
	IsGlbaCompliant                                   *SecurityCloudAppInfoState `json:"isGlbaCompliant,omitempty"`
	IsHipaaCompliant                                  *SecurityCloudAppInfoState `json:"isHipaaCompliant,omitempty"`
	IsHitrustCsfCompliant                             *SecurityCloudAppInfoState `json:"isHitrustCsfCompliant,omitempty"`
	IsHttpSecurityHeadersContentSecurityPolicy        *SecurityCloudAppInfoState `json:"isHttpSecurityHeadersContentSecurityPolicy,omitempty"`
	IsHttpSecurityHeadersStrictTransportSecurity      *SecurityCloudAppInfoState `json:"isHttpSecurityHeadersStrictTransportSecurity,omitempty"`
	IsHttpSecurityHeadersXContentTypeOptions          *SecurityCloudAppInfoState `json:"isHttpSecurityHeadersXContentTypeOptions,omitempty"`
	IsHttpSecurityHeadersXFrameOptions                *SecurityCloudAppInfoState `json:"isHttpSecurityHeadersXFrameOptions,omitempty"`
	IsHttpSecurityHeadersXXssProtection               *SecurityCloudAppInfoState `json:"isHttpSecurityHeadersXXssProtection,omitempty"`
	IsIPAddressRestriction                            *SecurityCloudAppInfoState `json:"isIpAddressRestriction,omitempty"`
	IsIsae3402Compliant                               *SecurityCloudAppInfoState `json:"isIsae3402Compliant,omitempty"`
	IsIso27001Compliant                               *SecurityCloudAppInfoState `json:"isIso27001Compliant,omitempty"`
	IsIso27017Compliant                               *SecurityCloudAppInfoState `json:"isIso27017Compliant,omitempty"`
	IsIso27018Compliant                               *SecurityCloudAppInfoState `json:"isIso27018Compliant,omitempty"`
	IsItarCompliant                                   *SecurityCloudAppInfoState `json:"isItarCompliant,omitempty"`
	IsMultiFactorAuthentication                       *SecurityCloudAppInfoState `json:"isMultiFactorAuthentication,omitempty"`
	IsPasswordPolicy                                  *SecurityCloudAppInfoState `json:"isPasswordPolicy,omitempty"`
	IsPasswordPolicyChangePasswordPeriod              *SecurityCloudAppInfoState `json:"isPasswordPolicyChangePasswordPeriod,omitempty"`
	IsPasswordPolicyCharacterCombination              *SecurityCloudAppInfoState `json:"isPasswordPolicyCharacterCombination,omitempty"`
	IsPasswordPolicyPasswordHistoryAndReuse           *SecurityCloudAppInfoState `json:"isPasswordPolicyPasswordHistoryAndReuse,omitempty"`
	IsPasswordPolicyPasswordLengthLimit               *SecurityCloudAppInfoState `json:"isPasswordPolicyPasswordLengthLimit,omitempty"`
	IsPasswordPolicyPersonalInformationUse            *SecurityCloudAppInfoState `json:"isPasswordPolicyPersonalInformationUse,omitempty"`
	IsPenetrationTesting                              *SecurityCloudAppInfoState `json:"isPenetrationTesting,omitempty"`
	IsPrivacyShieldCompliant                          *SecurityCloudAppInfoState `json:"isPrivacyShieldCompliant,omitempty"`
	IsRememberPassword                                *SecurityCloudAppInfoState `json:"isRememberPassword,omitempty"`
	IsRequiresUserAuthentication                      *SecurityCloudAppInfoState `json:"isRequiresUserAuthentication,omitempty"`
	IsSoc1Compliant                                   *SecurityCloudAppInfoState `json:"isSoc1Compliant,omitempty"`
	IsSoc2Compliant                                   *SecurityCloudAppInfoState `json:"isSoc2Compliant,omitempty"`
	IsSoc3Compliant                                   *SecurityCloudAppInfoState `json:"isSoc3Compliant,omitempty"`
	IsSoxCompliant                                    *SecurityCloudAppInfoState `json:"isSoxCompliant,omitempty"`
	IsSp80053Compliant                                *SecurityCloudAppInfoState `json:"isSp80053Compliant,omitempty"`
	IsSsae16Compliant                                 *SecurityCloudAppInfoState `json:"isSsae16Compliant,omitempty"`
	IsSupportsSaml                                    *SecurityCloudAppInfoState `json:"isSupportsSaml,omitempty"`
	IsTrustedCertificate                              *SecurityCloudAppInfoState `json:"isTrustedCertificate,omitempty"`
	IsUserAuditTrail                                  *SecurityCloudAppInfoState `json:"isUserAuditTrail,omitempty"`
	IsUserCanUploadData                               *SecurityCloudAppInfoState `json:"isUserCanUploadData,omitempty"`
	IsUserRolesSupport                                *SecurityCloudAppInfoState `json:"isUserRolesSupport,omitempty"`
	IsValidCertificateName                            *SecurityCloudAppInfoState `json:"isValidCertificateName,omitempty"`

	// Indicates the last date of the data breach for the company.
	LatestBreachDateTime nullable.Type[string] `json:"latestBreachDateTime,omitempty"`

	// Indicates the URL that users can use to sign into the app.
	LogonUrls *[]string `json:"logonUrls,omitempty"`

	PciDssVersion *SecurityAppInfoPciDssVersion `json:"pciDssVersion,omitempty"`

	// Indicates the app vendor.
	Vendor nullable.Type[string] `json:"vendor,omitempty"`

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

func (s SecurityDiscoveredCloudAppInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityDiscoveredCloudAppInfo{}

func (s SecurityDiscoveredCloudAppInfo) MarshalJSON() ([]byte, error) {
	type wrapper SecurityDiscoveredCloudAppInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityDiscoveredCloudAppInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDiscoveredCloudAppInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.discoveredCloudAppInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityDiscoveredCloudAppInfo: %+v", err)
	}

	return encoded, nil
}
