package dto

import (
	"github.com/ravielze/oculi/common/model/dao"
	time "github.com/ravielze/oculi/constant/time"
)

type (
	BaseModel struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	BaseModelSoftDelete struct {
		BaseModel
		DeletedAt string `json:"deleted_at"`
	}

	Map map[string]interface{}
)

func (m Map) ToMap() map[string]interface{} {
	return m
}

func NewBaseModel(item dao.BaseModel) BaseModel {
	return BaseModel{
		CreatedAt: item.CreatedAt.Format(time.DATETIME_LAYOUT),
		UpdatedAt: item.UpdatedAt.Format(time.DATETIME_LAYOUT),
	}
}

func NewBaseModelSoftDelete(item dao.BaseModelSoftDelete) BaseModelSoftDelete {
	result := BaseModelSoftDelete{
		BaseModel: NewBaseModel(item.BaseModel),
		DeletedAt: "",
	}
	if item.DeletedAt.Valid {
		result.DeletedAt = item.DeletedAt.Time.Format(time.DATETIME_LAYOUT)
	}
	return result
}
