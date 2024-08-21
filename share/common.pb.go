// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package share is a generated protocol buffer package.

It is generated from these files:
	common.proto
	controller_service.proto
	enforcer_service.proto
	scan.proto
	scanner_service.proto

It has these top-level messages:
	RPCVoid
	CLUSProfilingRequest
	CLUSBoolean
	CLUSMetry
	CLUSStats
	AdapterScanImageRequest
	GetScannersResponse
	ScannerRegisterData
	ScannerDeregisterData
	ControllerCaps
	ScannerAvailable
	CLUSFilePacket
	CLUSAdmissionRequest
	CLUSAdmissionResponse
	CLUSProcProfileReq
	CLUSProcProfileArray
	CLUSFileAccessRuleReq
	CLUSFileAccessRuleArray
	CLUSConnection
	CLUSConnectionArray
	CLUSReportResponse
	CLUSSyncRequest
	CLUSSyncReply
	CLUSControllerCounter
	CLUSGraphOps
	CLUSPolicyRuleCheck
	CLUSPolicyRuleMismatch
	CLUSPolicySyncStatus
	CLUSStoreWatcherInfo
	CLUSKickLoginSessionsRequest
	CLUSLoginTokenInfo
	CLUSKubernetesResInfo
	CLUSKick
	CLUSFilter
	CLUSSession
	CLUSSessionArray
	CLUSSessionCounter
	CLUSDatapathCounter
	CLUSDerivedPolicyApp
	CLUSDerivedPolicyRule
	CLUSDerivedPolicyRuleArray
	CLUSDerivedPolicyRuleMap
	CLUSProbeSummary
	CLUSProbeProcess
	CLUSProbeProcessArray
	CLUSProbeContainer
	CLUSProbeContainerArray
	CLUSFileMonitorFile
	CLUSFileMonitorFileArray
	CLUSSnifferRequest
	CLUSSnifferResponse
	CLUSSnifferFilter
	CLUSSniffer
	CLUSSnifferArray
	CLUSSnifferDownload
	CLUSSnifferPcap
	CLUSContainerLogReq
	CLUSContainerLogRes
	CLUSProcess
	CLUSProcessArray
	CLUSDerivedDlpRule
	CLUSDerivedDlpRuleArray
	CLUSDerivedDlpRuleMap
	CLUSDerivedDlpRuleEntry
	CLUSDerivedDlpRuleEntryArray
	CLUSDerivedDlpRuleMac
	CLUSDerivedDlpRuleMacArray
	CLUSDerivedProcessRule
	CLUSDerivedProcessRuleArray
	CLUSDerivedFileRule
	CLUSDerivedFileRuleArray
	CLUSWorkloadInterceptPort
	CLUSWorkloadIntercept
	CLUSMeter
	CLUSMeterArray
	CLUSWlIDArray
	ScanTypeMap
	ScanVulnerability
	ScanLayerResult
	ScanModule
	ScanModuleVul
	ScanSecretLog
	ScanSecretResult
	ScanSetIdPermLog
	ScanResult
	ScanSignatureInfo
	ScanRunningRequest
	ScanData
	ScanAppPackage
	ScanAppRequest
	ScanAwsLambdaRequest
	ScannerSettings
	ScanImageRequest
	SigstoreRootOfTrust
	SigstoreVerifier
	SigstoreKeypairOptions
	SigstoreKeylessOptions
	ScanCacheStatRes
	ScanCacheDataRes
*/
package share

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProfilingCmd int32

const (
	ProfilingCmd_StartProfiling ProfilingCmd = 0
	ProfilingCmd_StopProfiling  ProfilingCmd = 1
)

var ProfilingCmd_name = map[int32]string{
	0: "StartProfiling",
	1: "StopProfiling",
}
var ProfilingCmd_value = map[string]int32{
	"StartProfiling": 0,
	"StopProfiling":  1,
}

func (x ProfilingCmd) String() string {
	return proto.EnumName(ProfilingCmd_name, int32(x))
}
func (ProfilingCmd) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProfilingMethod int32

const (
	ProfilingMethod_CPU    ProfilingMethod = 0
	ProfilingMethod_Memory ProfilingMethod = 1
)

var ProfilingMethod_name = map[int32]string{
	0: "CPU",
	1: "Memory",
}
var ProfilingMethod_value = map[string]int32{
	"CPU":    0,
	"Memory": 1,
}

