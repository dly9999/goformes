

-- 裁剪生产数据查询(牛晓东)
create TEMPORARY table ProductOrdertemp(select * from  ProductOrder where {0}); 
create TEMPORARY table OrderDetailtemp(select * from  OrderDetail where {0});

select distinct    
			a.Customer, a.CutDate, a.OrderCode, b.Odid, a.PlanCode, 
            b.MesSort, b.ByWay, b.ByWayDate,  case  b.IfStop when 1 then '暂停生产' else '正在生产' end as IfStop, b.StopRemark, b.CardState, b.StepCode,-- b.CutFinishDate b.Finished,
            c.EmpCode,c.LoadDate,
            d.TeamCode,d.DeptCode,d.EmployeeName,
            e.TeamName,
            f.DeptName
from 		ProductOrdertemp a 
inner join  OrderDetailtemp b on a.SysCode=b.SysCode
inner join  BrushCard_Material c on c.odid=b.odid and c.stepcode=b.stepcode
inner join  HR_Employee d on c.EmpCode=d.EmployeeCode
inner join  HR_Team e on e.TeamCode=d.TeamCode
inner join  HR_Depart f on f.DeptCode=d.DeptCode
where  {0};
drop table ProductOrdertemp;
drop table OrderDetailtemp;

--缝制生产数据查询(牛晓东)

create TEMPORARY table ProductOrdertemp(select * from  ProductOrder where {0});  
create TEMPORARY table OrderDetailtemp(select * from  OrderDetail where {0});  

select distinct    a.Customer, a.SewingDate, a.OrderCode, b.Odid, a.PlanCode, 
                b.MesSort, b.ByWay, b.ByWayDate,  case  b.IfStop when 1 then '暂停生产' else '正在生产' end as IfStop, b.StopRemark, b.CardState, b.StepCode,-- b.SewingFinishDate Finished,
                c.EmpCode,c.LoadDate,
                d.TeamCode,d.DeptCode,d.EmployeeName,
                e.TeamName,
                f.DeptName
from 		ProductOrdertemp a 
inner join  OrderDetailtemp b on a.SysCode=b.SysCode
inner join  BrushCard_Material c on c.odid=b.odid and c.stepcode=b.stepcode
inner join  HR_Employee d on c.EmpCode=d.EmployeeCode
inner join  HR_Team e on e.TeamCode=d.TeamCode
inner join  HR_Depart f on f.DeptCode=d.DeptCode
where  {0};
drop table ProductOrdertemp;
drop table OrderDetailtemp;

--整烫生产数据查询(牛晓东)
create TEMPORARY table ProductOrdertemp(select * from  ProductOrder where {0});  
create TEMPORARY table OrderDetailtemp(select * from  OrderDetail where {0});  

select distinct    a.Customer, a.LroningDate , a.OrderCode, b.Odid, a.PlanCode, 
                b.MesSort, b.ByWay, b.ByWayDate,  case  b.IfStop when 1 then '暂停生产' else '正在生产' end as IfStop, b.StopRemark, b.CardState, b.StepCode, -- b.LroningFinishDate Finished
                c.EmpCode,c.LoadDate,
                d.TeamCode,d.DeptCode,d.EmployeeName,
                e.TeamName,
                f.DeptName
from 		ProductOrdertemp a 
inner join  OrderDetailtemp b on a.SysCode=b.SysCode
inner join  BrushCard_Material c on c.odid=b.odid and c.stepcode=b.stepcode
inner join  HR_Employee d on c.EmpCode=d.EmployeeCode
inner join  HR_Team e on e.TeamCode=d.TeamCode
inner join  HR_Depart f on f.DeptCode=d.DeptCode
where  {0};
drop table ProductOrdertemp;
drop table OrderDetailtemp;

--包装生产数据查询(苟晓军)
create TEMPORARY table ProductOrdertemp(select * from  ProductOrder where {0});  
create TEMPORARY table OrderDetailtemp(select * from  OrderDetail where {0}); 

