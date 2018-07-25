// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apiserver/apiserver.proto

/*
Package apiserver is a generated protocol buffer package.

It is generated from these files:
	pkg/apiserver/apiserver.proto

It has these top-level messages:
	WorkflowIdentifier
	SearchWorkflowResponse
	InvocationListQuery
	WorkflowInvocationIdentifier
	WorkflowInvocationList
	Health
*/
package apiserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fission_workflows_types "github.com/fission/fission-workflows/pkg/types"
import fission_workflows_version "github.com/fission/fission-workflows/pkg/version"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowIdentifier struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowIdentifier) Reset()                    { *m = WorkflowIdentifier{} }
func (m *WorkflowIdentifier) String() string            { return proto.CompactTextString(m) }
func (*WorkflowIdentifier) ProtoMessage()               {}
func (*WorkflowIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WorkflowIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchWorkflowResponse struct {
	Workflows []string `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
}

func (m *SearchWorkflowResponse) Reset()                    { *m = SearchWorkflowResponse{} }
func (m *SearchWorkflowResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchWorkflowResponse) ProtoMessage()               {}
func (*SearchWorkflowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchWorkflowResponse) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type InvocationListQuery struct {
	Workflows []string `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
}

func (m *InvocationListQuery) Reset()                    { *m = InvocationListQuery{} }
func (m *InvocationListQuery) String() string            { return proto.CompactTextString(m) }
func (*InvocationListQuery) ProtoMessage()               {}
func (*InvocationListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InvocationListQuery) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type WorkflowInvocationIdentifier struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowInvocationIdentifier) Reset()                    { *m = WorkflowInvocationIdentifier{} }
func (m *WorkflowInvocationIdentifier) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationIdentifier) ProtoMessage()               {}
func (*WorkflowInvocationIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkflowInvocationIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WorkflowInvocationList struct {
	Invocations []string `protobuf:"bytes,1,rep,name=invocations" json:"invocations,omitempty"`
}

func (m *WorkflowInvocationList) Reset()                    { *m = WorkflowInvocationList{} }
func (m *WorkflowInvocationList) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationList) ProtoMessage()               {}
func (*WorkflowInvocationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WorkflowInvocationList) GetInvocations() []string {
	if m != nil {
		return m.Invocations
	}
	return nil
}

type Health struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Health) Reset()                    { *m = Health{} }
func (m *Health) String() string            { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()               {}
func (*Health) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowIdentifier)(nil), "fission.workflows.apiserver.WorkflowIdentifier")
	proto.RegisterType((*SearchWorkflowResponse)(nil), "fission.workflows.apiserver.SearchWorkflowResponse")
	proto.RegisterType((*InvocationListQuery)(nil), "fission.workflows.apiserver.InvocationListQuery")
	proto.RegisterType((*WorkflowInvocationIdentifier)(nil), "fission.workflows.apiserver.WorkflowInvocationIdentifier")
	proto.RegisterType((*WorkflowInvocationList)(nil), "fission.workflows.apiserver.WorkflowInvocationList")
	proto.RegisterType((*Health)(nil), "fission.workflows.apiserver.Health")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkflowAPI service

