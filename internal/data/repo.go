package data

import (
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/biz"
	"gitee.com/moyusir/data-processing/internal/conf"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"strings"
	"time"
)

// Repo redis数据库操作对象，可以理解为dao
type Repo struct {
	redisClient    *RedisData
	influxdbClient *InfluxdbData
}

// NewRepo 实例化redis数据库操作对象
func NewRepo(redisData *RedisData, influxdbData *InfluxdbData) biz.UnionRepo {
	return &Repo{
		redisClient:    redisData,
		influxdbClient: influxdbData,
	}
}

// GetDeviceConfig 查询保存在指定key的hash的field中的设备配置信息
func (r *Repo) GetDeviceConfig(key, field string) ([]byte, error) {
	result, err := r.redisClient.HGet(context.Background(), key, field).Result()
	if err != nil {
		return nil, errors.Newf(
			500, "Repo_Config_Error",
			"向redis发出配置查询请求时发生了错误:%v", err)
	}

	// 二进制信息统一以十六进制字符串的信息在redis中保存，因此需要转换
	var config []byte
	_, err = fmt.Sscanf(result, "%x", &config)
	if err != nil {
		return nil, errors.Newf(
			500, "Repo_Config_Error",
			"将配置信息的十六进制字符串转为二进制信息时发生了错误:%v", err)
	}

	return config, nil
}

// RunWarningDetectTask 依据预警字段注册的预警规则，创建并运行下采样设备状态信息数据的task
func (r *Repo) RunWarningDetectTask(config *biz.WarningDetectTaskConfig) (*domain.Run, error) {
	tasksAPI := r.influxdbClient.TasksAPI()

	var flux string
	if config.AggregateType != utilApi.DeviceStateRegisterInfo_NONE {
		fluxFormat := `
		data = from(bucket: "%s")
			|> range(start: -task.every)
			|> filter(fn: (r) => r["deviceClassID"] == "%d")
			|> filter(fn: (r) => r["_field"] == "%s")
		
		data
			|> aggregateWindow(every: task.every, fn: %s, createEmpty: false)
			|> to(bucket: "%s", org: "%s")`
		var fn string
		switch config.AggregateType {
		case utilApi.DeviceStateRegisterInfo_AVG:
			fn = "mean"
		case utilApi.DeviceStateRegisterInfo_MAX:
			fn = "max"
		case utilApi.DeviceStateRegisterInfo_MIN:
			fn = "min"
		case utilApi.DeviceStateRegisterInfo_SUM:
			fn = "sum"
		}
		flux = fmt.Sprintf(fluxFormat,
			conf.Username,
			config.DeviceClassID,
			config.FieldName,
			fn, config.TargetBucket, r.influxdbClient.org,
		)
	} else {
		fluxFormat := `
		data = from(bucket: "%s")
			|> range(start: -task.every)
			|> filter(fn: (r) => r["deviceClassID"] == "%d")
			|> filter(fn: (r) => r["_field"] == "%s")
		
		data
			|> to(bucket: "%s", org: "%s")`

		flux = fmt.Sprintf(fluxFormat,
			conf.Username,
			config.DeviceClassID,
			config.FieldName,
			config.TargetBucket, r.influxdbClient.org,
		)
	}

	task, err := tasksAPI.CreateTaskWithEvery(
		context.Background(),
		config.Name,
		flux,
		config.Every.String(),
		r.influxdbClient.orgId,
	)
	if err != nil {
		if task != nil {
			tasksAPI.DeleteTask(context.Background(), task)
		}
		return nil, errors.Newf(
			500, "Repo_State_Error",
			"启动influx task时发生了错误:%v", err)
	}

	return tasksAPI.RunManually(context.Background(), task)
}

// StopWarningDetectTask 停止运行指定的task
func (r *Repo) StopWarningDetectTask(run *domain.Run) error {
	if run == nil {
		return nil
	}
	tasksAPI := r.influxdbClient.TasksAPI()
	tasksAPI.CancelRun(context.Background(), run)
	err := tasksAPI.DeleteTaskWithID(context.Background(), *run.TaskID)
	if err != nil {
		return errors.Newf(
			500, "Repo_State_Error",
			"停止influx task的运行时发生了错误:%v", err)
	}
	return nil
}

// GetRecordCount 依据查询条件获取记录数
func (r *Repo) GetRecordCount(option *biz.QueryOption) (int64, error) {
	option.Limit = 0
	option.CountQuery = true
	flux := buildFluxQuery(option)

	queryAPI := r.influxdbClient.QueryAPI(r.influxdbClient.org)
	tableResult, err := queryAPI.Query(context.Background(), flux)
	if err != nil {
		return 0, errors.Newf(
			500, "Repo_State_Error",
			"查询记录数量时发生了错误:%v", err)
	}
	defer tableResult.Close()

	tableResult.Next()
	return tableResult.Record().Value().(int64), nil
}

