// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quran/quran.proto

package quran

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DefaultResponse_Status int32

const (
	DefaultResponse_UNKOWN                DefaultResponse_Status = 0
	DefaultResponse_OK                    DefaultResponse_Status = 200
	DefaultResponse_CREATED               DefaultResponse_Status = 201
	DefaultResponse_NO_CONTENT            DefaultResponse_Status = 204
	DefaultResponse_BAD_REQUEST           DefaultResponse_Status = 400
	DefaultResponse_UNAUTHORIZED          DefaultResponse_Status = 401
	DefaultResponse_FORBIDDEN             DefaultResponse_Status = 403
	DefaultResponse_NOT_FOUND             DefaultResponse_Status = 404
	DefaultResponse_METHOD_NOT_ALLOWED    DefaultResponse_Status = 405
	DefaultResponse_GONE                  DefaultResponse_Status = 410
	DefaultResponse_UNPROCESSABLE_ENTITY  DefaultResponse_Status = 422
	DefaultResponse_INTERNAL_SERVER_ERROR DefaultResponse_Status = 500
	DefaultResponse_BAD_GATEWAY           DefaultResponse_Status = 502
)

var DefaultResponse_Status_name = map[int32]string{
	0:   "UNKOWN",
	200: "OK",
	201: "CREATED",
	204: "NO_CONTENT",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	405: "METHOD_NOT_ALLOWED",
	410: "GONE",
	422: "UNPROCESSABLE_ENTITY",
	500: "INTERNAL_SERVER_ERROR",
	502: "BAD_GATEWAY",
}

var DefaultResponse_Status_value = map[string]int32{
	"UNKOWN":                0,
	"OK":                    200,
	"CREATED":               201,
	"NO_CONTENT":            204,
	"BAD_REQUEST":           400,
	"UNAUTHORIZED":          401,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"METHOD_NOT_ALLOWED":    405,
	"GONE":                  410,
	"UNPROCESSABLE_ENTITY":  422,
	"INTERNAL_SERVER_ERROR": 500,
	"BAD_GATEWAY":           502,
}

func (x DefaultResponse_Status) String() string {
	return proto.EnumName(DefaultResponse_Status_name, int32(x))
}

func (DefaultResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{1, 0}
}

type EmptyPayload struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyPayload) Reset()         { *m = EmptyPayload{} }
func (m *EmptyPayload) String() string { return proto.CompactTextString(m) }
func (*EmptyPayload) ProtoMessage()    {}
func (*EmptyPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{0}
}

func (m *EmptyPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyPayload.Unmarshal(m, b)
}
func (m *EmptyPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyPayload.Marshal(b, m, deterministic)
}
func (m *EmptyPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyPayload.Merge(m, src)
}
func (m *EmptyPayload) XXX_Size() int {
	return xxx_messageInfo_EmptyPayload.Size(m)
}
func (m *EmptyPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyPayload proto.InternalMessageInfo

type DefaultResponse struct {
	Status               DefaultResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=quran.DefaultResponse_Status" json:"status,omitempty"`
	Message              string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DefaultResponse) Reset()         { *m = DefaultResponse{} }
func (m *DefaultResponse) String() string { return proto.CompactTextString(m) }
func (*DefaultResponse) ProtoMessage()    {}
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{1}
}

