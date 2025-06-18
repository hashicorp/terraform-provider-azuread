package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Entity interface {
	Entity() BaseEntityImpl
}

var _ Entity = BaseEntityImpl{}

type BaseEntityImpl struct {
	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEntityImpl) Entity() BaseEntityImpl {
	return s
}

var _ Entity = RawEntityImpl{}

// RawEntityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntityImpl struct {
	entity BaseEntityImpl
	Type   string
	Values map[string]interface{}
}

func (s RawEntityImpl) Entity() BaseEntityImpl {
	return s.entity
}

var _ json.Marshaler = BaseEntityImpl{}

func (s BaseEntityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEntityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEntityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEntityImpl: %+v", err)
	}

	delete(decoded, "id")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEntityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalEntityImplementation(input []byte) (Entity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Entity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackage") {
		var out AccessPackage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignment") {
		var out AccessPackageAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentPolicy") {
		var out AccessPackageAssignmentPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentRequest") {
		var out AccessPackageAssignmentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageCatalog") {
		var out AccessPackageCatalog
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageCatalog: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageQuestion") {
		var out AccessPackageQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageQuestion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResource") {
		var out AccessPackageResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceEnvironment") {
		var out AccessPackageResourceEnvironment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceEnvironment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceRequest") {
		var out AccessPackageResourceRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceRole") {
		var out AccessPackageResourceRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceRole: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceRoleScope") {
		var out AccessPackageResourceRoleScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceRoleScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceScope") {
		var out AccessPackageResourceScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageSubject") {
		var out AccessPackageSubject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageSubject: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewHistoryDefinition") {
		var out AccessReviewHistoryDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewHistoryDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewHistoryInstance") {
		var out AccessReviewHistoryInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewHistoryInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstance") {
		var out AccessReviewInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItem") {
		var out AccessReviewInstanceDecisionItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewReviewer") {
		var out AccessReviewReviewer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewReviewer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewScheduleDefinition") {
		var out AccessReviewScheduleDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewScheduleDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewSet") {
		var out AccessReviewSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewStage") {
		var out AccessReviewStage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewStage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.activityHistoryItem") {
		var out ActivityHistoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityHistoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.adminConsentRequestPolicy") {
		var out AdminConsentRequestPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AdminConsentRequestPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.adminMicrosoft365Apps") {
		var out AdminMicrosoft365Apps
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AdminMicrosoft365Apps: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.adminReportSettings") {
		var out AdminReportSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AdminReportSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreement") {
		var out Agreement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Agreement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementAcceptance") {
		var out AgreementAcceptance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementAcceptance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementFileProperties") {
		var out AgreementFileProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementFileProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteraction") {
		var out AiInteraction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteraction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionAttachment") {
		var out AiInteractionAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionAttachment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionContext") {
		var out AiInteractionContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionContext: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionHistory") {
		var out AiInteractionHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionLink") {
		var out AiInteractionLink
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionLink: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionMention") {
		var out AiInteractionMention
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionMention: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.aiUser") {
		var out AiUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.alert") {
		var out Alert
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Alert: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.allowedValue") {
		var out AllowedValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllowedValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appCatalogs") {
		var out AppCatalogs
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppCatalogs: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appConsentApprovalRoute") {
		var out AppConsentApprovalRoute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppConsentApprovalRoute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appConsentRequest") {
		var out AppConsentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppConsentRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appLogCollectionRequest") {
		var out AppLogCollectionRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppLogCollectionRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appScope") {
		var out AppScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.applePushNotificationCertificate") {
		var out ApplePushNotificationCertificate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplePushNotificationCertificate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.applicationTemplate") {
		var out ApplicationTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approval") {
		var out Approval
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Approval: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalStage") {
		var out ApprovalStage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalStage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attachment") {
		var out Attachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Attachment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attachmentBase") {
		var out AttachmentBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttachmentBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attachmentSession") {
		var out AttachmentSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttachmentSession: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attackSimulationRoot") {
		var out AttackSimulationRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttackSimulationRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attendanceRecord") {
		var out AttendanceRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttendanceRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeMappingFunctionSchema") {
		var out AttributeMappingFunctionSchema
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeMappingFunctionSchema: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.attributeSet") {
		var out AttributeSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.audioRoutingGroup") {
		var out AudioRoutingGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AudioRoutingGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.auditEvent") {
		var out AuditEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuditEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.auditLogRoot") {
		var out AuditLogRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuditLogRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authentication") {
		var out Authentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Authentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationCombinationConfiguration") {
		var out AuthenticationCombinationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationCombinationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationContextClassReference") {
		var out AuthenticationContextClassReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationContextClassReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationEventListener") {
		var out AuthenticationEventListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationEventListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationEventsFlow") {
		var out AuthenticationEventsFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationEventsFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationFlowsPolicy") {
		var out AuthenticationFlowsPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationFlowsPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethod") {
		var out AuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodConfiguration") {
		var out AuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodModeDetail") {
		var out AuthenticationMethodModeDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodModeDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodTarget") {
		var out AuthenticationMethodTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodsPolicy") {
		var out AuthenticationMethodsPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodsPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationMethodsRoot") {
		var out AuthenticationMethodsRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationMethodsRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationStrengthPolicy") {
		var out AuthenticationStrengthPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationStrengthPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authenticationStrengthRoot") {
		var out AuthenticationStrengthRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthenticationStrengthRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authoredNote") {
		var out AuthoredNote
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthoredNote: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.backupRestoreRoot") {
		var out BackupRestoreRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BackupRestoreRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.baseItem") {
		var out BaseItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BaseItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.baseItemVersion") {
		var out BaseItemVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BaseItemVersion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bitlocker") {
		var out Bitlocker
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Bitlocker: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bitlockerRecoveryKey") {
		var out BitlockerRecoveryKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BitlockerRecoveryKey: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingAppointment") {
		var out BookingAppointment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingAppointment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingBusiness") {
		var out BookingBusiness
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingBusiness: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCurrency") {
		var out BookingCurrency
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCurrency: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomQuestion") {
		var out BookingCustomQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomQuestion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomerBase") {
		var out BookingCustomerBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomerBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingService") {
		var out BookingService
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingService: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingStaffMemberBase") {
		var out BookingStaffMemberBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingStaffMemberBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSharedCookie") {
		var out BrowserSharedCookie
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSharedCookie: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSite") {
		var out BrowserSite
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSite: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.browserSiteList") {
		var out BrowserSiteList
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrowserSiteList: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bulkUpload") {
		var out BulkUpload
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BulkUpload: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendar") {
		var out Calendar
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Calendar: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarGroup") {
		var out CalendarGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.calendarPermission") {
		var out CalendarPermission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CalendarPermission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.call") {
		var out Call
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Call: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecording") {
		var out CallRecording
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecording: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.callRecord") {
		var out CallRecordsCallRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsCallRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.participantBase") {
		var out CallRecordsParticipantBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsParticipantBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.segment") {
		var out CallRecordsSegment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsSegment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.session") {
		var out CallRecordsSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsSession: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callTranscript") {
		var out CallTranscript
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallTranscript: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.canvasLayout") {
		var out CanvasLayout
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CanvasLayout: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateBasedAuthConfiguration") {
		var out CertificateBasedAuthConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateBasedAuthConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.changeTrackedEntity") {
		var out ChangeTrackedEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChangeTrackedEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channel") {
		var out Channel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Channel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chat") {
		var out Chat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Chat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessage") {
		var out ChatMessage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageInfo") {
		var out ChatMessageInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.checklistItem") {
		var out ChecklistItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChecklistItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudClipboardItem") {
		var out CloudClipboardItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudClipboardItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudClipboardRoot") {
		var out CloudClipboardRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudClipboardRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPC") {
		var out CloudPC
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPC: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcAuditEvent") {
		var out CloudPCAuditEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCAuditEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcDeviceImage") {
		var out CloudPCDeviceImage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCDeviceImage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcGalleryImage") {
		var out CloudPCGalleryImage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCGalleryImage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcOnPremisesConnection") {
		var out CloudPCOnPremisesConnection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCOnPremisesConnection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcProvisioningPolicy") {
		var out CloudPCProvisioningPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCProvisioningPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcProvisioningPolicyAssignment") {
		var out CloudPCProvisioningPolicyAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCProvisioningPolicyAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcUserSetting") {
		var out CloudPCUserSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCUserSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcUserSettingAssignment") {
		var out CloudPCUserSettingAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCUserSettingAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.columnDefinition") {
		var out ColumnDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ColumnDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.columnLink") {
		var out ColumnLink
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ColumnLink: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.commsOperation") {
		var out CommsOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommsOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.community") {
		var out Community
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Community: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.companySubscription") {
		var out CompanySubscription
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CompanySubscription: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.complianceManagementPartner") {
		var out ComplianceManagementPartner
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ComplianceManagementPartner: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessPolicy") {
		var out ConditionalAccessPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessRoot") {
		var out ConditionalAccessRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessTemplate") {
		var out ConditionalAccessTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.connectedOrganization") {
		var out ConnectedOrganization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectedOrganization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contactFolder") {
		var out ContactFolder
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContactFolder: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentSharingSession") {
		var out ContentSharingSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentSharingSession: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contentType") {
		var out ContentType
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentType: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversation") {
		var out Conversation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Conversation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversationMember") {
		var out ConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConversationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversationThread") {
		var out ConversationThread
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConversationThread: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.copilotAdmin") {
		var out CopilotAdmin
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CopilotAdmin: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.copilotAdminLimitedMode") {
		var out CopilotAdminLimitedMode
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CopilotAdminLimitedMode: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.copilotAdminSetting") {
		var out CopilotAdminSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CopilotAdminSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicyConfigurationDefault") {
		var out CrossTenantAccessPolicyConfigurationDefault
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicyConfigurationDefault: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customCalloutExtension") {
		var out CustomCalloutExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomCalloutExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customExtensionStageSetting") {
		var out CustomExtensionStageSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomExtensionStageSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customSecurityAttributeDefinition") {
		var out CustomSecurityAttributeDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSecurityAttributeDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dataPolicyOperation") {
		var out DataPolicyOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataPolicyOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminAccessAssignment") {
		var out DelegatedAdminAccessAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminAccessAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminCustomer") {
		var out DelegatedAdminCustomer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminCustomer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminRelationship") {
		var out DelegatedAdminRelationship
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminRelationship: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminRelationshipOperation") {
		var out DelegatedAdminRelationshipOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminRelationshipOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminRelationshipRequest") {
		var out DelegatedAdminRelationshipRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminRelationshipRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedAdminServiceManagementDetail") {
		var out DelegatedAdminServiceManagementDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedAdminServiceManagementDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.delegatedPermissionClassification") {
		var out DelegatedPermissionClassification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelegatedPermissionClassification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deletedChat") {
		var out DeletedChat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeletedChat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deletedItemContainer") {
		var out DeletedItemContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeletedItemContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deletedTeam") {
		var out DeletedTeam
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeletedTeam: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deltaParticipants") {
		var out DeltaParticipants
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeltaParticipants: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.detectedApp") {
		var out DetectedApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DetectedApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAppManagement") {
		var out DeviceAppManagement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAppManagement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCategory") {
		var out DeviceCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceActionItem") {
		var out DeviceComplianceActionItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceActionItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceDeviceOverview") {
		var out DeviceComplianceDeviceOverview
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceDeviceOverview: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceDeviceStatus") {
		var out DeviceComplianceDeviceStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceDeviceStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicy") {
		var out DeviceCompliancePolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyAssignment") {
		var out DeviceCompliancePolicyAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyDeviceStateSummary") {
		var out DeviceCompliancePolicyDeviceStateSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyDeviceStateSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicySettingStateSummary") {
		var out DeviceCompliancePolicySettingStateSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicySettingStateSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyState") {
		var out DeviceCompliancePolicyState
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyState: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScheduledActionForRule") {
		var out DeviceComplianceScheduledActionForRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScheduledActionForRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceSettingState") {
		var out DeviceComplianceSettingState
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceSettingState: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceUserOverview") {
		var out DeviceComplianceUserOverview
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceUserOverview: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceUserStatus") {
		var out DeviceComplianceUserStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceUserStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfiguration") {
		var out DeviceConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationAssignment") {
		var out DeviceConfigurationAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationDeviceOverview") {
		var out DeviceConfigurationDeviceOverview
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationDeviceOverview: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationDeviceStateSummary") {
		var out DeviceConfigurationDeviceStateSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationDeviceStateSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationDeviceStatus") {
		var out DeviceConfigurationDeviceStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationDeviceStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationState") {
		var out DeviceConfigurationState
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationState: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationUserOverview") {
		var out DeviceConfigurationUserOverview
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationUserOverview: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationUserStatus") {
		var out DeviceConfigurationUserStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationUserStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentConfiguration") {
		var out DeviceEnrollmentConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceInstallState") {
		var out DeviceInstallState
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceInstallState: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceLocalCredential") {
		var out DeviceLocalCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceLocalCredential: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceLocalCredentialInfo") {
		var out DeviceLocalCredentialInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceLocalCredentialInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceLogCollectionResponse") {
		var out DeviceLogCollectionResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceLogCollectionResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagement") {
		var out DeviceManagement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCachedReportConfiguration") {
		var out DeviceManagementCachedReportConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCachedReportConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExchangeConnector") {
		var out DeviceManagementExchangeConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExchangeConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementExportJob") {
		var out DeviceManagementExportJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementExportJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementPartner") {
		var out DeviceManagementPartner
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementPartner: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementReports") {
		var out DeviceManagementReports
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementReports: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTroubleshootingEvent") {
		var out DeviceManagementTroubleshootingEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTroubleshootingEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceRegistrationPolicy") {
		var out DeviceRegistrationPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceRegistrationPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directory") {
		var out Directory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Directory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryAudit") {
		var out DirectoryAudit
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryAudit: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryDefinition") {
		var out DirectoryDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryObject") {
		var out DirectoryObject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryObject: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domain") {
		var out Domain
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Domain: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsRecord") {
		var out DomainDnsRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.eBookInstallSummary") {
		var out EBookInstallSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EBookInstallSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.edge") {
		var out Edge
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Edge: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignment") {
		var out EducationAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentDefaults") {
		var out EducationAssignmentDefaults
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentDefaults: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentResource") {
		var out EducationAssignmentResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentSettings") {
		var out EducationAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationCategory") {
		var out EducationCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationClass") {
		var out EducationClass
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationClass: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationGradingCategory") {
		var out EducationGradingCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationGradingCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationModule") {
		var out EducationModule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationModule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationModuleResource") {
		var out EducationModuleResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationModuleResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationOrganization") {
		var out EducationOrganization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationOrganization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationOutcome") {
		var out EducationOutcome
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationOutcome: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationRubric") {
		var out EducationRubric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationRubric: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSubmission") {
		var out EducationSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSubmission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSubmissionResource") {
		var out EducationSubmissionResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSubmissionResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationUser") {
		var out EducationUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.employeeExperienceUser") {
		var out EmployeeExperienceUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmployeeExperienceUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endUserNotification") {
		var out EndUserNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndUserNotification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endUserNotificationDetail") {
		var out EndUserNotificationDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndUserNotificationDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentConfigurationAssignment") {
		var out EnrollmentConfigurationAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentConfigurationAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enterpriseCodeSigningCertificate") {
		var out EnterpriseCodeSigningCertificate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnterpriseCodeSigningCertificate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.entitlementManagement") {
		var out EntitlementManagement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EntitlementManagement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.entitlementManagementSettings") {
		var out EntitlementManagementSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EntitlementManagementSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extension") {
		var out Extension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Extension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.connectionOperation") {
		var out ExternalConnectorsConnectionOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsConnectionOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalActivity") {
		var out ExternalConnectorsExternalActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalConnection") {
		var out ExternalConnectorsExternalConnection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalConnection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalGroup") {
		var out ExternalConnectorsExternalGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalItem") {
		var out ExternalConnectorsExternalItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.identity") {
		var out ExternalConnectorsIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.schema") {
		var out ExternalConnectorsSchema
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsSchema: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalDomainName") {
		var out ExternalDomainName
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalDomainName: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.featureRolloutPolicy") {
		var out FeatureRolloutPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureRolloutPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.federatedIdentityCredential") {
		var out FederatedIdentityCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FederatedIdentityCredential: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fieldValueSet") {
		var out FieldValueSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FieldValueSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileStorage") {
		var out FileStorage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileStorage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileStorageContainer") {
		var out FileStorageContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileStorageContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.filterOperatorSchema") {
		var out FilterOperatorSchema
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FilterOperatorSchema: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.governanceInsight") {
		var out GovernanceInsight
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GovernanceInsight: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupLifecyclePolicy") {
		var out GroupLifecyclePolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupLifecyclePolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupSetting") {
		var out GroupSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.horizontalSection") {
		var out HorizontalSection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HorizontalSection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.horizontalSectionColumn") {
		var out HorizontalSectionColumn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HorizontalSectionColumn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityApiConnector") {
		var out IdentityApiConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityApiConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityContainer") {
		var out IdentityContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.insights") {
		var out IdentityGovernanceInsights
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceInsights: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.lifecycleManagementSettings") {
		var out IdentityGovernanceLifecycleManagementSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceLifecycleManagementSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.lifecycleWorkflowsContainer") {
		var out IdentityGovernanceLifecycleWorkflowsContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceLifecycleWorkflowsContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.run") {
		var out IdentityGovernanceRun
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceRun: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.task") {
		var out IdentityGovernanceTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskDefinition") {
		var out IdentityGovernanceTaskDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskProcessingResult") {
		var out IdentityGovernanceTaskProcessingResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskProcessingResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.taskReport") {
		var out IdentityGovernanceTaskReport
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTaskReport: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.userProcessingResult") {
		var out IdentityGovernanceUserProcessingResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceUserProcessingResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.workflowTemplate") {
		var out IdentityGovernanceWorkflowTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceWorkflowTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityProvider") {
		var out IdentityProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityProviderBase") {
		var out IdentityProviderBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityProviderBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityUserFlow") {
		var out IdentityUserFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityUserFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityUserFlowAttribute") {
		var out IdentityUserFlowAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityUserFlowAttribute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityUserFlowAttributeAssignment") {
		var out IdentityUserFlowAttributeAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityUserFlowAttributeAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedWindowsAutopilotDeviceIdentity") {
		var out ImportedWindowsAutopilotDeviceIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedWindowsAutopilotDeviceIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.importedWindowsAutopilotDeviceIdentityUpload") {
		var out ImportedWindowsAutopilotDeviceIdentityUpload
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedWindowsAutopilotDeviceIdentityUpload: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inferenceClassification") {
		var out InferenceClassification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InferenceClassification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inferenceClassificationOverride") {
		var out InferenceClassificationOverride
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InferenceClassificationOverride: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.insightsSettings") {
		var out InsightsSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InsightsSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.internetExplorerMode") {
		var out InternetExplorerMode
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InternetExplorerMode: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.invitation") {
		var out Invitation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Invitation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppProvisioningConfigurationAssignment") {
		var out IosLobAppProvisioningConfigurationAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppProvisioningConfigurationAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosUpdateDeviceStatus") {
		var out IosUpdateDeviceStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosUpdateDeviceStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemActivity") {
		var out ItemActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemActivityStat") {
		var out ItemActivityStat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemActivityStat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemAnalytics") {
		var out ItemAnalytics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemAnalytics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemRetentionLabel") {
		var out ItemRetentionLabel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemRetentionLabel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.landingPage") {
		var out LandingPage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LandingPage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.landingPageDetail") {
		var out LandingPageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LandingPageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningContent") {
		var out LearningContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningCourseActivity") {
		var out LearningCourseActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningCourseActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningProvider") {
		var out LearningProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.licenseDetails") {
		var out LicenseDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LicenseDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.linkedResource") {
		var out LinkedResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LinkedResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.localizedNotificationMessage") {
		var out LocalizedNotificationMessage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocalizedNotificationMessage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.loginPage") {
		var out LoginPage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LoginPage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.longRunningOperation") {
		var out LongRunningOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LongRunningOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.m365AppsInstallationOptions") {
		var out M365AppsInstallationOptions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into M365AppsInstallationOptions: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailFolder") {
		var out MailFolder
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailFolder: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.malwareStateForWindowsDevice") {
		var out MalwareStateForWindowsDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MalwareStateForWindowsDevice: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppOperation") {
		var out ManagedAppOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppPolicy") {
		var out ManagedAppPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppPolicyDeploymentSummary") {
		var out ManagedAppPolicyDeploymentSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppPolicyDeploymentSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppRegistration") {
		var out ManagedAppRegistration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppRegistration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppStatus") {
		var out ManagedAppStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDevice") {
		var out ManagedDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDevice: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfiguration") {
		var out ManagedDeviceMobileAppConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationAssignment") {
		var out ManagedDeviceMobileAppConfigurationAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatus") {
		var out ManagedDeviceMobileAppConfigurationDeviceStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationDeviceStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceSummary") {
		var out ManagedDeviceMobileAppConfigurationDeviceSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationDeviceSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationUserStatus") {
		var out ManagedDeviceMobileAppConfigurationUserStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationUserStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationUserSummary") {
		var out ManagedDeviceMobileAppConfigurationUserSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationUserSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceOverview") {
		var out ManagedDeviceOverview
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceOverview: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedEBook") {
		var out ManagedEBook
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedEBook: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedEBookAssignment") {
		var out ManagedEBookAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedEBookAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedMobileApp") {
		var out ManagedMobileApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedMobileApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingAttendanceReport") {
		var out MeetingAttendanceReport
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingAttendanceReport: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messageRule") {
		var out MessageRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileApp") {
		var out MobileApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppAssignment") {
		var out MobileAppAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppCategory") {
		var out MobileAppCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppContent") {
		var out MobileAppContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppContentFile") {
		var out MobileAppContentFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppContentFile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingEvent") {
		var out MobileAppTroubleshootingEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileContainedApp") {
		var out MobileContainedApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileContainedApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileThreatDefenseConnector") {
		var out MobileThreatDefenseConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileThreatDefenseConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiTenantOrganization") {
		var out MultiTenantOrganization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiTenantOrganization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiTenantOrganizationIdentitySyncPolicyTemplate") {
		var out MultiTenantOrganizationIdentitySyncPolicyTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiTenantOrganizationIdentitySyncPolicyTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiTenantOrganizationJoinRequestRecord") {
		var out MultiTenantOrganizationJoinRequestRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiTenantOrganizationJoinRequestRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiTenantOrganizationPartnerConfigurationTemplate") {
		var out MultiTenantOrganizationPartnerConfigurationTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiTenantOrganizationPartnerConfigurationTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiValueLegacyExtendedProperty") {
		var out MultiValueLegacyExtendedProperty
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiValueLegacyExtendedProperty: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.namedLocation") {
		var out NamedLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NamedLocation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.notificationMessageTemplate") {
		var out NotificationMessageTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NotificationMessageTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oAuth2PermissionGrant") {
		var out OAuth2PermissionGrant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OAuth2PermissionGrant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.officeGraphInsights") {
		var out OfficeGraphInsights
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeGraphInsights: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesConditionalAccessSettings") {
		var out OnPremisesConditionalAccessSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesConditionalAccessSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onPremisesDirectorySynchronization") {
		var out OnPremisesDirectorySynchronization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremisesDirectorySynchronization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenote") {
		var out Onenote
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Onenote: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteEntityBaseModel") {
		var out OnenoteEntityBaseModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteEntityBaseModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onlineMeetingBase") {
		var out OnlineMeetingBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnlineMeetingBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.operation") {
		var out Operation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Operation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.organizationalBrandingProperties") {
		var out OrganizationalBrandingProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizationalBrandingProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outlookCategory") {
		var out OutlookCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outlookItem") {
		var out OutlookItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outlookUser") {
		var out OutlookUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutlookUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.participant") {
		var out Participant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Participant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.participantJoiningNotification") {
		var out ParticipantJoiningNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ParticipantJoiningNotification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.participantLeftNotification") {
		var out ParticipantLeftNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ParticipantLeftNotification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners") {
		var out Partners
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Partners: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.azureUsage") {
		var out PartnersBillingAzureUsage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingAzureUsage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.billedReconciliation") {
		var out PartnersBillingBilledReconciliation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingBilledReconciliation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.billedUsage") {
		var out PartnersBillingBilledUsage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingBilledUsage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.billing") {
		var out PartnersBillingBilling
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingBilling: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.billingReconciliation") {
		var out PartnersBillingBillingReconciliation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingBillingReconciliation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.manifest") {
		var out PartnersBillingManifest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingManifest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.operation") {
		var out PartnersBillingOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.unbilledUsage") {
		var out PartnersBillingUnbilledUsage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingUnbilledUsage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.payload") {
		var out Payload
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Payload: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.peopleAdminSettings") {
		var out PeopleAdminSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PeopleAdminSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permission") {
		var out Permission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Permission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionGrantConditionSet") {
		var out PermissionGrantConditionSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionGrantConditionSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.person") {
		var out Person
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Person: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pinnedChatMessageInfo") {
		var out PinnedChatMessageInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PinnedChatMessageInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.place") {
		var out Place
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Place: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.planner") {
		var out Planner
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Planner: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerAssignedToTaskBoardTaskFormat") {
		var out PlannerAssignedToTaskBoardTaskFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerAssignedToTaskBoardTaskFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBucket") {
		var out PlannerBucket
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucket: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBucketTaskBoardTaskFormat") {
		var out PlannerBucketTaskBoardTaskFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucketTaskBoardTaskFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerGroup") {
		var out PlannerGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlan") {
		var out PlannerPlan
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlan: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanDetails") {
		var out PlannerPlanDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerProgressTaskBoardTaskFormat") {
		var out PlannerProgressTaskBoardTaskFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerProgressTaskBoardTaskFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTask") {
		var out PlannerTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTaskDetails") {
		var out PlannerTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerUser") {
		var out PlannerUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyRoot") {
		var out PolicyRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyTemplate") {
		var out PolicyTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.presence") {
		var out Presence
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Presence: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printConnector") {
		var out PrintConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printDocument") {
		var out PrintDocument
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintDocument: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printJob") {
		var out PrintJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printOperation") {
		var out PrintOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printService") {
		var out PrintService
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintService: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printServiceEndpoint") {
		var out PrintServiceEndpoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintServiceEndpoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printTask") {
		var out PrintTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printTaskDefinition") {
		var out PrintTaskDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintTaskDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printTaskTrigger") {
		var out PrintTaskTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintTaskTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printUsage") {
		var out PrintUsage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintUsage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printerBase") {
		var out PrinterBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrinterBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroup") {
		var out PrivilegedAccessGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessRoot") {
		var out PrivilegedAccessRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessSchedule") {
		var out PrivilegedAccessSchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessSchedule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessScheduleInstance") {
		var out PrivilegedAccessScheduleInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessScheduleInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profileCardProperty") {
		var out ProfileCardProperty
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfileCardProperty: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.profilePhoto") {
		var out ProfilePhoto
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProfilePhoto: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pronounsSettings") {
		var out PronounsSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PronounsSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectionPolicyBase") {
		var out ProtectionPolicyBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectionPolicyBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectionRuleBase") {
		var out ProtectionRuleBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectionRuleBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectionUnitBase") {
		var out ProtectionUnitBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectionUnitBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectionUnitsBulkJobBase") {
		var out ProtectionUnitsBulkJobBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectionUnitsBulkJobBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisioningObjectSummary") {
		var out ProvisioningObjectSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisioningObjectSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.publicKeyInfrastructureRoot") {
		var out PublicKeyInfrastructureRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PublicKeyInfrastructureRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rbacApplication") {
		var out RbacApplication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RbacApplication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.relyingPartyDetailedSummary") {
		var out RelyingPartyDetailedSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RelyingPartyDetailedSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.remoteAssistancePartner") {
		var out RemoteAssistancePartner
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RemoteAssistancePartner: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.remoteDesktopSecurityConfiguration") {
		var out RemoteDesktopSecurityConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RemoteDesktopSecurityConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.request") {
		var out Request
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Request: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceOperation") {
		var out ResourceOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restoreArtifactBase") {
		var out RestoreArtifactBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestoreArtifactBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restoreArtifactsBulkRequestBase") {
		var out RestoreArtifactsBulkRequestBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestoreArtifactsBulkRequestBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restorePoint") {
		var out RestorePoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestorePoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.restoreSessionBase") {
		var out RestoreSessionBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RestoreSessionBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskDetection") {
		var out RiskDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskDetection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyServicePrincipal") {
		var out RiskyServicePrincipal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyServicePrincipal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyUser") {
		var out RiskyUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleAssignment") {
		var out RoleAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleDefinition") {
		var out RoleDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.schedule") {
		var out Schedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Schedule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.schemaExtension") {
		var out SchemaExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SchemaExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.scopedRoleMembership") {
		var out ScopedRoleMembership
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScopedRoleMembership: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.searchEntity") {
		var out SearchEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.searchAnswer") {
		var out SearchSearchAnswer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchSearchAnswer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secureScore") {
		var out SecureScore
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecureScore: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secureScoreControlProfile") {
		var out SecureScoreControlProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecureScoreControlProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security") {
		var out Security
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Security: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.alert") {
		var out SecurityAlert
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlert: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.article") {
		var out SecurityArticle
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityArticle: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.artifact") {
		var out SecurityArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityArtifact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.case") {
		var out SecurityCase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.caseOperation") {
		var out SecurityCaseOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCaseOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.casesRoot") {
		var out SecurityCasesRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCasesRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataSet") {
		var out SecurityDataSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataSource") {
		var out SecurityDataSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataSourceContainer") {
		var out SecurityDataSourceContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataSourceContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dispositionReviewStage") {
		var out SecurityDispositionReviewStage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDispositionReviewStage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryCaseSettings") {
		var out SecurityEdiscoveryCaseSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryCaseSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanDescriptor") {
		var out SecurityFilePlanDescriptor
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanDescriptor: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanDescriptorTemplate") {
		var out SecurityFilePlanDescriptorTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanDescriptorTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.healthIssue") {
		var out SecurityHealthIssue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHealthIssue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostPair") {
		var out SecurityHostPair
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostPair: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostPort") {
		var out SecurityHostPort
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostPort: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostReputation") {
		var out SecurityHostReputation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostReputation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.identityContainer") {
		var out SecurityIdentityContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIdentityContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.incident") {
		var out SecurityIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.indicator") {
		var out SecurityIndicator
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIndicator: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.intelligenceProfile") {
		var out SecurityIntelligenceProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIntelligenceProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.labelsRoot") {
		var out SecurityLabelsRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityLabelsRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.networkAdapter") {
		var out SecurityNetworkAdapter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityNetworkAdapter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityReportsRoot") {
		var out SecurityReportsRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityReportsRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionEvent") {
		var out SecurityRetentionEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionEventType") {
		var out SecurityRetentionEventType
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionEventType: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionLabel") {
		var out SecurityRetentionLabel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionLabel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.search") {
		var out SecuritySearch
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySearch: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sensor") {
		var out SecuritySensor
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySensor: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.subdomain") {
		var out SecuritySubdomain
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySubdomain: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.tag") {
		var out SecurityTag
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTag: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.threatIntelligence") {
		var out SecurityThreatIntelligence
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityThreatIntelligence: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.triggerTypesRoot") {
		var out SecurityTriggerTypesRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTriggerTypesRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.triggersRoot") {
		var out SecurityTriggersRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTriggersRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vulnerability") {
		var out SecurityVulnerability
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVulnerability: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vulnerabilityComponent") {
		var out SecurityVulnerabilityComponent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVulnerabilityComponent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.whoisBaseRecord") {
		var out SecurityWhoisBaseRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWhoisBaseRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceAnnouncement") {
		var out ServiceAnnouncement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceAnnouncement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceAnnouncementAttachment") {
		var out ServiceAnnouncementAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceAnnouncementAttachment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceAnnouncementBase") {
		var out ServiceAnnouncementBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceAnnouncementBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceApp") {
		var out ServiceApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceHealth") {
		var out ServiceHealth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceHealth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalRiskDetection") {
		var out ServicePrincipalRiskDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalRiskDetection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.settingStateDeviceSummary") {
		var out SettingStateDeviceSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SettingStateDeviceSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedInsight") {
		var out SharedInsight
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedInsight: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharepoint") {
		var out Sharepoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Sharepoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharepointSettings") {
		var out SharepointSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharepointSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.signIn") {
		var out SignIn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SignIn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulation") {
		var out Simulation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Simulation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationAutomation") {
		var out SimulationAutomation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationAutomation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationAutomationRun") {
		var out SimulationAutomationRun
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationAutomationRun: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleValueLegacyExtendedProperty") {
		var out SingleValueLegacyExtendedProperty
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleValueLegacyExtendedProperty: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.softwareUpdateStatusSummary") {
		var out SoftwareUpdateStatusSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SoftwareUpdateStatusSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.storageQuotaBreakdown") {
		var out StorageQuotaBreakdown
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StorageQuotaBreakdown: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.storageSettings") {
		var out StorageSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StorageSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequest") {
		var out SubjectRightsRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subscribedSku") {
		var out SubscribedSku
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubscribedSku: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subscription") {
		var out Subscription
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Subscription: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronization") {
		var out Synchronization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Synchronization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationJob") {
		var out SynchronizationJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationSchema") {
		var out SynchronizationSchema
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationSchema: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.synchronizationTemplate") {
		var out SynchronizationTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynchronizationTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetDeviceGroup") {
		var out TargetDeviceGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetDeviceGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppPolicyAssignment") {
		var out TargetedManagedAppPolicyAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppPolicyAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.team") {
		var out Team
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Team: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamInfo") {
		var out TeamInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsApp") {
		var out TeamsApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppDefinition") {
		var out TeamsAppDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppInstallation") {
		var out TeamsAppInstallation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppInstallation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppSettings") {
		var out TeamsAppSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAsyncOperation") {
		var out TeamsAsyncOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAsyncOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsLicensingDetails") {
		var out TeamsLicensingDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsLicensingDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsTab") {
		var out TeamsTab
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsTab: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsTemplate") {
		var out TeamsTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamwork") {
		var out Teamwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Teamwork: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkBot") {
		var out TeamworkBot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkBot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkHostedContent") {
		var out TeamworkHostedContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkHostedContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkTag") {
		var out TeamworkTag
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkTag: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkTagMember") {
		var out TeamworkTagMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkTagMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.telecomExpenseManagementPartner") {
		var out TelecomExpenseManagementPartner
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TelecomExpenseManagementPartner: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.group") {
		var out TermStoreGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.relation") {
		var out TermStoreRelation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreRelation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.set") {
		var out TermStoreSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.store") {
		var out TermStoreStore
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreStore: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termStore.term") {
		var out TermStoreTerm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermStoreTerm: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditions") {
		var out TermsAndConditions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditions: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditionsAcceptanceStatus") {
		var out TermsAndConditionsAcceptanceStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditionsAcceptanceStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsAndConditionsAssignment") {
		var out TermsAndConditionsAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsAndConditionsAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.termsOfUseContainer") {
		var out TermsOfUseContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TermsOfUseContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.threatAssessmentRequest") {
		var out ThreatAssessmentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThreatAssessmentRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.threatAssessmentResult") {
		var out ThreatAssessmentResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThreatAssessmentResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.thumbnailSet") {
		var out ThumbnailSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThumbnailSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.todo") {
		var out Todo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Todo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.todoTask") {
		var out TodoTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TodoTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.todoTaskList") {
		var out TodoTaskList
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TodoTaskList: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.training") {
		var out Training
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Training: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trainingLanguageDetail") {
		var out TrainingLanguageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrainingLanguageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trending") {
		var out Trending
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Trending: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRbacResourceAction") {
		var out UnifiedRbacResourceAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRbacResourceAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRbacResourceNamespace") {
		var out UnifiedRbacResourceNamespace
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRbacResourceNamespace: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignment") {
		var out UnifiedRoleAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleDefinition") {
		var out UnifiedRoleDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicy") {
		var out UnifiedRoleManagementPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyAssignment") {
		var out UnifiedRoleManagementPolicyAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyRule") {
		var out UnifiedRoleManagementPolicyRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleScheduleBase") {
		var out UnifiedRoleScheduleBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleScheduleBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleScheduleInstanceBase") {
		var out UnifiedRoleScheduleInstanceBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleScheduleInstanceBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedStorageQuota") {
		var out UnifiedStorageQuota
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedStorageQuota: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.usedInsight") {
		var out UsedInsight
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UsedInsight: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userActivity") {
		var out UserActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByOSVersion") {
		var out UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthApplicationPerformance") {
		var out UserExperienceAnalyticsAppHealthApplicationPerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthApplicationPerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthDeviceModelPerformance") {
		var out UserExperienceAnalyticsAppHealthDeviceModelPerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthDeviceModelPerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformance") {
		var out UserExperienceAnalyticsAppHealthDevicePerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthDevicePerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformanceDetails") {
		var out UserExperienceAnalyticsAppHealthDevicePerformanceDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthDevicePerformanceDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsAppHealthOSVersionPerformance") {
		var out UserExperienceAnalyticsAppHealthOSVersionPerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsAppHealthOSVersionPerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsBaseline") {
		var out UserExperienceAnalyticsBaseline
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsBaseline: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsCategory") {
		var out UserExperienceAnalyticsCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDevicePerformance") {
		var out UserExperienceAnalyticsDevicePerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDevicePerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceScores") {
		var out UserExperienceAnalyticsDeviceScores
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceScores: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceStartupHistory") {
		var out UserExperienceAnalyticsDeviceStartupHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceStartupHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcess") {
		var out UserExperienceAnalyticsDeviceStartupProcess
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceStartupProcess: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcessPerformance") {
		var out UserExperienceAnalyticsDeviceStartupProcessPerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsDeviceStartupProcessPerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsMetric") {
		var out UserExperienceAnalyticsMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsMetric: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsMetricHistory") {
		var out UserExperienceAnalyticsMetricHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsMetricHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsModelScores") {
		var out UserExperienceAnalyticsModelScores
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsModelScores: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsOverview") {
		var out UserExperienceAnalyticsOverview
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsOverview: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsScoreHistory") {
		var out UserExperienceAnalyticsScoreHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsScoreHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereDevice") {
		var out UserExperienceAnalyticsWorkFromAnywhereDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereDevice: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric") {
		var out UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric") {
		var out UserExperienceAnalyticsWorkFromAnywhereMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereMetric: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereModelPerformance") {
		var out UserExperienceAnalyticsWorkFromAnywhereModelPerformance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserExperienceAnalyticsWorkFromAnywhereModelPerformance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userFlowLanguageConfiguration") {
		var out UserFlowLanguageConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserFlowLanguageConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userFlowLanguagePage") {
		var out UserFlowLanguagePage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserFlowLanguagePage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userInsightsSettings") {
		var out UserInsightsSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserInsightsSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userInstallStateSummary") {
		var out UserInstallStateSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserInstallStateSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userRegistrationDetails") {
		var out UserRegistrationDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserRegistrationDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSettings") {
		var out UserSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSolutionRoot") {
		var out UserSolutionRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSolutionRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userStorage") {
		var out UserStorage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserStorage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userTeamwork") {
		var out UserTeamwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserTeamwork: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verticalSection") {
		var out VerticalSection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerticalSection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEndpoint") {
		var out VirtualEndpoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEndpoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEvent") {
		var out VirtualEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventPresenter") {
		var out VirtualEventPresenter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventPresenter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistration") {
		var out VirtualEventRegistration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationConfiguration") {
		var out VirtualEventRegistrationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationQuestionBase") {
		var out VirtualEventRegistrationQuestionBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationQuestionBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventsRoot") {
		var out VirtualEventsRoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventsRoot: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.vppToken") {
		var out VppToken
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VppToken: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webPart") {
		var out WebPart
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebPart: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeploymentProfile") {
		var out WindowsAutopilotDeploymentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeploymentProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeploymentProfileAssignment") {
		var out WindowsAutopilotDeploymentProfileAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeploymentProfileAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeviceIdentity") {
		var out WindowsAutopilotDeviceIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeviceIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDeviceMalwareState") {
		var out WindowsDeviceMalwareState
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDeviceMalwareState: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionAppLearningSummary") {
		var out WindowsInformationProtectionAppLearningSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionAppLearningSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionAppLockerFile") {
		var out WindowsInformationProtectionAppLockerFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionAppLockerFile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionNetworkLearningSummary") {
		var out WindowsInformationProtectionNetworkLearningSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionNetworkLearningSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMalwareInformation") {
		var out WindowsMalwareInformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMalwareInformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsProtectionState") {
		var out WindowsProtectionState
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsProtectionState: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsSetting") {
		var out WindowsSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsSettingInstance") {
		var out WindowsSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbook") {
		var out Workbook
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Workbook: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookApplication") {
		var out WorkbookApplication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookApplication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChart") {
		var out WorkbookChart
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChart: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartAreaFormat") {
		var out WorkbookChartAreaFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartAreaFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartAxes") {
		var out WorkbookChartAxes
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartAxes: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartAxis") {
		var out WorkbookChartAxis
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartAxis: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartAxisFormat") {
		var out WorkbookChartAxisFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartAxisFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartAxisTitle") {
		var out WorkbookChartAxisTitle
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartAxisTitle: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartAxisTitleFormat") {
		var out WorkbookChartAxisTitleFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartAxisTitleFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartDataLabelFormat") {
		var out WorkbookChartDataLabelFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartDataLabelFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartDataLabels") {
		var out WorkbookChartDataLabels
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartDataLabels: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartFill") {
		var out WorkbookChartFill
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartFill: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartFont") {
		var out WorkbookChartFont
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartFont: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartGridlines") {
		var out WorkbookChartGridlines
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartGridlines: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartGridlinesFormat") {
		var out WorkbookChartGridlinesFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartGridlinesFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartLegend") {
		var out WorkbookChartLegend
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartLegend: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartLegendFormat") {
		var out WorkbookChartLegendFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartLegendFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartLineFormat") {
		var out WorkbookChartLineFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartLineFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartPoint") {
		var out WorkbookChartPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartPointFormat") {
		var out WorkbookChartPointFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartPointFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartSeries") {
		var out WorkbookChartSeries
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartSeries: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartSeriesFormat") {
		var out WorkbookChartSeriesFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartSeriesFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartTitle") {
		var out WorkbookChartTitle
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartTitle: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookChartTitleFormat") {
		var out WorkbookChartTitleFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookChartTitleFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookComment") {
		var out WorkbookComment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookComment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookCommentReply") {
		var out WorkbookCommentReply
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookCommentReply: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookFilter") {
		var out WorkbookFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookFilter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookFormatProtection") {
		var out WorkbookFormatProtection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookFormatProtection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookFunctionResult") {
		var out WorkbookFunctionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookFunctionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookFunctions") {
		var out WorkbookFunctions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookFunctions: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookNamedItem") {
		var out WorkbookNamedItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookNamedItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookOperation") {
		var out WorkbookOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookPivotTable") {
		var out WorkbookPivotTable
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookPivotTable: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRange") {
		var out WorkbookRange
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRange: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeBorder") {
		var out WorkbookRangeBorder
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeBorder: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeFill") {
		var out WorkbookRangeFill
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeFill: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeFont") {
		var out WorkbookRangeFont
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeFont: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeFormat") {
		var out WorkbookRangeFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeSort") {
		var out WorkbookRangeSort
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeSort: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookRangeView") {
		var out WorkbookRangeView
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookRangeView: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTable") {
		var out WorkbookTable
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTable: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTableColumn") {
		var out WorkbookTableColumn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTableColumn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTableRow") {
		var out WorkbookTableRow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTableRow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookTableSort") {
		var out WorkbookTableSort
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookTableSort: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookWorksheet") {
		var out WorkbookWorksheet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookWorksheet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workbookWorksheetProtection") {
		var out WorkbookWorksheetProtection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkbookWorksheetProtection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workingTimeSchedule") {
		var out WorkingTimeSchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkingTimeSchedule: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityImpl: %+v", err)
	}

	return RawEntityImpl{
		entity: parent,
		Type:   value,
		Values: temp,
	}, nil

}
