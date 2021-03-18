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

import SearchIcon from '@material-ui/icons/Search';
import { EntCleaningroom } from '../../api/models/EntCleaningroom';
import { EntRoomdetail } from '../../api/models/EntRoomdetail';


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
    },
    table: {
      minWidth: 650,
    },
  }),
);

const searchcheck = {
  roomdetailsearchcheck: true,
  phonenumbersearchcheck: true
}

export default function SearchCleaningroom() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);

  const [cleaningrooms, setCleaningrooms] = useState<EntCleaningroom[]>([]);
  const [roomdetails, setRoomdetails] = useState<EntRoomdetail[]>([]);


  const [phonenumbersearch, setPhonenumberSearch] = useState(String);
  const [roomdetailsearch, setRoomdetailSearch] = useState(Number);

  useEffect(() => {
    const getRoomdetails = async () => {
      const res = await api.listRoomdetail();
      setLoading(false);
      setRoomdetails(res);
    };
    getRoomdetails();


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

  const SearchCleaningroom = async () => {
    const apiUrl = `http://localhost:8080/api/v1/searchcleaningrooms?phonenumber=${phonenumbersearch}&roomdetail=${roomdetailsearch}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setCleaningrooms([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setErrorMessege("พบข้อมูลที่ค้นหา");
            setAlertType("success");
            console.log(data.data)
            setCleaningrooms(data.data);
          }
        }

        setStatus(true);
      });

  }

  const RoomdetailSearchhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomdetailSearch(event.target.value as number);
  };

  const PhonenumberSearchhandleChange = (event: any) => {
    setPhonenumberSearch(event.target.value as string);
  };


  return (
    <Page>
      <Content>
        <InfoCard title={'ค้นหาข้อมูลแจ้งทำความสะอาด'}>

          <FormControl
            className={classes.margin}
            variant="standard"
          >
            <div className={classes.paper}><strong>เบอร์โทรศัพท์</strong></div>
            <TextField
              id="phonenumber"
              type="string"
              size="medium"
              value={phonenumbersearch}
              onChange={PhonenumberSearchhandleChange}
              style={{ width: 210, marginLeft: 8 }}
            />
          </FormControl>

          <FormControl
            className={classes.margin}
            variant="standard"
          >
            <div className={classes.paper}><strong>เลขห้อง</strong></div>
            <Select
              // labelId="label"
              id="roomdetail"
              value={roomdetailsearch}
              onChange={RoomdetailSearchhandleChange}
              style={{ width: 175, marginLeft: 8 }}

            >   <MenuItem value={0}>-</MenuItem>
              {roomdetails.map((item: EntRoomdetail) => (
                <MenuItem value={item.id}>{item.roomnumber}</MenuItem>
              ))}
            </Select>
          </FormControl>




          <FormControl
            className={classes.margin}
          >
            <Button
              style={{ width: 250, backgroundColor: "#5319e7", marginTop: 20 }}
              onClick={() => {
                SearchCleaningroom();
              }}
              variant="contained"
              color="primary"
              startIcon={<SearchIcon style={{ fontSize: 'medium' }} />}
            >
              ค้นหาข้อมูลแจ้งทำความสะอาด
             </Button>

          </FormControl>

          <div><br></br></div>
          <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell align="center">เลขที่</TableCell>
                  <TableCell align="center">ชื่อพนักงานที่ทำการบันทึก</TableCell>
                  <TableCell align="center">ชื่อพนักงานที่รับผิดชอบ</TableCell>
                  <TableCell align="center">จำนวนพนักงานในการทำความสะอาด</TableCell>
                  <TableCell align="center">เลขห้อง</TableCell>
                  <TableCell align="center">ระยะเวลา</TableCell>
                  <TableCell align="center">วันที่และเวลาเริ่มทำความสะอาด</TableCell>
                  <TableCell align="center">เบอร์โทรศัพท์ที่สามารถติดต่อได้</TableCell>
                  <TableCell align="center">เพิ่มเติม</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {cleaningrooms.map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item.edges?.Employee?.name}</TableCell>
                    <TableCell align="center">{item.edges?.Cleanername?.cleanername}</TableCell>
                    <TableCell align="center">{item.numofem}</TableCell>
                    <TableCell align="center">{item.edges?.Roomdetail?.roomnumber}</TableCell>
                    <TableCell align="center">{item.edges?.Lengthtime?.lengthtime}</TableCell>
                    <TableCell align="center">{item.dateandstarttime}</TableCell>
                    <TableCell align="center">{item.phonenumber}</TableCell>
                    <TableCell align="center">{item.note}</TableCell>
                    
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
                  <Alert severity={alerttype} style={{ width: 400, marginTop: 20 }} onClose={() => { setStatus(false) }}>
                    {errormessege}
                  </Alert>
                ) : null}
              </div>
            ) : null}</div>

      </Content>
    </Page>
  );
}