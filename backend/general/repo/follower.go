package repo

import (
	"general/entities"
	"general/infra"
	"general/types"
)

var _ Follower = (*FollowerImpl)(nil)

type FollowerImpl struct {
	db infra.RDB
}

func (r *FollowerImpl) Run() (err error) {
	return
}

func (r *FollowerImpl) Stop() {}

func (r *FollowerImpl) Create(userId, followerId uint64) (err error) {
	err = r.db.Orm().Create(
		&entities.Follower{
			UserId:     userId,
			FollowerId: followerId,
		},
	).Error
	return
}

func (r *FollowerImpl) GetByUserIdAndFollowerId(userId, followerId uint64) (follower entities.Follower, err error) {
	err = r.db.Orm().
		Where("user_id = ? and follower_id = ?", userId, followerId).
		Take(&follower).Error
	return
}

func (r *FollowerImpl) GetFollowersByUserId(userId uint64) (followers []types.Follower, err error) {
	err = r.db.Orm().
		Table("followers f").
		Select("f.follower_id, u.username, i.file_name").
		Joins("inner join users u on u.user_id = f.follower_id").
		Joins("left join images i on i.user_id = u.user_id").
		Where("f.user_id = ?", userId).Find(&followers).Error
	return
}

func (r *FollowerImpl) GetUsersByFollowerId(followerId uint64) (users []types.Follower, err error) {
	err = r.db.Orm().
		Table("followers f").
		Select("f.user_id, u.username, i.file_name").
		Joins("inner join users u on u.user_id = f.user_id").
		Joins("left join images i on i.user_id = u.user_id").
		Where("f.follower_id = ?", followerId).Find(&users).Error
	return
}

func (r *FollowerImpl) DeleteByUserIdAndFollowerId(userId, followerId uint64) (err error) {
	err = r.db.Orm().
		Where("user_id = ? and follower_id = ?", userId, followerId).
		Delete(new(entities.Follower)).Error
	return
}
