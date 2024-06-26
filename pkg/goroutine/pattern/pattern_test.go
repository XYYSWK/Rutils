package pattern

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	now := time.Now()
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	ctx3, cancel3 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel3()
	newCtx := Or(ctx1, ctx2, ctx3)
	<-newCtx.Done()
	require.True(t, time.Since(now) < time.Second+100*time.Millisecond) // 验证时间间隔是否小于 1.1 秒
}
