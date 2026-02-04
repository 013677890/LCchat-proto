package service

import (
	"ChatServer/apps/user/internal/repository"
	pb "ChatServer/apps/user/pb"
	"ChatServer/consts"
	"ChatServer/pkg/logger"
	"ChatServer/pkg/util"
	"context"
	"errors"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// deviceServiceImpl 设备会话服务实现
type deviceServiceImpl struct {
	deviceRepo repository.IDeviceRepository
}

// NewDeviceService 创建设备服务实例
func NewDeviceService(deviceRepo repository.IDeviceRepository) DeviceService {
	return &deviceServiceImpl{
		deviceRepo: deviceRepo,
	}
}

// GetDeviceList 获取设备列表
func (s *deviceServiceImpl) GetDeviceList(ctx context.Context, req *pb.GetDeviceListRequest) (*pb.GetDeviceListResponse, error) {
	userUUID := util.GetUserUUIDFromContext(ctx)
	if userUUID == "" {
		logger.Warn(ctx, "获取设备列表失败：user_uuid 为空")
		return nil, status.Error(codes.Unauthenticated, strconv.Itoa(consts.CodeUnauthorized))
	}

	deviceID := util.GetDeviceIDFromContext(ctx)

	sessions, err := s.deviceRepo.GetByUserUUID(ctx, userUUID)
	if err != nil {
		logger.Error(ctx, "获取设备列表失败",
			logger.String("user_uuid", userUUID),
			logger.ErrorField("error", err),
		)
		return nil, status.Error(codes.Internal, strconv.Itoa(consts.CodeInternalError))
	}

	deviceIDs := make([]string, 0, len(sessions))
	for _, session := range sessions {
		if session == nil {
			continue
		}
		deviceIDs = append(deviceIDs, session.DeviceId)
	}

	activeTimes, err := s.deviceRepo.GetActiveTimestamps(ctx, userUUID, deviceIDs)
	if err != nil {
		logger.Warn(ctx, "获取设备活跃时间失败，使用当前时间兜底",
			logger.String("user_uuid", userUUID),
			logger.ErrorField("error", err),
		)
		activeTimes = map[string]int64{}
	}

	devices := make([]*pb.DeviceItem, 0, len(sessions))
	for _, session := range sessions {
		if session == nil {
			continue
		}
		sec, ok := activeTimes[session.DeviceId]
		if !ok || sec <= 0 {
			sec = time.Now().Unix()
			if err := s.deviceRepo.SetActiveTimestamp(ctx, userUUID, session.DeviceId, sec); err != nil {
				logger.Warn(ctx, "补写设备活跃时间失败",
					logger.String("user_uuid", userUUID),
					logger.String("device_id", session.DeviceId),
					logger.ErrorField("error", err),
				)
			}
		}
		lastSeenAt := sec * 1000
		devices = append(devices, &pb.DeviceItem{
			DeviceId:        session.DeviceId,
			DeviceName:      session.DeviceName,
			Platform:        session.Platform,
			AppVersion:      session.AppVersion,
			IsCurrentDevice: deviceID != "" && session.DeviceId == deviceID,
			Status:          int32(session.Status),
			LastSeenAt:      lastSeenAt,
		})
	}

	return &pb.GetDeviceListResponse{Devices: devices}, nil
}

// KickDevice 踢出设备
func (s *deviceServiceImpl) KickDevice(ctx context.Context, req *pb.KickDeviceRequest) error {
	userUUID := util.GetUserUUIDFromContext(ctx)
	if userUUID == "" {
		logger.Warn(ctx, "踢出设备失败：user_uuid 为空")
		return status.Error(codes.Unauthenticated, strconv.Itoa(consts.CodeUnauthorized))
	}

	if req == nil || req.DeviceId == "" {
		return status.Error(codes.InvalidArgument, strconv.Itoa(consts.CodeParamError))
	}

	currentDeviceID := util.GetDeviceIDFromContext(ctx)
	if currentDeviceID != "" && currentDeviceID == req.DeviceId {
		return status.Error(codes.FailedPrecondition, strconv.Itoa(consts.CodeCannotKickCurrent))
	}

	session, err := s.deviceRepo.GetByDeviceID(ctx, userUUID, req.DeviceId)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return status.Error(codes.NotFound, strconv.Itoa(consts.CodeDeviceNotFound))
		}
		logger.Error(ctx, "踢出设备失败：查询设备会话失败",
			logger.String("user_uuid", userUUID),
			logger.String("device_id", req.DeviceId),
			logger.ErrorField("error", err),
		)
		return status.Error(codes.Internal, strconv.Itoa(consts.CodeInternalError))
	}
	if session == nil {
		return status.Error(codes.NotFound, strconv.Itoa(consts.CodeDeviceNotFound))
	}

	// 幂等语义：无论 token 是否已删除，都返回成功；仅 Redis 异常才报错。
	if err := s.deviceRepo.DeleteTokens(ctx, userUUID, req.DeviceId); err != nil {
		logger.Error(ctx, "踢出设备失败：删除设备 Token 失败",
			logger.String("user_uuid", userUUID),
			logger.String("device_id", req.DeviceId),
			logger.ErrorField("error", err),
		)
		return status.Error(codes.Internal, strconv.Itoa(consts.CodeInternalError))
	}

	// status 语义：0=在线, 1=离线(被踢), 2=注销。
	// 注销设备保持 2；在线设备切到 1；已离线设备幂等成功。
	if session.Status == 0 {
		if err := s.deviceRepo.UpdateOnlineStatus(ctx, userUUID, req.DeviceId, 1); err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				return status.Error(codes.NotFound, strconv.Itoa(consts.CodeDeviceNotFound))
			}
			logger.Error(ctx, "踢出设备失败：更新设备状态失败",
				logger.String("user_uuid", userUUID),
				logger.String("device_id", req.DeviceId),
				logger.ErrorField("error", err),
			)
			return status.Error(codes.Internal, strconv.Itoa(consts.CodeInternalError))
		}
	}

	logger.Info(ctx, "踢出设备成功",
		logger.String("user_uuid", userUUID),
		logger.String("device_id", req.DeviceId),
		logger.Int("before_status", int(session.Status)),
	)

	return nil
}

// GetOnlineStatus 获取用户在线状态
func (s *deviceServiceImpl) GetOnlineStatus(ctx context.Context, req *pb.GetOnlineStatusRequest) (*pb.GetOnlineStatusResponse, error) {
	return nil, status.Error(codes.Unimplemented, "获取用户在线状态功能暂未实现")
}

// BatchGetOnlineStatus 批量获取在线状态
func (s *deviceServiceImpl) BatchGetOnlineStatus(ctx context.Context, req *pb.BatchGetOnlineStatusRequest) (*pb.BatchGetOnlineStatusResponse, error) {
	return nil, status.Error(codes.Unimplemented, "批量获取在线状态功能暂未实现")
}
