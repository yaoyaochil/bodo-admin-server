package system

import (
	"errors"

	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system/request"
	"gorm.io/gorm"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: DeleteSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type DictionaryService struct{}

func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(global.WUSHI_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.WUSHI_DB.Create(&sysDictionary).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: DeleteSysDictionary
//@description: 删除字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = global.WUSHI_DB.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.WUSHI_DB.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}

	if sysDictionary.SysDictionaryDetails != nil {
		return global.WUSHI_DB.Where("sys_dictionary_id=?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: UpdateSysDictionary
//@description: 更新字典数据
//@param: sysDictionary *model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := global.WUSHI_DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Type != sysDictionary.Type {
		if !errors.Is(global.WUSHI_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = db.Updates(sysDictionaryMap).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetSysDictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDictionary model.SysDictionary

func (dictionaryService *DictionaryService) GetSysDictionary(Type string, Id uint, status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = global.WUSHI_DB.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).Preload("SysDictionaryDetails", "status = ?", true).First(&sysDictionary).Error
	return
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetSysDictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (dictionaryService *DictionaryService) GetSysDictionaryInfoList(info request.SysDictionarySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.WUSHI_DB.Model(&system.SysDictionary{})
	var sysDictionarys []system.SysDictionary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return sysDictionarys, total, err
}
