package repo

import (
	"errors"
	"security/entities"
	"security/infra"
	"security/utils"
)

var _ Profile = (*ProfileImpl)(nil)

var ErrInvalidColumns = errors.New("column name invalid")

type ProfileImpl struct {
	db infra.RDB
}

func NewProfile(db infra.RDB) *ProfileImpl {
	return &ProfileImpl{
		db: db,
	}
}

func (r *ProfileImpl) GetPasswordByUserId(userId uint64) (hashedPassword string, err error) {
	var user entities.User
	err = r.db.Orm().
		Select("password").
		Where("user_id = ?", userId).
		First(&user).Error
	hashedPassword = user.Password
	return
}

func (r *ProfileImpl) UpdatePassword(userId uint64, password string) (int64, error) {
	result := r.db.Orm().
		Model(new(entities.User)).
		Where("user_id = ?", userId).
		Update("password", password)
	return result.RowsAffected, result.Error
}

func (r *ProfileImpl) UpsertProfileImageInfo(userId uint64, fileName string, desc string) error {
	image := entities.Image{
		UserId:   userId,
		FileName: fileName,
		Descript: desc,
	}

	return r.db.Orm().
		Table(image.TableName()).
		Where("user_id = ?", userId).
		Assign(map[string]interface{}{
			"file_name": fileName,
			"descript":  desc,
		}).
		FirstOrCreate(&image).Error
}

func (r *ProfileImpl) UpdateColumnsById(data interface{}, id uint64) (int64, error) {
	updateMap := utils.ConstructMapFromStruct(data)
	if len(updateMap) == 0 {
		return 0, ErrInvalidColumns
	}

	result := r.db.Orm().
		Model(new(entities.User)).
		Where("user_id = ?", id).
		Updates(updateMap)
	return result.RowsAffected, result.Error
}
