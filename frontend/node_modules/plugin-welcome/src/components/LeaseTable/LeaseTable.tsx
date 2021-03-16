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
import { ControllersLease, EntLease } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsLeaseTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [leases, setLeases] = useState<EntLease[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getLeases = async () => {
      const res = await http.listLease({ offset: 0 });
      setLoading(true);
      setLeases(res);
      setLoading(true);
      console.log(res);
    };
    getLeases();
  }, [loading]);
  
 
  
 return (
  <Page theme={pageTheme.tool}>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">Contract No.</TableCell>
           <TableCell align="center">Room no.</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Age</TableCell>
           <TableCell align="center">Identified ID</TableCell>
           <TableCell align="center">Phone No.</TableCell>
           <TableCell align="center">Wifi Status</TableCell>
		   <TableCell align="center">Time</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {leases.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.roomdetail?.roomnumber}</TableCell>
             <TableCell align="center">{item.tenant}</TableCell>
             <TableCell align="center">{item.agetenant}</TableCell>
             <TableCell align="center">{item.idtenant}</TableCell>
             <TableCell align="center">{item.numbtenant}</TableCell>
             <TableCell align="center">{item.edges?.wifi?.wifiname}</TableCell>
			 <TableCell align="center">{item.addedtime}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}