package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gin-api/utils"
	"strings"
)

type CondWhere struct {
	Where  map[string]interface{}
	Offset int
	Limit  int
	Order  string
}

func CondWhereAssemble(c *gin.Context) *CondWhere {
	_, per, offset := utils.GetPaginate(c)

	q := c.Request.URL.Query()

	order := "created_at DESC"

	where := make(map[string]interface{})

	for s, v := range q {
		if s == "page" || s == "per" {
			continue
		}
		if s == "sort" {
			sort := "desc"
			a := strings.Split(v[0], ",")
			if a[1] == "asc" {
				sort = "asc"
			}
			order = fmt.Sprintf("%s %s", a[0], sort)
			continue
		}
		where[s] = v[0]
	}

	return &CondWhere{
		Where:  where,
		Offset: offset,
		Limit:  per,
		Order:  order,
	}
}

func FilterSql(cw *CondWhere, db *gorm.DB) *gorm.DB {
	for k, v := range cw.Where {
		if bool := strings.Contains(k, "_matches"); bool {
			db = db.Where(strings.Split(k, "_matches")[0]+" like ?", fmt.Sprintf("%%%s%%", v))
		} else if bool := strings.Contains(k, "_lt"); bool {
			db = db.Where(strings.Split(k, "_lt")[0]+" < ?", v)
		} else if bool := strings.Contains(k, "_gt"); bool {
			db = db.Where(strings.Split(k, "_gt")[0]+" > ?", v)
		} else {
			db = db.Where(k+" = ?", v)
		}
	}

	return db
}