func (x ProfilingMethod) String() string {
	return proto.EnumName(ProfilingMethod_name, int32(x))
}
func (ProfilingMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type RPCVoid struct {
}

func (m *RPCVoid) Reset()                    { *m = RPCVoid{} }
func (m *RPCVoid) String() string            { return proto.CompactTextString(m) }
func (*RPCVoid) ProtoMessage()               {}
func (*RPCVoid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CLUSProfilingRequest struct {
	Cmd      ProfilingCmd      `protobuf:"varint,1,opt,name=Cmd,enum=share.ProfilingCmd" json:"Cmd,omitempty"`
	Methods  []ProfilingMethod `protobuf:"varint,2,rep,packed,name=Methods,enum=share.ProfilingMethod" json:"Methods,omitempty"`
	Duration uint32            `protobuf:"varint,3,opt,name=Duration" json:"Duration,omitempty"`
}

func (m *CLUSProfilingRequest) Reset()                    { *m = CLUSProfilingRequest{} }
func (m *CLUSProfilingRequest) String() string            { return proto.CompactTextString(m) }
func (*CLUSProfilingRequest) ProtoMessage()               {}
func (*CLUSProfilingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CLUSProfilingRequest) GetCmd() ProfilingCmd {
	if m != nil {
		return m.Cmd
	}
	return ProfilingCmd_StartProfiling
}

func (m *CLUSProfilingRequest) GetMethods() []ProfilingMethod {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *CLUSProfilingRequest) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type CLUSBoolean struct {
	Value bool `protobuf:"varint,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *CLUSBoolean) Reset()                    { *m = CLUSBoolean{} }
func (m *CLUSBoolean) String() string            { return proto.CompactTextString(m) }
func (*CLUSBoolean) ProtoMessage()               {}
func (*CLUSBoolean) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CLUSBoolean) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type CLUSMetry struct {
	CPU           float64 `protobuf:"fixed64,1,opt,name=CPU" json:"CPU,omitempty"`
	Memory        uint64  `protobuf:"varint,2,opt,name=Memory" json:"Memory,omitempty"`
	SessionIn     uint32  `protobuf:"varint,3,opt,name=SessionIn" json:"SessionIn,omitempty"`
	SessionOut    uint32  `protobuf:"varint,4,opt,name=SessionOut" json:"SessionOut,omitempty"`
	SessionCurIn  uint32  `protobuf:"varint,5,opt,name=SessionCurIn" json:"SessionCurIn,omitempty"`
	SessionCurOut uint32  `protobuf:"varint,6,opt,name=SessionCurOut" json:"SessionCurOut,omitempty"`
	PacketIn      uint64  `protobuf:"varint,7,opt,name=PacketIn" json:"PacketIn,omitempty"`
	PacketOut     uint64  `protobuf:"varint,8,opt,name=PacketOut" json:"PacketOut,omitempty"`
	ByteIn        uint64  `protobuf:"varint,9,opt,name=ByteIn" json:"ByteIn,omitempty"`
	ByteOut       uint64  `protobuf:"varint,10,opt,name=ByteOut" json:"ByteOut,omitempty"`
}

func (m *CLUSMetry) Reset()                    { *m = CLUSMetry{} }
func (m *CLUSMetry) String() string            { return proto.CompactTextString(m) }
func (*CLUSMetry) ProtoMessage()               {}
func (*CLUSMetry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CLUSMetry) GetCPU() float64 {
	if m != nil {
		return m.CPU
	}
	return 0
}

func (m *CLUSMetry) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *CLUSMetry) GetSessionIn() uint32 {
	if m != nil {
		return m.SessionIn
	}
	return 0
}

func (m *CLUSMetry) GetSessionOut() uint32 {
	if m != nil {
		return m.SessionOut
	}
	return 0
}

func (m *CLUSMetry) GetSessionCurIn() uint32 {
	if m != nil {
		return m.SessionCurIn
	}
	return 0
}

func (m *CLUSMetry) GetSessionCurOut() uint32 {
	if m != nil {
		return m.SessionCurOut
	}
	return 0
}

func (m *CLUSMetry) GetPacketIn() uint64 {
	if m != nil {
		return m.PacketIn
	}
	return 0
}

func (m *CLUSMetry) GetPacketOut() uint64 {
	if m != nil {
		return m.PacketOut
	}
	return 0
}

func (m *CLUSMetry) GetByteIn() uint64 {
	if m != nil {
		return m.ByteIn
	}
	return 0
}

func (m *CLUSMetry) GetByteOut() uint64 {
	if m != nil {
		return m.ByteOut
	}
	return 0
}

type CLUSStats struct {
	ReadAt   int64      `protobuf:"varint,1,opt,name=ReadAt" json:"ReadAt,omitempty"`
	Interval uint32     `protobuf:"varint,2,opt,name=Interval" json:"Interval,omitempty"`
	Total    *CLUSMetry `protobuf:"bytes,3,opt,name=Total" json:"Total,omitempty"`
	Span1    *CLUSMetry `protobuf:"bytes,4,opt,name=Span1" json:"Span1,omitempty"`
	Span12   *CLUSMetry `protobuf:"bytes,5,opt,name=Span12" json:"Span12,omitempty"`
	Span60   *CLUSMetry `protobuf:"bytes,6,opt,name=Span60" json:"Span60,omitempty"`
}

func (m *CLUSStats) Reset()                    { *m = CLUSStats{} }
func (m *CLUSStats) String() string            { return proto.CompactTextString(m) }
func (*CLUSStats) ProtoMessage()               {}
func (*CLUSStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CLUSStats) GetReadAt() int64 {
	if m != nil {
		return m.ReadAt
	}
	return 0
}

func (m *CLUSStats) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *CLUSStats) GetTotal() *CLUSMetry {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *CLUSStats) GetSpan1() *CLUSMetry {
	if m != nil {
		return m.Span1
	}
	return nil
}

func (m *CLUSStats) GetSpan12() *CLUSMetry {
	if m != nil {
		return m.Span12
	}
	return nil
}

func (m *CLUSStats) GetSpan60() *CLUSMetry {
	if m != nil {
		return m.Span60
	}
	return nil
}

func init() {
	proto.RegisterType((*RPCVoid)(nil), "share.RPCVoid")
	proto.RegisterType((*CLUSProfilingRequest)(nil), "share.CLUSProfilingRequest")
	proto.RegisterType((*CLUSBoolean)(nil), "share.CLUSBoolean")
	proto.RegisterType((*CLUSMetry)(nil), "share.CLUSMetry")
	proto.RegisterType((*CLUSStats)(nil), "share.CLUSStats")
	proto.RegisterEnum("share.ProfilingCmd", ProfilingCmd_name, ProfilingCmd_value)
	proto.RegisterEnum("share.ProfilingMethod", ProfilingMethod_name, ProfilingMethod_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xeb, 0x26, 0x4e, 0xa6, 0x49, 0x30, 0x43, 0x55, 0x59, 0x08, 0xa1, 0xc8, 0x40, 0x65,
	0xf5, 0x10, 0x85, 0x20, 0x7a, 0xa7, 0xe6, 0x62, 0x89, 0x88, 0x68, 0x4d, 0x7b, 0x5f, 0xea, 0x85,
	0x5a, 0xd8, 0xbb, 0x61, 0xbd, 0x46, 0xca, 0x33, 0xf0, 0x24, 0x3c, 0x17, 0x2f, 0x82, 0x76, 0xfc,
	0x93, 0xb6, 0x2a, 0xb7, 0xf9, 0x7e, 0xc6, 0x9e, 0xfd, 0x76, 0x16, 0x26, 0x37, 0xaa, 0x2c, 0x95,
	0x5c, 0x6c, 0xb5, 0x32, 0x0a, 0x07, 0xd5, 0x2d, 0xd7, 0x22, 0x1c, 0x83, 0xc7, 0x36, 0xf1, 0xb5,
	0xca, 0xb3, 0xf0, 0xb7, 0x03, 0x27, 0xf1, 0xa7, 0xab, 0x74, 0xa3, 0xd5, 0xb7, 0xbc, 0xc8, 0xe5,
	0x77, 0x26, 0x7e, 0xd6, 0xa2, 0x32, 0xf8, 0x06, 0xdc, 0xb8, 0xcc, 0x02, 0x67, 0xee, 0x44, 0xb3,
	0xd5, 0xb3, 0x05, 0x35, 0x2e, 0x7a, 0x57, 0x5c, 0x66, 0xcc, 0xea, 0xb8, 0x04, 0x6f, 0x2d, 0xcc,
	0xad, 0xca, 0xaa, 0xe0, 0x70, 0xee, 0x46, 0xb3, 0xd5, 0xe9, 0x43, 0x6b, 0x23, 0xb3, 0xce, 0x86,
	0xcf, 0x61, 0xf4, 0xb1, 0xd6, 0xdc, 0xe4, 0x4a, 0x06, 0xee, 0xdc, 0x89, 0xa6, 0xac, 0xc7, 0xe1,
	0x2b, 0x38, 0xb6, 0xc3, 0x5c, 0x2a, 0x55, 0x08, 0x2e, 0xf1, 0x04, 0x06, 0xd7, 0xbc, 0xa8, 0x05,
	0x4d, 0x31, 0x62, 0x0d, 0x08, 0xff, 0x1c, 0xc2, 0xd8, 0xba, 0xd6, 0xc2, 0xe8, 0x1d, 0xfa, 0xe0,
	0xc6, 0x9b, 0x2b, 0x72, 0x38, 0xcc, 0x96, 0x78, 0x0a, 0xc3, 0xb5, 0x28, 0x95, 0xde, 0x05, 0x87,
	0x73, 0x27, 0x3a, 0x62, 0x2d, 0xc2, 0x17, 0x30, 0x4e, 0x45, 0x55, 0xe5, 0x4a, 0x26, 0xdd, 0x9f,
	0xf7, 0x04, 0xbe, 0x04, 0x68, 0xc1, 0xe7, 0xda, 0x04, 0x47, 0x24, 0xdf, 0x61, 0x30, 0x84, 0x49,
	0x8b, 0xe2, 0x5a, 0x27, 0x32, 0x18, 0x90, 0xe3, 0x1e, 0x87, 0xaf, 0x61, 0xba, 0xc7, 0xf6, 0x33,
	0x43, 0x32, 0xdd, 0x27, 0x6d, 0x00, 0x1b, 0x7e, 0xf3, 0x43, 0x98, 0x44, 0x06, 0x1e, 0x4d, 0xd8,
	0x63, 0x3b, 0x63, 0x53, 0xdb, 0xee, 0x11, 0x89, 0x7b, 0xc2, 0x9e, 0xec, 0x72, 0x67, 0x44, 0x22,
	0x83, 0x71, 0x73, 0xb2, 0x06, 0x61, 0x00, 0x9e, 0xad, 0x6c, 0x0f, 0x90, 0xd0, 0xc1, 0xf0, 0xaf,
	0xd3, 0x64, 0x95, 0x1a, 0x6e, 0x2a, 0xdb, 0xcf, 0x04, 0xcf, 0x3e, 0x18, 0x8a, 0xcb, 0x65, 0x2d,
	0xb2, 0x13, 0x25, 0xd2, 0x08, 0xfd, 0x8b, 0x17, 0x94, 0xd9, 0x94, 0xf5, 0x18, 0xcf, 0x60, 0xf0,
	0x45, 0x19, 0x5e, 0x50, 0x62, 0xc7, 0x2b, 0xbf, 0xbd, 0xde, 0xfe, 0x02, 0x58, 0x23, 0x5b, 0x5f,
	0xba, 0xe5, 0xf2, 0x2d, 0x45, 0xf7, 0xa8, 0x8f, 0x64, 0x8c, 0x60, 0x48, 0xc5, 0x8a, 0x12, 0x7c,
	0xcc, 0xd8, 0xea, 0x9d, 0xf3, 0x62, 0x49, 0x31, 0xfe, 0xd7, 0x79, 0xb1, 0x3c, 0x7f, 0x0f, 0x93,
	0xbb, 0x9b, 0x89, 0x08, 0xb3, 0xd4, 0x70, 0x6d, 0x7a, 0xd2, 0x3f, 0xc0, 0xa7, 0x30, 0x4d, 0x8d,
	0xda, 0xee, 0x29, 0xe7, 0xfc, 0x0c, 0x9e, 0x3c, 0xd8, 0x52, 0xf4, 0x68, 0x9b, 0xfc, 0x03, 0x84,
	0x6e, 0x89, 0x7c, 0xe7, 0xeb, 0x90, 0x1e, 0xcf, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45,
	0xad, 0x79, 0xaa, 0x4c, 0x03, 0x00, 0x00,
}
