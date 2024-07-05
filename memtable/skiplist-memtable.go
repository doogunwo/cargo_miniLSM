package memtable

import (
    "fmt"
    "github.com/sean-public/fast-skiplist"
)

type MemTable struct {
    table *skiplist.SkilList
}

func NewMEmTable() *MemTable {
    return &MemTable {
        table: skiplist.New(),
    }
}

func (mt *MemTable) Put(key string, value string) {
    mt.table.Set(key, value)
}

func (mt *MemTable) Get(key string) (string, bool) {
    if entry, found := mt.table.Get(key); found {
        return entry.(string), true
    }
    return "",false
}

func main() {
    memTable := NewMemTable()

    memTable.Put("key1", "value1")
    memTable.Put("key2", "value2")

    if value, found := memTable.Get("key1"); found {
        fmt.Println("Found key1:", value)
    } else {
        fmt.Println("key1 not found")
    }

    if value, found := memTable.Get("key3"); found {
        fmt.Println("Found key3:", value)
    } else {
        fmt.Println("key3 not found")
    }
}

