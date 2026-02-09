package manager

import "sync"

// ConnectionManager 管理所有在线 WebSocket 连接。
// 维护两套索引：
// - byKey(user_uuid:device_id) 用于精确定位单设备连接；
// - byUser(user_uuid -> device_id -> client) 用于按用户广播。
type ConnectionManager struct {
	mu       sync.RWMutex
	byKey    map[string]*Client
	byUser   map[string]map[string]*Client
	shutdown bool
}

// NewConnectionManager 创建连接管理器实例。
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		byKey:  make(map[string]*Client),
		byUser: make(map[string]map[string]*Client),
	}
}

// Register 注册一个设备连接。
// 返回值 replaced 表示被新连接替换掉的旧连接（如果存在）。
// 调用方通常应主动关闭 replaced，确保同设备最多一个活跃连接。
func (m *ConnectionManager) Register(client *Client) (replaced *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.shutdown {
		return nil
	}

	key := client.Key()
	if old, ok := m.byKey[key]; ok && old != client {
		replaced = old
	}

	m.byKey[key] = client
	userConns, ok := m.byUser[client.UserUUID()]
	if !ok {
		userConns = make(map[string]*Client)
		m.byUser[client.UserUUID()] = userConns
	}
	userConns[client.DeviceID()] = client
	return replaced
}

// Unregister 注销一个连接。
// 只有当 map 中当前连接与入参完全一致时才删除，防止并发替换时误删新连接。
func (m *ConnectionManager) Unregister(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := client.Key()
	current, ok := m.byKey[key]
	if !ok || current != client {
		return
	}

	delete(m.byKey, key)
	if userConns, ok := m.byUser[client.UserUUID()]; ok {
		delete(userConns, client.DeviceID())
		if len(userConns) == 0 {
			delete(m.byUser, client.UserUUID())
		}
	}
}

// SendToDevice 向指定用户的指定设备发送消息。
// 返回 false 表示目标连接不存在或写队列不可用。
func (m *ConnectionManager) SendToDevice(userUUID, deviceID string, msg []byte) bool {
	m.mu.RLock()
	client := m.byKey[buildKey(userUUID, deviceID)]
	m.mu.RUnlock()
	if client == nil {
		return false
	}
	return client.Enqueue(msg)
}

// SendToUser 向用户的所有在线设备广播消息。
// 返回成功入队的设备数量，可用于统计下行投递率。
func (m *ConnectionManager) SendToUser(userUUID string, msg []byte) int {
	m.mu.RLock()
	userConns, ok := m.byUser[userUUID]
	if !ok || len(userConns) == 0 {
		m.mu.RUnlock()
		return 0
	}
	clients := make([]*Client, 0, len(userConns))
	for _, client := range userConns {
		clients = append(clients, client)
	}
	m.mu.RUnlock()

	sent := 0
	for _, client := range clients {
		if client.Enqueue(msg) {
			sent++
		}
	}
	return sent
}

// Count 返回当前在线连接数（按 user_uuid+device_id 去重后）。
func (m *ConnectionManager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.byKey)
}

// Shutdown 关闭全部连接并阻止后续注册。
// 用于进程优雅退出阶段，确保不再接收新连接并尽快释放资源。
func (m *ConnectionManager) Shutdown() {
	m.mu.Lock()
	if m.shutdown {
		m.mu.Unlock()
		return
	}
	m.shutdown = true

	clients := make([]*Client, 0, len(m.byKey))
	for _, client := range m.byKey {
		clients = append(clients, client)
	}
	m.byKey = make(map[string]*Client)
	m.byUser = make(map[string]map[string]*Client)
	m.mu.Unlock()

	for _, client := range clients {
		client.Close()
	}
}

// buildKey 统一构造设备连接键。
func buildKey(userUUID, deviceID string) string {
	return userUUID + ":" + deviceID
}
