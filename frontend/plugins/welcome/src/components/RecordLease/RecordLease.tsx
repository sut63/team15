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
import { EntWifi } from '../../api/models/EntWifi'; // import interface Wifi
import { EntRoomdetail } from '../../api/models/EntRoomdetail'; // import interface RoomEntRoomdetail
import { EntLease } from '../../api/models/EntLease'; // import interface Lease
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface EmEntEmployee
import ComponanceLeaseTable from '../LeaseTable';
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

export default function RecordLease() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [lease, setLease] = React.useState<EntLease[]>([]);
 const [employees, setEmployees] = React.useState<EntEmployee[]>([]); 
 const [wifis, setWifis] = React.useState<EntWifi[]>([]);
 const [roomdetails, setRoomdetails] = React.useState<EntRoomdetail[]>([]);

 const [status, setStatus] = useState(false);
 const [status2, setStatus2] = useState(false);
 const [alert, setAlert] = useState(false);
 const [alert2, setAlerts] = useState(true);
 const [alerttype, setAlertType] = useState(String);
 const [errormessege, setErrorMessege] = useState(String);

 const [loading, setLoading] = useState(true);

 const [tenanterror, setTenanterror] = React.useState('');
 const [numbtenanterror, setNumbtenanterror] = React.useState('');
 const [idtenanterror, setIdtenanterror] = React.useState('');
 const [agetenanterror, setAgetenanterror] = React.useState('');

 const [employee, setEmployee] = useState(Number);
 const [tenant, setTenant] = useState(String);
 const [numbtenant, setNumbtenant] = useState(String);
 const [idtenant, setIdtenant] = useState(String);
 const [agetenant, setAgetenant] = useState(Number);
 const [added, setAdded] = useState(String);
 const [wifi, setWifi] = useState(Number);
 const [roomdetail, setRoomdetail] = useState(Number);

 useEffect(() => {
  const getEmployees = async () => {
    const res = await http.listEmployee();
    setLoading(false);
    setEmployees(res);
    console.log(res);
  };
  getEmployees();

  const getWifis = async () => {
    const res = await http.listWifi({ offset: 0 });
    setLoading(false);
    setWifis(res);
    console.log(res);
  };
  getWifis();

  const getRoomdetails = async () => {
    const res = await http.listRoomdetail();
    setLoading(false);
    setRoomdetails(res);
    console.log(res);
  };
  getRoomdetails();

const checkEmployee = async () => {
  const edata = JSON.parse(String(localStorage.getItem("employeedata")));
  setLoading(false);
  setEmployee(edata)
console.log(edata);
};
checkEmployee();

}, [loading]);

const validateNumbtenant = (val: string) => {
  return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
}

const validateIdtenant = (val: string) => {
  return val.match("^([0-9]{1})-([0-9]{4})-([0-9]{5})-([0-9]{2})-([0-9]{1})$");
}

const validateAgetenant = (val: number) => {
  return (val <=200 && val >=1) ? true:false
}

const checkPattern  = (id: string, value:string) => {
  console.log(value);
  switch(id) {
    case 'numbtenant':
      validateNumbtenant(value) ? setNumbtenanterror('') : setNumbtenanterror ('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
    return;
    case 'agetenant':
        validateAgetenant(Number(value)) ? setAgetenanterror('') : setAgetenanterror('อายุต้องมากกว่า 0 แต่ไม่เกิน 200 ปี');
        return;
    case 'idtenant':
      validateIdtenant(value) ? setIdtenanterror('') : setIdtenanterror ('รูปแบบของเลขบัตรประจำตัวประชาชนต้องเป็น x-xxxx-xxxxx-xx-x');
    return;
      default:
        return;
  }
}

const getLease = async () => {
  const res = await http.listLease({ offset: 0 });
  setLease(res);
};

  const LeasehandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>) => {
    setTenant(event.target.value as string);
  };

  const IdtenanthandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof idtenant;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setIdtenant(event.target.value as string);
  };

  const AgetenanthandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('agetenant', validateValue)
    setAgetenant(event.target.value as number);
  };

  const NumbtenanthandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof numbtenant;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setNumbtenant(event.target.value as string);
  };

  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const AddedhandleChange = (event: any) => {
    setAdded(event.target.value as string);
  };

  const WifihandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setWifi(event.target.value as number);
  };

  const RoomdetailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomdetail(event.target.value as number);
  };

  const checkCaseSaveError = (field: string) => {
    if (field == "numbtenant") { setErrorMessege("ข้อมูลเบอร์โทรศัพท์ผิด"); }
        else if (field == "idtenant") { setErrorMessege("ข้อมูลรหัสบัตรประชาชนผิด"); }
        else if (field == "tenant") { setErrorMessege("โปรดกรอกชื่อผู้ทำสัญญาเช่า"); }
        else if (field == "agetenant") { setErrorMessege("ข้อมูลอายุผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ/ใส่ข้อมูลไม่ครบ"); }
  }

  


