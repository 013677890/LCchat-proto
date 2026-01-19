package consts

// 通用错误码
const (
	CodeSuccess = 0 // 成功
)

// 客户端错误 (1xxxx)
const (
	CodeParamError       = 10001 // 参数验证失败
	CodeBodyError        = 10002 // 请求体格式错误
	CodeResourceNotFound = 10003 // 资源不存在
	CodeMethodNotAllowed = 10004 // 请求方法不允许
	CodeTooManyRequests  = 10005 // 请求过于频繁
	CodeBodyTooLarge     = 10006 // 请求体过大
)

// 认证错误 (2xxxx)
const (
	CodeUnauthorized   = 20001 // 未认证
	CodeInvalidToken   = 20002 // Token 无效
	CodeTokenExpired   = 20003 // Token 已过期
	CodePermissionDeny = 20004 // 权限不足
)

// 用户模块错误 (11xxx)
const (
	CodeUserNotFound          = 11001 // 用户不存在
	CodeUserAlreadyExist      = 11002 // 用户已存在
	CodePasswordError         = 11003 // 密码错误
	CodeUserDisabled          = 11004 // 用户已被禁用
	CodeEmailFormatError      = 11005 // 邮箱格式错误
	CodeVerifyCodeError       = 11006 // 验证码错误
	CodeVerifyCodeExpire      = 11007 // 验证码已过期
	CodePhoneFormatError      = 11008 // 手机号格式错误
	CodePasswordSameAsOld     = 11009 // 新密码不能与旧密码相同
	CodeNicknameAlreadyExist  = 11010 // 昵称已被使用
	CodeFileFormatNotSupport  = 11011 // 文件格式不支持
	CodeFileUploadFail        = 11012 // 文件上传失败
	CodeQRCodeFormatError     = 11013 // 二维码格式错误
	CodeQRCodeExpired         = 11014 // 二维码已过期
	CodeEmailAlreadyExist     = 11015 // 邮箱已被使用
	CodeTelephoneAlreadyExist = 11016 // 手机号已被使用
	CodeAccountNotFound       = 11017 // 账号不存在(邮箱或手机号)
	CodeVerifyCodeTypeInvalid = 11018 // 验证码类型无效
	CodePasswordFormatError   = 11019 // 密码格式错误
	CodeNicknameFormatError   = 11020 // 昵称格式错误
	CodeSignatureTooLong      = 11021 // 个性签名过长
	CodeBirthdayFormatError   = 11022 // 生日格式错误
	CodeGenderInvalid         = 11023 // 性别值无效
	CodeRemarkTooLong         = 11024 // 备注过长
	CodeReasonTooLong         = 11025 // 理由过长
)

// 好友模块错误 (12xxx)
const (
	CodeAlreadyFriend         = 12001 // 已经是好友
	CodeFriendRequestSent     = 12002 // 好友申请已发送
	CodeNotFriend             = 12003 // 不存在该好友关系
	CodeIsBlacklist           = 12004 // 已经是黑名单(保留未使用)
	CodeApplyNotFoundOrHandle = 12005 // 申请不存在或已处理
	CodeNoPermissionHandle    = 12006 // 无权限处理该申请
	CodeCannotAddSelf         = 12007 // 不能添加自己为好友
	CodeFriendLimitExceeded   = 12008 // 好友数量已达上限
	CodeApplyExpired          = 12009 // 申请已过期
	CodeTagNameInvalid        = 12010 // 标签名称无效
	CodeSourceInvalid         = 12011 // 来源参数无效
)

// 消息模块错误 (13xxx)
const (
	CodeMessageNotFound       = 13001 // 消息不存在
	CodeMessageSendFail       = 13002 // 消息发送失败
	CodeMessageTypeNotSupport = 13003 // 消息类型不支持
	CodeConversationNotFound  = 13004 // 会话不存在
	CodeMessageContentEmpty   = 13005 // 消息内容为空
	CodeMessageTooLong        = 13006 // 消息内容过长
	CodeMessageRevoked        = 13007 // 消息已撤回
	CodeMessageDeleted        = 13008 // 消息已删除
)

