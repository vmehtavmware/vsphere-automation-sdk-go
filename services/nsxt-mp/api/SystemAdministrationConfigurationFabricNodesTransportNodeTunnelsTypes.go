// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricNodesTransportNodeTunnels.
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

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#getTunnel.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_GET_TUNNEL_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#getTunnel.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_GET_TUNNEL_SOURCE_CACHED = "cached"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_0 = "0"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_NO_DIAGNOSTIC = "NO_DIAGNOSTIC"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_1 = "1"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_CONTROL_DETECTION_TIME_EXPIRED = "CONTROL_DETECTION_TIME_EXPIRED"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_2 = "2"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_ECHO_FUNCTION_FAILED = "ECHO_FUNCTION_FAILED"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_3 = "3"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_NEIGHBOR_SIGNALED_SESSION_DOWN = "NEIGHBOR_SIGNALED_SESSION_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_4 = "4"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_FORWARDING_PLANE_RESET = "FORWARDING_PLANE_RESET"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_5 = "5"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_PATH_DOWN = "PATH_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_6 = "6"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_CONCATENATED_PATH_DOWN = "CONCATENATED_PATH_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_7 = "7"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_ADMINISTRATIVELY_DOWN = "ADMINISTRATIVELY_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_8 = "8"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_BFD_DIAGNOSTIC_CODE_REVERSE_CONCATENATED_PATH_DOWN = "REVERSE_CONCATENATED_PATH_DOWN"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_SOURCE_CACHED = "cached"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_STATUS_UP = "UP"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeTunnels#queryTunnels.
const SystemAdministrationConfigurationFabricNodesTransportNodeTunnels_QUERY_TUNNELS_STATUS_DOWN = "DOWN"

func systemAdministrationConfigurationFabricNodesTransportNodeTunnelsGetTunnelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["tunnel_name"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["tunnel_name"] = "TunnelName"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeTunnelsGetTunnelOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TunnelPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeTunnelsGetTunnelRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["tunnel_name"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["tunnel_name"] = "TunnelName"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["tunnel_name"] = bindings.NewStringType()
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	paramsTypeMap["tunnelName"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	pathParams["tunnel_name"] = "tunnelName"
	queryParams["source"] = "source"
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
		"/api/v1/transport-nodes/{nodeId}/tunnels/{tunnelName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeTunnelsQueryTunnelsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["bfd_diagnostic_code"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["remote_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["bfd_diagnostic_code"] = "BfdDiagnosticCode"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["remote_node_id"] = "RemoteNodeId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeTunnelsQueryTunnelsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TunnelListBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeTunnelsQueryTunnelsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["bfd_diagnostic_code"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["remote_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["bfd_diagnostic_code"] = "BfdDiagnosticCode"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["remote_node_id"] = "RemoteNodeId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["remote_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["bfd_diagnostic_code"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["bfd_diagnostic_code"] = "bfd_diagnostic_code"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["page_size"] = "page_size"
	queryParams["remote_node_id"] = "remote_node_id"
	queryParams["status"] = "status"
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
		"/api/v1/transport-nodes/{nodeId}/tunnels",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
