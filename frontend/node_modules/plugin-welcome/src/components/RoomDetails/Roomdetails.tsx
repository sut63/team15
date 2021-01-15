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


import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import { EntStaytype, EntBedtype, EntPetrule, EntPledge } from '../../api';
import RoomDetails from '.';
import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import ComponanceRoomdetailsTable from '../RoomdetailsTable';

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
      width: 400 ,
      marginLeft:7,
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
      marginLeft:7,
    },
    fieldLabel: {
      marginLeft:8,
      marginRight: 20,

    }
  }),
);


export default function CreateRoomdetail() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);

  //เก็บข้อมูลที่จะดึงมา
  const [roomdetail, setRoomdetail] = useState<EntRoomdetail[]>([]);
  const [bedtypes, setBedtypes] = useState<EntBedtype[]>([]);
  const [staytypes, setStaytypes] = useState<EntStaytype[]>([]);
  const [petrules, setPetrules] = useState<EntPetrule[]>([]);
  const [pledges, setPledges] = useState<EntPledge[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  const [price, setRoomprice] = useState(String);
  const [noroom, setRoomnumber] = useState(String);
  const [roomname, setRoomtypename] = useState(String);
  const [sleep, setSleep] = useState(String);
  const [bed, setBed] = useState(String);
  const [staytype, setStaytype] = useState(Number);
  const [bedtype, setBedtype] = useState(Number);
  const [petrule, setPetrule] = useState(Number);
  const [pledge, setPledge] = useState(Number);
  const [employeeid, setEmployee] = useState(Number);

  

  useEffect(() => {
    const getRoomdetais = async () => {
      const rd = await api.listRoomdetail();
      setLoading(false);
      setRoomdetail(rd);
    };
    getRoomdetais();
 
    const getBedtypes = async () => {
    const bt = await api.listBedtype();
      setLoading(false);
      setBedtypes(bt);
    };
    getBedtypes();

 
    const getStaytypes = async () => {
    const st = await api.listStaytype({ offset: 0 });
      setLoading(false);
      setStaytypes(st);
    };
    getStaytypes();
 
    const getPetrules = async () => {
     const pr = await api.listPetrule();
       setLoading(false);
       setPetrules(pr);
     };
     getPetrules();

     const getPledges = async () => {
      const p = await api.listPledge();
        setLoading(false);
        setPledges(p);
      };
      getPledges();

        const getEmployees = async () => {
          const em = await api.listEmployee();
            setLoading(false);
            setEmployees(em);
          };
          getEmployees();
        
        
         {const checkEmployeeLoginData = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if(jobdata != "พนักงานหอพัก"){
              localStorage.setItem("employeedata", JSON.stringify(null));
              localStorage.setItem("jobpositiondata", JSON.stringify(null));
              history.pushState("", "", "./");
              window.location.reload(false);    
            }
            else{
              setEmployee(Number(localStorage.getItem("employeedata")))
            }
          }
        checkEmployeeLoginData();} 
      
      }, [loading]);
      console.log(employeeid)
      console.log("testing")


  const handleRoompriceChange = (event: any) => {
    setRoomprice(event.target.value as string);
  };

  const handleNoroomChange = (event: any) => {
    setRoomnumber(event.target.value as string);
  };
  
  
  const handleRoomtypenameChange = (event: any) => {
    setRoomtypename(event.target.value as string);
  };

  const handleSleepChange = (event: any) => {
    setSleep(event.target.value as string);
  };
  
  const handleBedChange = (event: any) => {
    setBed(event.target.value as string);
  };
  
  
  const BedtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBedtype(event.target.value as number);
  };

  const StaytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStaytype(event.target.value as number);
  };

  const PetrulehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPetrule(event.target.value as number);
  };

  const PledgehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPledge(event.target.value as number);
  };

  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const listRoom = () => {
    window.location.href ="http://localhost:3000/RoomDetails";
  };
          const forCheck = () => {
            var i = 0;
            for (const rdt of roomdetail){
            i++;
            }
            if(i === 0){
              CreateRoomdetail();
          }
          else{
            for (const rdt of roomdetail){
              if(noroom === rdt.roomnumber){
                setStatus(true);
                setAlert(false);
                setAlerts(false);
                //window.location.reload(false);
        }
        else{
          CreateRoomdetail();
        }
     }
   }
  }

  const CreateRoomdetail = async () => {
    if ((noroom != "") && (noroom != null) && (roomname != "") && (roomname != null) && (price != "") && (price != null)
    && (sleep != null) && (sleep != null)  && (bed != null) && (bed != null) 
    && (bedtype != null) && (petrule != null) && (pledge != null) && (staytype != null)){
    
      const roomdetail = {
      roomnumber: noroom,
      roomprice: price,
      roomtypename: roomname,
      sleep: sleep,
      bed: bed,
      bedtype: bedtype,
      petrule: petrule,
      pledge: pledge,
      staytype: staytype,
      employee: employeeid,
    };
    console.log(roomdetail);
    const res: any = await api.createRoomdetail({ roomdetail: roomdetail });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      //window.location.reload(false);
    }
  }
     else {
      setAlert(false);
      setStatus(true);
    }
  };


  return (
 <Page theme={pageTheme.service}>
      <Content>
        <InfoCard title="Add room details" subheader="เพิ่มรายละเอียดห้องพักเข้าสู่ระบบ">
        <div className={classes.root}>
          <form noValidate autoComplete="off">


        <Grid container>
          
    <Grid item xs container>
      <Grid item xs={12}>
        <InfoCard title="กรอกข้อมูลห้องพัก">
          <FormControl>

              {/*///////// Field กรอก ///////////*/}
              <div>
                <div>
                  <strong className={classes.fieldLabel}>เลขห้อง(No room)</strong>
              <FormControl>
              <TextField 
                style={{ width: 70, marginBottom: 5}}
                id="roomnumber"
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                value={noroom}
                onChange={handleNoroomChange}
              /> 
              </FormControl>

                 <strong className={classes.fieldLabel}>ชื่อประเภทห้อง(Room type)</strong>
              <FormControl>
              <TextField 
                style={{ width: 250}}
                id="roomtypename"
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                value={roomname}
                onChange={handleRoomtypenameChange}
              /> 
              </FormControl>
                </div>

                <strong className={classes.fieldLabel}>ราคาห้อง(Price)</strong>
              <FormControl>
              <TextField 
                style={{ width: 100}}
                id="roomprice"
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                value={price}
                onChange={handleRoompriceChange}
              />
              </FormControl>
                <strong className={classes.fieldLabel}>จำนวนผู้พักได้สูงสุด(Sleeps)</strong>
              <FormControl>
              <TextField 
                style={{ width: 50}}
                id="sleep"
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                value={sleep}
                onChange={handleSleepChange}
              />
              </FormControl>

              <strong className={classes.fieldLabel}>จำนวนเตียง(Beds)</strong>
              <FormControl>
              <TextField 
                style={{ width: 50}}
                id="bed"
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                value={bed}
                onChange={handleBedChange}
              />
              </FormControl>
              
            </div> 
          </FormControl>
        </InfoCard>
      </Grid>

    </Grid>
    

    <Grid item xs={5}>
      <InfoCard
        title="เลือกข้อมูลห้องพัก"
      >
      {/*///////// Combobox ///////////*/}

      <div className={classes.paper}><strong>ประเภทเตียงนอน(Bed type)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                labelId="bedtype-label"
                id="bedtype"
                value={bedtype}
                variant="outlined"
                onChange={BedtypehandleChange}
              >
                {bedtypes.map((item: EntBedtype) => (
                  <MenuItem value={item.id}>{item.bedtypename}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ประเภทการเข้าพัก(Stay type)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="staytype"
                value={staytype}
                variant="outlined"
                onChange={StaytypehandleChange}
              >
                {staytypes.map((item: EntStaytype) => (
                  <MenuItem value={item.id}>{item.staytype}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ข้อจำกัดการจ่ายมัดจำ</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="pledge"
                value={pledge}
                variant="outlined"
                onChange={PledgehandleChange}
              >
                {pledges.map((item: EntPledge) => (
                  <MenuItem value={item.id}>{item.provision}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ข้อจำกัดในการเลี้ยงสัตว์</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="petrule"
                value={petrule}
                variant="outlined"
                onChange={PetrulehandleChange}
              >
                {petrules.map((item: EntPetrule) => (
                  <MenuItem value={item.id}>{item.petrule}</MenuItem>
                ))}
              </Select>

                <div className={classes.paper}><strong>พนักงานที่ทำการบันทึก(Employee)</strong></div>
                <TextField className={classes.select}
                    id="employee"
                    size="medium"
                    value={employees.filter((filter:EntEmployee) => filter.id == employeeid).map((item:EntEmployee) => `${item.name} (${item.email}) ตำแหน่ง (${item.edges?.jobposition?.positionname})`)}
                    style={{ width: 400 }}/>
             
              <Button
                style={{ width: 200, backgroundColor: "#5319e7",marginTop: 30,marginLeft: 7}}
                onClick={() => {
                  forCheck();
                }}
                variant="contained"
                color="primary"
              >
                บันทึกข้อมูลห้อง
             </Button>

              { /*<Button
                   style={{ width: 150, backgroundColor: "#C1FF3C",marginLeft: 20}}
                component={RouterLink}
                to="/ComponanceRoomdetailsTable"
                variant="contained"
              >
                ดูข้อมูลที่บันทึก
              </Button> */}
              <div>
              {status ? (
                        <div>
                    {(!alert2) ?
                          <Alert severity="warning" style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => {window.location.reload(false)}}>
                          มีข้อมูลนี้อยู่ในระบบแล้ว
                          </Alert>
                      :
                      (alert) ? (
                        <Alert severity="success" style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => {listRoom()}}>
                            บันทึกสำเร็จ
                        </Alert>
                    ) : (
                            <Alert severity="warning" style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => {setStatus(false)}}>
                                บันทึกไม่สำเร็จ กรุณาใส่ข้อมูลให้ครบ
                            </Alert>
                        )
                    }
                        </div>
                      ) : null}</div>
      </InfoCard>
    </Grid>
  </Grid>
          </form>
        </div>
        </InfoCard>
        <div>
          <InfoCard className={classes.cardtable} title="Room details table" subheader="ตารางรายละเอียดข้อมูลห้อง">
        <ComponanceRoomdetailsTable></ComponanceRoomdetailsTable>
        </InfoCard>
        </div>

        
      </Content>
    </Page>
  );
}
