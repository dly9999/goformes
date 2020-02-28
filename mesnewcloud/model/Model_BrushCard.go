package model

import (

	//	"container/list"
	//	"encoding/json"
	"fmt"
	//"math"
	"goformescloud/mesnewcloud/pkg"
	"strconv"
	"strings"

	//	"goformescloud/mesnewcloud/config"
	"goformescloud/mesnewcloud/db"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

type Model_BrushCard struct {
	CutDate      string
	OrderCode    string
	MesSort      string
	StepCode     string
	TeamCode     string
	TeamName     string
	EmpCode      string
	EmployeeName string
	LoadTime     string
	Odid         int
	OdState      int
}

type ModelList_BrushCard struct {
	CurrentPage int
	PageSize    int
	PageCount   int
	TotalRows   int
	ListParent  []Model_BrushCard
}

func Method_QueryBrushCardList(tenantId string, orderCode string, OrderSort string, pageSize int, currentPage int) *ModelList_BrushCard {

	listmo := &ModelList_BrushCard{}
	var builderSQLTotal strings.Builder
	builderSQLTotal.WriteString(" select count(1)")
	builderSQLTotal.WriteString(" from ProductOrder t1 left join OrderDetail t2 on t1.`SysCode`=t2.`SysCode` and t1.`tenantId`=t2.`tenantId` ")
	builderSQLTotal.WriteString(" left join BrushCard_Material t3 on t2.`Odid`=t3.`OdID`")
	builderSQLTotal.WriteString(" left join HR_Employee t4 on t4.`EmployeeCode`=t3.`EmpCode`")
	builderSQLTotal.WriteString(" left join HR_Team t5 on t5.`TeamCode`=t4.`TeamCode`")
	builderSQLTotal.WriteString(" where t1.`tenantId`='" + tenantId + "'")
	if orderCode != "" {
		builderSQLTotal.WriteString(" and t1.`OrderCode`='" + orderCode + "'")
	}
	if OrderSort != "" {
		builderSQLTotal.WriteString(" and t2.`MesSort`='" + OrderSort + "'")
	}
	fmt.Println(builderSQLTotal.String())
	paging := pkg.Method_QueryDataTotal(builderSQLTotal.String(), pageSize, currentPage)
	if paging.TotalRows <= 0 {
		return listmo
	} else {
		listmo.CurrentPage = paging.CurrentPage
		listmo.PageCount = paging.PageCount
		listmo.PageSize = pageSize
		listmo.TotalRows = paging.TotalRows
	}

	var builderSQL strings.Builder
	builderSQL.WriteString(" select 	t1.OrderCode,	t2.MesSort,	t2.odid,")
	builderSQL.WriteString(" t3.StepCode,	t3.OdState,	t4.TeamCode, ")
	builderSQL.WriteString(" t5.TeamName,	t3.EmpCode,	t4.EmployeeName,")
	builderSQL.WriteString(" t3.LoadDate,t1.CutDate ")
	builderSQL.WriteString(" from ProductOrder t1 inner join OrderDetail t2 on t1.`SysCode`=t2.`SysCode` and t1.`tenantId`=t2.`tenantId` ")
	builderSQL.WriteString(" inner join BrushCard_Material t3 on t2.`Odid`=t3.`OdID`")
	builderSQL.WriteString(" inner join HR_Employee t4 on t4.`EmployeeCode`=t3.`EmpCode`")
	builderSQL.WriteString(" inner join HR_Team t5 on t5.`TeamCode`=t4.`TeamCode`")
	builderSQL.WriteString(" where t1.`tenantId`='" + tenantId + "'")
	if orderCode != "" {
		builderSQL.WriteString(" and t1.`OrderCode`='" + orderCode + "'")
	}
	if OrderSort != "" {
		builderSQL.WriteString(" and t2.`MesSort`='" + OrderSort + "'")
	}
	builderSQL.WriteString(" limit " + strconv.Itoa(listmo.PageSize))
	builderSQL.WriteString(" offset  " + strconv.Itoa((listmo.CurrentPage-1)*listmo.PageSize))
	fmt.Println(builderSQL.String())
	rowsArray, err := Db.ReDb().Sql.Query(builderSQL.String())
	if err != nil {
		fmt.Println("SQL" + err.Error())
	}
	for rowsArray.Next() {
		ins_BrushCard := Model_BrushCard{}
		err = rowsArray.Scan(&ins_BrushCard.OrderCode, &ins_BrushCard.MesSort, &ins_BrushCard.Odid,
			&ins_BrushCard.StepCode, &ins_BrushCard.OdState, &ins_BrushCard.TeamCode,
			&ins_BrushCard.TeamName, &ins_BrushCard.EmpCode, &ins_BrushCard.EmployeeName,
			&ins_BrushCard.LoadTime, &ins_BrushCard.CutDate)
		listmo.ListParent = append(listmo.ListParent, ins_BrushCard)
		if err != nil {
			fmt.Println("Get Data from Row failed" + err.Error())
			//loger.Output(2, "Get Data from Row failed"+err.Error()+"\r\n")
			break
		}
	}
	return listmo
}
