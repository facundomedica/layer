// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/oracle/micro_report.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MicroReport struct {
	Reporter  string `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	Power     int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	Qid       string `protobuf:"bytes,3,opt,name=qid,proto3" json:"qid,omitempty"`
	Value     string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp uint64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *MicroReport) Reset()         { *m = MicroReport{} }
func (m *MicroReport) String() string { return proto.CompactTextString(m) }
func (*MicroReport) ProtoMessage()    {}
func (*MicroReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39350954f878191, []int{0}
}
func (m *MicroReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MicroReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MicroReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MicroReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicroReport.Merge(m, src)
}
func (m *MicroReport) XXX_Size() int {
	return m.Size()
}
func (m *MicroReport) XXX_DiscardUnknown() {
	xxx_messageInfo_MicroReport.DiscardUnknown(m)
}

var xxx_messageInfo_MicroReport proto.InternalMessageInfo

func (m *MicroReport) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *MicroReport) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *MicroReport) GetQid() string {
	if m != nil {
		return m.Qid
	}
	return ""
}

func (m *MicroReport) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *MicroReport) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*MicroReport)(nil), "layer.oracle.MicroReport")
}

func init() { proto.RegisterFile("layer/oracle/micro_report.proto", fileDescriptor_c39350954f878191) }

var fileDescriptor_c39350954f878191 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0x8f, 0x2f,
	0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x2b, 0xd0,
	0x83, 0x28, 0x50, 0x6a, 0x65, 0xe4, 0xe2, 0xf6, 0x05, 0x29, 0x0a, 0x02, 0xab, 0x11, 0x92, 0xe2,
	0xe2, 0x80, 0xa8, 0x4e, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x44,
	0xb8, 0x58, 0x0b, 0xf2, 0xcb, 0x53, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c,
	0x21, 0x01, 0x2e, 0xe6, 0xc2, 0xcc, 0x14, 0x09, 0x66, 0xb0, 0x62, 0x10, 0x13, 0xa4, 0xae, 0x2c,
	0x31, 0xa7, 0x34, 0x55, 0x82, 0x05, 0x2c, 0x06, 0xe1, 0x08, 0xc9, 0x70, 0x71, 0x96, 0x64, 0xe6,
	0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x48, 0xb0, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x21, 0x04, 0x9c,
	0x9c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x33, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x24, 0x35, 0x27, 0x27, 0xbf, 0x48, 0x37, 0x33,
	0x5f, 0x1f, 0xe2, 0xcb, 0x0a, 0x98, 0x3f, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x3e,
	0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xa4, 0x02, 0x3e, 0x04, 0x01, 0x00, 0x00,
}

func (m *MicroReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MicroReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MicroReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintMicroReport(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Qid) > 0 {
		i -= len(m.Qid)
		copy(dAtA[i:], m.Qid)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.Qid)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Power != 0 {
		i = encodeVarintMicroReport(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMicroReport(dAtA []byte, offset int, v uint64) int {
	offset -= sovMicroReport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MicroReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovMicroReport(uint64(m.Power))
	}
	l = len(m.Qid)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovMicroReport(uint64(m.Timestamp))
	}
	return n
}

func sovMicroReport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMicroReport(x uint64) (n int) {
	return sovMicroReport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MicroReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMicroReport
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
			return fmt.Errorf("proto: MicroReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MicroReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Qid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMicroReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMicroReport
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
func skipMicroReport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMicroReport
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
					return 0, ErrIntOverflowMicroReport
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
					return 0, ErrIntOverflowMicroReport
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
				return 0, ErrInvalidLengthMicroReport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMicroReport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMicroReport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMicroReport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMicroReport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMicroReport = fmt.Errorf("proto: unexpected end of group")
)
