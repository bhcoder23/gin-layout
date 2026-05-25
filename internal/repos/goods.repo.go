package repos

import (
	"github.com/bhcoder23/gin-layout/internal/components"
	"github.com/bhcoder23/gin-layout/internal/interfaces"
	"github.com/bhcoder23/gin-layout/internal/models"
)

var Goods *GoodsRepo

type GoodsRepo struct {
	interfaces.Repo[*models.Goods]
}

func NewGoodsRepo() {
	Goods = &GoodsRepo{
		Repo: *interfaces.NewRepo[*models.Goods](components.DB),
	}
}

func init() {
	RegisterRepos(NewGoodsRepo)
}

func (r *GoodsRepo) GetByID(id uint) (res models.Goods, err error) {
	err = r.Repo.DB.Where(`id=?`, id).First(&res).Error
	if err != nil {
		return
	}
	return
}