select distinct    a.Customer, a.PackingDate PackingDate, a.OrderCode, b.Odid, a.PlanCode, 
                b.MesSort, b.ByWay, b.ByWayDate,  case  b.IfStop when 1 then '暂停生产' else '正在生产' end as IfStop, b.StopRemark, b.CardState, b.StepCode,-- b.PackingFinishDate Finished,
                c.EmpCode,c.LoadDate,
                d.TeamCode,d.DeptCode,d.EmployeeName,
                e.TeamName,
                f.DeptName
from 		ProductOrdertemp a 
inner join  OrderDetailtemp b on a.SysCode=b.SysCode
inner join  BrushCard_Material c on c.odid=b.odid and c.stepcode=b.stepcode
inner join  HR_Employee d on c.EmpCode=d.EmployeeCode
inner join  HR_Team e on e.TeamCode=d.TeamCode
inner join  HR_Depart f on f.DeptCode=d.DeptCode
where  {0};
drop table ProductOrdertemp;
drop table OrderDetailtemp;


--订单状态信息查询(苟晓军)

create TEMPORARY table ProOrdtemp
(
        select o.Odid, p.SysCode, p.OrderCode, p.Customer, o.MesSort Sort, p.Counts,o.FabricNo, o.StyleNo, o.Finished, o.CardState, o.StepCode, case(o.ByWay) when 1 then '已过通道' else '未过通道' end ByWay, o.ByWayDate, o.IfStop, o.StopRemark, p.PlanDate, p.CutDate, p.SewingDate, p.LroningDate, p.PackingDate, p.DeliveryDate 
        from  ProductOrder p inner join  OrderDetail o on p.SysCode =o.SysCode and p.tenantId=o.tenantId where {0}
);  
select a.Odid, a.SysCode, a.OrderCode, a.Customer,  a.Sort, a.Counts,a.FabricNo, a.StyleNo, a.Finished, a.CardState, a.StepCode, case(a.ByWay) when 1 then '已过通道' else '未过通道' end ByWay, a.ByWayDate, a.IfStop, a.StopRemark, a.PlanDate, a.CutDate, a.SewingDate, a.LroningDate, a.PackingDate, a.DeliveryDate ,b.EmpCode,c.EmployeeName ,c.DeptCode,c.TeamCode,d.TeamName,e.DeptName
from  ProOrdtemp a 
    inner join  BrushCard_Material b on a.Odid =b.Odid  and a.StepCode=b.StepCode
    inner join  HR_Employee c on   b.EmpCode=c.EmployeeCode
    inner join  HR_Team d on   c.TeamCode=d.TeamCode
    inner join  HR_Depart e on   e.DeptCode=c.DeptCode;
drop table ProOrdtemp;
                        
                       

--where  1=1 and p.Isbulk<>'1' and p.tenantId='{0}'", base.tenantId);


-- 订单刷卡信息查询 (苟晓军)
create TEMPORARY table ProOrdtemp(select a.tenantId,a.CutDate,a.OrderCode, b.Odid,b.MesSort,b.stepcode from  ProductOrder a inner join  OrderDetail b on a.SysCode =b.SysCode and a.tenantId=b.tenantId where {0});  
create TEMPORARY table HR_Employeetemp(select * from  HR_Employee where {0});  

select distinct    a.CutDate,a.OrderCode, 
                a.Odid,a.MesSort,
                c.EmpCode,c.StepCode,c.LoadDate,{0} OdState,
                d.TeamCode,d.EmployeeName,
                e.TeamName
                
from 		ProOrdtemp a 
inner join  {1} c on c.odid=a.odid and c.stepcode=a.stepcode
inner join  HR_Employeetemp d on c.EmpCode=d.EmployeeCode
inner join  HR_Team e on e.TeamCode=d.TeamCode

where  {2};
drop table ProOrdtemp;
drop table HR_Employeetemp;

-- 返工记录查询（苟晓军）
create TEMPORARY table ProOrdertemp( select a.PlanCode, a.OrderCode, b.Odid,b.MesSort 
                    from  ProductOrder a inner join OrderDetail b on a.SysCode=b.SysCode  and a.tenantId=b.tenantId  where {0}
                  ); 
