package model

import (

	//	"container/list"
	//	"encoding/json"
	"fmt"
	//"math"
	"goformescloud/mesnewcloud/pkg"
	"strconv"
	"strings"

	//	"time"

	//	"goformescloud/mesnewcloud/config"
	"goformescloud/mesnewcloud/db"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

type Model_OrderState struct {
	Odid     int
	MesSort  string
	Counts   string
	FabricNo string
	StyleNo  string
	Finished string

	CardState string
	StepCode  string
	ByWay     string

	ByWayDate  string
	IfStop     string
	StopRemark string
	SysCode    string
	OrderCode  string
	Customer   string

	PlanDate     string
	CutDate      string
	SewingDate   string
	LroningDate  string
	PackingDate  string
	DeliveryDate string

	EmpCode      string
	EmployeeName string
	DeptCode     string
	TeamCode     string
	TeamName     string
	DeptName     string
}

type ModelList_OrderState struct {
	CurrentPage int
	PageSize    int
	PageCount   int
	TotalRows   int
	ListParent  []Model_OrderState
}

func Method_QueryOrderStateList(tenantId string, orderCode string, OrderSort string, pageSize int, currentPage int) *ModelList_OrderState {

	listmo := &ModelList_OrderState{}
	var builderSQLTotal strings.Builder
	builderSQLTotal.WriteString(" select count(1)")
	builderSQLTotal.WriteString(" from  ProductOrder t2 inner join  OrderDetail t1 on t2.SysCode =t1.SysCode and t2.tenantId=t1.tenantId  ")
	builderSQLTotal.WriteString(" inner join  BrushCard_Material  t3 on t1.Odid =t3.Odid  and t1.StepCode=t3.StepCode")
	builderSQLTotal.WriteString(" inner join  HR_Employee t4 on   t3.EmpCode=t4.EmployeeCode")
	builderSQLTotal.WriteString(" inner join  HR_Team t5 on   t4.TeamCode=t5.TeamCode")
	builderSQLTotal.WriteString(" inner join  HR_Depart t6 on   t6.DeptCode=t4.DeptCode")
	builderSQLTotal.WriteString(" where t1.`tenantId`='" + tenantId + "'")
	if orderCode != "" {
		builderSQLTotal.WriteString(" and t2.`OrderCode`='" + orderCode + "'")
	}
	if OrderSort != "" {
		builderSQLTotal.WriteString(" and t1.`MesSort`='" + OrderSort + "'")
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
	builderSQL.WriteString(" select  t1.Odid, t1.MesSort, ifnull(t2.Counts,'') as Counts, ifnull(t1.FabricNo,'') as FabricNo,")
	builderSQL.WriteString(" ifnull(t1.StyleNo,'') as StyleNo, ifnull(t1.Finished,'') as Finished, ifnull(t1.CardState,'') as CardState, ")
	builderSQL.WriteString(" t1.StepCode,case(t1.ByWay) when 1 then '已过通道'  else  '未过通道' end ByWay, ")
	builderSQL.WriteString(" ifnull(t1.ByWayDate,'') as ByWayDate, ifnull(t1.IfStop,'') as IfStop, ifnull(t1.StopRemark,'') as StopRemark,   ")
	builderSQL.WriteString(" t2.SysCode, t2.OrderCode, t2.Customer,t2.PlanDate, t2.CutDate, t2.SewingDate, t2.LroningDate, t2.PackingDate,")
	builderSQL.WriteString(" t2.DeliveryDate , t3.EmpCode,t4.EmployeeName ,t4.DeptCode,t4.TeamCode,t5.TeamName,t6.DeptName ")
	builderSQL.WriteString(" from  ProductOrder t2 inner join  OrderDetail t1 on t2.SysCode =t1.SysCode and t2.tenantId=t1.tenantId  ")
	builderSQL.WriteString(" inner join  BrushCard_Material  t3 on t1.Odid =t3.Odid  and t1.StepCode=t3.StepCode")
	builderSQL.WriteString(" inner join  HR_Employee t4 on   t3.EmpCode=t4.EmployeeCode")
	builderSQL.WriteString(" inner join  HR_Team t5 on   t4.TeamCode=t5.TeamCode")
	builderSQL.WriteString(" inner join  HR_Depart t6 on   t6.DeptCode=t4.DeptCode")
	builderSQL.WriteString(" where t1.`tenantId`='" + tenantId + "'")
	if orderCode != "" {
		builderSQL.WriteString(" and t2.`OrderCode`='" + orderCode + "'")
	}
	if OrderSort != "" {
		builderSQL.WriteString(" and t1.`MesSort`='" + OrderSort + "'")
	}
	builderSQL.WriteString(" limit " + strconv.Itoa(listmo.PageSize))
	builderSQL.WriteString(" offset  " + strconv.Itoa((listmo.CurrentPage-1)*listmo.PageSize))
	fmt.Println(builderSQL.String())
	rowsArray, err := Db.ReDb().Sql.Query(builderSQL.String())
	if err != nil {
		fmt.Println("SQL" + err.Error())
	}
	for rowsArray.Next() {
		ins_OrderState := Model_OrderState{}
		err = rowsArray.Scan(&ins_OrderState.Odid, &ins_OrderState.MesSort, &ins_OrderState.Counts,
			&ins_OrderState.FabricNo, &ins_OrderState.StyleNo, &ins_OrderState.Finished,
			&ins_OrderState.CardState, &ins_OrderState.StepCode, &ins_OrderState.ByWay,
			&ins_OrderState.ByWayDate, &ins_OrderState.IfStop, &ins_OrderState.StopRemark,
			&ins_OrderState.SysCode, &ins_OrderState.OrderCode, &ins_OrderState.Customer,
			&ins_OrderState.PlanDate, &ins_OrderState.CutDate, &ins_OrderState.SewingDate,
			&ins_OrderState.LroningDate, &ins_OrderState.PackingDate, &ins_OrderState.DeliveryDate,
			&ins_OrderState.EmpCode, &ins_OrderState.EmployeeName, &ins_OrderState.DeptCode,
			&ins_OrderState.TeamCode, &ins_OrderState.TeamName, &ins_OrderState.DeptName)
		listmo.ListParent = append(listmo.ListParent, ins_OrderState)
		if err != nil {
			fmt.Println("Get Data from Row failed" + err.Error())
			//loger.Output(2, "Get Data from Row failed"+err.Error()+"\r\n")
			break
		}
	}
	return listmo
}
