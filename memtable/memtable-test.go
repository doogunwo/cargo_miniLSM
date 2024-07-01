package memtable

import (
  "fmt"
  "testing"
)

func Testmemdb(t *testing.T){
  memtable := NewMemTable()

  memtable.Put("key1", []byte("v1"))
  memtable.Put("key1", []byte("v2"))

  val, exists := memtable.Get("key1")
  if !exists || string(val) != "v1" {
    t.Errorf("v1 for k1 for %s", string(val))
  }
  
  fmt.Println("passed")

}
