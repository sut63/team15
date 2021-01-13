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
import { EntEquipment, EntFacilitie, EntQuantity, EntNearbyplace, EntStaytype } from '../../api';
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
    },
    cardtable: {
      marginTop: theme.spacing(5),
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
  const [quantitys, setQuantitys] = useState<EntQuantity[]>([]);
  const [staytypes, setStaytypes] = useState<EntStaytype[]>([]);
  const [equipments, setEquipments] = useState<EntEquipment[]>([]);
  const [facilities, setFacilities] = useState<EntFacilitie[]>([]);
  const [nearbyplaces, setNearbyplaces] = useState<EntNearbyplace[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  const [price, setRoomprice] = useState(String);
  const [noroom, setRoomnumber] = useState(String);
  const [roomname, setRoomtypename] = useState(String);
  const [quantity, setQuantity] = useState(Number);
  const [staytype, setStaytype] = useState(Number);
  const [equipment, setEquipment] = useState(Number);
  const [facilitie, setFacilitie] = useState(Number);
  const [nearbyplace, setNearbyplace] = useState(Number);
  const [employeeid, setEmployee] = useState(Number);

  

  useEffect(() => {


    const getRoomdetais = async () => {
 
      const rd = await api.listRoomdetail();
      setLoading(false);
      setRoomdetail(rd);
    };
    getRoomdetais();
 

    const getQuantitys = async () => {
 
      const qu = await api.listQuantity({ offset: 0 });
      setLoading(false);
      setQuantitys(qu);
    };
    getQuantitys();
 
    const getStaytypes = async () => {
 
    const st = await api.listStaytype({ offset: 0 });
      setLoading(false);
      setStaytypes(st);
    };
    getStaytypes();
 
    const getEquipments = async () => {
 
     const eq = await api.listEquipment({ offset: 0 });
       setLoading(false);
       setEquipments(eq);
     };
     getEquipments();

     const getFacilitys = async () => {
 
      const fa = await api.listFacilitie({ offset: 0 });
        setLoading(false);
        setFacilities(fa);
      };
      getFacilitys();

      const getNearbyplaces = async () => {
 
        const np = await api.listNearbyplace({ offset: 0 });
          setLoading(false);
          setNearbyplaces(np);
        };
        getNearbyplaces();

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
  
  const QuantityhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setQuantity(event.target.value as number);
  };

  const StaytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStaytype(event.target.value as number);
  };

  const EquipmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEquipment(event.target.value as number);
  };

  const FacilitiehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFacilitie(event.target.value as number);
  };

  const NearbyplacehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNearbyplace(event.target.value as number);
  };

  const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployee(event.target.value as number);
  };

  const listRoom = () => {
    window.location.href ="http://localhost:3000/RoomDetails";
  };
  const forCheck = () => {
    for (const color of roomdetail){
      if(noroom === color.roomnumber){
             setStatus(true);
             setAlert(false);
             setAlerts(false);
             //window.location.reload(false);
      }
      else{
        CreateRoomdetail();
      }
    }
  };

  const CreateRoomdetail = async () => {
    if ((noroom != "") && (noroom != null) && (roomname != "") && (roomname != null) && (price != "") && (price != null)
    && (quantity != null) && (quantity != null) && (equipment != null) && (facilitie != null) && (staytype != null)
    && (nearbyplace != null)){
    
      const roomdetail = {
      roomnumber: noroom,
      roomprice: price,
      roomtypename: roomname,
      quantity: quantity,
      equipment: equipment,
      facilitie: facilitie,
      nearbyplace: nearbyplace,
      staytype: staytype,
      employee: employeeid,
    };
    console.log(roomdetail);
    const res: any = await api.createRoomdetail({ roomdetail: roomdetail });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
     // window.location.reload(false);
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
            <FormControl
              variant="outlined"
            >

            <div className={classes.paper}><strong>เลขห้อง(No room)</strong></div>
              <TextField className={classes.textField}
                id="roomnumber"
                label=""
                variant="outlined"
                //color="secondary"
                type="string"
                size="medium"
                value={noroom}
                onChange={handleNoroomChange}
              />      
               <div className={classes.paper}><strong>ประเภทห้อง(Room type)</strong></div>
              <TextField className={classes.textField}
                id="roomtypename"
                label=""
                variant="outlined"
                //color="secondary"
                type="string"
                size="medium"
                value={roomname}
                onChange={handleRoomtypenameChange}
              />

            <div className={classes.paper}><strong>ราคาห้อง(Price)</strong></div>
              <TextField className={classes.textField}
                id="roomprice"
                label=""
                variant="outlined"
                //color="secondary"
                type="string"
                size="medium"
                value={price}
                onChange={handleRoompriceChange}
              />

              <div className={classes.paper}><strong>จำนวนผู้ที่สามารถเข้าพักได้(Quantity)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                labelId="quantity-label"
                id="quantity"
                value={quantity}
                onChange={QuantityhandleChange}
              >
                <InputLabel className={classes.insideLabel} id="faculty-label">Quantity</InputLabel>

                {quantitys.map((item: EntQuantity) => (
                  <MenuItem value={item.id}>{item.quantity}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ประเภทการเข้าพัก(Stay type)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="staytype"
                value={staytype}
                onChange={StaytypehandleChange}
              >
                <InputLabel className={classes.insideLabel}>Stay type</InputLabel>

                {staytypes.map((item: EntStaytype) => (
                  <MenuItem value={item.id}>{item.staytype}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>สิ่งที่ติดตั้งให้ภายในห้อง(Equipment)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="equipment"
                value={equipment}
                onChange={EquipmenthandleChange}
              >
                <InputLabel className={classes.insideLabel}>Equipment</InputLabel>

                {equipments.map((item: EntEquipment) => (
                  <MenuItem value={item.id}>{item.equipment}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>สิ่งอำนวยความสะดวกทั่วไป(General facilities)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="facilitie"
                value={facilitie}
                onChange={FacilitiehandleChange}
              >
                <InputLabel className={classes.insideLabel}>Facilities</InputLabel>

                {facilities.map((item: EntFacilitie) => (
                  <MenuItem value={item.id}>{item.facilitie}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>สถานที่ใกล้เคียง(Nearby place)</strong></div>
              <Select className={classes.select}
                //color="secondary"
                id="nearbyplace"
                value={nearbyplace}
                onChange={NearbyplacehandleChange}
              >
                <InputLabel className={classes.insideLabel}>Nearby place</InputLabel>

                {nearbyplaces.map((item: EntNearbyplace) => (
                  <MenuItem value={item.id}>{item.nearbyplace}</MenuItem>
                ))}
              </Select>

                <div className={classes.paper}><strong>พนักงานที่ทำการบันทึก(Employee)</strong></div>
                <TextField className={classes.select}
                    id="employee"
                    size="medium"
                    value={employees.filter((filter:EntEmployee) => filter.id == employeeid).map((item:EntEmployee) => `${item.name} (${item.email}) ตำแหน่ง (${item.edges?.jobposition?.positionname})`)}
                    style={{ width: 500 }}/>
             


              <Button
                style={{ width: 500, backgroundColor: "#5319e7",marginTop: 30,marginLeft: 7}}
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
                          <Alert severity="warning" style={{ marginTop: 20, marginLeft:5 }} onClose={() => {window.location.reload(false)}}>
                          มีข้อมูลนี้อยู่ในระบบแล้ว
                          </Alert>
                      :
                      (alert) ? (
                        <Alert severity="success" style={{ marginTop: 20, marginLeft:5 }} onClose={() => {listRoom()}}>
                            บันทึกสำเร็จ
                        </Alert>
                    ) : (
                            <Alert severity="warning" style={{ marginTop: 20, marginLeft:5 }} onClose={() => {setStatus(false)}}>
                                บันทึกไม่สำเร็จ กรุณาใส่ข้อมูลให้ครบ
                            </Alert>
                        )
                    }
                        </div>
                      ) : null}</div>
          
            </FormControl>
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
