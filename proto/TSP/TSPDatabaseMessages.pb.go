// Code generated by protoc-gen-go.
// source: TSPDatabaseMessages.proto
// DO NOT EDIT!

/*
Package TSP is a generated protocol buffer package.

It is generated from these files:
	TSPDatabaseMessages.proto

It has these top-level messages:
	DatabaseData
	DatabaseDataArchive
	DatabaseImageDataArchive
*/
package TSP

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DatabaseImageDataArchive_ImageType int32

const (
	DatabaseImageDataArchive_unknown DatabaseImageDataArchive_ImageType = 0
	DatabaseImageDataArchive_bitmap  DatabaseImageDataArchive_ImageType = 1
	DatabaseImageDataArchive_pdf     DatabaseImageDataArchive_ImageType = 2
)

var DatabaseImageDataArchive_ImageType_name = map[int32]string{
	0: "unknown",
	1: "bitmap",
	2: "pdf",
}
var DatabaseImageDataArchive_ImageType_value = map[string]int32{
	"unknown": 0,
	"bitmap":  1,
	"pdf":     2,
}

func (x DatabaseImageDataArchive_ImageType) Enum() *DatabaseImageDataArchive_ImageType {
	p := new(DatabaseImageDataArchive_ImageType)
	*p = x
	return p
}
func (x DatabaseImageDataArchive_ImageType) String() string {
	return proto.EnumName(DatabaseImageDataArchive_ImageType_name, int32(x))
}
func (x *DatabaseImageDataArchive_ImageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatabaseImageDataArchive_ImageType_value, data, "DatabaseImageDataArchive_ImageType")
	if err != nil {
		return err
	}
	*x = DatabaseImageDataArchive_ImageType(value)
	return nil
}

type DatabaseData struct {
	Data             *DataReference `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *DatabaseData) Reset()         { *m = DatabaseData{} }
func (m *DatabaseData) String() string { return proto.CompactTextString(m) }
func (*DatabaseData) ProtoMessage()    {}

func (m *DatabaseData) GetData() *DataReference {
	if m != nil {
		return m.Data
	}
	return nil
}

type DatabaseDataArchive struct {
	Data             *Reference `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	AppRelativePath  *string    `protobuf:"bytes,2,opt,name=app_relative_path" json:"app_relative_path,omitempty"`
	DisplayName      *string    `protobuf:"bytes,3,req,name=display_name" json:"display_name,omitempty"`
	Length           *uint64    `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
	Hash             *uint32    `protobuf:"varint,5,opt,name=hash" json:"hash,omitempty"`
	Sharable         *bool      `protobuf:"varint,6,req,name=sharable,def=1" json:"sharable,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DatabaseDataArchive) Reset()         { *m = DatabaseDataArchive{} }
func (m *DatabaseDataArchive) String() string { return proto.CompactTextString(m) }
func (*DatabaseDataArchive) ProtoMessage()    {}

const Default_DatabaseDataArchive_Sharable bool = true

func (m *DatabaseDataArchive) GetData() *Reference {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DatabaseDataArchive) GetAppRelativePath() string {
	if m != nil && m.AppRelativePath != nil {
		return *m.AppRelativePath
	}
	return ""
}

func (m *DatabaseDataArchive) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *DatabaseDataArchive) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *DatabaseDataArchive) GetHash() uint32 {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return 0
}

func (m *DatabaseDataArchive) GetSharable() bool {
	if m != nil && m.Sharable != nil {
		return *m.Sharable
	}
	return Default_DatabaseDataArchive_Sharable
}

type DatabaseImageDataArchive struct {
	Super            *DatabaseDataArchive                `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Type             *DatabaseImageDataArchive_ImageType `protobuf:"varint,2,req,name=type,enum=DatabaseImageDataArchive_ImageType" json:"type,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *DatabaseImageDataArchive) Reset()         { *m = DatabaseImageDataArchive{} }
func (m *DatabaseImageDataArchive) String() string { return proto.CompactTextString(m) }
func (*DatabaseImageDataArchive) ProtoMessage()    {}

func (m *DatabaseImageDataArchive) GetSuper() *DatabaseDataArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *DatabaseImageDataArchive) GetType() DatabaseImageDataArchive_ImageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return DatabaseImageDataArchive_unknown
}

func init() {
	proto.RegisterEnum("DatabaseImageDataArchive_ImageType", DatabaseImageDataArchive_ImageType_name, DatabaseImageDataArchive_ImageType_value)
}