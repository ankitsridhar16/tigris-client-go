// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: server/v1/search.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	CreateOrUpdateIndex(ctx context.Context, in *CreateOrUpdateIndexRequest, opts ...grpc.CallOption) (*CreateOrUpdateIndexResponse, error)
	GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error)
	// Retrieves one or more documents by id. The response is an array of documents in the same order it is requests.
	// A null is returned for the documents that are not found.
	Get(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	// CreateById is used for indexing a single document. The API expects a single document. An "id" is optional
	// and the server can automatically generate it for you in case it is missing. In cases an id is provided in
	// the document and the document already exists then that document will not be indexed and an error is returned
	// with HTTP status code 409.
	CreateById(ctx context.Context, in *CreateByIdRequest, opts ...grpc.CallOption) (*CreateByIdResponse, error)
	// Create is used for indexing a single or multiple documents. The API expects an array of documents.
	// Each document is a JSON object. An "id" is optional and the server can automatically generate it for you in
	// case it is missing. In cases when an id is provided in the document and the document already exists then that
	// document will not be indexed and in the response there will be an error corresponding to that document id other
	// documents will succeed. Returns an array of status indicating the status of each document.
	Create(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*CreateDocumentResponse, error)
	// Creates or replaces one or more documents. Each document is a JSON object. A document is replaced
	// if it already exists. An "id" is generated automatically in case it is missing in the document. The
	// document is created if "id" doesn't exists otherwise it is replaced. Returns an array of status indicating
	// the status of each document.
	CreateOrReplace(ctx context.Context, in *CreateOrReplaceDocumentRequest, opts ...grpc.CallOption) (*CreateOrReplaceDocumentResponse, error)
	// Updates one or more documents by "id". Each document is required to have the
	// "id" field in it. Returns an array of status indicating the status of each document. Each status
	// has an error field that is set to null in case document is updated successfully otherwise the error
	// field is set with a code and message.
	Update(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*UpdateDocumentResponse, error)
	// Delete one or more documents by id. Returns an array of status indicating the status of each document. Each status
	// has an error field that is set to null in case document is deleted successfully otherwise it will non null with
	// an error code and message.
	Delete(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error)
	// DeleteByQuery is used to delete documents that match the filter. A filter is required. To delete document by id,
	// you can pass the filter as follows ```{"id": "test"}```. Returns a count of number of documents deleted.
	DeleteByQuery(ctx context.Context, in *DeleteByQueryRequest, opts ...grpc.CallOption) (*DeleteByQueryResponse, error)
	// Searches an index for the documents matching the query. A search can be a term search or a phrase search.
	// Search API allows filtering the result set using filters as documented
	// <a href="https://docs.tigrisdata.com/overview/query#specification-1" title="here">here</a>. You can also perform
	// a faceted search by passing the fields in the facet parameter. You can find more detailed documentation of the
	// Search API with multiple examples <a href="https://docs.tigrisdata.com/overview/search" title="here">here</a>.
	Search(ctx context.Context, in *SearchIndexRequest, opts ...grpc.CallOption) (Search_SearchClient, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) CreateOrUpdateIndex(ctx context.Context, in *CreateOrUpdateIndexRequest, opts ...grpc.CallOption) (*CreateOrUpdateIndexResponse, error) {
	out := new(CreateOrUpdateIndexResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/CreateOrUpdateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error) {
	out := new(GetIndexResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/GetIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error) {
	out := new(ListIndexesResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/ListIndexes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Get(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) CreateById(ctx context.Context, in *CreateByIdRequest, opts ...grpc.CallOption) (*CreateByIdResponse, error) {
	out := new(CreateByIdResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/CreateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Create(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*CreateDocumentResponse, error) {
	out := new(CreateDocumentResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) CreateOrReplace(ctx context.Context, in *CreateOrReplaceDocumentRequest, opts ...grpc.CallOption) (*CreateOrReplaceDocumentResponse, error) {
	out := new(CreateOrReplaceDocumentResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/CreateOrReplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Update(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*UpdateDocumentResponse, error) {
	out := new(UpdateDocumentResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Delete(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error) {
	out := new(DeleteDocumentResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) DeleteByQuery(ctx context.Context, in *DeleteByQueryRequest, opts ...grpc.CallOption) (*DeleteByQueryResponse, error) {
	out := new(DeleteByQueryResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.search.v1.Search/DeleteByQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Search(ctx context.Context, in *SearchIndexRequest, opts ...grpc.CallOption) (Search_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Search_ServiceDesc.Streams[0], "/tigrisdata.search.v1.Search/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Search_SearchClient interface {
	Recv() (*SearchIndexResponse, error)
	grpc.ClientStream
}

type searchSearchClient struct {
	grpc.ClientStream
}

func (x *searchSearchClient) Recv() (*SearchIndexResponse, error) {
	m := new(SearchIndexResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchServer is the server API for Search service.
// All implementations should embed UnimplementedSearchServer
// for forward compatibility
type SearchServer interface {
	CreateOrUpdateIndex(context.Context, *CreateOrUpdateIndexRequest) (*CreateOrUpdateIndexResponse, error)
	GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error)
	// Retrieves one or more documents by id. The response is an array of documents in the same order it is requests.
	// A null is returned for the documents that are not found.
	Get(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	// CreateById is used for indexing a single document. The API expects a single document. An "id" is optional
	// and the server can automatically generate it for you in case it is missing. In cases an id is provided in
	// the document and the document already exists then that document will not be indexed and an error is returned
	// with HTTP status code 409.
	CreateById(context.Context, *CreateByIdRequest) (*CreateByIdResponse, error)
	// Create is used for indexing a single or multiple documents. The API expects an array of documents.
	// Each document is a JSON object. An "id" is optional and the server can automatically generate it for you in
	// case it is missing. In cases when an id is provided in the document and the document already exists then that
	// document will not be indexed and in the response there will be an error corresponding to that document id other
	// documents will succeed. Returns an array of status indicating the status of each document.
	Create(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error)
	// Creates or replaces one or more documents. Each document is a JSON object. A document is replaced
	// if it already exists. An "id" is generated automatically in case it is missing in the document. The
	// document is created if "id" doesn't exists otherwise it is replaced. Returns an array of status indicating
	// the status of each document.
	CreateOrReplace(context.Context, *CreateOrReplaceDocumentRequest) (*CreateOrReplaceDocumentResponse, error)
	// Updates one or more documents by "id". Each document is required to have the
	// "id" field in it. Returns an array of status indicating the status of each document. Each status
	// has an error field that is set to null in case document is updated successfully otherwise the error
	// field is set with a code and message.
	Update(context.Context, *UpdateDocumentRequest) (*UpdateDocumentResponse, error)
	// Delete one or more documents by id. Returns an array of status indicating the status of each document. Each status
	// has an error field that is set to null in case document is deleted successfully otherwise it will non null with
	// an error code and message.
	Delete(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	// DeleteByQuery is used to delete documents that match the filter. A filter is required. To delete document by id,
	// you can pass the filter as follows ```{"id": "test"}```. Returns a count of number of documents deleted.
	DeleteByQuery(context.Context, *DeleteByQueryRequest) (*DeleteByQueryResponse, error)
	// Searches an index for the documents matching the query. A search can be a term search or a phrase search.
	// Search API allows filtering the result set using filters as documented
	// <a href="https://docs.tigrisdata.com/overview/query#specification-1" title="here">here</a>. You can also perform
	// a faceted search by passing the fields in the facet parameter. You can find more detailed documentation of the
	// Search API with multiple examples <a href="https://docs.tigrisdata.com/overview/search" title="here">here</a>.
	Search(*SearchIndexRequest, Search_SearchServer) error
}

// UnimplementedSearchServer should be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (UnimplementedSearchServer) CreateOrUpdateIndex(context.Context, *CreateOrUpdateIndexRequest) (*CreateOrUpdateIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateIndex not implemented")
}
func (UnimplementedSearchServer) GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndex not implemented")
}
func (UnimplementedSearchServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedSearchServer) ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIndexes not implemented")
}
func (UnimplementedSearchServer) Get(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSearchServer) CreateById(context.Context, *CreateByIdRequest) (*CreateByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateById not implemented")
}
func (UnimplementedSearchServer) Create(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSearchServer) CreateOrReplace(context.Context, *CreateOrReplaceDocumentRequest) (*CreateOrReplaceDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrReplace not implemented")
}
func (UnimplementedSearchServer) Update(context.Context, *UpdateDocumentRequest) (*UpdateDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSearchServer) Delete(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSearchServer) DeleteByQuery(context.Context, *DeleteByQueryRequest) (*DeleteByQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByQuery not implemented")
}
func (UnimplementedSearchServer) Search(*SearchIndexRequest, Search_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_CreateOrUpdateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).CreateOrUpdateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/CreateOrUpdateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).CreateOrUpdateIndex(ctx, req.(*CreateOrUpdateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_GetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).GetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/GetIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).GetIndex(ctx, req.(*GetIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_ListIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).ListIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/ListIndexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).ListIndexes(ctx, req.(*ListIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Get(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_CreateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).CreateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/CreateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).CreateById(ctx, req.(*CreateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Create(ctx, req.(*CreateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_CreateOrReplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrReplaceDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).CreateOrReplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/CreateOrReplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).CreateOrReplace(ctx, req.(*CreateOrReplaceDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Update(ctx, req.(*UpdateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Delete(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_DeleteByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).DeleteByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.search.v1.Search/DeleteByQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).DeleteByQuery(ctx, req.(*DeleteByQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchIndexRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServer).Search(m, &searchSearchServer{stream})
}

type Search_SearchServer interface {
	Send(*SearchIndexResponse) error
	grpc.ServerStream
}

type searchSearchServer struct {
	grpc.ServerStream
}

func (x *searchSearchServer) Send(m *SearchIndexResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tigrisdata.search.v1.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateIndex",
			Handler:    _Search_CreateOrUpdateIndex_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _Search_GetIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _Search_DeleteIndex_Handler,
		},
		{
			MethodName: "ListIndexes",
			Handler:    _Search_ListIndexes_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Search_Get_Handler,
		},
		{
			MethodName: "CreateById",
			Handler:    _Search_CreateById_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Search_Create_Handler,
		},
		{
			MethodName: "CreateOrReplace",
			Handler:    _Search_CreateOrReplace_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Search_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Search_Delete_Handler,
		},
		{
			MethodName: "DeleteByQuery",
			Handler:    _Search_DeleteByQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _Search_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/v1/search.proto",
}
