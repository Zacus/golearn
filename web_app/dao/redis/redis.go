package redis

import (
	"context"
	"fmt"
	"golearn/web_app/settings"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// 初始化连接
func Init(conf *settings.RedisConfig) (err error) {
	// rdb = redis.NewClient(&redis.Options{
	// 	Addr:     conf.Host,
	// 	Password: conf.Password,     // no password set
	// 	DB:       conf.MDb,          // use default DB
	// 	PoolSize: conf.MaxOpenConns, // 连接池大小
	// })

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password, // no password set
		DB:       conf.MDb,      // use default DB
		PoolSize: conf.MaxOpenConns,
	})
	//redis://<user>:<pass>@localhost:6379/<db>
	// dsn := fmt.Sprintf("redis://%s:%s@%s:%d/%d",
	// 	conf.User,
	// 	conf.Password,
	// 	conf.Host,
	// 	conf.Port,
	// 	conf.MDb,
	// )

	// opt, err := redis.ParseURL(dsn)
	// if err != nil {
	// 	panic(err)
	// }

	// rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err

}

func Close() {
	_ = rdb.Close()
}
