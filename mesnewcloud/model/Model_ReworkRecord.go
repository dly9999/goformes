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

type Model_ReworkRecord struct {
	Odid         int
	MesSort      string
	OrderCode    string
	PlanCode     string
	QCCheckID    int
	OutQualityID int
	QCCheckDate  string
	StepCode     string
	DutyGX       string
	Memo1        string
	BrushDate    string
	EmpCode      string
	EmployeeName string
	DeptCode     string
	TeamCode     string
	TeamName     string
	DeptName     string
}

type ModelList_ReworkRecord struct {
	CurrentPage int
	PageSize    int
	PageCount   int
	TotalRows   int
	ListParent  []Model_ReworkRecord
}

func Method_QueryReworkRecordList(tenantId string, orderCode string, OrderSort string, PlanCode string, pageSize int, currentPage int) *ModelList_ReworkRecord {
	listmo := &ModelList_ReworkRecord{}
	var builderSQLTotal strings.Builder
	builderSQLTotal.WriteString(" select count(1)")
	builderSQLTotal.WriteString(" from  ProductOrder t1 inner join OrderDetail t2 on t1.SysCode=t2.SysCode  and t1.tenantId=t2.tenantId   ")
	builderSQLTotal.WriteString(" inner join QC_CheckMain  t4 on t2.Odid=t4.Odid")
	builderSQLTotal.WriteString(" inner join QC_OutCheckReason t5 on t4.IsOutCheck=false and t5.QCCheckID=t4.QCCheckID ")
	builderSQLTotal.WriteString(" inner join QC_OutQuality t6 on t6.Memo6=false and t6.OutQualityID=t5.OutQualityID")
	builderSQLTotal.WriteString(" inner join BrushCard_Material t7 on t7.StepCode=t4.`StepCode` and t7.Odid=t2.Odid")
	builderSQLTotal.WriteString(" inner join HR_Employee t8 on t8.`EmployeeCode` =t7.`EmpCode`")
	builderSQLTotal.WriteString(" inner join HR_Team t9 on t8.`TeamCode` = t9.`TeamCode`")
	builderSQLTotal.WriteString(" inner join HR_Depart t10 on t8.`DeptCode` = t10.`DeptCode`")
	builderSQLTotal.WriteString(" where t1.`tenantId`='" + tenantId + "'")
	if orderCode != "" {
		builderSQLTotal.WriteString(" and t1.`OrderCode`='" + orderCode + "'")
	}
	if OrderSort != "" {
		builderSQLTotal.WriteString(" and t2.`MesSort`='" + OrderSort + "'")
	}
	if PlanCode != "" {
		builderSQLTotal.WriteString(" and t1.`PlanCode`='" + PlanCode + "'")
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
	builderSQL.WriteString(" select   t1.PlanCode, t1.OrderCode, t2.Odid,t2.MesSort ,")
	builderSQL.WriteString(" t4.QCCheckDate,t4.QCCheckID,t4.StepCode,t6.OutQualityID,")
	builderSQL.WriteString(" t6.DutyGX,t6.Memo1,t7.BrushDate,t7.EmpCode,")
	builderSQL.WriteString(" t8.EmployeeName,t8.DeptCode,t8.TeamCode,t9.TeamName,")
	builderSQL.WriteString(" t10.DeptName")
	builderSQL.WriteString(" from  ProductOrder t1 inner join OrderDetail t2 on t1.SysCode=t2.SysCode  and t1.tenantId=t2.tenantId   ")
	builderSQL.WriteString(" inner join QC_CheckMain  t4 on t2.Odid=t4.Odid")
	builderSQL.WriteString(" inner join QC_OutCheckReason t5 on t4.IsOutCheck=false and t5.QCCheckID=t4.QCCheckID ")
	builderSQL.WriteString(" inner join QC_OutQuality t6 on t6.Memo6=false and t6.OutQualityID=t5.OutQualityID")
	builderSQL.WriteString(" inner join BrushCard_Material t7 on t7.StepCode=t4.`StepCode` and t7.Odid=t2.Odid")
	builderSQL.WriteString(" inner join HR_Employee t8 on t8.`EmployeeCode` =t7.`EmpCode`")
	builderSQL.WriteString(" inner join HR_Team t9 on t8.`TeamCode` = t9.`TeamCode`")
	builderSQL.WriteString(" inner join HR_Depart t10 on t8.`DeptCode` = t10.`DeptCode`")
	builderSQLTotal.WriteString(" where t1.`tenantId`='" + tenantId + "'")
	if orderCode != "" {
		builderSQLTotal.WriteString(" and t1.`OrderCode`='" + orderCode + "'")
	}
	if OrderSort != "" {
		builderSQLTotal.WriteString(" and t2.`MesSort`='" + OrderSort + "'")
	}
	if PlanCode != "" {
		builderSQLTotal.WriteString(" and t1.`PlanCode`='" + PlanCode + "'")
	}
	builderSQL.WriteString(" limit " + strconv.Itoa(listmo.PageSize))
	builderSQL.WriteString(" offset  " + strconv.Itoa((listmo.CurrentPage-1)*listmo.PageSize))
	fmt.Println(builderSQL.String())
	rowsArray, err := Db.ReDb().Sql.Query(builderSQL.String())
	if err != nil {
		fmt.Println("SQL" + err.Error())
	}
	for rowsArray.Next() {
		ins_ReworkRecord := Model_ReworkRecord{}
		err = rowsArray.Scan(&ins_ReworkRecord.PlanCode, &ins_ReworkRecord.OrderCode, &ins_ReworkRecord.Odid,
			&ins_ReworkRecord.MesSort, &ins_ReworkRecord.QCCheckDate, &ins_ReworkRecord.QCCheckID,
			&ins_ReworkRecord.StepCode, &ins_ReworkRecord.OutQualityID, &ins_ReworkRecord.DutyGX,
			&ins_ReworkRecord.Memo1, &ins_ReworkRecord.BrushDate, &ins_ReworkRecord.EmpCode,
			&ins_ReworkRecord.EmployeeName, &ins_ReworkRecord.DeptCode, &ins_ReworkRecord.TeamCode,
			&ins_ReworkRecord.TeamName, &ins_ReworkRecord.DeptName)
		listmo.ListParent = append(listmo.ListParent, ins_ReworkRecord)
		if err != nil {
			fmt.Println("Get Data from Row failed" + err.Error())
			//loger.Output(2, "Get Data from Row failed"+err.Error()+"\r\n")
			break
		}
	}
	return listmo
}
