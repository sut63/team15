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
import SearchIcon from '@material-ui/icons/Search';

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
 const [alert2, setAlerts] = useState(true);
 const [alerttype, setAlertType] = useState(String);
 const [errormessege, setErrorMessege] = useState(String);

 const [loading, setLoading] = useState(true);

 const [tellerror, setTellerror] = React.useState('');
 const [taxpayererror, setTaxpayererror] = React.useState('');
 const [totalerror, setTotalerror] = React.useState('');

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

const validateTotal = (val: string) => {
  return val.match("^(([0-9]{1}).([0-9]{2}))$|^(([0-9]{2}).([0-9]{2}))$|^(([0-9]{3}).([0-9]{2}))$|^(([0-9]{4}).([0-9]{2}))$");
}

const validateTell = (val: string) => {
  return val.match("^([0-9]{3})-([0-9]{3})-([0-9]{4})$");
}

const validateTaxpayer = (val: string) => {
  return val.match("^([0-9]{1})-([0-9]{4})-([0-9]{3})-([0-9]{4})-([0-9]{1})$");
}

const checkPattern  = (id: string, value:string) => {
  console.log(value);
  switch(id) {
    case 'tell':
      validateTell(value) ? setTellerror('') : setTellerror ('รูปแบบของหมายเลขโทรศัพท์ต้องเป็น xxx-xxx-xxxx');
    return;
    case 'taxpayer':
      validateTaxpayer(value) ? setTaxpayererror('') : setTaxpayererror ('รูปแบบของหมายเลขภาษีต้องเป็น x-xxxx-xxx-xxxx-x');
    return;
    case 'total':
      validateTotal(value) ? setTotalerror('') : setTotalerror ('ใส่รูปแบบจำนวนเงิน xxxx.xx');
    return;
      default:
        return;
  }
}

const getBill = async () => {
  const res = await http.listBill({ offset: 0 });
  setBill(res);
};

const TellhandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
  const id = event.target.id as  typeof tell;
  const { value } = event.target;
  const validateValue = value.toString()
  checkPattern(id, validateValue)
  setTell(event.target.value as string);
};

const TaxpayerhandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
  const id = event.target.id as  typeof taxpayer;
  const { value } = event.target;
  const validateValue = value.toString()
  checkPattern(id, validateValue)
  setTaxpayer(event.target.value as string);
};

const TotalhandleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
  const id = event.target.id as  typeof total;
  const { value } = event.target;
  const validateValue = value.toString()
  checkPattern(id, validateValue)
  setTotal(event.target.value as string);
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
 
  const checkCaseSaveError = (field: string) => {
    if (field == "tell") { setErrorMessege("ข้อมูล field เบอร์โทรศัพท์ผิด"); }
        else if (field == "taxpayer") { setErrorMessege("ข้อมูล field เลขภาษีผิด"); }
        else if (field == "total") { setErrorMessege("ข้อมูลfield จำนวนเงินผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ ใส่ข้อมูลไม่ครบ"); }
  }

  const CreateBill = async () => {
    if ((lease != null) && (payment != null) && (situation != null)){
	  const apiUrl = 'http://localhost:8080/api/v1/bills';
    const bill = {
      added: added + ":00+07:00",
      situation: situation,
      payment: payment,
      total: total,
      tell: tell,
      taxpayer: taxpayer,
      lease: lease,
    };
    console.log(bill);
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bill),
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
    <Page theme={pageTheme.tool}>

      <Content>
        <ContentHeader title="Bill invoice"> 
        <Button
                style={{ width: 200, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                component={RouterLink} to="/SearchBill"
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                Search Bill
             </Button>
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
          error = {totalerror ? true : false}
				  helperText= {totalerror}
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
          error = {tellerror ? true : false}
				  helperText= {tellerror}
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
          error = {taxpayererror ? true : false}
				  helperText= {taxpayererror}
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

              <div>
					<Button onClick={() => {CreateBill();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new bill </Button>
			  </div>

			  <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20, marginLeft:6 }} >
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>
		
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
