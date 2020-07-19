package disk

import (
  "fmt"
)

type DiskStats struct {
  Total uint64 `json:"Total"`
  Used  uint64 `json:"used"`
  Free  uint64 `json:"free"`
}
/*
func Usage(path string) (HDD DiskStats) {
  fs := syscall.Statfs_t{}
  err := syscall.Statfs(path, &fs)
  if err != nil {
    return
  }
  HDD.Total  = fs.Blocks * uint64(fs.Bsize)
  HDD.Free   = fs.Bfree * uint64(fs.Bsize)
  HDD.Used   = HDD.Total - HDD.Free
  return
}
*/
const (
  B = 1
  KB = 1024 * B
  MB = 1024 * KB
  GB = 1024 * MB
)

func CheckDisk() {
  //HDD := Usage("/")
  //fmt.Println("Total: %.2f GB\n", float64(HDD.Total)/float64(GB))
  //fmt.Println("Used: %.2f GB\n", float64(HDD.Used)/float64(GB))
  //fmt.Println("Free: %.2f GB\n", float64(HDD.Free)/float64(GB))
  fmt.Println("WIP")
}
