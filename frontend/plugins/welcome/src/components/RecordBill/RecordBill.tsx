import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';
import { Alert, AlertTitle } from '@material-ui/lab';

import React, { useState, useEffect, FC } from 'react';  
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
  Typography,
  Toolbar,
  AppBar,
} from '@material-ui/core';

import { DefaultApi } from '../../api/apis';
import { EntSituation } from '../../api/models/EntSituation'; // import interface Situation
import { EntBill } from '../../api/models/EntBill'; // import interface Bill
import ComponanceBillTable from '../ฺBillTable';
import { EntPayment } from '../../api/models/EntPayment'; // import interface Payment
import { EntLease } from '../../api/models/EntLease'; // import interface Lease

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
      //marginTop:10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

export default function recordBill() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [bill, setBill] = React.useState<Partial<recordBill>>({});

 const [situations, setSituations] = React.useState<EntSituation[]>([]);
 const [payments, setPayments] = React.useState<EntPayment[]>([]);
 const [leases, setLeases] = React.useState<EntLease[]>([]);

 const [status, setStatus] = useState(false);
 const [status2, setStatus2] = useState(false);
 const [alert, setAlert] = useState(false);

 const [loading, setLoading] = useState(true);

 const [situation, setSituation] = useState(Number);
 const [payment, setPayment] = useState(Number);
 const [lease, setLease] = useState(Number);

 const [total, setTotal] = useState(String);
 const [tell, setTell] = useState(String);
 const [taxpayer, setTaxpayer] = useState(String);
 const [added, setAdded] = useState(String);

 useEffect(() => { 

  const getLeases = async () => {
    const res = await http.listLease({ offset: 0 });
    setLoading(false);
    setLeases(res);
    console.log(res);
  };
  getLeases();

    const getSituations = async () => {
    const si = await http.listSituation({ offset: 0 });
      setLoading(false);
      setSituations(si);
    };
    getSituations();

    const getPayments = async () => {
    const p = await http.listPayment({ offset: 0 });
      setLoading(false);
      setPayments(p);
    };
    getPayments();
    


}, [loading]);


const getBill = async () => {
  const res = await http.listBill({ offset: 0 });
  setBill(res);
};

const TotalhandleChange = (
  event: React.ChangeEvent<{ name: string; value: string }>,) => {
  setTotal(event.target.value as string);
};

const TellhandleChange = (
  event: React.ChangeEvent<{ name: string; value: string }>,) => {
  setTell(event.target.value as string);
};

const TaxpayerhandleChange = (
  event: React.ChangeEvent<{ name: string; value: string }>,) => {
  setTaxpayer(event.target.value as string);
};

const AddedhandleChange = (event: any) => {
  setAdded(event.target.value as string);
};

  const SituationhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSituation(event.target.value as number);
  };

  const LeasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLease(event.target.value as number);
  };

  const PaymenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPayment(event.target.value as number);
  };
 

 
// create bill
const CreateBill = async () => {
  const bills = {
    added: added + ":00+07:00",
    situation: situation,
    payment: payment,
    total: total,
    tell: tell,
    taxpayer: taxpayer,
    lease: lease,
  };
  console.log(bills);
  const timer2 = setTimeout(() => {
     setStatus2(true);
  }, 2000);
  const timer3 = setTimeout(() => {
     setStatus2(false);
  }, 6000);

  const res: any = await http.createBill({ bill: bills });
  if (res.id != '') {
    setAlert(true);
	setStatus2(false);
  }
  setStatus(true);
  const timer = setTimeout(() => {
     setStatus(false);
	 //window.location.reload(false);
  }, 7000);
  
  console.log(bills);
};

    return (
    <Page theme={pageTheme.tool}>

      <Content>
        <ContentHeader title="Bill invoice"> 
              <Button onClick={() => {CreateBill();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new bill invoice </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              //fullWidth
              //className={classes.margin}
              variant="outlined"
            >

<div className={classes.paper}><strong>Room</strong></div>
			  <Select
                  name="Lease"
				  id="lease"
                  value={lease} 
                  onChange={LeasehandleChange}
                >
                {leases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.tenant}</MenuItem>
                    );
                  })}
              </Select>
              
              <div className={classes.paper}><strong> Situation</strong></div>
			  <Select
                  name="Situation"
				  id="situation"
                  value={situation} 
                  onChange={SituationhandleChange}
                >
                 {situations.map((item: EntSituation) => (
                  <MenuItem value={item.id}>{item.situationname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong> Payment</strong></div>
			  <Select
                  name="Payment"
				  id="payment"
                  value={payment} 
                  onChange={PaymenthandleChange}
                >
                 {payments.map((item: EntPayment) => (
                  <MenuItem value={item.id}>{item.paymentname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>Date</strong></div>
			  <form className={classes.container} noValidate>
                <TextField
                  
                  name="added"
				  id="added"
                  type="datetime-local"
                  value={added} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={AddedhandleChange}
                />
              </form>

			  <div className={classes.paper}><strong>Total</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="total"
				  id="total"
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={total} 
                  onChange={TotalhandleChange}
                />
              </form>

              <div className={classes.paper}><strong>Tell</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="tell"
				  id="tell"
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={tell} 
                  onChange={TellhandleChange}
                />
              </form>

              <div className={classes.paper}><strong>Taxpayer</strong></div>
			  <form className={classes.root} noValidate>
                <TextField
                  
                  name="taxpayer"
				  id="taxpayer"
				  style={{ margin: 8 }}
                  type="text"
				  fullWidth
				  margin="normal"
				  InputLabelProps={{
					shrink: true,
				  }}
                  value={taxpayer} 
                  onChange={TaxpayerhandleChange}
                />
              </form>

			  {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
			  <Alert severity="success" style={{ marginTop: 20, marginLeft:5 }} >Successfull Save</Alert>
                      ) 
              : null} </div>
            ) : null}
			{status2 ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      null) 
              : (<Alert severity="warning" style={{ marginTop: 20, marginLeft:5 }}> Unsuccessfull Save!! </Alert>)} </div>
            ) : null}
		
            </FormControl>
			<div>
			<ComponanceBillTable></ComponanceBillTable>
			</div>

          </form>
        </div>
      </Content>
    </Page>
  );
 }
