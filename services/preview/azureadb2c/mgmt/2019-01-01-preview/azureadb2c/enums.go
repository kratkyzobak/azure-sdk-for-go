package azureadb2c

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// B2CResourceSKUName enumerates the values for b2c resource sku name.
type B2CResourceSKUName string

const (
	// PremiumP1 Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	PremiumP1 B2CResourceSKUName = "PremiumP1"
	// PremiumP2 Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	PremiumP2 B2CResourceSKUName = "PremiumP2"
	// Standard Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users
	// (MAU) billing model.
	Standard B2CResourceSKUName = "Standard"
)

// PossibleB2CResourceSKUNameValues returns an array of possible values for the B2CResourceSKUName const type.
func PossibleB2CResourceSKUNameValues() []B2CResourceSKUName {
	return []B2CResourceSKUName{PremiumP1, PremiumP2, Standard}
}

// B2CResourceSKUTier enumerates the values for b2c resource sku tier.
type B2CResourceSKUTier string

const (
	// A0 The SKU tier used for all Azure AD B2C tenants.
	A0 B2CResourceSKUTier = "A0"
)

// PossibleB2CResourceSKUTierValues returns an array of possible values for the B2CResourceSKUTier const type.
func PossibleB2CResourceSKUTierValues() []B2CResourceSKUTier {
	return []B2CResourceSKUTier{A0}
}

// BillingType enumerates the values for billing type.
type BillingType string

const (
	// Auths Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	Auths BillingType = "Auths"
	// MAU Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users (MAU)
	// billing model.
	MAU BillingType = "MAU"
)

// PossibleBillingTypeValues returns an array of possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{Auths, MAU}
}

// NameAvailabilityReasonType enumerates the values for name availability reason type.
type NameAvailabilityReasonType string

const (
	// AlreadyExists The name is already in use and is therefore unavailable.
	AlreadyExists NameAvailabilityReasonType = "AlreadyExists"
	// Invalid The name provided does not match the resource provider’s naming requirements (incorrect length,
	// unsupported characters, etc.).
	Invalid NameAvailabilityReasonType = "Invalid"
)

// PossibleNameAvailabilityReasonTypeValues returns an array of possible values for the NameAvailabilityReasonType const type.
func PossibleNameAvailabilityReasonTypeValues() []NameAvailabilityReasonType {
	return []NameAvailabilityReasonType{AlreadyExists, Invalid}
}

// StatusType enumerates the values for status type.
type StatusType string

const (
	// Failed The operation failed.
	Failed StatusType = "Failed"
	// Pending The operation is pending.
	Pending StatusType = "Pending"
	// Succeeded The operation succeeded.
	Succeeded StatusType = "Succeeded"
)

// PossibleStatusTypeValues returns an array of possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{Failed, Pending, Succeeded}
}

// TypeValue enumerates the values for type value.
type TypeValue string

const (
	// MicrosoftAzureActiveDirectoryb2cDirectories The resource type for Azure AD B2C tenant resource.
	MicrosoftAzureActiveDirectoryb2cDirectories TypeValue = "Microsoft.AzureActiveDirectory/b2cDirectories"
)

// PossibleTypeValueValues returns an array of possible values for the TypeValue const type.
func PossibleTypeValueValues() []TypeValue {
	return []TypeValue{MicrosoftAzureActiveDirectoryb2cDirectories}
}