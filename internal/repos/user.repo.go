package repos

import (
	"github.com/bhcoder23/gin-layout/internal/components"
	"github.com/bhcoder23/gin-layout/internal/interfaces"
	"github.com/bhcoder23/gin-layout/internal/models"
)

var User *UserRepo

type UserRepo struct {
	interfaces.Repo[*models.User]
}

func NewUserRepo() {
	User = &UserRepo{
		Repo: *interfaces.NewRepo[*models.User](components.DB),
	}
}

func init() {
	RegisterRepos(NewUserRepo)
}
