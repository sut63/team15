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
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import { Link as RouterLink } from 'react-router-dom';
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
import { EntWifi } from '../../api/models/EntWifi';
import { EntLease } from '../../api/models/EntLease';
import SearchIcon from '@material-ui/icons/Search';


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
    },
    table: {
        minWidth: 650,
    }
  }),
);

const searchcheck = {
    tenantsearchcheck: true,
    wifisearchcheck: true,
	roomdetailsearchcheck: true
}

export default function Search() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [roomdetails, setRoomdetails] = useState<EntRoomdetail[]>([]);
    const [wifis, setWifis] = useState<EntWifi[]>([]);  
	const [lease, setLease] = useState<EntLease[]>([]);   

    const [tenantsearch, setTenantSearch] = useState(String);
    const [wifisearch, setWifiSearch] = useState(Number);
    const [roomdetailsearch, setRoomdetailSearch] = useState(Number);

  useEffect(() => {
    const getWifis = async () => {
    const res = await api.listWifi({ offset: 0 });
      setLoading(false);
      setWifis(res);
    };
    getWifis();

	const getRoomdetails = async () => {
    const res = await api.listRoomdetail();
      setLoading(false);
      setRoomdetails(res);
    };
    getRoomdetails();

    }, [loading]);

    const SearchLease = async () => {
		const apiUrl = `http://localhost:8080/api/v1/searchleases?lease=${tenantsearch}&roomdetail=${roomdetailsearch}&wifi=${wifisearch}`;
		const requestOptions = {
			method: 'GET',
		};
		fetch(apiUrl, requestOptions)
		.then(response => response.json())
		.then(data => {
		console.log(data.data)
		setErrorMessege("not found");
		setAlertType("error");
        setLease([]);
        if (data.data != null) {
			if(data.data.length >= 1) {
				setErrorMessege("search found");
                setAlertType("success");
                console.log(data.data)
                setLease(data.data);
                }
            }
    
            setStatus(true);
        });
    
     }

    const WifihandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setWifiSearch(event.target.value as number);
    };

	const RoomdetailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setRoomdetailSearch(event.target.value as number);
    };

    const TenantSearchhandleChange = (event: any) => {
        setTenantSearch(event.target.value as string);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="Search Lease">

			<FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>Room number</strong></div>
                    <Select
                       // labelId="label"
                        id="roomdetail"
                        value={roomdetailsearch}
                        onChange={RoomdetailhandleChange}
                        style={{ width: 200, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {roomdetails.map((item: EntRoomdetail) => (
                                <MenuItem value={item.id}>{item.roomnumber}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>Tenant's name</strong></div>
                    <TextField
                        id="tenantcodesearch"
                       // label="�����Ţ��ͧ"
                        type="string"
                        size="medium"
                        value={tenantsearch}
                        onChange={TenantSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>       

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>Wifi status</strong></div>
                    <Select
                       // labelId="label"
                        id="wifi"
                        value={wifisearch}
                        onChange={WifihandleChange}
                        style={{ width: 200, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {wifis.map((item: EntWifi) => (
                                <MenuItem value={item.id}>{item.wifiname}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

				<div>

                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchLease();
                }}
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                Search
             </Button>

			 <Button style={{ width: 100,marginTop: 49,marginLeft: 20}} 
			 component={RouterLink} to="/DormEmployee" variant="contained" startIcon={<CancelRoundedIcon/>}>  Home </Button>

			 </div>
        
             <div><br></br></div>
             <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                <TableCell align="center">Contract No.</TableCell>
                                <TableCell align="center">Room no.</TableCell>
                                <TableCell align="center">Name</TableCell>
                                 <TableCell align="center">Age</TableCell>
                                 <TableCell align="center">Identified ID</TableCell>
                                <TableCell align="center">Phone No.</TableCell>
                                <TableCell align="center">Wifi Status</TableCell>
		                        <TableCell align="center">Time</TableCell>
								
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {lease.map((item:any ) => (
                                <TableRow key={item.id}>
                                     <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.Roomdetail?.roomnumber}</TableCell>
                                     <TableCell align="center">{item.tenant}</TableCell>
                                    <TableCell align="center">{item.agetenant}</TableCell>
                                     <TableCell align="center">{item.idtenant}</TableCell>
                                    <TableCell align="center">{item.numbtenant}</TableCell>
                                    <TableCell align="center">{item.edges?.Wifi?.wifiname}</TableCell>
			                        <TableCell align="center">{item.addedtime}</TableCell>                         
                                </TableRow>
                                ))}
                            </TableBody>
                            </Table>
                        </TableContainer>

                        <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>

                </InfoCard>

                        

            </Content>
     </Page>
    );
}