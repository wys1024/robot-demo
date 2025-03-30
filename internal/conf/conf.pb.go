// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.3
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Server        *Server                `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Llm           *Llm                   `protobuf:"bytes,3,opt,name=llm,proto3" json:"llm,omitempty"`
	Jwt           *JWT                   `protobuf:"bytes,4,opt,name=jwt,proto3" json:"jwt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_conf_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetLlm() *Llm {
	if x != nil {
		return x.Llm
	}
	return nil
}

func (x *Bootstrap) GetJwt() *JWT {
	if x != nil {
		return x.Jwt
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *Data_Redis            `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_conf_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Llm struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Model         string                 `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	ApiHost       string                 `protobuf:"bytes,2,opt,name=api_host,json=apiHost,proto3" json:"api_host,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Llm) Reset() {
	*x = Llm{}
	mi := &file_conf_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Llm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Llm) ProtoMessage() {}

func (x *Llm) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Llm.ProtoReflect.Descriptor instead.
func (*Llm) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Llm) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Llm) GetApiHost() string {
	if x != nil {
		return x.ApiHost
	}
	return ""
}

type JWT struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secret        string                 `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JWT) Reset() {
	*x = JWT{}
	mi := &file_conf_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JWT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWT) ProtoMessage() {}

func (x *JWT) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWT.ProtoReflect.Descriptor instead.
func (*JWT) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *JWT) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type Server_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_conf_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout   *durationpb.Duration   `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout  *durationpb.Duration   `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	mi := &file_conf_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

var File_conf_conf_proto protoreflect.FileDescriptor

const file_conf_conf_proto_rawDesc = "" +
	"\n" +
	"\x0fconf/conf.proto\x12\n" +
	"kratos.api\x1a\x1egoogle/protobuf/duration.proto\"\xa3\x01\n" +
	"\tBootstrap\x12*\n" +
	"\x06server\x18\x01 \x01(\v2\x12.kratos.api.ServerR\x06server\x12$\n" +
	"\x04data\x18\x02 \x01(\v2\x10.kratos.api.DataR\x04data\x12!\n" +
	"\x03llm\x18\x03 \x01(\v2\x0f.kratos.api.LlmR\x03llm\x12!\n" +
	"\x03jwt\x18\x04 \x01(\v2\x0f.kratos.api.JWTR\x03jwt\"\xb8\x02\n" +
	"\x06Server\x12+\n" +
	"\x04http\x18\x01 \x01(\v2\x17.kratos.api.Server.HTTPR\x04http\x12+\n" +
	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x1ai\n" +
	"\x04HTTP\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\x1ai\n" +
	"\x04GRPC\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"\xdd\x02\n" +
	"\x04Data\x125\n" +
	"\bdatabase\x18\x01 \x01(\v2\x19.kratos.api.Data.DatabaseR\bdatabase\x12,\n" +
	"\x05redis\x18\x02 \x01(\v2\x16.kratos.api.Data.RedisR\x05redis\x1a:\n" +
	"\bDatabase\x12\x16\n" +
	"\x06driver\x18\x01 \x01(\tR\x06driver\x12\x16\n" +
	"\x06source\x18\x02 \x01(\tR\x06source\x1a\xb3\x01\n" +
	"\x05Redis\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x12<\n" +
	"\fread_timeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\vreadTimeout\x12>\n" +
	"\rwrite_timeout\x18\x04 \x01(\v2\x19.google.protobuf.DurationR\fwriteTimeout\"6\n" +
	"\x03Llm\x12\x14\n" +
	"\x05model\x18\x01 \x01(\tR\x05model\x12\x19\n" +
	"\bapi_host\x18\x02 \x01(\tR\aapiHost\"\x1d\n" +
	"\x03JWT\x12\x16\n" +
	"\x06secret\x18\x01 \x01(\tR\x06secretB\x1fZ\x1drobot-demo/internal/conf;confb\x06proto3"

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData []byte
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)))
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_conf_conf_proto_goTypes = []any{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Data)(nil),                // 2: kratos.api.Data
	(*Llm)(nil),                 // 3: kratos.api.Llm
	(*JWT)(nil),                 // 4: kratos.api.JWT
	(*Server_HTTP)(nil),         // 5: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 6: kratos.api.Server.GRPC
	(*Data_Database)(nil),       // 7: kratos.api.Data.Database
	(*Data_Redis)(nil),          // 8: kratos.api.Data.Redis
	(*durationpb.Duration)(nil), // 9: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	2,  // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	3,  // 2: kratos.api.Bootstrap.llm:type_name -> kratos.api.Llm
	4,  // 3: kratos.api.Bootstrap.jwt:type_name -> kratos.api.JWT
	5,  // 4: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	6,  // 5: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	7,  // 6: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	8,  // 7: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	9,  // 8: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	9,  // 9: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	9,  // 10: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	9,  // 11: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
