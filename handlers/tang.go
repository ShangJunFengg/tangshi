package handlers

import (
	"MyTang/db"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	var name = c.Param("name")
	var sql = "SELECT * FROM poetries p " +
		" left join poets s on p.poet_id = s.id " +
		" WHERE p.title LIKE CONCAT ('%',?,'%' ) or p.content like CONCAT ('%',?,'%' ) or s.name like CONCAT ('%',?,'%' )"
	json, _ := db.Engine.Sql(sql, name, name, name).QueryWithDateFormat("2006-01-02 15:04:05").Json()
	res, _ := simplejson.NewJson([]byte(json))
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": res})
}
