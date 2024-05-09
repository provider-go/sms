package models

import (
	"github.com/provider-go/sms/global"
	"time"
)

type SMSLog struct {
	Id          int32     `json:"id" gorm:"auto_increment;primary_key;comment:主键"`
	Type        string    `json:"type" gorm:"column:type;type:varchar(20);not null;default:'';comment:短信类型"`
	Supplier    string    `json:"supplier" gorm:"column:supplier;type:varchar(20);not null;default:'';comment:短信供应商"`
	CountryCode string    `json:"countryCode" gorm:"column:country_code;type:varchar(10);not null;default:'0086';comment:国家编码"`
	Phone       string    `json:"phone" gorm:"column:phone;type:varchar(20);not null;default:'';comment:电话号码"`
	Message     string    `json:"message" gorm:"column:message;type:varchar(200);not null;default:'';comment:短信内容"`
	CreateTime  time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime  time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateSMSLog(smsType, supplier, phone, message string) error {
	return global.DB.Table("sms_logs").Create(&SMSLog{Type: smsType, Supplier: supplier, Phone: phone, Message: message}).Error
}

func ListSMSLog(pageSize, pageNum int) ([]*SMSLog, int64, error) {
	var rows []*SMSLog
	//计算列表数量
	var count int64
	global.DB.Table("sms_logs").Count(&count)

	if err := global.DB.Table("sms_logs").Order("id desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}
