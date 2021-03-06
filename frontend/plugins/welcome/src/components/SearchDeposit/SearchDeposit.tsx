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

import { EntDeposit } from '../../api/models/EntDeposit';
import { EntStatusd } from '../../api/models/EntStatusd';
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
    statusdsearchcheck: true,
	leasesearchcheck: true
}

export default function Search() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [deposit, setDeposit] = useState<EntDeposit[]>([]);
    const [statusds, setStatusds] = useState<EntStatusd[]>([]);  
	const [leases, setLeases] = useState<EntLease[]>([]);   

    const [parcelcodesearch, setDepositParcelcodeSearch] = useState(String);
    const [statusdsearch, setStatusdSearch] = useState(Number);
    const [leasesearch, setLeaseSearch] = useState(Number);

  useEffect(() => {
    const getStatusds = async () => {
    const res = await api.listStatusd({ offset: 0 });
      setLoading(false);
      setStatusds(res);
    };
    getStatusds();

	const getLeases = async () => {
    const res = await api.listLease({ offset: 0 });
      setLoading(false);
      setLeases(res);
    };
    getLeases();

    }, [loading]);

	const SearchDeposit = async () => {
		const apiUrl = `http://localhost:8080/api/v1/searchdeposits?lease=${leasesearch}&parcelcode=${parcelcodesearch}&statusd=${statusdsearch}`;
		const requestOptions = {
			method: 'GET',
		};
		fetch(apiUrl, requestOptions)
		.then(response => response.json())
		.then(data => {
		console.log(data.data)
		setErrorMessege("not found");
		setAlertType("error");
        setDeposit([]);
        if (data.data != null) {
			if(data.data.length >= 1) {
				setErrorMessege("search found");
                setAlertType("success");
                console.log(data.data)
                setDeposit(data.data);
                }
            }
    
            setStatus(true);
        });
    
     }

    const StatusdhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setStatusdSearch(event.target.value as number);
    };

	const LeasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setLeaseSearch(event.target.value as number);
    };

    const ParcelcodeSearchhandleChange = (event: any) => {
        setDepositParcelcodeSearch(event.target.value as string);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="Search Deposit">

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
                    <div className={classes.paper}><strong>Parcel Code</strong></div>
                    <TextField
                        id="parcelcodesearch"
                       // label="�����Ţ��ͧ"
                        type="string"
                        size="medium"
                        value={parcelcodesearch}
                        onChange={ParcelcodeSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>       

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                     <div className={classes.paper}><strong>Status</strong></div>
                    <Select
                       // labelId="label"
                        id="statusd"
                        value={statusdsearch}
                        onChange={StatusdhandleChange}
                        style={{ width: 200, marginLeft: 8 }}

                    >   <MenuItem value={0}>-</MenuItem>
                         {statusds.map((item: EntStatusd) => (
                                <MenuItem value={item.id}>{item.statusdname}</MenuItem>
                            ))}
                    </Select>
                </FormControl>

				<div>

                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchDeposit();
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
								<TableCell align="center">Recipient Tell</TableCell>
								<TableCell align="center">Info</TableCell>
								<TableCell align="center">Parcel Code</TableCell>
								<TableCell align="center">Depositor Name</TableCell>
								<TableCell align="center">Depositor Tell</TableCell>
								<TableCell align="center">Status</TableCell>
								<TableCell align="center">Time</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {deposit.map((item:any ) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
									<TableCell align="center">{item.edges?.Lease?.tenant}</TableCell>
									<TableCell align="center">{item.recipienttell}</TableCell>
									<TableCell align="center">{item.info}</TableCell>
									<TableCell align="center">{item.parcelcode}</TableCell>
									<TableCell align="center">{item.depositorname}</TableCell>
									<TableCell align="center">{item.depositortell}</TableCell>
									<TableCell align="center">{item.edges?.Statusd?.statusdname}</TableCell>
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