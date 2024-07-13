package dto

import "matchlove-services/internal/model"

type RecordLoginActivityDto struct {
	LoginActivity *model.LoginActivity
	DevicesInfo   *model.DevicesInfo
}
