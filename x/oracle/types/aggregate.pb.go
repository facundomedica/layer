// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/oracle/aggregate.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Aggregate struct {
	QueryId           string               `protobuf:"bytes,1,opt,name=queryId,proto3" json:"queryId,omitempty"`
	AggregateValue    string               `protobuf:"bytes,2,opt,name=aggregateValue,proto3" json:"aggregateValue,omitempty"`
	AggregateReporter string               `protobuf:"bytes,3,opt,name=aggregateReporter,proto3" json:"aggregateReporter,omitempty"`
	ReporterPower     int64                `protobuf:"varint,4,opt,name=reporterPower,proto3" json:"reporterPower,omitempty"`
	StandardDeviation float64              `protobuf:"fixed64,5,opt,name=standardDeviation,proto3" json:"standardDeviation,omitempty"`
	Reporters         []*AggregateReporter `protobuf:"bytes,6,rep,name=reporters,proto3" json:"reporters,omitempty"`
	Flagged           bool                 `protobuf:"varint,7,opt,name=flagged,proto3" json:"flagged,omitempty"`
	Nonce             int64                `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *Aggregate) Reset()         { *m = Aggregate{} }
func (m *Aggregate) String() string { return proto.CompactTextString(m) }
func (*Aggregate) ProtoMessage()    {}
func (*Aggregate) Descriptor() ([]byte, []int) {
	return fileDescriptor_788ad347f505f8a6, []int{0}
}
func (m *Aggregate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Aggregate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Aggregate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Aggregate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aggregate.Merge(m, src)
}
func (m *Aggregate) XXX_Size() int {
	return m.Size()
}
func (m *Aggregate) XXX_DiscardUnknown() {
	xxx_messageInfo_Aggregate.DiscardUnknown(m)
}

var xxx_messageInfo_Aggregate proto.InternalMessageInfo

func (m *Aggregate) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Aggregate) GetAggregateValue() string {
	if m != nil {
		return m.AggregateValue
	}
	return ""
}

func (m *Aggregate) GetAggregateReporter() string {
	if m != nil {
		return m.AggregateReporter
	}
	return ""
}

func (m *Aggregate) GetReporterPower() int64 {
	if m != nil {
		return m.ReporterPower
	}
	return 0
}

func (m *Aggregate) GetStandardDeviation() float64 {
	if m != nil {
		return m.StandardDeviation
	}
	return 0
}

func (m *Aggregate) GetReporters() []*AggregateReporter {
	if m != nil {
		return m.Reporters
	}
	return nil
}

func (m *Aggregate) GetFlagged() bool {
	if m != nil {
		return m.Flagged
	}
	return false
}

func (m *Aggregate) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type AvailableTimestamps struct {
	Timestamps []time.Time `protobuf:"bytes,1,rep,name=timestamps,proto3,stdtime" json:"timestamps"`
}

func (m *AvailableTimestamps) Reset()         { *m = AvailableTimestamps{} }
func (m *AvailableTimestamps) String() string { return proto.CompactTextString(m) }
func (*AvailableTimestamps) ProtoMessage()    {}
func (*AvailableTimestamps) Descriptor() ([]byte, []int) {
	return fileDescriptor_788ad347f505f8a6, []int{1}
}
func (m *AvailableTimestamps) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AvailableTimestamps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AvailableTimestamps.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AvailableTimestamps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableTimestamps.Merge(m, src)
}
func (m *AvailableTimestamps) XXX_Size() int {
	return m.Size()
}
func (m *AvailableTimestamps) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableTimestamps.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableTimestamps proto.InternalMessageInfo

func (m *AvailableTimestamps) GetTimestamps() []time.Time {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func init() {
	proto.RegisterType((*Aggregate)(nil), "layer.oracle.Aggregate")
	proto.RegisterType((*AvailableTimestamps)(nil), "layer.oracle.AvailableTimestamps")
}

func init() { proto.RegisterFile("layer/oracle/aggregate.proto", fileDescriptor_788ad347f505f8a6) }

var fileDescriptor_788ad347f505f8a6 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xbd, 0xce, 0xd3, 0x30,
	0x14, 0x8d, 0xbf, 0xf2, 0x7d, 0x6d, 0xcd, 0x8f, 0x84, 0xe9, 0x60, 0x55, 0x28, 0x89, 0x2a, 0x40,
	0x41, 0x02, 0x47, 0x82, 0x99, 0xa1, 0xa5, 0x0b, 0x1b, 0x8a, 0x10, 0x03, 0x0c, 0xc8, 0x69, 0x6e,
	0x4d, 0x24, 0x37, 0x0e, 0x8e, 0x53, 0xe8, 0x5b, 0xf4, 0x89, 0x98, 0x3b, 0x76, 0x64, 0x02, 0xd4,
	0xbe, 0x08, 0xaa, 0x53, 0x87, 0xd2, 0xb2, 0xf9, 0xdc, 0x73, 0x8e, 0xef, 0xd1, 0x3d, 0xf8, 0xa1,
	0xe4, 0x2b, 0xd0, 0xb1, 0xd2, 0x7c, 0x26, 0x21, 0xe6, 0x42, 0x68, 0x10, 0xdc, 0x00, 0x2b, 0xb5,
	0x32, 0x8a, 0xdc, 0xb1, 0x2c, 0x6b, 0xd8, 0xe1, 0x40, 0x28, 0xa1, 0x2c, 0x11, 0x1f, 0x5e, 0x8d,
	0x66, 0x18, 0x08, 0xa5, 0x84, 0x84, 0xd8, 0xa2, 0xb4, 0x9e, 0xc7, 0x26, 0x5f, 0x40, 0x65, 0xf8,
	0xa2, 0x3c, 0x0a, 0x1e, 0xff, 0x7f, 0xc5, 0x27, 0x0d, 0xa5, 0xd2, 0x06, 0x74, 0x23, 0x1b, 0x7d,
	0xbf, 0xc2, 0xfd, 0xb1, 0x23, 0x09, 0xc5, 0xdd, 0x2f, 0x35, 0xe8, 0xd5, 0x9b, 0x8c, 0xa2, 0x10,
	0x45, 0xfd, 0xc4, 0x41, 0xf2, 0x04, 0xdf, 0x6b, 0xff, 0x78, 0xcf, 0x65, 0x0d, 0xf4, 0xca, 0x0a,
	0xce, 0xa6, 0xe4, 0x19, 0xbe, 0xdf, 0x4e, 0x92, 0xe3, 0x2a, 0xda, 0xb1, 0xd2, 0x4b, 0x82, 0x3c,
	0xc2, 0x77, 0x5d, 0x9e, 0xb7, 0xea, 0x2b, 0x68, 0x7a, 0x2b, 0x44, 0x51, 0x27, 0xf9, 0x77, 0x78,
	0xf8, 0xb3, 0x32, 0xbc, 0xc8, 0xb8, 0xce, 0xa6, 0xb0, 0xcc, 0xb9, 0xc9, 0x55, 0x41, 0xaf, 0x43,
	0x14, 0xa1, 0xe4, 0x92, 0x20, 0xaf, 0x70, 0xdf, 0xd9, 0x2b, 0x7a, 0x13, 0x76, 0xa2, 0xdb, 0x2f,
	0x02, 0x76, 0x7a, 0x51, 0x36, 0x3e, 0xcf, 0x91, 0xfc, 0x75, 0x1c, 0x4e, 0x30, 0x97, 0x5c, 0x08,
	0xc8, 0x68, 0x37, 0x44, 0x51, 0x2f, 0x71, 0x90, 0x0c, 0xf0, 0x75, 0xa1, 0x8a, 0x19, 0xd0, 0x9e,
	0x0d, 0xd9, 0x80, 0xd1, 0x47, 0xfc, 0x60, 0xbc, 0xe4, 0xb9, 0xe4, 0xa9, 0x84, 0x77, 0xae, 0x83,
	0x8a, 0x4c, 0x31, 0x6e, 0x1b, 0xa9, 0x28, 0xb2, 0x31, 0x86, 0xac, 0x29, 0x8d, 0xb9, 0xd2, 0x58,
	0x6b, 0x98, 0xf4, 0x36, 0x3f, 0x03, 0x6f, 0xfd, 0x2b, 0x40, 0xc9, 0x89, 0x6f, 0xf2, 0x7a, 0xb3,
	0xf3, 0xd1, 0x76, 0xe7, 0xa3, 0xdf, 0x3b, 0x1f, 0xad, 0xf7, 0xbe, 0xb7, 0xdd, 0xfb, 0xde, 0x8f,
	0xbd, 0xef, 0x7d, 0x78, 0x2a, 0x72, 0xf3, 0xb9, 0x4e, 0xd9, 0x4c, 0x2d, 0x62, 0x03, 0x52, 0x2a,
	0xfd, 0x3c, 0x57, 0x71, 0xd3, 0xf9, 0x37, 0xd7, 0xba, 0x59, 0x95, 0x50, 0xa5, 0x37, 0x76, 0xdd,
	0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x0f, 0xda, 0xd1, 0x75, 0x02, 0x00, 0x00,
}

func (m *Aggregate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Aggregate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Aggregate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintAggregate(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x40
	}
	if m.Flagged {
		i--
		if m.Flagged {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Reporters) > 0 {
		for iNdEx := len(m.Reporters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reporters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAggregate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.StandardDeviation != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.StandardDeviation))))
		i--
		dAtA[i] = 0x29
	}
	if m.ReporterPower != 0 {
		i = encodeVarintAggregate(dAtA, i, uint64(m.ReporterPower))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AggregateReporter) > 0 {
		i -= len(m.AggregateReporter)
		copy(dAtA[i:], m.AggregateReporter)
		i = encodeVarintAggregate(dAtA, i, uint64(len(m.AggregateReporter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AggregateValue) > 0 {
		i -= len(m.AggregateValue)
		copy(dAtA[i:], m.AggregateValue)
		i = encodeVarintAggregate(dAtA, i, uint64(len(m.AggregateValue)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintAggregate(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AvailableTimestamps) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AvailableTimestamps) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AvailableTimestamps) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Timestamps) > 0 {
		for iNdEx := len(m.Timestamps) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamps[iNdEx], dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamps[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintAggregate(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAggregate(dAtA []byte, offset int, v uint64) int {
	offset -= sovAggregate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Aggregate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovAggregate(uint64(l))
	}
	l = len(m.AggregateValue)
	if l > 0 {
		n += 1 + l + sovAggregate(uint64(l))
	}
	l = len(m.AggregateReporter)
	if l > 0 {
		n += 1 + l + sovAggregate(uint64(l))
	}
	if m.ReporterPower != 0 {
		n += 1 + sovAggregate(uint64(m.ReporterPower))
	}
	if m.StandardDeviation != 0 {
		n += 9
	}
	if len(m.Reporters) > 0 {
		for _, e := range m.Reporters {
			l = e.Size()
			n += 1 + l + sovAggregate(uint64(l))
		}
	}
	if m.Flagged {
		n += 2
	}
	if m.Nonce != 0 {
		n += 1 + sovAggregate(uint64(m.Nonce))
	}
	return n
}

func (m *AvailableTimestamps) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Timestamps) > 0 {
		for _, e := range m.Timestamps {
			l = github_com_cosmos_gogoproto_types.SizeOfStdTime(e)
			n += 1 + l + sovAggregate(uint64(l))
		}
	}
	return n
}

func sovAggregate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAggregate(x uint64) (n int) {
	return sovAggregate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Aggregate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAggregate
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
			return fmt.Errorf("proto: Aggregate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Aggregate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateReporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateReporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReporterPower", wireType)
			}
			m.ReporterPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReporterPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardDeviation", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.StandardDeviation = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporters = append(m.Reporters, &AggregateReporter{})
			if err := m.Reporters[len(m.Reporters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flagged", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Flagged = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAggregate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAggregate
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
func (m *AvailableTimestamps) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAggregate
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
			return fmt.Errorf("proto: AvailableTimestamps: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AvailableTimestamps: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamps = append(m.Timestamps, time.Time{})
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&(m.Timestamps[len(m.Timestamps)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAggregate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAggregate
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
func skipAggregate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAggregate
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
					return 0, ErrIntOverflowAggregate
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
					return 0, ErrIntOverflowAggregate
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
				return 0, ErrInvalidLengthAggregate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAggregate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAggregate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAggregate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAggregate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAggregate = fmt.Errorf("proto: unexpected end of group")
)