// 群组模块错误 (14xxx)
const (
	CodeGroupNotFound       = 14001 // 群组不存在
	CodeNotGroupMember      = 14002 // 不是群成员
	CodeNoPermission        = 14003 // 没有权限
	CodeGroupFull           = 14004 // 群成员已满
	CodeGroupNameTooLong    = 14005 // 群名称过长
	CodeGroupNoticeTooLong  = 14006 // 群公告过长
	CodeGroupAlreadyDismiss = 14007 // 群组已解散
	CodeGroupMemberNotFound = 14008 // 群成员不存在
	CodeCannotKickOwner     = 14009 // 不能踢出群主
	CodeCannotKickAdmin     = 14010 // 不能踢出管理员
	CodeAlreadyGroupMember  = 14011 // 已经是群成员
	CodeGroupApplyNotFound  = 14012 // 入群申请不存在
	CodeGroupInviteLimit    = 14013 // 邀请人数超限
	CodeCannotQuitAsOwner   = 14014 // 群主不能退群
	CodeAdminLimitExceeded  = 14015 // 管理员数量已达上限
)

// 设备会话错误 (15xxx)
const (
	CodeDeviceCreateFail    = 15001 // 设备会话创建失败
	CodeDeviceAlreadyExist  = 15002 // 设备会话已存在
	CodeDeviceUpdateFail    = 15003 // 设备会话更新失败
	CodeDeviceNotFound      = 15004 // 设备会话不存在
	CodeCannotKickCurrent   = 15005 // 不能踢出当前设备
	CodeDeviceLimitExceeded = 15006 // 超过最大设备数限制
	CodeDeviceOffline       = 15007 // 设备已离线
	CodeDeviceInfoInvalid   = 15008 // 设备信息无效
	CodePlatformNotSupport  = 15009 // 平台不支持
)

// 黑名单错误 (16xxx)
const (
	CodePeerBlacklistYou    = 16001 // 对方已将你拉黑
	CodeYouBlacklistPeer    = 16002 // 你已将对方拉黑
	CodeAlreadyInBlacklist  = 16003 // 已在黑名单中
	CodeNotInBlacklist      = 16004 // 不在黑名单中
	CodeCannotBlacklistSelf = 16005 // 不能拉黑自己
)

// 服务端错误 (3xxxx)
const (
	CodeInternalError      = 30001 // 服务器内部错误
	CodeServiceUnavailable = 30002 // 服务暂不可用
	CodeTimeoutError       = 30003 // 超时错误
)