func (m *DefaultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultResponse.Unmarshal(m, b)
}
func (m *DefaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultResponse.Marshal(b, m, deterministic)
}
func (m *DefaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultResponse.Merge(m, src)
}
func (m *DefaultResponse) XXX_Size() int {
	return xxx_messageInfo_DefaultResponse.Size(m)
}
func (m *DefaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultResponse proto.InternalMessageInfo

func (m *DefaultResponse) GetStatus() DefaultResponse_Status {
	if m != nil {
		return m.Status
	}
	return DefaultResponse_UNKOWN
}

func (m *DefaultResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Tafseer struct {
	TafseerSlug          string   `protobuf:"bytes,1,opt,name=tafseerSlug,proto3" json:"tafseerSlug,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Tafseer              string   `protobuf:"bytes,3,opt,name=tafseer,proto3" json:"tafseer,omitempty"`
	Citations            string   `protobuf:"bytes,4,opt,name=citations,proto3" json:"citations,omitempty"`
	Uri                  string   `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	MediaBase64          string   `protobuf:"bytes,6,opt,name=mediaBase64,proto3" json:"mediaBase64,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tafseer) Reset()         { *m = Tafseer{} }
func (m *Tafseer) String() string { return proto.CompactTextString(m) }
func (*Tafseer) ProtoMessage()    {}
func (*Tafseer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{2}
}

func (m *Tafseer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tafseer.Unmarshal(m, b)
}
func (m *Tafseer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tafseer.Marshal(b, m, deterministic)
}
func (m *Tafseer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tafseer.Merge(m, src)
}
func (m *Tafseer) XXX_Size() int {
	return xxx_messageInfo_Tafseer.Size(m)
}
func (m *Tafseer) XXX_DiscardUnknown() {
	xxx_messageInfo_Tafseer.DiscardUnknown(m)
}

var xxx_messageInfo_Tafseer proto.InternalMessageInfo

func (m *Tafseer) GetTafseerSlug() string {
	if m != nil {
		return m.TafseerSlug
	}
	return ""
}

func (m *Tafseer) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Tafseer) GetTafseer() string {
	if m != nil {
		return m.Tafseer
	}
	return ""
}

func (m *Tafseer) GetCitations() string {
	if m != nil {
		return m.Citations
	}
	return ""
}

func (m *Tafseer) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Tafseer) GetMediaBase64() string {
	if m != nil {
		return m.MediaBase64
	}
	return ""
}

type AddTafseerRequest struct {
	AyahSlug             string   `protobuf:"bytes,1,opt,name=ayahSlug,proto3" json:"ayahSlug,omitempty"`
	Tafseer              *Tafseer `protobuf:"bytes,2,opt,name=tafseer,proto3" json:"tafseer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTafseerRequest) Reset()         { *m = AddTafseerRequest{} }
func (m *AddTafseerRequest) String() string { return proto.CompactTextString(m) }
func (*AddTafseerRequest) ProtoMessage()    {}
func (*AddTafseerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{3}
}

func (m *AddTafseerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTafseerRequest.Unmarshal(m, b)
}
func (m *AddTafseerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTafseerRequest.Marshal(b, m, deterministic)
}
func (m *AddTafseerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTafseerRequest.Merge(m, src)
}
func (m *AddTafseerRequest) XXX_Size() int {
	return xxx_messageInfo_AddTafseerRequest.Size(m)
}
func (m *AddTafseerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTafseerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTafseerRequest proto.InternalMessageInfo

func (m *AddTafseerRequest) GetAyahSlug() string {
	if m != nil {
		return m.AyahSlug
	}
	return ""
}

func (m *AddTafseerRequest) GetTafseer() *Tafseer {
	if m != nil {
		return m.Tafseer
	}
	return nil
}

type Ayah struct {
	// surah-ayahnumber
	AyahSlug string `protobuf:"bytes,1,opt,name=ayahSlug,proto3" json:"ayahSlug,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// Audio mp3 to json utf8
	Recitation           []string   `protobuf:"bytes,3,rep,name=recitation,proto3" json:"recitation,omitempty"`
	Tafseer              []*Tafseer `protobuf:"bytes,4,rep,name=tafseer,proto3" json:"tafseer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Ayah) Reset()         { *m = Ayah{} }
func (m *Ayah) String() string { return proto.CompactTextString(m) }
func (*Ayah) ProtoMessage()    {}
func (*Ayah) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{4}
}

func (m *Ayah) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ayah.Unmarshal(m, b)
}
func (m *Ayah) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ayah.Marshal(b, m, deterministic)
}
func (m *Ayah) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ayah.Merge(m, src)
}
func (m *Ayah) XXX_Size() int {
	return xxx_messageInfo_Ayah.Size(m)
}
func (m *Ayah) XXX_DiscardUnknown() {
	xxx_messageInfo_Ayah.DiscardUnknown(m)
}

var xxx_messageInfo_Ayah proto.InternalMessageInfo

func (m *Ayah) GetAyahSlug() string {
	if m != nil {
		return m.AyahSlug
	}
	return ""
}

func (m *Ayah) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Ayah) GetRecitation() []string {
	if m != nil {
		return m.Recitation
	}
	return nil
}

func (m *Ayah) GetTafseer() []*Tafseer {
	if m != nil {
		return m.Tafseer
	}
	return nil
}

