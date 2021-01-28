import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { ControllersBill, EntBill, EntLease, EntRoomdetail } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsBillTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [Bills, setBills] = useState<EntBill[]>([]);
  const [roomdetail, setRoomdetails] = useState<EntRoomdetail[]>([]);
  const [leases, setLeases] = useState<EntLease[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getBills = async () => {
      const res = await http.listBill({ offset: 0 });

      setLoading(true);
      setBills(res);
      console.log(res);
    };
    getBills();
  }, [loading]);
 
  
 return (
  <Page theme={pageTheme.tool}>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Room</TableCell>
           <TableCell align="center">Total</TableCell>
           <TableCell align="center">Situation</TableCell>
           <TableCell align="center">Payment</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {Bills.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell> 
             <TableCell align="center">{item.edges?.lease?.tenant}</TableCell> 
             {leases.filter((setfilterid:any) => setfilterid.id === item.edges?.lease?.id).map((item2:any) => (
                  <TableCell align="center">{item2.edges?.roomdetail?.roomnumber}</TableCell>
              ))}
               <TableCell align="center">{item.total}</TableCell>
             <TableCell align="center">{item.edges?.situation?.situationname}</TableCell>
             <TableCell align="center">{item.edges?.payment?.paymentname}</TableCell>
            
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}