// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/dispute/dispute_params.proto

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

type DisputeParams struct {
	Report   *MicroReport    `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	Category DisputeCategory `protobuf:"varint,2,opt,name=category,proto3,enum=layer.dispute.DisputeCategory" json:"category,omitempty"`
}

func (m *DisputeParams) Reset()         { *m = DisputeParams{} }
func (m *DisputeParams) String() string { return proto.CompactTextString(m) }
func (*DisputeParams) ProtoMessage()    {}
func (*DisputeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b782315e5e42770, []int{0}
}
func (m *DisputeParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisputeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisputeParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisputeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputeParams.Merge(m, src)
}
func (m *DisputeParams) XXX_Size() int {
	return m.Size()
}
func (m *DisputeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputeParams.DiscardUnknown(m)
}

var xxx_messageInfo_DisputeParams proto.InternalMessageInfo

func (m *DisputeParams) GetReport() *MicroReport {
	if m != nil {
		return m.Report
	}
	return nil
}

func (m *DisputeParams) GetCategory() DisputeCategory {
	if m != nil {
		return m.Category
	}
	return Unspecified
}

func init() {
	proto.RegisterType((*DisputeParams)(nil), "layer.dispute.DisputeParams")
}

func init() {
	proto.RegisterFile("layer/dispute/dispute_params.proto", fileDescriptor_2b782315e5e42770)
}

var fileDescriptor_2b782315e5e42770 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0x4f, 0xc9, 0x2c, 0x2e, 0x28, 0x2d, 0x49, 0x85, 0xd1, 0xf1, 0x05, 0x89, 0x45, 0x89,
	0xb9, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x60, 0x35, 0x7a, 0x50, 0x39, 0x29,
	0x69, 0xac, 0x5a, 0x20, 0x6a, 0x95, 0xea, 0xb9, 0x78, 0x5d, 0x20, 0x02, 0x01, 0x60, 0x23, 0x84,
	0x8c, 0xb8, 0xd8, 0x8a, 0x52, 0x0b, 0xf2, 0x8b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d,
	0xa4, 0xf4, 0x50, 0x4c, 0xd3, 0xf3, 0xcd, 0x4c, 0x2e, 0xca, 0x0f, 0x02, 0xab, 0x08, 0x82, 0xaa,
	0x14, 0xb2, 0xe2, 0xe2, 0x48, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0x94, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0x33, 0x92, 0x43, 0xd3, 0x05, 0xb5, 0xc3, 0x19, 0xaa, 0x2a, 0x08, 0xae, 0xde, 0xc9,
	0xe5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x4b, 0x52, 0x73, 0x72, 0xf2, 0x8b, 0x74, 0x33, 0xf3,
	0xf5, 0x21, 0x9e, 0xa9, 0x80, 0x7b, 0xa7, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x1b,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xd8, 0xbc, 0x13, 0x1f, 0x01, 0x00, 0x00,
}

func (m *DisputeParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisputeParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisputeParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Category != 0 {
		i = encodeVarintDisputeParams(dAtA, i, uint64(m.Category))
		i--
		dAtA[i] = 0x10
	}
	if m.Report != nil {
		{
			size, err := m.Report.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDisputeParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDisputeParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovDisputeParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DisputeParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Report != nil {
		l = m.Report.Size()
		n += 1 + l + sovDisputeParams(uint64(l))
	}
	if m.Category != 0 {
		n += 1 + sovDisputeParams(uint64(m.Category))
	}
	return n
}

func sovDisputeParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDisputeParams(x uint64) (n int) {
	return sovDisputeParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DisputeParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDisputeParams
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
			return fmt.Errorf("proto: DisputeParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisputeParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Report", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisputeParams
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
				return ErrInvalidLengthDisputeParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDisputeParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Report == nil {
				m.Report = &MicroReport{}
			}
			if err := m.Report.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			m.Category = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisputeParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Category |= DisputeCategory(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDisputeParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDisputeParams
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
func skipDisputeParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDisputeParams
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
					return 0, ErrIntOverflowDisputeParams
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
					return 0, ErrIntOverflowDisputeParams
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
				return 0, ErrInvalidLengthDisputeParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDisputeParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDisputeParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDisputeParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDisputeParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDisputeParams = fmt.Errorf("proto: unexpected end of group")
)
