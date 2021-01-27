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
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';
import HomeOutlinedIcon from '@material-ui/icons/HomeOutlined';
import PermIdentityIcon from '@material-ui/icons/PermIdentity';
import MeetingRoomIcon from '@material-ui/icons/MeetingRoom';
import { Box, Chip, Grid, Link, Typography } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import SaveIcon from '@material-ui/icons/Save';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import { EntStaytype, EntBedtype, EntPetrule, EntPledge } from '../../api';
import RoomDetails from '.';

import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import ComponanceRoomdetailsTable from '../RoomdetailsTable';
import { EntCleanerName } from '../../api/models/EntCleanerName';
import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import { EntLengthTime } from '../../api/models/EntLengthTime'
import Cleaningroom from '.';
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
  }),
);

export default function CreateRoomdetail() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);

  //เก็บข้อมูลที่จะดึงมา
  const [roomdetails, setRoomdetail] = useState<EntRoomdetail[]>([]);
  const [cleanernames, setCleanername] = useState<EntCleanerName[]>([]);
  const [lengthtimes, setLengthtime] = useState<EntLengthTime[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  //เอาไว้เก็บค่าจาก user
  const [roomdetail, setRoomdetails] = useState(Number);
  const [cleanername, setCleanernames] = useState(Number);
  const [lengthtime, setLengthtimes] = useState(Number);
  const [employeeid, setEmployee] = useState(Number);
  const [note, setNote] = useState(String);
  const [timedate, setTime] = useState(String);



  useEffect(() => {

    //ดึงDetailRoom
    const getRoomdetail = async () => {
      const res = await api.listRoomdetail()
      setLoading(false);
      setRoomdetail(res);
    }
    getRoomdetail();

    //ดึงชื่อคนทำความสะอาด
    const getCleanername = async () => {
      const res = await api.listCleanername({ limit: 15, offset: 0 })
      setLoading(false);
      setCleanername(res);
    }
    getCleanername();

    //ดึงช่วงเวลา
    const getLegnthtime = async () => {
      const res = await api.listLengthtime({ limit: 15, offset: 0 })
      setLoading(false);
      setLengthtime(res);
    }
    getLegnthtime();

    //ดึง Employee
    const getEmployees = async () => {
      const em = await api.listEmployee();
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();


    {
      const checkEmployeeLoginData = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
        setLoading(false);
        if (jobdata != "พนักงานหอพัก") {
          localStorage.setItem("employeedata", JSON.stringify(null));
          localStorage.setItem("jobpositiondata", JSON.stringify(null));
          history.pushState("", "", "./");
          window.location.reload(false);
        }
        else {
          setEmployee(Number(localStorage.getItem("employeedata")))
        }
      }
      checkEmployeeLoginData();
    }
  }, [loading]);
  // console.log(employeeid)
  // console.log("testing")

  const roomdetailhandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomdetails(event.target.value as number)
  };

  const cleanernamehandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCleanernames(event.target.value as number)
  };

  const lengthtimehandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLengthtimes(event.target.value as number)
  };


  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const noteehandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNote(event.target.value as string)
  };

  const datetimehandChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTime(event.target.value as string)
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

  const CreateCleaningroom = async () => {
    if (timedate === "" || note === "" || employeeid === 0 || lengthtime === 0 || cleanername === 0 || roomdetail === 0) {
      Toast.fire({
        icon: 'warning',
        title: 'โปรดระบุข้อมูลให้ครบ',
      });
    } else {
      const cleaningroom = {
        dateandstarttime: timedate,
        note: note,
        cleanername: cleanername,
        lengthtime: lengthtime,
        roomdetail: roomdetail,
        employee: employeeid,
      }

      console.log(cleaningroom);
      Toast.fire({
        icon: 'success',
        title: 'บันทึกข้อมูลเส็จ',
      });
      const res: any = await api.createCleaningroom({ cleaningroom: cleaningroom });
      setStatus(true);
      
    }    
  };


  return (
    <Page theme={pageTheme.service}>
      <Content>
        <InfoCard title="แจ้งทำความสะอาด">
          <div className={classes.root}>
            <form noValidate autoComplete="off">
              <Grid container>
                <FormControl>
                  <div className={classes.paper}><strong>พนักงานที่ทำการบันทึก(Employee)</strong></div>
                  <TextField className={classes.select}
                    id="employee"
                    variant="outlined"
                    size="medium"
                    value={employees.filter((filter: EntEmployee) => filter.id == employeeid).map((item: EntEmployee) => `${item.name}`)}
                    onChange={EmployeehandleChange}
                    style={{ width: 400 }} />

                  <div className={classes.paper}><strong>พนักงานทำความสะอาด</strong></div>
                  <Select className={classes.select}
                    //color="secondary"
                    labelId="cleanername-label"
                    id="cleanername"
                    value={cleanername}
                    variant="outlined"
                    onChange={cleanernamehandChange}
                  >
                    {cleanernames.map((item: EntCleanerName) => (
                      <MenuItem value={item.id}>{item.cleanername}</MenuItem>
                    ))}
                  </Select>

                  <div className={classes.paper}><strong>ห้องพัก</strong></div>
                  <Select className={classes.select}
                    //color="secondary"
                    labelId="roomdetail-label"
                    id="roomdetail"
                    value={roomdetail}
                    variant="outlined"
                    onChange={roomdetailhandChange}
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
                    {lengthtimes.map((item: EntLengthTime) => (
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

                  <div className={classes.paper}><strong>Note</strong></div>
                  <TextField className={classes.select}
                    id="note"
                    variant="outlined"
                    size="medium"
                    value={note}
                    onChange={noteehandChange}
                    style={{ width: 400 }} />
                  <br />
                  <Button
                    variant="contained"
                    color="primary"
                    size="large"
                    className={classes.paper}
                    startIcon={<SaveIcon />}
                    onClick={CreateCleaningroom}
                  >Save</Button>

                </FormControl>
              </Grid>
            </form>
          </div>
        </InfoCard>

      </Content>
    </Page>
  );
}
