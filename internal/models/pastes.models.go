package models

type Paste struct {
  PasteId   string `redis:"PasteId"`
  BurnAfter uint64 `redis:"BurnAfter"`
  ReadCount uint64 `redis:"ReadCount"`
  DeepUrl   uint8  `redis:"DeepUrl"`
  Secret    string `redis:"Secret"`
}
