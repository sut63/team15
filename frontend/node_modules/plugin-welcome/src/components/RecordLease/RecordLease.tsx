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
 const [errormessege, setErrorMessege] = useState(String);
 const [loading, setLoading] = useState(true);

 const [employee, setEmployee] = useState(Number);
 const [tenant, setTenant] = useState(String);
 const [numbtenant, setNumbtenant] = useState(String);
 const [pettenant, setPettenant] = useState(String);
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
	if(edata == ''){
    history.pushState("","","./");
	   window.location.href = "http://localhost:3000";
	}
    setEmployee(edata)
	console.log(edata);
  };
  checkEmployee();

}, [loading]);

const getLease = async () => {
  const res = await http.listLease({ offset: 0 });
  setLease(res);
};

  const LeasehandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>) => {
    setTenant(event.target.value as string);
  };

  const PettenanthandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>) => {
    setPettenant(event.target.value as string);
  };

  const NumbtenanthandleChange = (
    event: React.ChangeEvent<{ name: string; value: string }>) => {
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
    switch(field) {
      case 'tenant':
        setErrorMessege("โปรดกรอกชื่อผู้ทำสัญญาเช่า");
        return;
      case 'numbtenant':
        setErrorMessege("โปรดกรอกชื่อผู้ร่วมพักอาศัย");
        return;
      case 'pettenant':
        setErrorMessege("โปรดกรอกชนิดของสัตว์เลี้ยง");
        return;
      default:
        setErrorMessege("บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  };
  


// create lease
const CreateLease = async () => {
  const leases = {
    added: added + ":00+07:00",
    tenant: tenant,
    numbtenant: numbtenant,
    pettenant: pettenant,
    wifi: wifi,
    roomdetail: roomdetail,
    employee: employee,
  };
  console.log(leases);

  if (tenant == ""){
    checkCaseSaveError("tenant");
  }
  else {
    if(numbtenant == ""){
      checkCaseSaveError("numbtenant");
    }
  else {
    if(pettenant == ""){
        checkCaseSaveError("pettenant");
    }
    else {setErrorMessege("บันทึกข้อมูลไม่สำเร็จ");}
  }
}

  const timer2 = setTimeout(() => {
     setStatus2(true);
  }, 2000);
  const timer3 = setTimeout(() => {
     setStatus2(false);
  }, 6000);
  const res: any = await http.createLease({ lease: leases });
  if (res.id != '') {
    setAlert(true);
	setStatus2(false);
  }
  setStatus(true);
  const timer = setTimeout(() => {
     setStatus(false);
  }, 3000);
  
  console.log(leases);
};
 

    return (
    <Page theme={pageTheme.tool}>


      <Content>
        <ContentHeader title="Contract"> 
              <Button onClick={() => {CreateLease();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new Contract </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/LeaseTable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
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
        <div className={classes.paper}><strong>Parter's Name</strong></div>
        <form className={classes.root} noValidate>
                <TextField
                  name="numbtenant"
				          id="numbtenant"
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
          <div className={classes.paper}><strong>Pet Type</strong></div>
          <form className={classes.root} noValidate>
                <TextField
                  name="pettenant"
				          id="pettenant"
				          style={{ margin: 8 }}
                  type="text"
				          fullWidth
				          margin="normal"
				          InputLabelProps={{
					        shrink: true,
				         }}
                  value={pettenant} 
                  onChange={PettenanthandleChange}
                />
          </form>

              {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
			  <Alert severity="success" style={{ marginTop: 20, marginLeft:5 }} >Successful Save</Alert>
                      ) 
              : null} </div>
            ) : null}
			{status2 ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      null) 
              : (<Alert severity="warning" style={{ marginTop: 20, marginLeft:5 }}> {errormessege} </Alert>)} </div>
            ) : null}
		
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
