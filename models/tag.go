package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var tags []Tag
	var err error
	if pageNum > 0 && pageSize > 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return tag.ID > 0, nil
}

func AddTag(name string, state int, createdBy string) error {
	err := db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? and deleted_on = ?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return tag.ID > 0, nil
}

func DeleteTag(id int) error {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ? and deleted_on = ?", id, 0).Update(data).Error
	if err != nil {
		return err
	}
	return nil
}

//硬删除
func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})
	return true
}
