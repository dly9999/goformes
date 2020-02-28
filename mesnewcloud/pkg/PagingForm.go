package pkg

import (
	"fmt"
)

func Method_QueryFormTotal(TotalRows, pageSize, currentPage int) PagingModel {
	var paging PagingModel
	paging.TotalRows = TotalRows
	if paging.TotalRows <= 0 {
		paging.CurrentPage = 0
		paging.PageCount = 0
		paging.PageSize = pageSize
		return paging
	} else if pageSize*currentPage > paging.TotalRows {
		fmt.Println("大于")
		paging.PageCount = (paging.TotalRows + pageSize - 1) / pageSize
		paging.CurrentPage = paging.PageCount
		paging.PageSize = pageSize
		return paging
	} else {
		fmt.Println("小于等于")
		paging.PageCount = (paging.TotalRows + pageSize - 1) / pageSize
		paging.CurrentPage = currentPage
		paging.PageSize = pageSize
		return paging
	}
}
