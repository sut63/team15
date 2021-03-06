import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
    InfoCard,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { Alert } from '@material-ui/lab';

import { makeStyles, Theme, createStyles, ThemeProvider } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import { EntStaytype } from '../../api/models/EntStaytype';
import SearchIcon from '@material-ui/icons/Search';
import { EntBedtype } from '../../api/models/EntBedtype';
import { EntPetrule } from '../../api/models/EntPetrule';


// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(1.5),
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
    },
    table: {
        minWidth: 650,
    },
  }),
);

const searchcheck = {
    namesearchcheck: true,
    pricesearchcheck: true,
    staytypesearchcheck: true,
    bedtypesearchcheck: true,
    petrulesearchcheck: true
}

export default function SearchRoom() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [roomdetail, setRoomdetail] = useState<EntRoomdetail[]>([]);
    const [staytypes, setStaytypes] = useState<EntStaytype[]>([]);
    const [bedtypes, setBedtypes] = useState<EntBedtype[]>([]);
    const [petrules, setPetrules] = useState<EntPetrule[]>([]);
    

    const [namesearch, setRoomtypenameSearch] = useState(String);
    const [pricesearch, setRoompriceSearch] = useState(Number);
    const [staytypesearch, setStaytypeSearch] = useState(Number);
    const [bedtypesearch, setBedtypeSearch] = useState(Number);
    const [petrulesearch, setPetruleSearch] = useState(Number);

  useEffect(() => {
    const getStaytypes = async () => {
    const res = await api.listStaytype({ offset: 0 });
      setLoading(false);
      setStaytypes(res);
    };
    getStaytypes();

    const getBedtypes = async () => {
        const res = await api.listBedtype();
          setLoading(false);
          setBedtypes(res);
        };
        getBedtypes();

        const getPetrules = async () => {
            const res = await api.listPetrule();
              setLoading(false);
              setPetrules(res);
            };
            getPetrules();    

    {/*const checkEmployeeLoginData = async () => {
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
    checkEmployeeLoginData();*/} 

    }, [loading]);

        const SearchRoomdetail = async () => {
            const apiUrl = `http://localhost:8080/api/v1/searchroomdetails?price=${pricesearch}&roomtypename=${namesearch}&bedtype=${bedtypesearch}&staytype=${staytypesearch}&petrule=${petrulesearch}`;
            const requestOptions = {
                method: 'GET',
            };
            fetch(apiUrl, requestOptions)
                .then(response => response.json())
                .then(data => {
                    console.log(data.data)
                    setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
                    setAlertType("error");
                    setRoomdetail([]);
                    if (data.data != null) {
                        if(data.data.length >= 1) {
                            setErrorMessege("พบข้อมูลที่ค้นหา");
                            setAlertType("success");
                            console.log(data.data)
                            setRoomdetail(data.data);
                        }
                    }
    
                    setStatus(true);
                });
    
        }
    

    const StaytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setStaytypeSearch(event.target.value as number);
    };

    const BedtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setBedtypeSearch(event.target.value as number);
    };

    const PetrulehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPetruleSearch(event.target.value as number);
    };

    const PriceSearchhandleChange = (event: any) => {
        setRoompriceSearch(event.target.value as number);
    };

    const NameSearchhandleChange = (event: any) => {
        setRoomtypenameSearch(event.target.value as string);
    };
    

    return (
        <Page>
            <Content>
            <InfoCard title={'ค้นหาข้อมูลห้องพัก'}>

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>ชื่อประเภทห้อง(Room type name)</strong></div>
                    <TextField
                        id="roomtypenamesearch"
                       // label="ค้นหาเลขห้อง"
                        type="string"
                        size="medium"
                        value={namesearch}
                        onChange={NameSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>ราคาห้องพักเริ่มต้น(Price)</strong></div>
                    <TextField
                        id="pricesearch"
                       // label="ค้นหาเลขห้อง"
                        type="number"
                        size="medium"
                        value={pricesearch}
                        onChange={PriceSearchhandleChange}
                        style={{ width: 150,marginLeft: 8}}
                    />
                </FormControl>

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>ประเภทการเข้าพัก(Stay type)</strong></div>
                    <Select
                       // labelId="label"
                        id="staytype"
                        value={staytypesearch}
                        onChange={StaytypehandleChange}
                        style={{ width: 175, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {staytypes.map((item: EntStaytype) => (
                                <MenuItem value={item.id}>{item.staytype}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>ประเภทเตียงนอน(Bed type)</strong></div>
                    <Select
                       // labelId="label"
                        id="bedtype"
                        value={bedtypesearch}
                        onChange={BedtypehandleChange}
                        style={{ width: 170, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {bedtypes.map((item: EntBedtype) => (
                                <MenuItem value={item.id}>{item.bedtypename}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>ข้อจำกัดในการเลี้ยงสัตว์</strong></div>
                    <Select
                       // labelId="label"
                        id="petrule"
                        value={petrulesearch}
                        onChange={PetrulehandleChange}
                        style={{ width: 140, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {petrules.map((item: EntPetrule) => (
                                <MenuItem value={item.id}>{item.petrule}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

                <FormControl
                    className={classes.margin}
                >
                    <Button
                style={{ width: 165, backgroundColor: "#5319e7", marginTop: 20}}
                onClick={() => {
                  SearchRoomdetail();
                }}
                variant="contained"
                color="primary"
                startIcon={<SearchIcon style={{fontSize: 'medium'}} />}
              >
                ค้นหาข้อมูลห้องพัก
             </Button>
             
                </FormControl>
        
             <div><br></br></div>
             <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                <TableCell align="center">เลขห้อง</TableCell>
                                <TableCell align="center">ชื่อประเภทห้อง</TableCell>
                                <TableCell align="center">ราคาห้อง/ประเภทการเข้าพัก</TableCell>
                                <TableCell align="center">จำนวนผู้เข้าพักสูงสุด</TableCell>
                                <TableCell align="center">ประเภทเตียงนอน/จำนวนเตียง</TableCell>
                                <TableCell align="center">ข้อจำกัดการจ่ายมัดจำ</TableCell>
                                <TableCell align="center">ข้อจำกัดในการเลี้ยงสัตว์</TableCell>
                                <TableCell align="center">เบอร์โทรศัพท์</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {roomdetail.map((item:any ) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.roomnumber}</TableCell>
                                    <TableCell align="center">{item.roomtypename}</TableCell>
                                    <TableCell align="center">{item.roomprice} / {item.edges?.Staytype?.staytype}</TableCell>
                                    <TableCell align="center">{item.sleep} คน</TableCell>
                                    <TableCell align="center">{item.edges?.Bedtype?.bedtypename} / {item.bed} เตียง</TableCell>
                                    <TableCell align="center">{item.edges?.Pledge?.provision}</TableCell>
                                    <TableCell align="center">{item.edges?.Petrule?.petrule}</TableCell>
                                    <TableCell align="center">{item.phone}</TableCell>
                                
                                </TableRow>
                                ))}
                            </TableBody>
                            </Table>
                        </TableContainer>

                        

                </InfoCard>

                        <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20}} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>

            </Content>
     </Page>
    );
}