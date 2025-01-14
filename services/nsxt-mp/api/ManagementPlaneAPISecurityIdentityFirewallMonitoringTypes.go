// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPISecurityIdentityFirewallMonitoring.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func managementPlaneAPISecurityIdentityFirewallMonitoringGetComputeCollectionStatusByIdInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compute_collection_ext_id"] = bindings.NewStringType()
	fieldNameMap["compute_collection_ext_id"] = "ComputeCollectionExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringGetComputeCollectionStatusByIdOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwComputeCollectionStatusBindingType)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringGetComputeCollectionStatusByIdRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["compute_collection_ext_id"] = bindings.NewStringType()
	fieldNameMap["compute_collection_ext_id"] = "ComputeCollectionExtId"
	paramsTypeMap["compute_collection_ext_id"] = bindings.NewStringType()
	paramsTypeMap["computeCollectionExtId"] = bindings.NewStringType()
	pathParams["compute_collection_ext_id"] = "computeCollectionExtId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/idfw/compute-collections/{computeCollectionExtId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListComputeCollectionStatusesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListComputeCollectionStatusesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwComputeCollectionListResultBindingType)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListComputeCollectionStatusesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/idfw/compute-collections/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListTransportNodeStatusesByComputeCollectionIdInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListTransportNodeStatusesByComputeCollectionIdOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwTransportNodeStatusListResultBindingType)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListTransportNodeStatusesByComputeCollectionIdRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/idfw/compute-collections/{ccExtId}/transport-nodes/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListVirtualMachineStatusesByTransportNodeIdInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListVirtualMachineStatusesByTransportNodeIdOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwVirtualMachineStatusListResultBindingType)
}

func managementPlaneAPISecurityIdentityFirewallMonitoringListVirtualMachineStatusesByTransportNodeIdRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/idfw/transport-nodes/{transportNodeId}/vms/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
