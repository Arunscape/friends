package datatypes

import (
    "time"
)

type Message struct {
  Body string
  Sender string
  Timestamp uint64
}

func (m *Message) SetTimestamp() {
  m.Timestamp = uint64(time.Now().UnixNano() / int64(time.Millisecond))
}
