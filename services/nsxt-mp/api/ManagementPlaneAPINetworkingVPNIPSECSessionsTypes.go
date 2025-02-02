// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingVPNIPSECSessions.
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

// Possible value for ``sessionType`` of method ManagementPlaneAPINetworkingVPNIPSECSessions#listIPSecVPNSessions.
const ManagementPlaneAPINetworkingVPNIPSECSessions_LIST_IPSEC_VPNSESSIONS_SESSION_TYPE_POLICYBASEDIPSECVPNSESSION = "PolicyBasedIPSecVPNSession"

// Possible value for ``sessionType`` of method ManagementPlaneAPINetworkingVPNIPSECSessions#listIPSecVPNSessions.
const ManagementPlaneAPINetworkingVPNIPSECSessions_LIST_IPSEC_VPNSESSIONS_SESSION_TYPE_ROUTEBASEDIPSECVPNSESSION = "RouteBasedIPSecVPNSession"

func managementPlaneAPINetworkingVPNIPSECSessionsCreateIPSecVPNSessionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_sec_VPN_session"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
	fieldNameMap["ip_sec_VPN_session"] = "IpSecVPNSession"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsCreateIPSecVPNSessionOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingVPNIPSECSessionsCreateIPSecVPNSessionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_sec_VPN_session"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
	fieldNameMap["ip_sec_VPN_session"] = "IpSecVPNSession"
	paramsTypeMap["ip_sec_VPN_session"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
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
		"ip_sec_VPN_session",
		"POST",
		"/api/v1/vpn/ipsec/sessions",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECSessionsDeleteIPSecVPNSessionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsDeleteIPSecVPNSessionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingVPNIPSECSessionsDeleteIPSecVPNSessionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["ipsec_vpn_session_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnSessionId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_session_id"] = "ipsecVpnSessionId"
	queryParams["force"] = "force"
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
		"DELETE",
		"/api/v1/vpn/ipsec/sessions/{ipsecVpnSessionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetIPSecVPNSessionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetIPSecVPNSessionOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetIPSecVPNSessionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	paramsTypeMap["ipsec_vpn_session_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnSessionId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_session_id"] = "ipsecVpnSessionId"
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
		"/api/v1/vpn/ipsec/sessions/{ipsecVpnSessionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetIPSecVPNSessionStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetIPSecVPNSessionStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNSessionStateBindingType)
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetIPSecVPNSessionStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	paramsTypeMap["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ipsec_vpn_session_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnSessionId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_session_id"] = "ipsecVpnSessionId"
	queryParams["barrier_id"] = "barrier_id"
	queryParams["request_id"] = "request_id"
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
		"/api/v1/vpn/ipsec/sessions/{ipsecVpnSessionId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetPeerConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetPeerConfigOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func managementPlaneAPINetworkingVPNIPSECSessionsGetPeerConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	paramsTypeMap["ipsec_vpn_session_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnSessionId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_session_id"] = "ipsecVpnSessionId"
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
		"/api/v1/vpn/ipsec/sessions/{ipsecVpnSessionId}/peer-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECSessionsListIPSecVPNSessionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ipsec_vpn_service_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["session_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ipsec_vpn_service_id"] = "IpsecVpnServiceId"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["session_type"] = "SessionType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsListIPSecVPNSessionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNSessionListResultBindingType)
}

func managementPlaneAPINetworkingVPNIPSECSessionsListIPSecVPNSessionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ipsec_vpn_service_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["session_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ipsec_vpn_service_id"] = "IpsecVpnServiceId"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["session_type"] = "SessionType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["session_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["ipsec_vpn_service_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["logical_router_id"] = "logical_router_id"
	queryParams["ipsec_vpn_service_id"] = "ipsec_vpn_service_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["session_type"] = "session_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/vpn/ipsec/sessions",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECSessionsUpdateIPSecVPNSessionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fields["ip_sec_VPN_session"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	fieldNameMap["ip_sec_VPN_session"] = "IpSecVPNSession"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECSessionsUpdateIPSecVPNSessionOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingVPNIPSECSessionsUpdateIPSecVPNSessionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_session_id"] = bindings.NewStringType()
	fields["ip_sec_VPN_session"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
	fieldNameMap["ipsec_vpn_session_id"] = "IpsecVpnSessionId"
	fieldNameMap["ip_sec_VPN_session"] = "IpSecVPNSession"
	paramsTypeMap["ip_sec_VPN_session"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IPSecVPNSessionBindingType)}, bindings.REST)
	paramsTypeMap["ipsec_vpn_session_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnSessionId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_session_id"] = "ipsecVpnSessionId"
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
		"ip_sec_VPN_session",
		"PUT",
		"/api/v1/vpn/ipsec/sessions/{ipsecVpnSessionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
