// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPITroubleshootingAndMonitoringPortConnection.
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

func managementPlaneAPITroubleshootingAndMonitoringPortConnectionGetForwardingPathInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["peer_port_id"] = bindings.NewStringType()
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["peer_port_id"] = "PeerPortId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringPortConnectionGetForwardingPathOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortConnectionEntitiesBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringPortConnectionGetForwardingPathRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["peer_port_id"] = bindings.NewStringType()
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["peer_port_id"] = "PeerPortId"
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["peer_port_id"] = bindings.NewStringType()
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["peer_port_id"] = "peer_port_id"
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
		"/api/v1/logical-ports/{lportId}/forwarding-path",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
