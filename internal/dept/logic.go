package dept

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nysynysy2/go_web_demo/api"
	"github.com/nysynysy2/go_web_demo/internal/db"
)

type Dept struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
}

var deptDb = db.DB

func GetDeptBody(ctx *gin.Context) (*Dept, error) {
	nd := Dept{}
	if err := ctx.ShouldBindJSON(&nd); err != nil {
		ctx.JSON(http.StatusBadRequest, api.RError(err))
		return nil, err
	}
	return &nd, nil
}

func GetIdParam(ctx *gin.Context) (int, error) {
	idStr, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.RError(err))
		return 0, err
	}
	return id, nil
}

func GetAll(ctx *gin.Context) {
	depts := []Dept{}
	fmt.Printf("%#v", depts)
	res := deptDb.Table("dept").Find(&depts)
	if res.Error != nil {
		ctx.JSON(500, api.RError(res.Error))
		return
	}
	ctx.JSON(200, api.RSuccess(depts))
}

func Update(ctx *gin.Context) {
	nd, err := GetDeptBody(ctx)
	if err != nil {
		return
	}
	if err := deptDb.Table("dept").Updates(nd).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, api.RError(err))
	}
	ctx.JSON(http.StatusOK, api.RSuccess(nil))
}
func Create(ctx *gin.Context) {
	nd, err := GetDeptBody(ctx)
	if err != nil {
		return
	}
	err = deptDb.Table("dept").Create(nd).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.RError(err))
	}
	ctx.JSON(http.StatusOK, api.RSuccess(nil))
}

func GetById(ctx *gin.Context) {
	id, err := GetIdParam(ctx)
	if err != nil {
		return
	}
	d := Dept{}
	d.ID = id

	if err := deptDb.Table("dept").First(&d).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, api.RError(err))
	}
	ctx.JSON(http.StatusOK, api.RSuccess(d))
}

func DeleteById(ctx *gin.Context) {
	id, err := GetIdParam(ctx)
	if err != nil {
		return
	}
	d := &Dept{}
	d.ID = id
	err = deptDb.Table("dept").Delete(&d).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.RError(err))
	}
	ctx.JSON(http.StatusOK, api.RSuccess(nil))
}
