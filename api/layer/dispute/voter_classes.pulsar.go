// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package dispute

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_VoterClasses              protoreflect.MessageDescriptor
	fd_VoterClasses_validators   protoreflect.FieldDescriptor
	fd_VoterClasses_tokenHolders protoreflect.FieldDescriptor
	fd_VoterClasses_users        protoreflect.FieldDescriptor
	fd_VoterClasses_team         protoreflect.FieldDescriptor
)

func init() {
	file_layer_dispute_voter_classes_proto_init()
	md_VoterClasses = File_layer_dispute_voter_classes_proto.Messages().ByName("VoterClasses")
	fd_VoterClasses_validators = md_VoterClasses.Fields().ByName("validators")
	fd_VoterClasses_tokenHolders = md_VoterClasses.Fields().ByName("tokenHolders")
	fd_VoterClasses_users = md_VoterClasses.Fields().ByName("users")
	fd_VoterClasses_team = md_VoterClasses.Fields().ByName("team")
}

var _ protoreflect.Message = (*fastReflection_VoterClasses)(nil)

type fastReflection_VoterClasses VoterClasses

func (x *VoterClasses) ProtoReflect() protoreflect.Message {
	return (*fastReflection_VoterClasses)(x)
}

func (x *VoterClasses) slowProtoReflect() protoreflect.Message {
	mi := &file_layer_dispute_voter_classes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_VoterClasses_messageType fastReflection_VoterClasses_messageType
var _ protoreflect.MessageType = fastReflection_VoterClasses_messageType{}

type fastReflection_VoterClasses_messageType struct{}

func (x fastReflection_VoterClasses_messageType) Zero() protoreflect.Message {
	return (*fastReflection_VoterClasses)(nil)
}
func (x fastReflection_VoterClasses_messageType) New() protoreflect.Message {
	return new(fastReflection_VoterClasses)
}
func (x fastReflection_VoterClasses_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_VoterClasses
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_VoterClasses) Descriptor() protoreflect.MessageDescriptor {
	return md_VoterClasses
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_VoterClasses) Type() protoreflect.MessageType {
	return _fastReflection_VoterClasses_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_VoterClasses) New() protoreflect.Message {
	return new(fastReflection_VoterClasses)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_VoterClasses) Interface() protoreflect.ProtoMessage {
	return (*VoterClasses)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_VoterClasses) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Validators != "" {
		value := protoreflect.ValueOfString(x.Validators)
		if !f(fd_VoterClasses_validators, value) {
			return
		}
	}
	if x.TokenHolders != "" {
		value := protoreflect.ValueOfString(x.TokenHolders)
		if !f(fd_VoterClasses_tokenHolders, value) {
			return
		}
	}
	if x.Users != "" {
		value := protoreflect.ValueOfString(x.Users)
		if !f(fd_VoterClasses_users, value) {
			return
		}
	}
	if x.Team != "" {
		value := protoreflect.ValueOfString(x.Team)
		if !f(fd_VoterClasses_team, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_VoterClasses) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "layer.dispute.VoterClasses.validators":
		return x.Validators != ""
	case "layer.dispute.VoterClasses.tokenHolders":
		return x.TokenHolders != ""
	case "layer.dispute.VoterClasses.users":
		return x.Users != ""
	case "layer.dispute.VoterClasses.team":
		return x.Team != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.dispute.VoterClasses"))
		}
		panic(fmt.Errorf("message layer.dispute.VoterClasses does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VoterClasses) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "layer.dispute.VoterClasses.validators":
		x.Validators = ""
	case "layer.dispute.VoterClasses.tokenHolders":
		x.TokenHolders = ""
	case "layer.dispute.VoterClasses.users":
		x.Users = ""
	case "layer.dispute.VoterClasses.team":
		x.Team = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.dispute.VoterClasses"))
		}
		panic(fmt.Errorf("message layer.dispute.VoterClasses does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_VoterClasses) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "layer.dispute.VoterClasses.validators":
		value := x.Validators
		return protoreflect.ValueOfString(value)
	case "layer.dispute.VoterClasses.tokenHolders":
		value := x.TokenHolders
		return protoreflect.ValueOfString(value)
	case "layer.dispute.VoterClasses.users":
		value := x.Users
		return protoreflect.ValueOfString(value)
	case "layer.dispute.VoterClasses.team":
		value := x.Team
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.dispute.VoterClasses"))
		}
		panic(fmt.Errorf("message layer.dispute.VoterClasses does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VoterClasses) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "layer.dispute.VoterClasses.validators":
		x.Validators = value.Interface().(string)
	case "layer.dispute.VoterClasses.tokenHolders":
		x.TokenHolders = value.Interface().(string)
	case "layer.dispute.VoterClasses.users":
		x.Users = value.Interface().(string)
	case "layer.dispute.VoterClasses.team":
		x.Team = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.dispute.VoterClasses"))
		}
		panic(fmt.Errorf("message layer.dispute.VoterClasses does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VoterClasses) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "layer.dispute.VoterClasses.validators":
		panic(fmt.Errorf("field validators of message layer.dispute.VoterClasses is not mutable"))
	case "layer.dispute.VoterClasses.tokenHolders":
		panic(fmt.Errorf("field tokenHolders of message layer.dispute.VoterClasses is not mutable"))
	case "layer.dispute.VoterClasses.users":
		panic(fmt.Errorf("field users of message layer.dispute.VoterClasses is not mutable"))
	case "layer.dispute.VoterClasses.team":
		panic(fmt.Errorf("field team of message layer.dispute.VoterClasses is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.dispute.VoterClasses"))
		}
		panic(fmt.Errorf("message layer.dispute.VoterClasses does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_VoterClasses) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "layer.dispute.VoterClasses.validators":
		return protoreflect.ValueOfString("")
	case "layer.dispute.VoterClasses.tokenHolders":
		return protoreflect.ValueOfString("")
	case "layer.dispute.VoterClasses.users":
		return protoreflect.ValueOfString("")
	case "layer.dispute.VoterClasses.team":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.dispute.VoterClasses"))
		}
		panic(fmt.Errorf("message layer.dispute.VoterClasses does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_VoterClasses) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in layer.dispute.VoterClasses", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_VoterClasses) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VoterClasses) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_VoterClasses) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_VoterClasses) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*VoterClasses)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Validators)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TokenHolders)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Users)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Team)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*VoterClasses)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Team) > 0 {
			i -= len(x.Team)
			copy(dAtA[i:], x.Team)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Team)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Users) > 0 {
			i -= len(x.Users)
			copy(dAtA[i:], x.Users)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Users)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TokenHolders) > 0 {
			i -= len(x.TokenHolders)
			copy(dAtA[i:], x.TokenHolders)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TokenHolders)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Validators) > 0 {
			i -= len(x.Validators)
			copy(dAtA[i:], x.Validators)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Validators)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*VoterClasses)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VoterClasses: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VoterClasses: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TokenHolders", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TokenHolders = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Users = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Team", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Team = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: layer/dispute/voter_classes.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VoterClasses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators   string `protobuf:"bytes,1,opt,name=validators,proto3" json:"validators,omitempty"`
	TokenHolders string `protobuf:"bytes,2,opt,name=tokenHolders,proto3" json:"tokenHolders,omitempty"`
	Users        string `protobuf:"bytes,3,opt,name=users,proto3" json:"users,omitempty"`
	Team         string `protobuf:"bytes,4,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *VoterClasses) Reset() {
	*x = VoterClasses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_dispute_voter_classes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoterClasses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoterClasses) ProtoMessage() {}

// Deprecated: Use VoterClasses.ProtoReflect.Descriptor instead.
func (*VoterClasses) Descriptor() ([]byte, []int) {
	return file_layer_dispute_voter_classes_proto_rawDescGZIP(), []int{0}
}

func (x *VoterClasses) GetValidators() string {
	if x != nil {
		return x.Validators
	}
	return ""
}

func (x *VoterClasses) GetTokenHolders() string {
	if x != nil {
		return x.TokenHolders
	}
	return ""
}

func (x *VoterClasses) GetUsers() string {
	if x != nil {
		return x.Users
	}
	return ""
}

func (x *VoterClasses) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

var File_layer_dispute_voter_classes_proto protoreflect.FileDescriptor

var file_layer_dispute_voter_classes_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65, 0x2f,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x75,
	0x74, 0x65, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde,
	0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x60, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde,
	0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x52, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x50, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x42, 0x9b, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65,
	0x42, 0x11, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x64, 0x69,
	0x73, 0x70, 0x75, 0x74, 0x65, 0xa2, 0x02, 0x03, 0x4c, 0x44, 0x58, 0xaa, 0x02, 0x0d, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65, 0xca, 0x02, 0x0d, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x5c, 0x44, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65, 0xe2, 0x02, 0x19, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x5c, 0x44, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x3a,
	0x3a, 0x44, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_layer_dispute_voter_classes_proto_rawDescOnce sync.Once
	file_layer_dispute_voter_classes_proto_rawDescData = file_layer_dispute_voter_classes_proto_rawDesc
)

func file_layer_dispute_voter_classes_proto_rawDescGZIP() []byte {
	file_layer_dispute_voter_classes_proto_rawDescOnce.Do(func() {
		file_layer_dispute_voter_classes_proto_rawDescData = protoimpl.X.CompressGZIP(file_layer_dispute_voter_classes_proto_rawDescData)
	})
	return file_layer_dispute_voter_classes_proto_rawDescData
}

var file_layer_dispute_voter_classes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_layer_dispute_voter_classes_proto_goTypes = []interface{}{
	(*VoterClasses)(nil), // 0: layer.dispute.VoterClasses
}
var file_layer_dispute_voter_classes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_layer_dispute_voter_classes_proto_init() }
func file_layer_dispute_voter_classes_proto_init() {
	if File_layer_dispute_voter_classes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_layer_dispute_voter_classes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoterClasses); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_layer_dispute_voter_classes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_layer_dispute_voter_classes_proto_goTypes,
		DependencyIndexes: file_layer_dispute_voter_classes_proto_depIdxs,
		MessageInfos:      file_layer_dispute_voter_classes_proto_msgTypes,
	}.Build()
	File_layer_dispute_voter_classes_proto = out.File
	file_layer_dispute_voter_classes_proto_rawDesc = nil
	file_layer_dispute_voter_classes_proto_goTypes = nil
	file_layer_dispute_voter_classes_proto_depIdxs = nil
}