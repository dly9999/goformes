package model

import (
	"database/sql"
	"fmt"
	"goformescloud/mesnewcloud/pkg"
	"strings"

	"goformescloud/mesnewcloud/logprint"

	"github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type EmployeeOutPut struct {
	TenantId   sql.NullString `json:"tenantId"`
	OrgCode    sql.NullString `json:"OrgCode"`
	OrgName    sql.NullString `json:"OrgName"`
	OrgType    sql.NullString `json:"OrgType"`
	OutputSum  sql.NullString `json:"OutputSum"`
	MESSort    sql.NullString `json:"MESSort"`
	SortType   sql.NullString `json:"SortType"`
	OwnDay     sql.NullString `json:"OwnDay"`
	OwnMonth   sql.NullString `json:"OwnMonth"`
	StepCode   sql.NullString `json:"StepCode"`
	CreateBy   sql.NullString `json:"CreateBy"`
	CreateDate mysql.NullTime `json:"CreateDate"`
	UpdateBy   sql.NullString `json:"UpdateBy"`
	UpdateDate mysql.NullTime `json:"UpdateDate"`
	DeptCode   sql.NullString `json:"DeptCode"`
	TeamCode   sql.NullString `json:"TeamCode"`
}
type EmployeeOutPutParameter struct {
	TenantId    string `json:tenantId`
	OwnDayFrom  string `json:OwnDayFrom`
	OwnDayTo    string `json:OwnDayTo`
	OwnMonth    string `json:OwnMonth`
	OwnDay      string `json:OwnDay`
	DeptCode    string `json:DeptCode`
	TeamCode    string `json:TeamCode`
	DateType    string `json:DateType`
	PageSize    int    `json:PageSize`
	CurrentPage int    `json:CurrentPage`
}

