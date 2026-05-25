package interfaces

import (
	"github.com/bhcoder23/gin-layout/internal/utils/response"
	"github.com/gin-gonic/gin"
)

type IService[T IModel] interface {
	PageList(c *gin.Context, filter *IFilter) (res *response.PageListT[T], err error)
	PageListWithSelectOption(c *gin.Context, filter *IFilter, selectOpt []string) (res *response.PageListT[T], err error)
	One(c *gin.Context, id uint) (res T, err error)
	OneWithSelectOption(c *gin.Context, id uint, selectOpt []string) (res T, err error)
	OneByName(c *gin.Context, name string) (res T, err error)
	OneByNameWithSelectOption(c *gin.Context, name string, selectOpt []string) (res T, err error)
	Add(c *gin.Context, model T) (newID uint, err error)
	Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error)
	Delete(c *gin.Context, id uint) (deleted bool, err error)
}

type Service[T IModel] struct {
	Repo *IRepo[T]
}

func NewService[T IModel](r IRepo[T]) *Service[T] {
	return &Service[T]{
		Repo: &r,
	}
}

func (s *Service[T]) PageList(c *gin.Context, filter *IFilter) (res *response.PageListT[T], err error) {
	repo := *s.Repo
	return repo.PageList(c, filter)
}

func (s *Service[T]) PageListWithSelectOption(c *gin.Context, filter *IFilter, selectOpt []string) (res *response.PageListT[T], err error) {
	repo := *s.Repo
	return repo.PageListWithSelectOption(c, filter, selectOpt)
}

func (s *Service[T]) One(c *gin.Context, id uint) (res T, err error) {
	repo := *s.Repo
	return repo.One(c, id)
}

func (s *Service[T]) OneWithSelectOption(c *gin.Context, id uint, selectOpt []string) (res T, err error) {
	repo := *s.Repo
	return repo.OneWithSelectOption(c, id, selectOpt)
}

func (s *Service[T]) OneByName(c *gin.Context, name string) (res T, err error) {
	repo := *s.Repo
	return repo.OneByName(c, name)
}

func (s *Service[T]) OneByNameWithSelectOption(c *gin.Context, name string, selectOpt []string) (res T, err error) {
	repo := *s.Repo
	return repo.OneByNameWithSelectOption(c, name, selectOpt)
}

func (s *Service[T]) Add(c *gin.Context, model T) (newID uint, err error) {
	repo := *s.Repo
	return repo.Add(c, model)
}

func (s *Service[T]) Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error) {
	delete(updateFields, `created_at`)
	delete(updateFields, `updated_at`)
	delete(updateFields, `deleted_at`)
	repo := *s.Repo
	return repo.Update(c, updateFields, id)
}

func (s *Service[T]) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	repo := *s.Repo
	return repo.Delete(c, id)
}