// create lease
const CreateLease = async () => {
  if ((lease != null) && (employee != null) && (tenant != null)){
	const apiUrl = 'http://localhost:8080/api/v1/leases';
  const lease = {
    added: added + ":00+07:00",
    tenant: tenant,
    numbtenant: numbtenant,
    idtenant: idtenant,
    agetenant: Number(agetenant),
    wifi: wifi,
    roomdetail: roomdetail,
    employee: employee,
  };
  console.log(lease);

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(lease),
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
  setErrorMessege("บันทึกไม่สำเร็จ/ใส่ข้อมูลไม่ครบ");
  setAlertType("error");
  setStatus(true);
}
     };


    return (
    <Page theme={pageTheme.tool}>
      <Content>
        <ContentHeader title="Contract"> 
        <Button
                style={{ width: 200, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                component={RouterLink} to="/SearchLease"
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                Search Lease
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
              <div className={classes.paper}><strong>Room no.</strong></div>
			  <Select
                  name="Roomdetail"
				          id="roomdetail"
                  value={roomdetail} 
                  onChange={RoomdetailhandleChange}
                >
                {roomdetails.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.roomnumber}</MenuItem>
                    );
                  })}
              </Select>

              <div className={classes.paper}><strong>Wifi Status</strong></div>
			  <Select
                  name="Wifi"
				          id="wifi"
                  value={wifi} 
                  onChange={WifihandleChange}
                >
                {wifis.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.wifiname}</MenuItem>
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

			  <div className={classes.paper}><strong>Name</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  name="tenant"
				          id="tenant"
				          style={{ margin: 8 }}
                  type="text"
				          fullWidth
				          margin="normal"
				          InputLabelProps={{
					        shrink: true,
				         }}
                  value={tenant} 
                  onChange={LeasehandleChange}
                />
        </form>
        <div className={classes.paper}><strong>Age</strong></div>
        <form className={classes.root} noValidate>
                <TextField
                  name="agetenant"
				          id="agetenant"
				          error = {agetenanterror ? true : false}
				          helperText= {agetenanterror}
				          style={{ margin: 8 }}
                  type="number"
				          fullWidth
				          margin="normal"
				          InputLabelProps={{
					        shrink: true,
				          }}
                  value={agetenant} 
                  onChange={AgetenanthandleChange}
                />
          </form>
        <div className={classes.paper}><strong>Phone Number</strong></div>
        <form className={classes.root} noValidate>
                <TextField
                  name="numbtenant"
				          id="numbtenant"
				          error = {numbtenanterror ? true : false}
				          helperText= {numbtenanterror}
				          style={{ margin: 8 }}
                  type="text"
				          fullWidth
				          margin="normal"
				          InputLabelProps={{
					        shrink: true,
				          }}
                  value={numbtenant} 
                  onChange={NumbtenanthandleChange}
                />
          </form>
          <div className={classes.paper}><strong>Identification code</strong></div>
          <form className={classes.root} noValidate>
                <TextField
                  name="idtenant"
				          id="idtenant"
				          error = {idtenanterror ? true : false}
				          helperText= {idtenanterror}
				          style={{ margin: 8 }}
                  type="text"
				          fullWidth
				          margin="normal"
				          InputLabelProps={{
					        shrink: true,
				          }}
                  value={idtenant} 
                  onChange={IdtenanthandleChange}
                />
          </form>
          <div>
					<Button onClick={() => {CreateLease();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new Contract </Button>
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
			<ComponanceLeaseTable></ComponanceLeaseTable>
			</div>

          </form>
        </div>
      </Content>
    </Page>
  );
 }         
