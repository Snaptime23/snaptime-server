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
	IsEncoding           bool      `protobuf:"varint,9,opt,name=IsEncoding,json=is_encoding,proto3" json:"is_encoding"`
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

func (m *VideoInfo) GetIsEncoding() bool {
	if m != nil {
		return m.IsEncoding
	}
	return false
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
	Description          string   `protobuf:"bytes,1,opt,name=Description,json=description,proto3" json:"description"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"title"`
	VideoTag             []string `protobuf:"bytes,3,rep,name=VideoTag,json=video_tag,proto3" json:"video_tag"`
	FileExtension        string   `protobuf:"bytes,4,opt,name=FileExtension,json=file_extension,proto3" json:"file_extension"`
	UserId               string   `protobuf:"bytes,5,opt,name=UserId,json=user_id,proto3" json:"user_id"`
	VideoId              string   `protobuf:"bytes,6,opt,name=VideoId,json=video_id,proto3" json:"video_id"`
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

func (m *UploadVideoReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UploadVideoReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UploadVideoReq) GetVideoTag() []string {
	if m != nil {
		return m.VideoTag
	}
	return nil
}

func (m *UploadVideoReq) GetFileExtension() string {
	if m != nil {
		return m.FileExtension
	}
	return ""
}

func (m *UploadVideoReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UploadVideoReq) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

type UploadVideoResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,json=token,proto3" json:"token"`
	ResourceKey          string   `protobuf:"bytes,2,opt,name=ResourceKey,json=resource_key,proto3" json:"resource_key"`
	VideoId              string   `protobuf:"bytes,3,opt,name=VideoId,json=video_id,proto3" json:"video_id"`
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

func (m *UploadVideoResp) GetResourceKey() string {
	if m != nil {
		return m.ResourceKey
	}
	return ""
}

func (m *UploadVideoResp) GetVideoId() string {
	if m != nil {
		return m.VideoId
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

type GetVideoInfoByIdReq struct {
	VideoId              string   `protobuf:"bytes,1,opt,name=VideoId,json=video_id,proto3" json:"video_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVideoInfoByIdReq) Reset()         { *m = GetVideoInfoByIdReq{} }
func (m *GetVideoInfoByIdReq) String() string { return proto.CompactTextString(m) }
func (*GetVideoInfoByIdReq) ProtoMessage()    {}
func (*GetVideoInfoByIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{8}
}
func (m *GetVideoInfoByIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoInfoByIdReq.Unmarshal(m, b)
}
func (m *GetVideoInfoByIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoInfoByIdReq.Marshal(b, m, deterministic)
}
func (m *GetVideoInfoByIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoInfoByIdReq.Merge(m, src)
}
func (m *GetVideoInfoByIdReq) XXX_Size() int {
	return xxx_messageInfo_GetVideoInfoByIdReq.Size(m)
}
func (m *GetVideoInfoByIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoInfoByIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoInfoByIdReq proto.InternalMessageInfo

func (m *GetVideoInfoByIdReq) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

type GetVideoInfoByIdResp struct {
	Video                *VideoInfo `protobuf:"bytes,1,opt,name=Video,json=video,proto3" json:"video"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetVideoInfoByIdResp) Reset()         { *m = GetVideoInfoByIdResp{} }
func (m *GetVideoInfoByIdResp) String() string { return proto.CompactTextString(m) }
func (*GetVideoInfoByIdResp) ProtoMessage()    {}
func (*GetVideoInfoByIdResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{9}
}
func (m *GetVideoInfoByIdResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoInfoByIdResp.Unmarshal(m, b)
}
func (m *GetVideoInfoByIdResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoInfoByIdResp.Marshal(b, m, deterministic)
}
func (m *GetVideoInfoByIdResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoInfoByIdResp.Merge(m, src)
}
func (m *GetVideoInfoByIdResp) XXX_Size() int {
	return xxx_messageInfo_GetVideoInfoByIdResp.Size(m)
}
func (m *GetVideoInfoByIdResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoInfoByIdResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoInfoByIdResp proto.InternalMessageInfo

func (m *GetVideoInfoByIdResp) GetVideo() *VideoInfo {
	if m != nil {
		return m.Video
	}
	return nil
}

type RebackOneReq struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,json=title,proto3" json:"title"`
	CoverUrl             string   `protobuf:"bytes,2,opt,name=CoverUrl,json=cover_url,proto3" json:"cover_url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RebackOneReq) Reset()         { *m = RebackOneReq{} }
func (m *RebackOneReq) String() string { return proto.CompactTextString(m) }
func (*RebackOneReq) ProtoMessage()    {}
func (*RebackOneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{10}
}
func (m *RebackOneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RebackOneReq.Unmarshal(m, b)
}
func (m *RebackOneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RebackOneReq.Marshal(b, m, deterministic)
}
func (m *RebackOneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RebackOneReq.Merge(m, src)
}
func (m *RebackOneReq) XXX_Size() int {
	return xxx_messageInfo_RebackOneReq.Size(m)
}
func (m *RebackOneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RebackOneReq.DiscardUnknown(m)
}

var xxx_messageInfo_RebackOneReq proto.InternalMessageInfo

func (m *RebackOneReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RebackOneReq) GetCoverUrl() string {
	if m != nil {
		return m.CoverUrl
	}
	return ""
}

type RebackOneResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RebackOneResp) Reset()         { *m = RebackOneResp{} }
func (m *RebackOneResp) String() string { return proto.CompactTextString(m) }
func (*RebackOneResp) ProtoMessage()    {}
func (*RebackOneResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{11}
}
func (m *RebackOneResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RebackOneResp.Unmarshal(m, b)
}
func (m *RebackOneResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RebackOneResp.Marshal(b, m, deterministic)
}
func (m *RebackOneResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RebackOneResp.Merge(m, src)
}
func (m *RebackOneResp) XXX_Size() int {
	return xxx_messageInfo_RebackOneResp.Size(m)
}
func (m *RebackOneResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RebackOneResp.DiscardUnknown(m)
}

var xxx_messageInfo_RebackOneResp proto.InternalMessageInfo

type RebackTwoReq struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	ResourceKey          []string `protobuf:"bytes,2,rep,name=ResourceKey,proto3" json:"ResourceKey,omitempty"`
	Type                 []string `protobuf:"bytes,3,rep,name=Type,proto3" json:"Type,omitempty"`
	VideoId              string   `protobuf:"bytes,4,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	Quality              []int64  `protobuf:"varint,5,rep,packed,name=Quality,proto3" json:"Quality,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RebackTwoReq) Reset()         { *m = RebackTwoReq{} }
func (m *RebackTwoReq) String() string { return proto.CompactTextString(m) }
func (*RebackTwoReq) ProtoMessage()    {}
func (*RebackTwoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{12}
}
func (m *RebackTwoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RebackTwoReq.Unmarshal(m, b)
}
func (m *RebackTwoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RebackTwoReq.Marshal(b, m, deterministic)
}
func (m *RebackTwoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RebackTwoReq.Merge(m, src)
}
func (m *RebackTwoReq) XXX_Size() int {
	return xxx_messageInfo_RebackTwoReq.Size(m)
}
func (m *RebackTwoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RebackTwoReq.DiscardUnknown(m)
}

var xxx_messageInfo_RebackTwoReq proto.InternalMessageInfo

func (m *RebackTwoReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RebackTwoReq) GetResourceKey() []string {
	if m != nil {
		return m.ResourceKey
	}
	return nil
}

func (m *RebackTwoReq) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RebackTwoReq) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

func (m *RebackTwoReq) GetQuality() []int64 {
	if m != nil {
		return m.Quality
	}
	return nil
}

type RebackTwoResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RebackTwoResp) Reset()         { *m = RebackTwoResp{} }
func (m *RebackTwoResp) String() string { return proto.CompactTextString(m) }
func (*RebackTwoResp) ProtoMessage()    {}
func (*RebackTwoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{13}
}
func (m *RebackTwoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RebackTwoResp.Unmarshal(m, b)
}
func (m *RebackTwoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RebackTwoResp.Marshal(b, m, deterministic)
}
func (m *RebackTwoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RebackTwoResp.Merge(m, src)
}
func (m *RebackTwoResp) XXX_Size() int {
	return xxx_messageInfo_RebackTwoResp.Size(m)
}
func (m *RebackTwoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RebackTwoResp.DiscardUnknown(m)
}

var xxx_messageInfo_RebackTwoResp proto.InternalMessageInfo

type PublishListReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,json=user_id,proto3" json:"user_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishListReq) Reset()         { *m = PublishListReq{} }
func (m *PublishListReq) String() string { return proto.CompactTextString(m) }
func (*PublishListReq) ProtoMessage()    {}
func (*PublishListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{14}
}
func (m *PublishListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishListReq.Unmarshal(m, b)
}
func (m *PublishListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishListReq.Marshal(b, m, deterministic)
}
func (m *PublishListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishListReq.Merge(m, src)
}
func (m *PublishListReq) XXX_Size() int {
	return xxx_messageInfo_PublishListReq.Size(m)
}
func (m *PublishListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishListReq.DiscardUnknown(m)
}

