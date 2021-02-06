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

import { EntBill } from '../../api/models/EntBill';
import { EntSituation } from '../../api/models/EntSituation';
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
    parcelcodesearchcheck: true,
    situationsearchcheck: true,
	leasesearchcheck: true
}

export default function Search() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [bill, setBill] = useState<EntBill[]>([]);
    const [situations, setSituations] = useState<EntSituation[]>([]);  
	const [leases, setLeases] = useState<EntLease[]>([]);   

    const [situationsearch, setSituationSearch] = useState(Number);
    const [leasesearch, setLeaseSearch] = useState(Number);

  useEffect(() => {
    const getSituations = async () => {
    const res = await api.listSituation({ offset: 0 });
      setLoading(false);
      setSituations(res);
    };
    getSituations();

	const getLeases = async () => {
    const res = await api.listLease({ offset: 0 });
      setLoading(false);
      setLeases(res);
    };
    getLeases();

    }, [loading]);

    const SearchBill = async () => {
        const res = await api.listBill();
		const leasesearch = LeaseSearch(res);
        const situationsearch = SituationSearch(leasesearch);
        
        setErrorMessege("not found");
        setAlertType("error");
        setBill([]);
        if(situationsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("search found");
                    setAlertType("success");
                    setBill(situationsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.leasesearchcheck = true;
		searchcheck.situationsearchcheck = true;
    }

    const SituationSearch = (res: any) => {
        const data = res.filter((filter: EntBill) => filter.edges?.situation?.id == situationsearch)
        //console.log(data.length)
        if (data.length != 0) {
            return data;
        }
        else {
            searchcheck.situationsearchcheck = false;
            if(situationsearch == 0){
                return res;
            }
            else{
                return data;
            }
        }
    }

	const LeaseSearch = (res: any) => {
        const data = res.filter((filter: EntBill) => filter.edges?.lease?.id == leasesearch)
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

    const SituationhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setSituationSearch(event.target.value as number);
    };

	const LeasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setLeaseSearch(event.target.value as number);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="Search Bill">

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
                     <div className={classes.paper}><strong>Status</strong></div>
                    <Select
                       // labelId="label"
                        id="situation"
                        value={situationsearch}
                        onChange={SituationhandleChange}
                        style={{ width: 200, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {situations.map((item: EntSituation) => (
                                <MenuItem value={item.id}>{item.situationname}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

				<div>

                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchBill();
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
                                <TableCell align="center">No</TableCell>
                                <TableCell align="center">Name</TableCell>
                                <TableCell align="center">Total</TableCell>
                                <TableCell align="center">Situation</TableCell>
                                <TableCell align="center">Payment</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {bill.map((item:any ) => (
                                <TableRow key={item.id}>
                                       <TableCell align="center">{item.id}</TableCell> 
                                        <TableCell align="center">{item.edges?.lease?.tenant}</TableCell> 
                                        <TableCell align="center">{item.total}</TableCell>
                                        <TableCell align="center">{item.edges?.situation?.situationname}</TableCell>
                                        <TableCell align="center">{item.edges?.payment?.paymentname}</TableCell>
                                
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