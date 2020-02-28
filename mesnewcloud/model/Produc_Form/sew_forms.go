package Produc_Form

import (
	"database/sql"
	"fmt"
	"goformescloud/mesnewcloud/db"
	"goformescloud/mesnewcloud/logprint"
	"goformescloud/mesnewcloud/pkg"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	//	"github.com/goinggo/mapstructure"
)

type Sew_Form struct {
	SewDate      mysql.NullTime `json:SewDate`
	OrderCode    sql.NullString `json:ordercode`
	Customer     sql.NullString `json:customer`
	Odid         sql.NullInt64  `json:Odid`
	Sort         sql.NullString `json:sort`
	Finished     sql.NullString `json:finished`
	IfStop       sql.NullString `json:ifstop`
	ByWayDate    mysql.NullTime `json:bywaydate`
	StopRemark   sql.NullString `json:stopremark`
	ReWorkTime   sql.NullString `json:reworktime`
	StepCode     sql.NullString `json:stepcode`
	LoadDate     mysql.NullTime `json:LoadDate`
	EmployeeName sql.NullString `json:employeename`
	TeamName     sql.NullString `json:teamname`
	DeptName     sql.NullString `json:DeptName`
}
type Sew_Condition struct {
	SewDate_Max string `json:Sewdate_max`
	SewDate_Min string `json:Sewdate_min`
	OrderCode   string `json:ordercode`
	Sort        string `json:sort`
	IfFinish    string `json:IfFinish`
	IfStop      string `json:IfStop`
	Ifbulk      string `json:Ifbulk`
	CurrentPage int    `json:CurrentPage`
	TenantId    string `json:tenantId`
	PageSize    int    `json:pagesize`
	//	ToatlPage   int    `json:ToatlPage`

}
type Forms struct {
	Sewforms           []map[string]interface{}
	CurrentTotalPieces int
	CurrentPage        int
	TotalPage          int
}

func Dbsewforms(conditon Sew_Condition) (sews Forms, err error) {
	var prams []interface{}
	var buldsql strings.Builder
	var buldCount strings.Builder
	var toatalpieces int
	buldCount.WriteString("select count(1) ")
	buldCount.WriteString("from (")
	/*	buldCount.WriteString(" inner join  OrderDetail b on a.SysCode=b.SysCode")
		buldCount.WriteString(" inner join  BrushCard_Material c on c.odid=b.odid and c.stepcode=b.stepcode")
		buldCount.WriteString(" inner join  HR_Employee d on c.EmpCode=d.EmployeeCode")
		buldCount.WriteString(" inner join  HR_Team e on e.TeamCode=d.TeamCode")
		buldCount.WriteString(" inner join  HR_Depart f on f.DeptCode=d.DeptCode where ")*/

	buldsql.WriteString("select distinct a.Customer,a.SewingDate,a.OrderCode,b.MesSort,b.Odid,b.ByWayDate,")
	buldsql.WriteString("b.ReWorkTime, ")
	buldsql.WriteString("case  b.IfStop when 1 then '暂停生产' else '正在生产' end as IfStop,")
	buldsql.WriteString("b.StopRemark, b.StepCode,")
	buldsql.WriteString("case b.ByWay when 1 then '已完成' else '未完成' end as Finished,")
	buldsql.WriteString("c.LoadDate,f.DeptName,d.EmployeeName,e.TeamName ")
	buldsql.WriteString("from ProductOrder as a ")
	buldsql.WriteString(" inner join  OrderDetail b on a.SysCode=b.SysCode")
	buldsql.WriteString(" inner join  BrushCard_Material c on c.odid=b.odid and c.stepcode=b.stepcode")
	buldsql.WriteString(" inner join  HR_Employee d on c.EmpCode=d.EmployeeCode")
	buldsql.WriteString(" inner join  HR_Team e on e.TeamCode=d.TeamCode")
	buldsql.WriteString(" inner join  HR_Depart f on f.DeptCode=d.DeptCode where ")
	if conditon.TenantId != "" {
		buldsql.WriteString("b.tenantId = ? ")
		prams = append(prams, conditon.TenantId)
	}
	if conditon.SewDate_Max != "" && conditon.SewDate_Min != "" {
		buldsql.WriteString(" and a.CutDate > ? and a.CutDate < ?   ")
		prams = append(prams, func() time.Time {
			tim, _ := time.Parse("2006-01-02 15:04:05", conditon.SewDate_Min)
			return tim
		}())
		prams = append(prams, func() time.Time {
			tim, _ := time.Parse("2006-01-02 15:04:05", conditon.SewDate_Max)
			return tim
		}())
	}
	if conditon.Sort != "" {
		buldsql.WriteString(" and b.Sort like CONCAT('%',?,'%') ")
		prams = append(prams, conditon.Sort)
	}
	buldCount.WriteString(buldsql.String() + ") as temp")
	rowcount, er := Db.ReDb().Sql.Query(buldCount.String(), prams...)
	if er != nil {
		logprint.Logprint("sewform :", er)
		logprint.Logprint("buldCount", buldCount.String())
		return
	}
	for rowcount.Next() {
		er := rowcount.Scan(&toatalpieces)
		if er != nil {
			logprint.Logprint("rowcount :", er)
		}
	}
	pagemodel := pkg.Method_QueryFormTotal(toatalpieces, conditon.PageSize, conditon.CurrentPage)
	buldsql.WriteString(" limit ? ") //strconv.Itoa(listmo.PageSize))
	prams = append(prams, pagemodel.PageSize)
	buldsql.WriteString(" offset ? ") //+ strconv.Itoa((listmo.CurrentPage-1)*listmo.PageSize))
	prams = append(prams, (pagemodel.CurrentPage-1)*pagemodel.PageSize)
	row, er := Db.ReDb().Sql.Query(buldsql.String(), prams...)
	if er != nil {
		logprint.Logprint("cutform :", er)
		logprint.Logprint("buldsql", buldsql.String())
		//	err = er
		return
	}
	for row.Next() {
		var sew Sew_Form

		//var cout Sewforms
		er := row.Scan(&sew.Customer, &sew.SewDate, &sew.OrderCode, &sew.Sort,
			&sew.Odid, &sew.ByWayDate, &sew.ReWorkTime, &sew.IfStop, &sew.StopRemark, &sew.StepCode, &sew.Finished, &sew.LoadDate, &sew.DeptName, &sew.EmployeeName, &sew.TeamName)
		if er != nil {
			fmt.Println(er)
		}
		sewmap := pkg.StructToMapDemo(sew)
		sews.Sewforms = append(sews.Sewforms, sewmap)
		/*if er := mapstructure.Decode(cutmap, &cout); er != nil {
			logprint.Logprint("Cutforms", er)
		}
		cuts.Cutforms = append(cuts.Cutforms, cout)*/

		//var maptest=
	}
	defer row.Close()
	sews.TotalPage = pagemodel.PageCount
	sews.CurrentTotalPieces = len(sews.Sewforms)
	sews.CurrentPage = pagemodel.CurrentPage
	return
}
