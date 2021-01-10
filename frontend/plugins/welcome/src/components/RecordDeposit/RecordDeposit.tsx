import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';
import { Alert, AlertTitle } from '@material-ui/lab';

import React, { useState, useEffect, FC } from 'react';  
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
  Typography,
  Toolbar,
  AppBar,
} from '@material-ui/core';

import { DefaultApi } from '../../api/apis';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntStatusd } from '../../api/models/EntStatusd'; // import interface Statusd
import { EntDeposit } from '../../api/models/EntDeposit'; // import interface Deposit

// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
      //marginTop:10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

export default function recordDeposit() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [deposit, setDeposit] = React.useState<Partial<recordDeposit>>({});

 const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
 const [statusds, setStatusds] = React.useState<EntStatusd[]>([]);

 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(false);

 const [loading, setLoading] = useState(true);

 const [employee, setEmployee] = useState(Number);
 const [info, setInfo] = useState(String);
 const [added, setAdded] = useState(String);
 const [statusd, setStatusd] = useState(Number);

 useEffect(() => {
  const getEmployees = async () => {
    const res = await http.listEmployee({ offset: 0 });
    setLoading(false);
    setEmployees(res);
    console.log(res);
  };
  getEmployees();

  const getStatusds = async () => {
    const res = await http.listStatusd({ offset: 0 });
    setLoading(false);
    setStatusds(res);
    console.log(res);
  };
  getStatusds();

  const checkEmployee = async () => {
    const edata = JSON.parse(Number(localStorage.getItem("employeedata")));
    setLoading(false);
    setEmployee(edata)
	console.log(edata);
  };
  checkEmployee();

}, [loading]);


const getDeposit = async () => {
  const res = await http.listDeposit({ offset: 0 });
  setDeposit(res);
};

  const InfohandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>,) => {
    setInfo(event.target.value as string);
  };

  const AddedhandleChange = (event: any) => {
    setAdded(event.target.value as string);
  };

  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };
  const StatusdhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatusd(event.target.value as number);
  };

// create deposit
const CreateDeposit = async () => {
  const deposits = {
    added: added + ":00+07:00",
    info: info,
    statusd: statusd,
    employee: employee,
  };

   setStatus(true);
   const timer = setTimeout(() => {
     setStatus(false);
   }, 3000);
  console.log(deposits);
  const res: any = await http.createDeposit({ deposit: deposits });
  
  if (res.id != '') {
    setAlert(true);
  } 
};

    return (
    <Page theme={pageTheme.tool}>

      <Header
        title={`Create deposit invoice`}
        type="Deposit systems"> 
      </Header>

      <Content>
        <ContentHeader title="Deposit invoice"> 
              <Button onClick={() => {CreateDeposit();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new deposit invoice </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/DepositTable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              //fullWidth
              //className={classes.margin}
              variant="outlined"
            >

			  <div className={classes.paper}><strong>Employee</strong>
                <FormControl className={classes.root} noValidate>
                <TextField
                  id="employee"
                  label="employee"
                  type="string"
				  style={{ width: 400 }}
                  value={employees.filter((filter: EntEmployee) => filter.id == employee).map((item: EntEmployee) => `${item.employeename}`)}
                />
               </FormControl>
              </div>

              <div className={classes.paper}><strong>Status Deposit</strong></div>
			  <Select
                  name="Statusd"
				  id="statusd"
                  value={statusd} 
                  onChange={StatusdhandleChange}
                >
                {statusds.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.statusdname}</MenuItem>
                    );
                  })}
              </Select>

			  <div className={classes.paper}><strong>Date</strong></div>
			  <form className={classes.container} noValidate>
                <TextField
                  
                  name="added"
				  id="added"
                  type="datetime-local"
                  value={added} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={AddedhandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>Info</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="info"
				  id="info"
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={info} 
                  onChange={InfohandleChange}
                />
              </form>

			  {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      <Alert severity="success"> <AlertTitle>Success</AlertTitle> Successfull Save </Alert>) 
              : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Unsuccessfull Save!! </Alert>)} </div>
            ) : null}
		
            </FormControl>

          </form>
        </div>
      </Content>
    </Page>
  );
 }
