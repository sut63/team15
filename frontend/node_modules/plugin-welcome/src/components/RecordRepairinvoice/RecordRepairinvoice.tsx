import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  HeaderLabel,
  InfoCard,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';


import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntRentalstatus } from '../../api';
import Repairinvoice from '.';
import { EntRepairinvoice } from '../../api/models/EntRepairinvoice';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
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
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    center: {
      marginTop: theme.spacing(2),
      marginLeft: theme.spacing(23),
    }
  }),
);


export default function CreateRepairinvoice() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [repairinvoices, setRepairinvoice] = React.useState<EntRepairinvoice[]>([]);

  const [status, setStatus] = useState(false);
  const [status2, setStatus2] = useState(false);
 const [alert, setAlert] = useState(false);
  //เก็บข้อมูลที่จะดึงมา
  const [rentalstatuss, setQuantitys] = useState<EntRentalstatus[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  const [bequipment, setBequipment] = useState(String);
  const [rentalstatus, setRentalstatus] = useState(Number);
  const [employeeid, setEmployee] = useState(Number);

  

  useEffect(() => {

    const getRentalstatuss = async () => {
 
      const rs = await api.listRentalstatus({ offset: 0 });
      setLoading(false);
      setQuantitys(rs);
    };
    getRentalstatuss();

        const getEmployees = async () => {
 
          const em = await api.listEmployee();
            setLoading(false);
            setEmployees(em);
          };
          getEmployees();
          const checkEmployee = async () => {
            const edata = JSON.parse(String(localStorage.getItem("employeedata")));
            setLoading(false);
            setEmployee(edata)
          console.log(edata);
          };
          checkEmployee();
        
        }, [loading]);
      console.log(employeeid)
      console.log("testing")
  
      const handleBequipmentChange = (event: any) => {
        setBequipment(event.target.value as string);
      };
  const RentalstatushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRentalstatus(event.target.value as number);
  };
  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const CreateRepairinvoice = async () => {
    const repairinvoices = {
      bequipment: bequipment,
      rentalstatus: rentalstatus,
      employee: employeeid,
    };
    const timer2 = setTimeout(() => {
      setStatus2(true);
   }, 2000);
   const timer3 = setTimeout(() => {
      setStatus2(false);
   }, 6000);
    console.log(repairinvoices);
    const res: any = await api.createRepairinvoice({ repairinvoice : repairinvoices });
    if (res.id != '') {
      setAlert(true);
    setStatus2(false);
    }
    setStatus(true);
    const timer = setTimeout(() => {
       setStatus(false);
       //window.location.reload(false);
    }, 3000);
    
    console.log(repairinvoices);
  };

  
  return (
 <Page theme={pageTheme.service}>
      <Content>
        <InfoCard title="Add Repair in voice" subheader="เพิ่มซ่อมอุปกรณ์ที่ชำรุดสู่ระบบ">
        
        {status ? (
             <div>
                 {alert ? (
                      <Alert severity="success" onClose={() => { }} >บันทึกสำเร็จ</Alert>
                            ) : (
                                <Alert severity="warning" onClose={() => { setStatus(false) }} style={{marginTop: 20}}>กรุณากรอกข้อมูลให้ครบถ้วน</Alert>
                            )}
                        </div>
                    ) : null}
          
          <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              variant="outlined"
            >
               <div className={classes.paper}><strong>อุปกรณ์ที่ชำรุด(Broken Equipment)</strong></div>
              <TextField className={classes.textField}
                id="bequipment"
                label=""
                variant="outlined"
                //color="secondary"
                type="string"
                size="medium"
                value={bequipment}
                onChange={handleBequipmentChange}
              />

              <div className={classes.paper}><strong>สภานะการเข้าพัก(Rental Status)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                labelId="rentalstatus-label"
                id="rentalstatus"
                value={rentalstatus}
                onChange={RentalstatushandleChange}
              >
                <InputLabel className={classes.insideLabel} id="faculty-label">Rental Status</InputLabel>

                {rentalstatuss.map((item: EntRentalstatus) => (
                  <MenuItem value={item.id}>{item.rentalstatus}</MenuItem>
                ))}
              </Select>

              

                <div className={classes.paper}><strong>พนักงานที่ทำการบันทึก(Employee)</strong></div>
                <TextField className={classes.select}
                    id="employee"
                    size="medium"
                    value={employees.filter((filter:EntEmployee) => filter.id == employeeid).map((item:EntEmployee) => `${item.name}`)}
                    style={{ width: 500 }}/>
              <div className={classes.center}>


              <Button
                style={{ width: 150, backgroundColor: "#1814E5",marginLeft: 5 }}
                onClick={() => {
                  CreateRepairinvoice();
                }}
                variant="contained"
                color="primary"
              >
                บันทึกข้อมูลอุปกรณ์ที่ชำรุด
             </Button>


              <Button
                   style={{ width: 150, backgroundColor: "#C1FF3C",marginLeft: 20}}
                component={RouterLink}
                to="/RepairinvoiceTable"
                variant="contained"
              >
                ดูข้อมูลที่บันทึก
             </Button>
             {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      <strong> Successfull Save </strong>) 
              : null} </div>
            ) : null}
			{status2 ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      null) 
              : (<strong> Unsuccessfull Save!! </strong>)} </div>
            ) : null}
            </div>
            </FormControl>
            <div>
			<ComponanceRepairinvoiceTable></ComponanceRepairinvoiceTable>
			</div>
          </form>
        </div>
        </InfoCard>
      </Content>
    </Page>
  );
}

