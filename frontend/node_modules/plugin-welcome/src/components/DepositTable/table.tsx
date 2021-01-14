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
import { ControllersDeposit, EntDeposit } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsDepositTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [deposits, setDeposits] = useState<EntDeposit[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getDeposits = async () => {
      const res = await http.listDeposit({ offset: 0 });
	  const edata = JSON.parse(Number(localStorage.getItem("employeedata")));
      setLoading(true);
      setDeposits(res);
      console.log(res);
    };
    getDeposits();
  }, [loading]);
 
  
 return (
  <Page theme={pageTheme.tool}>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Info</TableCell>
           <TableCell align="center">Employees</TableCell>
           <TableCell align="center">Status</TableCell>
		   <TableCell align="center">Time</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {deposits.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.info}</TableCell>
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.edges?.statusd?.statusdname}</TableCell>
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