create TEMPORARY table QC_4(
        select a.Odid,b.QCCheckDate,b.EmployCode, b.EmployeeName Employee,
               b.QCCheckID,b.StepCode,d.OutQualityID,d.DutyGX,d.Memo1,e.BrushDate,e.EmpCode
        from  		   ProOrdertemp a 
            inner join QC_CheckMain  b on a.Odid=b.Odid
            inner join QC_OutCheckReason c on b.IsOutCheck=false and b.QCCheckID=c.QCCheckID
            inner join QC_OutQuality d on d.Memo6=false and d.OutQualityID=c.OutQualityID
            inner join BrushCard_Material e on e.StepCode=DutyGX and e.Odid=a.Odid
            
          where {0}
       );

create TEMPORARY table DeptTeamEmp(
                        select c.EmployeeName,a.DeptCode,b.TeamCode,a.DeptName,b.TeamName,c.EmployeeCode
                        from  	   HR_Depart a 
                        inner join HR_Team b on a.DeptCode=b.DeptCode
                        inner join HR_Employee c on a.DeptCode=c.DeptCode and  b.TeamCode=c.TeamCode
                        
                    ); 

select distinct    a.PlanCode, a.OrderCode, a.MesSort,
b.Odid,b.QCCheckDate,b.Employee,b.QCCheckID,b.StepCode,b.OutQualityID,b.DutyGX,b.Memo1,b.BrushDate,b.EmpCode,
c.EmployeeName,c.DeptCode,c.TeamCode,c.DeptName,c.TeamName

from 		ProOrdertemp a 
inner join  QC_4         b on a.Odid=b.Odid
inner join  DeptTeamEmp c on c.EmployeeCode=b.EmpCode;

drop table ProOrdertemp;
drop table QC_4;
drop table DeptTeamEmp;


--车间暂压查询 （王晓明）
create TEMPORARY table workoverstockTemp(select * from  Work_TeamOverStock where {0});
select a.*,b.DeptName
from  workoverstockTemp a
	inner join HR_Depart b on a.OrganizeCode=b.DeptCode
where  {1} ;

drop table workoverstockTemp;




--班组暂压查询  （王晓明）
create TEMPORARY table workoverstockTemp(select * from  Work_TeamOverStock where {0});
select a.*, c.DeptCode,c.DeptName
from  workoverstockTemp a
	inner join HR_Team b on a.OrganizeCode=b.DeptCode
	inner join HR_Depart c on b.DeptCode=c.DeptCode
where  {1} ;

drop table workoverstockTemp;


--车间产量  （王晓明）
create TEMPORARY table workcountTemp(select * from  {0} where {1});
select a.*
from  workcountTemp a
where  {2} ;

drop table workcountTemp;

--班组产量  （王晓明）
create TEMPORARY table workcountTemp(select * from  {0} where {1});
select a.*,c.DeptName
from  workcountTemp a
	inner join HR_Team b on a.OrgCode=b.DeptCode
	inner join HR_Depart c on b.DeptCode=c.DeptCode
where  {2} ;

drop table workcountTemp;

--人员产量    （孙娜）

create TEMPORARY table workcountTemp(select * from  {0} where {1});
select a.*,c.DeptCode,d.TeamCode
from  workcountTemp a 
inner join  HR_Employee b on a.OrgCode=b.EmployeeCode
inner join  HR_Depart c on b.DeptCode=c.DeptCode
inner join  HR_Team d on b.TeamCode=d.TeamCode
where  {2} ;

drop table workcountTemp;

--质量缺陷查询 （孙娜）
create TEMPORARY table OdidTabletemp(select Odid from  ProductOrder a inner join OrderDetail b on a.SysCode=b.SysCode where {0});  
create TEMPORARY table ProOrdertemp( select b.FabricNo, a.OrderCode, b.Sort, a.IsBulk, a.SysCode, a.CutDate, b.Odid 
                                    from  ProductOrder a inner join OrderDetail b on a.SysCode=b.SysCode where {0}
                                  );  