// 错误消息映射
var CodeMessage = map[int32]string{
	CodeSuccess: "success",

	// 客户端错误
	CodeParamError:       "参数验证失败",
	CodeBodyError:        "请求体格式错误",
	CodeResourceNotFound: "资源不存在",
	CodeMethodNotAllowed: "请求方法不允许",
	CodeTooManyRequests:  "请求过于频繁",
	CodeBodyTooLarge:     "请求体过大",

	// 认证错误
	CodeUnauthorized:   "未认证",
	CodeInvalidToken:   "Token 无效",
	CodeTokenExpired:   "Token 已过期",
	CodePermissionDeny: "权限不足",

	// 用户模块
	CodeUserNotFound:          "用户不存在",
	CodeUserAlreadyExist:      "用户已存在",
	CodePasswordError:         "密码错误",
	CodeUserDisabled:          "用户已被禁用",
	CodeEmailFormatError:      "邮箱格式错误",
	CodeVerifyCodeError:       "验证码错误",
	CodeVerifyCodeExpire:      "验证码已过期",
	CodePhoneFormatError:      "手机号格式错误",
	CodePasswordSameAsOld:     "新密码不能与旧密码相同",
	CodeNicknameAlreadyExist:  "昵称已被使用",
	CodeFileFormatNotSupport:  "文件格式不支持",
	CodeFileUploadFail:        "文件上传失败",
	CodeQRCodeFormatError:     "二维码格式错误",
	CodeQRCodeExpired:         "二维码已过期",
	CodeEmailAlreadyExist:     "邮箱已被使用",
	CodeTelephoneAlreadyExist: "手机号已被使用",
	CodeAccountNotFound:       "账号不存在",
	CodeVerifyCodeTypeInvalid: "验证码类型无效",
	CodePasswordFormatError:   "密码格式错误",
	CodeNicknameFormatError:   "昵称格式错误",
	CodeSignatureTooLong:      "个性签名过长",
	CodeBirthdayFormatError:   "生日格式错误",
	CodeGenderInvalid:         "性别值无效",
	CodeRemarkTooLong:         "备注过长",
	CodeReasonTooLong:         "理由过长",

	// 好友模块
	CodeAlreadyFriend:         "已经是好友",
	CodeFriendRequestSent:     "好友申请已发送",
	CodeNotFriend:             "不存在该好友关系",
	CodeIsBlacklist:           "已经是黑名单",
	CodeApplyNotFoundOrHandle: "申请不存在或已处理",
	CodeNoPermissionHandle:    "无权限处理该申请",
	CodeCannotAddSelf:         "不能添加自己为好友",
	CodeFriendLimitExceeded:   "好友数量已达上限",
	CodeApplyExpired:          "申请已过期",
	CodeTagNameInvalid:        "标签名称无效",
	CodeSourceInvalid:         "来源参数无效",

	// 消息模块
	CodeMessageNotFound:       "消息不存在",
	CodeMessageSendFail:       "消息发送失败",
	CodeMessageTypeNotSupport: "消息类型不支持",
	CodeConversationNotFound:  "会话不存在",
	CodeMessageContentEmpty:   "消息内容为空",
	CodeMessageTooLong:        "消息内容过长",
	CodeMessageRevoked:        "消息已撤回",
	CodeMessageDeleted:        "消息已删除",

	// 群组模块
	CodeGroupNotFound:       "群组不存在",
	CodeNotGroupMember:      "不是群成员",
	CodeNoPermission:        "没有权限",
	CodeGroupFull:           "群成员已满",
	CodeGroupNameTooLong:    "群名称过长",
	CodeGroupNoticeTooLong:  "群公告过长",
	CodeGroupAlreadyDismiss: "群组已解散",
	CodeGroupMemberNotFound: "群成员不存在",
	CodeCannotKickOwner:     "不能踢出群主",
	CodeCannotKickAdmin:     "不能踢出管理员",
	CodeAlreadyGroupMember:  "已经是群成员",
	CodeGroupApplyNotFound:  "入群申请不存在",
	CodeGroupInviteLimit:    "邀请人数超限",
	CodeCannotQuitAsOwner:   "群主不能退群",
	CodeAdminLimitExceeded:  "管理员数量已达上限",

	// 设备会话
	CodeDeviceCreateFail:    "设备会话创建失败",
	CodeDeviceAlreadyExist:  "设备会话已存在",
	CodeDeviceUpdateFail:    "设备会话更新失败",
	CodeDeviceNotFound:      "设备会话不存在",
	CodeCannotKickCurrent:   "不能踢出当前设备",
	CodeDeviceLimitExceeded: "超过最大设备数限制",
	CodeDeviceOffline:       "设备已离线",
	CodeDeviceInfoInvalid:   "设备信息无效",
	CodePlatformNotSupport:  "平台不支持",

	// 黑名单
	CodePeerBlacklistYou:    "对方已将你拉黑",
	CodeYouBlacklistPeer:    "你已将对方拉黑",
	CodeAlreadyInBlacklist:  "已在黑名单中",
	CodeNotInBlacklist:      "不在黑名单中",
	CodeCannotBlacklistSelf: "不能拉黑自己",

	// 服务端错误
	CodeInternalError:      "服务器内部错误",
	CodeServiceUnavailable: "服务暂不可用",
	CodeTimeoutError:       "超时错误",
}

// GetMessage 根据错误码获取错误消息
func GetMessage(code int32) string {
	if msg, ok := CodeMessage[code]; ok {
		return msg
	}
	return "未知错误"
}