var xxx_messageInfo_PublishListReq proto.InternalMessageInfo

func (m *PublishListReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type PublishListResp struct {
	Video                []*VideoInfo `protobuf:"bytes,1,rep,name=Video,json=video,proto3" json:"video"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PublishListResp) Reset()         { *m = PublishListResp{} }
func (m *PublishListResp) String() string { return proto.CompactTextString(m) }
func (*PublishListResp) ProtoMessage()    {}
func (*PublishListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{15}
}
func (m *PublishListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishListResp.Unmarshal(m, b)
}
func (m *PublishListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishListResp.Marshal(b, m, deterministic)
}
func (m *PublishListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishListResp.Merge(m, src)
}
func (m *PublishListResp) XXX_Size() int {
	return xxx_messageInfo_PublishListResp.Size(m)
}
func (m *PublishListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishListResp.DiscardUnknown(m)
}

var xxx_messageInfo_PublishListResp proto.InternalMessageInfo

func (m *PublishListResp) GetVideo() []*VideoInfo {
	if m != nil {
		return m.Video
	}
	return nil
}

type SearchVideoByVideoTagReq struct {
	VideoTag             string   `protobuf:"bytes,1,opt,name=VideoTag,json=video_tag,proto3" json:"video_tag"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchVideoByVideoTagReq) Reset()         { *m = SearchVideoByVideoTagReq{} }
func (m *SearchVideoByVideoTagReq) String() string { return proto.CompactTextString(m) }
func (*SearchVideoByVideoTagReq) ProtoMessage()    {}
func (*SearchVideoByVideoTagReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{16}
}
func (m *SearchVideoByVideoTagReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchVideoByVideoTagReq.Unmarshal(m, b)
}
func (m *SearchVideoByVideoTagReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchVideoByVideoTagReq.Marshal(b, m, deterministic)
}
func (m *SearchVideoByVideoTagReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchVideoByVideoTagReq.Merge(m, src)
}
func (m *SearchVideoByVideoTagReq) XXX_Size() int {
	return xxx_messageInfo_SearchVideoByVideoTagReq.Size(m)
}
func (m *SearchVideoByVideoTagReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchVideoByVideoTagReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchVideoByVideoTagReq proto.InternalMessageInfo

func (m *SearchVideoByVideoTagReq) GetVideoTag() string {
	if m != nil {
		return m.VideoTag
	}
	return ""
}

type SearchVideoByVideoTagResp struct {
	Video                []*VideoInfo `protobuf:"bytes,1,rep,name=Video,json=video_info,proto3" json:"video_info"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchVideoByVideoTagResp) Reset()         { *m = SearchVideoByVideoTagResp{} }
func (m *SearchVideoByVideoTagResp) String() string { return proto.CompactTextString(m) }
func (*SearchVideoByVideoTagResp) ProtoMessage()    {}
func (*SearchVideoByVideoTagResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{17}
}
func (m *SearchVideoByVideoTagResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchVideoByVideoTagResp.Unmarshal(m, b)
}
func (m *SearchVideoByVideoTagResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchVideoByVideoTagResp.Marshal(b, m, deterministic)
}
func (m *SearchVideoByVideoTagResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchVideoByVideoTagResp.Merge(m, src)
}
func (m *SearchVideoByVideoTagResp) XXX_Size() int {
	return xxx_messageInfo_SearchVideoByVideoTagResp.Size(m)
}
func (m *SearchVideoByVideoTagResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchVideoByVideoTagResp.DiscardUnknown(m)
}

var xxx_messageInfo_SearchVideoByVideoTagResp proto.InternalMessageInfo

func (m *SearchVideoByVideoTagResp) GetVideo() []*VideoInfo {
	if m != nil {
		return m.Video
	}
	return nil
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
	proto.RegisterType((*GetVideoInfoByIdReq)(nil), "video.GetVideoInfoByIdReq")
	proto.RegisterType((*GetVideoInfoByIdResp)(nil), "video.GetVideoInfoByIdResp")
	proto.RegisterType((*RebackOneReq)(nil), "video.RebackOneReq")
	proto.RegisterType((*RebackOneResp)(nil), "video.RebackOneResp")
	proto.RegisterType((*RebackTwoReq)(nil), "video.RebackTwoReq")
	proto.RegisterType((*RebackTwoResp)(nil), "video.RebackTwoResp")
	proto.RegisterType((*PublishListReq)(nil), "video.PublishListReq")
	proto.RegisterType((*PublishListResp)(nil), "video.PublishListResp")
	proto.RegisterType((*SearchVideoByVideoTagReq)(nil), "video.SearchVideoByVideoTagReq")
	proto.RegisterType((*SearchVideoByVideoTagResp)(nil), "video.SearchVideoByVideoTagResp")
}

func init() { proto.RegisterFile("video.proto", fileDescriptor_0ad4ea8866efb1e3) }

var fileDescriptor_0ad4ea8866efb1e3 = []byte{
	// 1114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdf, 0x4e, 0xe3, 0xc6,
	0x17, 0x56, 0x30, 0x09, 0x64, 0x9c, 0x10, 0x76, 0x80, 0x9f, 0xfc, 0x4b, 0x2f, 0x1c, 0x59, 0x95,
	0x16, 0xad, 0x54, 0xe8, 0x82, 0x84, 0x5a, 0x6d, 0xd5, 0x96, 0xc0, 0x52, 0xa1, 0x45, 0x5b, 0x3a,
	0x1b, 0xaa, 0x6a, 0x7b, 0x11, 0x19, 0x67, 0x12, 0x2c, 0x1c, 0x8f, 0xd7, 0x7f, 0x42, 0xf3, 0x12,
	0x55, 0x6f, 0xdb, 0x07, 0xe9, 0xe3, 0xf8, 0x01, 0x7c, 0xd9, 0x27, 0xa8, 0xe6, 0xcc, 0xd8, 0x1e,
	0x07, 0x92, 0x72, 0xe5, 0x73, 0xbe, 0x39, 0xdf, 0xf8, 0xf8, 0x9c, 0xef, 0x9c, 0x04, 0xe9, 0x33,
	0x77, 0x44, 0xd9, 0x41, 0x10, 0xb2, 0x98, 0xe1, 0x3a, 0x38, 0xdd, 0x2f, 0x26, 0x6e, 0x7c, 0x97,
	0xdc, 0x1e, 0x38, 0x6c, 0x7a, 0x38, 0x61, 0x13, 0x76, 0x08, 0xa7, 0xb7, 0xc9, 0x18, 0x3c, 0x70,
	0xc0, 0x12, 0x2c, 0xeb, 0x1f, 0x0d, 0x6d, 0xde, 0x44, 0x34, 0xbc, 0xf4, 0xc7, 0x0c, 0x7f, 0x8e,
	0x1a, 0x60, 0x8f, 0x8c, 0x5a, 0xaf, 0xb6, 0xdf, 0xec, 0xeb, 0x59, 0x6a, 0x6e, 0x24, 0x11, 0x0d,
	0x87, 0xee, 0x88, 0xe4, 0x06, 0x7e, 0x25, 0x18, 0xef, 0xed, 0x29, 0x35, 0xd6, 0x20, 0xae, 0x9d,
	0xa5, 0x66, 0x13, 0x8e, 0x7d, 0x7b, 0x4a, 0x49, 0x69, 0xe2, 0x63, 0xa4, 0x5f, 0x30, 0xcf, 0x63,
	0x0f, 0x67, 0x2c, 0xf1, 0x63, 0x43, 0xeb, 0xd5, 0xf6, 0xb5, 0xfe, 0x76, 0x96, 0x9a, 0xad, 0x31,
	0xc0, 0x43, 0x87, 0xe3, 0xa4, 0xe2, 0xe1, 0xaf, 0x51, 0x5b, 0x90, 0x68, 0x28, 0x68, 0xeb, 0x40,
	0xc3, 0x59, 0x6a, 0x6e, 0x8d, 0xe5, 0x81, 0x24, 0x2e, 0xf8, 0x3c, 0xb7, 0xcb, 0x48, 0x90, 0x8d,
	0x3a, 0xb0, 0x20, 0x37, 0x37, 0x1a, 0x8a, 0x40, 0x52, 0x9a, 0xd8, 0x42, 0x8d, 0xd3, 0x99, 0x1d,
	0xdb, 0xa1, 0xd1, 0x80, 0xaf, 0x40, 0x59, 0x6a, 0x36, 0x6c, 0x40, 0x88, 0x7c, 0xe2, 0x2f, 0x11,
	0xba, 0x4e, 0x6e, 0x3d, 0x37, 0xba, 0x7b, 0x9f, 0x4c, 0x8d, 0x0d, 0xb8, 0xb1, 0x93, 0xa5, 0xa6,
	0x1e, 0x08, 0x74, 0xe8, 0x27, 0x53, 0xa2, 0x3a, 0xf8, 0x04, 0xb5, 0x2e, 0xec, 0x19, 0x4b, 0x42,
	0x37, 0xa6, 0x9c, 0xb3, 0x09, 0x9c, 0x17, 0x59, 0x6a, 0xb6, 0xc7, 0x39, 0x0e, 0xac, 0xaa, 0x8b,
	0x5f, 0xa2, 0x8d, 0x2b, 0xf7, 0x1e, 0x28, 0x4d, 0xa0, 0xb4, 0xb2, 0xd4, 0xdc, 0xf4, 0xdc, 0x7b,
	0x11, 0x5d, 0x58, 0xf8, 0x14, 0x75, 0x08, 0x75, 0xa8, 0x3b, 0xa3, 0xa3, 0x9c, 0x80, 0x80, 0xb0,
	0x97, 0xa5, 0xe6, 0x8b, 0x50, 0x1e, 0x0d, 0x0b, 0xe6, 0x63, 0xc8, 0xfa, 0x5b, 0x43, 0xcd, 0x9f,
	0xb9, 0x5a, 0xa0, 0xeb, 0x2f, 0xd1, 0x86, 0x70, 0xce, 0x65, 0xdb, 0xe1, 0xcd, 0xa0, 0x26, 0xde,
	0xf7, 0xc2, 0xc2, 0xc7, 0xa8, 0x71, 0x9a, 0xc4, 0x77, 0x2c, 0x84, 0xb6, 0xeb, 0x47, 0x9d, 0x03,
	0xa1, 0xbf, 0x5c, 0x3f, 0xb2, 0x82, 0x10, 0x42, 0xe4, 0x93, 0xdf, 0x7e, 0xed, 0xd9, 0xf3, 0x9b,
	0xd0, 0x83, 0xee, 0xcb, 0xdb, 0x03, 0xcf, 0x9e, 0x0f, 0x93, 0xd0, 0x23, 0x85, 0xc5, 0x5b, 0x77,
	0xc6, 0x66, 0x34, 0xe4, 0x91, 0xeb, 0xa5, 0xac, 0x1c, 0x8e, 0x41, 0x68, 0x69, 0x82, 0x42, 0xec,
	0x19, 0xe3, 0xc5, 0x13, 0x0a, 0xa9, 0x2b, 0x0a, 0x91, 0x07, 0x85, 0x42, 0x2a, 0x3e, 0xef, 0xcf,
	0x19, 0x9b, 0x4e, 0xa9, 0x1f, 0x0b, 0x66, 0xa3, 0xec, 0x8f, 0x23, 0x70, 0x49, 0xac, 0xba, 0x5c,
	0x09, 0x97, 0x51, 0xfe, 0x52, 0x55, 0x09, 0x5c, 0x50, 0x12, 0x26, 0xaa, 0x83, 0x4d, 0x54, 0x1f,
	0xb8, 0xb1, 0x47, 0x41, 0x02, 0xcd, 0x7e, 0x33, 0x4b, 0xcd, 0x7a, 0xcc, 0x01, 0x22, 0x1e, 0xe2,
	0xca, 0xb7, 0xbe, 0xc3, 0x46, 0xae, 0x3f, 0x81, 0xae, 0x6f, 0x16, 0x57, 0x52, 0x09, 0x13, 0xd5,
	0xb1, 0xbe, 0x47, 0x2d, 0x68, 0xd5, 0x05, 0xa5, 0x23, 0x42, 0x3f, 0xf1, 0x1b, 0xae, 0xec, 0x98,
	0x46, 0xf1, 0xc0, 0x9d, 0x52, 0xe8, 0x9e, 0x4c, 0xca, 0x03, 0x74, 0x18, 0xbb, 0x53, 0x4a, 0x54,
	0xc7, 0xba, 0x46, 0x6d, 0xe5, 0x86, 0x28, 0xc0, 0xdf, 0x49, 0x29, 0x5c, 0xb9, 0x51, 0x6c, 0xd4,
	0x7a, 0xda, 0xbe, 0x7e, 0xb4, 0x2d, 0xfb, 0x5a, 0x48, 0xa4, 0xbf, 0x95, 0xa5, 0x26, 0x92, 0x3a,
	0xf0, 0xc7, 0x8c, 0x28, 0xb6, 0xf5, 0xd7, 0x1a, 0xda, 0xba, 0x09, 0x3c, 0x66, 0x8f, 0x20, 0x9e,
	0xa7, 0xf5, 0x1a, 0xe9, 0xe7, 0x34, 0x72, 0x42, 0x37, 0x88, 0x5d, 0xe6, 0x4b, 0x55, 0x41, 0x5e,
	0xa3, 0x12, 0x26, 0xaa, 0x53, 0x16, 0x6b, 0x6d, 0x49, 0xb1, 0x5e, 0xa1, 0x4d, 0xb8, 0x7f, 0x60,
	0x4f, 0x0c, 0xad, 0xa7, 0xe5, 0xf2, 0x10, 0x89, 0xc4, 0xf6, 0x84, 0x94, 0x26, 0xc8, 0xc3, 0xf5,
	0xe8, 0xdb, 0xdf, 0x62, 0xea, 0x47, 0x3c, 0x03, 0xa1, 0x27, 0x21, 0x0f, 0xd7, 0xa3, 0x43, 0x9a,
	0x9f, 0x90, 0x05, 0x5f, 0x59, 0x81, 0xf5, 0x15, 0x2b, 0xb0, 0x18, 0x99, 0x91, 0xdc, 0x1d, 0x4b,
	0x46, 0xc6, 0xfa, 0xa3, 0x86, 0x3a, 0x95, 0xe2, 0x44, 0x01, 0x7c, 0x2a, 0xbb, 0xa7, 0x79, 0x5d,
	0xc4, 0xa7, 0x72, 0x80, 0x88, 0x07, 0x5f, 0x9a, 0x84, 0x46, 0x2c, 0x09, 0x1d, 0xfa, 0x8e, 0xce,
	0x65, 0x45, 0x60, 0x69, 0x86, 0x12, 0x1e, 0xde, 0xd3, 0x39, 0xa9, 0x78, 0x6a, 0x4a, 0xda, 0xca,
	0x94, 0xda, 0x48, 0x3f, 0x67, 0x0f, 0x3e, 0xcf, 0x89, 0xd0, 0x4f, 0xd6, 0x21, 0x6a, 0x71, 0xf7,
	0x0a, 0xdc, 0x67, 0x64, 0x67, 0x7d, 0x8b, 0x76, 0x7e, 0xa0, 0x71, 0xa9, 0x8d, 0xf9, 0x25, 0x48,
	0x51, 0x79, 0xff, 0xca, 0x2d, 0x62, 0x5d, 0xa2, 0xdd, 0xc7, 0xfc, 0x28, 0xc0, 0xaf, 0x51, 0x1d,
	0x40, 0xa0, 0x3f, 0x25, 0x42, 0x48, 0x05, 0x40, 0x22, 0x1e, 0xd6, 0xaf, 0xa8, 0x45, 0xe8, 0xad,
	0xed, 0xdc, 0xff, 0xe8, 0x53, 0x9e, 0x43, 0x21, 0xa2, 0xda, 0x72, 0x11, 0x15, 0x3b, 0x66, 0x6d,
	0xf5, 0x8e, 0xb1, 0x3a, 0xa8, 0xad, 0x5c, 0x1e, 0x05, 0xd6, 0xef, 0xb5, 0xfc, 0x75, 0x83, 0x07,
	0x90, 0xf9, 0x6e, 0xe5, 0x75, 0x44, 0x38, 0xb8, 0xb7, 0xd8, 0x3d, 0x6d, 0xbf, 0x49, 0x54, 0x08,
	0x63, 0xb4, 0x3e, 0x98, 0x07, 0x54, 0xc8, 0x98, 0x80, 0x8d, 0x8d, 0xb2, 0x7c, 0x20, 0x56, 0x92,
	0xbb, 0xfc, 0xe4, 0xa7, 0xc4, 0xf6, 0xdc, 0x78, 0x6e, 0xd4, 0x7b, 0xda, 0xbe, 0x46, 0x72, 0xb7,
	0xcc, 0x10, 0xf2, 0x89, 0x02, 0xeb, 0x04, 0x6d, 0xc9, 0x5f, 0x2b, 0x3e, 0xcd, 0x3c, 0xc5, 0x67,
	0xfd, 0xa2, 0x5b, 0xe7, 0xa8, 0x53, 0xe1, 0x55, 0xbb, 0xa1, 0x3d, 0xb3, 0x1b, 0x17, 0xc8, 0xf8,
	0x40, 0xed, 0xd0, 0xb9, 0x83, 0xa0, 0xfe, 0x3c, 0x1f, 0x57, 0x9e, 0x87, 0x3a, 0xbd, 0xb5, 0xb2,
	0xf0, 0x4f, 0x4d, 0xaf, 0xf5, 0x0b, 0xfa, 0xff, 0x92, 0x7b, 0xa2, 0x00, 0xbf, 0xf9, 0xaf, 0xbc,
	0x56, 0xac, 0xaa, 0xa3, 0x3f, 0xd7, 0xe5, 0xfe, 0xfc, 0x40, 0xc3, 0x99, 0xeb, 0x50, 0x7c, 0x22,
	0x97, 0x1f, 0xdf, 0x86, 0x78, 0x47, 0xbd, 0x4b, 0x6e, 0xd8, 0xee, 0xee, 0x63, 0x30, 0x0a, 0xf0,
	0x37, 0x48, 0x57, 0xa6, 0x1a, 0xef, 0xe5, 0x3f, 0x84, 0x95, 0x35, 0xd8, 0xfd, 0xdf, 0x53, 0x70,
	0x14, 0xe0, 0xaf, 0x50, 0x3b, 0x1f, 0x39, 0xc1, 0xc7, 0x32, 0x50, 0x99, 0xcb, 0xee, 0x8e, 0x82,
	0x15, 0xc3, 0xf9, 0x0e, 0x6d, 0x2f, 0xce, 0x0e, 0xee, 0xca, 0xc0, 0x27, 0x86, 0xb2, 0xfb, 0xd9,
	0xd2, 0x33, 0x48, 0x43, 0x3f, 0xb3, 0x3d, 0x4f, 0x4a, 0xbc, 0xf8, 0x7c, 0x75, 0xa2, 0x8a, 0xcf,
	0xaf, 0x4c, 0x82, 0xca, 0x1c, 0x3c, 0xb0, 0x05, 0xa6, 0x18, 0x8e, 0x05, 0xa6, 0x54, 0x28, 0x2f,
	0x9c, 0xa2, 0xb4, 0xa2, 0x70, 0x55, 0xd5, 0x16, 0x85, 0x5b, 0x14, 0xe5, 0x47, 0xb4, 0xf7, 0xa4,
	0x32, 0xb0, 0x29, 0x09, 0xcb, 0xf4, 0xd7, 0xed, 0xad, 0x0e, 0x88, 0x82, 0x7e, 0xfb, 0xa3, 0x7e,
	0x70, 0xf8, 0x06, 0xa2, 0x4e, 0x03, 0xf7, 0xb6, 0x01, 0x7f, 0x8f, 0x8f, 0xff, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x87, 0x25, 0x13, 0x08, 0x63, 0x0b, 0x00, 0x00,
}
