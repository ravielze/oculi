package dto

import (
	"github.com/ravielze/oculi/common/constant/format"
	"github.com/ravielze/oculi/common/model/dao"
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
)

func NewBaseModel(item dao.BaseModel) BaseModel {
	return BaseModel{
		CreatedAt: item.CreatedAt.Format(format.DATETIME),
		UpdatedAt: item.UpdatedAt.Format(format.DATETIME),
	}
}

func NewBaseModelSoftDelete(item dao.BaseModelSoftDelete) BaseModelSoftDelete {
	result := BaseModelSoftDelete{
		BaseModel: NewBaseModel(item.BaseModel),
		DeletedAt: "",
	}
	if item.DeletedAt.Valid {
		result.DeletedAt = item.DeletedAt.Time.Format(format.DATETIME)
	}
	return result
}
