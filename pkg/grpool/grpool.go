package grpool

import (
	"sync"

	"github.com/baojingh/prctl/internal/logger"
	"github.com/panjf2000/ants/v2"
)

var log = logger.New()

var once sync.Once
var pool *ants.Pool

func init() {
	once.Do(func() {
		var err error
		pool, err = ants.NewPool(5)
		if err != nil {
			log.Error("Failed to init ants Pool")
			return
		}
		log.Info("Init ants pool success")
	})
}

func SubmitTask(task func()) error {
	if pool == nil {
		panic("gr pool is not init success")
	}
	res := pool.Submit(task)
	return res
}

func ShutdownPool() {
	if pool != nil {
		pool.Release()
		pool = nil
		log.Info("Shutdown ants pool success")
	}
}
