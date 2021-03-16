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
import { EntRentalstatus } from '../../api/models/EntRentalstatus'; // import interface Rentalstatus
import { EntRepairinvoice } from '../../api/models/EntRepairinvoice'; // import interface Repairinvoice
import { EntLease } from '../../api/models/EntLease'; // import interface Lease
import ComponanceRepairinvoiceTable from '../RepairinvoiceTable';

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

export default function recordRepairinvoice() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [repairinvoice, setRepairinvoice] = React.useState<Partial<recordRepairinvoice>>({});

 const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
 const [rentalstatuss, setRentalstatuss] = React.useState<[EntRentalstatus]>([]);
 const [leases, setLeases] = React.useState<EntLease[]>([]);

 const [status, setStatus] = useState(false);
 const [status2, setStatus2] = useState(false);
 const [alert, setAlert] = useState(false);
 const [alert2, setAlerts] = useState(true);
 const [alerttype, setAlertType] = useState(String);
 const [errormessege, setErrorMessege] = useState(String);

 const [loading, setLoading] = useState(true);

 const [emtellerror, setEmtellerror] = React.useState('');
 const [numerror, setNumerror] = React.useState('');
 const [bequipmenterror, setBequipmenterror] = React.useState('');

 const [employee, setEmployee] = useState(Number);
 const [emtell, setEmtell] = useState(String);
 const [num, setNum] = useState(Number);
 const [bequipment, setBequipment] = useState(String);
 const [rentalstatus, setRentalstatus] = useState(Number);
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

  const getRentalstatuss = async () => {
    const res = await http.listRentalstatus({ offset: 0 });
    setLoading(false);
    setRentalstatuss(res);
    console.log(res);
  };
  getRentalstatuss();

  const checkEmployee = async () => {
    const edata = JSON.parse(String(localStorage.getItem("employeedata")));
    setLoading(false);
    setEmployee(edata)
	console.log(edata);
  };
  checkEmployee();

}, [loading]);

  const validateBequipment = (val: string) => {
    return val.match("^([A-Z]{3})-([0-9]{4})$");
  }

  const validateEmtell = (val: string) => {
    return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
  }

  const validateNum = (val: number) => {
    return val <= 10 && val >=1 ? true:false
  }

  const checkPattern  = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'emtell':
        validateEmtell(value) ? setEmtellerror('') : setEmtellerror ('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
      return;
      case 'num':
        validateNum(value) ? setNumerror('') : setNumerror ('จำนวนอุปกรณ์ต้องไม่เกิน 1-10 ชิ้น');
      return;
      case 'bequipment':
        validateBequipment(value) ? setBequipmenterror('') : setBequipmenterror ('ใส่รูปแบบรหัสอุปกรณ์ให้ถูกต้อง XXX-xxxx');
      return;
        default:
          return;
    }
  }


const getRepairinvoice = async () => {
  const res = await http.listRepairinvoice({ offset: 0 });
  setRepairinvoice(res);
};

  const EmtellhandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof emtell;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setEmtell(event.target.value as string);
  };

  const NumhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('num', validateValue)
    setNum(event.target.value as number);
  };

  const BequipmenthandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof bequipment;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setBequipment(event.target.value as string);
  };

  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const RentalstatushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRentalstatus(event.target.value as number);
  };

  const LeasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLease(event.target.value as number);
  };

  const checkCaseSaveError = (field: string) => {
    if (field == "emtell") { setErrorMessege("ข้อมูล field เบอร์โทรศัพท์ของพนักงานผิด"); }
        else if (field == "num") { setErrorMessege("ข้อมูล field จำนวนอุปกรณ์ผิด"); }
        else if (field == "bequipment") { setErrorMessege("ข้อมูลfield เลขรหัสอุปกรณ์ผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ"); }
  }

const CreateRepairinvoice = async () => {
    if ((lease != null) && (employee != null) && (rentalstatus != null)){
	  const apiUrl = 'http://localhost:8080/api/v1/repairinvoices';
      const repairinvoice = {
		emtell: emtell,
    num: Number(num),
		bequipment: bequipment,
		rentalstatus: rentalstatus,
		lease: lease,
		employee: employee,
    };
    console.log(repairinvoice);
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(repairinvoice),
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
        <ContentHeader title="Repairinvoice invoice"> 
              <Button
                style={{ width: 200, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                component={RouterLink} to="/SearchRepairinvoice"
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                Search Repairinvoice
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

              <div className={classes.paper}><strong>Rental Status</strong></div>
			  <Select
                  name="Rentalstatus"
				  id="rentalstatus"
                  value={rentalstatus} 
                  onChange={RentalstatushandleChange}
                >
                {rentalstatuss.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.rentalstatus}</MenuItem>
                    );
                  })}
              </Select>

			  <div className={classes.paper}><strong>Empolyee Tell</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="emtell"
				  id="emtell"
				  error = {emtellerror ? true : false}
				  helperText= {emtellerror}
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={emtell} 
                  onChange={EmtellhandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>The number of defective devices</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="num"
				  id="num"
				  error = {numerror ? true : false}
				  helperText= {numerror}
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={num} 
                  onChange={NumhandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>Broken Equipment</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="bequipment"
				  id="bequipment"
				  error = {bequipmenterror ? true : false}
				  helperText= {bequipmenterror}
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={bequipment} 
                  onChange={BequipmenthandleChange}
                />
              </form>

			 

			  <div>
					<Button onClick={() => {CreateRepairinvoice();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new Repairinvoice invoice </Button>
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
			<ComponanceRepairinvoiceTable></ComponanceRepairinvoiceTable>
			</div>

          </form>
        </div>
      </Content>
    </Page>
  );
 }