// BatchGetDeviceStateInfo 批量查询某一类设备的状态信息
func (r *Repo) BatchGetDeviceStateInfo(deviceClassID int, option *biz.QueryOption) ([]*v1.DeviceState, error) {
	if option.Filter == nil {
		option.Filter = make(map[string]string)
	}

	// 填充查询的参数，并将关于deviceID的filter转换为对_measurement的filter
	if deviceID, ok := option.Filter["deviceID"]; ok {
		delete(option.Filter, "deviceID")
		option.Filter["_measurement"] = deviceID
	}

	// 查询设备状态信息时，以deviceClassID和_measurement即设备id以及时间_time作为group key
	option.GroupBy = append(option.GroupBy, "deviceClassID", "_measurement", "_time")

	// 添加关于deviceClassID的filter
	option.Filter["deviceClassID"] = strconv.Itoa(deviceClassID)

	queryApi := r.influxdbClient.QueryAPI(r.influxdbClient.org)
	tableResult, err := queryApi.Query(context.Background(), buildFluxQuery(option))
	if err != nil {
		return nil, errors.Newf(
			500, "Repo_State_Error",
			"批量查询设备状态信息时发生了错误:%v", err)
	}
	defer tableResult.Close()

	var result []*v1.DeviceState

	for i := 0; tableResult.Next(); i++ {
		pos := i / option.GroupCount
		if pos == -1 {
			break
		}
		// result达到容量上限，则需要扩容
		if pos == len(result) {
			result = append(result, &v1.DeviceState{
				Fields: make(map[string]float64),
				Tags:   make(map[string]string),
			})
		}

		record := tableResult.Record()
		if result[pos].DeviceId == "" {
			result[pos].DeviceId = record.Measurement()
			result[pos].DeviceClassId = int32(deviceClassID)
			// 数据库中的时间为utc时间，需要转换
			result[pos].Time = timestamppb.New(record.Time().Add(8 * time.Hour))
			// 解析tag
			for k, v := range record.Values() {
				// 除了系统字段、table字段以及deviceClassID字段，其余都视作tag
				if !strings.HasPrefix(k, "_") && k != "deviceClassID" && k != "table" {
					result[pos].Tags[k] = fmt.Sprintf("%v", v)
				}
			}
		}
		// 解析state的field
		field := record.ValueByKey("_field")
		if field != nil {
			result[pos].Fields[field.(string)] = record.Value().(float64)
		}
	}

	return result, nil
}

// BatchGetDeviceWarningDetectField 批量查询某一类设备某个字段的信息
func (r *Repo) BatchGetDeviceWarningDetectField(deviceClassID int, fieldName string, option *biz.QueryOption) (*api.QueryTableResult, error) {
	if option.Filter == nil {
		option.Filter = make(map[string]string)
	}

	// 添加关于deviceClassID和fieldName的filter
	option.Filter["deviceClassID"] = strconv.Itoa(deviceClassID)
	option.Filter["_field"] = fieldName

	queryApi := r.influxdbClient.QueryAPI(r.influxdbClient.org)
	tableResult, err := queryApi.Query(context.Background(), buildFluxQuery(option))
	if err != nil {
		return nil, errors.Newf(
			500, "Repo_State_Error",
			"查询设备状态字段信息时发生了错误:%v", err)
	}

	return tableResult, nil
}

func (r *Repo) GetWarningMessage(option *biz.QueryOption) ([]*v1.BatchGetWarningReply_Warning, error) {
	// 以设备id和设备字段名和设备类别号以及_time作为group key
	option.GroupBy = append(
		option.GroupBy, "deviceClassID", "_measurement", "deviceFieldName", "_time")

	queryApi := r.influxdbClient.QueryAPI(r.influxdbClient.org)
	tableResult, err := queryApi.Query(context.Background(), buildFluxQuery(option))
	if err != nil {
		return nil, errors.Newf(
			500, "Repo_State_Error",
			"批量查询警告信息时发生了错误:%v", err)
	}

	var warnings []*v1.BatchGetWarningReply_Warning
	for i := 0; tableResult.Next(); i++ {
		pos := i / option.GroupCount
		// 达到容量上限，则需要扩容
		if pos == len(warnings) {
			warnings = append(
				warnings, &v1.BatchGetWarningReply_Warning{
					Tags:   make(map[string]string),
					Fields: make(map[string]string),
				})
		}

		record := tableResult.Record()
		if warnings[pos].DeviceId == "" {
			// 解析tag字段，包括deviceId、deviceClassId、deviceFieldName
			deviceClassID, err := strconv.Atoi(record.ValueByKey("deviceClassID").(string))
			if err != nil {
				return nil, errors.Newf(
					500, "Repo_State_Error",
					"批量查询警告信息时发生了错误:%v", err)
			}
			warnings[pos].DeviceClassId = int32(deviceClassID)
			warnings[pos].DeviceId = record.Measurement()
			warnings[pos].Tags["deviceFieldName"] = record.ValueByKey("deviceFieldName").(string)
		}

		// 解析field，field包括start、end以及警告信息message
		field := record.ValueByKey("_field")
		value := record.Value()
		if field != nil {
			switch field.(string) {
			case "start":
				parse, err := time.Parse(time.RFC3339, value.(string))
				if err != nil {
					return nil, errors.Newf(
						500, "Repo_State_Error",
						"解析警告信息的start字段时发生了错误:%v", err)
				}
				warnings[pos].Start = timestamppb.New(parse.Add(8 * time.Hour))
			case "end":
				parse, err := time.Parse(time.RFC3339, value.(string))
				if err != nil {
					return nil, errors.Newf(
						500, "Repo_State_Error",
						"解析警告信息的end字段时发生了错误:%v", err)
				}
				warnings[pos].End = timestamppb.New(parse.Add(8 * time.Hour))
			case "message":
				warnings[pos].Fields["warningMessage"] = value.(string)
			}
		}
	}

	return warnings, nil
}

