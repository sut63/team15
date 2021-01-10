import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
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
 }

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
  }, [loading]);

  const EmailthandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };

  const PasswordthandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const LoginChecker = async () => {
    employees.map((item: any) => {
      if ((item.employeeemail == email) && (item.password == password)) {
	    localStorage.setItem("employeedata",JSON.stringify(item.id));
        window.location.href = "http://localhost:3000/Home";
      }
      else {
        setStatus(true);
        setAlert(false);
      }

    })
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.tool}>

<Header
       title="Signin" type=""> 
     </Header>

     <Content>
  <div className={classes.paper}> <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar></div>
     <div className={classes.head}><strong></strong></div>

     <form noValidate autoComplete="off">
     <div><TextField className={classes.textField}
                id="email"
                label="Email"
                variant="outlined"
                type="string"
                size="medium"
                value={email}
                onChange={EmailthandleChange}
              /></div>
      <div><TextField className={classes.textField}
                id="password"
                label="Password"
                variant="outlined"
                type="password"
                size="medium"
                value={password}
                onChange={PasswordthandleChange}
              /></div></form>

            <div> 
            <Button
              onClick={() => {LoginChecker();}}
              type="submit"
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Sign In
            </Button></div>

            {status ? ( 
                      <div className={classes.signin}>
              { alert ? ( 
                    <Link to="/Home" />   ) 
              : ( <Alert variant="outlined" severity="info"> Incorrect Username or Password </Alert> )} </div>
            ) : null}
     </Content>
    </Page>
  );
}