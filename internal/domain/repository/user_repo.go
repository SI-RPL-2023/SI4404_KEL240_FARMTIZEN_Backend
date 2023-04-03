package repository

import (
	"errors"

	"github.com/SI-RPL-2023/SI4404_KEL240_FARMTIZEN_Backend/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entity.User) error {
	if r.db.Create(user).Error != nil {
		return r.db.Create(user).Error
	}
	return errors.New("failed to create user")
}

func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if r.db.Where("email = ?", email).First(&user).Error != nil {
		return nil, r.db.Where("email = ?", email).First(&user).Error
	}
	return &user, nil
}

func (r *UserRepo) FindByID(id string) (*entity.User, error) {
	var user entity.User
	if r.db.Where("id = ?", id).First(&user).Error != nil {
		return nil, r.db.Where("id = ?", id).First(&user).Error
	}
	return &user, nil
}

func (r *UserRepo) FindAll() ([]entity.User, error) {
	var users []entity.User
	if r.db.Find(&users).Error != nil {
		return nil, r.db.Find(&users).Error
	}
	return users, nil
}

func (r *UserRepo) Update(user *entity.User) error {
	if r.db.Save(user).Error != nil {
		return r.db.Save(user).Error
	}
	return errors.New("failed to update user")
}

func (r *UserRepo) Delete(user *entity.User) error {
	if r.db.Delete(user).Error != nil {
		return r.db.Delete(user).Error
	}
	return errors.New("failed to delete user")
}

func (r *UserRepo) FindUserTokenByToken(token string) (*entity.UserToken, error) {
	var userToken entity.UserToken
	if r.db.Where("token = ?", token).First(&userToken).Error != nil {
		return nil, r.db.Where("token = ?", token).First(&userToken).Error
	}
	return &userToken, nil
}

func (r *UserRepo) FindUserTokenByUserID(userID string) (*entity.UserToken, error) {
	var userToken entity.UserToken
	if r.db.Where("user_id = ?", userID).First(&userToken).Error != nil {
		return nil, r.db.Where("user_id = ?", userID).First(&userToken).Error
	}
	return &userToken, nil
}

func (r *UserRepo) CreateUserToken(userToken *entity.UserToken) error {
	if r.db.Create(userToken).Error != nil {
		return r.db.Create(userToken).Error
	}
	return errors.New("failed to create user token")
}

func (r *UserRepo) DeleteUserToken(userToken *entity.UserToken) error {
	if r.db.Delete(userToken).Error != nil {
		return r.db.Delete(userToken).Error
	}
	return errors.New("failed to delete user token")
}

func (r *UserRepo) FindRoleByName(name string) (*entity.Role, error) {
	var role entity.Role
	if r.db.Where("name = ?", name).First(&role).Error != nil {
		return nil, r.db.Where("name = ?", name).First(&role).Error
	}
	return &role, nil
}

func (r *UserRepo) FindRoleByID(id string) (*entity.Role, error) {
	var role entity.Role
	if r.db.Where("id = ?", id).First(&role).Error != nil {
		return nil, r.db.Where("id = ?", id).First(&role).Error
	}
	return &role, nil
}
