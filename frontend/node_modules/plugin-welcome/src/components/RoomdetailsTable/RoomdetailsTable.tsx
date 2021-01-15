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
  

 }, [loading]);
 
 const deleteRoomdetail = async (id: number) => {
  const res = await api.deleteRoomdetail({ id: id });
  setLoading(true);
};
 
 return (
    <Page theme={pageTheme.service}>
     <Content>
       
    <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขห้อง</TableCell>
           <TableCell align="center">ชื่อประเภทห้อง</TableCell>
           <TableCell align="center">ราคาห้อง/ประเภทการเข้าพัก</TableCell>
           <TableCell align="center">จำนวนผู้เข้าพักสูงสุด</TableCell>
           <TableCell align="center">ประเภทเตียงนอน/จำนวนเตียง</TableCell>
           <TableCell align="center">ประเภทการเข้าพัก</TableCell>
           <TableCell align="center">ข้อจำกัดการจ่ายมัดจำ</TableCell>
           <TableCell align="center">ลบข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {roomdetails.map((item:any ) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.roomnumber}</TableCell>
             <TableCell align="center">{item.roomtypename}</TableCell>
             <TableCell align="center">{item.roomprice} / {item.edges?.staytype?.staytype}</TableCell>
             <TableCell align="center">{item.sleep} คน</TableCell>
             <TableCell align="center">{item.edges?.bedtype?.bedtypename} / {item.bed} เตียง</TableCell>
             <TableCell align="center">{item.edges?.pledge?.provision}</TableCell>
             <TableCell align="center">{item.edges?.petrule?.petrule}</TableCell>
             <TableCell align="center">
               <Button 
               onClick={() => {
                   deleteRoomdetail(item.id);
                 }}
                 style={{ marginLeft: 10 , backgroundColor: "#adf279" }}
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
 </Content>
</Page>
  
 );
}