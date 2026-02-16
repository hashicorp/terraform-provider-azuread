package transitivemember

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ApplicationOperationPredicate struct {
}

func (p ApplicationOperationPredicate) Matches(input beta.Application) bool {

	return true
}

type DeviceOperationPredicate struct {
}

func (p DeviceOperationPredicate) Matches(input beta.Device) bool {

	return true
}

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input beta.DirectoryObject) bool {

	return true
}

type GroupOperationPredicate struct {
}

func (p GroupOperationPredicate) Matches(input beta.Group) bool {

	return true
}

type OrgContactOperationPredicate struct {
}

func (p OrgContactOperationPredicate) Matches(input beta.OrgContact) bool {

	return true
}

type ServicePrincipalOperationPredicate struct {
}

func (p ServicePrincipalOperationPredicate) Matches(input beta.ServicePrincipal) bool {

	return true
}

type UserOperationPredicate struct {
}

func (p UserOperationPredicate) Matches(input beta.User) bool {

	return true
}
