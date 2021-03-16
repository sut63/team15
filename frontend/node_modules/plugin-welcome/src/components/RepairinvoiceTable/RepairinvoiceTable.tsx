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

import { EntRepairinvoice } from '../../api/models/EntRepairinvoice';
import { EntEmployee } from '../../api/models/EntEmployee';
import { InfoCard, Content, Header, pageTheme, Page } from '@backstage/core';
import { EntRentalstatus } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [repairinvoices, setRepairinvoices] = useState<EntRepairinvoice[]>([]);
 const [employees, setEmployees] = useState<EntEmployee[]>([]);
 const [loading, setLoading] = useState(true);

 const [employeeid, setEmployee] = useState(Number);
 
 useEffect(() => {
   const getRepairinvoices = async () => {
     const res = await api.listRepairinvoice({ offset: 0 });
     setLoading(false);
     setRepairinvoices(res);
   };
   getRepairinvoices();
 }, [loading]);
 
 
 return (
    <Page theme={pageTheme.service}>
     <Content>
 <InfoCard> <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">พนักงาน</TableCell>
           <TableCell align="center">อุปกรณ์ที่ชำรุด</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์ของพนักงาน</TableCell>
           <TableCell align="center">จำนวนอุปกรณ์ที่ชำรุด</TableCell>
           <TableCell align="center">สถานะการเข้าพัก</TableCell>
           
         </TableRow>
       </TableHead>
       <TableBody>
         {repairinvoices.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.bequipment}</TableCell>
             <TableCell align="center">{item.emtell}</TableCell>
             <TableCell align="center">{item.num}</TableCell>
             <TableCell align="center">{item.edges?.rentalstatus?.rentalstatus}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
</InfoCard></Content>
</Page>
  
 );
}