package dto

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ravielze/oculi/common/model/dao"
	"github.com/ravielze/oculi/constant/oculi_time"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewBaseModel(t *testing.T) {
	t.Run("Base model", func(t *testing.T) {
		oculi_time.Mock(time.Date(2021, time.June, 19, 2, 58, 30, 0, time.Local))
		baseModel := dao.BaseModel{
			CreatedAt: oculi_time.Now(),
			UpdatedAt: oculi_time.Now().Add(time.Hour * 1),
		}
		converted := NewBaseModel(baseModel)
		assert.Equal(t, "19-06-2021 02:58:30", converted.CreatedAt)
		assert.Equal(t, "19-06-2021 03:58:30", converted.UpdatedAt)
		buff, _ := json.MarshalIndent(converted, "", "\t")
		assert.Equal(t, string(buff),
			`{
	"created_at": "19-06-2021 02:58:30",
	"updated_at": "19-06-2021 03:58:30"
}`)
	})

	t.Run("Base model soft delete, not deleted", func(t *testing.T) {
		oculi_time.Mock(time.Date(2021, time.June, 19, 2, 38, 30, 0, time.Local))
		baseModel := dao.BaseModelSoftDelete{
			BaseModel: dao.BaseModel{
				CreatedAt: oculi_time.Now(),
				UpdatedAt: oculi_time.Now().Add(time.Hour * 1),
			},
			DeletedAt: gorm.DeletedAt{
				Valid: false,
				Time:  time.Time{},
			},
		}
		converted := NewBaseModelSoftDelete(baseModel)
		assert.Equal(t, "", converted.DeletedAt)
		assert.Equal(t, "19-06-2021 02:38:30", converted.CreatedAt)
		assert.Equal(t, "19-06-2021 03:38:30", converted.UpdatedAt)
		buff, _ := json.MarshalIndent(converted, "", "\t")
		assert.Equal(t, string(buff),
			`{
	"created_at": "19-06-2021 02:38:30",
	"updated_at": "19-06-2021 03:38:30",
	"deleted_at": ""
}`)
	})

	t.Run("Base model soft delete, deleted", func(t *testing.T) {
		oculi_time.Mock(time.Date(2021, time.June, 19, 2, 38, 30, 0, time.Local))
		baseModel := dao.BaseModelSoftDelete{
			BaseModel: dao.BaseModel{
				CreatedAt: oculi_time.Now(),
				UpdatedAt: oculi_time.Now().Add(time.Hour * 1),
			},
			DeletedAt: gorm.DeletedAt{
				Valid: true,
				Time:  oculi_time.Now().Add(time.Minute * 140),
			},
		}
		converted := NewBaseModelSoftDelete(baseModel)
		assert.Equal(t, "19-06-2021 04:58:30", converted.DeletedAt)
		assert.Equal(t, "19-06-2021 02:38:30", converted.CreatedAt)
		assert.Equal(t, "19-06-2021 03:38:30", converted.UpdatedAt)
		buff, _ := json.MarshalIndent(converted, "", "\t")
		assert.Equal(t, string(buff),
			`{
	"created_at": "19-06-2021 02:38:30",
	"updated_at": "19-06-2021 03:38:30",
	"deleted_at": "19-06-2021 04:58:30"
}`)
	})
}