// SaveWarningMessage 保存警告信息，默认以警告信息的检测区间start为警告的发生时间
// 除了警告信息的message字段、start字段、end字段作为field，deviceID作为_measurement，其余的都作为tag
func (r *Repo) SaveWarningMessage(bucket string, warnings ...*utilApi.Warning) error {
	writeAPI := r.influxdbClient.WriteAPI(r.influxdbClient.org, bucket)
	defer writeAPI.Flush()

	for _, w := range warnings {
		start := w.Start.AsTime().UTC()
		end := w.End.AsTime().UTC()

		point := write.NewPointWithMeasurement(w.DeviceId)
		point.SetTime(w.Start.AsTime().UTC()).
			AddTag("deviceClassID", strconv.FormatInt(int64(w.DeviceClassId), 10)).
			AddTag("deviceFieldName", w.DeviceFieldName).
			AddField("start", start.Format(time.RFC3339)).
			AddField("end", end.Format(time.RFC3339)).
			AddField("message", w.WarningMessage).
			SortFields().
			SortTags()
		writeAPI.WritePoint(point)
	}

	return nil
}

func buildFluxQuery(option *biz.QueryOption) string {
	flux := fmt.Sprintf(`from(bucket: "%s")`, option.Bucket)

	// 选取一定时间范围内的数据
	if option.Past != 0 {
		rangeFilter := `|> range(start: -%s)`
		flux += fmt.Sprintf(rangeFilter, option.Past.String())
	} else if option.Start != nil {
		rangeFilter := `|> range(start: %s, stop: %s)`

		// 注意时间都需要使用UTC时间
		start := strconv.FormatInt(option.Start.UTC().Unix(), 10)
		stop := "now()"
		if option.Stop != nil {
			stop = strconv.FormatInt(option.Stop.UTC().Unix(), 10)
		}

		flux += fmt.Sprintf(rangeFilter, start, stop)
	} else {
		flux += `|> range(start: -30m)`
	}

	// 添加过滤条件
	if len(option.Filter) != 0 {
		filterFormat := `|> filter(fn: (r) => %s)`
		filters := make([]string, 0, len(option.Filter))

		for k, v := range option.Filter {
			format := `r["%s"] == "%s"`
			filters = append(filters, fmt.Sprintf(format, k, v))
		}

		flux += fmt.Sprintf(filterFormat, strings.Join(filters, " and "))
	}

	// 将分组条件进行格式转换
	group := make([]string, len(option.GroupBy))
	for i, g := range option.GroupBy {
		group[i] = fmt.Sprintf(`"%s"`, g)
	}

	// 以分组条件group进行排序，让逻辑上视为一条记录排列在一起
	flux += fmt.Sprintf(
		`|> group()|> sort(columns: [%s])`, strings.Join(group, ","))

	// 分页查询
	if option.Limit != 0 {
		// 利用group聚合，然后利用sort将逻辑上视为同一记录的信息排列在一起，
		// 从而可以利用limit函数进行分页，但是要注意，limit和offset要进行放缩
		// 分页
		flux += fmt.Sprintf(`|> limit(n: %d,offset: %d)`, option.Limit, option.Offset)
	}

	// 计数查询，直接利用group聚合后再利用count计数
	// 注意此时得到的记录数量是measurement的field的数量乘以实际的记录数，因此还要除以field的数量
	if option.CountQuery {
		flux += "|> group()|> count()"
	}
	//else {
	//	format := `|> group(columns: [%s])`
	//	flux += fmt.Sprintf(format, strings.Join(group, ", "))
	//}

	return flux
}
