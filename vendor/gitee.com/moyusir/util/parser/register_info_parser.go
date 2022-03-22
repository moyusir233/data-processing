package parser

import (
	v1 "gitee.com/moyusir/util/api/util/v1"
)

// RegisterInfoParser 注册信息解析
type RegisterInfoParser struct {
	// 依据注册信息建立的，对应各个设备类别预警字段相关信息
	warningDetectInfo [][]WarningDetectField
	// 注册信息
	Info []v1.DeviceStateRegisterInfo
}

// WarningDetectField 预警字段
type WarningDetectField struct {
	// 预警字段名
	Name string
	// 字段的预警函数
	Func WarningFunc
	// 字段完整的预警规则
	Rule *v1.DeviceStateRegisterInfo_WarningRule
}

// NewRegisterInfoParser 构造器函数
func NewRegisterInfoParser(registerInfo []v1.DeviceStateRegisterInfo) (*RegisterInfoParser, error) {
	info := make([][]WarningDetectField, 0, len(registerInfo))

	// 处理注册信息
	for _, r := range registerInfo {
		var fields []WarningDetectField
		for _, f := range r.Fields {
			// 将预警字段信息处理为方便预警检测服务使用的形式
			if f.WarningRule != nil {
				warningFunc, err := GetWarningFunc(f.Type, f.WarningRule.CmpRule)
				if err != nil {
					return nil, err
				}
				wf := WarningDetectField{
					Name: f.Name,
					Func: warningFunc,
					Rule: f.WarningRule,
				}
				fields = append(fields, wf)
			}
		}
		info = append(info, fields)
	}
	return &RegisterInfoParser{warningDetectInfo: info, Info: registerInfo}, nil
}

// GetWarningDetectFieldNames 获得指定设备类别号下所有需要进行预警规则的字段名
func (p *RegisterInfoParser) GetWarningDetectFieldNames(deviceClassID int) (fields []string) {
	for _, f := range p.warningDetectInfo[deviceClassID] {
		fields = append(fields, f.Name)
	}
	return
}

// GetWarningDetectFields 获得指定设备类别号下所有的预警字段信息
func (p *RegisterInfoParser) GetWarningDetectFields(deviceClassID int) []WarningDetectField {
	return p.warningDetectInfo[deviceClassID]
}