//查询多行
func QueryMultiEmployeeOutPut(DB *sqlx.DB, parameter EmployeeOutPutParameter) (employeeOutPutList []map[string]interface{}) {
	fmt.Println("参数实体__", parameter)
	var toatalpieces int
	var employee EmployeeOutPut
	var prams []interface{}
	var builderSQLTotal strings.Builder
	var builderSQL strings.Builder
	builderSQLTotal.WriteString(" select count(1) from  ( ")
	if parameter.DateType == "1" {
		builderSQL.WriteString(" select a.tenantId, a.OrgCode, a.OrgName, a.OutputSum, a.MESSort, a.SortType, a.OrgType, a.OwnDay, a.StepCode, a.CreateBy, a.CreateDate, a.UpdateBy, a.UpdateDate,b.DeptCode, d.TeamCode from Work_DayCountEmployee a ")
	}
	if parameter.DateType == "2" {
		builderSQL.WriteString(" select a.tenantId, a.OrgCode, a.OrgName, a.OutputSum, a.MESSort, a.SortType, a.OrgType, a.OwnMonth, a.StepCode, a.CreateBy, a.CreateDate, a.UpdateBy, a.UpdateDate,b.DeptCode, d.TeamCode from  Work_MonthCountEmployee a  ")
	}
	builderSQL.WriteString(" inner join  HR_Employee b on a.OrgCode=b.EmployeeCode ")
	builderSQL.WriteString(" inner join  HR_Depart c on b.DeptCode=c.DeptCode ")
	builderSQL.WriteString(" inner join  HR_Team d on b.TeamCode=d.TeamCode ")
	builderSQL.WriteString(" where  1=1 ")

	if parameter.TenantId != "" {
		builderSQL.WriteString(" and a.tenantId= ? ")
		prams = append(prams, parameter.TenantId)
	}

	if parameter.DateType == "1" {
		if parameter.OwnDayFrom != "" {
			builderSQL.WriteString(" and OwnDay >= ? ")
			fmt.Println("起止时间：", parameter.OwnDayFrom)
			prams = append(prams, parameter.OwnDayFrom)
		}
		if parameter.OwnDayTo != "" {
			builderSQL.WriteString(" and OwnDay <= ? ")
			prams = append(prams, parameter.OwnDayTo)
		}
	}
	if parameter.DateType == "2" {

		if parameter.OwnDayFrom != "" {
			builderSQL.WriteString(" and OwnMonth >= ? ")
			/*Year := parameter.OwnDayFrom.Unix(dynamic.UpdateTime/1000, 0).Year().String()
			Month := parameter.OwnDayFrom.Unix(dynamic.UpdateTime/1000, 0).Month().String()*/
			fmt.Println("起止时间：", parameter.OwnDayFrom)
			prams = append(prams, parameter.OwnDayFrom)
		}
		if parameter.OwnDayTo != "" {
			builderSQL.WriteString(" and OwnMonth <= ? ")
			/*Year := OwnDayTo.Unix(dynamic.UpdateTime/1000, 0).Year().String()
			Month := OwnDayTo.Unix(dynamic.UpdateTime/1000, 0).Month().String()*/
			prams = append(prams, parameter.OwnDayTo)
		}
	}

	if parameter.DeptCode != "" {

		builderSQL.WriteString(" and c.DeptCode >= ? ")
		prams = append(prams, parameter.DeptCode)
	}
	if parameter.TeamCode != "" {
		builderSQL.WriteString(" and d.TeamCode >= ? ")
		prams = append(prams, parameter.TeamCode)
	}
	builderSQLTotal.WriteString(builderSQL.String() + " ) as t")
	fmt.Println("分页sql>>>", builderSQLTotal.String())

	rowcount, er := DB.Query(builderSQLTotal.String(), prams...)
	if er != nil {
		logprint.Logprint("cutform :", er)
		logprint.Logprint("builderSQLTotal", builderSQLTotal.String())
		//	err = er
		return
	}
	for rowcount.Next() {
		er := rowcount.Scan(&toatalpieces)
		if er != nil {
			logprint.Logprint("rowcount :", er)
		}
	}
	fmt.Println("分页参数--", toatalpieces, parameter.PageSize, parameter.CurrentPage)

	pagemodel := pkg.Method_QueryFormTotal(toatalpieces, parameter.PageSize, parameter.CurrentPage)

	builderSQL.WriteString(" limit ? ") //strconv.Itoa(listmo.PageSize))
	prams = append(prams, pagemodel.PageSize)
	builderSQL.WriteString(" offset ? ") //+ strconv.Itoa((listmo.CurrentPage-1)*listmo.PageSize))
	prams = append(prams, (pagemodel.CurrentPage-1)*pagemodel.PageSize)

	fmt.Println("sql>>>>>>", builderSQL.String())
	fmt.Println(" CurrentPage---", pagemodel.CurrentPage)
	fmt.Println(" PageSize---", pagemodel.PageSize)

	fmt.Println(">>>>>>>>listParms>>>>", prams)
	rows, err := DB.Query(builderSQL.String(), prams...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	for rows.Next() {
		if parameter.DateType == "1" {
			err = rows.Scan(&employee.TenantId, &employee.OrgCode, &employee.OrgName, &employee.OutputSum, &employee.MESSort, &employee.SortType, &employee.OrgType, &employee.OwnDay, &employee.StepCode, &employee.CreateBy, &employee.CreateDate, &employee.UpdateBy, &employee.UpdateDate, &employee.DeptCode, &employee.TeamCode)
			if err != nil {
				fmt.Printf("Scan failed,err:%v", err)
				return
			}
			map_sg := pkg.StructToMapDemo(employee)
			employeeOutPutList = append(employeeOutPutList, map_sg)
		}
		if parameter.DateType == "2" {
			err = rows.Scan(&employee.TenantId, &employee.OrgCode, &employee.OrgName, &employee.OutputSum, &employee.MESSort, &employee.SortType, &employee.OrgType, &employee.OwnMonth, &employee.StepCode, &employee.CreateBy, &employee.CreateDate, &employee.UpdateBy, &employee.UpdateDate, &employee.DeptCode, &employee.TeamCode)
			if err != nil {
				fmt.Printf("Scan failed,err:%v", err)
				return
			}
			map_sg := pkg.StructToMapDemo(employee)
			employeeOutPutList = append(employeeOutPutList, map_sg)
		}

	}
	fmt.Println("---------结果实体——————", employeeOutPutList)
	return

}
