package repo

import (
	"errors"
	"fmt"
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
	err = r.db.DB().QueryRow("select password from users where user_id = ? ", userId).Scan(&hashedPassword)
	return
}

func (r *ProfileImpl) UpdatePassword(userId uint64, password string) (int64, error) {
	sqlRes, err := r.db.DB().Exec("update users set password = ? where user_id = ?", password, userId)
	if err != nil {
		return 0, err
	}

	return sqlRes.RowsAffected()
}

func (r *ProfileImpl) UpsertProfileImageInfo(userId uint64, fileName string, desc string) error {
	var count int64
	if err := r.db.DB().QueryRow("select count(user_id) from images where user_id = ?", userId).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		_, err := r.db.DB().Exec("update images set file_name = ?, descript = ? where user_id = ?", fileName, desc, userId)
		return err
	}

	_, err := r.db.DB().Exec("insert into images (user_id, file_name, descript) values (?, ?, ?)", userId, fileName, desc)
	return err
}

func (r *ProfileImpl) UpdateColumnsById(data interface{}, id uint64) (int64, error) {
	cols, args := utils.ConstructParamsFromStruct(data)
	if cols == "" {
		return 0, ErrInvalidColumns
	}

	stmt := fmt.Sprintf("update users set %s where user_id = ?", cols)
	args = append(args, id)
	sqlRes, err := r.db.DB().Exec(stmt, args...)
	if err != nil {
		return 0, err
	}

	return sqlRes.RowsAffected()
}
