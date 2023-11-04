// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: video.proto

// protoc -I=. -I=$GOPATH/pkg/mod --go-grpc_out=. --gogo_out=. *.proto

package videoApi

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UserInfo struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,json=user_id,proto3" json:"user_id"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,json=user_name,proto3" json:"user_name"`
	FollowCount          int64    `protobuf:"varint,3,opt,name=FollowCount,json=follow_count,proto3" json:"follow_count"`
	FollowerCount        int64    `protobuf:"varint,4,opt,name=FollowerCount,json=follower_count,proto3" json:"follower_count"`
	IsFollow             int64    `protobuf:"varint,5,opt,name=IsFollow,json=is_follow,proto3" json:"is_follow"`
	Avatar               string   `protobuf:"bytes,6,opt,name=Avatar,json=avatar,proto3" json:"avatar"`
	PublishNum           int64    `protobuf:"varint,7,opt,name=PublishNum,json=publish_num,proto3" json:"publish_num"`
	FavouriteNum         int64    `protobuf:"varint,8,opt,name=FavouriteNum,json=favourite_num,proto3" json:"favourite_num"`
	LikeNum              int64    `protobuf:"varint,9,opt,name=LikeNum,json=like_num,proto3" json:"like_num"`
	ReceivedLikeNum      int64    `protobuf:"varint,10,opt,name=ReceivedLikeNum,json=received_like_num,proto3" json:"received_like_num"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{0}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetFollowCount() int64 {
	if m != nil {
		return m.FollowCount
	}
	return 0
}

func (m *UserInfo) GetFollowerCount() int64 {
	if m != nil {
		return m.FollowerCount
	}
	return 0
}

func (m *UserInfo) GetIsFollow() int64 {
	if m != nil {
		return m.IsFollow
	}
	return 0
}

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserInfo) GetPublishNum() int64 {
	if m != nil {
		return m.PublishNum
	}
	return 0
}

func (m *UserInfo) GetFavouriteNum() int64 {
	if m != nil {
		return m.FavouriteNum
	}
	return 0
}

func (m *UserInfo) GetLikeNum() int64 {
	if m != nil {
		return m.LikeNum
	}
	return 0
}

func (m *UserInfo) GetReceivedLikeNum() int64 {
	if m != nil {
		return m.ReceivedLikeNum
	}
	return 0
}

