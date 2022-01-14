// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliance/provisional_model.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type ProvisionalModel struct {
	Vid               int32  `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Pid               int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	SoftwareVersion   uint32 `protobuf:"varint,3,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty"`
	CertificationType string `protobuf:"bytes,4,opt,name=certification_type,json=certificationType,proto3" json:"certification_type,omitempty"`
	Value             bool   `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ProvisionalModel) Reset()         { *m = ProvisionalModel{} }
func (m *ProvisionalModel) String() string { return proto.CompactTextString(m) }
func (*ProvisionalModel) ProtoMessage()    {}
func (*ProvisionalModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_228e3edf2a47c81f, []int{0}
}
func (m *ProvisionalModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProvisionalModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProvisionalModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProvisionalModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionalModel.Merge(m, src)
}
func (m *ProvisionalModel) XXX_Size() int {
	return m.Size()
}
func (m *ProvisionalModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionalModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionalModel proto.InternalMessageInfo

func (m *ProvisionalModel) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *ProvisionalModel) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProvisionalModel) GetSoftwareVersion() uint32 {
	if m != nil {
		return m.SoftwareVersion
	}
	return 0
}

func (m *ProvisionalModel) GetCertificationType() string {
	if m != nil {
		return m.CertificationType
	}
	return ""
}

func (m *ProvisionalModel) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*ProvisionalModel)(nil), "zigbeealliance.distributedcomplianceledger.compliance.ProvisionalModel")
}

func init() {
	proto.RegisterFile("compliance/provisional_model.proto", fileDescriptor_228e3edf2a47c81f)
}

var fileDescriptor_228e3edf2a47c81f = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xeb, 0xdb, 0x5b, 0x04, 0x96, 0x10, 0xc5, 0x62, 0xc8, 0x64, 0x45, 0x9d, 0xc2, 0x90,
	0x66, 0x40, 0xbc, 0x00, 0x2b, 0x42, 0x42, 0x11, 0x62, 0x60, 0xa9, 0x9c, 0xf8, 0x34, 0x1c, 0xc9,
	0x8d, 0x2d, 0xc7, 0x09, 0x94, 0xa7, 0xe0, 0x29, 0x78, 0x16, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x11,
	0x64, 0xac, 0x2a, 0x65, 0xb3, 0x3f, 0xfb, 0xfc, 0x3a, 0xff, 0x47, 0x17, 0xa5, 0xde, 0x18, 0x85,
	0xa2, 0x2e, 0x21, 0x33, 0x56, 0x77, 0xd8, 0xa0, 0xae, 0x85, 0x5a, 0x6d, 0xb4, 0x04, 0xb5, 0x34,
	0x56, 0x3b, 0xcd, 0xae, 0xdf, 0xb0, 0x2a, 0x00, 0x84, 0x0a, 0xff, 0x96, 0x12, 0x1b, 0x67, 0xb1,
	0x68, 0x1d, 0xc8, 0x71, 0x5a, 0x81, 0xac, 0xc0, 0x2e, 0x47, 0xb0, 0xf8, 0x20, 0x74, 0x7e, 0x3f,
	0x46, 0xde, 0xf9, 0x44, 0x36, 0xa7, 0xd3, 0x0e, 0x65, 0x44, 0x62, 0x92, 0xcc, 0x72, 0x7f, 0xf4,
	0xc4, 0xa0, 0x8c, 0xfe, 0x05, 0x62, 0x50, 0xb2, 0x4b, 0x3a, 0x6f, 0xf4, 0xda, 0xbd, 0x08, 0x0b,
	0xab, 0x0e, 0xac, 0x1f, 0x8f, 0xa6, 0x31, 0x49, 0x4e, 0xf3, 0xb3, 0x3d, 0x7f, 0x0c, 0x98, 0xa5,
	0x94, 0x95, 0x60, 0x1d, 0xae, 0xb1, 0x14, 0x0e, 0x75, 0xbd, 0x72, 0x5b, 0x03, 0xd1, 0xff, 0x98,
	0x24, 0x27, 0xf9, 0xf9, 0x9f, 0x97, 0x87, 0xad, 0x01, 0x76, 0x41, 0x67, 0x9d, 0x50, 0x2d, 0x44,
	0xb3, 0x98, 0x24, 0xc7, 0x79, 0xb8, 0xdc, 0xc0, 0x67, 0xcf, 0xc9, 0xae, 0xe7, 0xe4, 0xbb, 0xe7,
	0xe4, 0x7d, 0xe0, 0x93, 0xdd, 0xc0, 0x27, 0x5f, 0x03, 0x9f, 0x3c, 0xdd, 0x56, 0xe8, 0x9e, 0xdb,
	0xc2, 0xb7, 0xca, 0x82, 0x84, 0x74, 0x6f, 0x21, 0x3b, 0xb0, 0x90, 0x8e, 0xad, 0xd3, 0xe0, 0x21,
	0x7b, 0xcd, 0x0e, 0xc4, 0xfa, 0xa5, 0x9a, 0xe2, 0xe8, 0xd7, 0xe6, 0xd5, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0x5d, 0x89, 0xd6, 0x73, 0x01, 0x00, 0x00,
}

func (m *ProvisionalModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProvisionalModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProvisionalModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.CertificationType) > 0 {
		i -= len(m.CertificationType)
		copy(dAtA[i:], m.CertificationType)
		i = encodeVarintProvisionalModel(dAtA, i, uint64(len(m.CertificationType)))
		i--
		dAtA[i] = 0x22
	}
	if m.SoftwareVersion != 0 {
		i = encodeVarintProvisionalModel(dAtA, i, uint64(m.SoftwareVersion))
		i--
		dAtA[i] = 0x18
	}
	if m.Pid != 0 {
		i = encodeVarintProvisionalModel(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x10
	}
	if m.Vid != 0 {
		i = encodeVarintProvisionalModel(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvisionalModel(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvisionalModel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProvisionalModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovProvisionalModel(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovProvisionalModel(uint64(m.Pid))
	}
	if m.SoftwareVersion != 0 {
		n += 1 + sovProvisionalModel(uint64(m.SoftwareVersion))
	}
	l = len(m.CertificationType)
	if l > 0 {
		n += 1 + l + sovProvisionalModel(uint64(l))
	}
	if m.Value {
		n += 2
	}
	return n
}

func sovProvisionalModel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvisionalModel(x uint64) (n int) {
	return sovProvisionalModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProvisionalModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvisionalModel
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
			return fmt.Errorf("proto: ProvisionalModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProvisionalModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvisionalModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvisionalModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersion", wireType)
			}
			m.SoftwareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvisionalModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvisionalModel
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
				return ErrInvalidLengthProvisionalModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvisionalModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvisionalModel
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
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProvisionalModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvisionalModel
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
func skipProvisionalModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvisionalModel
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
					return 0, ErrIntOverflowProvisionalModel
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
					return 0, ErrIntOverflowProvisionalModel
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
				return 0, ErrInvalidLengthProvisionalModel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvisionalModel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvisionalModel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvisionalModel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvisionalModel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvisionalModel = fmt.Errorf("proto: unexpected end of group")
)
