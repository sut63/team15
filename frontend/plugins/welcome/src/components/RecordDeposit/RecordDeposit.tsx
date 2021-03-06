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
import SearchIcon from '@material-ui/icons/Search';

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
import { EntLease } from '../../api/models/EntLease'; // import interface Lease
import ComponanceDepositTable from '../DepositTable';

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
 const [leases, setLeases] = React.useState<EntLease[]>([]);

 const [status, setStatus] = useState(false);
 const [status2, setStatus2] = useState(false);
 const [alert, setAlert] = useState(false);
 const [alert2, setAlerts] = useState(true);
 const [alerttype, setAlertType] = useState(String);
 const [errormessege, setErrorMessege] = useState(String);

 const [loading, setLoading] = useState(true);

 const [infoerror, setInfoerror] = React.useState('');
 const [depositornameerror, setDepositornameerror] = React.useState('');
 const [depositortellerror, setDepositortellerror] = React.useState('');
 const [recipienttellerror, setRecipienttellerror] = React.useState('');
 const [parcelcodeerror, setParcelcodeerror] = React.useState('');

 const [employee, setEmployee] = useState(Number);
 const [info, setInfo] = useState(String);
 const [depositorname, setDepositorname] = useState(String);
 const [depositortell, setDepositortell] = useState(String);
 const [recipienttell, setRecipienttell] = useState(String);
 const [parcelcode, setParcelcode] = useState(String);
 const [added, setAdded] = useState(String);
 const [statusd, setStatusd] = useState(Number);
 const [lease, setLease] = useState(Number);

 useEffect(() => {
  const getLeases = async () => {
    const res = await http.listLease({ offset: 0 });
    setLoading(false);
    setLeases(res);
    console.log(res);
  };
  getLeases();

  const getEmployees = async () => {
    const res = await http.listEmployee();
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
    const edata = JSON.parse(String(localStorage.getItem("employeedata")));
    setLoading(false);
    setEmployee(edata)
	console.log(edata);
  };
  checkEmployee();

}, [loading]);

  const validateParcelcode = (val: string) => {
    return val.match("^([A-Z]{3})-([0-9]{4})$");
  }

  const validateDepositortell = (val: string) => {
    return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
  }

  const validateRecipienttell = (val: string) => {
    return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
  }

  const checkPattern  = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'depositortell':
        validateDepositortell(value) ? setDepositortellerror('') : setDepositortellerror ('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
      return;
      case 'recipienttell':
        validateRecipienttell(value) ? setRecipienttellerror('') : setRecipienttellerror ('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
      return;
      case 'parcelcode':
        validateParcelcode(value) ? setParcelcodeerror('') : setParcelcodeerror ('ใส่รูปแบบรหัสพัสดุถูกต้อง XXX-xxxx');
      return;
        default:
          return;
    }
  }


const getDeposit = async () => {
  const res = await http.listDeposit({ offset: 0 });
  setDeposit(res);
};

  const InfohandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>,) => {
    setInfo(event.target.value as string);
  };

  const DepositornamehandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>,) => {
    setDepositorname(event.target.value as string);
  };

  const DepositortellhandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof depositortell;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setDepositortell(event.target.value as string);
  };

  const RecipienttellhandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof recipienttell;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setRecipienttell(event.target.value as string);
  };

  const ParcelcodehandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof parcelcode;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setParcelcode(event.target.value as string);
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

  const LeasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLease(event.target.value as number);
  };

  const checkCaseSaveError = (field: string) => {
    if (field == "depositortell") { setErrorMessege("ข้อมูล field เบอร์โทรศัพท์ของผู้ส่งผิด"); }
        else if (field == "recipienttell") { setErrorMessege("ข้อมูล field เบอร์โทรศัพท์ของผู้รับผิด"); }
        else if (field == "parcelcode") { setErrorMessege("ข้อมูลfield เลขรหัสสินค้าผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ"); }
  }

const CreateDeposit = async () => {
    if ((lease != null) && (employee != null) && (statusd != null) && (info != null)){
	  const apiUrl = 'http://localhost:8080/api/v1/deposits';
      const deposit = {
        added: added + ":00+07:00",
		info: info,
		depositorname: depositorname,
		depositortell: depositortell,
		recipienttell: recipienttell,
		parcelcode: parcelcode,
		statusd: statusd,
		lease: lease,
		employee: employee,
    };
    console.log(deposit);
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(deposit),
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        setStatus(true);
        if (data.status === true) {
          setErrorMessege("บันทึกข้อมูลสำเร็จ");
          setAlertType("success");
        }
        else {
          checkCaseSaveError(data.error.Name);
          setAlertType("error");
        }
      });
  }
  else{
    setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ");
    setAlertType("error");
    setStatus(true);
  }
       };

    return (
    <Page theme={pageTheme.tool}>

      <Content>
        <ContentHeader title="Deposit invoice"> 
              <Button
                style={{ width: 200, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                component={RouterLink} to="/SearchDeposit"
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                Search Deposit
             </Button>
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
                  value={employees.filter((filter: EntEmployee) => filter.id == employee).map((item: EntEmployee) => `${item.name}`)}
                />
               </FormControl>
              </div>

              <div className={classes.paper}><strong>Name</strong></div>
			  <Select
                  name="Lease"
				  id="lease"
                  value={lease} 
                  onChange={LeasehandleChange}
                >
                {leases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.tenant}</MenuItem>
                    );
                  })}
              </Select>

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

			  <div className={classes.paper}><strong>Depositor Name</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="depositorname"
				  id="depositorname"
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={depositorname} 
                  onChange={DepositornamehandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>Depositor Tell</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="depositortell"
				  id="depositortell"
				  error = {depositortellerror ? true : false}
				  helperText= {depositortellerror}
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={depositortell} 
                  onChange={DepositortellhandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>Recipient Tell</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="recipienttell"
				  id="recipienttell"
				  error = {recipienttellerror ? true : false}
				  helperText= {recipienttellerror}
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={recipienttell} 
                  onChange={RecipienttellhandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>Parcel Code</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="parcelcode"
				  id="parcelcode"
				  error = {parcelcodeerror ? true : false}
				  helperText= {parcelcodeerror}
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={parcelcode} 
                  onChange={ParcelcodehandleChange}
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

			  <div>
					<Button onClick={() => {CreateDeposit();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new deposit invoice </Button>
			  </div>

			  <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20, marginLeft:6 }} >
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>
		
            </FormControl>
			<div>
			<ComponanceDepositTable></ComponanceDepositTable>
			</div>

          </form>
        </div>
      </Content>
    </Page>
  );
 }
