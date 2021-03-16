import React, { useState, useEffect } from 'react';
import {
  Content,
  Page,
  pageTheme,
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
import { EntStaytype, EntBedtype, EntPetrule, EntPledge } from '../../api';
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

  //เก็บข้อมูลที่จะดึงมา
  const [roomdetail, setRoomdetail] = useState<EntRoomdetail[]>([]);
  const [bedtypes, setBedtypes] = useState<EntBedtype[]>([]);
  const [staytypes, setStaytypes] = useState<EntStaytype[]>([]);
  const [petrules, setPetrules] = useState<EntPetrule[]>([]);
  const [pledges, setPledges] = useState<EntPledge[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  const [noroomerror, setRoomnumbererror] = React.useState('');
  const [roomnameerror, setRoomnameerror] = React.useState('');
  const [phoneerror, setPhoneerror] = React.useState('');
  const [sleeperror, setSleeperror] = React.useState('');
  const [bederror, setBederror] = React.useState('');
  const [priceerror, setRoompriceerror] = React.useState('');
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);

  const [price, setRoomprice] = useState(Number);
  const [noroom, setRoomnumber] = useState(String);
  const [roomname, setRoomtypename] = useState(String);
  const [phone, setPhone] = useState(String);
  const [sleep, setSleep] = useState(Number);
  const [bed, setBed] = useState(Number);

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

      
  const validateBed = (val: number) => {
    return val <= 4 && val >=1 ? true:false
  }

  const validateSleep = (val: number) => {
    return val <= 10 && val >=1 ? true:false
  }

  const validateNoroom = (val: string) => {
    return val.match("^[ABD]{1}[0-9]{3}$");
  }

  const validateRoomname = (val: string) => {
    return val.match("^[A-Z]{1}[0-9]{3}$");
  }

  const validatePhone = (val: string) => {
    return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
  }

  const validatePrice = (val: number) => {
    return val > 0 ? true:false
  }



  const checkPattern  = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'bed':
        validateBed(Number(value)) ? setBederror('') : setBederror('จำนวนเตียงต่อห้องต้องไม่เกิน 1-4 เตียง');
        return;
      case 'sleep':
        validateSleep(Number(value)) ? setSleeperror('') : setSleeperror('จำนวนผู้เข้าพักต่อห้องต้องไม่เกิน 1-10 คน');
      return;
      case 'roomnumber':
        validateNoroom(value) ? setRoomnumbererror('') : setRoomnumbererror('เลขห้องต้องขึ้นต้นด้วยตัวอักษร A,B,D และตามด้วยเลขจำนวน 3 หลักเท่านั้น');
      return;
      case 'roomname':
        validateRoomname(value) ? setRoomnameerror('') : setRoomnameerror ('');
      return;
      case 'phone':
        validatePhone(value) ? setPhoneerror('') : setPhoneerror ('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
      return;
      case 'roomprice':
        validatePrice(Number(value)) ? setRoompriceerror('') : setRoompriceerror ('ใส่รูปแบบราคาไม่ถูกต้อง ราคาต้องไม่เท่ากับ 0 หรือติดลบ');
      return;
        default:
          return;
    }
  }


  const handleRoompriceChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('roomprice', validateValue)
    setRoomprice(event.target.value as number);
  };

  const handleNoroomChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof noroom;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setRoomnumber(event.target.value as string);
  };
  
  
  const handleRoomtypenameChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof roomname;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setRoomtypename(event.target.value as string);
  };

  const handlePhoneChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof phone;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setPhone(event.target.value as string);
  };

  const handleSleepChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('sleep', validateValue)
    setSleep(event.target.value as number);
  };
  
  const handleBedChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('bed', validateValue)
    setBed(event.target.value as number);
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
    setStatus(false);
    if(errormessege == "บันทึกข้อมูลสำเร็จ"){
      window.location.href ="http://localhost:3000/DormEmployee"
    }
  };

  const clearData = () => {
    setRoomprice(0);
    setRoomnumber('');
    setRoomtypename('');
    setPhone('');
    setBed(0);
    setSleep(0);
    setBedtype(0);
    setStaytype(0);
    setPledge(0);
    setPetrule(0);
  }

  
  const forCheck = () => {
    var i = 0;
    var check = 0;
    for (const rdt of roomdetail){
      console.log(rdt);
     i++;
    }
    console.log(i);
    if(i === 0){
      CreateRoomdetail();
   }
   else{
    for (const rdt of roomdetail){
      if(noroom === rdt.roomnumber){
        //console.log("ซ้ำ");
        setAlertType("error");
        setErrorMessege("มีข้อมูลห้องในระบบแล้ว")
        clearData();
        setStatus(true);
        check = 1;
        break;
        //window.location.reload(false);
     }
 }
 if(check != 1){
  CreateRoomdetail();
}
   }
  }

  const checkCaseSaveError = (field: string) => {
    if (field == "bed") { setErrorMessege("ข้อมูลfield จำนวนเตียงผิด"); }
        else if (field == "sleep") { setErrorMessege("ข้อมูลfield จำนวนผู้เข้าพักผิด"); }
        else if (field == "roomnumber") { setErrorMessege("ข้อมูลfield เลขห้องผิด"); }
        else if (field == "roomtypename") { setErrorMessege("ข้อมูลfield ชื่อประเภทห้องพักผิด"); }
        else if (field == "roomprice") { setErrorMessege("ข้อมูลfield ราคาห้องผิด"); }
        else if (field == "phone") { setErrorMessege("ข้อมูลfield เบอร์โทรศัพท์ผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ"); }
  }

  const CreateRoomdetail = async () => {
    if ((bedtype != null) && (petrule != null) && (pledge != null) && (staytype != null)){
      const apiUrl = 'http://localhost:8080/api/v1/roomdetails';
      const roomdetail = {
      roomnumber: noroom,
      roomprice: Number(price),
      roomtypename: roomname,
      phone: phone,
      sleep: Number(sleep),
      bed: Number(bed),
      bedtype: bedtype,
      petrule: petrule,
      pledge: pledge,
      staytype: staytype,
      employee: employeeid,
    };
   
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(roomdetail),
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
                error = {noroomerror ? true : false}
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                helperText= {noroomerror}
                value={noroom}
                onChange={handleNoroomChange}
              /> 
              </FormControl>

                 <strong className={classes.fieldLabel}>ชื่อประเภทห้อง(Room type)</strong>
              <FormControl>
              <TextField 
                style={{ width: 250}}
                id="roomtypename"
                error = {roomnameerror ? true : false}
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                helperText= {roomnameerror}
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
                error = {priceerror ? true : false}
                label=""
                variant="standard"
                //color="secondary"
                //type="number"
                size="medium"
                helperText= {priceerror}
                value={price}
                onChange={handleRoompriceChange}
              />
              </FormControl>
                <strong className={classes.fieldLabel}>จำนวนผู้พักได้สูงสุด(Sleeps)</strong>
              <FormControl>
              <TextField 
                style={{ width: 50}}
                id="sleep"
                error = {sleeperror ? true : false}
                label=""
                variant="standard"
                //color="secondary"
                type="number"
                size="medium"
                helperText= {sleeperror}
                value={sleep}
                onChange={handleSleepChange}
              />
              </FormControl>

              <strong className={classes.fieldLabel}>จำนวนเตียง(Beds)</strong>
              <FormControl>
              <TextField 
                style={{ width: 50}}
                id="bed"
                error = {bederror ? true : false}
                label=""
                variant="standard"
                //color="secondary"
                type="number"
                size="medium"
                helperText= {bederror}
                value={bed}
                onChange={handleBedChange}
              />
              </FormControl>

              <strong className={classes.fieldLabel}>เบอร์โทรศัพท์ที่สามารถติดต่อได้(Phone)</strong>
              <FormControl>
              <TextField 
                style={{ width: 110}}
                id="phone"
                error = {phoneerror ? true : false}
                label=""
                variant="standard"
                //color="secondary"
                type="string"
                size="medium"
                helperText= {phoneerror}
                value={phone}
                onChange={handlePhoneChange}
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
                style={{ width: 130, backgroundColor: "#5319e7",marginTop: 30,marginLeft: 7}}
                onClick={() => {
                  forCheck();
                }}
                variant="contained"
                color="primary"
              >
                บันทึกข้อมูลห้อง
             </Button>

             <Button
                style={{ width: 130, backgroundColor: "#C70039",marginTop: 30,marginLeft: 7}}
                onClick={() => { clearData(); }}
                //defaultValue="Reset"
                type="reset"
                variant="contained"
                color="primary"
              >
                ล้างข้อมูลที่กรอก
             </Button>

          

              { /*<Button
                   style={{ width: 150, backgroundColor: "#C1FF3C",marginLeft: 20}}
                component={RouterLink}
                to="/ComponanceRoomdetailsTable"
                variant="contained"
              >
                ดูข้อมูลที่บันทึก
              </Button> */}
              
      </InfoCard>
      <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20 }} onClose={() => { listRoom() }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>
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
