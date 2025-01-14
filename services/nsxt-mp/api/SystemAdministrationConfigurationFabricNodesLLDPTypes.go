// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricNodesLLDP.
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

func systemAdministrationConfigurationFabricNodesLLDPListFabricNodeNeighborPropertiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fabric_node_id"] = bindings.NewStringType()
	fieldNameMap["fabric_node_id"] = "FabricNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesLLDPListFabricNodeNeighborPropertiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InterfaceNeighborPropertyListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesLLDPListFabricNodeNeighborPropertiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["fabric_node_id"] = bindings.NewStringType()
	fieldNameMap["fabric_node_id"] = "FabricNodeId"
	paramsTypeMap["fabric_node_id"] = bindings.NewStringType()
	paramsTypeMap["fabricNodeId"] = bindings.NewStringType()
	pathParams["fabric_node_id"] = "fabricNodeId"
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
		"/api/v1/lldp/fabric-nodes/{fabricNodeId}/interfaces",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesLLDPListNeighborPropertiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesLLDPListNeighborPropertiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InterfaceNeighborPropertyListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesLLDPListNeighborPropertiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
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
		"/api/v1/lldp/transport-nodes/{nodeId}/interfaces",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesLLDPReadFabricNodeNeighborPropertiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fabric_node_id"] = bindings.NewStringType()
	fields["interface_name"] = bindings.NewStringType()
	fieldNameMap["fabric_node_id"] = "FabricNodeId"
	fieldNameMap["interface_name"] = "InterfaceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesLLDPReadFabricNodeNeighborPropertiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InterfaceNeighborPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesLLDPReadFabricNodeNeighborPropertiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["fabric_node_id"] = bindings.NewStringType()
	fields["interface_name"] = bindings.NewStringType()
	fieldNameMap["fabric_node_id"] = "FabricNodeId"
	fieldNameMap["interface_name"] = "InterfaceName"
	paramsTypeMap["fabric_node_id"] = bindings.NewStringType()
	paramsTypeMap["interface_name"] = bindings.NewStringType()
	paramsTypeMap["fabricNodeId"] = bindings.NewStringType()
	paramsTypeMap["interfaceName"] = bindings.NewStringType()
	pathParams["interface_name"] = "interfaceName"
	pathParams["fabric_node_id"] = "fabricNodeId"
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
		"/api/v1/lldp/fabric-nodes/{fabricNodeId}/interfaces/{interfaceName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesLLDPReadNeighborPropertiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["interface_name"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["interface_name"] = "InterfaceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesLLDPReadNeighborPropertiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InterfaceNeighborPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesLLDPReadNeighborPropertiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["interface_name"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["interface_name"] = "InterfaceName"
	paramsTypeMap["interface_name"] = bindings.NewStringType()
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	paramsTypeMap["interfaceName"] = bindings.NewStringType()
	pathParams["interface_name"] = "interfaceName"
	pathParams["node_id"] = "nodeId"
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
		"/api/v1/lldp/transport-nodes/{nodeId}/interfaces/{interfaceName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
