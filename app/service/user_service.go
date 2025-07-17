package service

import (
	"errors"
	"fmt"
	"janx-admin/app/model"
	"janx-admin/app/vo"
	"janx-admin/global"
	"janx-admin/pkg/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type UserServiceInterface interface {
	Create(*model.User) error
	Update(id uint, user *model.User) error
	Delete(ids vo.UserDeleteReq) error
	List(r *vo.UserListReq) (list []*model.User, total int64, err error)
	ValidateUser(username string, password string) (model.User, error)
	GetCurrentUser(c *gin.Context) (model.User, error)
	GetUserById(id uint) (model.User, error)
}

type UserService struct {
}

var userServiceCache = cache.New(24*time.Hour, 48*time.Hour) // 缓存2小时

func NewUserService() UserServiceInterface {
	return &UserService{}
}

func (u *UserService) Create(user *model.User) error {
	users := []model.User{}
	global.Db.Where("username = ? or phone = ?", user.Username, user.Phone).Find(&users)
	if len(users) > 0 {
		return errors.New("用户已存在(用户名或手机号已存在)，请勿重新添加")
	}
	err := global.Db.Create(&user).Error
	return err
}

func (u *UserService) Update(id uint, user *model.User) error {
	err := global.Db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error
	return err
}

func (u *UserService) Delete(ids vo.UserDeleteReq) error {
	err := global.Db.Where("id in (?)", ids.Ids).Delete(&model.User{}).Error
	return err
}

func (u *UserService) List(r *vo.UserListReq) (list []*model.User, total int64, err error) {
	db := global.Db.Model(&model.User{}).Order("created_at desc")
	if strings.TrimSpace(r.Username) != "" {
		db.Where("username like ?", "%"+r.Username+"%").Find(&list)
	}
	if strings.TrimSpace(r.Nickname) != "" {
		db.Where("nickname like ?", "%"+r.Nickname+"%").Find(&list)
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(r.PageNum)
	pageSize := int(r.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}

func (u *UserService) ValidateUser(username string, password string) (model.User, error) {
	user := model.User{}
	err := global.Db.Where("username = ?", username).First(&user).Error
	if user.ID <= 0 {
		return user, errors.New("用户不存在")
	}

	if err != nil {
		return user, fmt.Errorf("获取用户错误：%s", err.Error())
	}
	if user.Status == 0 {
		return user, errors.New("用户已被禁用，请联系管理员")
	}
	ok := utils.CheckPasswordHash(password, user.Password)
	if !ok {
		return user, errors.New("用户名或密码错误")
	}
	return user, nil
}

func (u *UserService) GetUserById(id uint) (model.User, error) {
	user := model.User{}
	err := global.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *UserService) GetCurrentUser(c *gin.Context) (model.User, error) {
	var newUser model.User
	ctxUser, exist := c.Get("user")
	if !exist {
		return newUser, errors.New("获取用户信息失败")
	}
	us, _ := ctxUser.(model.User)

	cacheUser, found := userServiceCache.Get(us.Username)
	var user model.User
	var err error
	if found {
		user = cacheUser.(model.User)
		err = nil
	} else {
		user, err = u.GetUserById(us.ID)
		if err != nil {
			userServiceCache.Delete(us.Username)
		} else {
			userServiceCache.Set(us.Username, user, cache.DefaultExpiration)
		}
	}

	return user, err
}
