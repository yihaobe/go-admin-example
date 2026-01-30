package service

import (
	"errors"
	"go-admin/app/user/models"
	"go-admin/app/user/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type User struct {
	service.Service
}

// GetPage 获取User列表
func (e *User) GetPage(c *dto.UserGetPageReq, p *actions.DataPermission, list *[]models.User, count *int64) error {
	var err error
	var data models.User

	err = e.Orm.Debug().Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetUserPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取User对象
func (e *User) Get(d *dto.UserGetReq, p *actions.DataPermission, model *models.User) error {
	var data models.User

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建User对象
func (e *User) Insert(c *dto.UserInsertReq) error {
    var err error
    var data models.User
    c.Generate(&data)
    err = e.Orm.Create(&data).Error
    if err != nil {
        e.Log.Errorf("db error:%s", err)
        return err
    }
    return nil
}

// Update 修改User对象
func (e *User) Update(c *dto.UserUpdateReq, p *actions.DataPermission) error {
	var err error
	var model = models.User{}
	e.Orm.Debug().First(&model, c.GetId())
	c.Generate(&model)
	
	db := e.Orm.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateUser error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除User
func (e *User) Remove(d *dto.UserDeleteReq, p *actions.DataPermission) error {
	var data models.User

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveUser error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
