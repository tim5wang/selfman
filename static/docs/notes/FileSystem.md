
# 文件系统

## Golang 中对文件系统的抽象

```go
type FS interface {
	Open(name string) (File, error)
}

type GlobFS interface {
    FS
    Glob(pattern string) ([]string, error) // 搜索文件
}

type File interface {
    Stat() (FileInfo, error)
    Read([]byte) (int, error)
    Close() error
}

type FileInfo interface {
    Name() string       
    Size() int64        
    Mode() FileMode     
    ModTime() time.Time 
    IsDir() bool        
    Sys() interface{}   
}

type FileMode uint32

type DirEntry interface {
    Name() string
    IsDir() bool
    Type() FileMode
    Info() (FileInfo, error)
}

type ReadDirFile interface {
    File
    ReadDir(n int) ([]DirEntry, error)
}
type DirEntry interface {
    Name() string
    IsDir() bool
    Type() FileMode
    Info() (FileInfo, error)
}

// 写文件
func (f *File) Write(b []byte) (n int, err error)
// 带偏移量的写文件
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
// 设置文件的读写光标位置
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
// whence的值，在os包中定义了相应的常量，应该使用这些常量
const (
    SEEK_SET int = 0 // seek relative to the origin of the file
    SEEK_CUR int = 1 // seek relative to the current offset
    SEEK_END int = 2 // seek relative to the end
)
```

## NFS

## HDFS

## GFS

## [SeaweedFS](https://github.com/chrislusf/seaweedfs)
