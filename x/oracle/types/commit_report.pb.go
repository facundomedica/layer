// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/oracle/commit_report.proto

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

type Commit struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	QueryId   []byte `protobuf:"bytes,2,opt,name=queryId,proto3" json:"queryId,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56672e00badb6a1, []int{0}
}
func (m *Commit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return m.Size()
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Commit) GetQueryId() []byte {
	if m != nil {
		return m.QueryId
	}
	return nil
}

func (m *Commit) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type CommitReport struct {
	Report *Commit `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	Block  int64   `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *CommitReport) Reset()         { *m = CommitReport{} }
func (m *CommitReport) String() string { return proto.CompactTextString(m) }
func (*CommitReport) ProtoMessage()    {}
func (*CommitReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56672e00badb6a1, []int{1}
}
func (m *CommitReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitReport.Merge(m, src)
}
func (m *CommitReport) XXX_Size() int {
	return m.Size()
}
func (m *CommitReport) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitReport.DiscardUnknown(m)
}

var xxx_messageInfo_CommitReport proto.InternalMessageInfo

func (m *CommitReport) GetReport() *Commit {
	if m != nil {
		return m.Report
	}
	return nil
}

func (m *CommitReport) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

type CommitsByHeight struct {
	Commits []*Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
}

func (m *CommitsByHeight) Reset()         { *m = CommitsByHeight{} }
func (m *CommitsByHeight) String() string { return proto.CompactTextString(m) }
func (*CommitsByHeight) ProtoMessage()    {}
func (*CommitsByHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56672e00badb6a1, []int{2}
}
func (m *CommitsByHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitsByHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitsByHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitsByHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitsByHeight.Merge(m, src)
}
func (m *CommitsByHeight) XXX_Size() int {
	return m.Size()
}
func (m *CommitsByHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitsByHeight.DiscardUnknown(m)
}

var xxx_messageInfo_CommitsByHeight proto.InternalMessageInfo

func (m *CommitsByHeight) GetCommits() []*Commit {
	if m != nil {
		return m.Commits
	}
	return nil
}

func init() {
	proto.RegisterType((*Commit)(nil), "layer.oracle.Commit")
	proto.RegisterType((*CommitReport)(nil), "layer.oracle.CommitReport")
	proto.RegisterType((*CommitsByHeight)(nil), "layer.oracle.CommitsByHeight")
}

func init() { proto.RegisterFile("layer/oracle/commit_report.proto", fileDescriptor_b56672e00badb6a1) }

var fileDescriptor_b56672e00badb6a1 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0x9b, 0xaf, 0x7c, 0x1d, 0x26, 0x16, 0x84, 0x30, 0x42, 0x11, 0x09, 0xa5, 0xab, 0x11,
	0x34, 0x05, 0x7d, 0x02, 0x67, 0x36, 0xba, 0xed, 0x72, 0x36, 0xd2, 0xd6, 0xd0, 0x09, 0xa6, 0xde,
	0x7a, 0x9b, 0xc2, 0xf4, 0x2d, 0x7c, 0x2c, 0x97, 0xb3, 0x74, 0x29, 0xed, 0x8b, 0x88, 0xc9, 0xd4,
	0x3f, 0x0b, 0x97, 0x27, 0xf9, 0xdd, 0x7b, 0xce, 0x3d, 0x34, 0xd6, 0x79, 0x2f, 0x31, 0x05, 0xcc,
	0x4b, 0x2d, 0xd3, 0x12, 0xea, 0x5a, 0x99, 0x7b, 0x94, 0x0d, 0xa0, 0x11, 0x0d, 0x82, 0x01, 0x16,
	0x5a, 0x42, 0x38, 0xe2, 0xf4, 0xe4, 0x17, 0x6f, 0x76, 0x0e, 0x4a, 0x36, 0x34, 0x58, 0xdb, 0x59,
	0x16, 0xd1, 0x59, 0x89, 0x32, 0x37, 0x80, 0x11, 0x89, 0xc9, 0x72, 0x9e, 0x4d, 0xf2, 0xf3, 0xe7,
	0xb9, 0x93, 0xd8, 0xdf, 0x3d, 0x44, 0xff, 0x62, 0xb2, 0x0c, 0xb3, 0x49, 0xb2, 0x33, 0x3a, 0x6f,
	0x55, 0xf5, 0x94, 0x9b, 0x0e, 0x65, 0xe4, 0xdb, 0xa9, 0xef, 0x87, 0x24, 0xa3, 0xa1, 0xdb, 0x9d,
	0xd9, 0x58, 0xec, 0x82, 0x06, 0x2e, 0xa0, 0x35, 0x38, 0xba, 0x5a, 0x88, 0x9f, 0x09, 0xc5, 0x81,
	0x3d, 0x30, 0x6c, 0x41, 0xff, 0x17, 0x1a, 0xca, 0x47, 0xeb, 0xe9, 0x67, 0x4e, 0x24, 0x37, 0xf4,
	0xd8, 0x71, 0xed, 0xaa, 0xbf, 0x95, 0xaa, 0xda, 0x1a, 0x26, 0xe8, 0xcc, 0x9d, 0xdf, 0x46, 0x24,
	0xf6, 0xff, 0xdc, 0x3b, 0x41, 0xab, 0xf5, 0xeb, 0xc0, 0xc9, 0x7e, 0xe0, 0xe4, 0x7d, 0xe0, 0xe4,
	0x65, 0xe4, 0xde, 0x7e, 0xe4, 0xde, 0xdb, 0xc8, 0xbd, 0xcd, 0x79, 0xa5, 0xcc, 0xb6, 0x2b, 0x44,
	0x09, 0x75, 0x6a, 0xa4, 0xd6, 0x80, 0x97, 0x0a, 0x52, 0x57, 0xdc, 0xee, 0xab, 0xba, 0xbe, 0x91,
	0x6d, 0x11, 0xd8, 0xfa, 0xae, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x47, 0xeb, 0x6b, 0x26, 0x87,
	0x01, 0x00, 0x00,
}

func (m *Commit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintCommitReport(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintCommitReport(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCommitReport(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommitReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintCommitReport(dAtA, i, uint64(m.Block))
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
			i = encodeVarintCommitReport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommitsByHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitsByHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitsByHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commits) > 0 {
		for iNdEx := len(m.Commits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommitReport(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommitReport(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommitReport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Commit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCommitReport(uint64(l))
	}
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovCommitReport(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovCommitReport(uint64(l))
	}
	return n
}

func (m *CommitReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Report != nil {
		l = m.Report.Size()
		n += 1 + l + sovCommitReport(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovCommitReport(uint64(m.Block))
	}
	return n
}

func (m *CommitsByHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Commits) > 0 {
		for _, e := range m.Commits {
			l = e.Size()
			n += 1 + l + sovCommitReport(uint64(l))
		}
	}
	return n
}

func sovCommitReport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommitReport(x uint64) (n int) {
	return sovCommitReport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Commit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommitReport
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
			return fmt.Errorf("proto: Commit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitReport
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
				return ErrInvalidLengthCommitReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommitReport
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = append(m.QueryId[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryId == nil {
				m.QueryId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitReport
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
				return ErrInvalidLengthCommitReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommitReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommitReport
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
func (m *CommitReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommitReport
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
			return fmt.Errorf("proto: CommitReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Report", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitReport
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
				return ErrInvalidLengthCommitReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommitReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Report == nil {
				m.Report = &Commit{}
			}
			if err := m.Report.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommitReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommitReport
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
func (m *CommitsByHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommitReport
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
			return fmt.Errorf("proto: CommitsByHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitsByHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitReport
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
				return ErrInvalidLengthCommitReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommitReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commits = append(m.Commits, &Commit{})
			if err := m.Commits[len(m.Commits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommitReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommitReport
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
func skipCommitReport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommitReport
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
					return 0, ErrIntOverflowCommitReport
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
					return 0, ErrIntOverflowCommitReport
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
				return 0, ErrInvalidLengthCommitReport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommitReport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommitReport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommitReport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommitReport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommitReport = fmt.Errorf("proto: unexpected end of group")
)