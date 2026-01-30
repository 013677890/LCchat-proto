package service

import (
	"ChatServer/apps/gateway/internal/dto"
	"ChatServer/apps/gateway/internal/pb"
	"ChatServer/apps/gateway/internal/utils"
	userpb "ChatServer/apps/user/pb"
	"ChatServer/consts"
	"ChatServer/pkg/logger"
	"context"
	"time"
)

// FriendServiceImpl 好友服务实现
type FriendServiceImpl struct {
	userClient pb.UserServiceClient
}

// NewFriendService 创建好友服务实例
// userClient: 用户服务 gRPC 客户端
func NewFriendService(userClient pb.UserServiceClient) FriendService {
	return &FriendServiceImpl{
		userClient: userClient,
	}
}

// SearchUserByKeywordAndPageAndSize 搜索用户
// ctx: 请求上下文
// req: 搜索用户请求
// 返回: 搜索用户响应
func (s *FriendServiceImpl) SearchUserByKeywordAndPageAndSize(ctx context.Context, req *dto.SearchUserRequest) (*dto.SearchUserResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoSearchUserRequest(req)

	// 2. 调用好友服务搜索用户(gRPC)
	grpcResp, err := s.userClient.SearchUser(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用好友服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	logger.Info(ctx, "搜索用户成功",
		logger.String("keyword", req.Keyword),
		logger.Int32("page", req.Page),
		logger.Int32("page_size", req.PageSize),
		logger.Duration("duration", time.Since(startTime)),
	)

	return dto.ConvertSearchUserResponseFromProto(grpcResp), nil
}


// SendFriendApply 发送好友申请
func (s *FriendServiceImpl) SendFriendApply(ctx context.Context, req *dto.SendFriendApplyRequest) (*dto.SendFriendApplyResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoSendFriendApplyRequest(req)

	// 2. 调用用户服务发送好友申请(gRPC)
	grpcResp, err := s.userClient.SendFriendApply(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertFriendApplyResponseFromProto(grpcResp), nil
}

// GetFriendApplyList 获取好友申请列表
func (s *FriendServiceImpl) GetFriendApplyList(ctx context.Context, req *dto.GetFriendApplyListRequest) (*dto.GetFriendApplyListResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := &userpb.GetFriendApplyListRequest{
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 2. 调用用户服务获取好友申请列表(gRPC)
	grpcResp, err := s.userClient.GetFriendApplyList(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertGetFriendApplyListResponseFromProto(grpcResp), nil
}

// GetSentApplyList 获取发出的申请列表
func (s *FriendServiceImpl) GetSentApplyList(ctx context.Context, req *dto.GetSentApplyListRequest) (*dto.GetSentApplyListResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := &userpb.GetSentApplyListRequest{
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 2. 调用用户服务获取发出的申请列表(gRPC)
	grpcResp, err := s.userClient.GetSentApplyList(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertGetSentApplyListResponseFromProto(grpcResp), nil
}

// HandleFriendApply 处理好友申请
func (s *FriendServiceImpl) HandleFriendApply(ctx context.Context, req *dto.HandleFriendApplyRequest) (*dto.HandleFriendApplyResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoHandleFriendApplyRequest(req)

	// 2. 调用用户服务处理好友申请(gRPC)
	_, err := s.userClient.HandleFriendApply(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertHandleFriendApplyResponseFromProto(nil), nil
}

// GetUnreadApplyCount 获取未读申请数量
func (s *FriendServiceImpl) GetUnreadApplyCount(ctx context.Context, req *dto.GetUnreadApplyCountRequest) (*dto.GetUnreadApplyCountResponse, error) {
	startTime := time.Now()

	// 1. 调用用户服务获取未读申请数量(gRPC)
	grpcResp, err := s.userClient.GetUnreadApplyCount(ctx, &userpb.GetUnreadApplyCountRequest{})
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 2. gRPC 调用成功，返回结果
	return dto.ConvertGetUnreadApplyCountResponseFromProto(grpcResp), nil
}

// MarkApplyAsRead 标记申请已读
func (s *FriendServiceImpl) MarkApplyAsRead(ctx context.Context, req *dto.MarkApplyAsReadRequest) (*dto.MarkApplyAsReadResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoMarkApplyAsReadRequest(req)

	// 2. 调用用户服务标记申请已读(gRPC)
	_, err := s.userClient.MarkApplyAsRead(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertMarkApplyAsReadResponseFromProto(nil), nil
}

// GetFriendList 获取好友列表
func (s *FriendServiceImpl) GetFriendList(ctx context.Context, req *dto.GetFriendListRequest) (*dto.GetFriendListResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := &userpb.GetFriendListRequest{
		GroupTag: req.GroupTag,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 2. 调用用户服务获取好友列表(gRPC)
	grpcResp, err := s.userClient.GetFriendList(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertGetFriendListResponseFromProto(grpcResp), nil
}

// SyncFriendList 好友增量同步
func (s *FriendServiceImpl) SyncFriendList(ctx context.Context, req *dto.SyncFriendListRequest) (*dto.SyncFriendListResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := &userpb.SyncFriendListRequest{
		Version: req.Version,
		Limit:   req.Limit,
	}

	// 2. 调用用户服务同步好友列表(gRPC)
	grpcResp, err := s.userClient.SyncFriendList(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertSyncFriendListResponseFromProto(grpcResp), nil
}

// DeleteFriend 删除好友
func (s *FriendServiceImpl) DeleteFriend(ctx context.Context, req *dto.DeleteFriendRequest) (*dto.DeleteFriendResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoDeleteFriendRequest(req)

	// 2. 调用用户服务删除好友(gRPC)
	_, err := s.userClient.DeleteFriend(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertDeleteFriendResponseFromProto(nil), nil
}

// SetFriendRemark 设置好友备注
func (s *FriendServiceImpl) SetFriendRemark(ctx context.Context, req *dto.SetFriendRemarkRequest) (*dto.SetFriendRemarkResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoSetFriendRemarkRequest(req)

	// 2. 调用用户服务设置好友备注(gRPC)
	_, err := s.userClient.SetFriendRemark(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertSetFriendRemarkResponseFromProto(nil), nil
}

// SetFriendTag 设置好友标签
func (s *FriendServiceImpl) SetFriendTag(ctx context.Context, req *dto.SetFriendTagRequest) (*dto.SetFriendTagResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoSetFriendTagRequest(req)

	// 2. 调用用户服务设置好友标签(gRPC)
	_, err := s.userClient.SetFriendTag(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertSetFriendTagResponseFromProto(nil), nil
}

// GetTagList 获取标签列表
func (s *FriendServiceImpl) GetTagList(ctx context.Context, req *dto.GetTagListRequest) (*dto.GetTagListResponse, error) {
	startTime := time.Now()

	// 1. 调用用户服务获取标签列表(gRPC)
	grpcResp, err := s.userClient.GetTagList(ctx, &userpb.GetTagListRequest{})
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 2. gRPC 调用成功，返回结果
	return dto.ConvertGetTagListResponseFromProto(grpcResp), nil
}

// CheckIsFriend 判断是否好友
func (s *FriendServiceImpl) CheckIsFriend(ctx context.Context, req *dto.CheckIsFriendRequest) (*dto.CheckIsFriendResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoCheckIsFriendRequest(req)

	// 2. 调用用户服务判断是否好友(gRPC)
	grpcResp, err := s.userClient.CheckIsFriend(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertCheckIsFriendResponseFromProto(grpcResp), nil
}

// GetRelationStatus 获取关系状态
func (s *FriendServiceImpl) GetRelationStatus(ctx context.Context, req *dto.GetRelationStatusRequest) (*dto.GetRelationStatusResponse, error) {
	startTime := time.Now()

	// 1. 转换 DTO 为 Protobuf 请求
	grpcReq := dto.ConvertToProtoGetRelationStatusRequest(req)

	// 2. 调用用户服务获取关系状态(gRPC)
	grpcResp, err := s.userClient.GetRelationStatus(ctx, grpcReq)
	if err != nil {
		// gRPC 调用失败，提取业务错误码
		code := utils.ExtractErrorCode(err)
		// 记录错误日志
		logger.Error(ctx, "调用用户服务 gRPC 失败",
			logger.ErrorField("error", err),
			logger.Int("business_code", code),
			logger.String("business_message", consts.GetMessage(code)),
			logger.Duration("duration", time.Since(startTime)),
		)
		// 返回业务错误（作为 Go error 返回，由 Handler 层处理）
		return nil, err
	}

	// 3. gRPC 调用成功，返回结果
	return dto.ConvertGetRelationStatusResponseFromProto(grpcResp), nil
}
