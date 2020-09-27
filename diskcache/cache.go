package diskcache

import (
	"context"
)

// 磁盘缓存接口
type Interface interface {
	Store(ctx context.Context, key string, value []byte) error
	Load(ctx context.Context, key string) (value []byte, exist bool, err error)
	Delete(ctx context.Context, key string) error
}
