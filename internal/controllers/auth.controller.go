// Package controllers defines HTTP handlers.
package controllers

import (
	"net/http"

	"github.com/bhcoder23/gin-layout/internal/components"
	"github.com/bhcoder23/gin-layout/internal/errors"
	"github.com/bhcoder23/gin-layout/internal/models"
	"github.com/bhcoder23/gin-layout/internal/services"
	"github.com/bhcoder23/gin-layout/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type inputUser struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Passwd   string `json:"passwd" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	input := inputUser{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			`error`: err.Error(),
		})
		return
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(input.Passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: "failed to hash password",
		})
		return
	}
	user := models.User{
		Username: input.Username,
		PassWd:   string(hashedPwd),
	}
	if err := components.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		`message`: "user created successfully",
		`user`: gin.H{
			`id`:       user.ID,
			`username`: user.Username,
		},
	})
}

// Login authenticates a user and returns a JWT.
func Login(c *gin.Context) {
	input := inputUser{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			`error`: err.Error(),
		})
		return
	}
	var user models.User
	err := components.DB.Where(`username = ?`, input.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PassWd), []byte(input.Passwd))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func Me(c *gin.Context) {
	user, err := services.User.GetByID(c)
	if err != nil {
		if biz, ok := err.(*errors.BizError); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": biz.Code,
				"msg":  biz.Msg,
			})
			return
		}
		fields := []zap.Field{zap.String("trace_id", c.GetString(`trace_id`))}
		zap.L().Error(err.Error(), fields...)
		c.JSON(http.StatusInternalServerError, gin.H{
			`code`: http.StatusInternalServerError,
			`msg`:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "query successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}
