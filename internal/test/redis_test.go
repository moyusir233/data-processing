package test

import (
	"bytes"
	"context"
	"fmt"
	"gitee.com/moyusir/data-processing/internal/biz"
	"gitee.com/moyusir/data-processing/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"os"
	"testing"
	"time"
)

func TestData_RedisRepo(t *testing.T) {
	bootstrap, err := generalInit()
	if err != nil {
		t.Fatal(err)
	}

	// 实例化数据源和redisRepo
	logger := log.NewStdLogger(os.Stdout)
	client, clearnUp, err := data.NewData(bootstrap.Data, logger)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(clearnUp)

	redisRepo := data.NewRedisRepo(client, logger)

	// 测试配置查询
	t.Run("Test_GetDeviceConfig", func(t *testing.T) {
		// 初始化测试用例
		testCases := make([]struct {
			key, field string
			value      []byte
		}, 6)
		// 将测试用例分布在两个hash中
		for i := 0; i < 6; i++ {
			testCases[i].key = fmt.Sprintf("test%d", i%2)
			testCases[i].field = fmt.Sprintf("test%d", i)
			testCases[i].value = []byte("test")
		}

		// 注册清理函数以及向redis中写入数据
		t.Cleanup(func() {
			for _, c := range testCases {
				client.Del(context.Background(), c.key)
			}
		})

		for _, c := range testCases {
			err := client.HSet(
				context.Background(),
				c.key,
				c.field,
				// 二进制数据需要以十六进制字符串的形式存储，查询时函数会进行转换
				fmt.Sprintf("%x", c.value),
			).Err()
			if err != nil {
				t.Error(err)
				return
			}
		}

		for _, c := range testCases {
			value, err := redisRepo.GetDeviceConfig(c.key, c.field)
			if err != nil {
				t.Error(err)
				return
			}
			if !bytes.Equal(value, c.value) {
				t.Errorf(
					"The result of the query configuration does not match the expectation:%s %s",
					string(value), string(c.value),
				)
			}
		}
	})

	// 测试设备状态信息的批量查询
	t.Run("Test_BatchGetDeviceStateInfo", func(t *testing.T) {
		// 准备用于测试的数据
		var (
			key       = t.Name()
			testCases = make([]struct {
				value []byte
				time  time.Time
			}, 5)
			baseTime = time.Now()
		)
		for i := 0; i < 5; i++ {
			v := fmt.Sprintf("test%d", i)
			testCases[i].value = []byte(v)
			testCases[i].time = baseTime.Add(time.Duration(i) * time.Second)
		}

		// 注册清理函数并写入到redis中
		t.Cleanup(func() {
			client.Del(context.Background(), key)
		})
		for _, c := range testCases {
			err := client.ZAdd(context.Background(), key, &redis.Z{
				Score:  float64(c.time.Unix()),
				Member: fmt.Sprintf("%x", c.value),
			}).Err()
			if err != nil {
				t.Error(err)
				return
			}
		}

		// 测试分页查询
		t.Run("PageQuery", func(t *testing.T) {
			count := 2
			for offset := 0; offset < 5; offset += 2 {
				query, err := redisRepo.BatchGetDeviceStateInfo(key, &biz.QueryOption{
					Begin:  0,
					End:    0,
					Offset: int64(offset),
					Count:  int64(count),
				})
				if err != nil {
					t.Error(err)
				}

				// 检查查询得到的数据是否正确
				end := offset + count
				if end > 5 {
					end = 5
				}
				// 对用例也执行类似分页的分片，然后进行比较
				cases := testCases[offset:end]
				for i, c := range cases {
					if !bytes.Equal(c.value, query[i]) {
						t.Errorf(
							"The result of the query configuration does not match the expectation:%s %s",
							string(query[i]), string(c.value),
						)
					}
				}
			}
		})

		// 测试结合时间戳筛选的分页查询
		t.Run("PageQueryWithTimestamp", func(t *testing.T) {
			// 时间戳筛选的范围正好对应testCases[1:5]
			count := 2
			for offset := 0; offset < 4; offset += 2 {
				query, err := redisRepo.BatchGetDeviceStateInfo(key, &biz.QueryOption{
					Begin:  baseTime.Add(time.Second).Unix(),
					End:    baseTime.Add(4 * time.Second).Unix(),
					Offset: int64(offset),
					Count:  int64(count),
				})
				if err != nil {
					t.Error(err)
				}

				// 检查查询得到的数据是否正确
				end := offset + count
				if end > 5 {
					end = 5
				}
				// 对用例也执行类似分页的分片，然后进行比较
				cases := testCases[1:][offset:end]
				for i, c := range cases {
					if !bytes.Equal(c.value, query[i]) {
						t.Errorf(
							"The result of the query configuration does not match the expectation:%s %s",
							string(query[i]), string(c.value),
						)
					}
				}
			}
		})
	})

	// 测试设备预警字段的批量查询
	t.Run("Test_BatchGetDeviceWarningDetectField", func(t *testing.T) {
		// 准备测试数据
		var (
			label     = "test"
			testCases = make([]struct {
				key   string
				time  time.Time
				value float64
			}, 5)
		)
		for i := 0; i < 5; i++ {
			testCases[i].key = fmt.Sprintf("test%d", i)
			testCases[i].time = time.Now()
			testCases[i].value = float64(i)
		}

		// 写入redis中，并注册清理函数
		t.Cleanup(func() {
			for _, c := range testCases {
				client.Del(context.Background(), c.key)
			}
		})

		for _, c := range testCases {
			command := []interface{}{"TS.ADD"}
			command = append(command, c.key, c.time.Unix(), c.value)
			command = append(command, "LABELS", biz.WarningDetectFieldLabelName, label)
			err := client.Do(context.Background(), command...).Err()
			if err != nil {
				t.Error(err)
				return
			}
		}

		// 通过label进行批量查询
		field, err := redisRepo.BatchGetDeviceWarningDetectField(label, &biz.TSQueryOption{
			Begin:           0,
			End:             0,
			AggregationType: "",
			TimeBucket:      0,
		})
		if err != nil {
			t.Error(err)
			return
		}

		for i, f := range field {
			if f.Key != testCases[i].key {
				t.Errorf(
					"The result of the query key does not match the expectation:%s %s",
					f.Key, testCases[i].key,
				)
			}
			if f.Value != testCases[i].value {
				t.Errorf(
					"The result of the query value does not match the expectation:%f %f",
					f.Value, testCases[i].value,
				)
			}
			if f.Timestamp != testCases[i].time.Unix() {
				t.Errorf(
					"The result of the query timestamp does not match the expectation:%d %d",
					f.Timestamp, testCases[i].time.Unix(),
				)
			}
		}
	})

	// 警告消息相关的api即将数据存储在ZSet中并查询出来，因此不再额外编写测试
}
