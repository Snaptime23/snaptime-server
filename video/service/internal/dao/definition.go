package dao

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao/model"
)

func CreateVideoDefinition(ctx context.Context, definition *model.Definition) (err error) {
	tx := DB.Begin().WithContext(ctx)
	if err = tx.Create(definition).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}
