//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AsyncOperationDetail
type AsyncOperationDetail struct {
	// The operation end time.
	EndTime *string `json:"endTime,omitempty"`

	// The error.
	Error *ErrorDetail `json:"error,omitempty"`

	// The operation id.
	ID *string `json:"id,omitempty"`

	// The operation name.
	Name *string `json:"name,omitempty"`

	// The operation start time.
	StartTime *string `json:"startTime,omitempty"`

	// The operation status.
	Status *string `json:"status,omitempty"`
}

// ConnectionState information.
type ConnectionState struct {
	// Actions required (if any).
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// Description of the connection state.
	Description *string `json:"description,omitempty"`

	// Status of the connection.
	Status *PersistedConnectionStatus `json:"status,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData
// error response format.).
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// The error object.
	InnerError *ErrorDetail `json:"error,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// GroupConnectivityInformation
type GroupConnectivityInformation struct {
	// Specifies the customer visible FQDNs of the group connectivity information.
	CustomerVisibleFqdns []*string `json:"customerVisibleFqdns,omitempty"`

	// Specifies the group id of the group connectivity information.
	GroupID *string `json:"groupId,omitempty"`

	// Specifies the internal FQDN of the group connectivity information.
	InternalFqdn *string `json:"internalFqdn,omitempty"`

	// Specifies the member name of the group connectivity information.
	MemberName *string `json:"memberName,omitempty"`

	// Specifies the ARM region of the group connectivity information.
	PrivateLinkServiceArmRegion *string `json:"privateLinkServiceArmRegion,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type GroupConnectivityInformation.
func (g GroupConnectivityInformation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "customerVisibleFqdns", g.CustomerVisibleFqdns)
	populate(objectMap, "groupId", g.GroupID)
	populate(objectMap, "internalFqdn", g.InternalFqdn)
	populate(objectMap, "memberName", g.MemberName)
	populate(objectMap, "privateLinkServiceArmRegion", g.PrivateLinkServiceArmRegion)
	return json.Marshal(objectMap)
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write", "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine", "Restart Virtual
	// Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// PowerBIResourcesCreateOptions contains the optional parameters for the PowerBIResources.Create method.
type PowerBIResourcesCreateOptions struct {
	// The client tenant id in header. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
	ClientTenantID *string
}

// PowerBIResourcesDeleteOptions contains the optional parameters for the PowerBIResources.Delete method.
type PowerBIResourcesDeleteOptions struct {
	// placeholder for future optional parameters
}

// PowerBIResourcesListByResourceNameOptions contains the optional parameters for the PowerBIResources.ListByResourceName method.
type PowerBIResourcesListByResourceNameOptions struct {
	// placeholder for future optional parameters
}

// PowerBIResourcesUpdateOptions contains the optional parameters for the PowerBIResources.Update method.
type PowerBIResourcesUpdateOptions struct {
	// The client tenant id in header. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
	ClientTenantID *string
}

// PrivateEndpoint
type PrivateEndpoint struct {
	// Specifies the id of private endpoint.
	ID *string `json:"id,omitempty"`
}

// PrivateEndpointConnection
type PrivateEndpointConnection struct {
	// Specifies the properties of the private endpoint connection.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; Specifies the id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionListResult - List of private endpoint connections.
type PrivateEndpointConnectionListResult struct {
	// Specifies the name of the private endpoint connection.
	Value []*PrivateEndpointConnection `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// PrivateEndpointConnectionProperties
type PrivateEndpointConnectionProperties struct {
	// Specifies the private endpoint.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// Specifies the connection state.
	PrivateLinkServiceConnectionState *ConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *ResourceProvisioningState `json:"provisioningState,omitempty"`
}

// PrivateEndpointConnectionsBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnections.BeginDelete method.
type PrivateEndpointConnectionsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsCreateOptions contains the optional parameters for the PrivateEndpointConnections.Create method.
type PrivateEndpointConnectionsCreateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsGetOptions contains the optional parameters for the PrivateEndpointConnections.Get method.
type PrivateEndpointConnectionsGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsListByResourceOptions contains the optional parameters for the PrivateEndpointConnections.ListByResource method.
type PrivateEndpointConnectionsListByResourceOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkConnectionDetail
type PrivateLinkConnectionDetail struct {
	// Specifies the group id of the connection detail.
	GroupID *string `json:"groupId,omitempty"`

	// Specifies the type of the connection detail.
	ID *string `json:"id,omitempty"`

	// Specifies the link id of the connection detail.
	LinkIdentifier *string `json:"linkIdentifier,omitempty"`

	// Specifies the member name of the connection detail.
	MemberName *string `json:"memberName,omitempty"`

	// Specifies the private ip address of the connection detail.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Fully qualified identifier of the resource.
	ID *string `json:"id,omitempty"`

	// Name of the resource.
	Name *string `json:"name,omitempty"`

	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`

	// Type of the resource.
	Type *string `json:"type,omitempty"`
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty"`

	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// PrivateLinkResourcesGetOptions contains the optional parameters for the PrivateLinkResources.Get method.
type PrivateLinkResourcesGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesListByResourceOptions contains the optional parameters for the PrivateLinkResources.ListByResource method.
type PrivateLinkResourcesListByResourceOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesListResult - Specifies list of the private link resource.
type PrivateLinkResourcesListResult struct {
	// A collection of private endpoint connection resources.
	Value []*PrivateLinkResource `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourcesListResult.
func (p PrivateLinkResourcesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// PrivateLinkServiceConnection
type PrivateLinkServiceConnection struct {
	// Specifies the group ids of the private link service connection.
	GroupIDs []*string `json:"groupIds,omitempty"`

	// Specifies the name of the private link service connection.
	Name *string `json:"name,omitempty"`

	// Specifies the request message of the private link service connection.
	RequestMessage *string `json:"requestMessage,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkServiceConnection.
func (p PrivateLinkServiceConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupIds", p.GroupIDs)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "requestMessage", p.RequestMessage)
	return json.Marshal(objectMap)
}

// PrivateLinkServiceProxy
type PrivateLinkServiceProxy struct {
	// Specifies the group connectivity information of the private link service proxy.
	GroupConnectivityInformation []*GroupConnectivityInformation `json:"groupConnectivityInformation,omitempty"`

	// Specifies the id of the private link service proxy.
	ID *string `json:"id,omitempty"`

	// Specifies the private endpoint connection of the private link service proxy.
	RemotePrivateEndpointConnection *RemotePrivateEndpointConnection `json:"remotePrivateEndpointConnection,omitempty"`

	// Specifies the connection state of the private link service proxy.
	RemotePrivateLinkServiceConnectionState *ConnectionState `json:"remotePrivateLinkServiceConnectionState,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkServiceProxy.
func (p PrivateLinkServiceProxy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupConnectivityInformation", p.GroupConnectivityInformation)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "remotePrivateEndpointConnection", p.RemotePrivateEndpointConnection)
	populate(objectMap, "remotePrivateLinkServiceConnectionState", p.RemotePrivateLinkServiceConnectionState)
	return json.Marshal(objectMap)
}

// PrivateLinkServiceResourceOperationResultsBeginGetOptions contains the optional parameters for the PrivateLinkServiceResourceOperationResults.BeginGet
// method.
type PrivateLinkServiceResourceOperationResultsBeginGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServicesForPowerBIListBySubscriptionIDOptions contains the optional parameters for the PrivateLinkServicesForPowerBI.ListBySubscriptionID
// method.
type PrivateLinkServicesForPowerBIListBySubscriptionIDOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServicesListByResourceGroupOptions contains the optional parameters for the PrivateLinkServices.ListByResourceGroup method.
type PrivateLinkServicesListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// RemotePrivateEndpointConnection
type RemotePrivateEndpointConnection struct {
	// Specifies the id of private endpoint connection.
	ID *string `json:"id,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TenantProperties
type TenantProperties struct {
	// Specifies the private endpoint connections of the resource.
	PrivateEndpointConnections []*PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`

	// Specifies the tenant id of the resource.
	TenantID *string `json:"tenantId,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TenantProperties.
func (t TenantProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "privateEndpointConnections", t.PrivateEndpointConnections)
	populate(objectMap, "tenantId", t.TenantID)
	return json.Marshal(objectMap)
}

// TenantResource
type TenantResource struct {
	// Specifies the location of the resource.
	Location *string `json:"location,omitempty"`

	// Specifies the properties of the resource.
	Properties *TenantProperties `json:"properties,omitempty"`

	// Specifies the tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Specifies the resource identifier of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TenantResource.
func (t TenantResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "properties", t.Properties)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