type WorkflowAPIClient interface {
	Create(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*WorkflowIdentifier, error)
	List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*SearchWorkflowResponse, error)
	Get(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.Workflow, error)
	Delete(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Validate(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type workflowAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowAPIClient(cc *grpc.ClientConn) WorkflowAPIClient {
	return &workflowAPIClient{cc}
}

func (c *workflowAPIClient) Create(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*WorkflowIdentifier, error) {
	out := new(WorkflowIdentifier)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*SearchWorkflowResponse, error) {
	out := new(SearchWorkflowResponse)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Get(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.Workflow, error) {
	out := new(fission_workflows_types.Workflow)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Delete(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Validate(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowAPI service

type WorkflowAPIServer interface {
	Create(context.Context, *fission_workflows_types.WorkflowSpec) (*WorkflowIdentifier, error)
	List(context.Context, *google_protobuf1.Empty) (*SearchWorkflowResponse, error)
	Get(context.Context, *WorkflowIdentifier) (*fission_workflows_types.Workflow, error)
	Delete(context.Context, *WorkflowIdentifier) (*google_protobuf1.Empty, error)
	Validate(context.Context, *fission_workflows_types.WorkflowSpec) (*google_protobuf1.Empty, error)
}

func RegisterWorkflowAPIServer(s *grpc.Server, srv WorkflowAPIServer) {
	s.RegisterService(&_WorkflowAPI_serviceDesc, srv)
}

func _WorkflowAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Create(ctx, req.(*fission_workflows_types.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).List(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Get(ctx, req.(*WorkflowIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Delete(ctx, req.(*WorkflowIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Validate(ctx, req.(*fission_workflows_types.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.WorkflowAPI",
	HandlerType: (*WorkflowAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WorkflowAPI_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowAPI_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WorkflowAPI_Delete_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for WorkflowInvocationAPI service

type WorkflowInvocationAPIClient interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*WorkflowInvocationIdentifier, error)
	InvokeSync(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	List(ctx context.Context, in *InvocationListQuery, opts ...grpc.CallOption) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error)
	Validate(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type workflowInvocationAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowInvocationAPIClient(cc *grpc.ClientConn) WorkflowInvocationAPIClient {
	return &workflowInvocationAPIClient{cc}
}

func (c *workflowInvocationAPIClient) Invoke(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*WorkflowInvocationIdentifier, error) {
	out := new(WorkflowInvocationIdentifier)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) InvokeSync(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error) {
	out := new(fission_workflows_types.WorkflowInvocation)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/InvokeSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Cancel(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) List(ctx context.Context, in *InvocationListQuery, opts ...grpc.CallOption) (*WorkflowInvocationList, error) {
	out := new(WorkflowInvocationList)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Get(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error) {
	out := new(fission_workflows_types.WorkflowInvocation)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Validate(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowInvocationAPI service

type WorkflowInvocationAPIServer interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(context.Context, *fission_workflows_types.WorkflowInvocationSpec) (*WorkflowInvocationIdentifier, error)
	InvokeSync(context.Context, *fission_workflows_types.WorkflowInvocationSpec) (*fission_workflows_types.WorkflowInvocation, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(context.Context, *WorkflowInvocationIdentifier) (*google_protobuf1.Empty, error)
	List(context.Context, *InvocationListQuery) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(context.Context, *WorkflowInvocationIdentifier) (*fission_workflows_types.WorkflowInvocation, error)
	Validate(context.Context, *fission_workflows_types.WorkflowInvocationSpec) (*google_protobuf1.Empty, error)
}

func RegisterWorkflowInvocationAPIServer(s *grpc.Server, srv WorkflowInvocationAPIServer) {
	s.RegisterService(&_WorkflowInvocationAPI_serviceDesc, srv)
}

func _WorkflowInvocationAPI_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, req.(*fission_workflows_types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_InvokeSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/InvokeSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, req.(*fission_workflows_types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowInvocationIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, req.(*WorkflowInvocationIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).List(ctx, req.(*InvocationListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowInvocationIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, req.(*WorkflowInvocationIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Validate(ctx, req.(*fission_workflows_types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowInvocationAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.WorkflowInvocationAPI",
	HandlerType: (*WorkflowInvocationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _WorkflowInvocationAPI_Invoke_Handler,
		},
		{
			MethodName: "InvokeSync",
			Handler:    _WorkflowInvocationAPI_InvokeSync_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _WorkflowInvocationAPI_Cancel_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowInvocationAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowInvocationAPI_Get_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowInvocationAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for AdminAPI service

type AdminAPIClient interface {
	Status(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Health, error)
	Version(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*fission_workflows_version.Info, error)
}

type adminAPIClient struct {
	cc *grpc.ClientConn
}

func NewAdminAPIClient(cc *grpc.ClientConn) AdminAPIClient {
	return &adminAPIClient{cc}
}

func (c *adminAPIClient) Status(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.AdminAPI/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAPIClient) Version(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*fission_workflows_version.Info, error) {
	out := new(fission_workflows_version.Info)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.AdminAPI/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdminAPI service

type AdminAPIServer interface {
	Status(context.Context, *google_protobuf1.Empty) (*Health, error)
	Version(context.Context, *google_protobuf1.Empty) (*fission_workflows_version.Info, error)
}

func RegisterAdminAPIServer(s *grpc.Server, srv AdminAPIServer) {
	s.RegisterService(&_AdminAPI_serviceDesc, srv)
}

func _AdminAPI_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.AdminAPI/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Status(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.AdminAPI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Version(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.AdminAPI",
	HandlerType: (*AdminAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AdminAPI_Status_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _AdminAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

func init() { proto.RegisterFile("pkg/apiserver/apiserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe5, 0xb6, 0xf2, 0xaf, 0x99, 0xe8, 0x57, 0x95, 0x69, 0x09, 0x25, 0x6d, 0xd5, 0xb0,
	0x80, 0x54, 0x82, 0xf0, 0xa2, 0x56, 0x42, 0x22, 0x07, 0xa4, 0x52, 0x10, 0x44, 0xe2, 0x00, 0x0d,
	0x6a, 0xa5, 0xde, 0x5c, 0x67, 0x93, 0xac, 0x9a, 0x7a, 0x8d, 0xbd, 0x49, 0x15, 0x10, 0x97, 0x5e,
	0xb8, 0x70, 0xe3, 0xc8, 0x73, 0x70, 0xe0, 0xc6, 0x3b, 0xf0, 0x0a, 0x3c, 0x08, 0xca, 0x7a, 0xed,
	0xb8, 0xd8, 0x4e, 0x6a, 0xb8, 0xc4, 0xf1, 0xec, 0xce, 0x7c, 0xbe, 0x3b, 0x7f, 0xd6, 0xb0, 0xe9,
	0x9d, 0x76, 0xa9, 0xed, 0xf1, 0x80, 0xf9, 0x43, 0xe6, 0x4f, 0xfe, 0x59, 0x9e, 0x2f, 0xa4, 0xc0,
	0xf5, 0x0e, 0x0f, 0x02, 0x2e, 0x5c, 0xeb, 0x5c, 0xf8, 0xa7, 0x9d, 0xbe, 0x38, 0x0f, 0xac, 0x78,
	0x4b, 0xb5, 0xd1, 0xe5, 0xb2, 0x37, 0x38, 0xb1, 0x1c, 0x71, 0x46, 0xf5, 0xbe, 0xe8, 0xf9, 0x20,
	0xde, 0x4f, 0xc7, 0x00, 0x39, 0xf2, 0x58, 0x10, 0xfe, 0x86, 0x81, 0xab, 0x4f, 0xae, 0xec, 0x3b,
	0x64, 0xbe, 0x5a, 0xd5, 0x4f, 0xed, 0xbf, 0xde, 0x15, 0xa2, 0xdb, 0x67, 0x54, 0xbd, 0x9d, 0x0c,
	0x3a, 0x94, 0x9d, 0x79, 0x72, 0xa4, 0x17, 0x37, 0xf4, 0xa2, 0xed, 0x71, 0x6a, 0xbb, 0xae, 0x90,
	0xb6, 0xe4, 0xc2, 0xd5, 0x68, 0x72, 0x07, 0xf0, 0x48, 0x13, 0x9a, 0x6d, 0xe6, 0x4a, 0xde, 0xe1,
	0xcc, 0xc7, 0x25, 0x98, 0xe3, 0xed, 0x35, 0xa3, 0x66, 0x6c, 0x97, 0x0e, 0xe6, 0x78, 0x9b, 0x3c,
	0x82, 0x4a, 0x8b, 0xd9, 0xbe, 0xd3, 0x8b, 0xf6, 0x1e, 0xb0, 0xc0, 0x13, 0x6e, 0xc0, 0x70, 0x03,
	0x4a, 0xb1, 0xc2, 0x35, 0xa3, 0x36, 0xbf, 0x5d, 0x3a, 0x98, 0x18, 0xc8, 0x2e, 0xac, 0x34, 0xdd,
	0xa1, 0x70, 0x14, 0xf2, 0x15, 0x0f, 0xe4, 0x9b, 0x01, 0xf3, 0x47, 0x33, 0x9c, 0x2c, 0xd8, 0x88,
	0x25, 0xc5, 0xce, 0x53, 0xc4, 0x35, 0xa0, 0x92, 0xde, 0x3f, 0x86, 0x61, 0x0d, 0xca, 0x3c, 0xb6,
	0x44, 0xa4, 0xa4, 0x89, 0xd4, 0xc0, 0x7c, 0xc9, 0xec, 0xbe, 0xec, 0x61, 0x05, 0xcc, 0x40, 0xda,
	0x72, 0x10, 0xe8, 0xc8, 0xfa, 0x6d, 0xe7, 0xdb, 0x02, 0x94, 0xa3, 0xf0, 0x7b, 0xaf, 0x9b, 0x38,
	0x04, 0x73, 0xdf, 0x67, 0xb6, 0x64, 0x78, 0xd7, 0x4a, 0xf7, 0x43, 0x58, 0xd5, 0x68, 0x7f, 0xcb,
	0x63, 0x4e, 0x95, 0x5a, 0x53, 0xda, 0xc6, 0x4a, 0x27, 0x9f, 0xac, 0x5e, 0xfc, 0xfc, 0xf5, 0x65,
	0x6e, 0x89, 0x94, 0x68, 0xe4, 0xd0, 0x30, 0xea, 0xd8, 0x81, 0x05, 0x75, 0xa6, 0x8a, 0x15, 0xd6,
	0xd3, 0x8a, 0x8a, 0x6d, 0x3d, 0x1f, 0x17, 0xbb, 0xba, 0x3b, 0x15, 0x93, 0x5d, 0x3d, 0x72, 0x4d,
	0xa1, 0xca, 0x38, 0x41, 0xe1, 0x3b, 0x98, 0x7f, 0xc1, 0x24, 0x16, 0x55, 0x5d, 0xbd, 0x35, 0x33,
	0x1b, 0xa4, 0xa2, 0x68, 0xcb, 0xb8, 0x14, 0xd3, 0xe8, 0x07, 0xde, 0xfe, 0x88, 0x1c, 0xcc, 0x67,
	0xac, 0xcf, 0x24, 0x2b, 0x4e, 0xcd, 0xc9, 0x46, 0x84, 0xaa, 0xff, 0x89, 0xea, 0xc1, 0xe2, 0xa1,
	0xdd, 0xe7, 0xed, 0x02, 0xf5, 0xcb, 0x43, 0x6c, 0x2a, 0xc4, 0x0d, 0x82, 0x13, 0xc4, 0x50, 0x87,
	0x6e, 0x18, 0xf5, 0x9d, 0x1f, 0x26, 0x5c, 0x4f, 0xb7, 0xe5, 0xb8, 0x83, 0x3e, 0x1b, 0x60, 0x8e,
	0x2d, 0xa7, 0xd9, 0xe7, 0xbd, 0x2c, 0x61, 0xe2, 0xaa, 0xc4, 0x3c, 0xbe, 0x5a, 0x82, 0x32, 0xc6,
	0x26, 0x4a, 0x09, 0x29, 0xd3, 0xc9, 0x00, 0x8c, 0x1b, 0xeb, 0xab, 0x01, 0x10, 0xca, 0x69, 0x8d,
	0x5c, 0xa7, 0xb8, 0xa4, 0xfb, 0x05, 0x1c, 0x08, 0x55, 0x22, 0xee, 0x91, 0xe5, 0x84, 0x08, 0x1a,
	0x8c, 0x5c, 0xa7, 0x61, 0xd4, 0x8f, 0x11, 0x53, 0x66, 0x1c, 0x80, 0xb9, 0x6f, 0xbb, 0x0e, 0xeb,
	0xe3, 0xdf, 0x1f, 0x3d, 0xb7, 0x84, 0x6b, 0x4a, 0x0d, 0xd6, 0x2f, 0x61, 0x55, 0x9f, 0x5c, 0x18,
	0x7a, 0xdc, 0x1e, 0x4e, 0xa5, 0x66, 0x5c, 0x6e, 0x33, 0x06, 0x31, 0xfb, 0xa6, 0x22, 0x2b, 0x4a,
	0xc9, 0xff, 0x98, 0x2c, 0x0e, 0x7e, 0x32, 0xc2, 0x59, 0xfc, 0x87, 0x93, 0x17, 0x2a, 0x8e, 0x4e,
	0x07, 0xa6, 0xd3, 0x21, 0x13, 0x63, 0x53, 0xb8, 0x41, 0xf2, 0xb2, 0xbf, 0xa5, 0x70, 0x37, 0xc9,
	0x6a, 0x12, 0x97, 0x1c, 0xa1, 0xef, 0x06, 0x2c, 0xee, 0xb5, 0xcf, 0xb8, 0x9a, 0x9a, 0x23, 0x30,
	0x5b, 0xea, 0x46, 0xce, 0xbd, 0x01, 0x6f, 0x4f, 0x4d, 0x53, 0x78, 0xcd, 0x93, 0x65, 0x05, 0x05,
	0x5c, 0xa4, 0x3d, 0x65, 0x78, 0x8f, 0x6f, 0xe1, 0xbf, 0xc3, 0xf0, 0x6b, 0x9a, 0x1b, 0x79, 0x2b,
	0x23, 0x72, 0xf4, 0x05, 0x6e, 0xba, 0x1d, 0x91, 0x88, 0xaa, 0xcd, 0x4f, 0xcb, 0xc7, 0xa5, 0x98,
	0x7d, 0x62, 0xaa, 0x78, 0xbb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xd9, 0x98, 0xf7, 0x60,
	0x08, 0x00, 0x00,
}
