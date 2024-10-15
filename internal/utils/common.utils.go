package utils

import "strconv"

func HexStr(n int64) string {
  return strconv.FormatInt(int64(n), 16)
}
