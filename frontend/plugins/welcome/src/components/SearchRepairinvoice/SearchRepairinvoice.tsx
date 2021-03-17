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

import { EntRepairinvoice } from '../../api/models/EntRepairinvoice';
import { EntRentalstatus } from '../../api/models/EntRentalstatus';
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
    bequipmentsearchcheck: true,
	leasesearchcheck: true
}

export default function Search() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [repairinvoice, setRepairinvoice] = useState<EntRepairinvoice[]>([]);  
	const [leases, setLeases] = useState<EntLease[]>([]);   

    const [bequipmentsearch, setRepairBequipmentSearch] = useState(String);
    const [leasesearch, setLeaseSearch] = useState(Number);

  useEffect(() => {
	const getLeases = async () => {
    const res = await api.listLease({ offset: 0 });
      setLoading(false);
      setLeases(res);
    };
    getLeases();

    }, [loading]);

    const SearchRepairinvoice = async () => {
		const apiUrl = `http://localhost:8080/api/v1/searchrepairinvoices?lease=${leasesearch}&bequipment=${bequipmentsearch}`;
		const requestOptions = {
			method: 'GET',
		};
		fetch(apiUrl, requestOptions)
		.then(response => response.json())
		.then(data => {
		console.log(data.data)
		setErrorMessege("not found");
		setAlertType("error");
        setRepairinvoice([]);
        if (data.data != null) {
			if(data.data.length >= 1) {
				setErrorMessege("search found");
                setAlertType("success");
                console.log(data.data)
                setRepairinvoice(data.data);
                }
            }
            setStatus(true);
        });
    
     }

    const ReBequipmentSearch = (res: any) => {
        const data = res.filter((filter: EntRepairinvoice) => filter.bequipment?.includes(bequipmentsearch))
        //console.log(data.length)
        if (data.length != 0 && bequipmentsearch != "") {
            return data;
        }
        else {
            searchcheck.bequipmentsearchcheck = false;
            if(bequipmentsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

	const LeaseSearch = (res: any) => {
        const data = res.filter((filter: EntRepairinvoice) => filter.edges?.lease?.id == leasesearch)
        //console.log(data.length)
        if (data.length != 0) {
            return data;
        }
        else {
            searchcheck.leasesearchcheck = false;
            if(leasesearch == 0){
                return res;
            }
            else{
                return data;
            }
        }
    }

	const LeasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setLeaseSearch(event.target.value as number);
    };

    const BequipmentSearchhandleChange = (event: any) => {
        setRepairBequipmentSearch(event.target.value as string);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="Search Repairinvoice">

			<FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>Name</strong></div>
                    <Select
                       // labelId="label"
                        id="lease"
                        value={leasesearch}
                        onChange={LeasehandleChange}
                        style={{ width: 200, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {leases.map((item: EntLease) => (
                                <MenuItem value={item.id}>{item.tenant}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>Broken Equipment</strong></div>
                    <TextField
                        id="repairbequipmentsearch"
                       // label="�����Ţ��ͧ"
                        type="string"
                        size="medium"
                        value={bequipmentsearch}
                        onChange={BequipmentSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>       

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchRepairinvoice();
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
                                <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">ผู้เช่า</TableCell>
           <TableCell align="center">อุปกรณ์ที่ชำรุด</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์ของพนักงาน</TableCell>
           <TableCell align="center">จำนวนอุปกรณ์ที่ชำรุด</TableCell>
           
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {repairinvoice.map((item:any ) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.Lease?.tenant}</TableCell>
             <TableCell align="center">{item.bequipment}</TableCell>
             <TableCell align="center">{item.emtell}</TableCell>
             <TableCell align="center">{item.num}</TableCell>   
                    
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