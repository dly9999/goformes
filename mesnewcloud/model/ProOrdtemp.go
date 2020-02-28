package model

import (
	"database/sql"
	"fmt"
	"goformescloud/mesnewcloud/pkg"
	"strings"

	"goformescloud/mesnewcloud/logprint"

	"github.com/go-sql-driver/mysql"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProOrder struct {
	Odid         int            `json:"Odid" db:"Odid"`
	SysCode      sql.NullString `json:"SysCode" db:"SysCode"`
	OrderCode    sql.NullString `json:"OrderCode" db:"OrderCode"`
	Customer     sql.NullString `json:"Customer" db:"Customer"`
	MesSort      sql.NullString `json:"MesSort" db:"MesSort"`
	Counts       sql.NullString `json:"Counts" db:"Counts"`
	FabricNo     sql.NullString `json:"FabricNo" db:"FabricNo"`
	StyleNo      sql.NullString `json:"StyleNo" db:"StyleNo"`
	Finished     sql.NullString `json:"Finished" db:"Finished"`
	CardState    sql.NullString `json:"CardState" db:"CardState"`
	StepCode     sql.NullString `json:"StepCode" db:"StepCode"`
	ByWay        sql.NullString `json:"ByWay" db:"ByWay"`
	ByWayDate    mysql.NullTime `json:"ByWayDate" db:"ByWayDate"`
	IfStop       sql.NullString `json:"IfStop" db:"IfStop"`
	StopRemark   sql.NullString `json:"StopRemark" db:"StopRemark"`
	PlanDate     mysql.NullTime `json:"PlanDate" db:"PlanDate"`
	CutDate      mysql.NullTime `json:"CutDate" db:"CutDate"`
	SewingDate   mysql.NullTime `json:"SewingDate" db:"SewingDate"`
	LroningDate  mysql.NullTime `json:"LroningDate" db:"LroningDate"`
	PackingDate  mysql.NullTime `json:"PackingDate" db:"PackingDate"`
	DeliveryDate mysql.NullTime `json:"DeliveryDate" db:"DeliveryDate"`
	TenantId     sql.NullString `json:"tenantId" db:"tenantId"`
}

//查询单行
/*func QueryOne(DB *sqlx.DB) (p ProOrder) {
	var order ProOrder
	sqlText := "select a.tenantId,a.Odid, a.SysCode, a.OrderCode, a.Customer,  a.MesSort Sort, a.Counts,a.FabricNo, a.StyleNo, a.Finished, a.CardState, a.StepCode, case(a.ByWay) when 1 then '已过通道' else '未过通道' end ByWay, a.ByWayDate, a.IfStop, a.StopRemark, a.PlanDate, a.CutDate, a.SewingDate, a.LroningDate, a.PackingDate, a.DeliveryDate ,b.EmpCode,c.EmployeeName ,c.DeptCode,c.TeamCode,d.TeamName,e.DeptName from  ProOrdtemp a inner join  BrushCard_Material b on a.Odid =b.Odid  and a.StepCode=b.StepCode inner join  HR_Employee c on   b.EmpCode=c.EmployeeCode  inner join  HR_Team d on   c.TeamCode=d.TeamCode inner join  HR_Depart e on   e.DeptCode=c.DeptCode"

	fmt.Println("---------", sqlText)
	row := DB.QueryRow(sqlText)
	if err := row.Scan(&order.tenantId, &order.Odid, &order.SysCode, &order.OrderCode, &order.Customer, &order.Sort, &order.Counts, &order.FabricNo, &order.StyleNo, &order.Finished, &order.CardState, &order.StepCode, &order.ByWay, &order.ByWayDate, &order.IfStop, &order.StopRemark, &order.PlanDate, &order.CutDate, &order.SewingDate, &order.LroningDate, &order.PackingDate, &order.DeliveryDate, &order.EmpCode, &order.EmployeeName, &order.DeptCode, &order.TeamCode, &order.TeamName, &order.DeptName); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(">>>>>>", order)
	p = order
	return
}*/

//查询多行
func QueryMultiProOrder(DB *sqlx.DB, tenantId string, Isbulk string, orderCode string, customer string, sort string, deliveryDateFrom string, deliveryDateTo string, pageSize int, currentPage int) (orderprofList []map[string]interface{}) {
	fmt.Println("传入参数：", pageSize, currentPage)
	var toatalpieces int
	var order ProOrder
	var prams []interface{}
	var builderSQLTotal strings.Builder
	var builderSQL strings.Builder
	builderSQLTotal.WriteString(" select  count(1) from  ProductOrder p   inner join  OrderDetail o on p.SysCode =o.SysCode and  p.tenantId=o.tenantId where 1=1 ")
	builderSQL.WriteString(" select o.Odid, p.SysCode, p.OrderCode, p.Customer, o.MesSort, p.Counts,o.FabricNo, o.StyleNo, o.Finished, o.CardState, o.StepCode, case(o.ByWay) when 1 then '已过通道' else '未过通道' end ByWay, o.ByWayDate, o.IfStop, o.StopRemark, p.PlanDate, p.CutDate, p.SewingDate, p.LroningDate, p.PackingDate, p.DeliveryDate,p.tenantId ")
	builderSQL.WriteString(" from  ProductOrder p  ")
	builderSQL.WriteString(" inner join  OrderDetail o on p.SysCode =o.SysCode and p.tenantId=o.tenantId where 1=1 ")
	if tenantId != "" {
		builderSQLTotal.WriteString(" and o.tenantId= ? ")
		builderSQL.WriteString(" and o.tenantId= ? ")
		prams = append(prams, tenantId)
	}
	if Isbulk != "" {
		builderSQLTotal.WriteString(" and  p.Isbulk= ? ")
		builderSQL.WriteString(" and  p.Isbulk= ? ")
		prams = append(prams, Isbulk)
	}
	if orderCode != "" {
		builderSQLTotal.WriteString(" and  OrderCode like CONCAT('%',?,'%')")
		builderSQL.WriteString(" and  OrderCode like CONCAT('%',?,'%')")
		prams = append(prams, orderCode)
	}
	if customer != "" {
		builderSQLTotal.WriteString(" and Customer like  CONCAT('%',?,'%')")
		builderSQL.WriteString(" and Customer like  CONCAT('%',?,'%')")
		prams = append(prams, customer)
	}
	if sort != "" {
		builderSQLTotal.WriteString(" and Sort like  CONCAT('%',?,'%')")
		builderSQL.WriteString(" and Sort like  CONCAT('%',?,'%')")
		prams = append(prams, sort)
	}
	if deliveryDateFrom != "" {
		builderSQLTotal.WriteString(" and  DeliveryDate >= ?")
		builderSQL.WriteString(" and  DeliveryDate >= ?")
		prams = append(prams, deliveryDateFrom)
	}
	if deliveryDateTo != "" {
		builderSQLTotal.WriteString(" and DeliveryDate <= ?")
		builderSQL.WriteString(" and DeliveryDate <= ?")
		prams = append(prams, deliveryDateTo)
	}
	//2019-12-31 待优化
	rowcount, er := DB.Query(builderSQLTotal.String(), prams...)
	fmt.Println("总数：", rowcount)
	logprint.Logprint("builderSQLTotal", builderSQLTotal.String())
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
		err = rows.Scan(&order.Odid, &order.SysCode, &order.OrderCode, &order.Customer, &order.MesSort, &order.Counts, &order.FabricNo, &order.StyleNo, &order.Finished, &order.CardState, &order.StepCode, &order.ByWay, &order.ByWayDate, &order.IfStop, &order.StopRemark, &order.PlanDate, &order.CutDate, &order.SewingDate, &order.LroningDate, &order.PackingDate, &order.DeliveryDate, &order.TenantId)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		map_sg := pkg.StructToMapDemo(order)
		orderprofList = append(orderprofList, map_sg)

	}
	fmt.Println("---------结果实体——————", orderprofList)
	return
}
