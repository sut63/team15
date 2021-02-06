import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';

import { Content, pageTheme, Page } from '@backstage/core';
import { EntCleaningroom } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsCleaningroomTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [cleaningrooms, setCleaningrooms] = useState<EntCleaningroom[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
    const getCleaningrooms = async () => {
      const res = await api.listCleaningroom();
      setLoading(true);
      setCleaningrooms(res);
      console.log(res);
    };
    getCleaningrooms();
  }, [loading]);
 
 
 return (
    <Page theme={pageTheme.service}>
     <Content>
       
    <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">ชื่อพนักงานที่ทำการบันทึก</TableCell>
           <TableCell align="center">ชื่อพนักงานที่รับผิดชอบ</TableCell>
           <TableCell align="center">จำนวนพนักงานในการทำความสะอาด</TableCell>
           <TableCell align="center">เลขห้อง</TableCell>
           <TableCell align="center">ระยะเวลา</TableCell>
           <TableCell align="center">วันที่และเวลาเริ่มทำความสะอาด</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์ที่สามารถติดต่อได้</TableCell>
           <TableCell align="center">เพิ่มเติม</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {cleaningrooms.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.edges?.cleanername?.cleanername}</TableCell>
             <TableCell align="center">{item.numofem}</TableCell>
             <TableCell align="center">{item.edges?.roomdetail?.roomnumber}</TableCell>
             <TableCell align="center">{item.edges?.lengthtime?.lengthtime}</TableCell>
             <TableCell align="center">{item.dateandstarttime}</TableCell>
             <TableCell align="center">{item.phonenumber}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">
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