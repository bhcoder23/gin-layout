package interfaces

import (
	"errors"

	"github.com/bhcoder23/gin-layout/internal/utils/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRepo[T IModel] interface {
	PageList(c *gin.Context, query *IFilter) (res *response.PageListT[T], err error)
	PageListWithSelectOption(c *gin.Context, query *IFilter, selectOpt []string) (res *response.PageListT[T], err error)
	One(c *gin.Context, id uint) (res T, err error)
	OneWithSelectOption(c *gin.Context, id uint, selectOpt []string) (res T, err error)
	OneByName(c *gin.Context, name string) (res T, err error)
	OneByNameWithSelectOption(c *gin.Context, name string, selectOpt []string) (res T, err error)
	Add(c *gin.Context, model T) (newID uint, err error)
	Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error)
	Delete(c *gin.Context, id uint) (deleted bool, err error)
}

type Repo[T IModel] struct {
	DB *gorm.DB
}

func NewRepo[T IModel](db *gorm.DB) *Repo[T] {
	return &Repo[T]{
		DB: db,
	}
}

func (r *Repo[T]) PageList(c *gin.Context, f *IFilter) (res *response.PageListT[T], err error) {
	db := r.DB
	db = (*f).BuildPageListFilter(c, db)
	offset := ((*f).GetPage() - 1) * (*f).GetPageSize()
	db = db.Model(new(T)).Offset(int(offset)).Limit(int((*f).GetPageSize()))
	objs := make([]T, 0)
	err = db.Find(&objs).Error
	var count int64
	db.Offset(-1).Limit(-1).Select("count(id)").Count(&count)

	res = &response.PageListT[T]{
		List:  objs,
		Pages: response.MakePages(count, (*f).GetPage(), (*f).GetPageSize()),
	}

	return
}

func (r *Repo[T]) PageListWithSelectOption(c *gin.Context, f *IFilter, selectOpt []string) (res *response.PageListT[T], err error) {
	db := r.DB
	db = (*f).BuildPageListFilter(c, db)
	offset := ((*f).GetPage() - 1) * (*f).GetPageSize()
	db = db.Model(new(T)).Offset(int(offset)).Limit(int((*f).GetPageSize()))
	if len(selectOpt) > 0 {
		db = db.Select(selectOpt)
	}
	objs := make([]T, 0)
	err = db.Find(&objs).Error
	var count int64
	db.Offset(-1).Limit(-1).Select("count(id)").Count(&count)

	res = &response.PageListT[T]{
		List:  objs,
		Pages: response.MakePages(count, (*f).GetPage(), (*f).GetPageSize()),
	}

	return
}

func (r *Repo[T]) One(c *gin.Context, id uint) (res T, err error) {
	db := r.DB
	err = db.Model(new(T)).Where(`id=?`, id).First(&res).Error
	return
}

func (r *Repo[T]) OneWithSelectOption(c *gin.Context, id uint, selectOpt []string) (res T, err error) {
	db := r.DB
	db = db.Model(new(T)).Where(`id=?`, id)
	if len(selectOpt) > 0 {
		db = db.Select(selectOpt)
	}
	err = db.First(&res).Error
	return
}

func (r *Repo[T]) OneByName(c *gin.Context, name string) (res T, err error) {
	db := r.DB
	err = db.Model(new(T)).Where(`name=?`, name).First(&res).Error
	return
}

func (r *Repo[T]) OneByNameWithSelectOption(c *gin.Context, name string, selectOpt []string) (res T, err error) {
	db := r.DB
	db = db.Model(new(T)).Where(`name=?`, name)
	if len(selectOpt) > 0 {
		db = db.Select(selectOpt)
	}
	err = db.First(&res).Error
	return
}

func (r *Repo[T]) Add(c *gin.Context, model T) (newID uint, err error) {
	db := r.DB
	err = db.Create(model).Error
	newID = model.GetID()
	return
}

func (r *Repo[T]) Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error) {
	if id <= 0 {
		updated = false
		err = errors.New(`missing id`)
		return
	}
	_, err = r.One(c, id)
	if err != nil {
		return
	}
	db := r.DB
	err = db.Model(new(T)).Omit(`created_at`).Where(`id=?`, id).Updates(updateFields).Error
	if err == nil {
		updated = true
	}
	return
}

func (r *Repo[T]) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	if id <= 0 {
		deleted = false
		err = errors.New(`missing id`)
		return
	}
	db := r.DB
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	err = db.Model(new(T)).Where(`id=?`, id).Delete(&model).Error
	if err == nil {
		deleted = true
	}
	return
}
