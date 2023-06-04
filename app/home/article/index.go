package article

import (
	"gofly/app/model"
	"gofly/utils"
	"gofly/utils/results"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
*使用 Index 是省略路径中的index
*本路径为： /admin/user/login -省去了index
 */
type Index struct{}

func init() {
	fpath := Index{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 文章
func (api *Index) Get_list(c *gin.Context) {
	cid := c.DefaultQuery("cid", "0")
	name := c.DefaultQuery("name", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("gofly_article").Fields("id,cid,title,image,des,weigh,status,createtime").
		Where("status", 0)
	if cid != "0" {
		MDB.Where("cid", cid)
	}
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id asc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		for _, val := range list {
			catename, _ := model.DB().Table("gofly_article_cate").Where("id", val["cid"]).Value("name")
			val["catename"] = catename
		}
		var totalCount int64
		totalCount, _ = model.DB().Table("gofly_article").Count()
		results.Success(c, "文章分类列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}
