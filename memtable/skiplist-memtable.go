package memtable

import (
    "crypto/md5"
    "encoding/binary"
    "github.com/sean-public/fast-skiplist"
)
//스킵리스트 멤테이블
type MemTable struct {
    table *skiplist.SkipList
}
//멤테이블 생성
func NewMemTable() *MemTable {
    return &MemTable{
        table: skiplist.New(),
    }
}
//스킵리스트는 float64 키 요구함
func hashKey(key string) float64 {
    hasher := md5.New()
    hasher.Write([]byte(key))
    hashBytes := hasher.Sum(nil)
    return float64(binary.BigEndian.Uint64(hashBytes[:8]))
}

//키 삽입
func (mt *MemTable) Put(key string, value string) {
    hashedKey := hashKey(key)
    mt.table.Set(hashedKey, value)
}

//키 가져오기 조회임
func (mt *MemTable) Get(key string) (string, bool) {
    hashedKey := hashKey(key)
    entry := mt.table.Get(hashedKey)
    if entry != nil {
        return entry.Value().(string), true
    }
    return "", false
}

//멤테이블 길이 재기
func (mt *MemTable) Len() int {
    return mt.table.Len()
}

//이터레이터 생성
func (mt *MemTable) NewIterator() *skiplist.Iterator {
    return mt.table.NewIterator()
}
