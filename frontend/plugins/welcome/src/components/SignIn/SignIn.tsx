import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  InfoCard,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Avatar from '@material-ui/core/Avatar';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import { EntEmployee } from '../../api/models/EntEmployee';
import { EntJobposition } from '../../api';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  head: {
    marginLeft: theme.spacing(70),
    fontSize: '18px',
  },
  avatar: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(84),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(2, 0, 2),
    marginLeft: theme.spacing(83),
  },
  textField: {
    width: 350 ,
    marginLeft: theme.spacing(67),
   },
   margin: {
    margin: theme.spacing(2),
 },
 signin: {
   margin: theme.spacing(2, 0, 2),
   width: 350 ,
   marginLeft: theme.spacing(67),
 },
 root: {
  display: 'flex',
  flexWrap: 'wrap',
  justifyContent: 'center',
},
}));


export default function Login(props: any) {
  const classes = useStyles();
  const api = new DefaultApi();

  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);

  const [email, setEmail] = useState(String);
  const [password, setPassword] = useState(String);

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getEmployees = async () => {
      const res: any = await api.listEmployee();
      setLoading(false);
      setEmployees(res);
    }

    getEmployees();

    const resetEmployeeData = async () => {
      setLoading(false);
      localStorage.setItem("employeedata", JSON.stringify(null)); 
      localStorage.setItem("jobpositiondata", JSON.stringify(null));
    }
    resetEmployeeData();

  }, [loading]);

  const EmailthandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };

  const PasswordthandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const LoginChecker = async () => {
    employees.map((item: any) => {
      console.log(item.email);
      if ((item.email == email) && (item.password == password)) {
       setAlert(true);
        localStorage.setItem("employeedata",JSON.stringify(item.id));
        console.log(item.email);
        localStorage.setItem("jobpositiondata", JSON.stringify(item.edges?.jobposition?.positionname))
        console.log(item.edges?.jobposition?.positionname);
        
      if (item.edges?.jobposition?.positionname == "พนักงานหอพัก") {
          history.pushState("", "", "/DormEmployee");
        }
          window.location.reload(false);
        }
    })
    setStatus(true);
     //const timer = setTimeout(() => {
     //  setStatus(false);
     //}, 10000);
  };

  return (
    <Page theme={pageTheme.tool}>
  <Header
        title={`ยินดีต้อนรับสู่ ระบบหอพัก`}
        subtitle="หอพักทั่วไป"
      ></Header>

<Content>
       
        
        <InfoCard><div className={classes.root}> 
        <ContentHeader title="กรุณา Login ก่อนการใช้งาน">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success" onClose={() => { }}>
                  เข้าสู่ระบบสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="error" onClose={() => { setStatus(false) }}>
                    เข้าสู่ระบบไม่สำเร็จ กรุณาตรวจสอบ Email หรือ Password
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
          <form noValidate autoComplete="off">
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
              </FormControl>
            </div>
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="email"
                  label="Email"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={email}
                  onChange={EmailthandleChange}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="password"
                  label="Password"
                  variant="outlined"
                  type="password"
                  size="medium"
                  value={password}
                  onChange={PasswordthandleChange}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>

            <div className={classes.margin}>
              <Button
                style={{ width: 100, backgroundColor: "#6FD9FF",marginLeft: 300 }}
                onClick={() => {
                  LoginChecker();
                }}
                variant="contained"
                color="primary"
              >
                Login
             </Button>
            </div>
          </form>
        </div></InfoCard>
        
      </Content>
    </Page>
  );
}