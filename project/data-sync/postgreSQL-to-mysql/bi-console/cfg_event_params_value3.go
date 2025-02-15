package main

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
)

// CfgEventParamsValuePG3 From PostgreSQL fotoabledb data_cfg.cfg_event_params_value
type CfgEventParamsValuePG3 struct {
	//   "app_id" varchar(2000),
	//  "params" varchar(2000),
	//  "params_value" varchar(2000),
	//  "params_label" varchar(2000),
	//  "person" varchar(2000),
	//  "remark" varchar(200),
	//  "create_time" timestamp(0),
	//  "update_time" timestamp(0)
	AppID       string `gorm:"column:app_id;NOT NULL;" json:"app_id"`
	Params      string `gorm:"column:params;NOT NULL;" json:"params"`
	ParamsValue string `gorm:"column:params_value;NOT NULL;" json:"params_value"`
	ParamsLabel string `gorm:"column:params_label;NOT NULL;" json:"params_label"`
	Person      string `gorm:"column:person;NOT NULL;" json:"person"`
	Remark      string `gorm:"column:remark;NOT NULL;" json:"remark"`
	CreateTime  string `gorm:"column:create_time;NOT NULL;" json:"create_time"`
	UpdateTime  string `gorm:"column:update_time;NOT NULL;" json:"update_time"`
}

// CfgEventParamsValue3 From MySQL bi_console cfg_event_params_value
type CfgEventParamsValue3 struct {
	//  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
	//  `app_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用id',
	//  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '参数名称',
	//  `params_value` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '参数值',
	//  `params_label` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '参数值名称',
	//  `creator` int NOT NULL COMMENT '创建人id',
	//  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
	//  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID       string `gorm:"column:app_id;NOT NULL;" json:"app_id"`
	Params      string `gorm:"column:params;NOT NULL;" json:"params"`
	ParamsValue string `gorm:"column:params_value;NOT NULL;" json:"params_value"`
	ParamsLabel string `gorm:"column:params_label;NOT NULL;" json:"params_label"`
	Creator     int32  `gorm:"column:creator;NOT NULL;" json:"creator"`
	CreatedAt   int32  `gorm:"column:created_at;NOT NULL;" json:"created_at"`
	UpdatedAt   int32  `gorm:"column:updated_at;NOT NULL;" json:"updated_at"`
}

func FunCfgEventParamsValue3() {

	// 1、pg查数据
	cfgEventParamsValuePG := make([]*CfgEventParamsValuePG3, 0)
	err := db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").
		Find(&cfgEventParamsValuePG).Error
	if err != nil {
		fmt.Println("RunCfgEventParamsValue PostgreSQLClient Find err:", err)
		return
	}

	if len(cfgEventParamsValuePG) == 0 {
		fmt.Println("No more data to migrate.")
		return
	}

	// 2、转换数据并存入MySQL
	for i, v := range cfgEventParamsValuePG {
		cfgEventParamsValue := &CfgEventParamsValue3{
			AppID:       v.AppID,
			Params:      v.Params,
			ParamsValue: v.ParamsValue,
			ParamsLabel: v.ParamsLabel,
		}
		// 3、mysql存数据
		err = db2.MySQLClientBI.Table("cfg_event_params_value").Create(cfgEventParamsValue).Error
		if err != nil {
			fmt.Println("RunCfgEventParamsValue MySQLClientBI CreateInBatches err:", err)
			return
		}
		fmt.Println("第 ", i, " 条数据迁移完成")
	}
}

func main() {
	FunCfgEventParamsValue3()
	fmt.Println("Migration complete!")
}