create TEMPORARY table QC_4(
                        select a.Odid,b.QCCheckDate,b.EmployCode, b.EmployeeName Employee,b.QCCheckID,b.StepCode,d.OutQualityID,d.DutyGX,d.Memo1,f.OutTypeName,e.BrushDate,e.EmpCode
                        from  		   OdidTabletemp a 
                            inner join QC_CheckMain  b on a.Odid=b.Odid
                            inner join QC_OutCheckReason c on b.IsOutCheck=false and b.QCCheckID=c.QCCheckID
                            inner join QC_OutQuality d on d.Memo6=false and d.OutQualityID=c.OutQualityID
                            inner join BrushCard_Material e on e.StepCode=DutyGX and e.Odid=a.Odid
                            left  join QC_OutType f on d.Memo1=f.OutTypeID
                            
                          where {0}
                       );

create TEMPORARY table DeptTeamEmp(
                                        select c.EmployeeName,c.DeptCode,c.TeamCode,a.DeptName,b.TeamName,c.EmployeeCode
                                        from  	   HR_Depart a 
                                        inner join HR_Team b on a.DeptCode=b.DeptCode
                                        inner join HR_Employee c on a.DeptCode=c.DeptCode and  b.TeamCode=c.TeamCode
                                        where {0}
                                    );  

select distinct    a.OrderCode, a.IsBulk, a.SysCode, a.CutDate,a.FabricNo, a.Odid,a.Sort,
                b.QCCheckDate,b.Employee,b.QCCheckID,b.StepCode,b.OutQualityID,b.DutyGX,b.Memo1,b.OutTypeName,b.BrushDate,b.EmpCode,
                c.EmployeeName,c.DeptCode,c.TeamCode,c.DeptName,c.TeamName

from 		ProOrdertemp a 
inner join  QC_4         b on a.Odid=b.Odid
inner join  DeptTeamEmp c on c.EmployeeCode=b.EmpCode;

drop table OdidTabletemp;
drop table ProOrdertemp;
drop table QC_4;
drop table DeptTeamEmp;



-------订单状态维护 （孙娜）
create TEMPORARY table ProOrdtemp
(
        select o.Odid, p.SysCode, p.OrderCode, p.Customer, o.MesSort, p.Counts,o.FabricNo, o.StyleNo, o.Finished, o.CardState, o.StepCode, case(o.ByWay) when 1 then '已过通道' else '未过通道' end ByWay, o.ByWayDate, o.IfStop, o.StopRemark, p.PlanDate, p.CutDate, p.SewingDate, p.LroningDate, p.PackingDate, p.DeliveryDate,p.tenantId
        from  ProductOrder p inner join  OrderDetail o on p.SysCode =o.SysCode and p.tenantId=o.tenantId where {0}
);  
select a.tenantId,a.Odid, a.SysCode, a.OrderCode, a.Customer,  a.MesSort Sort, a.Counts,a.FabricNo, a.StyleNo, a.Finished, a.CardState, a.StepCode, case(a.ByWay) when 1 then '已过通道' else '未过通道' end ByWay, a.ByWayDate, a.IfStop, a.StopRemark, a.PlanDate, a.CutDate, a.SewingDate, a.LroningDate, a.PackingDate, a.DeliveryDate ,b.EmpCode,c.EmployeeName ,c.DeptCode,c.TeamCode,d.TeamName,e.DeptName
        from  ProOrdtemp a 
            inner join  BrushCard_Material b on a.Odid =b.Odid  and a.StepCode=b.StepCode
            inner join  HR_Employee c on   b.EmpCode=c.EmployeeCode
            inner join  HR_Team d on   c.TeamCode=d.TeamCode
            inner join  HR_Depart e on   e.DeptCode=c.DeptCode;
drop table ProOrdtemp;

-------质量缺陷维护
--对QC_OutQuality表的操作


--下计划 （王华）

--查询
select 
tenantId,SysCode,OrderCode,Customer,PlanCode,DeliveryDate,CutDate,SewingDate,LroningDate,PackingDate,OrderStatus,Counts,PlanDate,IsBulk,ByWay,ByWayDate,InsertDate,SwetDept,CutDept,CutSpanTime,Cutter,PlanState,Remark,AddPlanUser,AddPlanTime,AddCutUser,AddCutTime 
FROM SCT_51
--下计划所下车间
select * FROM PlanDepart

