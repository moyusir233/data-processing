package data

import (
	"context"
	"fmt"
	"gitee.com/moyusir/data-processing/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// RedisRepo redis数据库操作对象，可以理解为dao
type RedisRepo struct {
	client *Data
	logger *log.Helper
}

// NewRedisRepo 实例化redis数据库操作对象
func NewRedisRepo(data *Data, logger log.Logger) biz.UnionRepo {
	return &RedisRepo{
		client: data,
		logger: log.NewHelper(logger),
	}
}

// GetDeviceConfig 查询保存在指定key的hash的field中的设备配置信息
func (r *RedisRepo) GetDeviceConfig(key, field string) ([]byte, error) {
	result, err := r.client.HGet(context.Background(), key, field).Result()
	if err != nil {
		return nil, err
	}

	// 二进制信息统一以十六进制字符串的信息在redis中保存，因此需要转换
	var config []byte
	_, err = fmt.Sscanf(result, "%x", &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// BatchGetDeviceStateInfo 批量查询存储在ZSet中的设备状态的二进制信息
func (r *RedisRepo) BatchGetDeviceStateInfo(key string, option *biz.QueryOption) ([][]byte, error) {
	return r.batchQueryWithZset(key, option)
}

// BatchGetDeviceWarningDetectField 批量查询存储在ts中的预警字段信息
func (r *RedisRepo) BatchGetDeviceWarningDetectField(label string, option *biz.TSQueryOption) ([]biz.TsQueryResult, error) {
	command := []interface{}{"TS.MRANGE"}

	// 参数转换

	// 查询的时间戳范围
	if option.Begin != 0 {
		command = append(command, option.Begin)
	} else {
		command = append(command, "-")
	}
	if option.End != 0 {
		command = append(command, option.End)
	} else {
		command = append(command, "+")
	}

	// 配置时间桶的对齐
	command = append(command, "ALIGN", "start")

	// 配置聚合查询
	if option.AggregationType != "" {
		command = append(command,
			"AGGREGATION",
			option.AggregationType,
			option.TimeBucket.Milliseconds())
	}

	// 配置依据标签进行过滤
	command = append(command, "FILTER", biz.WarningDetectFieldLabelName+"="+label)

	slice, err := r.client.Do(context.Background(), command...).Slice()
	if err != nil {
		return nil, err
	}

	// 将查询结果进行处理
	// redis返回的是一个数组，其中每个元素仍是一个数组，
	// 每个元素数组的第一个元素是TS的key，第二个元素是label数组，第三个元素是由timestamp-value对组成的数组
	result := make([]biz.TsQueryResult, 0, len(slice))
	for _, sliceInterface := range slice {
		s := sliceInterface.([]interface{})
		if len(s) == 0 {
			continue
		}

		qr := biz.TsQueryResult{}
		qr.Key = s[0].(string)
		timestampAndValues := s[2].([]interface{})

		for _, pairInterface := range timestampAndValues {
			pair := pairInterface.([]interface{})

			qr.Timestamp = pair[0].(int64)

			value, err := strconv.ParseFloat(pair[1].(string), 64)
			if err != nil {
				return nil, err
			}
			qr.Value = value
			result = append(result, qr)
		}
	}

	return result, nil
}

// GetWarningMessage 查询存储在ZSet中的警告信息
func (r *RedisRepo) GetWarningMessage(key string, option *biz.QueryOption) ([][]byte, error) {
	return r.batchQueryWithZset(key, option)
}

// SaveWarningMessage 保存警告信息至ZSet中
func (r *RedisRepo) SaveWarningMessage(key string, timestamp int64, value []byte) error {
	return r.client.ZAdd(context.Background(), key, &redis.Z{
		Score:  float64(timestamp),
		Member: fmt.Sprintf("%x", value),
	}).Err()
}

func (r *RedisRepo) batchQueryWithZset(key string, option *biz.QueryOption) ([][]byte, error) {
	// 将传入参数进行转换
	o := new(redis.ZRangeBy)
	if option.Begin != 0 {
		o.Min = strconv.FormatInt(option.Begin, 10)
	} else {
		o.Min = "-inf"
	}
	if option.End != 0 {
		o.Max = strconv.FormatInt(option.End, 10)
	} else {
		o.Max = "+inf"
	}
	o.Count = option.Count
	o.Offset = option.Offset

	result, err := r.client.ZRangeByScore(context.Background(), key, o).Result()
	if err != nil {
		return nil, err
	}

	// 二进制信息统一以十六进制字符串的信息在redis中保存，因此需要转换
	states := make([][]byte, len(result))
	for i, state := range result {
		_, err := fmt.Sscanf(state, "%x", &(states[i]))
		if err != nil {
			return nil, err
		}
	}

	return states, nil
}
