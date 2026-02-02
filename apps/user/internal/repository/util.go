package repository

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

type friendMeta struct {
	Remark    string `json:"remark"`
	GroupTag  string `json:"group_tag"`
	Source    string `json:"source"`
	UpdatedAt int64  `json:"updated_at"`
}

func buildFriendMetaJSON(remark, groupTag, source string, updatedAt int64) string {
	meta := friendMeta{
		Remark:    remark,
		GroupTag:  groupTag,
		Source:    source,
		UpdatedAt: updatedAt,
	}
	data, err := json.Marshal(&meta)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func parseFriendMetaJSON(raw string) (*friendMeta, error) {
	var meta friendMeta
	if err := json.Unmarshal([]byte(raw), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func isRedisWrongType(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "WRONGTYPE")
}

// getRandomExpireTime 生成带随机抖动的过期时间
// baseExpire: 基础过期时间
// 返回: 基础过期时间 ± 10% 的随机时间
func getRandomExpireTime(baseExpire time.Duration) time.Duration {
	// 计算随机抖动范围（±10%）
	jitterRange := float64(baseExpire) * 0.1
	jitter := time.Duration(rand.Float64()*float64(jitterRange)*2 - float64(jitterRange))

	return baseExpire + jitter
}

// getRandomBool 生成随机布尔值
// probability: 概率
// 返回: 概率为probability的布尔值
func getRandomBool(probability float64) bool {
	return rand.Float64() < probability
}
