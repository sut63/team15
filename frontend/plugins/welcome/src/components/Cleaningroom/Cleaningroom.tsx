import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Page,
  pageTheme,
  ContentHeader,
  InfoCard,
} from '@backstage/core';
import { Grid } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import SearchIcon from '@material-ui/icons/Search';

import { EntCleaningroom } from '../../api/models/EntCleaningroom';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntCleanername } from '../../api/models/EntCleanername';
import { EntLengthtime } from '../../api/models/EntLengthtime';
import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import ComponanceCleaningroomTable from '../CleaningroomTable';
import Swal from 'sweetalert2';

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
      width: 500,
      marginLeft: 7,
      marginRight: -7,
    },
    select: {
      width: 400,
      marginLeft: 7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    center: {
      marginTop: theme.spacing(2),
      marginLeft: theme.spacing(23),
    },
    cardtable: {
      marginTop: theme.spacing(2),
    },
    fieldText: {
      width: 200,
      marginLeft: 7,
    },
    fieldLabel: {
      marginLeft: 8,
      marginRight: 20,

    }
  }),
);


export default function CreateCleaningroom() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);


  //เก็บข้อมูลที่จะดึงมา
  const [cleaningroom, setCleaningroom] = useState<EntCleaningroom[]>([]);
  const [cleanernames, setCleanernames] = useState<EntCleanername[]>([]);
  const [lengthtimes, setLengthtimes] = useState<EntLengthtime[]>([]);
  const [roomdetails, setRoomdetails] = useState<EntRoomdetail[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);


  const [numofemerror, setNumofemerror] = React.useState('');
  const [phonenumbererror, setPhonenumbererror] = React.useState('');
  const [noteerror, setNoteerror] = React.useState('');

  //เอาไว้เก็บค่าจาก user
  const [employeeid, setEmployee] = useState(Number);
  const [cleanername, setCleanername] = useState(Number);
  const [numofem, setNumofem] = useState(Number);
  const [roomdetail, setRoomdetail] = useState(Number);
  const [lengthtime, setLengthtime] = useState(Number);
  const [timedate, setTime] = useState(String);
  const [phonenumber, setPhonenumber] = useState(String);
  const [note, setNote] = useState(String);

  useEffect(() => {

    const getCleaningrooms = async () => {
      const cr = await api.listCleaningroom();
      setLoading(false);
      setCleaningroom(cr);
    };
    getCleaningrooms();

    //ดึงDetailRoom
    const getRoomdetails = async () => {
      const rd = await api.listRoomdetail()
      setLoading(false);
      setRoomdetails(rd);
    }
    getRoomdetails();

    //ดึงชื่อคนทำความสะอาด
    const getCleanernames = async () => {
      const cn = await api.listCleanername({ limit: 15, offset: 0 })
      setLoading(false);
      setCleanernames(cn);
    }
    getCleanernames();

    //ดึงช่วงเวลา
    const getLegnthtimes = async () => {
      const lt = await api.listLengthtime({ limit: 15, offset: 0 })
      setLoading(false);
      setLengthtimes(lt);
    }
    getLegnthtimes();

    //ดึง Employee
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



  const validateNumofem = (val: number) => {
    return val <= 4 && val >= 1 ? true : false
  }

  const validatePhonenumber = (val: string) => {
    return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
  }

  const ValidateNote = (val: string) => {
    return val.match("^[ก-๏\s]+$");
  }


  const checkPattern = (id: string, value: string) => {
    console.log(value);
    switch (id) {
      case 'numofem':
        validateNumofem(Number(value)) ? setNumofemerror('') : setNumofemerror('จำนวนพนักงานทำความสะอาดต้องไม่เกิน 1-4 คน');
        return;
      case 'phonenumber':
        validatePhonenumber(value) ? setPhonenumbererror('') : setPhonenumbererror('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
        return;
      case 'note':
        ValidateNote(value) ? setNoteerror('') : setNoteerror("กรุณากรอกข้อมูลเฉพาะภาษาไทย");
        return;
      default:
        return;
    }
  }

  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const cleanernamehandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCleanername(event.target.value as number)
  };

  const numofemhandChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('numofem', validateValue)
    setNumofem(event.target.value as number)
  };

  const roomdetailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomdetail(event.target.value as number);
  };

  const lengthtimehandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLengthtime(event.target.value as number)
  };

  const datetimehandChange = (event: any) => {
    setTime(event.target.value as string)
  };

  const phonenumberhandChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as typeof phonenumber;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setPhonenumber(event.target.value as string)
  };

  const notehandChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as typeof note;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setNote(event.target.value as string)
  };

  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: false,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const listCleaning = () => {
    setStatus(false);
    if (errormessege == "บันทึกข้อมูลสำเร็จ") {
      window.location.href = "http://localhost:3000/DormEmployee";
    }
  };

  const checkCaseSaveError = (field: string) => {
    if (field == "numofem") { setErrorMessege("ข้อมูล field จำนวนพนักงานทำความสะอาดผิด"); }
    else if (field == "phonenumber") { setErrorMessege("ข้อมูล field เบอร์โทร์ศัพท์ผิด"); }
    else if (field == "note") { setErrorMessege("กรุณากรอกข้อมูลเฉพาะภาษาไทย"); }
    else { setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ"); }
  }


  const CreateCleaningroom = async () => {
    if ((cleanername != null) && (roomdetail != null) && (lengthtime != null)) {
      const apiUrl = 'http://localhost:8080/api/v1/cleaningrooms';
      const cleaningroom = {
        employee: employeeid,
        cleanername: cleanername,
        numofem: Number(numofem),
        roomdetail: roomdetail,
        lengthtime: lengthtime,
        dateandstarttime: timedate + ":00+07:00",
        phonenumber: phonenumber,
        note: note,
      };

      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(cleaningroom),
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
    else {
      setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ");
      setAlertType("error");
      setStatus(true);
    }
  };


  return (
    <Page theme={pageTheme.service}>
      <Content>
        <InfoCard title="Add Cleaning Request">
          <div className={classes.root}>
            <form noValidate autoComplete="off">
              <Grid container>
                <FormControl>

                  <div className={classes.paper}><strong>พนักงานที่ทำการบันทึก(Employee)</strong></div>
                  <TextField className={classes.select}
                    id="employee"
                    label="employee"
                    type="string"
                    value={employees.filter((filter: EntEmployee) => filter.id == employeeid).map((item: EntEmployee) => `${item.name}`)}
                    style={{ width: 400 }} />

                  <div className={classes.paper}><strong>พนักงานทำความสะอาดที่รับผิดชอบ</strong></div>
                  <Select className={classes.select}
                    //color="secondary"
                    labelId="cleanername-label"
                    id="cleanername"
                    value={cleanername}
                    variant="outlined"
                    onChange={cleanernamehandChange}
                  >
                    {cleanernames.map((item: EntCleanername) => (
                      <MenuItem value={item.id}>{item.cleanername}</MenuItem>
                    ))}
                  </Select>

                  <div className={classes.paper}><strong>จำนวนพนักงานในการทำความสะอาด</strong></div>
                  <TextField className={classes.select}
                    id="numofem"
                    error={numofemerror ? true : false}
                    variant="outlined"
                    size="medium"
                    type="number"
                    helperText={numofemerror}
                    value={numofem}
                    onChange={numofemhandChange}
                    style={{ width: 400 }}
                  />

                  <div className={classes.paper}><strong>ห้องพัก</strong></div>
                  <Select className={classes.select}
                    //color="secondary"
                    labelId="roomdetail-label"
                    id="roomdetail"
                    value={roomdetail}
                    variant="outlined"
                    onChange={roomdetailhandleChange}
                  >
                    {roomdetails.map((item: EntRoomdetail) => (
                      <MenuItem value={item.id}>{item.roomnumber}</MenuItem>
                    ))}
                  </Select>

                  <div className={classes.paper}><strong>ระยะเวลา</strong></div>
                  <Select className={classes.select}
                    //color="secondary"
                    labelId="lengthtime-label"
                    id="lengthtime"
                    value={lengthtime}
                    variant="outlined"
                    onChange={lengthtimehandChange}
                  >
                    {lengthtimes.map((item: EntLengthtime) => (
                      <MenuItem value={item.id}>{item.lengthtime}</MenuItem>
                    ))}
                  </Select>

                  <div className={classes.paper}><strong>วันที่และเวลาเริ่มทำความสะอาด</strong></div>
                  <TextField
                    id="timedate-local"
                    type="datetime-local"
                    variant="outlined"
                    defaultValue="2017-05-24T10:30"
                    className={classes.textField}
                    InputLabelProps={{
                      shrink: true,
                    }}
                    value={timedate}
                    onChange={datetimehandChange}
                    style={{ width: 400 }}
                  />

                  <div className={classes.paper}><strong>เบอร์โทรศัพท์ที่สามารถติดต่อได้</strong></div>
                  <TextField className={classes.select}
                    id="phonenumber"
                    error={phonenumbererror ? true : false}
                    variant="outlined"
                    size="medium"
                    type="string"
                    helperText={phonenumbererror}
                    value={phonenumber}
                    onChange={phonenumberhandChange}
                    style={{ width: 400 }}
                  />

                  <div className={classes.paper}><strong>Note</strong></div>
                  <TextField className={classes.select}
                    id="note"
                    error={noteerror ? true : false}
                    variant="outlined"
                    size="medium"
                    type="string"
                    helperText={noteerror}
                    value={note}
                    onChange={notehandChange}
                    style={{ width: 400 }}
                  />


                  <br />
                  <div>
                    <Button
                      style={{ width: 200, backgroundColor: "#5319e7", marginLeft: 7 }}
                      onClick={() => {
                        CreateCleaningroom();
                      }}
                      variant="contained"
                      color="primary"
                    >
                      บันทึกข้อมูล
             </Button>


                  </div>

                  <div>
                    {status ? (
                      <div>
                        {alerttype != "" ? (
                          <Alert severity={alerttype} style={{ width: 400, marginTop: 20, marginLeft: 6 }} onClose={() => { listCleaning() }}>
                            {errormessege}
                          </Alert>
                        ) : null}
                      </div>
                    ) : null}</div>

                </FormControl>
              </Grid>
            </form>
          </div>
        </InfoCard>
        <div>
          <InfoCard className={classes.cardtable} title="Cleaningroom table" subheader="ตารางรายละเอียดแจ้งทำความสะอาดห้อง">
            <ComponanceCleaningroomTable></ComponanceCleaningroomTable>
          </InfoCard>
        </div>

      </Content>
    </Page>
  );
}
