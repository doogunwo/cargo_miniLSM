package memtable

import (
  "sync"
)

type MemTable struct {
    mu sync.RWMutex //뮤텍스
    table map[string][]byte // 테이블
}

func NewMemTable() *MemTable {
    return &MemTable {
      table: make(map[string][]byte),
    }
}
// memtable에 키에 해당하는 값 반환
func (m *MemTable) Put(key string, value []byte) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.table[key] = value
}

//get memtable에 해당하는 키 반환
func (m *MemTable) Get(key string) ([]byte, bool) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    value, exists := m.table[key]
    return value, exists
}

// Delete 는 MemTable에서 키 삭제
func (m *MemTable) Delete(key string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    delete(m.table, key)
}

func (m *MemTable) Flush() map[string][]byte {
  m.mu.Lock()
  defer m.mu.Unlock()
  flushedTable := m.table
  m.table = make(map[string][]byte)
  return flushedTable
}