--下计划更新sct51信息
--  update SCT_51 set  PlanCode= '{0}',CutDate='{1}',SewingDate='{2}',LroningDate='{3}',PackingDate='{4}',AddPlanUser='{5}', AddPlanTime='{6}',PlanState='{7}',Remark='{8}',OrderStatus='{9}',SwetDept='{10}',PlanDate='{11}'    where  SysCode ='{12}' and tenantId='{13}'


--拍裁床 （王华）

--查询
select  a.tenantId, a.SysCode, a.OrderCode, a.DeliveryDate, a.CutDate, a.SewingDate, a.LroningDate, a.PackingDate, a.PlanState, a.Remark, a.Cutter, a.SwetDept, a.CutDept,a.CutSpanTime, a.OrderStatus, b.P_DeptName  
FROM SCT_51 a left join PlanDepart b on a.Remark=b.P_DeptCode and a.tenantId=b.tenantId 

--拍裁床 （王华）

--更新sct51信息，同时把订单信息插入到orderdetail 和productorder
-- update SCT_51 set  Cutter= '{0}',CutDept='{1}',AddCutTime='{2}',AddCutUser='{3}',CutSpanTime='{4}',OrderStatus='{5}' where  SysCode ='{6}' and tenantId='{7}'
-- insert into OrderDetail (tenantId,SysCode,StyleNo,StyleName,Sort,MesSort,FabricNo,LineNum,BuckleNo,CuffNo,LocklineNo,Inneredge,Outeredge,LiningType,Odid,Finished,SpecialPlate,SpecialInfor,ClothingSpec,Count,Size) 
--                            values( '{0}','{1}','{2}','{3}','{4}','{5}','{6}','{7}','{8}','{9}','{10}','{11}','{12}','{13}','{14}','{15}','{16}','{17}','{18}','{19}','{20}'
--                            )
--insert into ProductOrder (tenantId,SysCode,OrderCode,Customer,PlanCode,DeliveryDate,CutDate,SewingDate,LroningDate,PackingDate,OrderStatus,Counts,PlanDate,IsBulk,InsertDate) 
--                            values( '{0}','{1}','{2}','{3}','{4}','{5}','{6}','{7}','{8}','{9}','{10}','{11}','{12}',{13},'{14}'
--                            ) 



--计划更改  （王华）

-- 查询
select 
a.tenantId, a.SysCode, a.OrderCode, a.DeliveryDate, a.CutDate, a.SewingDate, a.LroningDate, a.PackingDate, a.PlanState, a.Remark, a.Cutter, a.SwetDept, a.CutDept,a.CutSpanTime, a.OrderStatus, b.P_DeptName   
FROM SCT_51 a left join PlanDepart b on a.Remark=b.P_DeptCode and a.tenantId=b.tenantId

--备份数据
-- insert into SCT_51Modify (tenantId,SysCode,OrderCode,CutDate,SewingDate,LroningDate,PackingDate,SwetDept,CutDept,CutSpanTime,Cutter,DelTime,DelUser) values('{0}','{1}','{2}','{3}','{4}','{5}','{6}','{7}','{8}','{9}','{10}','{11}','{12}')
--更新sct51信息，同时更新productorder信息
-- update SCT_51 set CutDate='{0}',SewingDate='{1}',LroningDate='{2}',PackingDate='{3}',PlanState='{4}',Remark='{5}',Cutter='{6}',CutDept='{7}',CutSpanTime='{8}',SwetDept='{9}'  where SysCode='{10}' and tenantId='{11}'
-- update ProductOrder set  CutDate='{0}',SewingDate='{1}',LroningDate='{2}',PackingDate='{3}' where SysCode='{4}' and tenantId='{5}'


-- 工厂日历 （王华）
select * from  PlanDateMaintain

--维护工厂日历，修改是否休息，都是操作PlanDateMaintain



--用户管理
--增删改查 操作   UserInfo 
--权限管理
--增删改查 操作   GroupInfo 
--主菜单管理
--增删改查 操作   Parent_Menu 
--子菜单管理
--增删改查 操作   Second_Menu_New 



















