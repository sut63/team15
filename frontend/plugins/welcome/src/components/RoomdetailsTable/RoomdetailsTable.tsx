import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import { EntEmployee } from '../../api/models/EntEmployee';
import { InfoCard, Content, Header, pageTheme, Page } from '@backstage/core';
import { EntCleanerName, EntEquipment, EntFacilitie, EntNearbyplace } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [roomdetails, setRoomdetails] = useState<EntRoomdetail[]>([]);
 const [employees, setEmployees] = useState<EntEmployee[]>([]);
 const [loading, setLoading] = useState(true);

 const [employeeid, setEmployee] = useState(Number);
 
 useEffect(() => {
   const getRoomdetais = async () => {
     const res = await api.listRoomdetail();
     setLoading(false);
     setRoomdetails(res);
   };
   getRoomdetais();

   const getEmployees = async () => {
 
    const em = await api.listEmployee();
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();

   const checkEmployeeLoginData = async () => {
    const logindata = JSON.parse(String(localStorage.getItem("employeelogindata")));
    setLoading(false);
    if(logindata != "โรเจอร์"){
      localStorage.setItem("employeedata",JSON.stringify(null));
      localStorage.setItem("employeelogindata",JSON.stringify(null));
      history.pushState("","","./");
      window.location.reload(false);     
    }
    else{
      setEmployee(Number(localStorage.getItem("employeedata")))
    }
  }
  checkEmployeeLoginData();

 }, [loading]);
 
 const deleteRoomdetail = async (id: number) => {
  const res = await api.deleteRoomdetail({ id: id });
  setLoading(true);
};
 
 return (
    <Page theme={pageTheme.service}>
    <Header title="Room Detail table" subtitle="ตารางรายละเอียดข้อมูลห้องพัก">
    <Button
                   style={{ width: 230, backgroundColor: "#ADF279",marginLeft: 25}}
                component={RouterLink}
                to="/RoomDetails"
                variant="contained"
              >
                <strong>เพิ่มรายละเอียดห้องพักเข้าสู่ระบบ</strong>
             </Button>

</Header>
     <Content>
 <InfoCard> <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">ประเภทห้อง</TableCell>
           <TableCell align="center">ราคาห้อง</TableCell>
           <TableCell align="center">จำนวนผู้ที่สามารถเข้าพักได้</TableCell>
           <TableCell align="center">ประเภทการเข้าพัก</TableCell>
           <TableCell align="center">สิ่งที่ติดตั้งให้ภายในห้อง</TableCell>
           <TableCell align="center">สิ่งอำนวยความสะดวกทั่วไป</TableCell>
           <TableCell align="center">สถานที่ใกล้เคียง</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {roomdetails.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.roomtypename}</TableCell>
             <TableCell align="center">{item.roomprice}</TableCell>
             <TableCell align="center">{item.edges?.quantity?.quantity}</TableCell>
             <TableCell align="center">{item.edges?.staytype?.staytype}</TableCell>
             <TableCell align="center">{item.edges?.equipments?.equipment}</TableCell>
             <TableCell align="center">{item.edges?.facilities?.facilitie}</TableCell>
             <TableCell align="center">{item.edges?.nearbyplaces?.nearbyplace}</TableCell>
             <TableCell align="center">
               <Button 
               onClick={() => {
                   deleteRoomdetail(item.id);
                 }}
                 style={{ marginLeft: 10 , backgroundColor: "#140087" }}
                 variant="contained"
                 color="secondary"
               >
                 ลบข้อมูล
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
</InfoCard></Content>
</Page>
  
 );
}