type VideoInfo struct {
	VideoID              string    `protobuf:"bytes,1,opt,name=VideoID,json=video_id,proto3" json:"video_id"`
	Author               *UserInfo `protobuf:"bytes,2,opt,name=Author,json=author,proto3" json:"author"`
	PlayUrl              string    `protobuf:"bytes,3,opt,name=PlayUrl,json=play_url,proto3" json:"play_url"`
	CoverUrl             string    `protobuf:"bytes,4,opt,name=CoverUrl,json=cover_url,proto3" json:"cover_url"`
	FavoriteCount        int64     `protobuf:"varint,5,opt,name=FavoriteCount,json=favorite_count,proto3" json:"favorite_count"`
	CommentCount         int64     `protobuf:"varint,6,opt,name=CommentCount,json=comment_count,proto3" json:"comment_count"`
	IsFavorite           int64     `protobuf:"varint,7,opt,name=IsFavorite,json=is_favorite,proto3" json:"is_favorite"`
	Title                string    `protobuf:"bytes,8,opt,name=Title,json=title,proto3" json:"title"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VideoInfo) Reset()         { *m = VideoInfo{} }
func (m *VideoInfo) String() string { return proto.CompactTextString(m) }
func (*VideoInfo) ProtoMessage()    {}
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{1}
}
func (m *VideoInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoInfo.Unmarshal(m, b)
}
func (m *VideoInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoInfo.Marshal(b, m, deterministic)
}
func (m *VideoInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoInfo.Merge(m, src)
}
func (m *VideoInfo) XXX_Size() int {
	return xxx_messageInfo_VideoInfo.Size(m)
}
func (m *VideoInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VideoInfo proto.InternalMessageInfo

func (m *VideoInfo) GetVideoID() string {
	if m != nil {
		return m.VideoID
	}
	return ""
}

func (m *VideoInfo) GetAuthor() *UserInfo {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *VideoInfo) GetPlayUrl() string {
	if m != nil {
		return m.PlayUrl
	}
	return ""
}

func (m *VideoInfo) GetCoverUrl() string {
	if m != nil {
		return m.CoverUrl
	}
	return ""
}

func (m *VideoInfo) GetFavoriteCount() int64 {
	if m != nil {
		return m.FavoriteCount
	}
	return 0
}

func (m *VideoInfo) GetCommentCount() int64 {
	if m != nil {
		return m.CommentCount
	}
	return 0
}

func (m *VideoInfo) GetIsFavorite() int64 {
	if m != nil {
		return m.IsFavorite
	}
	return 0
}

func (m *VideoInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type VideoFeedReq struct {
	LatestTime           int64    `protobuf:"varint,1,opt,name=LatestTime,json=latest_time,proto3" json:"latest_time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoFeedReq) Reset()         { *m = VideoFeedReq{} }
func (m *VideoFeedReq) String() string { return proto.CompactTextString(m) }
func (*VideoFeedReq) ProtoMessage()    {}
func (*VideoFeedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{2}
}
func (m *VideoFeedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoFeedReq.Unmarshal(m, b)
}
func (m *VideoFeedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoFeedReq.Marshal(b, m, deterministic)
}
func (m *VideoFeedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoFeedReq.Merge(m, src)
}
func (m *VideoFeedReq) XXX_Size() int {
	return xxx_messageInfo_VideoFeedReq.Size(m)
}
func (m *VideoFeedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoFeedReq.DiscardUnknown(m)
}

var xxx_messageInfo_VideoFeedReq proto.InternalMessageInfo

func (m *VideoFeedReq) GetLatestTime() int64 {
	if m != nil {
		return m.LatestTime
	}
	return 0
}

type VideoFeedResp struct {
	VideoList            []*VideoInfo `protobuf:"bytes,1,rep,name=VideoList,json=video_info,proto3" json:"video_info"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VideoFeedResp) Reset()         { *m = VideoFeedResp{} }
func (m *VideoFeedResp) String() string { return proto.CompactTextString(m) }
func (*VideoFeedResp) ProtoMessage()    {}
func (*VideoFeedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{3}
}
func (m *VideoFeedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoFeedResp.Unmarshal(m, b)
}
func (m *VideoFeedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoFeedResp.Marshal(b, m, deterministic)
}
func (m *VideoFeedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoFeedResp.Merge(m, src)
}
func (m *VideoFeedResp) XXX_Size() int {
	return xxx_messageInfo_VideoFeedResp.Size(m)
}
func (m *VideoFeedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoFeedResp.DiscardUnknown(m)
}

var xxx_messageInfo_VideoFeedResp proto.InternalMessageInfo

func (m *VideoFeedResp) GetVideoList() []*VideoInfo {
	if m != nil {
		return m.VideoList
	}
	return nil
}

type UploadVideoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadVideoReq) Reset()         { *m = UploadVideoReq{} }
func (m *UploadVideoReq) String() string { return proto.CompactTextString(m) }
func (*UploadVideoReq) ProtoMessage()    {}
func (*UploadVideoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{4}
}
func (m *UploadVideoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadVideoReq.Unmarshal(m, b)
}
func (m *UploadVideoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadVideoReq.Marshal(b, m, deterministic)
}
func (m *UploadVideoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadVideoReq.Merge(m, src)
}
func (m *UploadVideoReq) XXX_Size() int {
	return xxx_messageInfo_UploadVideoReq.Size(m)
}
func (m *UploadVideoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadVideoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadVideoReq proto.InternalMessageInfo

type UploadVideoResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,json=token,proto3" json:"token"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadVideoResp) Reset()         { *m = UploadVideoResp{} }
func (m *UploadVideoResp) String() string { return proto.CompactTextString(m) }
func (*UploadVideoResp) ProtoMessage()    {}
func (*UploadVideoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{5}
}
func (m *UploadVideoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadVideoResp.Unmarshal(m, b)
}
func (m *UploadVideoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadVideoResp.Marshal(b, m, deterministic)
}
func (m *UploadVideoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadVideoResp.Merge(m, src)
}
func (m *UploadVideoResp) XXX_Size() int {
	return xxx_messageInfo_UploadVideoResp.Size(m)
}
func (m *UploadVideoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadVideoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UploadVideoResp proto.InternalMessageInfo

func (m *UploadVideoResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DownloadReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadReq) Reset()         { *m = DownloadReq{} }
func (m *DownloadReq) String() string { return proto.CompactTextString(m) }
func (*DownloadReq) ProtoMessage()    {}
func (*DownloadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{6}
}
func (m *DownloadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadReq.Unmarshal(m, b)
}
func (m *DownloadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadReq.Marshal(b, m, deterministic)
}
func (m *DownloadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadReq.Merge(m, src)
}
func (m *DownloadReq) XXX_Size() int {
	return xxx_messageInfo_DownloadReq.Size(m)
}
func (m *DownloadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadReq.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadReq proto.InternalMessageInfo

type DownLoadResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,json=token,proto3" json:"token"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownLoadResp) Reset()         { *m = DownLoadResp{} }
func (m *DownLoadResp) String() string { return proto.CompactTextString(m) }
func (*DownLoadResp) ProtoMessage()    {}
func (*DownLoadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{7}
}
func (m *DownLoadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownLoadResp.Unmarshal(m, b)
}
func (m *DownLoadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownLoadResp.Marshal(b, m, deterministic)
}
func (m *DownLoadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownLoadResp.Merge(m, src)
}
func (m *DownLoadResp) XXX_Size() int {
	return xxx_messageInfo_DownLoadResp.Size(m)
}
func (m *DownLoadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DownLoadResp.DiscardUnknown(m)
}

var xxx_messageInfo_DownLoadResp proto.InternalMessageInfo

func (m *DownLoadResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "video.UserInfo")
	proto.RegisterType((*VideoInfo)(nil), "video.VideoInfo")
	proto.RegisterType((*VideoFeedReq)(nil), "video.VideoFeedReq")
	proto.RegisterType((*VideoFeedResp)(nil), "video.VideoFeedResp")
	proto.RegisterType((*UploadVideoReq)(nil), "video.UploadVideoReq")
	proto.RegisterType((*UploadVideoResp)(nil), "video.UploadVideoResp")
	proto.RegisterType((*DownloadReq)(nil), "video.DownloadReq")
	proto.RegisterType((*DownLoadResp)(nil), "video.DownLoadResp")
}

func init() { proto.RegisterFile("video.proto", fileDescriptor_0ad4ea8866efb1e3) }

var fileDescriptor_0ad4ea8866efb1e3 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x25, 0x4b, 0xe3, 0x34, 0x72, 0xd2, 0xa4, 0xea, 0x3a, 0x4c, 0x5e, 0x5c, 0xcc, 0xa0, 0x65,
	0xb0, 0x64, 0xa4, 0x50, 0x18, 0x63, 0x6c, 0x49, 0x4b, 0xa1, 0x10, 0x4a, 0xd1, 0xda, 0x3d, 0xec,
	0xc5, 0x38, 0x89, 0xd2, 0x8a, 0xd8, 0x96, 0x6b, 0xcb, 0x2e, 0x7d, 0xdf, 0xc7, 0xed, 0x2b, 0xfc,
	0x01, 0x7e, 0xdc, 0x17, 0x0c, 0x5d, 0xc9, 0x89, 0xb3, 0xee, 0x61, 0x4f, 0xbe, 0xe7, 0xe8, 0x1e,
	0x49, 0x57, 0xe7, 0x60, 0x64, 0x66, 0x6c, 0x41, 0xf9, 0x20, 0x8a, 0xb9, 0xe0, 0xb8, 0x01, 0xa0,
	0xff, 0xfe, 0x9e, 0x89, 0x87, 0x74, 0x36, 0x98, 0xf3, 0x60, 0x78, 0xcf, 0xef, 0xf9, 0x10, 0x56,
	0x67, 0xe9, 0x12, 0x10, 0x00, 0xa8, 0x94, 0xca, 0xf9, 0x5d, 0x47, 0xbb, 0x77, 0x09, 0x8d, 0xaf,
	0xc2, 0x25, 0xc7, 0x6f, 0x91, 0x01, 0xf5, 0xc2, 0xaa, 0x1d, 0xd5, 0x4e, 0x5a, 0x13, 0xb3, 0xc8,
	0xed, 0x66, 0x9a, 0xd0, 0xd8, 0x65, 0x0b, 0x52, 0x16, 0xf8, 0x9d, 0x52, 0x5c, 0x7b, 0x01, 0xb5,
	0x5e, 0x41, 0x5f, 0xa7, 0xc8, 0xed, 0x16, 0x2c, 0x87, 0x5e, 0x40, 0xc9, 0xa6, 0xc4, 0xa7, 0xc8,
	0xbc, 0xe4, 0xbe, 0xcf, 0x9f, 0xce, 0x79, 0x1a, 0x0a, 0xab, 0x7e, 0x54, 0x3b, 0xa9, 0x4f, 0x7a,
	0x45, 0x6e, 0xb7, 0x97, 0x40, 0xbb, 0x73, 0xc9, 0x93, 0x2d, 0x84, 0x3f, 0xa2, 0x8e, 0x12, 0xd1,
	0x58, 0xc9, 0x76, 0x40, 0x86, 0x8b, 0xdc, 0xde, 0x5b, 0xea, 0x05, 0x2d, 0xfc, 0x0b, 0xcb, 0xbb,
	0x5d, 0x25, 0x4a, 0x6c, 0x35, 0x40, 0x05, 0x77, 0x63, 0x89, 0xab, 0x1a, 0xc9, 0xa6, 0xc4, 0x0e,
	0x32, 0xc6, 0x99, 0x27, 0xbc, 0xd8, 0x32, 0x60, 0x0a, 0x54, 0xe4, 0xb6, 0xe1, 0x01, 0x43, 0xf4,
	0x17, 0x7f, 0x40, 0xe8, 0x26, 0x9d, 0xf9, 0x2c, 0x79, 0xb8, 0x4e, 0x03, 0xab, 0x09, 0x3b, 0x76,
	0x8b, 0xdc, 0x36, 0x23, 0xc5, 0xba, 0x61, 0x1a, 0x90, 0x2a, 0xc0, 0x67, 0xa8, 0x7d, 0xe9, 0x65,
	0x3c, 0x8d, 0x99, 0xa0, 0x52, 0xb3, 0x0b, 0x9a, 0xfd, 0x22, 0xb7, 0x3b, 0xcb, 0x92, 0x07, 0xd5,
	0x36, 0xc4, 0xc7, 0xa8, 0x39, 0x65, 0x2b, 0x90, 0xb4, 0x40, 0xd2, 0x2e, 0x72, 0x7b, 0xd7, 0x67,
	0x2b, 0xd5, 0xbd, 0xae, 0xf0, 0x18, 0x75, 0x09, 0x9d, 0x53, 0x96, 0xd1, 0x45, 0x29, 0x40, 0x20,
	0x38, 0x2c, 0x72, 0x7b, 0x3f, 0xd6, 0x4b, 0xee, 0x5a, 0xf9, 0x92, 0x72, 0x7e, 0xd6, 0x51, 0xeb,
	0xbb, 0x4c, 0x0b, 0xb8, 0x7e, 0x8c, 0x9a, 0x0a, 0x5c, 0x68, 0xdb, 0xe1, 0x64, 0x48, 0x93, 0xf4,
	0x7d, 0x5d, 0xe1, 0x53, 0x64, 0x8c, 0x53, 0xf1, 0xc0, 0x63, 0xb0, 0xdd, 0x1c, 0x75, 0x07, 0x2a,
	0x7f, 0x65, 0x7e, 0xf4, 0x0b, 0x42, 0x0b, 0xd1, 0x5f, 0xb9, 0xfb, 0x8d, 0xef, 0x3d, 0xdf, 0xc5,
	0x3e, 0xb8, 0xaf, 0x77, 0x8f, 0x7c, 0xef, 0xd9, 0x4d, 0x63, 0x9f, 0xac, 0x2b, 0x69, 0xdd, 0x39,
	0xcf, 0x68, 0x2c, 0x3b, 0x77, 0x36, 0xb1, 0x9a, 0x4b, 0x0e, 0x5a, 0x37, 0x25, 0x24, 0xc4, 0xcb,
	0xb8, 0x7c, 0x3c, 0x95, 0x90, 0x46, 0x25, 0x21, 0x7a, 0x61, 0x9d, 0x90, 0x2d, 0x2c, 0xfd, 0x39,
	0xe7, 0x41, 0x40, 0x43, 0xa1, 0x94, 0xc6, 0xc6, 0x9f, 0xb9, 0xe2, 0xb5, 0x70, 0x1b, 0xca, 0x24,
	0x5c, 0x25, 0xe5, 0xa1, 0xd5, 0x24, 0xc8, 0x40, 0x69, 0x9a, 0x54, 0x01, 0xb6, 0x51, 0xe3, 0x96,
	0x09, 0x9f, 0x42, 0x04, 0x5a, 0x93, 0x56, 0x91, 0xdb, 0x0d, 0x21, 0x09, 0xa2, 0x3e, 0xce, 0x57,
	0xd4, 0x86, 0x87, 0xbf, 0xa4, 0x74, 0x41, 0xe8, 0xa3, 0x3c, 0x62, 0xea, 0x09, 0x9a, 0x88, 0x5b,
	0x16, 0x50, 0xf0, 0x42, 0x1f, 0xe1, 0x03, 0xeb, 0x0a, 0x16, 0x50, 0x52, 0x05, 0xce, 0x0d, 0xea,
	0x54, 0x76, 0x48, 0x22, 0xfc, 0x45, 0x1b, 0x3b, 0x65, 0x89, 0xb0, 0x6a, 0x47, 0xf5, 0x13, 0x73,
	0xd4, 0xd3, 0x2e, 0xad, 0x0d, 0x9f, 0xec, 0x15, 0xb9, 0x8d, 0xb4, 0xab, 0xe1, 0x92, 0x93, 0x4a,
	0xed, 0xf4, 0xd0, 0xde, 0x5d, 0xe4, 0x73, 0x6f, 0x01, 0xed, 0x84, 0x3e, 0x3a, 0x23, 0xd4, 0xdd,
	0x62, 0x92, 0x08, 0x26, 0xe3, 0x2b, 0x1a, 0xea, 0xbc, 0xa8, 0xc9, 0x24, 0x41, 0xd4, 0xc7, 0xe9,
	0x20, 0xf3, 0x82, 0x3f, 0x85, 0x52, 0x25, 0xb7, 0x18, 0xa2, 0xb6, 0x84, 0x53, 0x80, 0xff, 0xa1,
	0x1f, 0xfd, 0xaa, 0xe9, 0xa7, 0xf9, 0x46, 0xe3, 0x8c, 0xcd, 0x29, 0x3e, 0xd3, 0x73, 0xc9, 0x41,
	0xf1, 0x41, 0x75, 0x22, 0xfd, 0x78, 0xfd, 0xd7, 0x2f, 0xc9, 0x24, 0xc2, 0x63, 0xd4, 0xab, 0x5c,
	0x1e, 0x0e, 0xc5, 0x87, 0x65, 0x6c, 0xb7, 0xe6, 0xec, 0xbf, 0xf9, 0x17, 0x9d, 0x44, 0xf8, 0x33,
	0xc2, 0xe5, 0xe5, 0x2b, 0x9b, 0x60, 0xdd, 0x5d, 0x19, 0xb3, 0x7f, 0x50, 0xe1, 0xca, 0x59, 0x27,
	0x9d, 0x1f, 0xe6, 0x60, 0xf8, 0x09, 0x16, 0xc6, 0x11, 0x9b, 0x19, 0xf0, 0xdb, 0x3d, 0xfd, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0x30, 0xb2, 0x89, 0xbb, 0x05, 0x00, 0x00,
}
