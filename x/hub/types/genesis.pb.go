// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hub/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the hub module's genesis state.
type GenesisState struct {
	Params              Params               `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId              string               `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	StudentInfo         *StudentInfo         `protobuf:"bytes,3,opt,name=studentInfo,proto3" json:"studentInfo,omitempty"`
	TranscriptOfRecords *TranscriptOfRecords `protobuf:"bytes,4,opt,name=transcriptOfRecords,proto3" json:"transcriptOfRecords,omitempty"`
	PersonalInfo        *PersonalInfo        `protobuf:"bytes,5,opt,name=personalInfo,proto3" json:"personalInfo,omitempty"`
	ResidenceInfo       *ResidenceInfo       `protobuf:"bytes,6,opt,name=residenceInfo,proto3" json:"residenceInfo,omitempty"`
	ContactInfo         *ContactInfo         `protobuf:"bytes,7,opt,name=contactInfo,proto3" json:"contactInfo,omitempty"`
	TaxesInfo           *TaxesInfo           `protobuf:"bytes,8,opt,name=taxesInfo,proto3" json:"taxesInfo,omitempty"`
	ErasmusInfo         *ErasmusInfo         `protobuf:"bytes,9,opt,name=erasmusInfo,proto3" json:"erasmusInfo,omitempty"`
	StoredStudentList   []StoredStudent      `protobuf:"bytes,10,rep,name=storedStudentList,proto3" json:"storedStudentList"`
	ChainInfo           *ChainInfo           `protobuf:"bytes,11,opt,name=chainInfo,proto3" json:"chainInfo,omitempty"`
	UniversitiesList    []Universities       `protobuf:"bytes,12,rep,name=universitiesList,proto3" json:"universitiesList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_70299d674abbf152, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetStudentInfo() *StudentInfo {
	if m != nil {
		return m.StudentInfo
	}
	return nil
}

func (m *GenesisState) GetTranscriptOfRecords() *TranscriptOfRecords {
	if m != nil {
		return m.TranscriptOfRecords
	}
	return nil
}

func (m *GenesisState) GetPersonalInfo() *PersonalInfo {
	if m != nil {
		return m.PersonalInfo
	}
	return nil
}

func (m *GenesisState) GetResidenceInfo() *ResidenceInfo {
	if m != nil {
		return m.ResidenceInfo
	}
	return nil
}

func (m *GenesisState) GetContactInfo() *ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *GenesisState) GetTaxesInfo() *TaxesInfo {
	if m != nil {
		return m.TaxesInfo
	}
	return nil
}

func (m *GenesisState) GetErasmusInfo() *ErasmusInfo {
	if m != nil {
		return m.ErasmusInfo
	}
	return nil
}

func (m *GenesisState) GetStoredStudentList() []StoredStudent {
	if m != nil {
		return m.StoredStudentList
	}
	return nil
}

func (m *GenesisState) GetChainInfo() *ChainInfo {
	if m != nil {
		return m.ChainInfo
	}
	return nil
}

func (m *GenesisState) GetUniversitiesList() []Universities {
	if m != nil {
		return m.UniversitiesList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "hub.hub.GenesisState")
}

func init() { proto.RegisterFile("hub/genesis.proto", fileDescriptor_70299d674abbf152) }

var fileDescriptor_70299d674abbf152 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0xb6, 0x75, 0xd4, 0x2d, 0xda, 0x66, 0xca, 0x16, 0x55, 0x28, 0xab, 0x38, 0x8d,
	0x03, 0x2d, 0x02, 0x09, 0x09, 0x89, 0xd3, 0x26, 0x34, 0x0d, 0x21, 0x40, 0x2e, 0x5c, 0xb8, 0x44,
	0x6e, 0xe2, 0xb6, 0x96, 0x98, 0x1d, 0xd9, 0x0e, 0x1a, 0xdf, 0x82, 0x8f, 0xc4, 0x71, 0xc7, 0x1d,
	0x39, 0x21, 0xd4, 0x7e, 0x11, 0xe4, 0x37, 0x4e, 0xed, 0x36, 0xdc, 0x5a, 0x3f, 0xcf, 0xef, 0xfd,
	0x1f, 0x74, 0xb4, 0x28, 0xa7, 0xe3, 0x39, 0x13, 0x4c, 0x73, 0x3d, 0x2a, 0x94, 0x34, 0x12, 0xef,
	0x2f, 0xca, 0xe9, 0x68, 0x51, 0x4e, 0x07, 0xfd, 0xb9, 0x9c, 0x4b, 0x78, 0x1b, 0xdb, 0x5f, 0x95,
	0x3c, 0x38, 0xb4, 0x44, 0x41, 0x15, 0xbd, 0x76, 0xc0, 0xe0, 0xd8, 0xbe, 0x68, 0x53, 0xe6, 0x4c,
	0x98, 0x94, 0x8b, 0x59, 0xed, 0x3c, 0xb5, 0xef, 0x46, 0x51, 0xa1, 0x33, 0xc5, 0x0b, 0x93, 0xca,
	0x59, 0xaa, 0x58, 0x26, 0x55, 0x5e, 0x83, 0x27, 0x10, 0x8a, 0x29, 0x2d, 0x05, 0xfd, 0x16, 0x92,
	0xb1, 0x15, 0x14, 0xd3, 0x3c, 0x67, 0x22, 0x63, 0xa1, 0x02, 0xb9, 0x32, 0x29, 0x0c, 0xcd, 0x36,
	0x72, 0xf5, 0x21, 0x17, 0xbd, 0x61, 0xba, 0xe1, 0x66, 0x8a, 0xea, 0xeb, 0x52, 0x37, 0xe2, 0x6b,
	0x23, 0x15, 0xcb, 0x53, 0x57, 0x78, 0x18, 0x27, 0x5b, 0x50, 0x2e, 0x1a, 0x71, 0x4a, 0xc1, 0xbf,
	0x33, 0xa5, 0xb9, 0xe1, 0xcc, 0x35, 0xf0, 0xe4, 0xd7, 0x1e, 0xea, 0x5d, 0x56, 0xc3, 0x9b, 0x18,
	0x6a, 0x18, 0x7e, 0x86, 0xda, 0xd5, 0x68, 0xe2, 0x68, 0x18, 0x9d, 0x75, 0x5f, 0x1c, 0x8c, 0xdc,
	0x30, 0x47, 0x9f, 0xe0, 0xf9, 0x7c, 0xf7, 0xf6, 0xcf, 0x69, 0x8b, 0x38, 0x13, 0x3e, 0x41, 0xfb,
	0x85, 0x54, 0x26, 0xe5, 0x79, 0x7c, 0x6f, 0x18, 0x9d, 0x75, 0x48, 0xdb, 0xfe, 0xbd, 0xca, 0xf1,
	0x2b, 0xd4, 0x75, 0x75, 0x5d, 0x89, 0x99, 0x8c, 0x77, 0x20, 0x58, 0x7f, 0x1d, 0x6c, 0xe2, 0x35,
	0x12, 0x1a, 0xf1, 0x07, 0xf4, 0xd0, 0x0f, 0xfc, 0xe3, 0x8c, 0x54, 0xe3, 0x8e, 0x77, 0x81, 0x7f,
	0xbc, 0xe6, 0x3f, 0x37, 0x3d, 0xe4, 0x7f, 0x20, 0x7e, 0x8d, 0x7a, 0xf5, 0x7e, 0xa0, 0x90, 0x3d,
	0x08, 0xf4, 0xc8, 0x77, 0x15, 0x88, 0x64, 0xc3, 0x8a, 0xdf, 0xa0, 0x07, 0xeb, 0x0d, 0x02, 0xdb,
	0x06, 0xf6, 0x78, 0xcd, 0x92, 0x50, 0x25, 0x9b, 0x66, 0x3b, 0x00, 0xb7, 0x65, 0x60, 0xf7, 0xb7,
	0x06, 0x70, 0xe1, 0x35, 0x12, 0x1a, 0xf1, 0x73, 0xd4, 0x81, 0x2b, 0x00, 0xea, 0x3e, 0x50, 0xd8,
	0xb7, 0x5d, 0x2b, 0xc4, 0x9b, 0x6c, 0x26, 0x77, 0x21, 0xc0, 0x74, 0xb6, 0x32, 0xbd, 0xf5, 0x1a,
	0x09, 0x8d, 0xf8, 0x1d, 0x3a, 0xaa, 0x2e, 0xc8, 0x2d, 0xe3, 0x3d, 0xd7, 0x26, 0x46, 0xc3, 0x9d,
	0x8d, 0x1e, 0x27, 0xa1, 0xc3, 0x2d, 0xbf, 0x89, 0xd9, 0xaa, 0xe1, 0xe6, 0xa0, 0x82, 0xee, 0x56,
	0xd5, 0x17, 0xb5, 0x42, 0xbc, 0x09, 0x5f, 0xa2, 0xc3, 0xf0, 0x1e, 0x21, 0x79, 0x0f, 0x92, 0xfb,
	0xe5, 0x7c, 0x09, 0x0c, 0x2e, 0x77, 0x03, 0x3a, 0x7f, 0x7a, 0xbb, 0x4c, 0xa2, 0xbb, 0x65, 0x12,
	0xfd, 0x5d, 0x26, 0xd1, 0xcf, 0x55, 0xd2, 0xba, 0x5b, 0x25, 0xad, 0xdf, 0xab, 0xa4, 0xf5, 0xf5,
	0xc0, 0x1e, 0xfd, 0xcd, 0x18, 0x3e, 0xac, 0x1f, 0x05, 0xd3, 0xd3, 0x36, 0x1c, 0xfd, 0xcb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x07, 0x72, 0x22, 0x34, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UniversitiesList) > 0 {
		for iNdEx := len(m.UniversitiesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UniversitiesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.ChainInfo != nil {
		{
			size, err := m.ChainInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if len(m.StoredStudentList) > 0 {
		for iNdEx := len(m.StoredStudentList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoredStudentList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.ErasmusInfo != nil {
		{
			size, err := m.ErasmusInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.TaxesInfo != nil {
		{
			size, err := m.TaxesInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.ContactInfo != nil {
		{
			size, err := m.ContactInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.ResidenceInfo != nil {
		{
			size, err := m.ResidenceInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.PersonalInfo != nil {
		{
			size, err := m.PersonalInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.TranscriptOfRecords != nil {
		{
			size, err := m.TranscriptOfRecords.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.StudentInfo != nil {
		{
			size, err := m.StudentInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.StudentInfo != nil {
		l = m.StudentInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TranscriptOfRecords != nil {
		l = m.TranscriptOfRecords.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PersonalInfo != nil {
		l = m.PersonalInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ResidenceInfo != nil {
		l = m.ResidenceInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ContactInfo != nil {
		l = m.ContactInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TaxesInfo != nil {
		l = m.TaxesInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ErasmusInfo != nil {
		l = m.ErasmusInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.StoredStudentList) > 0 {
		for _, e := range m.StoredStudentList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ChainInfo != nil {
		l = m.ChainInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.UniversitiesList) > 0 {
		for _, e := range m.UniversitiesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StudentInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StudentInfo == nil {
				m.StudentInfo = &StudentInfo{}
			}
			if err := m.StudentInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TranscriptOfRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TranscriptOfRecords == nil {
				m.TranscriptOfRecords = &TranscriptOfRecords{}
			}
			if err := m.TranscriptOfRecords.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PersonalInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PersonalInfo == nil {
				m.PersonalInfo = &PersonalInfo{}
			}
			if err := m.PersonalInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResidenceInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResidenceInfo == nil {
				m.ResidenceInfo = &ResidenceInfo{}
			}
			if err := m.ResidenceInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContactInfo == nil {
				m.ContactInfo = &ContactInfo{}
			}
			if err := m.ContactInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxesInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TaxesInfo == nil {
				m.TaxesInfo = &TaxesInfo{}
			}
			if err := m.TaxesInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErasmusInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ErasmusInfo == nil {
				m.ErasmusInfo = &ErasmusInfo{}
			}
			if err := m.ErasmusInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredStudentList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoredStudentList = append(m.StoredStudentList, StoredStudent{})
			if err := m.StoredStudentList[len(m.StoredStudentList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainInfo == nil {
				m.ChainInfo = &ChainInfo{}
			}
			if err := m.ChainInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniversitiesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniversitiesList = append(m.UniversitiesList, Universities{})
			if err := m.UniversitiesList[len(m.UniversitiesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
