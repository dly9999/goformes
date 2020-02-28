package pkg

import (
	"fmt"
	"goformescloud/mesnewcloud/db"
)

type PagingModel struct {
	CurrentPage int //当前页
	PageSize    int //每页条数
	TotalRows   int //总条数
	PageCount   int //总页数
}

func Method_QueryDataTotal(strSQL string, pageSize int, currentPage int) PagingModel {
	var paging PagingModel
	err := Db.ReDb().Sql.QueryRow(strSQL).Scan(&paging.TotalRows)
	if err != nil {
		fmt.Println("SQL" + err.Error())
		paging.TotalRows = 0
		return paging
	} else {
		if paging.TotalRows <= 0 {
			paging.CurrentPage = 0
			paging.PageCount = 0
			paging.PageSize = pageSize
			return paging
		} else if pageSize*currentPage > paging.TotalRows {
			paging.PageCount = (paging.TotalRows + pageSize - 1) / pageSize
			paging.CurrentPage = paging.PageCount
			paging.PageSize = pageSize
			return paging
		} else {
			paging.PageCount = (paging.TotalRows + pageSize - 1) / pageSize
			paging.CurrentPage = currentPage
			paging.PageSize = pageSize
			return paging
		}
	}
}
