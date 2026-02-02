package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseCollectionPaginationCountResponse interface {
	BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl
}

var _ BaseCollectionPaginationCountResponse = BaseBaseCollectionPaginationCountResponseImpl{}

type BaseBaseCollectionPaginationCountResponseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseBaseCollectionPaginationCountResponseImpl) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return s
}

var _ BaseCollectionPaginationCountResponse = RawBaseCollectionPaginationCountResponseImpl{}

// RawBaseCollectionPaginationCountResponseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBaseCollectionPaginationCountResponseImpl struct {
	baseCollectionPaginationCountResponse BaseBaseCollectionPaginationCountResponseImpl
	Type                                  string
	Values                                map[string]interface{}
}

func (s RawBaseCollectionPaginationCountResponseImpl) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return s.baseCollectionPaginationCountResponse
}

func UnmarshalBaseCollectionPaginationCountResponseImplementation(input []byte) (BaseCollectionPaginationCountResponse, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCollectionPaginationCountResponse into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.aadUserConversationMemberCollectionResponse") {
		var out AadUserConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadUserConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAnswerChoiceCollectionResponse") {
		var out AccessPackageAnswerChoiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAnswerChoiceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAnswerCollectionResponse") {
		var out AccessPackageAnswerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAnswerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageApprovalStageCollectionResponse") {
		var out AccessPackageApprovalStageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageApprovalStageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentCollectionResponse") {
		var out AccessPackageAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentPolicyCollectionResponse") {
		var out AccessPackageAssignmentPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentRequestCollectionResponse") {
		var out AccessPackageAssignmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentRequestWorkflowExtensionCollectionResponse") {
		var out AccessPackageAssignmentRequestWorkflowExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentRequestWorkflowExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentWorkflowExtensionCollectionResponse") {
		var out AccessPackageAssignmentWorkflowExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentWorkflowExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageCatalogCollectionResponse") {
		var out AccessPackageCatalogCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageCatalogCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageCollectionResponse") {
		var out AccessPackageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageLocalizedTextCollectionResponse") {
		var out AccessPackageLocalizedTextCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageLocalizedTextCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageMultipleChoiceQuestionCollectionResponse") {
		var out AccessPackageMultipleChoiceQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageMultipleChoiceQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageQuestionCollectionResponse") {
		var out AccessPackageQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceAttributeCollectionResponse") {
		var out AccessPackageResourceAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceCollectionResponse") {
		var out AccessPackageResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceEnvironmentCollectionResponse") {
		var out AccessPackageResourceEnvironmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceEnvironmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceRequestCollectionResponse") {
		var out AccessPackageResourceRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceRoleCollectionResponse") {
		var out AccessPackageResourceRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceRoleScopeCollectionResponse") {
		var out AccessPackageResourceRoleScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceRoleScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceScopeCollectionResponse") {
		var out AccessPackageResourceScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageTextInputQuestionCollectionResponse") {
		var out AccessPackageTextInputQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageTextInputQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewApplyActionCollectionResponse") {
		var out AccessReviewApplyActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewApplyActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewHistoryDefinitionCollectionResponse") {
		var out AccessReviewHistoryDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewHistoryDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewHistoryInstanceCollectionResponse") {
		var out AccessReviewHistoryInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewHistoryInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceCollectionResponse") {
		var out AccessReviewInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItemCollectionResponse") {
		var out AccessReviewInstanceDecisionItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewNotificationRecipientItemCollectionResponse") {
		var out AccessReviewNotificationRecipientItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewNotificationRecipientItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewRecommendationInsightSettingCollectionResponse") {
		var out AccessReviewRecommendationInsightSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewRecommendationInsightSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewReviewerCollectionResponse") {
		var out AccessReviewReviewerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewReviewerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewReviewerScopeCollectionResponse") {
		var out AccessReviewReviewerScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewReviewerScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewScheduleDefinitionCollectionResponse") {
		var out AccessReviewScheduleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewScheduleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewScopeCollectionResponse") {
		var out AccessReviewScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewStageCollectionResponse") {
		var out AccessReviewStageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewStageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewStageSettingsCollectionResponse") {
		var out AccessReviewStageSettingsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewStageSettingsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.activityBasedTimeoutPolicyCollectionResponse") {
		var out ActivityBasedTimeoutPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityBasedTimeoutPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.activityHistoryItemCollectionResponse") {
		var out ActivityHistoryItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityHistoryItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.addInCollectionResponse") {
		var out AddInCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddInCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.addLargeGalleryViewOperationCollectionResponse") {
		var out AddLargeGalleryViewOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddLargeGalleryViewOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.administrativeUnitCollectionResponse") {
		var out AdministrativeUnitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AdministrativeUnitCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aggregationOptionCollectionResponse") {
		var out AggregationOptionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AggregationOptionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementAcceptanceCollectionResponse") {
		var out AgreementAcceptanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementAcceptanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementCollectionResponse") {
		var out AgreementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementFileLocalizationCollectionResponse") {
		var out AgreementFileLocalizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementFileLocalizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementFileVersionCollectionResponse") {
		var out AgreementFileVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementFileVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionAttachmentCollectionResponse") {
		var out AiInteractionAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionContextCollectionResponse") {
		var out AiInteractionContextCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionContextCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionLinkCollectionResponse") {
		var out AiInteractionLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionMentionCollectionResponse") {
		var out AiInteractionMentionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionMentionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiUserCollectionResponse") {
		var out AiUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alertCollectionResponse") {
		var out AlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alertDetectionCollectionResponse") {
		var out AlertDetectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AlertDetectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alertHistoryStateCollectionResponse") {
		var out AlertHistoryStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AlertHistoryStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alertTriggerCollectionResponse") {
		var out AlertTriggerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AlertTriggerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.allowedValueCollectionResponse") {
		var out AllowedValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllowedValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alteredQueryTokenCollectionResponse") {
		var out AlteredQueryTokenCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AlteredQueryTokenCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alternativeSecurityIdCollectionResponse") {
		var out AlternativeSecurityIdCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AlternativeSecurityIdCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidCompliancePolicyCollectionResponse") {
		var out AndroidCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidCustomConfigurationCollectionResponse") {
		var out AndroidCustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidCustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidGeneralDeviceConfigurationCollectionResponse") {
		var out AndroidGeneralDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidGeneralDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidLobAppCollectionResponse") {
		var out AndroidLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedAppProtectionCollectionResponse") {
		var out AndroidManagedAppProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedAppProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedAppRegistrationCollectionResponse") {
		var out AndroidManagedAppRegistrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedAppRegistrationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidStoreAppCollectionResponse") {
		var out AndroidStoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidStoreAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileCompliancePolicyCollectionResponse") {
		var out AndroidWorkProfileCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileCustomConfigurationCollectionResponse") {
		var out AndroidWorkProfileCustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileCustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileGeneralDeviceConfigurationCollectionResponse") {
		var out AndroidWorkProfileGeneralDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileGeneralDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.anonymousGuestConversationMemberCollectionResponse") {
		var out AnonymousGuestConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AnonymousGuestConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appConfigurationSettingItemCollectionResponse") {
		var out AppConfigurationSettingItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppConfigurationSettingItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appConsentRequestCollectionResponse") {
		var out AppConsentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppConsentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appConsentRequestScopeCollectionResponse") {
		var out AppConsentRequestScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppConsentRequestScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appListItemCollectionResponse") {
		var out AppListItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppListItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appLogCollectionRequestCollectionResponse") {
		var out AppLogCollectionRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppLogCollectionRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appManagementPolicyCollectionResponse") {
		var out AppManagementPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppManagementPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appRoleAssignmentCollectionResponse") {
		var out AppRoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppRoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appRoleCollectionResponse") {
		var out AppRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleDeviceFeaturesConfigurationBaseCollectionResponse") {
		var out AppleDeviceFeaturesConfigurationBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleDeviceFeaturesConfigurationBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleManagedIdentityProviderCollectionResponse") {
		var out AppleManagedIdentityProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleManagedIdentityProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.applicationCollectionResponse") {
		var out ApplicationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.applicationTemplateCollectionResponse") {
		var out ApplicationTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appliedConditionalAccessPolicyCollectionResponse") {
		var out AppliedConditionalAccessPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppliedConditionalAccessPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalCollectionResponse") {
		var out ApprovalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalStageCollectionResponse") {
		var out ApprovalStageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalStageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignedLabelCollectionResponse") {
		var out AssignedLabelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignedLabelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignedLicenseCollectionResponse") {
		var out AssignedLicenseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignedLicenseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignedPlanCollectionResponse") {
		var out AssignedPlanCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignedPlanCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignedTrainingInfoCollectionResponse") {
		var out AssignedTrainingInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignedTrainingInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.associatedTeamInfoCollectionResponse") {
		var out AssociatedTeamInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssociatedTeamInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attachmentBaseCollectionResponse") {
		var out AttachmentBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttachmentBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attachmentCollectionResponse") {
		var out AttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attachmentSessionCollectionResponse") {
		var out AttachmentSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttachmentSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attackSimulationOperationCollectionResponse") {
		var out AttackSimulationOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttackSimulationOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attendanceIntervalCollectionResponse") {
		var out AttendanceIntervalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttendanceIntervalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attendanceRecordCollectionResponse") {
		var out AttendanceRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttendanceRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attendeeAvailabilityCollectionResponse") {
		var out AttendeeAvailabilityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttendeeAvailabilityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attendeeCollectionResponse") {
		var out AttendeeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttendeeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeDefinitionCollectionResponse") {
		var out AttributeDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeDefinitionMetadataEntryCollectionResponse") {
		var out AttributeDefinitionMetadataEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeDefinitionMetadataEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeMappingCollectionResponse") {
		var out AttributeMappingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeMappingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeMappingFunctionSchemaCollectionResponse") {
		var out AttributeMappingFunctionSchemaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeMappingFunctionSchemaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeMappingParameterSchemaCollectionResponse") {
		var out AttributeMappingParameterSchemaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeMappingParameterSchemaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeSetCollectionResponse") {
		var out AttributeSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.audioRoutingGroupCollectionResponse") {
		var out AudioRoutingGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AudioRoutingGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.auditEventCollectionResponse") {
		var out AuditEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuditEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.auditPropertyCollectionResponse") {
		var out AuditPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuditPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.auditResourceCollectionResponse") {
		var out AuditResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuditResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationAttributeCollectionInputConfigurationCollectionResponse") {
		var out AuthenticationAttributeCollectionInputConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationAttributeCollectionInputConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationAttributeCollectionOptionConfigurationCollectionResponse") {
		var out AuthenticationAttributeCollectionOptionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationAttributeCollectionOptionConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationAttributeCollectionPageViewConfigurationCollectionResponse") {
		var out AuthenticationAttributeCollectionPageViewConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationAttributeCollectionPageViewConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationCombinationConfigurationCollectionResponse") {
		var out AuthenticationCombinationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationCombinationConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationConditionApplicationCollectionResponse") {
		var out AuthenticationConditionApplicationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationConditionApplicationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationContextClassReferenceCollectionResponse") {
		var out AuthenticationContextClassReferenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationContextClassReferenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationEventListenerCollectionResponse") {
		var out AuthenticationEventListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationEventListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationEventsFlowCollectionResponse") {
		var out AuthenticationEventsFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationEventsFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodCollectionResponse") {
		var out AuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodConfigurationCollectionResponse") {
		var out AuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodModeDetailCollectionResponse") {
		var out AuthenticationMethodModeDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodModeDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodTargetCollectionResponse") {
		var out AuthenticationMethodTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodsRegistrationCampaignIncludeTargetCollectionResponse") {
		var out AuthenticationMethodsRegistrationCampaignIncludeTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodsRegistrationCampaignIncludeTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationStrengthPolicyCollectionResponse") {
		var out AuthenticationStrengthPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationStrengthPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authoredNoteCollectionResponse") {
		var out AuthoredNoteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthoredNoteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authorizationPolicyCollectionResponse") {
		var out AuthorizationPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthorizationPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.availabilityItemCollectionResponse") {
		var out AvailabilityItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AvailabilityItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.averageComparativeScoreCollectionResponse") {
		var out AverageComparativeScoreCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AverageComparativeScoreCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureCommunicationServicesUserConversationMemberCollectionResponse") {
		var out AzureCommunicationServicesUserConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureCommunicationServicesUserConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.b2xIdentityUserFlowCollectionResponse") {
		var out B2xIdentityUserFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into B2xIdentityUserFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.baseItemCollectionResponse") {
		var out BaseItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BaseItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.baseSitePageCollectionResponse") {
		var out BaseSitePageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BaseSitePageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bitlockerRecoveryKeyCollectionResponse") {
		var out BitlockerRecoveryKeyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BitlockerRecoveryKeyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingAppointmentCollectionResponse") {
		var out BookingAppointmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingAppointmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingBusinessCollectionResponse") {
		var out BookingBusinessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingBusinessCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCurrencyCollectionResponse") {
		var out BookingCurrencyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCurrencyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomQuestionCollectionResponse") {
		var out BookingCustomQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomerBaseCollectionResponse") {
		var out BookingCustomerBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomerBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomerCollectionResponse") {
		var out BookingCustomerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomerInformationBaseCollectionResponse") {
		var out BookingCustomerInformationBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomerInformationBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingQuestionAnswerCollectionResponse") {
		var out BookingQuestionAnswerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingQuestionAnswerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingQuestionAssignmentCollectionResponse") {
		var out BookingQuestionAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingQuestionAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingReminderCollectionResponse") {
		var out BookingReminderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingReminderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingServiceCollectionResponse") {
		var out BookingServiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingServiceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingStaffMemberBaseCollectionResponse") {
		var out BookingStaffMemberBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingStaffMemberBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingStaffMemberCollectionResponse") {
		var out BookingStaffMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingStaffMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingWorkHoursCollectionResponse") {
		var out BookingWorkHoursCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingWorkHoursCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingWorkTimeSlotCollectionResponse") {
		var out BookingWorkTimeSlotCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingWorkTimeSlotCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingsAvailabilityWindowCollectionResponse") {
		var out BookingsAvailabilityWindowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingsAvailabilityWindowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSharedCookieCollectionResponse") {
		var out BrowserSharedCookieCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSharedCookieCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSharedCookieHistoryCollectionResponse") {
		var out BrowserSharedCookieHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSharedCookieHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSiteCollectionResponse") {
		var out BrowserSiteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSiteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSiteHistoryCollectionResponse") {
		var out BrowserSiteHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSiteHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSiteListCollectionResponse") {
		var out BrowserSiteListCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSiteListCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bucketAggregationRangeCollectionResponse") {
		var out BucketAggregationRangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BucketAggregationRangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.builtInIdentityProviderCollectionResponse") {
		var out BuiltInIdentityProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BuiltInIdentityProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarCollectionResponse") {
		var out CalendarCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarGroupCollectionResponse") {
		var out CalendarGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarPermissionCollectionResponse") {
		var out CalendarPermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarPermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarSharingMessageActionCollectionResponse") {
		var out CalendarSharingMessageActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarSharingMessageActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarSharingMessageCollectionResponse") {
		var out CalendarSharingMessageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarSharingMessageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callCollectionResponse") {
		var out CallCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callParticipantInfoCollectionResponse") {
		var out CallParticipantInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallParticipantInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecordingCollectionResponse") {
		var out CallRecordingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.administrativeUnitInfoCollectionResponse") {
		var out CallRecordsAdministrativeUnitInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsAdministrativeUnitInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.callRecordCollectionResponse") {
		var out CallRecordsCallRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsCallRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.mediaCollectionResponse") {
		var out CallRecordsMediaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsMediaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.mediaStreamCollectionResponse") {
		var out CallRecordsMediaStreamCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsMediaStreamCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.participantCollectionResponse") {
		var out CallRecordsParticipantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsParticipantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.segmentCollectionResponse") {
		var out CallRecordsSegmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsSegmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.sessionCollectionResponse") {
		var out CallRecordsSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.traceRouteHopCollectionResponse") {
		var out CallRecordsTraceRouteHopCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsTraceRouteHopCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRouteCollectionResponse") {
		var out CallRouteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRouteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callTranscriptCollectionResponse") {
		var out CallTranscriptCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallTranscriptCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cancelMediaProcessingOperationCollectionResponse") {
		var out CancelMediaProcessingOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CancelMediaProcessingOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateAuthorityCollectionResponse") {
		var out CertificateAuthorityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateAuthorityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateAuthorityDetailCollectionResponse") {
		var out CertificateAuthorityDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateAuthorityDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateBasedAuthConfigurationCollectionResponse") {
		var out CertificateBasedAuthConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateBasedAuthConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateBasedAuthPkiCollectionResponse") {
		var out CertificateBasedAuthPkiCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateBasedAuthPkiCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificationControlCollectionResponse") {
		var out CertificationControlCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificationControlCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.changeNotificationCollectionResponse") {
		var out ChangeNotificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChangeNotificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelCollectionResponse") {
		var out ChannelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatCollectionResponse") {
		var out ChatCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageAttachmentCollectionResponse") {
		var out ChatMessageAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageCollectionResponse") {
		var out ChatMessageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageHistoryItemCollectionResponse") {
		var out ChatMessageHistoryItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageHistoryItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageHostedContentCollectionResponse") {
		var out ChatMessageHostedContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageHostedContentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageMentionCollectionResponse") {
		var out ChatMessageMentionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageMentionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageReactionCollectionResponse") {
		var out ChatMessageReactionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageReactionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.checklistItemCollectionResponse") {
		var out ChecklistItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChecklistItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.claimsMappingPolicyCollectionResponse") {
		var out ClaimsMappingPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClaimsMappingPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudAppSecurityStateCollectionResponse") {
		var out CloudAppSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudAppSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudClipboardItemCollectionResponse") {
		var out CloudClipboardItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudClipboardItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudClipboardItemPayloadCollectionResponse") {
		var out CloudClipboardItemPayloadCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudClipboardItemPayloadCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcAuditEventCollectionResponse") {
		var out CloudPCAuditEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCAuditEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcAuditPropertyCollectionResponse") {
		var out CloudPCAuditPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCAuditPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcAuditResourceCollectionResponse") {
		var out CloudPCAuditResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCAuditResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPCCollectionResponse") {
		var out CloudPCCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcDeviceImageCollectionResponse") {
		var out CloudPCDeviceImageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCDeviceImageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcDomainJoinConfigurationCollectionResponse") {
		var out CloudPCDomainJoinConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCDomainJoinConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcGalleryImageCollectionResponse") {
		var out CloudPCGalleryImageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCGalleryImageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcOnPremisesConnectionCollectionResponse") {
		var out CloudPCOnPremisesConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCOnPremisesConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcOnPremisesConnectionHealthCheckCollectionResponse") {
		var out CloudPCOnPremisesConnectionHealthCheckCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCOnPremisesConnectionHealthCheckCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcProvisioningPolicyAssignmentCollectionResponse") {
		var out CloudPCProvisioningPolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCProvisioningPolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcProvisioningPolicyCollectionResponse") {
		var out CloudPCProvisioningPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCProvisioningPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcUserRoleScopeTagInfoCollectionResponse") {
		var out CloudPCUserRoleScopeTagInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCUserRoleScopeTagInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcUserSettingAssignmentCollectionResponse") {
		var out CloudPCUserSettingAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCUserSettingAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcUserSettingCollectionResponse") {
		var out CloudPCUserSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCUserSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.collapsePropertyCollectionResponse") {
		var out CollapsePropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CollapsePropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.columnDefinitionCollectionResponse") {
		var out ColumnDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ColumnDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.columnLinkCollectionResponse") {
		var out ColumnLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ColumnLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.commsNotificationCollectionResponse") {
		var out CommsNotificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommsNotificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.commsOperationCollectionResponse") {
		var out CommsOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommsOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsUserIdentityCollectionResponse") {
		var out CommunicationsUserIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsUserIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communityCollectionResponse") {
		var out CommunityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.companySubscriptionCollectionResponse") {
		var out CompanySubscriptionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CompanySubscriptionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.complianceInformationCollectionResponse") {
		var out ComplianceInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ComplianceInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.complianceManagementPartnerAssignmentCollectionResponse") {
		var out ComplianceManagementPartnerAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ComplianceManagementPartnerAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.complianceManagementPartnerCollectionResponse") {
		var out ComplianceManagementPartnerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ComplianceManagementPartnerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessPolicyCollectionResponse") {
		var out ConditionalAccessPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessTemplateCollectionResponse") {
		var out ConditionalAccessTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.connectedOrganizationCollectionResponse") {
		var out ConnectedOrganizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectedOrganizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contactCollectionResponse") {
		var out ContactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contactFolderCollectionResponse") {
		var out ContactFolderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContactFolderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentSharingSessionCollectionResponse") {
		var out ContentSharingSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentSharingSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentTypeCollectionResponse") {
		var out ContentTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentTypeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentTypeInfoCollectionResponse") {
		var out ContentTypeInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentTypeInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contractCollectionResponse") {
		var out ContractCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContractCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.controlScoreCollectionResponse") {
		var out ControlScoreCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ControlScoreCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversationCollectionResponse") {
		var out ConversationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConversationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversationMemberCollectionResponse") {
		var out ConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversationThreadCollectionResponse") {
		var out ConversationThreadCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConversationThreadCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.countryNamedLocationCollectionResponse") {
		var out CountryNamedLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CountryNamedLocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicyCollectionResponse") {
		var out CrossTenantAccessPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicyConfigurationPartnerCollectionResponse") {
		var out CrossTenantAccessPolicyConfigurationPartnerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicyConfigurationPartnerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicyTargetCollectionResponse") {
		var out CrossTenantAccessPolicyTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicyTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customAuthenticationExtensionCollectionResponse") {
		var out CustomAuthenticationExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAuthenticationExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customCalloutExtensionCollectionResponse") {
		var out CustomCalloutExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomCalloutExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customExtensionCalloutInstanceCollectionResponse") {
		var out CustomExtensionCalloutInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomExtensionCalloutInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customExtensionStageSettingCollectionResponse") {
		var out CustomExtensionStageSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomExtensionStageSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customSecurityAttributeDefinitionCollectionResponse") {
		var out CustomSecurityAttributeDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSecurityAttributeDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dataPolicyOperationCollectionResponse") {
		var out DataPolicyOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataPolicyOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dayNoteCollectionResponse") {
		var out DayNoteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DayNoteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.defaultManagedAppProtectionCollectionResponse") {
		var out DefaultManagedAppProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefaultManagedAppProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminAccessAssignmentCollectionResponse") {
		var out DelegatedAdminAccessAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminAccessAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminCustomerCollectionResponse") {
		var out DelegatedAdminCustomerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminCustomerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminRelationshipCollectionResponse") {
		var out DelegatedAdminRelationshipCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminRelationshipCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminRelationshipOperationCollectionResponse") {
		var out DelegatedAdminRelationshipOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminRelationshipOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminRelationshipRequestCollectionResponse") {
		var out DelegatedAdminRelationshipRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminRelationshipRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminServiceManagementDetailCollectionResponse") {
		var out DelegatedAdminServiceManagementDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminServiceManagementDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedPermissionClassificationCollectionResponse") {
		var out DelegatedPermissionClassificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedPermissionClassificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deletedChatCollectionResponse") {
		var out DeletedChatCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeletedChatCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deletedTeamCollectionResponse") {
		var out DeletedTeamCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeletedTeamCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.detectedAppCollectionResponse") {
		var out DetectedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DetectedAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceActionResultCollectionResponse") {
		var out DeviceActionResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceActionResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementRoleAssignmentCollectionResponse") {
		var out DeviceAndAppManagementRoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementRoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementRoleDefinitionCollectionResponse") {
		var out DeviceAndAppManagementRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCategoryCollectionResponse") {
		var out DeviceCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCollectionResponse") {
		var out DeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceActionItemCollectionResponse") {
		var out DeviceComplianceActionItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceActionItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceDeviceStatusCollectionResponse") {
		var out DeviceComplianceDeviceStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceDeviceStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyAssignmentCollectionResponse") {
		var out DeviceCompliancePolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyCollectionResponse") {
		var out DeviceCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicySettingStateCollectionResponse") {
		var out DeviceCompliancePolicySettingStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicySettingStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicySettingStateSummaryCollectionResponse") {
		var out DeviceCompliancePolicySettingStateSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicySettingStateSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyStateCollectionResponse") {
		var out DeviceCompliancePolicyStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScheduledActionForRuleCollectionResponse") {
		var out DeviceComplianceScheduledActionForRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScheduledActionForRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceSettingStateCollectionResponse") {
		var out DeviceComplianceSettingStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceSettingStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceUserStatusCollectionResponse") {
		var out DeviceComplianceUserStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceUserStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationAssignmentCollectionResponse") {
		var out DeviceConfigurationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationCollectionResponse") {
		var out DeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationDeviceStatusCollectionResponse") {
		var out DeviceConfigurationDeviceStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationDeviceStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationSettingStateCollectionResponse") {
		var out DeviceConfigurationSettingStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationSettingStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationStateCollectionResponse") {
		var out DeviceConfigurationStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationUserStatusCollectionResponse") {
		var out DeviceConfigurationUserStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationUserStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentConfigurationCollectionResponse") {
		var out DeviceEnrollmentConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentLimitConfigurationCollectionResponse") {
		var out DeviceEnrollmentLimitConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentLimitConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse") {
		var out DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentWindowsHelloForBusinessConfigurationCollectionResponse") {
		var out DeviceEnrollmentWindowsHelloForBusinessConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentWindowsHelloForBusinessConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceInstallStateCollectionResponse") {
		var out DeviceInstallStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceInstallStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceLocalCredentialCollectionResponse") {
		var out DeviceLocalCredentialCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceLocalCredentialCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceLocalCredentialInfoCollectionResponse") {
		var out DeviceLocalCredentialInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceLocalCredentialInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceLogCollectionResponseCollectionResponse") {
		var out DeviceLogCollectionResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceLogCollectionResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExchangeConnectorCollectionResponse") {
		var out DeviceManagementExchangeConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExchangeConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExportJobCollectionResponse") {
		var out DeviceManagementExportJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExportJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementPartnerAssignmentCollectionResponse") {
		var out DeviceManagementPartnerAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementPartnerAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementPartnerCollectionResponse") {
		var out DeviceManagementPartnerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementPartnerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTroubleshootingEventCollectionResponse") {
		var out DeviceManagementTroubleshootingEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTroubleshootingEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryAuditCollectionResponse") {
		var out DirectoryAuditCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryAuditCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryDefinitionCollectionResponse") {
		var out DirectoryDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryObjectCollectionResponse") {
		var out DirectoryObjectCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryObjectCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryObjectPartnerReferenceCollectionResponse") {
		var out DirectoryObjectPartnerReferenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryObjectPartnerReferenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryRoleCollectionResponse") {
		var out DirectoryRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryRoleTemplateCollectionResponse") {
		var out DirectoryRoleTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryRoleTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.displayNameLocalizationCollectionResponse") {
		var out DisplayNameLocalizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DisplayNameLocalizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.documentSetContentCollectionResponse") {
		var out DocumentSetContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentSetContentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.documentSetVersionCollectionResponse") {
		var out DocumentSetVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentSetVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.documentSetVersionItemCollectionResponse") {
		var out DocumentSetVersionItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentSetVersionItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainCollectionResponse") {
		var out DomainCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsCnameRecordCollectionResponse") {
		var out DomainDnsCnameRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsCnameRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsMxRecordCollectionResponse") {
		var out DomainDnsMxRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsMxRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsRecordCollectionResponse") {
		var out DomainDnsRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsSrvRecordCollectionResponse") {
		var out DomainDnsSrvRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsSrvRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsTxtRecordCollectionResponse") {
		var out DomainDnsTxtRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsTxtRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsUnavailableRecordCollectionResponse") {
		var out DomainDnsUnavailableRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsUnavailableRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveCollectionResponse") {
		var out DriveCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveItemCollectionResponse") {
		var out DriveItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveItemVersionCollectionResponse") {
		var out DriveItemVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveItemVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveProtectionRuleCollectionResponse") {
		var out DriveProtectionRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveProtectionRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveProtectionUnitCollectionResponse") {
		var out DriveProtectionUnitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveProtectionUnitCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveProtectionUnitsBulkAdditionJobCollectionResponse") {
		var out DriveProtectionUnitsBulkAdditionJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveProtectionUnitsBulkAdditionJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveRestoreArtifactCollectionResponse") {
		var out DriveRestoreArtifactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveRestoreArtifactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.driveRestoreArtifactsBulkAdditionRequestCollectionResponse") {
		var out DriveRestoreArtifactsBulkAdditionRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveRestoreArtifactsBulkAdditionRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.editionUpgradeConfigurationCollectionResponse") {
		var out EditionUpgradeConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EditionUpgradeConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentCollectionResponse") {
		var out EducationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentResourceCollectionResponse") {
		var out EducationAssignmentResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationCategoryCollectionResponse") {
		var out EducationCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationClassCollectionResponse") {
		var out EducationClassCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationClassCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationFeedbackOutcomeCollectionResponse") {
		var out EducationFeedbackOutcomeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationFeedbackOutcomeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationFeedbackResourceOutcomeCollectionResponse") {
		var out EducationFeedbackResourceOutcomeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationFeedbackResourceOutcomeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationGradingCategoryCollectionResponse") {
		var out EducationGradingCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationGradingCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationModuleCollectionResponse") {
		var out EducationModuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationModuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationModuleResourceCollectionResponse") {
		var out EducationModuleResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationModuleResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationOutcomeCollectionResponse") {
		var out EducationOutcomeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationOutcomeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationPointsOutcomeCollectionResponse") {
		var out EducationPointsOutcomeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationPointsOutcomeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationRubricCollectionResponse") {
		var out EducationRubricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationRubricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationRubricOutcomeCollectionResponse") {
		var out EducationRubricOutcomeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationRubricOutcomeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSchoolCollectionResponse") {
		var out EducationSchoolCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSchoolCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSubmissionCollectionResponse") {
		var out EducationSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSubmissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSubmissionResourceCollectionResponse") {
		var out EducationSubmissionResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSubmissionResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationUserCollectionResponse") {
		var out EducationUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailAddressCollectionResponse") {
		var out EmailAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailAuthenticationMethodCollectionResponse") {
		var out EmailAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailAuthenticationMethodConfigurationCollectionResponse") {
		var out EmailAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailFileAssessmentRequestCollectionResponse") {
		var out EmailFileAssessmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailFileAssessmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endUserNotificationCollectionResponse") {
		var out EndUserNotificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndUserNotificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endUserNotificationDetailCollectionResponse") {
		var out EndUserNotificationDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndUserNotificationDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endpointCollectionResponse") {
		var out EndpointCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndpointCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.engagementAsyncOperationCollectionResponse") {
		var out EngagementAsyncOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EngagementAsyncOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentConfigurationAssignmentCollectionResponse") {
		var out EnrollmentConfigurationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentConfigurationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentTroubleshootingEventCollectionResponse") {
		var out EnrollmentTroubleshootingEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentTroubleshootingEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.eventCollectionResponse") {
		var out EventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.eventMessageCollectionResponse") {
		var out EventMessageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventMessageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.eventMessageRequestCollectionResponse") {
		var out EventMessageRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventMessageRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.eventMessageResponseCollectionResponse") {
		var out EventMessageResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventMessageResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exchangeProtectionPolicyCollectionResponse") {
		var out ExchangeProtectionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExchangeProtectionPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exchangeRestoreSessionCollectionResponse") {
		var out ExchangeRestoreSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExchangeRestoreSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.excludeTargetCollectionResponse") {
		var out ExcludeTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExcludeTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extensionCollectionResponse") {
		var out ExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extensionPropertyCollectionResponse") {
		var out ExtensionPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtensionPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extensionSchemaPropertyCollectionResponse") {
		var out ExtensionSchemaPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtensionSchemaPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.aclCollectionResponse") {
		var out ExternalConnectorsAclCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsAclCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.connectionOperationCollectionResponse") {
		var out ExternalConnectorsConnectionOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsConnectionOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.displayTemplateCollectionResponse") {
		var out ExternalConnectorsDisplayTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsDisplayTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalActivityCollectionResponse") {
		var out ExternalConnectorsExternalActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalActivityResultCollectionResponse") {
		var out ExternalConnectorsExternalActivityResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalActivityResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalConnectionCollectionResponse") {
		var out ExternalConnectorsExternalConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalGroupCollectionResponse") {
		var out ExternalConnectorsExternalGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalItemCollectionResponse") {
		var out ExternalConnectorsExternalItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.identityCollectionResponse") {
		var out ExternalConnectorsIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.propertyCollectionResponse") {
		var out ExternalConnectorsPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.propertyRuleCollectionResponse") {
		var out ExternalConnectorsPropertyRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsPropertyRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.urlToItemResolverBaseCollectionResponse") {
		var out ExternalConnectorsUrlToItemResolverBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsUrlToItemResolverBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalDomainNameCollectionResponse") {
		var out ExternalDomainNameCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalDomainNameCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalUsersSelfServiceSignUpEventsFlowCollectionResponse") {
		var out ExternalUsersSelfServiceSignUpEventsFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalUsersSelfServiceSignUpEventsFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.featureRolloutPolicyCollectionResponse") {
		var out FeatureRolloutPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureRolloutPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.federatedIdentityCredentialCollectionResponse") {
		var out FederatedIdentityCredentialCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FederatedIdentityCredentialCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fido2AuthenticationMethodCollectionResponse") {
		var out Fido2AuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Fido2AuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fido2AuthenticationMethodConfigurationCollectionResponse") {
		var out Fido2AuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Fido2AuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fido2CombinationConfigurationCollectionResponse") {
		var out Fido2CombinationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Fido2CombinationConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileAssessmentRequestCollectionResponse") {
		var out FileAssessmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileAssessmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileAttachmentCollectionResponse") {
		var out FileAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileSecurityStateCollectionResponse") {
		var out FileSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileStorageContainerCollectionResponse") {
		var out FileStorageContainerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileStorageContainerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.filterClauseCollectionResponse") {
		var out FilterClauseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FilterClauseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.filterGroupCollectionResponse") {
		var out FilterGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FilterGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.filterOperatorSchemaCollectionResponse") {
		var out FilterOperatorSchemaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FilterOperatorSchemaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.genericErrorCollectionResponse") {
		var out GenericErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GenericErrorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceInsightCollectionResponse") {
		var out GovernanceInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.granularMailboxRestoreArtifactCollectionResponse") {
		var out GranularMailboxRestoreArtifactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GranularMailboxRestoreArtifactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupCollectionResponse") {
		var out GroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupLifecyclePolicyCollectionResponse") {
		var out GroupLifecyclePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupLifecyclePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupSettingCollectionResponse") {
		var out GroupSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupSettingTemplateCollectionResponse") {
		var out GroupSettingTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupSettingTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.homeRealmDiscoveryPolicyCollectionResponse") {
		var out HomeRealmDiscoveryPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HomeRealmDiscoveryPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.horizontalSectionCollectionResponse") {
		var out HorizontalSectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HorizontalSectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.horizontalSectionColumnCollectionResponse") {
		var out HorizontalSectionColumnCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HorizontalSectionColumnCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hostSecurityStateCollectionResponse") {
		var out HostSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HostSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ipNamedLocationCollectionResponse") {
		var out IPNamedLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPNamedLocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ipRangeCollectionResponse") {
		var out IPRangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPRangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityApiConnectorCollectionResponse") {
		var out IdentityApiConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityApiConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityBuiltInUserFlowAttributeCollectionResponse") {
		var out IdentityBuiltInUserFlowAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityBuiltInUserFlowAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityCollectionResponse") {
		var out IdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityCustomUserFlowAttributeCollectionResponse") {
		var out IdentityCustomUserFlowAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityCustomUserFlowAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.customTaskExtensionCollectionResponse") {
		var out IdentityGovernanceCustomTaskExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceCustomTaskExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.parameterCollectionResponse") {
		var out IdentityGovernanceParameterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceParameterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.runCollectionResponse") {
		var out IdentityGovernanceRunCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceRunCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskCollectionResponse") {
		var out IdentityGovernanceTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskDefinitionCollectionResponse") {
		var out IdentityGovernanceTaskDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskProcessingResultCollectionResponse") {
		var out IdentityGovernanceTaskProcessingResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskProcessingResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskReportCollectionResponse") {
		var out IdentityGovernanceTaskReportCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskReportCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.triggerAttributeCollectionResponse") {
		var out IdentityGovernanceTriggerAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTriggerAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.userProcessingResultCollectionResponse") {
		var out IdentityGovernanceUserProcessingResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceUserProcessingResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.workflowCollectionResponse") {
		var out IdentityGovernanceWorkflowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceWorkflowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.workflowTemplateCollectionResponse") {
		var out IdentityGovernanceWorkflowTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceWorkflowTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.workflowVersionCollectionResponse") {
		var out IdentityGovernanceWorkflowVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceWorkflowVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityProviderBaseCollectionResponse") {
		var out IdentityProviderBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityProviderBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityProviderCollectionResponse") {
		var out IdentityProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identitySecurityDefaultsEnforcementPolicyCollectionResponse") {
		var out IdentitySecurityDefaultsEnforcementPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentitySecurityDefaultsEnforcementPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identitySetCollectionResponse") {
		var out IdentitySetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentitySetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identitySourceCollectionResponse") {
		var out IdentitySourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentitySourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityUserFlowAttributeAssignmentCollectionResponse") {
		var out IdentityUserFlowAttributeAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityUserFlowAttributeAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityUserFlowAttributeCollectionResponse") {
		var out IdentityUserFlowAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityUserFlowAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedWindowsAutopilotDeviceIdentityCollectionResponse") {
		var out ImportedWindowsAutopilotDeviceIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedWindowsAutopilotDeviceIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inferenceClassificationOverrideCollectionResponse") {
		var out InferenceClassificationOverrideCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InferenceClassificationOverrideCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.integerRangeCollectionResponse") {
		var out IntegerRangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IntegerRangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.internalDomainFederationCollectionResponse") {
		var out InternalDomainFederationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InternalDomainFederationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.internetMessageHeaderCollectionResponse") {
		var out InternetMessageHeaderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InternetMessageHeaderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.investigationSecurityStateCollectionResponse") {
		var out InvestigationSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvestigationSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.invitationCollectionResponse") {
		var out InvitationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvitationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.invitationParticipantInfoCollectionResponse") {
		var out InvitationParticipantInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvitationParticipantInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inviteParticipantsOperationCollectionResponse") {
		var out InviteParticipantsOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InviteParticipantsOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCertificateProfileCollectionResponse") {
		var out IosCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCompliancePolicyCollectionResponse") {
		var out IosCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCustomConfigurationCollectionResponse") {
		var out IosCustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosDeviceFeaturesConfigurationCollectionResponse") {
		var out IosDeviceFeaturesConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosDeviceFeaturesConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosGeneralDeviceConfigurationCollectionResponse") {
		var out IosGeneralDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosGeneralDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosHomeScreenAppCollectionResponse") {
		var out IosHomeScreenAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosHomeScreenAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosHomeScreenFolderPageCollectionResponse") {
		var out IosHomeScreenFolderPageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosHomeScreenFolderPageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosHomeScreenItemCollectionResponse") {
		var out IosHomeScreenItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosHomeScreenItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosHomeScreenPageCollectionResponse") {
		var out IosHomeScreenPageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosHomeScreenPageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppCollectionResponse") {
		var out IosLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosManagedAppProtectionCollectionResponse") {
		var out IosManagedAppProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosManagedAppProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosManagedAppRegistrationCollectionResponse") {
		var out IosManagedAppRegistrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosManagedAppRegistrationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosMobileAppConfigurationCollectionResponse") {
		var out IosMobileAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosMobileAppConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosNetworkUsageRuleCollectionResponse") {
		var out IosNetworkUsageRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosNetworkUsageRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosNotificationSettingsCollectionResponse") {
		var out IosNotificationSettingsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosNotificationSettingsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosStoreAppCollectionResponse") {
		var out IosStoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosStoreAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosUpdateConfigurationCollectionResponse") {
		var out IosUpdateConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosUpdateConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosUpdateDeviceStatusCollectionResponse") {
		var out IosUpdateDeviceStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosUpdateDeviceStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppCollectionResponse") {
		var out IosVppAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppEBookAssignmentCollectionResponse") {
		var out IosVppEBookAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppEBookAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppEBookCollectionResponse") {
		var out IosVppEBookCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppEBookCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosiPadOSWebClipCollectionResponse") {
		var out IosiPadOSWebClipCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosiPadOSWebClipCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemActivityCollectionResponse") {
		var out ItemActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemActivityStatCollectionResponse") {
		var out ItemActivityStatCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemActivityStatCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemAttachmentCollectionResponse") {
		var out ItemAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyCredentialCollectionResponse") {
		var out KeyCredentialCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyCredentialCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyCredentialConfigurationCollectionResponse") {
		var out KeyCredentialConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyCredentialConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyValueCollectionResponse") {
		var out KeyValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyValuePairCollectionResponse") {
		var out KeyValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.landingPageCollectionResponse") {
		var out LandingPageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LandingPageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.landingPageDetailCollectionResponse") {
		var out LandingPageDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LandingPageDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningAssignmentCollectionResponse") {
		var out LearningAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningContentCollectionResponse") {
		var out LearningContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningContentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningCourseActivityCollectionResponse") {
		var out LearningCourseActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningCourseActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningProviderCollectionResponse") {
		var out LearningProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningSelfInitiatedCourseCollectionResponse") {
		var out LearningSelfInitiatedCourseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningSelfInitiatedCourseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.licenseAssignmentStateCollectionResponse") {
		var out LicenseAssignmentStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LicenseAssignmentStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.licenseDetailsCollectionResponse") {
		var out LicenseDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LicenseDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.linkedResourceCollectionResponse") {
		var out LinkedResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LinkedResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.listCollectionResponse") {
		var out ListCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ListCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.listItemCollectionResponse") {
		var out ListItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ListItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.listItemVersionCollectionResponse") {
		var out ListItemVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ListItemVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.localizedNotificationMessageCollectionResponse") {
		var out LocalizedNotificationMessageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocalizedNotificationMessageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.locationCollectionResponse") {
		var out LocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.locationConstraintItemCollectionResponse") {
		var out LocationConstraintItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocationConstraintItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.loginPageCollectionResponse") {
		var out LoginPageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LoginPageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.longRunningOperationCollectionResponse") {
		var out LongRunningOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LongRunningOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSCompliancePolicyCollectionResponse") {
		var out MacOSCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSCustomConfigurationCollectionResponse") {
		var out MacOSCustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSCustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSDeviceFeaturesConfigurationCollectionResponse") {
		var out MacOSDeviceFeaturesConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSDeviceFeaturesConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSDmgAppCollectionResponse") {
		var out MacOSDmgAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSDmgAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSGeneralDeviceConfigurationCollectionResponse") {
		var out MacOSGeneralDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSGeneralDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSIncludedAppCollectionResponse") {
		var out MacOSIncludedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSIncludedAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSLobAppCollectionResponse") {
		var out MacOSLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSLobChildAppCollectionResponse") {
		var out MacOSLobChildAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSLobChildAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSMicrosoftDefenderAppCollectionResponse") {
		var out MacOSMicrosoftDefenderAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSMicrosoftDefenderAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSMicrosoftEdgeAppCollectionResponse") {
		var out MacOSMicrosoftEdgeAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSMicrosoftEdgeAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSOfficeSuiteAppCollectionResponse") {
		var out MacOSOfficeSuiteAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSOfficeSuiteAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailAssessmentRequestCollectionResponse") {
		var out MailAssessmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailAssessmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailFolderCollectionResponse") {
		var out MailFolderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailFolderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailSearchFolderCollectionResponse") {
		var out MailSearchFolderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailSearchFolderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxProtectionRuleCollectionResponse") {
		var out MailboxProtectionRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxProtectionRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxProtectionUnitCollectionResponse") {
		var out MailboxProtectionUnitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxProtectionUnitCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxProtectionUnitsBulkAdditionJobCollectionResponse") {
		var out MailboxProtectionUnitsBulkAdditionJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxProtectionUnitsBulkAdditionJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxRestoreArtifactCollectionResponse") {
		var out MailboxRestoreArtifactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxRestoreArtifactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxRestoreArtifactsBulkAdditionRequestCollectionResponse") {
		var out MailboxRestoreArtifactsBulkAdditionRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxRestoreArtifactsBulkAdditionRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.malwareStateCollectionResponse") {
		var out MalwareStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MalwareStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.malwareStateForWindowsDeviceCollectionResponse") {
		var out MalwareStateForWindowsDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MalwareStateForWindowsDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAndroidLobAppCollectionResponse") {
		var out ManagedAndroidLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAndroidLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAndroidStoreAppCollectionResponse") {
		var out ManagedAndroidStoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAndroidStoreAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppCollectionResponse") {
		var out ManagedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppConfigurationCollectionResponse") {
		var out ManagedAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppOperationCollectionResponse") {
		var out ManagedAppOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppPolicyCollectionResponse") {
		var out ManagedAppPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppPolicyDeploymentSummaryPerAppCollectionResponse") {
		var out ManagedAppPolicyDeploymentSummaryPerAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppPolicyDeploymentSummaryPerAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppProtectionCollectionResponse") {
		var out ManagedAppProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppRegistrationCollectionResponse") {
		var out ManagedAppRegistrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppRegistrationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppStatusCollectionResponse") {
		var out ManagedAppStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppStatusRawCollectionResponse") {
		var out ManagedAppStatusRawCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppStatusRawCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceCollectionResponse") {
		var out ManagedDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationAssignmentCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatusCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationDeviceStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationDeviceStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationUserStatusCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationUserStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationUserStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedEBookAssignmentCollectionResponse") {
		var out ManagedEBookAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedEBookAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedEBookCollectionResponse") {
		var out ManagedEBookCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedEBookCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedIOSLobAppCollectionResponse") {
		var out ManagedIOSLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIOSLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedIOSStoreAppCollectionResponse") {
		var out ManagedIOSStoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIOSStoreAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedMobileAppCollectionResponse") {
		var out ManagedMobileAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedMobileAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedMobileLobAppCollectionResponse") {
		var out ManagedMobileLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedMobileLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mdmWindowsInformationProtectionPolicyCollectionResponse") {
		var out MdmWindowsInformationProtectionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MdmWindowsInformationProtectionPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mediaInfoCollectionResponse") {
		var out MediaInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MediaInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mediaStreamCollectionResponse") {
		var out MediaStreamCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MediaStreamCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingAttendanceReportCollectionResponse") {
		var out MeetingAttendanceReportCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingAttendanceReportCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingParticipantInfoCollectionResponse") {
		var out MeetingParticipantInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingParticipantInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingTimeSuggestionCollectionResponse") {
		var out MeetingTimeSuggestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingTimeSuggestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.membershipOutlierInsightCollectionResponse") {
		var out MembershipOutlierInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MembershipOutlierInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messageCollectionResponse") {
		var out MessageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messageRuleCollectionResponse") {
		var out MessageRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messageSecurityStateCollectionResponse") {
		var out MessageSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.metaDataKeyStringPairCollectionResponse") {
		var out MetaDataKeyStringPairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MetaDataKeyStringPairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAccountUserConversationMemberCollectionResponse") {
		var out MicrosoftAccountUserConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAccountUserConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodCollectionResponse") {
		var out MicrosoftAuthenticatorAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAuthenticatorAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse") {
		var out MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodTargetCollectionResponse") {
		var out MicrosoftAuthenticatorAuthenticationMethodTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAuthenticatorAuthenticationMethodTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftStoreForBusinessAppCollectionResponse") {
		var out MicrosoftStoreForBusinessAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftStoreForBusinessAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftTrainingAssignmentMappingCollectionResponse") {
		var out MicrosoftTrainingAssignmentMappingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTrainingAssignmentMappingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppAssignmentCollectionResponse") {
		var out MobileAppAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppCategoryCollectionResponse") {
		var out MobileAppCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppCollectionResponse") {
		var out MobileAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppContentCollectionResponse") {
		var out MobileAppContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppContentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppContentFileCollectionResponse") {
		var out MobileAppContentFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppContentFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingEventCollectionResponse") {
		var out MobileAppTroubleshootingEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileContainedAppCollectionResponse") {
		var out MobileContainedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileContainedAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileLobAppCollectionResponse") {
		var out MobileLobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileLobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileThreatDefenseConnectorCollectionResponse") {
		var out MobileThreatDefenseConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileThreatDefenseConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.modifiedPropertyCollectionResponse") {
		var out ModifiedPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModifiedPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiTenantOrganizationMemberCollectionResponse") {
		var out MultiTenantOrganizationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiTenantOrganizationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiValueLegacyExtendedPropertyCollectionResponse") {
		var out MultiValueLegacyExtendedPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiValueLegacyExtendedPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.muteParticipantOperationCollectionResponse") {
		var out MuteParticipantOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MuteParticipantOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.namedLocationCollectionResponse") {
		var out NamedLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NamedLocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkConnectionCollectionResponse") {
		var out NetworkConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.notebookCollectionResponse") {
		var out NotebookCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NotebookCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.notificationMessageTemplateCollectionResponse") {
		var out NotificationMessageTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NotificationMessageTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oAuth2PermissionGrantCollectionResponse") {
		var out OAuth2PermissionGrantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OAuth2PermissionGrantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.objectDefinitionCollectionResponse") {
		var out ObjectDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ObjectDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.objectDefinitionMetadataEntryCollectionResponse") {
		var out ObjectDefinitionMetadataEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ObjectDefinitionMetadataEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.objectIdentityCollectionResponse") {
		var out ObjectIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ObjectIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.objectMappingCollectionResponse") {
		var out ObjectMappingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ObjectMappingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.objectMappingMetadataEntryCollectionResponse") {
		var out ObjectMappingMetadataEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ObjectMappingMetadataEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.offerShiftRequestCollectionResponse") {
		var out OfferShiftRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfferShiftRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingCollectionResponse") {
		var out OmaSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionListenerCollectionResponse") {
		var out OnAttributeCollectionListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionStartCustomExtensionCollectionResponse") {
		var out OnAttributeCollectionStartCustomExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionStartCustomExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionStartListenerCollectionResponse") {
		var out OnAttributeCollectionStartListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionStartListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionSubmitCustomExtensionCollectionResponse") {
		var out OnAttributeCollectionSubmitCustomExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionSubmitCustomExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionSubmitListenerCollectionResponse") {
		var out OnAttributeCollectionSubmitListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionSubmitListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAuthenticationMethodLoadStartListenerCollectionResponse") {
		var out OnAuthenticationMethodLoadStartListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAuthenticationMethodLoadStartListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onInteractiveAuthFlowStartListenerCollectionResponse") {
		var out OnInteractiveAuthFlowStartListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnInteractiveAuthFlowStartListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesDirectorySynchronizationCollectionResponse") {
		var out OnPremisesDirectorySynchronizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesDirectorySynchronizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesProvisioningErrorCollectionResponse") {
		var out OnPremisesProvisioningErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesProvisioningErrorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onTokenIssuanceStartCustomExtensionCollectionResponse") {
		var out OnTokenIssuanceStartCustomExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnTokenIssuanceStartCustomExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onTokenIssuanceStartListenerCollectionResponse") {
		var out OnTokenIssuanceStartListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnTokenIssuanceStartListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onTokenIssuanceStartReturnClaimCollectionResponse") {
		var out OnTokenIssuanceStartReturnClaimCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnTokenIssuanceStartReturnClaimCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onUserCreateStartListenerCollectionResponse") {
		var out OnUserCreateStartListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnUserCreateStartListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oneDriveForBusinessProtectionPolicyCollectionResponse") {
		var out OneDriveForBusinessProtectionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OneDriveForBusinessProtectionPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oneDriveForBusinessRestoreSessionCollectionResponse") {
		var out OneDriveForBusinessRestoreSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OneDriveForBusinessRestoreSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteOperationCollectionResponse") {
		var out OnenoteOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenotePageCollectionResponse") {
		var out OnenotePageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenotePageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteResourceCollectionResponse") {
		var out OnenoteResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteSectionCollectionResponse") {
		var out OnenoteSectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteSectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onlineMeetingCollectionResponse") {
		var out OnlineMeetingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnlineMeetingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openShiftChangeRequestCollectionResponse") {
		var out OpenShiftChangeRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenShiftChangeRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openShiftCollectionResponse") {
		var out OpenShiftCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenShiftCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openTypeExtensionCollectionResponse") {
		var out OpenTypeExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenTypeExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.optionalClaimCollectionResponse") {
		var out OptionalClaimCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OptionalClaimCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.orgContactCollectionResponse") {
		var out OrgContactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrgContactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.organizationCollectionResponse") {
		var out OrganizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.organizationalBrandingLocalizationCollectionResponse") {
		var out OrganizationalBrandingLocalizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizationalBrandingLocalizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.osVersionCountCollectionResponse") {
		var out OsVersionCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OsVersionCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outlookCategoryCollectionResponse") {
		var out OutlookCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.participantCollectionResponse") {
		var out ParticipantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ParticipantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.blobCollectionResponse") {
		var out PartnersBillingBlobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingBlobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.exportSuccessOperationCollectionResponse") {
		var out PartnersBillingExportSuccessOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingExportSuccessOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.failedOperationCollectionResponse") {
		var out PartnersBillingFailedOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingFailedOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.manifestCollectionResponse") {
		var out PartnersBillingManifestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingManifestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.operationCollectionResponse") {
		var out PartnersBillingOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.runningOperationCollectionResponse") {
		var out PartnersBillingRunningOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingRunningOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passwordAuthenticationMethodCollectionResponse") {
		var out PasswordAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passwordCredentialCollectionResponse") {
		var out PasswordCredentialCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordCredentialCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passwordCredentialConfigurationCollectionResponse") {
		var out PasswordCredentialConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordCredentialConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.payloadCoachmarkCollectionResponse") {
		var out PayloadCoachmarkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PayloadCoachmarkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.payloadCollectionResponse") {
		var out PayloadCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PayloadCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionCollectionResponse") {
		var out PermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionGrantConditionSetCollectionResponse") {
		var out PermissionGrantConditionSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionGrantConditionSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionGrantPolicyCollectionResponse") {
		var out PermissionGrantPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionGrantPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionScopeCollectionResponse") {
		var out PermissionScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personCollectionResponse") {
		var out PersonCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.phoneAuthenticationMethodCollectionResponse") {
		var out PhoneAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PhoneAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.phoneCollectionResponse") {
		var out PhoneCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PhoneCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.physicalAddressCollectionResponse") {
		var out PhysicalAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PhysicalAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.physicalOfficeAddressCollectionResponse") {
		var out PhysicalOfficeAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PhysicalOfficeAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pinnedChatMessageInfoCollectionResponse") {
		var out PinnedChatMessageInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PinnedChatMessageInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pkcs12CertificateInformationCollectionResponse") {
		var out Pkcs12CertificateInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Pkcs12CertificateInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.placeCollectionResponse") {
		var out PlaceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlaceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBucketCollectionResponse") {
		var out PlannerBucketCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucketCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanCollectionResponse") {
		var out PlannerPlanCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTaskCollectionResponse") {
		var out PlannerTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.platformCredentialAuthenticationMethodCollectionResponse") {
		var out PlatformCredentialAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlatformCredentialAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.playPromptOperationCollectionResponse") {
		var out PlayPromptOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlayPromptOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyBaseCollectionResponse") {
		var out PolicyBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.postCollectionResponse") {
		var out PostCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PostCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.preAuthorizedApplicationCollectionResponse") {
		var out PreAuthorizedApplicationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PreAuthorizedApplicationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.presenceCollectionResponse") {
		var out PresenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PresenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printConnectorCollectionResponse") {
		var out PrintConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printDocumentCollectionResponse") {
		var out PrintDocumentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintDocumentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printJobCollectionResponse") {
		var out PrintJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printOperationCollectionResponse") {
		var out PrintOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printServiceCollectionResponse") {
		var out PrintServiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintServiceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printServiceEndpointCollectionResponse") {
		var out PrintServiceEndpointCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintServiceEndpointCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printTaskCollectionResponse") {
		var out PrintTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printTaskDefinitionCollectionResponse") {
		var out PrintTaskDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintTaskDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printTaskTriggerCollectionResponse") {
		var out PrintTaskTriggerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintTaskTriggerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printUsageByPrinterCollectionResponse") {
		var out PrintUsageByPrinterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintUsageByPrinterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printUsageByUserCollectionResponse") {
		var out PrintUsageByUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintUsageByUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printerCollectionResponse") {
		var out PrinterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrinterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printerCreateOperationCollectionResponse") {
		var out PrinterCreateOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrinterCreateOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printerShareCollectionResponse") {
		var out PrinterShareCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrinterShareCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupAssignmentScheduleCollectionResponse") {
		var out PrivilegedAccessGroupAssignmentScheduleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupAssignmentScheduleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupAssignmentScheduleInstanceCollectionResponse") {
		var out PrivilegedAccessGroupAssignmentScheduleInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupAssignmentScheduleInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupAssignmentScheduleRequestCollectionResponse") {
		var out PrivilegedAccessGroupAssignmentScheduleRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupAssignmentScheduleRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupEligibilityScheduleCollectionResponse") {
		var out PrivilegedAccessGroupEligibilityScheduleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilityScheduleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupEligibilityScheduleInstanceCollectionResponse") {
		var out PrivilegedAccessGroupEligibilityScheduleInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilityScheduleInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupEligibilityScheduleRequestCollectionResponse") {
		var out PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.processCollectionResponse") {
		var out ProcessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profileCardAnnotationCollectionResponse") {
		var out ProfileCardAnnotationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfileCardAnnotationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profileCardPropertyCollectionResponse") {
		var out ProfileCardPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfileCardPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profilePhotoCollectionResponse") {
		var out ProfilePhotoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfilePhotoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectionPolicyBaseCollectionResponse") {
		var out ProtectionPolicyBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectionPolicyBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectionUnitBaseCollectionResponse") {
		var out ProtectionUnitBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectionUnitBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisionedPlanCollectionResponse") {
		var out ProvisionedPlanCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisionedPlanCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisioningObjectSummaryCollectionResponse") {
		var out ProvisioningObjectSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisioningObjectSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisioningStepCollectionResponse") {
		var out ProvisioningStepCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisioningStepCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.proxiedDomainCollectionResponse") {
		var out ProxiedDomainCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProxiedDomainCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.publicErrorDetailCollectionResponse") {
		var out PublicErrorDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PublicErrorDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.recipientCollectionResponse") {
		var out RecipientCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecipientCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.recommendedActionCollectionResponse") {
		var out RecommendedActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecommendedActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.recordOperationCollectionResponse") {
		var out RecordOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecordOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.recycleBinCollectionResponse") {
		var out RecycleBinCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecycleBinCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.recycleBinItemCollectionResponse") {
		var out RecycleBinItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecycleBinItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.redirectUriSettingsCollectionResponse") {
		var out RedirectUriSettingsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RedirectUriSettingsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.referenceAttachmentCollectionResponse") {
		var out ReferenceAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ReferenceAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.referencedObjectCollectionResponse") {
		var out ReferencedObjectCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ReferencedObjectCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.registryKeyStateCollectionResponse") {
		var out RegistryKeyStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryKeyStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.relatedContactCollectionResponse") {
		var out RelatedContactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RelatedContactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.remoteAssistancePartnerCollectionResponse") {
		var out RemoteAssistancePartnerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RemoteAssistancePartnerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.requiredResourceAccessCollectionResponse") {
		var out RequiredResourceAccessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RequiredResourceAccessCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resellerDelegatedAdminRelationshipCollectionResponse") {
		var out ResellerDelegatedAdminRelationshipCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResellerDelegatedAdminRelationshipCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceAccessCollectionResponse") {
		var out ResourceAccessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceAccessCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceActionCollectionResponse") {
		var out ResourceActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceOperationCollectionResponse") {
		var out ResourceOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourcePermissionCollectionResponse") {
		var out ResourcePermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourcePermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceSpecificPermissionCollectionResponse") {
		var out ResourceSpecificPermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceSpecificPermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceSpecificPermissionGrantCollectionResponse") {
		var out ResourceSpecificPermissionGrantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceSpecificPermissionGrantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restorePointCollectionResponse") {
		var out RestorePointCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestorePointCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restorePointSearchResultCollectionResponse") {
		var out RestorePointSearchResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestorePointSearchResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restoreSessionBaseCollectionResponse") {
		var out RestoreSessionBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestoreSessionBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.retentionSettingCollectionResponse") {
		var out RetentionSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RetentionSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.richLongRunningOperationCollectionResponse") {
		var out RichLongRunningOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RichLongRunningOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskDetectionCollectionResponse") {
		var out RiskDetectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskDetectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyServicePrincipalCollectionResponse") {
		var out RiskyServicePrincipalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyServicePrincipalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyServicePrincipalHistoryItemCollectionResponse") {
		var out RiskyServicePrincipalHistoryItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyServicePrincipalHistoryItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyUserCollectionResponse") {
		var out RiskyUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyUserHistoryItemCollectionResponse") {
		var out RiskyUserHistoryItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyUserHistoryItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleAssignmentCollectionResponse") {
		var out RoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleDefinitionCollectionResponse") {
		var out RoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rolePermissionCollectionResponse") {
		var out RolePermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RolePermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roomCollectionResponse") {
		var out RoomCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoomCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roomListCollectionResponse") {
		var out RoomListCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoomListCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rubricCriterionCollectionResponse") {
		var out RubricCriterionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RubricCriterionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rubricLevelCollectionResponse") {
		var out RubricLevelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RubricLevelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rubricQualityCollectionResponse") {
		var out RubricQualityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RubricQualityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rubricQualityFeedbackModelCollectionResponse") {
		var out RubricQualityFeedbackModelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RubricQualityFeedbackModelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rubricQualitySelectedColumnModelCollectionResponse") {
		var out RubricQualitySelectedColumnModelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RubricQualitySelectedColumnModelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.samlOrWsFedExternalDomainFederationCollectionResponse") {
		var out SamlOrWsFedExternalDomainFederationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SamlOrWsFedExternalDomainFederationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.samlOrWsFedProviderCollectionResponse") {
		var out SamlOrWsFedProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SamlOrWsFedProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.scheduleItemCollectionResponse") {
		var out ScheduleItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScheduleItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.schedulingGroupCollectionResponse") {
		var out SchedulingGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SchedulingGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.schemaExtensionCollectionResponse") {
		var out SchemaExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SchemaExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.scopedRoleMembershipCollectionResponse") {
		var out ScopedRoleMembershipCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScopedRoleMembershipCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.scoredEmailAddressCollectionResponse") {
		var out ScoredEmailAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScoredEmailAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.acronymCollectionResponse") {
		var out SearchAcronymCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchAcronymCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.searchAggregationCollectionResponse") {
		var out SearchAggregationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchAggregationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.answerVariantCollectionResponse") {
		var out SearchAnswerVariantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchAnswerVariantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.bookmarkCollectionResponse") {
		var out SearchBookmarkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchBookmarkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.searchBucketCollectionResponse") {
		var out SearchBucketCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchBucketCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.searchHitCollectionResponse") {
		var out SearchHitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchHitCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.searchHitsContainerCollectionResponse") {
		var out SearchHitsContainerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchHitsContainerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.qnaCollectionResponse") {
		var out SearchQnaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchQnaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sectionGroupCollectionResponse") {
		var out SectionGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SectionGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secureScoreCollectionResponse") {
		var out SecureScoreCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecureScoreCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secureScoreControlProfileCollectionResponse") {
		var out SecureScoreControlProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecureScoreControlProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secureScoreControlStateUpdateCollectionResponse") {
		var out SecureScoreControlStateUpdateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecureScoreControlStateUpdateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.alertCollectionResponse") {
		var out SecurityAlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.alertCommentCollectionResponse") {
		var out SecurityAlertCommentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlertCommentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.alertEvidenceCollectionResponse") {
		var out SecurityAlertEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlertEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.articleCollectionResponse") {
		var out SecurityArticleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityArticleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.articleIndicatorCollectionResponse") {
		var out SecurityArticleIndicatorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityArticleIndicatorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.authorityTemplateCollectionResponse") {
		var out SecurityAuthorityTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAuthorityTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.caseOperationCollectionResponse") {
		var out SecurityCaseOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCaseOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.categoryTemplateCollectionResponse") {
		var out SecurityCategoryTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCategoryTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.citationTemplateCollectionResponse") {
		var out SecurityCitationTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCitationTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.containerEvidenceCollectionResponse") {
		var out SecurityContainerEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityContainerEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataSourceCollectionResponse") {
		var out SecurityDataSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.departmentTemplateCollectionResponse") {
		var out SecurityDepartmentTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDepartmentTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dispositionReviewStageCollectionResponse") {
		var out SecurityDispositionReviewStageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDispositionReviewStageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryAddToReviewSetOperationCollectionResponse") {
		var out SecurityEdiscoveryAddToReviewSetOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryAddToReviewSetOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryCaseCollectionResponse") {
		var out SecurityEdiscoveryCaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryCaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryCustodianCollectionResponse") {
		var out SecurityEdiscoveryCustodianCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryCustodianCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryEstimateOperationCollectionResponse") {
		var out SecurityEdiscoveryEstimateOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryEstimateOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryExportOperationCollectionResponse") {
		var out SecurityEdiscoveryExportOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryExportOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryHoldOperationCollectionResponse") {
		var out SecurityEdiscoveryHoldOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryHoldOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryIndexOperationCollectionResponse") {
		var out SecurityEdiscoveryIndexOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryIndexOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryNoncustodialDataSourceCollectionResponse") {
		var out SecurityEdiscoveryNoncustodialDataSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryNoncustodialDataSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryPurgeDataOperationCollectionResponse") {
		var out SecurityEdiscoveryPurgeDataOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryPurgeDataOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryReviewSetCollectionResponse") {
		var out SecurityEdiscoveryReviewSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryReviewSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryReviewSetQueryCollectionResponse") {
		var out SecurityEdiscoveryReviewSetQueryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryReviewSetQueryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryReviewTagCollectionResponse") {
		var out SecurityEdiscoveryReviewTagCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryReviewTagCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoverySearchCollectionResponse") {
		var out SecurityEdiscoverySearchCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoverySearchCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoverySearchExportOperationCollectionResponse") {
		var out SecurityEdiscoverySearchExportOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoverySearchExportOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryTagOperationCollectionResponse") {
		var out SecurityEdiscoveryTagOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryTagOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.eventPropagationResultCollectionResponse") {
		var out SecurityEventPropagationResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEventPropagationResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.eventQueryCollectionResponse") {
		var out SecurityEventQueryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEventQueryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exportFileMetadataCollectionResponse") {
		var out SecurityExportFileMetadataCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExportFileMetadataCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fileEvidenceCollectionResponse") {
		var out SecurityFileEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fileHashCollectionResponse") {
		var out SecurityFileHashCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileHashCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanReferenceTemplateCollectionResponse") {
		var out SecurityFilePlanReferenceTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanReferenceTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.healthIssueCollectionResponse") {
		var out SecurityHealthIssueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHealthIssueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostCollectionResponse") {
		var out SecurityHostCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostComponentCollectionResponse") {
		var out SecurityHostComponentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostComponentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostCookieCollectionResponse") {
		var out SecurityHostCookieCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostCookieCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostPairCollectionResponse") {
		var out SecurityHostPairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostPairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostPortBannerCollectionResponse") {
		var out SecurityHostPortBannerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostPortBannerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostPortCollectionResponse") {
		var out SecurityHostPortCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostPortCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostPortComponentCollectionResponse") {
		var out SecurityHostPortComponentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostPortComponentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostReputationRuleCollectionResponse") {
		var out SecurityHostReputationRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostReputationRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostSslCertificateCollectionResponse") {
		var out SecurityHostSslCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostSslCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostSslCertificatePortCollectionResponse") {
		var out SecurityHostSslCertificatePortCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostSslCertificatePortCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostTrackerCollectionResponse") {
		var out SecurityHostTrackerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostTrackerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostnameCollectionResponse") {
		var out SecurityHostnameCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostnameCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.huntingRowResultCollectionResponse") {
		var out SecurityHuntingRowResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHuntingRowResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hyperlinkCollectionResponse") {
		var out SecurityHyperlinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHyperlinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ipAddressCollectionResponse") {
		var out SecurityIPAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIPAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ipEvidenceCollectionResponse") {
		var out SecurityIPEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIPEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.incidentCollectionResponse") {
		var out SecurityIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIncidentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.intelligenceProfileCollectionResponse") {
		var out SecurityIntelligenceProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIntelligenceProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.intelligenceProfileCountryOrRegionOfOriginCollectionResponse") {
		var out SecurityIntelligenceProfileCountryOrRegionOfOriginCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIntelligenceProfileCountryOrRegionOfOriginCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.intelligenceProfileIndicatorCollectionResponse") {
		var out SecurityIntelligenceProfileIndicatorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIntelligenceProfileIndicatorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.kubernetesServicePortCollectionResponse") {
		var out SecurityKubernetesServicePortCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityKubernetesServicePortCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.loggedOnUserCollectionResponse") {
		var out SecurityLoggedOnUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityLoggedOnUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.networkAdapterCollectionResponse") {
		var out SecurityNetworkAdapterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityNetworkAdapterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.nicEvidenceCollectionResponse") {
		var out SecurityNicEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityNicEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.passiveDnsRecordCollectionResponse") {
		var out SecurityPassiveDnsRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPassiveDnsRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.processEvidenceCollectionResponse") {
		var out SecurityProcessEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProcessEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityResourceCollectionResponse") {
		var out SecurityResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionEventCollectionResponse") {
		var out SecurityRetentionEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionEventTypeCollectionResponse") {
		var out SecurityRetentionEventTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionEventTypeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionLabelCollectionResponse") {
		var out SecurityRetentionLabelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionLabelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sensorCollectionResponse") {
		var out SecuritySensorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySensorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.singlePropertySchemaCollectionResponse") {
		var out SecuritySinglePropertySchemaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySinglePropertySchemaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.siteSourceCollectionResponse") {
		var out SecuritySiteSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySiteSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sslCertificateCollectionResponse") {
		var out SecuritySslCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySslCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.subcategoryTemplateCollectionResponse") {
		var out SecuritySubcategoryTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySubcategoryTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.subdomainCollectionResponse") {
		var out SecuritySubdomainCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySubdomainCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.unifiedGroupSourceCollectionResponse") {
		var out SecurityUnifiedGroupSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUnifiedGroupSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.urlEvidenceCollectionResponse") {
		var out SecurityUrlEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUrlEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.userSourceCollectionResponse") {
		var out SecurityUserSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUserSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vulnerabilityCollectionResponse") {
		var out SecurityVulnerabilityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVulnerabilityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vulnerabilityComponentCollectionResponse") {
		var out SecurityVulnerabilityComponentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVulnerabilityComponentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.whoisHistoryRecordCollectionResponse") {
		var out SecurityWhoisHistoryRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWhoisHistoryRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.whoisNameserverCollectionResponse") {
		var out SecurityWhoisNameserverCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWhoisNameserverCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.whoisRecordCollectionResponse") {
		var out SecurityWhoisRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWhoisRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sendDtmfTonesOperationCollectionResponse") {
		var out SendDtmfTonesOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SendDtmfTonesOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sensitivityLabelAssignmentCollectionResponse") {
		var out SensitivityLabelAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SensitivityLabelAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceAnnouncementAttachmentCollectionResponse") {
		var out ServiceAnnouncementAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceAnnouncementAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceAppCollectionResponse") {
		var out ServiceAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceHealthCollectionResponse") {
		var out ServiceHealthCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceHealthCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceHealthIssueCollectionResponse") {
		var out ServiceHealthIssueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceHealthIssueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceHealthIssuePostCollectionResponse") {
		var out ServiceHealthIssuePostCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceHealthIssuePostCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePlanInfoCollectionResponse") {
		var out ServicePlanInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePlanInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalCollectionResponse") {
		var out ServicePrincipalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalRiskDetectionCollectionResponse") {
		var out ServicePrincipalRiskDetectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalRiskDetectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceProvisioningErrorCollectionResponse") {
		var out ServiceProvisioningErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceProvisioningErrorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceStorageQuotaBreakdownCollectionResponse") {
		var out ServiceStorageQuotaBreakdownCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceStorageQuotaBreakdownCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceUpdateMessageCollectionResponse") {
		var out ServiceUpdateMessageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceUpdateMessageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.settingSourceCollectionResponse") {
		var out SettingSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SettingSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.settingStateDeviceSummaryCollectionResponse") {
		var out SettingStateDeviceSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SettingStateDeviceSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.settingTemplateValueCollectionResponse") {
		var out SettingTemplateValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SettingTemplateValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.settingValueCollectionResponse") {
		var out SettingValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SettingValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointIdentitySetCollectionResponse") {
		var out SharePointIdentitySetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointIdentitySetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointProtectionPolicyCollectionResponse") {
		var out SharePointProtectionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointProtectionPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointRestoreSessionCollectionResponse") {
		var out SharePointRestoreSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointRestoreSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedDriveItemCollectionResponse") {
		var out SharedDriveItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedDriveItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedInsightCollectionResponse") {
		var out SharedInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedPCConfigurationCollectionResponse") {
		var out SharedPCConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedPCConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedWithChannelTeamInfoCollectionResponse") {
		var out SharedWithChannelTeamInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedWithChannelTeamInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharingDetailCollectionResponse") {
		var out SharingDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharingDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shiftActivityCollectionResponse") {
		var out ShiftActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shiftAvailabilityCollectionResponse") {
		var out ShiftAvailabilityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftAvailabilityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shiftCollectionResponse") {
		var out ShiftCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.signInCollectionResponse") {
		var out SignInCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SignInCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationAutomationCollectionResponse") {
		var out SimulationAutomationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationAutomationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationAutomationRunCollectionResponse") {
		var out SimulationAutomationRunCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationAutomationRunCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationCollectionResponse") {
		var out SimulationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationEventCollectionResponse") {
		var out SimulationEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleValueLegacyExtendedPropertyCollectionResponse") {
		var out SingleValueLegacyExtendedPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleValueLegacyExtendedPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteCollectionResponse") {
		var out SiteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sitePageCollectionResponse") {
		var out SitePageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SitePageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteProtectionRuleCollectionResponse") {
		var out SiteProtectionRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteProtectionRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteProtectionUnitCollectionResponse") {
		var out SiteProtectionUnitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteProtectionUnitCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteProtectionUnitsBulkAdditionJobCollectionResponse") {
		var out SiteProtectionUnitsBulkAdditionJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteProtectionUnitsBulkAdditionJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteRestoreArtifactCollectionResponse") {
		var out SiteRestoreArtifactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteRestoreArtifactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteRestoreArtifactsBulkAdditionRequestCollectionResponse") {
		var out SiteRestoreArtifactsBulkAdditionRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteRestoreArtifactsBulkAdditionRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.skypeForBusinessUserConversationMemberCollectionResponse") {
		var out SkypeForBusinessUserConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SkypeForBusinessUserConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.skypeUserConversationMemberCollectionResponse") {
		var out SkypeUserConversationMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SkypeUserConversationMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.smsAuthenticationMethodConfigurationCollectionResponse") {
		var out SmsAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SmsAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.smsAuthenticationMethodTargetCollectionResponse") {
		var out SmsAuthenticationMethodTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SmsAuthenticationMethodTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.socialIdentityProviderCollectionResponse") {
		var out SocialIdentityProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SocialIdentityProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.softwareOathAuthenticationMethodCollectionResponse") {
		var out SoftwareOathAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SoftwareOathAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.softwareOathAuthenticationMethodConfigurationCollectionResponse") {
		var out SoftwareOathAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SoftwareOathAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sortPropertyCollectionResponse") {
		var out SortPropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SortPropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.standardWebPartCollectionResponse") {
		var out StandardWebPartCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StandardWebPartCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.startHoldMusicOperationCollectionResponse") {
		var out StartHoldMusicOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StartHoldMusicOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stopHoldMusicOperationCollectionResponse") {
		var out StopHoldMusicOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StopHoldMusicOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#StringCollectionResponse") {
		var out StringCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StringCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stringKeyAttributeMappingSourceValuePairCollectionResponse") {
		var out StringKeyAttributeMappingSourceValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StringKeyAttributeMappingSourceValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stringKeyLongValuePairCollectionResponse") {
		var out StringKeyLongValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StringKeyLongValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stringKeyObjectValuePairCollectionResponse") {
		var out StringKeyObjectValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StringKeyObjectValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stringKeyStringValuePairCollectionResponse") {
		var out StringKeyStringValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StringKeyStringValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stsPolicyCollectionResponse") {
		var out StsPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StsPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestCollectionResponse") {
		var out SubjectRightsRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestHistoryCollectionResponse") {
		var out SubjectRightsRequestHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestStageDetailCollectionResponse") {
		var out SubjectRightsRequestStageDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestStageDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectSetCollectionResponse") {
		var out SubjectSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subscribeToToneOperationCollectionResponse") {
		var out SubscribeToToneOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubscribeToToneOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subscribedSkuCollectionResponse") {
		var out SubscribedSkuCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubscribedSkuCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subscriptionCollectionResponse") {
		var out SubscriptionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubscriptionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.swapShiftsChangeRequestCollectionResponse") {
		var out SwapShiftsChangeRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SwapShiftsChangeRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationJobCollectionResponse") {
		var out SynchronizationJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationJobSubjectCollectionResponse") {
		var out SynchronizationJobSubjectCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationJobSubjectCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationMetadataEntryCollectionResponse") {
		var out SynchronizationMetadataEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationMetadataEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationProgressCollectionResponse") {
		var out SynchronizationProgressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationProgressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationRuleCollectionResponse") {
		var out SynchronizationRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationSecretKeyStringValuePairCollectionResponse") {
		var out SynchronizationSecretKeyStringValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationSecretKeyStringValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationTemplateCollectionResponse") {
		var out SynchronizationTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetDeviceGroupCollectionResponse") {
		var out TargetDeviceGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetDeviceGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetResourceCollectionResponse") {
		var out TargetResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppConfigurationCollectionResponse") {
		var out TargetedManagedAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppPolicyAssignmentCollectionResponse") {
		var out TargetedManagedAppPolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppPolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppProtectionCollectionResponse") {
		var out TargetedManagedAppProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.taskFileAttachmentCollectionResponse") {
		var out TaskFileAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TaskFileAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamCollectionResponse") {
		var out TeamCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppCollectionResponse") {
		var out TeamsAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppDefinitionCollectionResponse") {
		var out TeamsAppDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppInstallationCollectionResponse") {
		var out TeamsAppInstallationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppInstallationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppResourceSpecificPermissionCollectionResponse") {
		var out TeamsAppResourceSpecificPermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppResourceSpecificPermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAsyncOperationCollectionResponse") {
		var out TeamsAsyncOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAsyncOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsTabCollectionResponse") {
		var out TeamsTabCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsTabCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsTemplateCollectionResponse") {
		var out TeamsTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkTagCollectionResponse") {
		var out TeamworkTagCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkTagCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkTagMemberCollectionResponse") {
		var out TeamworkTagMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkTagMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkUserIdentityCollectionResponse") {
		var out TeamworkUserIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkUserIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.telecomExpenseManagementPartnerCollectionResponse") {
		var out TelecomExpenseManagementPartnerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TelecomExpenseManagementPartnerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teleconferenceDeviceMediaQualityCollectionResponse") {
		var out TeleconferenceDeviceMediaQualityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeleconferenceDeviceMediaQualityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.temporaryAccessPassAuthenticationMethodCollectionResponse") {
		var out TemporaryAccessPassAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TemporaryAccessPassAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.temporaryAccessPassAuthenticationMethodConfigurationCollectionResponse") {
		var out TemporaryAccessPassAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TemporaryAccessPassAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tenantAppManagementPolicyCollectionResponse") {
		var out TenantAppManagementPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TenantAppManagementPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.groupCollectionResponse") {
		var out TermStoreGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.localizedDescriptionCollectionResponse") {
		var out TermStoreLocalizedDescriptionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreLocalizedDescriptionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.localizedLabelCollectionResponse") {
		var out TermStoreLocalizedLabelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreLocalizedLabelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.localizedNameCollectionResponse") {
		var out TermStoreLocalizedNameCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreLocalizedNameCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.relationCollectionResponse") {
		var out TermStoreRelationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreRelationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.setCollectionResponse") {
		var out TermStoreSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.storeCollectionResponse") {
		var out TermStoreStoreCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreStoreCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.termCollectionResponse") {
		var out TermStoreTermCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreTermCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditionsAcceptanceStatusCollectionResponse") {
		var out TermsAndConditionsAcceptanceStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditionsAcceptanceStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditionsAssignmentCollectionResponse") {
		var out TermsAndConditionsAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditionsAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditionsCollectionResponse") {
		var out TermsAndConditionsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditionsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.textWebPartCollectionResponse") {
		var out TextWebPartCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextWebPartCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.threatAssessmentRequestCollectionResponse") {
		var out ThreatAssessmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThreatAssessmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.threatAssessmentResultCollectionResponse") {
		var out ThreatAssessmentResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThreatAssessmentResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.thumbnailSetCollectionResponse") {
		var out ThumbnailSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThumbnailSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeCardBreakCollectionResponse") {
		var out TimeCardBreakCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeCardBreakCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeCardCollectionResponse") {
		var out TimeCardCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeCardCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOffCollectionResponse") {
		var out TimeOffCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOffCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOffReasonCollectionResponse") {
		var out TimeOffReasonCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOffReasonCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOffRequestCollectionResponse") {
		var out TimeOffRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOffRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeRangeCollectionResponse") {
		var out TimeRangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeRangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeSlotCollectionResponse") {
		var out TimeSlotCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeSlotCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.todoTaskCollectionResponse") {
		var out TodoTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TodoTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.todoTaskListCollectionResponse") {
		var out TodoTaskListCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TodoTaskListCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tokenIssuancePolicyCollectionResponse") {
		var out TokenIssuancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TokenIssuancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tokenLifetimePolicyCollectionResponse") {
		var out TokenLifetimePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TokenLifetimePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trainingCollectionResponse") {
		var out TrainingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrainingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trainingLanguageDetailCollectionResponse") {
		var out TrainingLanguageDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrainingLanguageDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trendingCollectionResponse") {
		var out TrendingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrendingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedApprovalStageCollectionResponse") {
		var out UnifiedApprovalStageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedApprovalStageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRbacResourceActionCollectionResponse") {
		var out UnifiedRbacResourceActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRbacResourceActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRbacResourceNamespaceCollectionResponse") {
		var out UnifiedRbacResourceNamespaceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRbacResourceNamespaceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignmentCollectionResponse") {
		var out UnifiedRoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignmentScheduleCollectionResponse") {
		var out UnifiedRoleAssignmentScheduleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignmentScheduleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignmentScheduleInstanceCollectionResponse") {
		var out UnifiedRoleAssignmentScheduleInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignmentScheduleInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignmentScheduleRequestCollectionResponse") {
		var out UnifiedRoleAssignmentScheduleRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignmentScheduleRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleCollectionResponse") {
		var out UnifiedRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleDefinitionCollectionResponse") {
		var out UnifiedRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleEligibilityScheduleCollectionResponse") {
		var out UnifiedRoleEligibilityScheduleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleEligibilityScheduleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleEligibilityScheduleInstanceCollectionResponse") {
		var out UnifiedRoleEligibilityScheduleInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleEligibilityScheduleInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleEligibilityScheduleRequestCollectionResponse") {
		var out UnifiedRoleEligibilityScheduleRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleEligibilityScheduleRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyApprovalRuleCollectionResponse") {
		var out UnifiedRoleManagementPolicyApprovalRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyApprovalRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyAssignmentCollectionResponse") {
		var out UnifiedRoleManagementPolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyAuthenticationContextRuleCollectionResponse") {
		var out UnifiedRoleManagementPolicyAuthenticationContextRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyAuthenticationContextRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyCollectionResponse") {
		var out UnifiedRoleManagementPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyEnablementRuleCollectionResponse") {
		var out UnifiedRoleManagementPolicyEnablementRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyEnablementRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyExpirationRuleCollectionResponse") {
		var out UnifiedRoleManagementPolicyExpirationRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyExpirationRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyNotificationRuleCollectionResponse") {
		var out UnifiedRoleManagementPolicyNotificationRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyNotificationRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyRuleCollectionResponse") {
		var out UnifiedRoleManagementPolicyRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRolePermissionCollectionResponse") {
		var out UnifiedRolePermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRolePermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unmuteParticipantOperationCollectionResponse") {
		var out UnmuteParticipantOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnmuteParticipantOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.updateRecordingStatusOperationCollectionResponse") {
		var out UpdateRecordingStatusOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UpdateRecordingStatusOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.uriClickSecurityStateCollectionResponse") {
		var out UriClickSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UriClickSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.urlAssessmentRequestCollectionResponse") {
		var out UrlAssessmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UrlAssessmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.usedInsightCollectionResponse") {
		var out UsedInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UsedInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userActivityCollectionResponse") {
		var out UserActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userAttributeValuesItemCollectionResponse") {
		var out UserAttributeValuesItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserAttributeValuesItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userCollectionResponse") {
		var out UserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userConsentRequestCollectionResponse") {
		var out UserConsentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserConsentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByOSVersionCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthApplicationPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthApplicationPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthApplicationPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthDeviceModelPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthDeviceModelPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthDeviceModelPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformanceCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthDevicePerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthDevicePerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformanceDetailsCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthDevicePerformanceDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthDevicePerformanceDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthOSVersionPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthOSVersionPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthOSVersionPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBaselineCollectionResponse") {
		var out UserExperienceAnalyticsBaselineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBaselineCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsCategoryCollectionResponse") {
		var out UserExperienceAnalyticsCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDevicePerformanceCollectionResponse") {
		var out UserExperienceAnalyticsDevicePerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDevicePerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceScoresCollectionResponse") {
		var out UserExperienceAnalyticsDeviceScoresCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceScoresCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceStartupHistoryCollectionResponse") {
		var out UserExperienceAnalyticsDeviceStartupHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceStartupHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcessCollectionResponse") {
		var out UserExperienceAnalyticsDeviceStartupProcessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceStartupProcessCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsInsightCollectionResponse") {
		var out UserExperienceAnalyticsInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsInsightValueCollectionResponse") {
		var out UserExperienceAnalyticsInsightValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsInsightValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsMetricCollectionResponse") {
		var out UserExperienceAnalyticsMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsMetricHistoryCollectionResponse") {
		var out UserExperienceAnalyticsMetricHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsMetricHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsModelScoresCollectionResponse") {
		var out UserExperienceAnalyticsModelScoresCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsModelScoresCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsScoreHistoryCollectionResponse") {
		var out UserExperienceAnalyticsScoreHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsScoreHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereDeviceCollectionResponse") {
		var out UserExperienceAnalyticsWorkFromAnywhereDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetricCollectionResponse") {
		var out UserExperienceAnalyticsWorkFromAnywhereMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereModelPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userFlowLanguageConfigurationCollectionResponse") {
		var out UserFlowLanguageConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserFlowLanguageConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userFlowLanguagePageCollectionResponse") {
		var out UserFlowLanguagePageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserFlowLanguagePageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userInstallStateSummaryCollectionResponse") {
		var out UserInstallStateSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserInstallStateSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userRegistrationDetailsCollectionResponse") {
		var out UserRegistrationDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserRegistrationDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userRegistrationFeatureCountCollectionResponse") {
		var out UserRegistrationFeatureCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserRegistrationFeatureCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userRegistrationMethodCountCollectionResponse") {
		var out UserRegistrationMethodCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserRegistrationMethodCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userScopeTeamsAppInstallationCollectionResponse") {
		var out UserScopeTeamsAppInstallationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserScopeTeamsAppInstallationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSecurityStateCollectionResponse") {
		var out UserSecurityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSecurityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSignInInsightCollectionResponse") {
		var out UserSignInInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSignInInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSimulationDetailsCollectionResponse") {
		var out UserSimulationDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSimulationDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSimulationEventInfoCollectionResponse") {
		var out UserSimulationEventInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSimulationEventInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userTrainingEventInfoCollectionResponse") {
		var out UserTrainingEventInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserTrainingEventInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userTrainingStatusInfoCollectionResponse") {
		var out UserTrainingStatusInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserTrainingStatusInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiedDomainCollectionResponse") {
		var out VerifiedDomainCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiedDomainCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventCollectionResponse") {
		var out VirtualEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventExternalInformationCollectionResponse") {
		var out VirtualEventExternalInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventExternalInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventPresenterCollectionResponse") {
		var out VirtualEventPresenterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventPresenterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationCollectionResponse") {
		var out VirtualEventRegistrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationCustomQuestionCollectionResponse") {
		var out VirtualEventRegistrationCustomQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationCustomQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationPredefinedQuestionCollectionResponse") {
		var out VirtualEventRegistrationPredefinedQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationPredefinedQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationQuestionAnswerCollectionResponse") {
		var out VirtualEventRegistrationQuestionAnswerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationQuestionAnswerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationQuestionBaseCollectionResponse") {
		var out VirtualEventRegistrationQuestionBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationQuestionBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventSessionCollectionResponse") {
		var out VirtualEventSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventTownhallCollectionResponse") {
		var out VirtualEventTownhallCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventTownhallCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventWebinarCollectionResponse") {
		var out VirtualEventWebinarCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventWebinarCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.voiceAuthenticationMethodConfigurationCollectionResponse") {
		var out VoiceAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VoiceAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vppTokenCollectionResponse") {
		var out VppTokenCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VppTokenCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vulnerabilityStateCollectionResponse") {
		var out VulnerabilityStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VulnerabilityStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webAppCollectionResponse") {
		var out WebAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webPartCollectionResponse") {
		var out WebPartCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebPartCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.websiteCollectionResponse") {
		var out WebsiteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebsiteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppCollectionResponse") {
		var out Win32LobAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppReturnCodeCollectionResponse") {
		var out Win32LobAppReturnCodeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppReturnCodeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppRuleCollectionResponse") {
		var out Win32LobAppRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10CompliancePolicyCollectionResponse") {
		var out Windows10CompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10CompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10CustomConfigurationCollectionResponse") {
		var out Windows10CustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10CustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EndpointProtectionConfigurationCollectionResponse") {
		var out Windows10EndpointProtectionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EndpointProtectionConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EnrollmentCompletionPageConfigurationCollectionResponse") {
		var out Windows10EnrollmentCompletionPageConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EnrollmentCompletionPageConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EnterpriseModernAppManagementConfigurationCollectionResponse") {
		var out Windows10EnterpriseModernAppManagementConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EnterpriseModernAppManagementConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10GeneralConfigurationCollectionResponse") {
		var out Windows10GeneralConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10GeneralConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10MobileCompliancePolicyCollectionResponse") {
		var out Windows10MobileCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10MobileCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10SecureAssessmentConfigurationCollectionResponse") {
		var out Windows10SecureAssessmentConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10SecureAssessmentConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10TeamGeneralConfigurationCollectionResponse") {
		var out Windows10TeamGeneralConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10TeamGeneralConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81CompliancePolicyCollectionResponse") {
		var out Windows81CompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81CompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81GeneralConfigurationCollectionResponse") {
		var out Windows81GeneralConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81GeneralConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAppXCollectionResponse") {
		var out WindowsAppXCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAppXCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeviceIdentityCollectionResponse") {
		var out WindowsAutopilotDeviceIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeviceIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfigurationCollectionResponse") {
		var out WindowsDefenderAdvancedThreatProtectionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDefenderAdvancedThreatProtectionConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDeviceMalwareStateCollectionResponse") {
		var out WindowsDeviceMalwareStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDeviceMalwareStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsHelloForBusinessAuthenticationMethodCollectionResponse") {
		var out WindowsHelloForBusinessAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsHelloForBusinessAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionAppCollectionResponse") {
		var out WindowsInformationProtectionAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionAppLearningSummaryCollectionResponse") {
		var out WindowsInformationProtectionAppLearningSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionAppLearningSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionAppLockerFileCollectionResponse") {
		var out WindowsInformationProtectionAppLockerFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionAppLockerFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionCollectionResponse") {
		var out WindowsInformationProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionIPRangeCollectionCollectionResponse") {
		var out WindowsInformationProtectionIPRangeCollectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionIPRangeCollectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionNetworkLearningSummaryCollectionResponse") {
		var out WindowsInformationProtectionNetworkLearningSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionNetworkLearningSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionPolicyCollectionResponse") {
		var out WindowsInformationProtectionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionProxiedDomainCollectionCollectionResponse") {
		var out WindowsInformationProtectionProxiedDomainCollectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionProxiedDomainCollectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionResourceCollectionCollectionResponse") {
		var out WindowsInformationProtectionResourceCollectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionResourceCollectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareCategoryCountCollectionResponse") {
		var out WindowsMalwareCategoryCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareCategoryCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareExecutionStateCountCollectionResponse") {
		var out WindowsMalwareExecutionStateCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareExecutionStateCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareInformationCollectionResponse") {
		var out WindowsMalwareInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareNameCountCollectionResponse") {
		var out WindowsMalwareNameCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareNameCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareSeverityCountCollectionResponse") {
		var out WindowsMalwareSeverityCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareSeverityCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareStateCountCollectionResponse") {
		var out WindowsMalwareStateCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareStateCountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMicrosoftEdgeAppCollectionResponse") {
		var out WindowsMicrosoftEdgeAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMicrosoftEdgeAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMobileMSICollectionResponse") {
		var out WindowsMobileMSICollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMobileMSICollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81CompliancePolicyCollectionResponse") {
		var out WindowsPhone81CompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81CompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81CustomConfigurationCollectionResponse") {
		var out WindowsPhone81CustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81CustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81GeneralConfigurationCollectionResponse") {
		var out WindowsPhone81GeneralConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81GeneralConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsSettingCollectionResponse") {
		var out WindowsSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsSettingInstanceCollectionResponse") {
		var out WindowsSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUniversalAppXCollectionResponse") {
		var out WindowsUniversalAppXCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUniversalAppXCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUniversalAppXContainedAppCollectionResponse") {
		var out WindowsUniversalAppXContainedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUniversalAppXContainedAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdateForBusinessConfigurationCollectionResponse") {
		var out WindowsUpdateForBusinessConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdateForBusinessConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsWebAppCollectionResponse") {
		var out WindowsWebAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsWebAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartCollectionResponse") {
		var out WorkbookChartCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartPointCollectionResponse") {
		var out WorkbookChartPointCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartPointCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartSeriesCollectionResponse") {
		var out WorkbookChartSeriesCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartSeriesCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookCommentCollectionResponse") {
		var out WorkbookCommentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookCommentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookCommentReplyCollectionResponse") {
		var out WorkbookCommentReplyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookCommentReplyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookNamedItemCollectionResponse") {
		var out WorkbookNamedItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookNamedItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookOperationCollectionResponse") {
		var out WorkbookOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookPivotTableCollectionResponse") {
		var out WorkbookPivotTableCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookPivotTableCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeBorderCollectionResponse") {
		var out WorkbookRangeBorderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeBorderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeViewCollectionResponse") {
		var out WorkbookRangeViewCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeViewCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookSortFieldCollectionResponse") {
		var out WorkbookSortFieldCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookSortFieldCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTableCollectionResponse") {
		var out WorkbookTableCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTableCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTableColumnCollectionResponse") {
		var out WorkbookTableColumnCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTableColumnCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTableRowCollectionResponse") {
		var out WorkbookTableRowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTableRowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookWorksheetCollectionResponse") {
		var out WorkbookWorksheetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookWorksheetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workforceIntegrationCollectionResponse") {
		var out WorkforceIntegrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkforceIntegrationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateAuthenticationMethodConfigurationCollectionResponse") {
		var out X509CertificateAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateCombinationConfigurationCollectionResponse") {
		var out X509CertificateCombinationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateCombinationConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateRuleCollectionResponse") {
		var out X509CertificateRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateUserBindingCollectionResponse") {
		var out X509CertificateUserBindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateUserBindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	var parent BaseBaseCollectionPaginationCountResponseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBaseCollectionPaginationCountResponseImpl: %+v", err)
	}

	return RawBaseCollectionPaginationCountResponseImpl{
		baseCollectionPaginationCountResponse: parent,
		Type:                                  value,
		Values:                                temp,
	}, nil

}
