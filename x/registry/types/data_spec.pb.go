// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/registry/data_spec.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DataSpec is a specification for how to interpret and aggregate data
type DataSpec struct {
	// ipfs hash of the data spec
	DocumentHash string `protobuf:"bytes,1,opt,name=document_hash,json=documentHash,proto3" json:"document_hash,omitempty"`
	// the value's datatype for decoding the value
	ValueType string `protobuf:"bytes,2,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	// how to aggregate the data (ie. average, median, mode, etc) for aggregating reports and arriving at final value
	AggregationMethod string `protobuf:"bytes,3,opt,name=aggregation_method,json=aggregationMethod,proto3" json:"aggregation_method,omitempty"`
	// address that originally registered the data spec
	Registrar string `protobuf:"bytes,4,opt,name=registrar,proto3" json:"registrar,omitempty"`
	// report_buffer_window specifies the duration of the time window following an initial report
	// during which additional reports can be submitted. This duration acts as a buffer, allowing
	// a collection of related reports in a defined time frame. The window ensures that all
	// pertinent reports are aggregated together before arriving at a final value. This defaults
	// to 0s if not specified.
	// extensions: treat as a golang time.duration, don't allow nil values, don't omit empty values
	ReportBufferWindow time.Duration `protobuf:"bytes,5,opt,name=report_buffer_window,json=reportBufferWindow,proto3,stdduration" json:"report_buffer_window"`
}

func (m *DataSpec) Reset()         { *m = DataSpec{} }
func (m *DataSpec) String() string { return proto.CompactTextString(m) }
func (*DataSpec) ProtoMessage()    {}
func (*DataSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c1d9edbb99f1378, []int{0}
}
func (m *DataSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSpec.Merge(m, src)
}
func (m *DataSpec) XXX_Size() int {
	return m.Size()
}
func (m *DataSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DataSpec proto.InternalMessageInfo

func (m *DataSpec) GetDocumentHash() string {
	if m != nil {
		return m.DocumentHash
	}
	return ""
}

func (m *DataSpec) GetValueType() string {
	if m != nil {
		return m.ValueType
	}
	return ""
}

func (m *DataSpec) GetAggregationMethod() string {
	if m != nil {
		return m.AggregationMethod
	}
	return ""
}

func (m *DataSpec) GetRegistrar() string {
	if m != nil {
		return m.Registrar
	}
	return ""
}

func (m *DataSpec) GetReportBufferWindow() time.Duration {
	if m != nil {
		return m.ReportBufferWindow
	}
	return 0
}

func init() {
	proto.RegisterType((*DataSpec)(nil), "layer.registry.DataSpec")
}

func init() { proto.RegisterFile("layer/registry/data_spec.proto", fileDescriptor_8c1d9edbb99f1378) }

var fileDescriptor_8c1d9edbb99f1378 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbd, 0x8a, 0xdb, 0x40,
	0x14, 0x85, 0x35, 0xce, 0x0f, 0xf1, 0x24, 0x0e, 0x78, 0x70, 0x21, 0x1b, 0x32, 0x36, 0x49, 0x63,
	0x12, 0xac, 0x81, 0x04, 0xd2, 0x47, 0x38, 0x90, 0x26, 0x8d, 0x1d, 0x08, 0xb8, 0x11, 0x23, 0x69,
	0x3c, 0x12, 0x48, 0xba, 0x62, 0x66, 0x14, 0x47, 0x6f, 0x91, 0x74, 0x79, 0x84, 0x94, 0x5b, 0xec,
	0x43, 0xb8, 0x34, 0x5b, 0x6d, 0xb5, 0xbb, 0xd8, 0xc5, 0xbe, 0xc6, 0xe2, 0x91, 0x8c, 0xb7, 0x11,
	0xba, 0xe7, 0x3b, 0x9c, 0xcb, 0x99, 0x8b, 0x69, 0xc6, 0x6b, 0xa1, 0x98, 0x12, 0x32, 0xd5, 0x46,
	0xd5, 0x2c, 0xe6, 0x86, 0x07, 0xba, 0x14, 0x91, 0x57, 0x2a, 0x30, 0x40, 0x5e, 0x5b, 0xee, 0x9d,
	0xf8, 0xa8, 0xcf, 0xf3, 0xb4, 0x00, 0x66, 0xbf, 0x8d, 0x65, 0x34, 0x8c, 0x40, 0xe7, 0xa0, 0x03,
	0x3b, 0xb1, 0x66, 0x68, 0xd1, 0x40, 0x82, 0x84, 0x46, 0x3f, 0xfe, 0xb5, 0x2a, 0x95, 0x00, 0x32,
	0x13, 0xcc, 0x4e, 0x61, 0xb5, 0x66, 0x71, 0xa5, 0xb8, 0x49, 0xa1, 0x68, 0xf8, 0xdb, 0xbf, 0x1d,
	0xfc, 0x62, 0xce, 0x0d, 0x5f, 0x96, 0x22, 0x22, 0xef, 0x70, 0x2f, 0x86, 0xa8, 0xca, 0x45, 0x61,
	0x82, 0x84, 0xeb, 0xc4, 0x45, 0x13, 0x34, 0xed, 0x2e, 0x5e, 0x9d, 0xc4, 0x6f, 0x5c, 0x27, 0xe4,
	0x0d, 0xc6, 0xbf, 0x78, 0x56, 0x89, 0xc0, 0xd4, 0xa5, 0x70, 0x3b, 0xd6, 0xd1, 0xb5, 0xca, 0x8f,
	0xba, 0x14, 0x64, 0x86, 0x09, 0x97, 0x52, 0x09, 0x69, 0xb7, 0x04, 0xb9, 0x30, 0x09, 0xc4, 0xee,
	0x13, 0x6b, 0xeb, 0x3f, 0x22, 0xdf, 0x2d, 0x20, 0x9f, 0x71, 0xb7, 0xed, 0xcb, 0x95, 0xfb, 0xf4,
	0xe8, 0xf2, 0xdd, 0xab, 0xcb, 0xd9, 0xa0, 0xad, 0xf6, 0x25, 0x8e, 0x95, 0xd0, 0x7a, 0x69, 0x54,
	0x5a, 0xc8, 0xc5, 0xd9, 0x4a, 0x56, 0x78, 0xa0, 0x44, 0x09, 0xca, 0x04, 0x61, 0xb5, 0x5e, 0x0b,
	0x15, 0x6c, 0xd2, 0x22, 0x86, 0x8d, 0xfb, 0x6c, 0x82, 0xa6, 0x2f, 0x3f, 0x0e, 0xbd, 0xa6, 0xb6,
	0x77, 0xaa, 0xed, 0xcd, 0xdb, 0xda, 0x7e, 0x6f, 0x7b, 0x33, 0x76, 0xfe, 0xdd, 0x8e, 0xd1, 0xff,
	0xfb, 0x8b, 0xf7, 0x68, 0x41, 0x9a, 0x14, 0xdf, 0x86, 0xfc, 0xb4, 0x19, 0xfe, 0xd7, 0xed, 0x9e,
	0xa2, 0xdd, 0x9e, 0xa2, 0xbb, 0x3d, 0x45, 0x7f, 0x0e, 0xd4, 0xd9, 0x1d, 0xa8, 0x73, 0x7d, 0xa0,
	0xce, 0xea, 0x83, 0x4c, 0x4d, 0x52, 0x85, 0x5e, 0x04, 0x39, 0x33, 0x22, 0xcb, 0x40, 0xcd, 0x52,
	0x60, 0xcd, 0x59, 0x7f, 0x9f, 0x0f, 0x7b, 0x7c, 0x19, 0x1d, 0x3e, 0xb7, 0xcb, 0x3f, 0x3d, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xb0, 0x3e, 0x27, 0x36, 0xf7, 0x01, 0x00, 0x00,
}

func (m *DataSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ReportBufferWindow, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ReportBufferWindow):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDataSpec(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Registrar) > 0 {
		i -= len(m.Registrar)
		copy(dAtA[i:], m.Registrar)
		i = encodeVarintDataSpec(dAtA, i, uint64(len(m.Registrar)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AggregationMethod) > 0 {
		i -= len(m.AggregationMethod)
		copy(dAtA[i:], m.AggregationMethod)
		i = encodeVarintDataSpec(dAtA, i, uint64(len(m.AggregationMethod)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValueType) > 0 {
		i -= len(m.ValueType)
		copy(dAtA[i:], m.ValueType)
		i = encodeVarintDataSpec(dAtA, i, uint64(len(m.ValueType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DocumentHash) > 0 {
		i -= len(m.DocumentHash)
		copy(dAtA[i:], m.DocumentHash)
		i = encodeVarintDataSpec(dAtA, i, uint64(len(m.DocumentHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataSpec(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataSpec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DocumentHash)
	if l > 0 {
		n += 1 + l + sovDataSpec(uint64(l))
	}
	l = len(m.ValueType)
	if l > 0 {
		n += 1 + l + sovDataSpec(uint64(l))
	}
	l = len(m.AggregationMethod)
	if l > 0 {
		n += 1 + l + sovDataSpec(uint64(l))
	}
	l = len(m.Registrar)
	if l > 0 {
		n += 1 + l + sovDataSpec(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ReportBufferWindow)
	n += 1 + l + sovDataSpec(uint64(l))
	return n
}

func sovDataSpec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataSpec(x uint64) (n int) {
	return sovDataSpec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataSpec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DocumentHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DocumentHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregationMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registrar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registrar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportBufferWindow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ReportBufferWindow, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataSpec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDataSpec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataSpec
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataSpec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDataSpec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataSpec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataSpec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataSpec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataSpec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataSpec = fmt.Errorf("proto: unexpected end of group")
)
