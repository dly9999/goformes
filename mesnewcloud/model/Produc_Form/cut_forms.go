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
	"github.com/goinggo/mapstructure"
)

type Cut_Formout struct {
	CutDate      string `json:cutdate`
	OrderCode    string `json:ordercode`
	Customer     string `json:customer`
	Odid         int    `json:Odid`
	Sort         string `json:sort`
	Finished     string `json:iffinish`
	IfStop       string `json:ifstop`
	ByWayDate    string `json:bywaydate`
	StopRemark   string `json:stopremark`
	ReWorkTime   string `json:reworktime`
	StepCode     string `json:stepcode`
	LoadDate     string `json:LoadDate`
	EmployeeName string `json:employeename`
	TeamName     string `json:teamname`
	DeptName     string `json:DeptName`
}
type Cut_Form struct {
	CutDate      mysql.NullTime `json:cutdate`
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
type Cut_Condition struct {
	CutDate_Max string `json:cutdate_max`
	CutDate_Min string `json:cutdate_min`
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
type Cutforms struct {
	Cutforms           []Cut_Formout
	CurrentTotalPieces int
	CurrentPage        int
	TotalPage          int
}

func Dbcutforms(conditon Cut_Condition) (cuts Cutforms, err error) {
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

	buldsql.WriteString("select distinct a.Customer,a.CutDate,a.OrderCode,b.MesSort,b.Odid,b.ByWayDate,")
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
	if conditon.CutDate_Max != "" && conditon.CutDate_Min != "" {
		buldsql.WriteString(" and a.CutDate > ? and a.CutDate < ?   ")
		prams = append(prams, func() time.Time {
			tim, _ := time.Parse("2006-01-02 15:04:05", conditon.CutDate_Min)
			return tim
		}())
		prams = append(prams, func() time.Time {
			tim, _ := time.Parse("2006-01-02 15:04:05", conditon.CutDate_Max)
			return tim
		}())
	}
	if conditon.Sort != "" {
		buldsql.WriteString(" and b.Sort like CONCAT('%',?,'%') ")
		prams = append(prams, conditon.Sort)
	}
	buldCount.WriteString(buldsql.String() + ") as temp")
	rowcount, er := Db.ReDb().Sql.Query(buldCount.String(), prams...)
	logprint.Logprint("buldCount", buldCount.String())
	if er != nil {
		logprint.Logprint("cutform :", er)
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
		var cut Cut_Form

		var cout Cut_Formout
		er := row.Scan(&cut.Customer, &cut.CutDate, &cut.OrderCode, &cut.Sort,
			&cut.Odid, &cut.ByWayDate, &cut.ReWorkTime, &cut.IfStop, &cut.StopRemark, &cut.StepCode, &cut.Finished, &cut.LoadDate, &cut.DeptName, &cut.EmployeeName, &cut.TeamName)
		if er != nil {
			fmt.Println(er)
		}
		cutmap := pkg.StructToMapDemo(cut)
		if er := mapstructure.Decode(cutmap, &cout); er != nil {
			logprint.Logprint("Cutforms", er)
		}
		cuts.Cutforms = append(cuts.Cutforms, cout)
		//var maptest=
	}
	defer row.Close()
	cuts.TotalPage = pagemodel.PageCount
	cuts.CurrentTotalPieces = len(cuts.Cutforms)
	cuts.CurrentPage = pagemodel.CurrentPage
	return
}
