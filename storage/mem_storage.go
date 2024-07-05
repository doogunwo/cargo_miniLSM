package lsmstorage

import (
    "fmt"
    "os"
    "sync"

    "main/storage" // 실제 프로젝트 경로로 바꿔주세요
)

type LsmStorage struct {
    mutable       *memtable.MemTable
    immutables    []*memtable.MemTable
    mu            sync.Mutex
    maxSize       int
    maxImmutables int
}

func NewLsmStorage(maxSize int, maxImmutables int) *LsmStorage {
    return &LsmStorage{
        mutable:       memtable.NewMemTable(),
        immutables:    []*memtable.MemTable{},
        maxSize:       maxSize,
        maxImmutables: maxImmutables,
    }
}

func (ls *LsmStorage) Put(key string, value string) {
    ls.mu.Lock()
    defer ls.mu.Unlock()

    ls.mutable.Put(key, value)
    if ls.mutable.Len() >= ls.maxSize {
        ls.freezeMemTable()
        if len(ls.immutables) >= ls.maxImmutables {
            ls.compact()
        }
    }
}

func (ls *LsmStorage) Get(key string) (string, bool) {
    ls.mu.Lock()
    defer ls.mu.Unlock()

    if value, found := ls.mutable.Get(key); found {
        return value, true
    }

    for _, immTable := range ls.immutables {
        if value, found := immTable.Get(key); found {
            return value, true
        }
    }
    return "", false
}

func (ls *LsmStorage) freezeMemTable() {
    frozenTable := ls.mutable
    ls.immutables = append(ls.immutables, frozenTable)
    ls.mutable = memtable.NewMemTable()
}

func (ls *LsmStorage) compact() {
    if len(ls.immutables) == 0 {
        return
    }

    immTable := ls.immutables[0]
    ls.immutables = ls.immutables[1:]

    filename := fmt.Sprintf("sst_file_%d.sst", len(ls.immutables))
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error creating SST file:", err)
        return
    }
    defer file.Close()

    iter := immTable.NewIterator()
    for iter.Next() {
        key := iter.Key().(float64)
        value := iter.Value().(string)
        _, err := fmt.Fprintf(file, "%f:%s\n", key, value)
        if err != nil {
            fmt.Println("Error writing to SST file:", err)
            return
        }
    }

    fmt.Println("Compaction completed, SST file created:", filename)
}
