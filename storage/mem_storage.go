package stroage

import (
    "bytes"
    "os"
    "sync"
)

const typeShift = 4

var _ [0]struct{} = [TypeAll >> typeShift]struct{}{}