type AddRecitatonToAyahRequest struct {
	AyahSlug string `protobuf:"bytes,1,opt,name=ayahSlug,proto3" json:"ayahSlug,omitempty"`
	// Audio mp3 to json utf8
	Recitation           string   `protobuf:"bytes,2,opt,name=recitation,proto3" json:"recitation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRecitatonToAyahRequest) Reset()         { *m = AddRecitatonToAyahRequest{} }
func (m *AddRecitatonToAyahRequest) String() string { return proto.CompactTextString(m) }
func (*AddRecitatonToAyahRequest) ProtoMessage()    {}
func (*AddRecitatonToAyahRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{5}
}

func (m *AddRecitatonToAyahRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRecitatonToAyahRequest.Unmarshal(m, b)
}
func (m *AddRecitatonToAyahRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRecitatonToAyahRequest.Marshal(b, m, deterministic)
}
func (m *AddRecitatonToAyahRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRecitatonToAyahRequest.Merge(m, src)
}
func (m *AddRecitatonToAyahRequest) XXX_Size() int {
	return xxx_messageInfo_AddRecitatonToAyahRequest.Size(m)
}
func (m *AddRecitatonToAyahRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRecitatonToAyahRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRecitatonToAyahRequest proto.InternalMessageInfo

func (m *AddRecitatonToAyahRequest) GetAyahSlug() string {
	if m != nil {
		return m.AyahSlug
	}
	return ""
}

func (m *AddRecitatonToAyahRequest) GetRecitation() string {
	if m != nil {
		return m.Recitation
	}
	return ""
}

type AddRecitatonToSurahRequest struct {
	SurahName string `protobuf:"bytes,1,opt,name=surahName,proto3" json:"surahName,omitempty"`
	// Audio mp3 to json utf8
	Recitation           string   `protobuf:"bytes,2,opt,name=recitation,proto3" json:"recitation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRecitatonToSurahRequest) Reset()         { *m = AddRecitatonToSurahRequest{} }
func (m *AddRecitatonToSurahRequest) String() string { return proto.CompactTextString(m) }
func (*AddRecitatonToSurahRequest) ProtoMessage()    {}
func (*AddRecitatonToSurahRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{6}
}

func (m *AddRecitatonToSurahRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRecitatonToSurahRequest.Unmarshal(m, b)
}
func (m *AddRecitatonToSurahRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRecitatonToSurahRequest.Marshal(b, m, deterministic)
}
func (m *AddRecitatonToSurahRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRecitatonToSurahRequest.Merge(m, src)
}
func (m *AddRecitatonToSurahRequest) XXX_Size() int {
	return xxx_messageInfo_AddRecitatonToSurahRequest.Size(m)
}
func (m *AddRecitatonToSurahRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRecitatonToSurahRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRecitatonToSurahRequest proto.InternalMessageInfo

func (m *AddRecitatonToSurahRequest) GetSurahName() string {
	if m != nil {
		return m.SurahName
	}
	return ""
}

func (m *AddRecitatonToSurahRequest) GetRecitation() string {
	if m != nil {
		return m.Recitation
	}
	return ""
}

type Surah struct {
	SruhName             string   `protobuf:"bytes,1,opt,name=sruhName,proto3" json:"sruhName,omitempty"`
	Ayah                 []*Ayah  `protobuf:"bytes,2,rep,name=ayah,proto3" json:"ayah,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Surah) Reset()         { *m = Surah{} }
func (m *Surah) String() string { return proto.CompactTextString(m) }
func (*Surah) ProtoMessage()    {}
func (*Surah) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{7}
}

func (m *Surah) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Surah.Unmarshal(m, b)
}
func (m *Surah) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Surah.Marshal(b, m, deterministic)
}
func (m *Surah) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Surah.Merge(m, src)
}
func (m *Surah) XXX_Size() int {
	return xxx_messageInfo_Surah.Size(m)
}
func (m *Surah) XXX_DiscardUnknown() {
	xxx_messageInfo_Surah.DiscardUnknown(m)
}

var xxx_messageInfo_Surah proto.InternalMessageInfo

func (m *Surah) GetSruhName() string {
	if m != nil {
		return m.SruhName
	}
	return ""
}

func (m *Surah) GetAyah() []*Ayah {
	if m != nil {
		return m.Ayah
	}
	return nil
}

// chapter is Goza`
type Chapter struct {
	ChapterId            int32    `protobuf:"varint,1,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
	Surahs               []*Surah `protobuf:"bytes,2,rep,name=surahs,proto3" json:"surahs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chapter) Reset()         { *m = Chapter{} }
func (m *Chapter) String() string { return proto.CompactTextString(m) }
func (*Chapter) ProtoMessage()    {}
func (*Chapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{8}
}

func (m *Chapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chapter.Unmarshal(m, b)
}
func (m *Chapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chapter.Marshal(b, m, deterministic)
}
func (m *Chapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chapter.Merge(m, src)
}
func (m *Chapter) XXX_Size() int {
	return xxx_messageInfo_Chapter.Size(m)
}
func (m *Chapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Chapter.DiscardUnknown(m)
}

var xxx_messageInfo_Chapter proto.InternalMessageInfo

func (m *Chapter) GetChapterId() int32 {
	if m != nil {
		return m.ChapterId
	}
	return 0
}

func (m *Chapter) GetSurahs() []*Surah {
	if m != nil {
		return m.Surahs
	}
	return nil
}

type Quran struct {
	Surahs               []*Surah `protobuf:"bytes,1,rep,name=surahs,proto3" json:"surahs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quran) Reset()         { *m = Quran{} }
func (m *Quran) String() string { return proto.CompactTextString(m) }
func (*Quran) ProtoMessage()    {}
func (*Quran) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{9}
}

func (m *Quran) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quran.Unmarshal(m, b)
}
func (m *Quran) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quran.Marshal(b, m, deterministic)
}
func (m *Quran) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quran.Merge(m, src)
}
func (m *Quran) XXX_Size() int {
	return xxx_messageInfo_Quran.Size(m)
}
func (m *Quran) XXX_DiscardUnknown() {
	xxx_messageInfo_Quran.DiscardUnknown(m)
}

var xxx_messageInfo_Quran proto.InternalMessageInfo

func (m *Quran) GetSurahs() []*Surah {
	if m != nil {
		return m.Surahs
	}
	return nil
}

type GetChapterRequest struct {
	ChapterId            int32    `protobuf:"varint,1,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChapterRequest) Reset()         { *m = GetChapterRequest{} }
func (m *GetChapterRequest) String() string { return proto.CompactTextString(m) }
func (*GetChapterRequest) ProtoMessage()    {}
func (*GetChapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{10}
}

func (m *GetChapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChapterRequest.Unmarshal(m, b)
}
func (m *GetChapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChapterRequest.Marshal(b, m, deterministic)
}
func (m *GetChapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChapterRequest.Merge(m, src)
}
func (m *GetChapterRequest) XXX_Size() int {
	return xxx_messageInfo_GetChapterRequest.Size(m)
}
func (m *GetChapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChapterRequest proto.InternalMessageInfo

func (m *GetChapterRequest) GetChapterId() int32 {
	if m != nil {
		return m.ChapterId
	}
	return 0
}

type GetSurahRequest struct {
	SurahName            string   `protobuf:"bytes,1,opt,name=surahName,proto3" json:"surahName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSurahRequest) Reset()         { *m = GetSurahRequest{} }
func (m *GetSurahRequest) String() string { return proto.CompactTextString(m) }
func (*GetSurahRequest) ProtoMessage()    {}
func (*GetSurahRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{11}
}

func (m *GetSurahRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSurahRequest.Unmarshal(m, b)
}
func (m *GetSurahRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSurahRequest.Marshal(b, m, deterministic)
}
func (m *GetSurahRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSurahRequest.Merge(m, src)
}
func (m *GetSurahRequest) XXX_Size() int {
	return xxx_messageInfo_GetSurahRequest.Size(m)
}
func (m *GetSurahRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSurahRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSurahRequest proto.InternalMessageInfo

func (m *GetSurahRequest) GetSurahName() string {
	if m != nil {
		return m.SurahName
	}
	return ""
}

type GetAyahRequest struct {
	AyahSlug             string   `protobuf:"bytes,1,opt,name=ayahSlug,proto3" json:"ayahSlug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAyahRequest) Reset()         { *m = GetAyahRequest{} }
func (m *GetAyahRequest) String() string { return proto.CompactTextString(m) }
func (*GetAyahRequest) ProtoMessage()    {}
func (*GetAyahRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24e86525da88af0, []int{12}
}

func (m *GetAyahRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAyahRequest.Unmarshal(m, b)
}
func (m *GetAyahRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAyahRequest.Marshal(b, m, deterministic)
}
func (m *GetAyahRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAyahRequest.Merge(m, src)
}
func (m *GetAyahRequest) XXX_Size() int {
	return xxx_messageInfo_GetAyahRequest.Size(m)
}
func (m *GetAyahRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAyahRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAyahRequest proto.InternalMessageInfo

func (m *GetAyahRequest) GetAyahSlug() string {
	if m != nil {
		return m.AyahSlug
	}
	return ""
}

func init() {
	proto.RegisterEnum("quran.DefaultResponse_Status", DefaultResponse_Status_name, DefaultResponse_Status_value)
	proto.RegisterType((*EmptyPayload)(nil), "quran.EmptyPayload")
	proto.RegisterType((*DefaultResponse)(nil), "quran.DefaultResponse")
	proto.RegisterType((*Tafseer)(nil), "quran.Tafseer")
	proto.RegisterType((*AddTafseerRequest)(nil), "quran.AddTafseerRequest")
	proto.RegisterType((*Ayah)(nil), "quran.Ayah")
	proto.RegisterType((*AddRecitatonToAyahRequest)(nil), "quran.AddRecitatonToAyahRequest")
	proto.RegisterType((*AddRecitatonToSurahRequest)(nil), "quran.AddRecitatonToSurahRequest")
	proto.RegisterType((*Surah)(nil), "quran.Surah")
	proto.RegisterType((*Chapter)(nil), "quran.Chapter")
	proto.RegisterType((*Quran)(nil), "quran.Quran")
	proto.RegisterType((*GetChapterRequest)(nil), "quran.GetChapterRequest")
	proto.RegisterType((*GetSurahRequest)(nil), "quran.GetSurahRequest")
	proto.RegisterType((*GetAyahRequest)(nil), "quran.GetAyahRequest")
}

func init() { proto.RegisterFile("quran/quran.proto", fileDescriptor_f24e86525da88af0) }

var fileDescriptor_f24e86525da88af0 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x8f, 0xdb, 0x44,
	0x18, 0xc6, 0xeb, 0x7c, 0x34, 0xef, 0x46, 0x59, 0xef, 0x5b, 0x5a, 0x5c, 0xab, 0x40, 0xb0, 0x38,
	0xe4, 0xc0, 0x6e, 0xcb, 0x02, 0x3d, 0x22, 0x79, 0xd7, 0xd3, 0x34, 0x6a, 0x6a, 0x6f, 0xc7, 0x0e,
	0x51, 0x7a, 0x89, 0x86, 0x66, 0xba, 0x1b, 0x29, 0x5f, 0xb5, 0xc7, 0x88, 0xdc, 0xf8, 0x09, 0x20,
	0xe0, 0xc2, 0x0f, 0xe0, 0x0f, 0xf0, 0x13, 0xb8, 0x80, 0xc4, 0x8d, 0xff, 0xc1, 0x09, 0x71, 0x46,
	0x1e, 0x8f, 0x63, 0xa7, 0x64, 0x97, 0xbd, 0x44, 0xf3, 0x3e, 0xef, 0xe3, 0xc7, 0xcf, 0xfb, 0xcc,
	0x8c, 0x03, 0x87, 0xaf, 0x93, 0x88, 0x2d, 0x1e, 0xc8, 0xdf, 0xe3, 0x55, 0xb4, 0x14, 0x4b, 0xac,
	0xca, 0xc2, 0x6e, 0x41, 0x93, 0xcc, 0x57, 0x62, 0x7d, 0xce, 0xd6, 0xb3, 0x25, 0x9b, 0xd8, 0x7f,
	0xee, 0xc1, 0x81, 0xcb, 0x5f, 0xb1, 0x64, 0x26, 0x28, 0x8f, 0x57, 0xcb, 0x45, 0xcc, 0xf1, 0x33,
	0xa8, 0xc5, 0x82, 0x89, 0x24, 0x36, 0xb5, 0xb6, 0xd6, 0x69, 0x9d, 0xbc, 0x7b, 0x9c, 0x09, 0xbd,
	0xc1, 0x3b, 0x0e, 0x24, 0x89, 0x2a, 0x32, 0x9a, 0x50, 0x9f, 0xf3, 0x38, 0x66, 0x17, 0xdc, 0xdc,
	0x6b, 0x6b, 0x9d, 0x06, 0xcd, 0x4b, 0xfb, 0x2f, 0x0d, 0x6a, 0x19, 0x19, 0x01, 0x6a, 0x03, 0xef,
	0xa9, 0x3f, 0xf4, 0x8c, 0xb7, 0xb0, 0x0e, 0x7b, 0xfe, 0x53, 0xe3, 0x37, 0x0d, 0x9b, 0x50, 0x3f,
	0xa3, 0xc4, 0x09, 0x89, 0x6b, 0xfc, 0xae, 0xe1, 0x01, 0x80, 0xe7, 0x8f, 0xcf, 0x7c, 0x2f, 0x24,
	0x5e, 0x68, 0xfc, 0xa1, 0xa1, 0x01, 0xfb, 0xa7, 0x8e, 0x3b, 0xa6, 0xe4, 0xf9, 0x80, 0x04, 0xa1,
	0xf1, 0xad, 0x8e, 0x87, 0xd0, 0x1c, 0x78, 0xce, 0x20, 0x7c, 0xe2, 0xd3, 0xde, 0x0b, 0xe2, 0x1a,
	0xdf, 0xe9, 0xd8, 0x82, 0xc6, 0x63, 0x9f, 0x9e, 0xf6, 0x5c, 0x97, 0x78, 0xc6, 0xf7, 0xb2, 0xf6,
	0xfc, 0x70, 0xfc, 0xd8, 0x1f, 0x78, 0xae, 0xf1, 0x83, 0x8e, 0xef, 0x00, 0x3e, 0x23, 0xe1, 0x13,
	0xdf, 0x1d, 0xa7, 0xb0, 0xd3, 0xef, 0xfb, 0x43, 0xe2, 0x1a, 0x3f, 0xea, 0xd8, 0x80, 0x4a, 0xd7,
	0xf7, 0x88, 0xf1, 0x93, 0x8e, 0xf7, 0xe0, 0xed, 0x81, 0x77, 0x4e, 0xfd, 0x33, 0x12, 0x04, 0xce,
	0x69, 0x9f, 0x8c, 0x89, 0x17, 0xf6, 0xc2, 0x91, 0xf1, 0xb3, 0x8e, 0x16, 0xdc, 0xe9, 0x79, 0x21,
	0xa1, 0x9e, 0xd3, 0x1f, 0x07, 0x84, 0x7e, 0x41, 0xe8, 0x98, 0x50, 0xea, 0x53, 0xe3, 0x6f, 0x3d,
	0xf7, 0xd7, 0x75, 0x42, 0x32, 0x74, 0x46, 0xc6, 0x3f, 0xba, 0xfd, 0x8b, 0x06, 0xf5, 0x90, 0xbd,
	0x8a, 0x39, 0x8f, 0xb0, 0x0d, 0xfb, 0x22, 0x5b, 0x06, 0xb3, 0xe4, 0x42, 0x46, 0xda, 0xa0, 0x65,
	0x08, 0x2d, 0xb8, 0x35, 0x63, 0x8b, 0x8b, 0xa4, 0x48, 0x6e, 0x53, 0xa7, 0xa1, 0x2a, 0xaa, 0xa9,
	0x67, 0xa1, 0xaa, 0x12, 0xef, 0x43, 0xe3, 0xe5, 0x54, 0x30, 0x31, 0x5d, 0x2e, 0x62, 0xb3, 0x22,
	0x7b, 0x05, 0x80, 0x06, 0xe8, 0x49, 0x34, 0x35, 0xab, 0x12, 0x4f, 0x97, 0xa9, 0x8f, 0x39, 0x9f,
	0x4c, 0xd9, 0x29, 0x8b, 0xf9, 0xa3, 0x4f, 0xcd, 0x5a, 0xe6, 0xa3, 0x04, 0xd9, 0x23, 0x38, 0x74,
	0x26, 0x13, 0xe5, 0x9b, 0xf2, 0xd7, 0x09, 0x8f, 0x45, 0x6a, 0x8e, 0xad, 0xd9, 0x65, 0xc9, 0xfb,
	0xa6, 0xc6, 0x4e, 0x61, 0x2e, 0xf5, 0xbd, 0x7f, 0xd2, 0x52, 0x27, 0x25, 0xd7, 0xc8, 0xdb, 0xf6,
	0x37, 0x1a, 0x54, 0x9c, 0x35, 0xbb, 0xbc, 0x56, 0x0e, 0xa1, 0x22, 0xf8, 0xd7, 0x42, 0x65, 0x20,
	0xd7, 0xf8, 0x1e, 0x40, 0xc4, 0xf3, 0xb1, 0x4c, 0xbd, 0xad, 0x77, 0x1a, 0xb4, 0x84, 0x94, 0x2d,
	0x54, 0xda, 0xfa, 0x75, 0x16, 0x86, 0x70, 0xcf, 0x99, 0x4c, 0x68, 0xf6, 0xe8, 0x72, 0x11, 0x2e,
	0x53, 0x3f, 0x37, 0x99, 0x72, 0xdb, 0x42, 0x66, 0xae, 0x84, 0xd8, 0x2f, 0xc0, 0xda, 0x16, 0x0e,
	0x92, 0xa8, 0x50, 0xbe, 0x0f, 0x8d, 0x38, 0xad, 0x3d, 0x36, 0xe7, 0x4a, 0xba, 0x00, 0xfe, 0x57,
	0xdb, 0x85, 0xaa, 0x54, 0x4b, 0x0d, 0xc6, 0x51, 0x52, 0x56, 0xd9, 0xd4, 0xf8, 0x3e, 0x54, 0x52,
	0xb3, 0xe6, 0x9e, 0x0c, 0x60, 0x5f, 0x05, 0x20, 0xc7, 0x93, 0x0d, 0xfb, 0x19, 0xd4, 0xcf, 0x2e,
	0xd9, 0x4a, 0xa8, 0x53, 0x93, 0x2d, 0x7b, 0x13, 0x29, 0x54, 0xa5, 0x05, 0x80, 0x1f, 0x42, 0x4d,
	0x7a, 0x8b, 0x95, 0x56, 0x53, 0x69, 0x65, 0x13, 0xa9, 0x9e, 0x7d, 0x04, 0xd5, 0xe7, 0x29, 0x5c,
	0xa2, 0x6b, 0xd7, 0xd0, 0x3f, 0x86, 0xc3, 0x2e, 0x17, 0xca, 0x40, 0x29, 0x96, 0xab, 0x7d, 0xd8,
	0x0f, 0xe0, 0xa0, 0xcb, 0xc5, 0xcd, 0x73, 0xb4, 0x3f, 0x82, 0x56, 0x97, 0x8b, 0x1b, 0xee, 0xe8,
	0xc9, 0xaf, 0x3a, 0x34, 0xe5, 0x04, 0x01, 0x8f, 0xbe, 0x9a, 0xbe, 0xe4, 0x78, 0x04, 0xb7, 0xba,
	0x5c, 0x64, 0x43, 0xdd, 0x56, 0x43, 0x94, 0x3f, 0x93, 0x56, 0x3e, 0x59, 0x46, 0x39, 0x82, 0xba,
	0x7a, 0x1b, 0xde, 0x51, 0x8d, 0xed, 0xb7, 0x5b, 0xe5, 0x4d, 0xc0, 0x87, 0x52, 0x3d, 0xdb, 0xc7,
	0xbb, 0x05, 0xbf, 0x3c, 0x9e, 0xb5, 0x15, 0x1d, 0x3e, 0x02, 0x28, 0x22, 0x43, 0xb3, 0x78, 0x66,
	0x3b, 0x45, 0x2b, 0x3f, 0xec, 0x39, 0xf3, 0x73, 0x80, 0xe2, 0x06, 0x6f, 0x9e, 0xfb, 0xcf, 0xa5,
	0xb6, 0xee, 0xee, 0xfe, 0xa2, 0xe3, 0xe8, 0xcd, 0x3b, 0x22, 0xed, 0x04, 0x22, 0xe2, 0x6c, 0x8e,
	0xed, 0x42, 0x6e, 0xf7, 0x2d, 0xba, 0x4a, 0xb6, 0xa3, 0x3d, 0xd4, 0xf0, 0x1c, 0x6e, 0xef, 0x90,
	0xc6, 0x0f, 0x76, 0x8a, 0x6e, 0x45, 0x73, 0x85, 0xea, 0x97, 0x35, 0xf9, 0xc7, 0xf6, 0xc9, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xa5, 0x15, 0xe2, 0xed, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuranServiceClient is the client API for QuranService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuranServiceClient interface {
	// Getters
	GetQuran(ctx context.Context, in *EmptyPayload, opts ...grpc.CallOption) (*Quran, error)
	GetAyah(ctx context.Context, in *GetAyahRequest, opts ...grpc.CallOption) (*Ayah, error)
	GetSurah(ctx context.Context, in *GetSurahRequest, opts ...grpc.CallOption) (*Surah, error)
	GetChapter(ctx context.Context, in *GetChapterRequest, opts ...grpc.CallOption) (*Chapter, error)
	AddTafseer(ctx context.Context, in *AddTafseerRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	AddRecitatonToSurahStream(ctx context.Context, opts ...grpc.CallOption) (QuranService_AddRecitatonToSurahStreamClient, error)
	AddRecitatonToSurah(ctx context.Context, in *AddRecitatonToSurahRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type quranServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuranServiceClient(cc *grpc.ClientConn) QuranServiceClient {
	return &quranServiceClient{cc}
}

func (c *quranServiceClient) GetQuran(ctx context.Context, in *EmptyPayload, opts ...grpc.CallOption) (*Quran, error) {
	out := new(Quran)
	err := c.cc.Invoke(ctx, "/quran.QuranService/GetQuran", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranServiceClient) GetAyah(ctx context.Context, in *GetAyahRequest, opts ...grpc.CallOption) (*Ayah, error) {
	out := new(Ayah)
	err := c.cc.Invoke(ctx, "/quran.QuranService/GetAyah", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranServiceClient) GetSurah(ctx context.Context, in *GetSurahRequest, opts ...grpc.CallOption) (*Surah, error) {
	out := new(Surah)
	err := c.cc.Invoke(ctx, "/quran.QuranService/GetSurah", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranServiceClient) GetChapter(ctx context.Context, in *GetChapterRequest, opts ...grpc.CallOption) (*Chapter, error) {
	out := new(Chapter)
	err := c.cc.Invoke(ctx, "/quran.QuranService/GetChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranServiceClient) AddTafseer(ctx context.Context, in *AddTafseerRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/quran.QuranService/AddTafseer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranServiceClient) AddRecitatonToSurahStream(ctx context.Context, opts ...grpc.CallOption) (QuranService_AddRecitatonToSurahStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QuranService_serviceDesc.Streams[0], "/quran.QuranService/AddRecitatonToSurahStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &quranServiceAddRecitatonToSurahStreamClient{stream}
	return x, nil
}

type QuranService_AddRecitatonToSurahStreamClient interface {
	Send(*AddRecitatonToAyahRequest) error
	Recv() (*DefaultResponse, error)
	grpc.ClientStream
}

type quranServiceAddRecitatonToSurahStreamClient struct {
	grpc.ClientStream
}

func (x *quranServiceAddRecitatonToSurahStreamClient) Send(m *AddRecitatonToAyahRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *quranServiceAddRecitatonToSurahStreamClient) Recv() (*DefaultResponse, error) {
	m := new(DefaultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *quranServiceClient) AddRecitatonToSurah(ctx context.Context, in *AddRecitatonToSurahRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/quran.QuranService/AddRecitatonToSurah", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuranServiceServer is the server API for QuranService service.
type QuranServiceServer interface {
	// Getters
	GetQuran(context.Context, *EmptyPayload) (*Quran, error)
	GetAyah(context.Context, *GetAyahRequest) (*Ayah, error)
	GetSurah(context.Context, *GetSurahRequest) (*Surah, error)
	GetChapter(context.Context, *GetChapterRequest) (*Chapter, error)
	AddTafseer(context.Context, *AddTafseerRequest) (*DefaultResponse, error)
	AddRecitatonToSurahStream(QuranService_AddRecitatonToSurahStreamServer) error
	AddRecitatonToSurah(context.Context, *AddRecitatonToSurahRequest) (*DefaultResponse, error)
}

func RegisterQuranServiceServer(s *grpc.Server, srv QuranServiceServer) {
	s.RegisterService(&_QuranService_serviceDesc, srv)
}

func _QuranService_GetQuran_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuranServiceServer).GetQuran(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quran.QuranService/GetQuran",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuranServiceServer).GetQuran(ctx, req.(*EmptyPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuranService_GetAyah_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAyahRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuranServiceServer).GetAyah(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quran.QuranService/GetAyah",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuranServiceServer).GetAyah(ctx, req.(*GetAyahRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuranService_GetSurah_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSurahRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuranServiceServer).GetSurah(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quran.QuranService/GetSurah",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuranServiceServer).GetSurah(ctx, req.(*GetSurahRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuranService_GetChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuranServiceServer).GetChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quran.QuranService/GetChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuranServiceServer).GetChapter(ctx, req.(*GetChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuranService_AddTafseer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTafseerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuranServiceServer).AddTafseer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quran.QuranService/AddTafseer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuranServiceServer).AddTafseer(ctx, req.(*AddTafseerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuranService_AddRecitatonToSurahStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QuranServiceServer).AddRecitatonToSurahStream(&quranServiceAddRecitatonToSurahStreamServer{stream})
}

type QuranService_AddRecitatonToSurahStreamServer interface {
	Send(*DefaultResponse) error
	Recv() (*AddRecitatonToAyahRequest, error)
	grpc.ServerStream
}

type quranServiceAddRecitatonToSurahStreamServer struct {
	grpc.ServerStream
}

func (x *quranServiceAddRecitatonToSurahStreamServer) Send(m *DefaultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *quranServiceAddRecitatonToSurahStreamServer) Recv() (*AddRecitatonToAyahRequest, error) {
	m := new(AddRecitatonToAyahRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _QuranService_AddRecitatonToSurah_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRecitatonToSurahRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuranServiceServer).AddRecitatonToSurah(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quran.QuranService/AddRecitatonToSurah",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuranServiceServer).AddRecitatonToSurah(ctx, req.(*AddRecitatonToSurahRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuranService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quran.QuranService",
	HandlerType: (*QuranServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuran",
			Handler:    _QuranService_GetQuran_Handler,
		},
		{
			MethodName: "GetAyah",
			Handler:    _QuranService_GetAyah_Handler,
		},
		{
			MethodName: "GetSurah",
			Handler:    _QuranService_GetSurah_Handler,
		},
		{
			MethodName: "GetChapter",
			Handler:    _QuranService_GetChapter_Handler,
		},
		{
			MethodName: "AddTafseer",
			Handler:    _QuranService_AddTafseer_Handler,
		},
		{
			MethodName: "AddRecitatonToSurah",
			Handler:    _QuranService_AddRecitatonToSurah_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddRecitatonToSurahStream",
			Handler:       _QuranService_AddRecitatonToSurahStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "quran/quran.proto",
}
