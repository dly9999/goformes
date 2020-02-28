package model

import (
	"database/sql"
	"fmt"
	"goformescloud/mesnewcloud/pkg"
	"strings"

	"github.com/go-sql-driver/mysql"

	"goformescloud/mesnewcloud/logprint"

	"github.com/jmoiron/sqlx"
)

//质量缺陷实体
type OrderDefect struct {
	TenantId     sql.NullString `json:"tenantId" db:"tenantId"`
	Odid         int            `json:"Odid" db:"Odid"`
	QCCheckDate  sql.NullString `json:"QCCheckDate" db:"QCCheckDate"`
	EmployCode   sql.NullString `json:"EmployCode" db:"EmployCode"`
	EmployeeName sql.NullString `json:"EmployeeName" db:"EmployeeName"`
	QCCheckID    sql.NullString `json:"QCCheckID" db:"QCCheckID"`
	StepCode     sql.NullString `json:"StepCode" db:"StepCode"`
	OutQualityID sql.NullString `json:"OutQualityID" db:"OutQualityID"`
	DutyGX       sql.NullString `json:"DutyGX" db:"DutyGX"`
	Memo1        sql.NullString `json:"Memo1" db:"Memo1"`
	OutTypeName  sql.NullString `json:"OutTypeName" db:"OutTypeName"`
	BrushDate    mysql.NullTime `json:"BrushDate" db:"BrushDate"`
	EmpCode      sql.NullString `json:"EmpCode" db:"EmpCode"`
	DeptCode     sql.NullString `json:"DeptCode" db:"DeptCode"`
	TeamCode     sql.NullString `json:"TeamCode" db:"TeamCode"`
	DeptName     sql.NullString `json:"DeptName" db:"DeptName"`
	TeamName     sql.NullString `json:"TeamName" db:"TeamName"`
}

//查询多行
func QueryMultiOrderDefectOne(DB *sqlx.DB, tenantId string, groupName string, sort string, outQualityName string, pageSize int, currentPage int) (orderdefList []map[string]interface{}) {
	fmt.Println("传入参数：", pageSize, currentPage)
	var toatalpieces int
	var orderdef OrderDefect
	var prams []interface{}
	var builderSQLTotal strings.Builder
	var builderSQL strings.Builder

	builderSQLTotal.WriteString(" select  count(1) from (")
	builderSQL.WriteString(" select distinct p.tenantId,p.Odid,b.QCCheckDate,b.EmployCode, b.EmployeeName,b.QCCheckID,b.StepCode, d.OutQualityID,d.DutyGX,d.Memo1,f.OutTypeName,e.BrushDate,e.EmpCode,g.DeptCode,h.TeamCode,m.DeptName,h.TeamName ")
	builderSQL.WriteString(" from ProductOrder a inner join OrderDetail p  on a.SysCode=p.SysCode ")
	builderSQL.WriteString(" inner join QC_CheckMain  b on p.Odid=b.Odid ")
	builderSQL.WriteString(" inner join QC_OutCheckReason c on b.IsOutCheck=false and b.QCCheckID=c.QCCheckID ")
	builderSQL.WriteString(" inner join QC_OutQuality d on d.Memo6=false and d.OutQualityID=c.OutQualityID ")
	builderSQL.WriteString(" inner join BrushCard_Material e on e.StepCode=DutyGX and e.Odid=p.Odid ")
	builderSQL.WriteString(" left  join QC_OutType f on d.Memo1=f.OutTypeID ")
	builderSQL.WriteString(" inner join HR_Employee g on g.EmployeeCode=e.EmpCode ")
	builderSQL.WriteString(" inner join HR_Team h on g.TeamCode=h.TeamCode ")
	builderSQL.WriteString(" inner join HR_Depart m on m.DeptCode=h.DeptCode   where 1=1 ")
	if tenantId != "" {
		builderSQL.WriteString(" and p.tenantId= ? ")
		prams = append(prams, tenantId)
	}
	if groupName != "" {
		builderSQL.WriteString(" and  GroupName like CONCAT('%',?,'%')")
		prams = append(prams, groupName)
	}
	if sort != "" {
		builderSQL.WriteString(" and  Sort like CONCAT('%',?,'%')")
		prams = append(prams, sort)
	}
	if outQualityName != "" {
		builderSQL.WriteString(" and  OutQualityName like CONCAT('%',?,'%')")
		prams = append(prams, outQualityName)
	}
	builderSQLTotal.WriteString(builderSQL.String() + ") as pro ")
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
	fmt.Println("分页参数--", toatalpieces, pageSize, currentPage)
	pagemodel := pkg.Method_QueryFormTotal(toatalpieces, pageSize, currentPage)

	builderSQL.WriteString(" limit ? ") //strconv.Itoa(listmo.PageSize))
	prams = append(prams, pagemodel.PageSize)
	builderSQL.WriteString(" offset ? ") //+ strconv.Itoa((listmo.CurrentPage-1)*listmo.PageSize))
	prams = append(prams, (pagemodel.CurrentPage-1)*pagemodel.PageSize)

	fmt.Println(" CurrentPage---", pagemodel.CurrentPage)
	fmt.Println(" PageSize---", pagemodel.PageSize)
	fmt.Println("sql>>>>>>", builderSQL.String())
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
		err = rows.Scan(&orderdef.TenantId, &orderdef.Odid, &orderdef.QCCheckDate, &orderdef.EmployCode, &orderdef.EmployeeName, &orderdef.QCCheckID, &orderdef.StepCode, &orderdef.OutQualityID, &orderdef.DutyGX, &orderdef.Memo1, &orderdef.OutTypeName, &orderdef.BrushDate, &orderdef.EmpCode, &orderdef.DeptCode, &orderdef.TeamCode, &orderdef.DeptName, &orderdef.TeamName)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		map_sg := pkg.StructToMapDemo(orderdef)
		orderdefList = append(orderdefList, map_sg)

	}
	fmt.Println("---------结果实体——————", orderdefList)
	return
}
