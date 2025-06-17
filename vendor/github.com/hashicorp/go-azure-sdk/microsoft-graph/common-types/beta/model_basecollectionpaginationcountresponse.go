package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
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

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentResourceRoleCollectionResponse") {
		var out AccessPackageAssignmentResourceRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentResourceRoleCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.accessPackageSubjectCollectionResponse") {
		var out AccessPackageSubjectCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageSubjectCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.accessReviewCollectionResponse") {
		var out AccessReviewCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewDecisionCollectionResponse") {
		var out AccessReviewDecisionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewDecisionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewErrorCollectionResponse") {
		var out AccessReviewErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewErrorCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.accountCollectionResponse") {
		var out AccountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aclCollectionResponse") {
		var out AclCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AclCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.actionItemCollectionResponse") {
		var out ActionItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActionItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.actionStepCollectionResponse") {
		var out ActionStepCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActionStepCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.actionUrlCollectionResponse") {
		var out ActionUrlCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActionUrlCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.activeDirectoryWindowsAutopilotDeploymentProfileCollectionResponse") {
		var out ActiveDirectoryWindowsAutopilotDeploymentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActiveDirectoryWindowsAutopilotDeploymentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.activeUsersMetricCollectionResponse") {
		var out ActiveUsersMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActiveUsersMetricCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.activityStatisticsCollectionResponse") {
		var out ActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityStatisticsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.advancedThreatProtectionOnboardingDeviceSettingStateCollectionResponse") {
		var out AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agedAccountsPayableCollectionResponse") {
		var out AgedAccountsPayableCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgedAccountsPayableCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agedAccountsReceivableCollectionResponse") {
		var out AgedAccountsReceivableCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgedAccountsReceivableCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionPluginCollectionResponse") {
		var out AiInteractionPluginCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionPluginCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiOnlineMeetingCollectionResponse") {
		var out AiOnlineMeetingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiOnlineMeetingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.airPrintDestinationCollectionResponse") {
		var out AirPrintDestinationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AirPrintDestinationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.allowedDataLocationCollectionResponse") {
		var out AllowedDataLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllowedDataLocationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidCertificateProfileBaseCollectionResponse") {
		var out AndroidCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidCertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerCertificateProfileBaseCollectionResponse") {
		var out AndroidDeviceOwnerCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerCertificateProfileBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerCompliancePolicyCollectionResponse") {
		var out AndroidDeviceOwnerCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerDelegatedScopeAppSettingCollectionResponse") {
		var out AndroidDeviceOwnerDelegatedScopeAppSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerDelegatedScopeAppSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerDerivedCredentialAuthenticationConfigurationCollectionResponse") {
		var out AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerEnrollmentProfileCollectionResponse") {
		var out AndroidDeviceOwnerEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerEnterpriseWiFiConfigurationCollectionResponse") {
		var out AndroidDeviceOwnerEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerGeneralDeviceConfigurationCollectionResponse") {
		var out AndroidDeviceOwnerGeneralDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerGeneralDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerImportedPFXCertificateProfileCollectionResponse") {
		var out AndroidDeviceOwnerImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerImportedPFXCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeAppPositionItemCollectionResponse") {
		var out AndroidDeviceOwnerKioskModeAppPositionItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeAppPositionItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeFolderItemCollectionResponse") {
		var out AndroidDeviceOwnerKioskModeFolderItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeFolderItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeManagedFolderCollectionResponse") {
		var out AndroidDeviceOwnerKioskModeManagedFolderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeManagedFolderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerPkcsCertificateProfileCollectionResponse") {
		var out AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerScepCertificateProfileCollectionResponse") {
		var out AndroidDeviceOwnerScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerScepCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerSilentCertificateAccessCollectionResponse") {
		var out AndroidDeviceOwnerSilentCertificateAccessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerSilentCertificateAccessCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerSystemUpdateFreezePeriodCollectionResponse") {
		var out AndroidDeviceOwnerSystemUpdateFreezePeriodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerSystemUpdateFreezePeriodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerTrustedRootCertificateCollectionResponse") {
		var out AndroidDeviceOwnerTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerVpnConfigurationCollectionResponse") {
		var out AndroidDeviceOwnerVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerWiFiConfigurationCollectionResponse") {
		var out AndroidDeviceOwnerWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidEasEmailProfileConfigurationCollectionResponse") {
		var out AndroidEasEmailProfileConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidEasEmailProfileConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidEnrollmentCompanyCodeCollectionResponse") {
		var out AndroidEnrollmentCompanyCodeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidEnrollmentCompanyCodeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidEnterpriseWiFiConfigurationCollectionResponse") {
		var out AndroidEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkAppCollectionResponse") {
		var out AndroidForWorkAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkAppConfigurationSchemaCollectionResponse") {
		var out AndroidForWorkAppConfigurationSchemaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkAppConfigurationSchemaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkAppConfigurationSchemaItemCollectionResponse") {
		var out AndroidForWorkAppConfigurationSchemaItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkAppConfigurationSchemaItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkCertificateProfileBaseCollectionResponse") {
		var out AndroidForWorkCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkCertificateProfileBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkCompliancePolicyCollectionResponse") {
		var out AndroidForWorkCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkCustomConfigurationCollectionResponse") {
		var out AndroidForWorkCustomConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkCustomConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkEasEmailProfileBaseCollectionResponse") {
		var out AndroidForWorkEasEmailProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkEasEmailProfileBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkEnrollmentProfileCollectionResponse") {
		var out AndroidForWorkEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkEnterpriseWiFiConfigurationCollectionResponse") {
		var out AndroidForWorkEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkGeneralDeviceConfigurationCollectionResponse") {
		var out AndroidForWorkGeneralDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkGeneralDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkGmailEasConfigurationCollectionResponse") {
		var out AndroidForWorkGmailEasConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkGmailEasConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkImportedPFXCertificateProfileCollectionResponse") {
		var out AndroidForWorkImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkImportedPFXCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkMobileAppConfigurationCollectionResponse") {
		var out AndroidForWorkMobileAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkMobileAppConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkNineWorkEasConfigurationCollectionResponse") {
		var out AndroidForWorkNineWorkEasConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkNineWorkEasConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkPkcsCertificateProfileCollectionResponse") {
		var out AndroidForWorkPkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkPkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkScepCertificateProfileCollectionResponse") {
		var out AndroidForWorkScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkScepCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkTrustedRootCertificateCollectionResponse") {
		var out AndroidForWorkTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkVpnConfigurationCollectionResponse") {
		var out AndroidForWorkVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkWiFiConfigurationCollectionResponse") {
		var out AndroidForWorkWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidFotaDeploymentAssignmentCollectionResponse") {
		var out AndroidFotaDeploymentAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidFotaDeploymentAssignmentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidImportedPFXCertificateProfileCollectionResponse") {
		var out AndroidImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidImportedPFXCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreAppCollectionResponse") {
		var out AndroidManagedStoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreAppConfigurationCollectionResponse") {
		var out AndroidManagedStoreAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreAppConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreAppConfigurationSchemaCollectionResponse") {
		var out AndroidManagedStoreAppConfigurationSchemaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreAppConfigurationSchemaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreAppConfigurationSchemaItemCollectionResponse") {
		var out AndroidManagedStoreAppConfigurationSchemaItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreAppConfigurationSchemaItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreAppTrackCollectionResponse") {
		var out AndroidManagedStoreAppTrackCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreAppTrackCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreWebAppCollectionResponse") {
		var out AndroidManagedStoreWebAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreWebAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidOmaCpConfigurationCollectionResponse") {
		var out AndroidOmaCpConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidOmaCpConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidPermissionActionCollectionResponse") {
		var out AndroidPermissionActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidPermissionActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidPkcsCertificateProfileCollectionResponse") {
		var out AndroidPkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidPkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidScepCertificateProfileCollectionResponse") {
		var out AndroidScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidScepCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidTrustedRootCertificateCollectionResponse") {
		var out AndroidTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidVpnConfigurationCollectionResponse") {
		var out AndroidVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWiFiConfigurationCollectionResponse") {
		var out AndroidWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileCertificateProfileBaseCollectionResponse") {
		var out AndroidWorkProfileCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileCertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileEasEmailProfileBaseCollectionResponse") {
		var out AndroidWorkProfileEasEmailProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileEasEmailProfileBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileEnterpriseWiFiConfigurationCollectionResponse") {
		var out AndroidWorkProfileEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileGmailEasConfigurationCollectionResponse") {
		var out AndroidWorkProfileGmailEasConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileGmailEasConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileNineWorkEasConfigurationCollectionResponse") {
		var out AndroidWorkProfileNineWorkEasConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileNineWorkEasConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfilePkcsCertificateProfileCollectionResponse") {
		var out AndroidWorkProfilePkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfilePkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileScepCertificateProfileCollectionResponse") {
		var out AndroidWorkProfileScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileScepCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileTrustedRootCertificateCollectionResponse") {
		var out AndroidWorkProfileTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileVpnConfigurationCollectionResponse") {
		var out AndroidWorkProfileVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileWiFiConfigurationCollectionResponse") {
		var out AndroidWorkProfileWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileWiFiConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerCertificateProfileBaseCollectionResponse") {
		var out AospDeviceOwnerCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerCertificateProfileBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerCompliancePolicyCollectionResponse") {
		var out AospDeviceOwnerCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerDeviceConfigurationCollectionResponse") {
		var out AospDeviceOwnerDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerEnterpriseWiFiConfigurationCollectionResponse") {
		var out AospDeviceOwnerEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerPkcsCertificateProfileCollectionResponse") {
		var out AospDeviceOwnerPkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerPkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerScepCertificateProfileCollectionResponse") {
		var out AospDeviceOwnerScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerScepCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerTrustedRootCertificateCollectionResponse") {
		var out AospDeviceOwnerTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aospDeviceOwnerWiFiConfigurationCollectionResponse") {
		var out AospDeviceOwnerWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AospDeviceOwnerWiFiConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.appCredentialSignInActivityCollectionResponse") {
		var out AppCredentialSignInActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppCredentialSignInActivityCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.appMetadataEntryCollectionResponse") {
		var out AppMetadataEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppMetadataEntryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.appScopeCollectionResponse") {
		var out AppScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appVulnerabilityManagedDeviceCollectionResponse") {
		var out AppVulnerabilityManagedDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppVulnerabilityManagedDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appVulnerabilityMobileAppCollectionResponse") {
		var out AppVulnerabilityMobileAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppVulnerabilityMobileAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appVulnerabilityTaskCollectionResponse") {
		var out AppVulnerabilityTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppVulnerabilityTaskCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.appleEnrollmentProfileAssignmentCollectionResponse") {
		var out AppleEnrollmentProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleEnrollmentProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleExpeditedCheckinConfigurationBaseCollectionResponse") {
		var out AppleExpeditedCheckinConfigurationBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleExpeditedCheckinConfigurationBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.appleOwnerTypeEnrollmentTypeCollectionResponse") {
		var out AppleOwnerTypeEnrollmentTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleOwnerTypeEnrollmentTypeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleUserInitiatedEnrollmentProfileCollectionResponse") {
		var out AppleUserInitiatedEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleUserInitiatedEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleVpnConfigurationCollectionResponse") {
		var out AppleVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleVppTokenTroubleshootingEventCollectionResponse") {
		var out AppleVppTokenTroubleshootingEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleVppTokenTroubleshootingEventCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.applicationSignInDetailedSummaryCollectionResponse") {
		var out ApplicationSignInDetailedSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationSignInDetailedSummaryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.appliedAuthenticationEventListenerCollectionResponse") {
		var out AppliedAuthenticationEventListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppliedAuthenticationEventListenerCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.approvalIdentitySetCollectionResponse") {
		var out ApprovalIdentitySetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalIdentitySetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalItemCollectionResponse") {
		var out ApprovalItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalItemRequestCollectionResponse") {
		var out ApprovalItemRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalItemRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalItemResponseCollectionResponse") {
		var out ApprovalItemResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalItemResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalOperationCollectionResponse") {
		var out ApprovalOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalOperationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.approvalStepCollectionResponse") {
		var out ApprovalStepCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalStepCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalWorkflowProviderCollectionResponse") {
		var out ApprovalWorkflowProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalWorkflowProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvedClientAppCollectionResponse") {
		var out ApprovedClientAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovedClientAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignedComputeInstanceDetailsCollectionResponse") {
		var out AssignedComputeInstanceDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignedComputeInstanceDetailsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.assignmentFilterEvaluationStatusDetailsCollectionResponse") {
		var out AssignmentFilterEvaluationStatusDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignmentFilterEvaluationStatusDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignmentFilterEvaluationSummaryCollectionResponse") {
		var out AssignmentFilterEvaluationSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignmentFilterEvaluationSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.assignmentFilterTypeAndEvaluationResultCollectionResponse") {
		var out AssignmentFilterTypeAndEvaluationResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssignmentFilterTypeAndEvaluationResultCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authenticationAppPolicyDetailsCollectionResponse") {
		var out AuthenticationAppPolicyDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationAppPolicyDetailsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authenticationContextCollectionResponse") {
		var out AuthenticationContextCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationContextCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationDetailCollectionResponse") {
		var out AuthenticationDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationDetailCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authenticationFailureCollectionResponse") {
		var out AuthenticationFailureCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationFailureCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationListenerCollectionResponse") {
		var out AuthenticationListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationListenerCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodDeviceCollectionResponse") {
		var out AuthenticationMethodDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodDeviceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authenticationRequirementPolicyCollectionResponse") {
		var out AuthenticationRequirementPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationRequirementPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authenticationsMetricCollectionResponse") {
		var out AuthenticationsMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationsMetricCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.authorizationSystemCollectionResponse") {
		var out AuthorizationSystemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthorizationSystemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authorizationSystemResourceCollectionResponse") {
		var out AuthorizationSystemResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthorizationSystemResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authorizationSystemTypeActionCollectionResponse") {
		var out AuthorizationSystemTypeActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthorizationSystemTypeActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authorizationSystemTypeServiceCollectionResponse") {
		var out AuthorizationSystemTypeServiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthorizationSystemTypeServiceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.awsAccessKeyCollectionResponse") {
		var out AwsAccessKeyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAccessKeyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAuthorizationSystemCollectionResponse") {
		var out AwsAuthorizationSystemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAuthorizationSystemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAuthorizationSystemResourceCollectionResponse") {
		var out AwsAuthorizationSystemResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAuthorizationSystemResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAuthorizationSystemTypeActionCollectionResponse") {
		var out AwsAuthorizationSystemTypeActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAuthorizationSystemTypeActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsEc2InstanceCollectionResponse") {
		var out AwsEc2InstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsEc2InstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsExternalSystemAccessFindingCollectionResponse") {
		var out AwsExternalSystemAccessFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsExternalSystemAccessFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsExternalSystemAccessRoleFindingCollectionResponse") {
		var out AwsExternalSystemAccessRoleFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsExternalSystemAccessRoleFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsGroupCollectionResponse") {
		var out AwsGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsIdentityAccessManagementKeyAgeFindingCollectionResponse") {
		var out AwsIdentityAccessManagementKeyAgeFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsIdentityAccessManagementKeyAgeFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsIdentityAccessManagementKeyUsageFindingCollectionResponse") {
		var out AwsIdentityAccessManagementKeyUsageFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsIdentityAccessManagementKeyUsageFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsIdentityCollectionResponse") {
		var out AwsIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsLambdaCollectionResponse") {
		var out AwsLambdaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsLambdaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsPolicyCollectionResponse") {
		var out AwsPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsRoleCollectionResponse") {
		var out AwsRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsSecretInformationAccessFindingCollectionResponse") {
		var out AwsSecretInformationAccessFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsSecretInformationAccessFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsSecurityToolAdministrationFindingCollectionResponse") {
		var out AwsSecurityToolAdministrationFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsSecurityToolAdministrationFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsStatementCollectionResponse") {
		var out AwsStatementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsStatementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsUserCollectionResponse") {
		var out AwsUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureADWindowsAutopilotDeploymentProfileCollectionResponse") {
		var out AzureADWindowsAutopilotDeploymentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureADWindowsAutopilotDeploymentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAuthorizationSystemCollectionResponse") {
		var out AzureAuthorizationSystemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAuthorizationSystemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAuthorizationSystemResourceCollectionResponse") {
		var out AzureAuthorizationSystemResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAuthorizationSystemResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAuthorizationSystemTypeActionCollectionResponse") {
		var out AzureAuthorizationSystemTypeActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAuthorizationSystemTypeActionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.azureGroupCollectionResponse") {
		var out AzureGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureIdentityCollectionResponse") {
		var out AzureIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureManagedIdentityCollectionResponse") {
		var out AzureManagedIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureManagedIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureRoleDefinitionCollectionResponse") {
		var out AzureRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureServerlessFunctionCollectionResponse") {
		var out AzureServerlessFunctionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureServerlessFunctionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureServicePrincipalCollectionResponse") {
		var out AzureServicePrincipalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureServicePrincipalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureUserCollectionResponse") {
		var out AzureUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.b2cIdentityUserFlowCollectionResponse") {
		var out B2cIdentityUserFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into B2cIdentityUserFlowCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.businessFlowCollectionResponse") {
		var out BusinessFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BusinessFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.businessFlowTemplateCollectionResponse") {
		var out BusinessFlowTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BusinessFlowTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.businessScenarioCollectionResponse") {
		var out BusinessScenarioCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BusinessScenarioCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.businessScenarioTaskCollectionResponse") {
		var out BusinessScenarioTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BusinessScenarioTaskCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.callActivityStatisticsCollectionResponse") {
		var out CallActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallActivityStatisticsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callAiInsightCollectionResponse") {
		var out CallAiInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallAiInsightCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cartToClassAssociationCollectionResponse") {
		var out CartToClassAssociationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CartToClassAssociationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateAuthorityAsEntityCollectionResponse") {
		var out CertificateAuthorityAsEntityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateAuthorityAsEntityCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.certificateBasedApplicationConfigurationCollectionResponse") {
		var out CertificateBasedApplicationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateBasedApplicationConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.certificateConnectorDetailsCollectionResponse") {
		var out CertificateConnectorDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateConnectorDetailsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.challengingWordCollectionResponse") {
		var out ChallengingWordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChallengingWordCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.chatActivityStatisticsCollectionResponse") {
		var out ChatActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatActivityStatisticsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.chromeOSDevicePropertyCollectionResponse") {
		var out ChromeOSDevicePropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChromeOSDevicePropertyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chromeOSOnboardingSettingsCollectionResponse") {
		var out ChromeOSOnboardingSettingsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChromeOSOnboardingSettingsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.classifcationErrorBaseCollectionResponse") {
		var out ClassifcationErrorBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassifcationErrorBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.classificationAttributeCollectionResponse") {
		var out ClassificationAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassificationAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.classificationErrorCollectionResponse") {
		var out ClassificationErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassificationErrorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.classificationJobResponseCollectionResponse") {
		var out ClassificationJobResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassificationJobResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudAppSecurityProfileCollectionResponse") {
		var out CloudAppSecurityProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudAppSecurityProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudCertificationAuthorityCollectionResponse") {
		var out CloudCertificationAuthorityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudCertificationAuthorityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudCertificationAuthorityLeafCertificateCollectionResponse") {
		var out CloudCertificationAuthorityLeafCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudCertificationAuthorityLeafCertificateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudLicensing.serviceCollectionResponse") {
		var out CloudLicensingServiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudLicensingServiceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudLicensing.usageRightCollectionResponse") {
		var out CloudLicensingUsageRightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudLicensingUsageRightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcAgentHealthCheckDetailCollectionResponse") {
		var out CloudPCAgentHealthCheckDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCAgentHealthCheckDetailCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkActionCollectionResponse") {
		var out CloudPCBulkActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkCreateSnapshotCollectionResponse") {
		var out CloudPCBulkCreateSnapshotCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkCreateSnapshotCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkDisasterRecoveryCollectionResponse") {
		var out CloudPCBulkDisasterRecoveryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkDisasterRecoveryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkDisasterRecoveryFailbackCollectionResponse") {
		var out CloudPCBulkDisasterRecoveryFailbackCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkDisasterRecoveryFailbackCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkDisasterRecoveryFailoverCollectionResponse") {
		var out CloudPCBulkDisasterRecoveryFailoverCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkDisasterRecoveryFailoverCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkModifyDiskEncryptionTypeCollectionResponse") {
		var out CloudPCBulkModifyDiskEncryptionTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkModifyDiskEncryptionTypeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkMoveCollectionResponse") {
		var out CloudPCBulkMoveCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkMoveCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkPowerOffCollectionResponse") {
		var out CloudPCBulkPowerOffCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkPowerOffCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkPowerOnCollectionResponse") {
		var out CloudPCBulkPowerOnCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkPowerOnCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkReinstallAgentCollectionResponse") {
		var out CloudPCBulkReinstallAgentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkReinstallAgentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkReprovisionCollectionResponse") {
		var out CloudPCBulkReprovisionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkReprovisionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkResizeCollectionResponse") {
		var out CloudPCBulkResizeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkResizeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkRestartCollectionResponse") {
		var out CloudPCBulkRestartCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkRestartCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkRestoreCollectionResponse") {
		var out CloudPCBulkRestoreCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkRestoreCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkSetReviewStatusCollectionResponse") {
		var out CloudPCBulkSetReviewStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkSetReviewStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkTroubleshootCollectionResponse") {
		var out CloudPCBulkTroubleshootCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkTroubleshootCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudPCConnectivityIssueCollectionResponse") {
		var out CloudPCConnectivityIssueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCConnectivityIssueCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudPcExportJobCollectionResponse") {
		var out CloudPCExportJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCExportJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcExternalPartnerSettingCollectionResponse") {
		var out CloudPCExternalPartnerSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCExternalPartnerSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcFrontLineServicePlanCollectionResponse") {
		var out CloudPCFrontLineServicePlanCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCFrontLineServicePlanCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudPcHealthCheckItemCollectionResponse") {
		var out CloudPCHealthCheckItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCHealthCheckItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudPcPartnerAgentInstallResultCollectionResponse") {
		var out CloudPCPartnerAgentInstallResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCPartnerAgentInstallResultCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.cloudPcRemoteActionResultCollectionResponse") {
		var out CloudPCRemoteActionResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCRemoteActionResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcServicePlanCollectionResponse") {
		var out CloudPCServicePlanCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCServicePlanCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcSnapshotCollectionResponse") {
		var out CloudPCSnapshotCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCSnapshotCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcSupportedRegionCollectionResponse") {
		var out CloudPCSupportedRegionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCSupportedRegionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.comanagementEligibleDeviceCollectionResponse") {
		var out ComanagementEligibleDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ComanagementEligibleDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.commandCollectionResponse") {
		var out CommandCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommandCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.companyCollectionResponse") {
		var out CompanyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CompanyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.companyInformationCollectionResponse") {
		var out CompanyInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CompanyInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.companyPortalBlockedActionCollectionResponse") {
		var out CompanyPortalBlockedActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CompanyPortalBlockedActionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessRuleSatisfiedCollectionResponse") {
		var out ConditionalAccessRuleSatisfiedCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessRuleSatisfiedCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.configManagerCollectionCollectionResponse") {
		var out ConfigManagerCollectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConfigManagerCollectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.configurationUriCollectionResponse") {
		var out ConfigurationUriCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConfigurationUriCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.connectionOperationCollectionResponse") {
		var out ConnectionOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectionOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.connectorCollectionResponse") {
		var out ConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.connectorGroupCollectionResponse") {
		var out ConnectorGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectorGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.connectorStatusDetailsCollectionResponse") {
		var out ConnectorStatusDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectorStatusDetailsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.contentActivityCollectionResponse") {
		var out ContentActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentClassificationCollectionResponse") {
		var out ContentClassificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentClassificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentModelCollectionResponse") {
		var out ContentModelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentModelCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.corsConfigurationCollectionResponse") {
		var out CorsConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CorsConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.corsConfiguration_v2CollectionResponse") {
		var out CorsConfigurationv2CollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CorsConfigurationv2CollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.countryRegionCollectionResponse") {
		var out CountryRegionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CountryRegionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.credentialCollectionResponse") {
		var out CredentialCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CredentialCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.credentialUserRegistrationDetailsCollectionResponse") {
		var out CredentialUserRegistrationDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CredentialUserRegistrationDetailsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.currencyCollectionResponse") {
		var out CurrencyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CurrencyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customAccessPackageWorkflowExtensionCollectionResponse") {
		var out CustomAccessPackageWorkflowExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAccessPackageWorkflowExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customAppScopeCollectionResponse") {
		var out CustomAppScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAppScopeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.customClaimBaseCollectionResponse") {
		var out CustomClaimBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomClaimBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customClaimConfigurationCollectionResponse") {
		var out CustomClaimConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomClaimConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customClaimTransformationCollectionResponse") {
		var out CustomClaimTransformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomClaimTransformationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.customExtensionHandlerCollectionResponse") {
		var out CustomExtensionHandlerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomExtensionHandlerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customExtensionHandlerInstanceCollectionResponse") {
		var out CustomExtensionHandlerInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomExtensionHandlerInstanceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.customQuestionAnswerCollectionResponse") {
		var out CustomQuestionAnswerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomQuestionAnswerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customSecurityAttributeAuditCollectionResponse") {
		var out CustomSecurityAttributeAuditCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSecurityAttributeAuditCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.customSecurityAttributeExemptionCollectionResponse") {
		var out CustomSecurityAttributeExemptionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSecurityAttributeExemptionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customSecurityAttributeStringValueExemptionCollectionResponse") {
		var out CustomSecurityAttributeStringValueExemptionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSecurityAttributeStringValueExemptionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customSubjectAlternativeNameCollectionResponse") {
		var out CustomSubjectAlternativeNameCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSubjectAlternativeNameCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customUpdateTimeWindowCollectionResponse") {
		var out CustomUpdateTimeWindowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomUpdateTimeWindowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customerCollectionResponse") {
		var out CustomerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customerPaymentCollectionResponse") {
		var out CustomerPaymentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomerPaymentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customerPaymentJournalCollectionResponse") {
		var out CustomerPaymentJournalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomerPaymentJournalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dailyInactiveUsersByApplicationMetricCollectionResponse") {
		var out DailyInactiveUsersByApplicationMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DailyInactiveUsersByApplicationMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dailyInactiveUsersMetricCollectionResponse") {
		var out DailyInactiveUsersMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DailyInactiveUsersMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dataLossPreventionPolicyCollectionResponse") {
		var out DataLossPreventionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataLossPreventionPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.dataSharingConsentCollectionResponse") {
		var out DataSharingConsentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataSharingConsentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.defaultDeviceCompliancePolicyCollectionResponse") {
		var out DefaultDeviceCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefaultDeviceCompliancePolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.defaultUserRoleOverrideCollectionResponse") {
		var out DefaultUserRoleOverrideCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefaultUserRoleOverrideCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.delegationSettingsCollectionResponse") {
		var out DelegationSettingsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegationSettingsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.depEnrollmentBaseProfileCollectionResponse") {
		var out DepEnrollmentBaseProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepEnrollmentBaseProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depEnrollmentProfileCollectionResponse") {
		var out DepEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depIOSEnrollmentProfileCollectionResponse") {
		var out DepIOSEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepIOSEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depMacOSEnrollmentProfileCollectionResponse") {
		var out DepMacOSEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepMacOSEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depOnboardingSettingCollectionResponse") {
		var out DepOnboardingSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepOnboardingSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depTvOSEnrollmentProfileCollectionResponse") {
		var out DepTvOSEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepTvOSEnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depVisionOSEnrollmentProfileCollectionResponse") {
		var out DepVisionOSEnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepVisionOSEnrollmentProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.detectedSensitiveContentCollectionResponse") {
		var out DetectedSensitiveContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DetectedSensitiveContentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementAssignedRoleDefinitionCollectionResponse") {
		var out DeviceAndAppManagementAssignedRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementAssignedRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementAssignmentFilterCollectionResponse") {
		var out DeviceAndAppManagementAssignmentFilterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementAssignmentFilterCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceAppManagementTaskCollectionResponse") {
		var out DeviceAppManagementTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAppManagementTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAssignmentItemCollectionResponse") {
		var out DeviceAssignmentItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAssignmentItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceComanagementAuthorityConfigurationCollectionResponse") {
		var out DeviceComanagementAuthorityConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComanagementAuthorityConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyPolicySetItemCollectionResponse") {
		var out DeviceCompliancePolicyPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScriptCollectionResponse") {
		var out DeviceComplianceScriptCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScriptCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScriptDeviceStateCollectionResponse") {
		var out DeviceComplianceScriptDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScriptDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScriptErrorCollectionResponse") {
		var out DeviceComplianceScriptErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScriptErrorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScriptRuleCollectionResponse") {
		var out DeviceComplianceScriptRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScriptRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScriptRuleErrorCollectionResponse") {
		var out DeviceComplianceScriptRuleErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScriptRuleErrorCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationConflictSummaryCollectionResponse") {
		var out DeviceConfigurationConflictSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationConflictSummaryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationGroupAssignmentCollectionResponse") {
		var out DeviceConfigurationGroupAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationGroupAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationPolicySetItemCollectionResponse") {
		var out DeviceConfigurationPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceCustomAttributeShellScriptCollectionResponse") {
		var out DeviceCustomAttributeShellScriptCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCustomAttributeShellScriptCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentNotificationConfigurationCollectionResponse") {
		var out DeviceEnrollmentNotificationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentNotificationConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentPlatformRestrictionConfigurationCollectionResponse") {
		var out DeviceEnrollmentPlatformRestrictionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentPlatformRestrictionConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptAssignmentCollectionResponse") {
		var out DeviceHealthScriptAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptCollectionResponse") {
		var out DeviceHealthScriptCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptDeviceStateCollectionResponse") {
		var out DeviceHealthScriptDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptParameterCollectionResponse") {
		var out DeviceHealthScriptParameterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptParameterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptPolicyStateCollectionResponse") {
		var out DeviceHealthScriptPolicyStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptPolicyStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptRemediationHistoryDataCollectionResponse") {
		var out DeviceHealthScriptRemediationHistoryDataCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptRemediationHistoryDataCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceKeyCollectionResponse") {
		var out DeviceKeyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceKeyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementAbstractComplexSettingDefinitionCollectionResponse") {
		var out DeviceManagementAbstractComplexSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAbstractComplexSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementAbstractComplexSettingInstanceCollectionResponse") {
		var out DeviceManagementAbstractComplexSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAbstractComplexSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagement.alertRecordCollectionResponse") {
		var out DeviceManagementAlertRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAlertRecordCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagement.alertRuleCollectionResponse") {
		var out DeviceManagementAlertRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAlertRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementAutopilotEventCollectionResponse") {
		var out DeviceManagementAutopilotEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAutopilotEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementBooleanSettingInstanceCollectionResponse") {
		var out DeviceManagementBooleanSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementBooleanSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCachedReportConfigurationCollectionResponse") {
		var out DeviceManagementCachedReportConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCachedReportConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCollectionSettingDefinitionCollectionResponse") {
		var out DeviceManagementCollectionSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCollectionSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCollectionSettingInstanceCollectionResponse") {
		var out DeviceManagementCollectionSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCollectionSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementComplexSettingDefinitionCollectionResponse") {
		var out DeviceManagementComplexSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementComplexSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementComplexSettingInstanceCollectionResponse") {
		var out DeviceManagementComplexSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementComplexSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementComplianceActionItemCollectionResponse") {
		var out DeviceManagementComplianceActionItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementComplianceActionItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCompliancePolicyCollectionResponse") {
		var out DeviceManagementCompliancePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCompliancePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementComplianceScheduledActionForRuleCollectionResponse") {
		var out DeviceManagementComplianceScheduledActionForRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementComplianceScheduledActionForRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationCategoryCollectionResponse") {
		var out DeviceManagementConfigurationCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationChoiceSettingCollectionDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingCollectionDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationChoiceSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingValueCollectionResponse") {
		var out DeviceManagementConfigurationChoiceSettingValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingValueTemplateCollectionResponse") {
		var out DeviceManagementConfigurationChoiceSettingValueTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingValueTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationDependentOnCollectionResponse") {
		var out DeviceManagementConfigurationDependentOnCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationDependentOnCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingValueCollectionResponse") {
		var out DeviceManagementConfigurationGroupSettingValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingValueTemplateCollectionResponse") {
		var out DeviceManagementConfigurationGroupSettingValueTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingValueTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationOptionDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationOptionDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationOptionDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationOptionDefinitionTemplateCollectionResponse") {
		var out DeviceManagementConfigurationOptionDefinitionTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationOptionDefinitionTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationPolicyAssignmentCollectionResponse") {
		var out DeviceManagementConfigurationPolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationPolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationPolicyCollectionResponse") {
		var out DeviceManagementConfigurationPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationPolicyPolicySetItemCollectionResponse") {
		var out DeviceManagementConfigurationPolicyPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationPolicyPolicySetItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationPolicyTemplateCollectionResponse") {
		var out DeviceManagementConfigurationPolicyTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationPolicyTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationRedirectSettingDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationRedirectSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationRedirectSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationReferredSettingInformationCollectionResponse") {
		var out DeviceManagementConfigurationReferredSettingInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationReferredSettingInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingCollectionResponse") {
		var out DeviceManagementConfigurationSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingDependedOnByCollectionResponse") {
		var out DeviceManagementConfigurationSettingDependedOnByCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingDependedOnByCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationSettingGroupCollectionDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingGroupCollectionDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingGroupDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationSettingGroupDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingGroupDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingInstanceCollectionResponse") {
		var out DeviceManagementConfigurationSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingInstanceTemplateCollectionResponse") {
		var out DeviceManagementConfigurationSettingInstanceTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingInstanceTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingTemplateCollectionResponse") {
		var out DeviceManagementConfigurationSettingTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationSimpleSettingCollectionDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingCollectionDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingDefinitionCollectionResponse") {
		var out DeviceManagementConfigurationSimpleSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingValueCollectionResponse") {
		var out DeviceManagementConfigurationSimpleSettingValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingValueTemplateCollectionResponse") {
		var out DeviceManagementConfigurationSimpleSettingValueTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingValueTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConstraintCollectionResponse") {
		var out DeviceManagementConstraintCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConstraintCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementDerivedCredentialSettingsCollectionResponse") {
		var out DeviceManagementDerivedCredentialSettingsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementDerivedCredentialSettingsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementDomainJoinConnectorCollectionResponse") {
		var out DeviceManagementDomainJoinConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementDomainJoinConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementEnumValueCollectionResponse") {
		var out DeviceManagementEnumValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementEnumValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExchangeAccessRuleCollectionResponse") {
		var out DeviceManagementExchangeAccessRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExchangeAccessRuleCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExchangeDeviceClassCollectionResponse") {
		var out DeviceManagementExchangeDeviceClassCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExchangeDeviceClassCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExchangeOnPremisesPolicyCollectionResponse") {
		var out DeviceManagementExchangeOnPremisesPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExchangeOnPremisesPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntegerSettingInstanceCollectionResponse") {
		var out DeviceManagementIntegerSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntegerSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentAssignmentCollectionResponse") {
		var out DeviceManagementIntentAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentCollectionResponse") {
		var out DeviceManagementIntentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentDeviceSettingStateSummaryCollectionResponse") {
		var out DeviceManagementIntentDeviceSettingStateSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentDeviceSettingStateSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentDeviceStateCollectionResponse") {
		var out DeviceManagementIntentDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentSettingCategoryCollectionResponse") {
		var out DeviceManagementIntentSettingCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentSettingCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentUserStateCollectionResponse") {
		var out DeviceManagementIntentUserStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentUserStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagement.notificationChannelCollectionResponse") {
		var out DeviceManagementNotificationChannelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementNotificationChannelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagement.notificationReceiverCollectionResponse") {
		var out DeviceManagementNotificationReceiverCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementNotificationReceiverCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementResourceAccessProfileAssignmentCollectionResponse") {
		var out DeviceManagementResourceAccessProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementResourceAccessProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementResourceAccessProfileBaseCollectionResponse") {
		var out DeviceManagementResourceAccessProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementResourceAccessProfileBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementReusablePolicySettingCollectionResponse") {
		var out DeviceManagementReusablePolicySettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementReusablePolicySettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagement.ruleConditionCollectionResponse") {
		var out DeviceManagementRuleConditionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementRuleConditionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptAssignmentCollectionResponse") {
		var out DeviceManagementScriptAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptCollectionResponse") {
		var out DeviceManagementScriptCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptDeviceStateCollectionResponse") {
		var out DeviceManagementScriptDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptGroupAssignmentCollectionResponse") {
		var out DeviceManagementScriptGroupAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptGroupAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptPolicySetItemCollectionResponse") {
		var out DeviceManagementScriptPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptPolicySetItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptUserStateCollectionResponse") {
		var out DeviceManagementScriptUserStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptUserStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingCategoryCollectionResponse") {
		var out DeviceManagementSettingCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingDefinitionCollectionResponse") {
		var out DeviceManagementSettingDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingDependencyCollectionResponse") {
		var out DeviceManagementSettingDependencyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingDependencyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingInsightsDefinitionCollectionResponse") {
		var out DeviceManagementSettingInsightsDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingInsightsDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingInstanceCollectionResponse") {
		var out DeviceManagementSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementStringSettingInstanceCollectionResponse") {
		var out DeviceManagementStringSettingInstanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementStringSettingInstanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTemplateCollectionResponse") {
		var out DeviceManagementTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTemplateInsightsDefinitionCollectionResponse") {
		var out DeviceManagementTemplateInsightsDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTemplateInsightsDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTemplateSettingCategoryCollectionResponse") {
		var out DeviceManagementTemplateSettingCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTemplateSettingCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTroubleshootingErrorResourceCollectionResponse") {
		var out DeviceManagementTroubleshootingErrorResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTroubleshootingErrorResourceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementUserRightsLocalUserOrGroupCollectionResponse") {
		var out DeviceManagementUserRightsLocalUserOrGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementUserRightsLocalUserOrGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceShellScriptCollectionResponse") {
		var out DeviceShellScriptCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceShellScriptCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceTemplateCollectionResponse") {
		var out DeviceTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dimensionCollectionResponse") {
		var out DimensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DimensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dimensionValueCollectionResponse") {
		var out DimensionValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DimensionValueCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.directorySettingCollectionResponse") {
		var out DirectorySettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectorySettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directorySettingTemplateCollectionResponse") {
		var out DirectorySettingTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectorySettingTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.discoveredSensitiveTypeCollectionResponse") {
		var out DiscoveredSensitiveTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DiscoveredSensitiveTypeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.dlpActionInfoCollectionResponse") {
		var out DlpActionInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DlpActionInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dlpEvaluatePoliciesJobResponseCollectionResponse") {
		var out DlpEvaluatePoliciesJobResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DlpEvaluatePoliciesJobResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.documentCommentCollectionResponse") {
		var out DocumentCommentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentCommentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.documentCommentReplyCollectionResponse") {
		var out DocumentCommentReplyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentCommentReplyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.documentProcessingJobCollectionResponse") {
		var out DocumentProcessingJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentProcessingJobCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.domainSecurityProfileCollectionResponse") {
		var out DomainSecurityProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainSecurityProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.easEmailProfileConfigurationBaseCollectionResponse") {
		var out EasEmailProfileConfigurationBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EasEmailProfileConfigurationBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.addToReviewSetOperationCollectionResponse") {
		var out EdiscoveryAddToReviewSetOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryAddToReviewSetOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseCollectionResponse") {
		var out EdiscoveryCaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseExportOperationCollectionResponse") {
		var out EdiscoveryCaseExportOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseExportOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseHoldOperationCollectionResponse") {
		var out EdiscoveryCaseHoldOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseHoldOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseIndexOperationCollectionResponse") {
		var out EdiscoveryCaseIndexOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseIndexOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseOperationCollectionResponse") {
		var out EdiscoveryCaseOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.custodianCollectionResponse") {
		var out EdiscoveryCustodianCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCustodianCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.dataSourceCollectionResponse") {
		var out EdiscoveryDataSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryDataSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.estimateStatisticsOperationCollectionResponse") {
		var out EdiscoveryEstimateStatisticsOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryEstimateStatisticsOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.legalHoldCollectionResponse") {
		var out EdiscoveryLegalHoldCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryLegalHoldCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.noncustodialDataSourceCollectionResponse") {
		var out EdiscoveryNoncustodialDataSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryNoncustodialDataSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.purgeDataOperationCollectionResponse") {
		var out EdiscoveryPurgeDataOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryPurgeDataOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.reviewSetCollectionResponse") {
		var out EdiscoveryReviewSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryReviewSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.reviewSetQueryCollectionResponse") {
		var out EdiscoveryReviewSetQueryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryReviewSetQueryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.siteSourceCollectionResponse") {
		var out EdiscoverySiteSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoverySiteSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.sourceCollectionCollectionResponse") {
		var out EdiscoverySourceCollectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoverySourceCollectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.tagCollectionResponse") {
		var out EdiscoveryTagCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryTagCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.tagOperationCollectionResponse") {
		var out EdiscoveryTagOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryTagOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.unifiedGroupSourceCollectionResponse") {
		var out EdiscoveryUnifiedGroupSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryUnifiedGroupSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.userSourceCollectionResponse") {
		var out EdiscoveryUserSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryUserSourceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.educationGradingSchemeCollectionResponse") {
		var out EducationGradingSchemeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationGradingSchemeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationGradingSchemeGradeCollectionResponse") {
		var out EducationGradingSchemeGradeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationGradingSchemeGradeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.educationalActivityCollectionResponse") {
		var out EducationalActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationalActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailActivityStatisticsCollectionResponse") {
		var out EmailActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailActivityStatisticsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.embeddedSIMActivationCodeCollectionResponse") {
		var out EmbeddedSIMActivationCodeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmbeddedSIMActivationCodeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.embeddedSIMActivationCodePoolAssignmentCollectionResponse") {
		var out EmbeddedSIMActivationCodePoolAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmbeddedSIMActivationCodePoolAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.embeddedSIMActivationCodePoolCollectionResponse") {
		var out EmbeddedSIMActivationCodePoolCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmbeddedSIMActivationCodePoolCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.embeddedSIMDeviceStateCollectionResponse") {
		var out EmbeddedSIMDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmbeddedSIMDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.employeeCollectionResponse") {
		var out EmployeeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmployeeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptedAwsStorageBucketFindingCollectionResponse") {
		var out EncryptedAwsStorageBucketFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptedAwsStorageBucketFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptedAzureStorageAccountFindingCollectionResponse") {
		var out EncryptedAzureStorageAccountFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptedAzureStorageAccountFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptedGcpStorageBucketFindingCollectionResponse") {
		var out EncryptedGcpStorageBucketFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptedGcpStorageBucketFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptionReportPolicyDetailsCollectionResponse") {
		var out EncryptionReportPolicyDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptionReportPolicyDetailsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.engagementRoleCollectionResponse") {
		var out EngagementRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EngagementRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.engagementRoleMemberCollectionResponse") {
		var out EngagementRoleMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EngagementRoleMemberCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.enrollmentProfileCollectionResponse") {
		var out EnrollmentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentRestrictionsConfigurationPolicySetItemCollectionResponse") {
		var out EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentTimeDeviceMembershipTargetStatusCollectionResponse") {
		var out EnrollmentTimeDeviceMembershipTargetStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentTimeDeviceMembershipTargetStatusCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.enterpriseCodeSigningCertificateCollectionResponse") {
		var out EnterpriseCodeSigningCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnterpriseCodeSigningCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.evaluateLabelJobResponseCollectionResponse") {
		var out EvaluateLabelJobResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EvaluateLabelJobResponseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.exactDataMatchStoreColumnCollectionResponse") {
		var out ExactDataMatchStoreColumnCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactDataMatchStoreColumnCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchDataStoreCollectionResponse") {
		var out ExactMatchDataStoreCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchDataStoreCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchDetectedSensitiveContentCollectionResponse") {
		var out ExactMatchDetectedSensitiveContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchDetectedSensitiveContentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchSessionCollectionResponse") {
		var out ExactMatchSessionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchSessionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchUploadAgentCollectionResponse") {
		var out ExactMatchUploadAgentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchUploadAgentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.expressionEvaluationDetailsCollectionResponse") {
		var out ExpressionEvaluationDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExpressionEvaluationDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extendedKeyUsageCollectionResponse") {
		var out ExtendedKeyUsageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtendedKeyUsageCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.externalAuthenticationMethodConfigurationCollectionResponse") {
		var out ExternalAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectionCollectionResponse") {
		var out ExternalConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.externalGroupCollectionResponse") {
		var out ExternalGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalIdentitiesPolicyCollectionResponse") {
		var out ExternalIdentitiesPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalIdentitiesPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalItemCollectionResponse") {
		var out ExternalItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalMeetingRegistrantCollectionResponse") {
		var out ExternalMeetingRegistrantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalMeetingRegistrantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalProfileCollectionResponse") {
		var out ExternalProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalUserProfileCollectionResponse") {
		var out ExternalUserProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalUserProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.externallyAccessibleAwsStorageBucketFindingCollectionResponse") {
		var out ExternallyAccessibleAwsStorageBucketFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternallyAccessibleAwsStorageBucketFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externallyAccessibleAzureBlobContainerFindingCollectionResponse") {
		var out ExternallyAccessibleAzureBlobContainerFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternallyAccessibleAzureBlobContainerFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externallyAccessibleGcpStorageBucketFindingCollectionResponse") {
		var out ExternallyAccessibleGcpStorageBucketFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternallyAccessibleGcpStorageBucketFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.federatedTokenValidationPolicyCollectionResponse") {
		var out FederatedTokenValidationPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FederatedTokenValidationPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.fileHashCollectionResponse") {
		var out FileHashCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileHashCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileSecurityProfileCollectionResponse") {
		var out FileSecurityProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileSecurityProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.findingCollectionResponse") {
		var out FindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.focusActivityStatisticsCollectionResponse") {
		var out FocusActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FocusActivityStatisticsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpAuthorizationSystemCollectionResponse") {
		var out GcpAuthorizationSystemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpAuthorizationSystemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpAuthorizationSystemResourceCollectionResponse") {
		var out GcpAuthorizationSystemResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpAuthorizationSystemResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpAuthorizationSystemTypeActionCollectionResponse") {
		var out GcpAuthorizationSystemTypeActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpAuthorizationSystemTypeActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpCloudFunctionCollectionResponse") {
		var out GcpCloudFunctionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpCloudFunctionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpGroupCollectionResponse") {
		var out GcpGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpIdentityCollectionResponse") {
		var out GcpIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpRoleCollectionResponse") {
		var out GcpRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpScopeCollectionResponse") {
		var out GcpScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpScopeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpServiceAccountCollectionResponse") {
		var out GcpServiceAccountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpServiceAccountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpUserCollectionResponse") {
		var out GcpUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpUserCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.generalLedgerEntryCollectionResponse") {
		var out GeneralLedgerEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GeneralLedgerEntryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.goalsExportJobCollectionResponse") {
		var out GoalsExportJobCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GoalsExportJobCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceCriteriaCollectionResponse") {
		var out GovernanceCriteriaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceCriteriaCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.governanceNotificationTemplateCollectionResponse") {
		var out GovernanceNotificationTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceNotificationTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governancePolicyTemplateCollectionResponse") {
		var out GovernancePolicyTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernancePolicyTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceResourceCollectionResponse") {
		var out GovernanceResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceRoleAssignmentCollectionResponse") {
		var out GovernanceRoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceRoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceRoleAssignmentRequestCollectionResponse") {
		var out GovernanceRoleAssignmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceRoleAssignmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceRoleDefinitionCollectionResponse") {
		var out GovernanceRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceRoleSettingCollectionResponse") {
		var out GovernanceRoleSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceRoleSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceRuleSettingCollectionResponse") {
		var out GovernanceRuleSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceRuleSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceSubjectCollectionResponse") {
		var out GovernanceSubjectCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceSubjectCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyCategoryCollectionResponse") {
		var out GroupPolicyCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyConfigurationAssignmentCollectionResponse") {
		var out GroupPolicyConfigurationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyConfigurationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyConfigurationCollectionResponse") {
		var out GroupPolicyConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyDefinitionCollectionResponse") {
		var out GroupPolicyDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyDefinitionFileCollectionResponse") {
		var out GroupPolicyDefinitionFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyDefinitionFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyDefinitionValueCollectionResponse") {
		var out GroupPolicyDefinitionValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyDefinitionValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyMigrationReportCollectionResponse") {
		var out GroupPolicyMigrationReportCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyMigrationReportCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyObjectFileCollectionResponse") {
		var out GroupPolicyObjectFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyObjectFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyOperationCollectionResponse") {
		var out GroupPolicyOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationCheckBoxCollectionResponse") {
		var out GroupPolicyPresentationCheckBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationCheckBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationCollectionResponse") {
		var out GroupPolicyPresentationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationComboBoxCollectionResponse") {
		var out GroupPolicyPresentationComboBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationComboBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationDecimalTextBoxCollectionResponse") {
		var out GroupPolicyPresentationDecimalTextBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationDecimalTextBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationDropdownListCollectionResponse") {
		var out GroupPolicyPresentationDropdownListCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationDropdownListCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationDropdownListItemCollectionResponse") {
		var out GroupPolicyPresentationDropdownListItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationDropdownListItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationListBoxCollectionResponse") {
		var out GroupPolicyPresentationListBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationListBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationLongDecimalTextBoxCollectionResponse") {
		var out GroupPolicyPresentationLongDecimalTextBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationLongDecimalTextBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationMultiTextBoxCollectionResponse") {
		var out GroupPolicyPresentationMultiTextBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationMultiTextBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationTextBoxCollectionResponse") {
		var out GroupPolicyPresentationTextBoxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationTextBoxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationTextCollectionResponse") {
		var out GroupPolicyPresentationTextCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationTextCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueBooleanCollectionResponse") {
		var out GroupPolicyPresentationValueBooleanCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueBooleanCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueCollectionResponse") {
		var out GroupPolicyPresentationValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueDecimalCollectionResponse") {
		var out GroupPolicyPresentationValueDecimalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueDecimalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueListCollectionResponse") {
		var out GroupPolicyPresentationValueListCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueListCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueLongDecimalCollectionResponse") {
		var out GroupPolicyPresentationValueLongDecimalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueLongDecimalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueMultiTextCollectionResponse") {
		var out GroupPolicyPresentationValueMultiTextCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueMultiTextCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueTextCollectionResponse") {
		var out GroupPolicyPresentationValueTextCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueTextCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicySettingMappingCollectionResponse") {
		var out GroupPolicySettingMappingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicySettingMappingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyUploadedDefinitionFileCollectionResponse") {
		var out GroupPolicyUploadedDefinitionFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyUploadedDefinitionFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyUploadedLanguageFileCollectionResponse") {
		var out GroupPolicyUploadedLanguageFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyUploadedLanguageFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyUploadedPresentationCollectionResponse") {
		var out GroupPolicyUploadedPresentationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyUploadedPresentationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareConfigurationAssignmentCollectionResponse") {
		var out HardwareConfigurationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareConfigurationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareConfigurationCollectionResponse") {
		var out HardwareConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareConfigurationDeviceStateCollectionResponse") {
		var out HardwareConfigurationDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareConfigurationDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareConfigurationUserStateCollectionResponse") {
		var out HardwareConfigurationUserStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareConfigurationUserStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareOathAuthenticationMethodCollectionResponse") {
		var out HardwareOathAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareOathAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareOathAuthenticationMethodConfigurationCollectionResponse") {
		var out HardwareOathAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareOathAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareOathTokenAuthenticationMethodDeviceCollectionResponse") {
		var out HardwareOathTokenAuthenticationMethodDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareOathTokenAuthenticationMethodDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwarePasswordDetailCollectionResponse") {
		var out HardwarePasswordDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwarePasswordDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwarePasswordInfoCollectionResponse") {
		var out HardwarePasswordInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwarePasswordInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.alertCollectionResponse") {
		var out HealthMonitoringAlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringAlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.alertConfigurationCollectionResponse") {
		var out HealthMonitoringAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.emailNotificationConfigurationCollectionResponse") {
		var out HealthMonitoringEmailNotificationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringEmailNotificationConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.resourceImpactSummaryCollectionResponse") {
		var out HealthMonitoringResourceImpactSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringResourceImpactSummaryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.hostSecurityProfileCollectionResponse") {
		var out HostSecurityProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HostSecurityProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.ipApplicationSegmentCollectionResponse") {
		var out IPApplicationSegmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPApplicationSegmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ipCategoryCollectionResponse") {
		var out IPCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPCategoryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.ipReferenceDataCollectionResponse") {
		var out IPReferenceDataCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPReferenceDataCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ipSecurityProfileCollectionResponse") {
		var out IPSecurityProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPSecurityProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iPv4RangeCollectionResponse") {
		var out IPv4RangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPv4RangeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.identityFindingCollectionResponse") {
		var out IdentityFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.identityUserFlowCollectionResponse") {
		var out IdentityUserFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityUserFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.impactedResourceCollectionResponse") {
		var out ImpactedResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImpactedResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedAppleDeviceIdentityCollectionResponse") {
		var out ImportedAppleDeviceIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedAppleDeviceIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedAppleDeviceIdentityResultCollectionResponse") {
		var out ImportedAppleDeviceIdentityResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedAppleDeviceIdentityResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedDeviceIdentityCollectionResponse") {
		var out ImportedDeviceIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedDeviceIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedDeviceIdentityResultCollectionResponse") {
		var out ImportedDeviceIdentityResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedDeviceIdentityResultCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.inactiveAwsResourceFindingCollectionResponse") {
		var out InactiveAwsResourceFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveAwsResourceFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveAwsRoleFindingCollectionResponse") {
		var out InactiveAwsRoleFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveAwsRoleFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveAzureServicePrincipalFindingCollectionResponse") {
		var out InactiveAzureServicePrincipalFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveAzureServicePrincipalFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveGcpServiceAccountFindingCollectionResponse") {
		var out InactiveGcpServiceAccountFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveGcpServiceAccountFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveGroupFindingCollectionResponse") {
		var out InactiveGroupFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveGroupFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveServerlessFunctionFindingCollectionResponse") {
		var out InactiveServerlessFunctionFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveServerlessFunctionFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveUserFindingCollectionResponse") {
		var out InactiveUserFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveUserFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inboundSharedUserProfileCollectionResponse") {
		var out InboundSharedUserProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InboundSharedUserProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.includeTargetCollectionResponse") {
		var out IncludeTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IncludeTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.administrativeUnitProvisioningFlowCollectionResponse") {
		var out IndustryDataAdministrativeUnitProvisioningFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataAdministrativeUnitProvisioningFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.apiDataConnectorCollectionResponse") {
		var out IndustryDataApiDataConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataApiDataConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.azureDataLakeConnectorCollectionResponse") {
		var out IndustryDataAzureDataLakeConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataAzureDataLakeConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.classGroupProvisioningFlowCollectionResponse") {
		var out IndustryDataClassGroupProvisioningFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataClassGroupProvisioningFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.fileDataConnectorCollectionResponse") {
		var out IndustryDataFileDataConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataFileDataConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.fileValidateOperationCollectionResponse") {
		var out IndustryDataFileValidateOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataFileValidateOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundApiFlowCollectionResponse") {
		var out IndustryDataInboundApiFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundApiFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundFileFlowCollectionResponse") {
		var out IndustryDataInboundFileFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundFileFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundFlowActivityCollectionResponse") {
		var out IndustryDataInboundFlowActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundFlowActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundFlowCollectionResponse") {
		var out IndustryDataInboundFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.industryDataActivityStatisticsCollectionResponse") {
		var out IndustryDataIndustryDataActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataIndustryDataActivityStatisticsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.industryDataConnectorCollectionResponse") {
		var out IndustryDataIndustryDataConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataIndustryDataConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.industryDataRunActivityCollectionResponse") {
		var out IndustryDataIndustryDataRunActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataIndustryDataRunActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.industryDataRunCollectionResponse") {
		var out IndustryDataIndustryDataRunCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataIndustryDataRunCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.industryDataRunRoleCountMetricCollectionResponse") {
		var out IndustryDataIndustryDataRunRoleCountMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataIndustryDataRunRoleCountMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.oneRosterApiDataConnectorCollectionResponse") {
		var out IndustryDataOneRosterApiDataConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOneRosterApiDataConnectorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.outboundFlowActivityCollectionResponse") {
		var out IndustryDataOutboundFlowActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOutboundFlowActivityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.outboundProvisioningFlowSetCollectionResponse") {
		var out IndustryDataOutboundProvisioningFlowSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOutboundProvisioningFlowSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.provisioningFlowCollectionResponse") {
		var out IndustryDataProvisioningFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataProvisioningFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.referenceDefinitionCollectionResponse") {
		var out IndustryDataReferenceDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataReferenceDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.roleGroupCollectionResponse") {
		var out IndustryDataRoleGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataRoleGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.roleReferenceValueCollectionResponse") {
		var out IndustryDataRoleReferenceValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataRoleReferenceValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.sectionRoleReferenceValueCollectionResponse") {
		var out IndustryDataSectionRoleReferenceValueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataSectionRoleReferenceValueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.securityGroupProvisioningFlowCollectionResponse") {
		var out IndustryDataSecurityGroupProvisioningFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataSecurityGroupProvisioningFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.sourceSystemDefinitionCollectionResponse") {
		var out IndustryDataSourceSystemDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataSourceSystemDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.userConfigurationCollectionResponse") {
		var out IndustryDataUserConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataUserConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.userMatchingSettingCollectionResponse") {
		var out IndustryDataUserMatchingSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataUserMatchingSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.userProvisioningFlowCollectionResponse") {
		var out IndustryDataUserProvisioningFlowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataUserProvisioningFlowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.validateOperationCollectionResponse") {
		var out IndustryDataValidateOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataValidateOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.yearTimePeriodDefinitionCollectionResponse") {
		var out IndustryDataYearTimePeriodDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataYearTimePeriodDefinitionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.informationProtectionActionCollectionResponse") {
		var out InformationProtectionActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InformationProtectionActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.informationProtectionLabelCollectionResponse") {
		var out InformationProtectionLabelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InformationProtectionLabelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.insightSummaryCollectionResponse") {
		var out InsightSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InsightSummaryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.intuneBrandingProfileAssignmentCollectionResponse") {
		var out IntuneBrandingProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IntuneBrandingProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.intuneBrandingProfileCollectionResponse") {
		var out IntuneBrandingProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IntuneBrandingProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.invalidLicenseAlertConfigurationCollectionResponse") {
		var out InvalidLicenseAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvalidLicenseAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.invalidLicenseAlertIncidentCollectionResponse") {
		var out InvalidLicenseAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvalidLicenseAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.invokeUserFlowListenerCollectionResponse") {
		var out InvokeUserFlowListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvokeUserFlowListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosBookmarkCollectionResponse") {
		var out IosBookmarkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosBookmarkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCertificateProfileBaseCollectionResponse") {
		var out IosCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosDerivedCredentialAuthenticationConfigurationCollectionResponse") {
		var out IosDerivedCredentialAuthenticationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosDerivedCredentialAuthenticationConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosEasEmailProfileConfigurationCollectionResponse") {
		var out IosEasEmailProfileConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosEasEmailProfileConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosEduDeviceConfigurationCollectionResponse") {
		var out IosEduDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosEduDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosEducationDeviceConfigurationCollectionResponse") {
		var out IosEducationDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosEducationDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosEnterpriseWiFiConfigurationCollectionResponse") {
		var out IosEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosExpeditedCheckinConfigurationCollectionResponse") {
		var out IosExpeditedCheckinConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosExpeditedCheckinConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosImportedPFXCertificateProfileCollectionResponse") {
		var out IosImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosImportedPFXCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppProvisioningConfigurationAssignmentCollectionResponse") {
		var out IosLobAppProvisioningConfigurationAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppProvisioningConfigurationAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppProvisioningConfigurationCollectionResponse") {
		var out IosLobAppProvisioningConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppProvisioningConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppProvisioningConfigurationPolicySetItemCollectionResponse") {
		var out IosLobAppProvisioningConfigurationPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppProvisioningConfigurationPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosPkcsCertificateProfileCollectionResponse") {
		var out IosPkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosPkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosScepCertificateProfileCollectionResponse") {
		var out IosScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosScepCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosTrustedRootCertificateCollectionResponse") {
		var out IosTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosTrustedRootCertificateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosVpnConfigurationCollectionResponse") {
		var out IosVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppAssignedDeviceLicenseCollectionResponse") {
		var out IosVppAppAssignedDeviceLicenseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppAssignedDeviceLicenseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppAssignedLicenseCollectionResponse") {
		var out IosVppAppAssignedLicenseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppAssignedLicenseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppAssignedUserLicenseCollectionResponse") {
		var out IosVppAppAssignedUserLicenseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppAssignedUserLicenseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppRevokeLicensesActionResultCollectionResponse") {
		var out IosVppAppRevokeLicensesActionResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppRevokeLicensesActionResultCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosWiFiConfigurationCollectionResponse") {
		var out IosWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosWiFiConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.iosikEv2VpnConfigurationCollectionResponse") {
		var out IosikEv2VpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosikEv2VpnConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.itemActivityOLDCollectionResponse") {
		var out ItemActivityOLDCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemActivityOLDCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.itemAddressCollectionResponse") {
		var out ItemAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemAddressCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.itemCategoryCollectionResponse") {
		var out ItemCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemCollectionResponse") {
		var out ItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemEmailCollectionResponse") {
		var out ItemEmailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemEmailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemPatentCollectionResponse") {
		var out ItemPatentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemPatentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemPhoneCollectionResponse") {
		var out ItemPhoneCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemPhoneCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemPublicationCollectionResponse") {
		var out ItemPublicationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemPublicationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.jobResponseBaseCollectionResponse") {
		var out JobResponseBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobResponseBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.journalCollectionResponse") {
		var out JournalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JournalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.journalLineCollectionResponse") {
		var out JournalLineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JournalLineCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.keyTypedValuePairCollectionResponse") {
		var out KeyTypedValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyTypedValuePairCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.labelActionBaseCollectionResponse") {
		var out LabelActionBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LabelActionBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.labelContentRightCollectionResponse") {
		var out LabelContentRightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LabelContentRightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.labelPolicyCollectionResponse") {
		var out LabelPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LabelPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.languageProficiencyCollectionResponse") {
		var out LanguageProficiencyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LanguageProficiencyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.localeInfoCollectionResponse") {
		var out LocaleInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocaleInfoCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.loggedOnUserCollectionResponse") {
		var out LoggedOnUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LoggedOnUserCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.logonUserCollectionResponse") {
		var out LogonUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogonUserCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.lookupResultRowCollectionResponse") {
		var out LookupResultRowCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LookupResultRowCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSAppleEventReceiverCollectionResponse") {
		var out MacOSAppleEventReceiverCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSAppleEventReceiverCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSAssociatedDomainsItemCollectionResponse") {
		var out MacOSAssociatedDomainsItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSAssociatedDomainsItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSCertificateProfileBaseCollectionResponse") {
		var out MacOSCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSCertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.macOSCustomAppConfigurationCollectionResponse") {
		var out MacOSCustomAppConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSCustomAppConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.macOSEndpointProtectionConfigurationCollectionResponse") {
		var out MacOSEndpointProtectionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSEndpointProtectionConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSEnterpriseWiFiConfigurationCollectionResponse") {
		var out MacOSEnterpriseWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSEnterpriseWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSExtensionsConfigurationCollectionResponse") {
		var out MacOSExtensionsConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSExtensionsConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSFirewallApplicationCollectionResponse") {
		var out MacOSFirewallApplicationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSFirewallApplicationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.macOSImportedPFXCertificateProfileCollectionResponse") {
		var out MacOSImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSImportedPFXCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.macOSKernelExtensionCollectionResponse") {
		var out MacOSKernelExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSKernelExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSLaunchItemCollectionResponse") {
		var out MacOSLaunchItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSLaunchItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.macOSPkcsCertificateProfileCollectionResponse") {
		var out MacOSPkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSPkcsCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSPkgAppCollectionResponse") {
		var out MacOSPkgAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSPkgAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSPrivacyAccessControlItemCollectionResponse") {
		var out MacOSPrivacyAccessControlItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSPrivacyAccessControlItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSScepCertificateProfileCollectionResponse") {
		var out MacOSScepCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSScepCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSoftwareUpdateAccountSummaryCollectionResponse") {
		var out MacOSSoftwareUpdateAccountSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSoftwareUpdateAccountSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSoftwareUpdateCategorySummaryCollectionResponse") {
		var out MacOSSoftwareUpdateCategorySummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSoftwareUpdateCategorySummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSoftwareUpdateConfigurationCollectionResponse") {
		var out MacOSSoftwareUpdateConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSoftwareUpdateConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSoftwareUpdateStateSummaryCollectionResponse") {
		var out MacOSSoftwareUpdateStateSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSoftwareUpdateStateSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSystemExtensionCollectionResponse") {
		var out MacOSSystemExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSystemExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSystemExtensionTypeMappingCollectionResponse") {
		var out MacOSSystemExtensionTypeMappingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSystemExtensionTypeMappingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSTrustedRootCertificateCollectionResponse") {
		var out MacOSTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSVpnConfigurationCollectionResponse") {
		var out MacOSVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSWebClipCollectionResponse") {
		var out MacOSWebClipCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSWebClipCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSWiFiConfigurationCollectionResponse") {
		var out MacOSWiFiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSWiFiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSWiredNetworkConfigurationCollectionResponse") {
		var out MacOSWiredNetworkConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSWiredNetworkConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOsVppAppAssignedLicenseCollectionResponse") {
		var out MacOsVppAppAssignedLicenseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOsVppAppAssignedLicenseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOsVppAppCollectionResponse") {
		var out MacOsVppAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOsVppAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOsVppAppRevokeLicensesActionResultCollectionResponse") {
		var out MacOsVppAppRevokeLicensesActionResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOsVppAppRevokeLicensesActionResultCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mailFolderOperationCollectionResponse") {
		var out MailFolderOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailFolderOperationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mailboxCollectionResponse") {
		var out MailboxCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxFolderCollectionResponse") {
		var out MailboxFolderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxFolderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxItemCollectionResponse") {
		var out MailboxItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedAllDeviceCertificateStateCollectionResponse") {
		var out ManagedAllDeviceCertificateStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAllDeviceCertificateStateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedAppLogCollectionRequestCollectionResponse") {
		var out ManagedAppLogCollectionRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppLogCollectionRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppLogUploadCollectionResponse") {
		var out ManagedAppLogUploadCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppLogUploadCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedAppProtectionPolicySetItemCollectionResponse") {
		var out ManagedAppProtectionPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppProtectionPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceCertificateStateCollectionResponse") {
		var out ManagedDeviceCertificateStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceCertificateStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceCleanupRuleCollectionResponse") {
		var out ManagedDeviceCleanupRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceCleanupRuleCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceEncryptionStateCollectionResponse") {
		var out ManagedDeviceEncryptionStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceEncryptionStateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationPolicySetItemCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationPolicySetItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationSettingStateCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationSettingStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationSettingStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationStateCollectionResponse") {
		var out ManagedDeviceMobileAppConfigurationStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationStateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceReportedAppCollectionResponse") {
		var out ManagedDeviceReportedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceReportedAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceWindowsOperatingSystemEditionCollectionResponse") {
		var out ManagedDeviceWindowsOperatingSystemEditionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceWindowsOperatingSystemEditionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceWindowsOperatingSystemImageCollectionResponse") {
		var out ManagedDeviceWindowsOperatingSystemImageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceWindowsOperatingSystemImageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceWindowsOperatingSystemUpdateCollectionResponse") {
		var out ManagedDeviceWindowsOperatingSystemUpdateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceWindowsOperatingSystemUpdateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedEBookCategoryCollectionResponse") {
		var out ManagedEBookCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedEBookCategoryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.aggregatedPolicyComplianceCollectionResponse") {
		var out ManagedTenantsAggregatedPolicyComplianceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsAggregatedPolicyComplianceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.alertDataReferenceStringCollectionResponse") {
		var out ManagedTenantsAlertDataReferenceStringCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsAlertDataReferenceStringCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.appPerformanceCollectionResponse") {
		var out ManagedTenantsAppPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsAppPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.auditEventCollectionResponse") {
		var out ManagedTenantsAuditEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsAuditEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.cloudPcConnectionCollectionResponse") {
		var out ManagedTenantsCloudPCConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsCloudPCConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.cloudPcDeviceCollectionResponse") {
		var out ManagedTenantsCloudPCDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsCloudPCDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.cloudPcOverviewCollectionResponse") {
		var out ManagedTenantsCloudPCOverviewCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsCloudPCOverviewCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.conditionalAccessPolicyCoverageCollectionResponse") {
		var out ManagedTenantsConditionalAccessPolicyCoverageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsConditionalAccessPolicyCoverageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.credentialUserRegistrationsSummaryCollectionResponse") {
		var out ManagedTenantsCredentialUserRegistrationsSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsCredentialUserRegistrationsSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.deviceAppPerformanceCollectionResponse") {
		var out ManagedTenantsDeviceAppPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsDeviceAppPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.deviceCompliancePolicySettingStateSummaryCollectionResponse") {
		var out ManagedTenantsDeviceCompliancePolicySettingStateSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsDeviceCompliancePolicySettingStateSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.deviceHealthStatusCollectionResponse") {
		var out ManagedTenantsDeviceHealthStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsDeviceHealthStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.emailCollectionResponse") {
		var out ManagedTenantsEmailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsEmailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedDeviceComplianceCollectionResponse") {
		var out ManagedTenantsManagedDeviceComplianceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedDeviceComplianceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedDeviceComplianceTrendCollectionResponse") {
		var out ManagedTenantsManagedDeviceComplianceTrendCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedDeviceComplianceTrendCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantAlertCollectionResponse") {
		var out ManagedTenantsManagedTenantAlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantAlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantAlertLogCollectionResponse") {
		var out ManagedTenantsManagedTenantAlertLogCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantAlertLogCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantAlertRuleCollectionResponse") {
		var out ManagedTenantsManagedTenantAlertRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantAlertRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantAlertRuleDefinitionCollectionResponse") {
		var out ManagedTenantsManagedTenantAlertRuleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantAlertRuleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantApiNotificationCollectionResponse") {
		var out ManagedTenantsManagedTenantApiNotificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantApiNotificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantEmailNotificationCollectionResponse") {
		var out ManagedTenantsManagedTenantEmailNotificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantEmailNotificationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantTicketingEndpointCollectionResponse") {
		var out ManagedTenantsManagedTenantTicketingEndpointCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantTicketingEndpointCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementActionCollectionResponse") {
		var out ManagedTenantsManagementActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementActionDeploymentStatusCollectionResponse") {
		var out ManagedTenantsManagementActionDeploymentStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementActionDeploymentStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementActionInfoCollectionResponse") {
		var out ManagedTenantsManagementActionInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementActionInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementActionTenantDeploymentStatusCollectionResponse") {
		var out ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementIntentCollectionResponse") {
		var out ManagedTenantsManagementIntentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementIntentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementIntentInfoCollectionResponse") {
		var out ManagedTenantsManagementIntentInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementIntentInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateCollectionCollectionResponse") {
		var out ManagedTenantsManagementTemplateCollectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateCollectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateCollectionResponse") {
		var out ManagedTenantsManagementTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateCollectionTenantSummaryCollectionResponse") {
		var out ManagedTenantsManagementTemplateCollectionTenantSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateCollectionTenantSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateDetailedInfoCollectionResponse") {
		var out ManagedTenantsManagementTemplateDetailedInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateDetailedInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateStepCollectionResponse") {
		var out ManagedTenantsManagementTemplateStepCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateStepCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateStepDeploymentCollectionResponse") {
		var out ManagedTenantsManagementTemplateStepDeploymentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateStepDeploymentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateStepTenantSummaryCollectionResponse") {
		var out ManagedTenantsManagementTemplateStepTenantSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateStepTenantSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managementTemplateStepVersionCollectionResponse") {
		var out ManagedTenantsManagementTemplateStepVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagementTemplateStepVersionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.myRoleCollectionResponse") {
		var out ManagedTenantsMyRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsMyRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.notificationTargetCollectionResponse") {
		var out ManagedTenantsNotificationTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsNotificationTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.roleAssignmentCollectionResponse") {
		var out ManagedTenantsRoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsRoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.roleDefinitionCollectionResponse") {
		var out ManagedTenantsRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.settingCollectionResponse") {
		var out ManagedTenantsSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.templateParameterCollectionResponse") {
		var out ManagedTenantsTemplateParameterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTemplateParameterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantCollectionResponse") {
		var out ManagedTenantsTenantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantContactInformationCollectionResponse") {
		var out ManagedTenantsTenantContactInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantContactInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantCustomizedInformationCollectionResponse") {
		var out ManagedTenantsTenantCustomizedInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantCustomizedInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantDetailedInformationCollectionResponse") {
		var out ManagedTenantsTenantDetailedInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantDetailedInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantGroupCollectionResponse") {
		var out ManagedTenantsTenantGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantInfoCollectionResponse") {
		var out ManagedTenantsTenantInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.tenantTagCollectionResponse") {
		var out ManagedTenantsTenantTagCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsTenantTagCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.windowsDeviceMalwareStateCollectionResponse") {
		var out ManagedTenantsWindowsDeviceMalwareStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsWindowsDeviceMalwareStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.windowsProtectionStateCollectionResponse") {
		var out ManagedTenantsWindowsProtectionStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsWindowsProtectionStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.workloadActionCollectionResponse") {
		var out ManagedTenantsWorkloadActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsWorkloadActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.workloadActionDeploymentStatusCollectionResponse") {
		var out ManagedTenantsWorkloadActionDeploymentStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsWorkloadActionDeploymentStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.workloadStatusCollectionResponse") {
		var out ManagedTenantsWorkloadStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsWorkloadStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managementCertificateWithThumbprintCollectionResponse") {
		var out ManagementCertificateWithThumbprintCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagementCertificateWithThumbprintCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.matchLocationCollectionResponse") {
		var out MatchLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MatchLocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.matchingDlpRuleCollectionResponse") {
		var out MatchingDlpRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MatchingDlpRuleCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mdmWindowsInformationProtectionPolicyPolicySetItemCollectionResponse") {
		var out MdmWindowsInformationProtectionPolicyPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MdmWindowsInformationProtectionPolicyPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.meetingActivityStatisticsCollectionResponse") {
		var out MeetingActivityStatisticsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingActivityStatisticsCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.meetingNoteCollectionResponse") {
		var out MeetingNoteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingNoteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingNoteSubpointCollectionResponse") {
		var out MeetingNoteSubpointCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingNoteSubpointCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.meetingRegistrantBaseCollectionResponse") {
		var out MeetingRegistrantBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingRegistrantBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingRegistrantCollectionResponse") {
		var out MeetingRegistrantCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingRegistrantCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingRegistrationQuestionCollectionResponse") {
		var out MeetingRegistrationQuestionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingRegistrationQuestionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingSpeakerCollectionResponse") {
		var out MeetingSpeakerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingSpeakerCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mentionCollectionResponse") {
		var out MentionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MentionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mentionEventCollectionResponse") {
		var out MentionEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MentionEventCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.messageEventCollectionResponse") {
		var out MessageEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messageRecipientCollectionResponse") {
		var out MessageRecipientCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageRecipientCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.messageTraceCollectionResponse") {
		var out MessageTraceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageTraceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.metaDataKeyValuePairCollectionResponse") {
		var out MetaDataKeyValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MetaDataKeyValuePairCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mfaCompletionMetricCollectionResponse") {
		var out MfaCompletionMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MfaCompletionMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mfaFailureCollectionResponse") {
		var out MfaFailureCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MfaFailureCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mfaTelecomFraudMetricCollectionResponse") {
		var out MfaTelecomFraudMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MfaTelecomFraudMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mfaUserCountMetricCollectionResponse") {
		var out MfaUserCountMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MfaUserCountMetricCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.microsoftStoreForBusinessContainedAppCollectionResponse") {
		var out MicrosoftStoreForBusinessContainedAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftStoreForBusinessContainedAppCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.microsoftTunnelConfigurationCollectionResponse") {
		var out MicrosoftTunnelConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTunnelConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftTunnelHealthThresholdCollectionResponse") {
		var out MicrosoftTunnelHealthThresholdCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTunnelHealthThresholdCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftTunnelServerCollectionResponse") {
		var out MicrosoftTunnelServerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTunnelServerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftTunnelServerLogCollectionResponseCollectionResponse") {
		var out MicrosoftTunnelServerLogCollectionResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTunnelServerLogCollectionResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftTunnelSiteCollectionResponse") {
		var out MicrosoftTunnelSiteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTunnelSiteCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mobileAppCatalogPackageCollectionResponse") {
		var out MobileAppCatalogPackageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppCatalogPackageCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mobileAppDependencyCollectionResponse") {
		var out MobileAppDependencyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppDependencyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppInstallStatusCollectionResponse") {
		var out MobileAppInstallStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppInstallStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppIntentAndStateCollectionResponse") {
		var out MobileAppIntentAndStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppIntentAndStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppIntentAndStateDetailCollectionResponse") {
		var out MobileAppIntentAndStateDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppIntentAndStateDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppPolicySetItemCollectionResponse") {
		var out MobileAppPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppPolicySetItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppProvisioningConfigGroupAssignmentCollectionResponse") {
		var out MobileAppProvisioningConfigGroupAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppProvisioningConfigGroupAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppRelationshipCollectionResponse") {
		var out MobileAppRelationshipCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppRelationshipCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppSupersedenceCollectionResponse") {
		var out MobileAppSupersedenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppSupersedenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppSupportedDeviceTypeCollectionResponse") {
		var out MobileAppSupportedDeviceTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppSupportedDeviceTypeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingHistoryItemCollectionResponse") {
		var out MobileAppTroubleshootingHistoryItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingHistoryItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.mobilityManagementPolicyCollectionResponse") {
		var out MobilityManagementPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobilityManagementPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.monthlyInactiveUsersByApplicationMetricCollectionResponse") {
		var out MonthlyInactiveUsersByApplicationMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MonthlyInactiveUsersByApplicationMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.monthlyInactiveUsersMetricCollectionResponse") {
		var out MonthlyInactiveUsersMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MonthlyInactiveUsersMetricCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.muteParticipantsOperationCollectionResponse") {
		var out MuteParticipantsOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MuteParticipantsOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mutualTlsOauthConfigurationCollectionResponse") {
		var out MutualTlsOauthConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MutualTlsOauthConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ndesConnectorCollectionResponse") {
		var out NdesConnectorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NdesConnectorCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.networkInterfaceCollectionResponse") {
		var out NetworkInterfaceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkInterfaceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkLocationDetailCollectionResponse") {
		var out NetworkLocationDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkLocationDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.alertActionCollectionResponse") {
		var out NetworkaccessAlertActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessAlertActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.alertCollectionResponse") {
		var out NetworkaccessAlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessAlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.associationCollectionResponse") {
		var out NetworkaccessAssociationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessAssociationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.branchSiteCollectionResponse") {
		var out NetworkaccessBranchSiteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessBranchSiteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.conditionalAccessPolicyCollectionResponse") {
		var out NetworkaccessConditionalAccessPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessConditionalAccessPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.connectionCollectionResponse") {
		var out NetworkaccessConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.connectivityConfigurationLinkCollectionResponse") {
		var out NetworkaccessConnectivityConfigurationLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessConnectivityConfigurationLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.deviceLinkCollectionResponse") {
		var out NetworkaccessDeviceLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessDeviceLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringPolicyCollectionResponse") {
		var out NetworkaccessFilteringPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringPolicyLinkCollectionResponse") {
		var out NetworkaccessFilteringPolicyLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringPolicyLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringProfileCollectionResponse") {
		var out NetworkaccessFilteringProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringRuleCollectionResponse") {
		var out NetworkaccessFilteringRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingPolicyCollectionResponse") {
		var out NetworkaccessForwardingPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingPolicyLinkCollectionResponse") {
		var out NetworkaccessForwardingPolicyLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingPolicyLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingProfileCollectionResponse") {
		var out NetworkaccessForwardingProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingRuleCollectionResponse") {
		var out NetworkaccessForwardingRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.fqdnFilteringRuleCollectionResponse") {
		var out NetworkaccessFqdnFilteringRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFqdnFilteringRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.internetAccessForwardingRuleCollectionResponse") {
		var out NetworkaccessInternetAccessForwardingRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessInternetAccessForwardingRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.localConnectivityConfigurationCollectionResponse") {
		var out NetworkaccessLocalConnectivityConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessLocalConnectivityConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.m365ForwardingRuleCollectionResponse") {
		var out NetworkaccessM365ForwardingRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessM365ForwardingRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.networkAccessTrafficCollectionResponse") {
		var out NetworkaccessNetworkAccessTrafficCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessNetworkAccessTrafficCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.policyLinkCollectionResponse") {
		var out NetworkaccessPolicyLinkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessPolicyLinkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.policyRuleCollectionResponse") {
		var out NetworkaccessPolicyRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessPolicyRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.privateAccessForwardingRuleCollectionResponse") {
		var out NetworkaccessPrivateAccessForwardingRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessPrivateAccessForwardingRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedResourceCollectionResponse") {
		var out NetworkaccessRelatedResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.remoteNetworkCollectionResponse") {
		var out NetworkaccessRemoteNetworkCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRemoteNetworkCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.remoteNetworkHealthEventCollectionResponse") {
		var out NetworkaccessRemoteNetworkHealthEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRemoteNetworkHealthEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.ruleDestinationCollectionResponse") {
		var out NetworkaccessRuleDestinationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRuleDestinationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.webCategoryCollectionResponse") {
		var out NetworkaccessWebCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessWebCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.webCategoryFilteringRuleCollectionResponse") {
		var out NetworkaccessWebCategoryFilteringRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessWebCategoryFilteringRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.newsLinkPageCollectionResponse") {
		var out NewsLinkPageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NewsLinkPageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noMfaOnRoleActivationAlertConfigurationCollectionResponse") {
		var out NoMfaOnRoleActivationAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoMfaOnRoleActivationAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noMfaOnRoleActivationAlertIncidentCollectionResponse") {
		var out NoMfaOnRoleActivationAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoMfaOnRoleActivationAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.notificationCollectionResponse") {
		var out NotificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NotificationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.numberRangeCollectionResponse") {
		var out NumberRangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NumberRangeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.officeSuiteAppCollectionResponse") {
		var out OfficeSuiteAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeSuiteAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oidcIdentityProviderCollectionResponse") {
		var out OidcIdentityProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OidcIdentityProviderCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.onEmailOtpSendListenerCollectionResponse") {
		var out OnEmailOtpSendListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnEmailOtpSendListenerCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.onOtpSendCustomExtensionCollectionResponse") {
		var out OnOtpSendCustomExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnOtpSendCustomExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPhoneMethodLoadStartListenerCollectionResponse") {
		var out OnPhoneMethodLoadStartListenerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPhoneMethodLoadStartListenerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesAgentCollectionResponse") {
		var out OnPremisesAgentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesAgentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesAgentGroupCollectionResponse") {
		var out OnPremisesAgentGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesAgentGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesApplicationSegmentCollectionResponse") {
		var out OnPremisesApplicationSegmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesApplicationSegmentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.onPremisesPublishingProfileCollectionResponse") {
		var out OnPremisesPublishingProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesPublishingProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.openAwsSecurityGroupFindingCollectionResponse") {
		var out OpenAwsSecurityGroupFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenAwsSecurityGroupFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openIdConnectIdentityProviderCollectionResponse") {
		var out OpenIdConnectIdentityProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenIdConnectIdentityProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openIdConnectProviderCollectionResponse") {
		var out OpenIdConnectProviderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenIdConnectProviderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openNetworkAzureSecurityGroupFindingCollectionResponse") {
		var out OpenNetworkAzureSecurityGroupFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenNetworkAzureSecurityGroupFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.operatingSystemVersionRangeCollectionResponse") {
		var out OperatingSystemVersionRangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperatingSystemVersionRangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.operationApprovalPolicyCollectionResponse") {
		var out OperationApprovalPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationApprovalPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.operationApprovalRequestCollectionResponse") {
		var out OperationApprovalRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationApprovalRequestCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.organizationalUnitCollectionResponse") {
		var out OrganizationalUnitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizationalUnitCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.outboundSharedUserProfileCollectionResponse") {
		var out OutboundSharedUserProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutboundSharedUserProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.outlookTaskCollectionResponse") {
		var out OutlookTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outlookTaskFolderCollectionResponse") {
		var out OutlookTaskFolderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookTaskFolderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outlookTaskGroupCollectionResponse") {
		var out OutlookTaskGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookTaskGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedAwsResourceFindingCollectionResponse") {
		var out OverprovisionedAwsResourceFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedAwsResourceFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedAwsRoleFindingCollectionResponse") {
		var out OverprovisionedAwsRoleFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedAwsRoleFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedAzureServicePrincipalFindingCollectionResponse") {
		var out OverprovisionedAzureServicePrincipalFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedAzureServicePrincipalFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedGcpServiceAccountFindingCollectionResponse") {
		var out OverprovisionedGcpServiceAccountFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedGcpServiceAccountFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedServerlessFunctionFindingCollectionResponse") {
		var out OverprovisionedServerlessFunctionFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedServerlessFunctionFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedUserFindingCollectionResponse") {
		var out OverprovisionedUserFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedUserFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pageTemplateCollectionResponse") {
		var out PageTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PageTemplateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.partner.security.activityLogCollectionResponse") {
		var out PartnerSecurityActivityLogCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityActivityLogCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.adminsMfaEnforcedSecurityRequirementCollectionResponse") {
		var out PartnerSecurityAdminsMfaEnforcedSecurityRequirementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityAdminsMfaEnforcedSecurityRequirementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.affectedResourceCollectionResponse") {
		var out PartnerSecurityAffectedResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityAffectedResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.customerInsightCollectionResponse") {
		var out PartnerSecurityCustomerInsightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityCustomerInsightCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.customersMfaEnforcedSecurityRequirementCollectionResponse") {
		var out PartnerSecurityCustomersMfaEnforcedSecurityRequirementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityCustomersMfaEnforcedSecurityRequirementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.customersSpendingBudgetSecurityRequirementCollectionResponse") {
		var out PartnerSecurityCustomersSpendingBudgetSecurityRequirementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityCustomersSpendingBudgetSecurityRequirementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.partnerSecurityAlertCollectionResponse") {
		var out PartnerSecurityPartnerSecurityAlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityPartnerSecurityAlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.responseTimeSecurityRequirementCollectionResponse") {
		var out PartnerSecurityResponseTimeSecurityRequirementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityResponseTimeSecurityRequirementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.securityRequirementCollectionResponse") {
		var out PartnerSecuritySecurityRequirementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecuritySecurityRequirementCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.securityScoreHistoryCollectionResponse") {
		var out PartnerSecuritySecurityScoreHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecuritySecurityScoreHistoryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.passkeyAuthenticationMethodTargetCollectionResponse") {
		var out PasskeyAuthenticationMethodTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasskeyAuthenticationMethodTargetCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.passwordSingleSignOnFieldCollectionResponse") {
		var out PasswordSingleSignOnFieldCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordSingleSignOnFieldCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse") {
		var out PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.payloadByFilterCollectionResponse") {
		var out PayloadByFilterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PayloadByFilterCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.payloadCompatibleAssignmentFilterCollectionResponse") {
		var out PayloadCompatibleAssignmentFilterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PayloadCompatibleAssignmentFilterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.payloadResponseCollectionResponse") {
		var out PayloadResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PayloadResponseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.paymentMethodCollectionResponse") {
		var out PaymentMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PaymentMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.paymentTermCollectionResponse") {
		var out PaymentTermCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PaymentTermCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pendingExternalUserProfileCollectionResponse") {
		var out PendingExternalUserProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PendingExternalUserProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.permissionGrantPreApprovalPolicyCollectionResponse") {
		var out PermissionGrantPreApprovalPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionGrantPreApprovalPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.permissionsCreepIndexDistributionCollectionResponse") {
		var out PermissionsCreepIndexDistributionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionsCreepIndexDistributionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionsDefinitionAwsPolicyCollectionResponse") {
		var out PermissionsDefinitionAwsPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionsDefinitionAwsPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionsDefinitionAzureRoleCollectionResponse") {
		var out PermissionsDefinitionAzureRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionsDefinitionAzureRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionsDefinitionGcpRoleCollectionResponse") {
		var out PermissionsDefinitionGcpRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionsDefinitionGcpRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionsRequestChangeCollectionResponse") {
		var out PermissionsRequestChangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionsRequestChangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personAnnotationCollectionResponse") {
		var out PersonAnnotationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonAnnotationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personAnnualEventCollectionResponse") {
		var out PersonAnnualEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonAnnualEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personAwardCollectionResponse") {
		var out PersonAwardCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonAwardCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personCertificationCollectionResponse") {
		var out PersonCertificationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonCertificationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.personDataSourceCollectionResponse") {
		var out PersonDataSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonDataSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personExtensionCollectionResponse") {
		var out PersonExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personInterestCollectionResponse") {
		var out PersonInterestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonInterestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personNameCollectionResponse") {
		var out PersonNameCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonNameCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personWebsiteCollectionResponse") {
		var out PersonWebsiteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonWebsiteCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.pictureCollectionResponse") {
		var out PictureCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PictureCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.plannerAssignedToTaskBoardTaskFormatCollectionResponse") {
		var out PlannerAssignedToTaskBoardTaskFormatCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerAssignedToTaskBoardTaskFormatCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.plannerBucketTaskBoardTaskFormatCollectionResponse") {
		var out PlannerBucketTaskBoardTaskFormatCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucketTaskBoardTaskFormatCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerDeltaCollectionResponse") {
		var out PlannerDeltaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerDeltaCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanConfigurationBucketDefinitionCollectionResponse") {
		var out PlannerPlanConfigurationBucketDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanConfigurationBucketDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanConfigurationBucketLocalizationCollectionResponse") {
		var out PlannerPlanConfigurationBucketLocalizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanConfigurationBucketLocalizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanConfigurationLocalizationCollectionResponse") {
		var out PlannerPlanConfigurationLocalizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanConfigurationLocalizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanDetailsCollectionResponse") {
		var out PlannerPlanDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerProgressTaskBoardTaskFormatCollectionResponse") {
		var out PlannerProgressTaskBoardTaskFormatCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerProgressTaskBoardTaskFormatCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerRosterCollectionResponse") {
		var out PlannerRosterCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerRosterCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerRosterMemberCollectionResponse") {
		var out PlannerRosterMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerRosterMemberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerRuleOverrideCollectionResponse") {
		var out PlannerRuleOverrideCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerRuleOverrideCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerSharedWithContainerCollectionResponse") {
		var out PlannerSharedWithContainerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerSharedWithContainerCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.plannerTaskDetailsCollectionResponse") {
		var out PlannerTaskDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTaskDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTaskRoleBasedRuleCollectionResponse") {
		var out PlannerTaskRoleBasedRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTaskRoleBasedRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerUserCollectionResponse") {
		var out PlannerUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerUserCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.policyLocationCollectionResponse") {
		var out PolicyLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyLocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policySetAssignmentCollectionResponse") {
		var out PolicySetAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicySetAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policySetCollectionResponse") {
		var out PolicySetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicySetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policySetItemCollectionResponse") {
		var out PolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.preApprovalDetailCollectionResponse") {
		var out PreApprovalDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PreApprovalDetailCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.printUsageCollectionResponse") {
		var out PrintUsageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintUsageCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationAwsResourceFindingCollectionResponse") {
		var out PrivilegeEscalationAwsResourceFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationAwsResourceFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationAwsRoleFindingCollectionResponse") {
		var out PrivilegeEscalationAwsRoleFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationAwsRoleFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationCollectionResponse") {
		var out PrivilegeEscalationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationFindingCollectionResponse") {
		var out PrivilegeEscalationFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationGcpServiceAccountFindingCollectionResponse") {
		var out PrivilegeEscalationGcpServiceAccountFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationGcpServiceAccountFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationUserFindingCollectionResponse") {
		var out PrivilegeEscalationUserFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationUserFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeManagementElevationCollectionResponse") {
		var out PrivilegeManagementElevationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeManagementElevationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeManagementElevationRequestCollectionResponse") {
		var out PrivilegeManagementElevationRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeManagementElevationRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessCollectionResponse") {
		var out PrivilegedAccessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.privilegedApprovalCollectionResponse") {
		var out PrivilegedApprovalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedApprovalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedOperationEventCollectionResponse") {
		var out PrivilegedOperationEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedOperationEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedRoleAssignmentCollectionResponse") {
		var out PrivilegedRoleAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedRoleAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedRoleAssignmentRequestCollectionResponse") {
		var out PrivilegedRoleAssignmentRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedRoleAssignmentRequestCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedRoleCollectionResponse") {
		var out PrivilegedRoleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedRoleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedSignupStatusCollectionResponse") {
		var out PrivilegedSignupStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedSignupStatusCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.processContentMetadataBaseCollectionResponse") {
		var out ProcessContentMetadataBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessContentMetadataBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.processingErrorCollectionResponse") {
		var out ProcessingErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessingErrorCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.profilePropertySettingCollectionResponse") {
		var out ProfilePropertySettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfilePropertySettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profileSourceAnnotationCollectionResponse") {
		var out ProfileSourceAnnotationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfileSourceAnnotationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profileSourceCollectionResponse") {
		var out ProfileSourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfileSourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profileSourceLocalizationCollectionResponse") {
		var out ProfileSourceLocalizationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfileSourceLocalizationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.programCollectionResponse") {
		var out ProgramCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProgramCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.programControlCollectionResponse") {
		var out ProgramControlCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProgramControlCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.programControlTypeCollectionResponse") {
		var out ProgramControlTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProgramControlTypeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.projectParticipationCollectionResponse") {
		var out ProjectParticipationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProjectParticipationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.propertyCollectionResponse") {
		var out PropertyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PropertyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.providerTenantSettingCollectionResponse") {
		var out ProviderTenantSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProviderTenantSettingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.publicErrorCollectionResponse") {
		var out PublicErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PublicErrorCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.publishedResourceCollectionResponse") {
		var out PublishedResourceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PublishedResourceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.purchaseInvoiceCollectionResponse") {
		var out PurchaseInvoiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PurchaseInvoiceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.purchaseInvoiceLineCollectionResponse") {
		var out PurchaseInvoiceLineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PurchaseInvoiceLineCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.qrCodePinAuthenticationMethodCollectionResponse") {
		var out QrCodePinAuthenticationMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into QrCodePinAuthenticationMethodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.qrCodePinAuthenticationMethodConfigurationCollectionResponse") {
		var out QrCodePinAuthenticationMethodConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into QrCodePinAuthenticationMethodConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rankedEmailAddressCollectionResponse") {
		var out RankedEmailAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RankedEmailAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rbacApplicationCollectionResponse") {
		var out RbacApplicationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RbacApplicationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.readingAssignmentSubmissionCollectionResponse") {
		var out ReadingAssignmentSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ReadingAssignmentSubmissionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.recommendationCollectionResponse") {
		var out RecommendationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecommendationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.redundantAssignmentAlertConfigurationCollectionResponse") {
		var out RedundantAssignmentAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RedundantAssignmentAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.redundantAssignmentAlertIncidentCollectionResponse") {
		var out RedundantAssignmentAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RedundantAssignmentAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.reflectCheckInResponseCollectionResponse") {
		var out ReflectCheckInResponseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ReflectCheckInResponseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.relatedPersonCollectionResponse") {
		var out RelatedPersonCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RelatedPersonCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.remoteActionAuditCollectionResponse") {
		var out RemoteActionAuditCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RemoteActionAuditCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.reputationCategoryCollectionResponse") {
		var out ReputationCategoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ReputationCategoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.requestActivityCollectionResponse") {
		var out RequestActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RequestActivityCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.responsibleSensitiveTypeCollectionResponse") {
		var out ResponsibleSensitiveTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResponsibleSensitiveTypeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.restrictedAppsViolationCollectionResponse") {
		var out RestrictedAppsViolationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestrictedAppsViolationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.retrievalExtractCollectionResponse") {
		var out RetrievalExtractCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RetrievalExtractCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.retrievalHitCollectionResponse") {
		var out RetrievalHitCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RetrievalHitCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.roleScopeTagAutoAssignmentCollectionResponse") {
		var out RoleScopeTagAutoAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleScopeTagAutoAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleScopeTagCollectionResponse") {
		var out RoleScopeTagCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleScopeTagCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleScopeTagInfoCollectionResponse") {
		var out RoleScopeTagInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleScopeTagInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationCollectionResponse") {
		var out RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertIncidentCollectionResponse") {
		var out RolesAssignedOutsidePrivilegedIdentityManagementAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RolesAssignedOutsidePrivilegedIdentityManagementAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.salesCreditMemoCollectionResponse") {
		var out SalesCreditMemoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesCreditMemoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesCreditMemoLineCollectionResponse") {
		var out SalesCreditMemoLineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesCreditMemoLineCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesInvoiceCollectionResponse") {
		var out SalesInvoiceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesInvoiceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesInvoiceLineCollectionResponse") {
		var out SalesInvoiceLineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesInvoiceLineCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesOrderCollectionResponse") {
		var out SalesOrderCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesOrderCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesOrderLineCollectionResponse") {
		var out SalesOrderLineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesOrderLineCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesQuoteCollectionResponse") {
		var out SalesQuoteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesQuoteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.salesQuoteLineCollectionResponse") {
		var out SalesQuoteLineCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SalesQuoteLineCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.scheduledPermissionsRequestCollectionResponse") {
		var out ScheduledPermissionsRequestCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScheduledPermissionsRequestCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.scopeBaseCollectionResponse") {
		var out ScopeBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScopeBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsResourceFindingCollectionResponse") {
		var out SecretInformationAccessAwsResourceFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsResourceFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsRoleFindingCollectionResponse") {
		var out SecretInformationAccessAwsRoleFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsRoleFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsServerlessFunctionFindingCollectionResponse") {
		var out SecretInformationAccessAwsServerlessFunctionFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsServerlessFunctionFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsUserFindingCollectionResponse") {
		var out SecretInformationAccessAwsUserFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsUserFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.securityActionCollectionResponse") {
		var out SecurityActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityActionStateCollectionResponse") {
		var out SecurityActionStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityActionStateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.analyzedEmailAttachmentCollectionResponse") {
		var out SecurityAnalyzedEmailAttachmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAnalyzedEmailAttachmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.analyzedEmailCollectionResponse") {
		var out SecurityAnalyzedEmailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAnalyzedEmailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.analyzedEmailDlpRuleInfoCollectionResponse") {
		var out SecurityAnalyzedEmailDlpRuleInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAnalyzedEmailDlpRuleInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.analyzedEmailExchangeTransportRuleInfoCollectionResponse") {
		var out SecurityAnalyzedEmailExchangeTransportRuleInfoCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAnalyzedEmailExchangeTransportRuleInfoCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.analyzedEmailUrlCollectionResponse") {
		var out SecurityAnalyzedEmailUrlCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAnalyzedEmailUrlCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.auditLogQueryCollectionResponse") {
		var out SecurityAuditLogQueryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAuditLogQueryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.auditLogRecordCollectionResponse") {
		var out SecurityAuditLogRecordCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAuditLogRecordCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineCategoryStateSummaryCollectionResponse") {
		var out SecurityBaselineCategoryStateSummaryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineCategoryStateSummaryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineContributingPolicyCollectionResponse") {
		var out SecurityBaselineContributingPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineContributingPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineDeviceStateCollectionResponse") {
		var out SecurityBaselineDeviceStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineDeviceStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineSettingStateCollectionResponse") {
		var out SecurityBaselineSettingStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineSettingStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineStateCollectionResponse") {
		var out SecurityBaselineStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineStateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineTemplateCollectionResponse") {
		var out SecurityBaselineTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineTemplateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.cloudAppDiscoveryReportCollectionResponse") {
		var out SecurityCloudAppDiscoveryReportCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCloudAppDiscoveryReportCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.compromiseIndicatorCollectionResponse") {
		var out SecurityCompromiseIndicatorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCompromiseIndicatorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityConfigurationTaskCollectionResponse") {
		var out SecurityConfigurationTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityConfigurationTaskCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.detectionRuleCollectionResponse") {
		var out SecurityDetectionRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDetectionRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.detonationChainCollectionResponse") {
		var out SecurityDetonationChainCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDetonationChainCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.discoveredCloudAppDeviceCollectionResponse") {
		var out SecurityDiscoveredCloudAppDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDiscoveredCloudAppDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.discoveredCloudAppIPAddressCollectionResponse") {
		var out SecurityDiscoveredCloudAppIPAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDiscoveredCloudAppIPAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.discoveredCloudAppUserCollectionResponse") {
		var out SecurityDiscoveredCloudAppUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDiscoveredCloudAppUserCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryCaseMemberCollectionResponse") {
		var out SecurityEdiscoveryCaseMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryCaseMemberCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryFileCollectionResponse") {
		var out SecurityEdiscoveryFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryFileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryHoldPolicyCollectionResponse") {
		var out SecurityEdiscoveryHoldPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryHoldPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.emailContentThreatSubmissionCollectionResponse") {
		var out SecurityEmailContentThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailContentThreatSubmissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.emailThreatSubmissionCollectionResponse") {
		var out SecurityEmailThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailThreatSubmissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.emailThreatSubmissionPolicyCollectionResponse") {
		var out SecurityEmailThreatSubmissionPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailThreatSubmissionPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.emailUrlThreatSubmissionCollectionResponse") {
		var out SecurityEmailUrlThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailUrlThreatSubmissionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.fileContentThreatSubmissionCollectionResponse") {
		var out SecurityFileContentThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileContentThreatSubmissionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.fileThreatSubmissionCollectionResponse") {
		var out SecurityFileThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileThreatSubmissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fileUrlThreatSubmissionCollectionResponse") {
		var out SecurityFileUrlThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileUrlThreatSubmissionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.impactedAssetCollectionResponse") {
		var out SecurityImpactedAssetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityImpactedAssetCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.informationProtectionActionCollectionResponse") {
		var out SecurityInformationProtectionActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityInformationProtectionActionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.keyValuePairCollectionResponse") {
		var out SecurityKeyValuePairCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityKeyValuePairCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.recommendedHuntingQueryCollectionResponse") {
		var out SecurityRecommendedHuntingQueryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRecommendedHuntingQueryCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.responseActionCollectionResponse") {
		var out SecurityResponseActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityResponseActionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.sensitivityLabelCollectionResponse") {
		var out SecuritySensitivityLabelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySensitivityLabelCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.submissionDetectedFileCollectionResponse") {
		var out SecuritySubmissionDetectedFileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySubmissionDetectedFileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.tenantAllowBlockListEntryResultCollectionResponse") {
		var out SecurityTenantAllowBlockListEntryResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTenantAllowBlockListEntryResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.threatDetectionDetailCollectionResponse") {
		var out SecurityThreatDetectionDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityThreatDetectionDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.timelineEventCollectionResponse") {
		var out SecurityTimelineEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTimelineEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsResourceAdministratorFindingCollectionResponse") {
		var out SecurityToolAwsResourceAdministratorFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsResourceAdministratorFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsRoleAdministratorFindingCollectionResponse") {
		var out SecurityToolAwsRoleAdministratorFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsRoleAdministratorFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsServerlessFunctionAdministratorFindingCollectionResponse") {
		var out SecurityToolAwsServerlessFunctionAdministratorFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsServerlessFunctionAdministratorFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsUserAdministratorFindingCollectionResponse") {
		var out SecurityToolAwsUserAdministratorFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsUserAdministratorFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.security.urlThreatSubmissionCollectionResponse") {
		var out SecurityUrlThreatSubmissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUrlThreatSubmissionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.selfServiceSignUpCollectionResponse") {
		var out SelfServiceSignUpCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelfServiceSignUpCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sensitiveContentEvidenceCollectionResponse") {
		var out SensitiveContentEvidenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SensitiveContentEvidenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sensitiveContentLocationCollectionResponse") {
		var out SensitiveContentLocationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SensitiveContentLocationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sensitiveTypeCollectionResponse") {
		var out SensitiveTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SensitiveTypeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sensitivityLabelCollectionResponse") {
		var out SensitivityLabelCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SensitivityLabelCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sequentialActivationRenewalsAlertConfigurationCollectionResponse") {
		var out SequentialActivationRenewalsAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SequentialActivationRenewalsAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sequentialActivationRenewalsAlertIncidentCollectionResponse") {
		var out SequentialActivationRenewalsAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SequentialActivationRenewalsAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.serviceLevelAgreementAttainmentCollectionResponse") {
		var out ServiceLevelAgreementAttainmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceLevelAgreementAttainmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceNowConnectionCollectionResponse") {
		var out ServiceNowConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceNowConnectionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalCreationConditionSetCollectionResponse") {
		var out ServicePrincipalCreationConditionSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalCreationConditionSetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalCreationPolicyCollectionResponse") {
		var out ServicePrincipalCreationPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalCreationPolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalSignInActivityCollectionResponse") {
		var out ServicePrincipalSignInActivityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalSignInActivityCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.serviceProvisioningResourceErrorDetailCollectionResponse") {
		var out ServiceProvisioningResourceErrorDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceProvisioningResourceErrorDetailCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sessionLifetimePolicyCollectionResponse") {
		var out SessionLifetimePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SessionLifetimePolicyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sharePointGroupCollectionResponse") {
		var out SharePointGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointGroupMemberCollectionResponse") {
		var out SharePointGroupMemberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointGroupMemberCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sharedAppleDeviceUserCollectionResponse") {
		var out SharedAppleDeviceUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedAppleDeviceUserCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sharedEmailDomainCollectionResponse") {
		var out SharedEmailDomainCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedEmailDomainCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedEmailDomainInvitationCollectionResponse") {
		var out SharedEmailDomainInvitationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedEmailDomainInvitationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.shiftsRoleDefinitionCollectionResponse") {
		var out ShiftsRoleDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftsRoleDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shiftsRolePermissionCollectionResponse") {
		var out ShiftsRolePermissionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftsRolePermissionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shipmentMethodCollectionResponse") {
		var out ShipmentMethodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShipmentMethodCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.skillProficiencyCollectionResponse") {
		var out SkillProficiencyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SkillProficiencyCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.sourcedAttributeCollectionResponse") {
		var out SourcedAttributeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SourcedAttributeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.staleSignInAlertConfigurationCollectionResponse") {
		var out StaleSignInAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StaleSignInAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.staleSignInAlertIncidentCollectionResponse") {
		var out StaleSignInAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StaleSignInAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.startRecordingOperationCollectionResponse") {
		var out StartRecordingOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StartRecordingOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.startTranscriptionOperationCollectionResponse") {
		var out StartTranscriptionOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StartTranscriptionOperationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.stopRecordingOperationCollectionResponse") {
		var out StopRecordingOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StopRecordingOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stopTranscriptionOperationCollectionResponse") {
		var out StopTranscriptionOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StopTranscriptionOperationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.superAwsResourceFindingCollectionResponse") {
		var out SuperAwsResourceFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperAwsResourceFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superAwsRoleFindingCollectionResponse") {
		var out SuperAwsRoleFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperAwsRoleFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superAzureServicePrincipalFindingCollectionResponse") {
		var out SuperAzureServicePrincipalFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperAzureServicePrincipalFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superGcpServiceAccountFindingCollectionResponse") {
		var out SuperGcpServiceAccountFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperGcpServiceAccountFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superServerlessFunctionFindingCollectionResponse") {
		var out SuperServerlessFunctionFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperServerlessFunctionFindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superUserFindingCollectionResponse") {
		var out SuperUserFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperUserFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppConfigurationPolicySetItemCollectionResponse") {
		var out TargetedManagedAppConfigurationPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppConfigurationPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.taxAreaCollectionResponse") {
		var out TaxAreaCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TaxAreaCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.taxGroupCollectionResponse") {
		var out TaxGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TaxGroupCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.teamTemplateCollectionResponse") {
		var out TeamTemplateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamTemplateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamTemplateDefinitionCollectionResponse") {
		var out TeamTemplateDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamTemplateDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAdministration.assignedTelephoneNumberCollectionResponse") {
		var out TeamsAdministrationAssignedTelephoneNumberCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAdministrationAssignedTelephoneNumberCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAdministration.effectivePolicyAssignmentCollectionResponse") {
		var out TeamsAdministrationEffectivePolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAdministrationEffectivePolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAdministration.teamsUserConfigurationCollectionResponse") {
		var out TeamsAdministrationTeamsUserConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAdministrationTeamsUserConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.teamsAppDashboardCardDefinitionCollectionResponse") {
		var out TeamsAppDashboardCardDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppDashboardCardDefinitionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.teamworkConfiguredPeripheralCollectionResponse") {
		var out TeamworkConfiguredPeripheralCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkConfiguredPeripheralCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkDeviceCollectionResponse") {
		var out TeamworkDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkDeviceOperationCollectionResponse") {
		var out TeamworkDeviceOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkDeviceOperationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkPeripheralCollectionResponse") {
		var out TeamworkPeripheralCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkPeripheralCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkPeripheralHealthCollectionResponse") {
		var out TeamworkPeripheralHealthCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkPeripheralHealthCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.tenantReferenceCollectionResponse") {
		var out TenantReferenceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TenantReferenceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tenantRelationshipAccessPolicyBaseCollectionResponse") {
		var out TenantRelationshipAccessPolicyBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TenantRelationshipAccessPolicyBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditionsGroupAssignmentCollectionResponse") {
		var out TermsAndConditionsGroupAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditionsGroupAssignmentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.tiIndicatorCollectionResponse") {
		var out TiIndicatorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TiIndicatorCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertConfigurationCollectionResponse") {
		var out TooManyGlobalAdminsAssignedToTenantAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TooManyGlobalAdminsAssignedToTenantAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertIncidentCollectionResponse") {
		var out TooManyGlobalAdminsAssignedToTenantAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TooManyGlobalAdminsAssignedToTenantAlertIncidentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trainingCampaignCollectionResponse") {
		var out TrainingCampaignCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrainingCampaignCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.translationLanguageOverrideCollectionResponse") {
		var out TranslationLanguageOverrideCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TranslationLanguageOverrideCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.trustFrameworkKeyCollectionResponse") {
		var out TrustFrameworkKeyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustFrameworkKeyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustFrameworkKeySetCollectionResponse") {
		var out TrustFrameworkKeySetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustFrameworkKeySetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustFrameworkKey_v2CollectionResponse") {
		var out TrustFrameworkKeyv2CollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustFrameworkKeyv2CollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustFrameworkPolicyCollectionResponse") {
		var out TrustFrameworkPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustFrameworkPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustedCertificateAuthorityAsEntityBaseCollectionResponse") {
		var out TrustedCertificateAuthorityAsEntityBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustedCertificateAuthorityAsEntityBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustedCertificateAuthorityBaseCollectionResponse") {
		var out TrustedCertificateAuthorityBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustedCertificateAuthorityBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.typedEmailAddressCollectionResponse") {
		var out TypedEmailAddressCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TypedEmailAddressCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unenforcedMfaAwsUserFindingCollectionResponse") {
		var out UnenforcedMfaAwsUserFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnenforcedMfaAwsUserFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignmentMultipleCollectionResponse") {
		var out UnifiedRoleAssignmentMultipleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignmentMultipleCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementAlertCollectionResponse") {
		var out UnifiedRoleManagementAlertCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementAlertCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementAlertConfigurationCollectionResponse") {
		var out UnifiedRoleManagementAlertConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementAlertConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementAlertDefinitionCollectionResponse") {
		var out UnifiedRoleManagementAlertDefinitionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementAlertDefinitionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementAlertIncidentCollectionResponse") {
		var out UnifiedRoleManagementAlertIncidentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementAlertIncidentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.unitOfMeasureCollectionResponse") {
		var out UnitOfMeasureCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnitOfMeasureCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unmanagedDeviceCollectionResponse") {
		var out UnmanagedDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnmanagedDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unmanagedDeviceDiscoveryTaskCollectionResponse") {
		var out UnmanagedDeviceDiscoveryTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnmanagedDeviceDiscoveryTaskCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.unsupportedDeviceConfigurationCollectionResponse") {
		var out UnsupportedDeviceConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnsupportedDeviceConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unsupportedDeviceConfigurationDetailCollectionResponse") {
		var out UnsupportedDeviceConfigurationDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnsupportedDeviceConfigurationDetailCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unsupportedGroupPolicyExtensionCollectionResponse") {
		var out UnsupportedGroupPolicyExtensionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnsupportedGroupPolicyExtensionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.updateAllMessagesReadStateOperationCollectionResponse") {
		var out UpdateAllMessagesReadStateOperationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UpdateAllMessagesReadStateOperationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.usageRightCollectionResponse") {
		var out UsageRightCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UsageRightCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userAccountCollectionResponse") {
		var out UserAccountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserAccountCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userAccountInformationCollectionResponse") {
		var out UserAccountInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserAccountInformationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userConfigurationCollectionResponse") {
		var out UserConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userCountMetricCollectionResponse") {
		var out UserCountMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserCountMetricCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userCredentialUsageDetailsCollectionResponse") {
		var out UserCredentialUsageDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserCredentialUsageDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAnomalyCollectionResponse") {
		var out UserExperienceAnalyticsAnomalyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAnomalyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAnomalyCorrelationGroupFeatureCollectionResponse") {
		var out UserExperienceAnalyticsAnomalyCorrelationGroupFeatureCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAnomalyCorrelationGroupFeatureCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse") {
		var out UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAnomalyDeviceCollectionResponse") {
		var out UserExperienceAnalyticsAnomalyDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAnomalyDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionCollectionResponse") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBatteryHealthAppImpactCollectionResponse") {
		var out UserExperienceAnalyticsBatteryHealthAppImpactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBatteryHealthAppImpactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBatteryHealthDeviceAppImpactCollectionResponse") {
		var out UserExperienceAnalyticsBatteryHealthDeviceAppImpactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBatteryHealthDeviceAppImpactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponse") {
		var out UserExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponse") {
		var out UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBatteryHealthModelPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsBatteryHealthModelPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBatteryHealthModelPerformanceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBatteryHealthOsPerformanceCollectionResponse") {
		var out UserExperienceAnalyticsBatteryHealthOsPerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBatteryHealthOsPerformanceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceBatteryDetailCollectionResponse") {
		var out UserExperienceAnalyticsDeviceBatteryDetailCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceBatteryDetailCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceScopeCollectionResponse") {
		var out UserExperienceAnalyticsDeviceScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceScopeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceTimelineEventCollectionResponse") {
		var out UserExperienceAnalyticsDeviceTimelineEventCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceTimelineEventCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponse") {
		var out UserExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsImpactingProcessCollectionResponse") {
		var out UserExperienceAnalyticsImpactingProcessCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsImpactingProcessCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponse") {
		var out UserExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsRemoteConnectionCollectionResponse") {
		var out UserExperienceAnalyticsRemoteConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsRemoteConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsResourcePerformanceCollectionResponse") {
		var out UserExperienceAnalyticsResourcePerformanceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsResourcePerformanceCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userIdentityCollectionResponse") {
		var out UserIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserIdentityCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userPFXCertificateCollectionResponse") {
		var out UserPFXCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserPFXCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userRegistrationCountCollectionResponse") {
		var out UserRegistrationCountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserRegistrationCountCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userRequestsMetricCollectionResponse") {
		var out UserRequestsMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserRequestsMetricCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userSecurityProfileCollectionResponse") {
		var out UserSecurityProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSecurityProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userSetCollectionResponse") {
		var out UserSetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSetCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.userSignUpMetricCollectionResponse") {
		var out UserSignUpMetricCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSignUpMetricCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.validationResultCollectionResponse") {
		var out ValidationResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValidationResultCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vendorCollectionResponse") {
		var out VendorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VendorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiableCredentialClaimBindingCollectionResponse") {
		var out VerifiableCredentialClaimBindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiableCredentialClaimBindingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiableCredentialTypeCollectionResponse") {
		var out VerifiableCredentialTypeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiableCredentialTypeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiedCredentialDataCollectionResponse") {
		var out VerifiedCredentialDataCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiedCredentialDataCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.videoNewsLinkPageCollectionResponse") {
		var out VideoNewsLinkPageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VideoNewsLinkPageCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.virtualMachineDetailsCollectionResponse") {
		var out VirtualMachineDetailsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineDetailsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualMachineWithAwsStorageBucketAccessFindingCollectionResponse") {
		var out VirtualMachineWithAwsStorageBucketAccessFindingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineWithAwsStorageBucketAccessFindingCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.voiceAuthenticationMethodTargetCollectionResponse") {
		var out VoiceAuthenticationMethodTargetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VoiceAuthenticationMethodTargetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vpnConfigurationCollectionResponse") {
		var out VpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vpnDnsRuleCollectionResponse") {
		var out VpnDnsRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VpnDnsRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vpnOnDemandRuleCollectionResponse") {
		var out VpnOnDemandRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VpnOnDemandRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vpnRouteCollectionResponse") {
		var out VpnRouteCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VpnRouteCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vpnServerCollectionResponse") {
		var out VpnServerCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VpnServerCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vpnTrafficRuleCollectionResponse") {
		var out VpnTrafficRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VpnTrafficRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vppTokenActionResultCollectionResponse") {
		var out VppTokenActionResultCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VppTokenActionResultCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.vulnerableManagedDeviceCollectionResponse") {
		var out VulnerableManagedDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VulnerableManagedDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webAccountCollectionResponse") {
		var out WebAccountCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebAccountCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.webApplicationSegmentCollectionResponse") {
		var out WebApplicationSegmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebApplicationSegmentCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.webauthnPublicKeyCredentialDescriptorCollectionResponse") {
		var out WebauthnPublicKeyCredentialDescriptorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebauthnPublicKeyCredentialDescriptorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webauthnPublicKeyCredentialParametersCollectionResponse") {
		var out WebauthnPublicKeyCredentialParametersCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebauthnPublicKeyCredentialParametersCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.win32CatalogAppCollectionResponse") {
		var out Win32CatalogAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32CatalogAppCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppDetectionCollectionResponse") {
		var out Win32LobAppDetectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppDetectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppRequirementCollectionResponse") {
		var out Win32LobAppRequirementCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppRequirementCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.win32MobileAppCatalogPackageCollectionResponse") {
		var out Win32MobileAppCatalogPackageCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32MobileAppCatalogPackageCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.winGetAppCollectionResponse") {
		var out WinGetAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WinGetAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10AssociatedAppsCollectionResponse") {
		var out Windows10AssociatedAppsCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10AssociatedAppsCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10CertificateProfileBaseCollectionResponse") {
		var out Windows10CertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10CertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windows10DeviceFirmwareConfigurationInterfaceCollectionResponse") {
		var out Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EasEmailProfileConfigurationCollectionResponse") {
		var out Windows10EasEmailProfileConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EasEmailProfileConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse") {
		var out Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windows10ImportedPFXCertificateProfileCollectionResponse") {
		var out Windows10ImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10ImportedPFXCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windows10NetworkBoundaryConfigurationCollectionResponse") {
		var out Windows10NetworkBoundaryConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10NetworkBoundaryConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10PFXImportCertificateProfileCollectionResponse") {
		var out Windows10PFXImportCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10PFXImportCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10PkcsCertificateProfileCollectionResponse") {
		var out Windows10PkcsCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10PkcsCertificateProfileCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windows10VpnConfigurationCollectionResponse") {
		var out Windows10VpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10VpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XCertificateProfileCollectionResponse") {
		var out Windows10XCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XCustomSubjectAlternativeNameCollectionResponse") {
		var out Windows10XCustomSubjectAlternativeNameCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XCustomSubjectAlternativeNameCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XSCEPCertificateProfileCollectionResponse") {
		var out Windows10XSCEPCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XSCEPCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XTrustedRootCertificateCollectionResponse") {
		var out Windows10XTrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XTrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XVpnConfigurationCollectionResponse") {
		var out Windows10XVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XVpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XWifiConfigurationCollectionResponse") {
		var out Windows10XWifiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XWifiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81CertificateProfileBaseCollectionResponse") {
		var out Windows81CertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81CertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windows81SCEPCertificateProfileCollectionResponse") {
		var out Windows81SCEPCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81SCEPCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81TrustedRootCertificateCollectionResponse") {
		var out Windows81TrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81TrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81VpnConfigurationCollectionResponse") {
		var out Windows81VpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81VpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81WifiImportConfigurationCollectionResponse") {
		var out Windows81WifiImportConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81WifiImportConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeploymentProfileAssignmentCollectionResponse") {
		var out WindowsAutopilotDeploymentProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeploymentProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeploymentProfileCollectionResponse") {
		var out WindowsAutopilotDeploymentProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeploymentProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeploymentProfilePolicySetItemCollectionResponse") {
		var out WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsCertificateProfileBaseCollectionResponse") {
		var out WindowsCertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsCertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse") {
		var out WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyCollectionResponse") {
		var out WindowsDefenderApplicationControlSupplementalPolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDefenderApplicationControlSupplementalPolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponse") {
		var out WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDeliveryOptimizationConfigurationCollectionResponse") {
		var out WindowsDeliveryOptimizationConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDeliveryOptimizationConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsDomainJoinConfigurationCollectionResponse") {
		var out WindowsDomainJoinConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDomainJoinConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDriverUpdateInventoryCollectionResponse") {
		var out WindowsDriverUpdateInventoryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDriverUpdateInventoryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDriverUpdateProfileAssignmentCollectionResponse") {
		var out WindowsDriverUpdateProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDriverUpdateProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDriverUpdateProfileCollectionResponse") {
		var out WindowsDriverUpdateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDriverUpdateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsFeatureUpdateCatalogItemCollectionResponse") {
		var out WindowsFeatureUpdateCatalogItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsFeatureUpdateCatalogItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsFeatureUpdateProfileAssignmentCollectionResponse") {
		var out WindowsFeatureUpdateProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsFeatureUpdateProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsFeatureUpdateProfileCollectionResponse") {
		var out WindowsFeatureUpdateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsFeatureUpdateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsFirewallRuleCollectionResponse") {
		var out WindowsFirewallRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsFirewallRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsHealthMonitoringConfigurationCollectionResponse") {
		var out WindowsHealthMonitoringConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsHealthMonitoringConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsIdentityProtectionConfigurationCollectionResponse") {
		var out WindowsIdentityProtectionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsIdentityProtectionConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionDeviceRegistrationCollectionResponse") {
		var out WindowsInformationProtectionDeviceRegistrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionDeviceRegistrationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionWipeActionCollectionResponse") {
		var out WindowsInformationProtectionWipeActionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionWipeActionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskAppBaseCollectionResponse") {
		var out WindowsKioskAppBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskAppBaseCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskConfigurationCollectionResponse") {
		var out WindowsKioskConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskProfileCollectionResponse") {
		var out WindowsKioskProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskUserCollectionResponse") {
		var out WindowsKioskUserCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskUserCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsManagedAppProtectionCollectionResponse") {
		var out WindowsManagedAppProtectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsManagedAppProtectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsManagedAppRegistrationCollectionResponse") {
		var out WindowsManagedAppRegistrationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsManagedAppRegistrationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsManagedDeviceCollectionResponse") {
		var out WindowsManagedDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsManagedDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsManagementAppHealthStateCollectionResponse") {
		var out WindowsManagementAppHealthStateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsManagementAppHealthStateCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsPackageInformationCollectionResponse") {
		var out WindowsPackageInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPackageInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81AppXBundleCollectionResponse") {
		var out WindowsPhone81AppXBundleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81AppXBundleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81AppXCollectionResponse") {
		var out WindowsPhone81AppXCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81AppXCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81CertificateProfileBaseCollectionResponse") {
		var out WindowsPhone81CertificateProfileBaseCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81CertificateProfileBaseCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81ImportedPFXCertificateProfileCollectionResponse") {
		var out WindowsPhone81ImportedPFXCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81ImportedPFXCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81SCEPCertificateProfileCollectionResponse") {
		var out WindowsPhone81SCEPCertificateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81SCEPCertificateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81StoreAppCollectionResponse") {
		var out WindowsPhone81StoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81StoreAppCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81TrustedRootCertificateCollectionResponse") {
		var out WindowsPhone81TrustedRootCertificateCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81TrustedRootCertificateCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81VpnConfigurationCollectionResponse") {
		var out WindowsPhone81VpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81VpnConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhoneEASEmailProfileConfigurationCollectionResponse") {
		var out WindowsPhoneEASEmailProfileConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhoneEASEmailProfileConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhoneXAPCollectionResponse") {
		var out WindowsPhoneXAPCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhoneXAPCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPrivacyDataAccessControlItemCollectionResponse") {
		var out WindowsPrivacyDataAccessControlItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPrivacyDataAccessControlItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdateApprovalSettingCollectionResponse") {
		var out WindowsQualityUpdateApprovalSettingCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdateApprovalSettingCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdateCatalogItemCollectionResponse") {
		var out WindowsQualityUpdateCatalogItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdateCatalogItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdateCatalogProductRevisionCollectionResponse") {
		var out WindowsQualityUpdateCatalogProductRevisionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdateCatalogProductRevisionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdatePolicyAssignmentCollectionResponse") {
		var out WindowsQualityUpdatePolicyAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdatePolicyAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdatePolicyCollectionResponse") {
		var out WindowsQualityUpdatePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdatePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdateProfileAssignmentCollectionResponse") {
		var out WindowsQualityUpdateProfileAssignmentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdateProfileAssignmentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdateProfileCollectionResponse") {
		var out WindowsQualityUpdateProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdateProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsRestoreDeviceEnrollmentConfigurationCollectionResponse") {
		var out WindowsRestoreDeviceEnrollmentConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsRestoreDeviceEnrollmentConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsStoreAppCollectionResponse") {
		var out WindowsStoreAppCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsStoreAppCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdateCatalogItemCollectionResponse") {
		var out WindowsUpdateCatalogItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdateCatalogItemCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.applicableContentCollectionResponse") {
		var out WindowsUpdatesApplicableContentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesApplicableContentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.applicableContentDeviceMatchCollectionResponse") {
		var out WindowsUpdatesApplicableContentDeviceMatchCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesApplicableContentDeviceMatchCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.assignedGroupCollectionResponse") {
		var out WindowsUpdatesAssignedGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesAssignedGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.azureADDeviceCollectionResponse") {
		var out WindowsUpdatesAzureADDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesAzureADDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.catalogEntryCollectionResponse") {
		var out WindowsUpdatesCatalogEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesCatalogEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.complianceChangeCollectionResponse") {
		var out WindowsUpdatesComplianceChangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesComplianceChangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.complianceChangeRuleCollectionResponse") {
		var out WindowsUpdatesComplianceChangeRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesComplianceChangeRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.contentApprovalCollectionResponse") {
		var out WindowsUpdatesContentApprovalCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesContentApprovalCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.cveInformationCollectionResponse") {
		var out WindowsUpdatesCveInformationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesCveInformationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.deploymentAudienceCollectionResponse") {
		var out WindowsUpdatesDeploymentAudienceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDeploymentAudienceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.deploymentCollectionResponse") {
		var out WindowsUpdatesDeploymentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDeploymentCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.deploymentStateReasonCollectionResponse") {
		var out WindowsUpdatesDeploymentStateReasonCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDeploymentStateReasonCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.driverUpdateCatalogEntryCollectionResponse") {
		var out WindowsUpdatesDriverUpdateCatalogEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDriverUpdateCatalogEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.editionCollectionResponse") {
		var out WindowsUpdatesEditionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesEditionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.featureUpdateCatalogEntryCollectionResponse") {
		var out WindowsUpdatesFeatureUpdateCatalogEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesFeatureUpdateCatalogEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.knownIssueCollectionResponse") {
		var out WindowsUpdatesKnownIssueCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesKnownIssueCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.knownIssueHistoryItemCollectionResponse") {
		var out WindowsUpdatesKnownIssueHistoryItemCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesKnownIssueHistoryItemCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.monitoringRuleCollectionResponse") {
		var out WindowsUpdatesMonitoringRuleCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesMonitoringRuleCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.operationalInsightsConnectionCollectionResponse") {
		var out WindowsUpdatesOperationalInsightsConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesOperationalInsightsConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.productCollectionResponse") {
		var out WindowsUpdatesProductCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesProductCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.productRevisionCollectionResponse") {
		var out WindowsUpdatesProductRevisionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesProductRevisionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntryCollectionResponse") {
		var out WindowsUpdatesQualityUpdateCatalogEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesQualityUpdateCatalogEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.resourceConnectionCollectionResponse") {
		var out WindowsUpdatesResourceConnectionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesResourceConnectionCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.safeguardProfileCollectionResponse") {
		var out WindowsUpdatesSafeguardProfileCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesSafeguardProfileCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.servicingPeriodCollectionResponse") {
		var out WindowsUpdatesServicingPeriodCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesServicingPeriodCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.softwareUpdateCatalogEntryCollectionResponse") {
		var out WindowsUpdatesSoftwareUpdateCatalogEntryCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesSoftwareUpdateCatalogEntryCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.updatableAssetCollectionResponse") {
		var out WindowsUpdatesUpdatableAssetCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesUpdatableAssetCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.updatableAssetErrorCollectionResponse") {
		var out WindowsUpdatesUpdatableAssetErrorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesUpdatableAssetErrorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.updatableAssetGroupCollectionResponse") {
		var out WindowsUpdatesUpdatableAssetGroupCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesUpdatableAssetGroupCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.updatePolicyCollectionResponse") {
		var out WindowsUpdatesUpdatePolicyCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesUpdatePolicyCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsVpnConfigurationCollectionResponse") {
		var out WindowsVpnConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsVpnConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.windowsWifiConfigurationCollectionResponse") {
		var out WindowsWifiConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsWifiConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsWifiEnterpriseEAPConfigurationCollectionResponse") {
		var out WindowsWifiEnterpriseEAPConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsWifiEnterpriseEAPConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsWiredNetworkConfigurationCollectionResponse") {
		var out WindowsWiredNetworkConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsWiredNetworkConfigurationCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workPositionCollectionResponse") {
		var out WorkPositionCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkPositionCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.workbookDocumentTaskChangeCollectionResponse") {
		var out WorkbookDocumentTaskChangeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookDocumentTaskChangeCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookDocumentTaskCollectionResponse") {
		var out WorkbookDocumentTaskCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookDocumentTaskCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookEmailIdentityCollectionResponse") {
		var out WorkbookEmailIdentityCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookEmailIdentityCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.workplaceSensorCollectionResponse") {
		var out WorkplaceSensorCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkplaceSensorCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workplaceSensorDeviceCollectionResponse") {
		var out WorkplaceSensorDeviceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkplaceSensorDeviceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workspaceCollectionResponse") {
		var out WorkspaceCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkspaceCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.wslDistributionConfigurationCollectionResponse") {
		var out WslDistributionConfigurationCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WslDistributionConfigurationCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateAuthorityScopeCollectionResponse") {
		var out X509CertificateAuthorityScopeCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateAuthorityScopeCollectionResponse: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.zebraFotaArtifactCollectionResponse") {
		var out ZebraFotaArtifactCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ZebraFotaArtifactCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.zebraFotaDeploymentCollectionResponse") {
		var out ZebraFotaDeploymentCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ZebraFotaDeploymentCollectionResponse: %+v", err)
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
