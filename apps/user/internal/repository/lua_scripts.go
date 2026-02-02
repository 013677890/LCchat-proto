package repository

const (
	// luaIncrementWithExpire 递增计数器，仅在首次创建时设置过期时间
	// KEYS[1]: 计数器 key
	// ARGV[1]: 过期时间（秒）
	// 返回: 递增后的值
	luaIncrementWithExpire = `
local key = KEYS[1]
local expire = tonumber(ARGV[1])
local current = redis.call('INCR', key)

-- 如果是第一次创建值为1,则设置过期时间
if current == 1 then
	redis.call('EXPIRE', key, expire)
end

return current
`

	// luaAddPendingApplyIfExists 申请写入（仅在 key 存在时增量更新）
	// KEYS[1]: 待处理申请 ZSet
	// ARGV[1]: score(created_at unix)
	// ARGV[2]: member(applicant_uuid)
	// ARGV[3]: 过期时间（秒）
	// 返回: 1 表示写入成功，0 表示 key 不存在
	luaAddPendingApplyIfExists = `
if redis.call('EXISTS', KEYS[1]) == 1 then
	redis.call('ZREM', KEYS[1], '__EMPTY__')
	redis.call('ZADD', KEYS[1], ARGV[1], ARGV[2])
	redis.call('EXPIRE', KEYS[1], ARGV[3])
	return 1
end
return 0
`

	// luaRemoveFriendIfExists 删除好友缓存（仅在 key 存在时增量更新）
	// KEYS[1]: 好友关系 Set
	// ARGV[1]: friend_uuid
	// ARGV[2]: 过期时间（秒）
	// 返回: 1 表示写入成功，0 表示 key 不存在
	luaRemoveFriendIfExists = `
if redis.call('EXISTS', KEYS[1]) == 1 then
	redis.call('SREM', KEYS[1], ARGV[1])
	redis.call('SREM', KEYS[1], '__EMPTY__')
	if redis.call('SCARD', KEYS[1]) == 0 then
		redis.call('SADD', KEYS[1], '__EMPTY__')
	end
	redis.call('EXPIRE', KEYS[1], ARGV[2])
	return 1
end
return 0
`
)
