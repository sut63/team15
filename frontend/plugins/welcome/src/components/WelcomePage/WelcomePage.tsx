import React, { FC, useState, useEffect } from 'react';
import { Typography, Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  TabbedCard, 
  CardTab ,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import ComponanceRoomDetails from '../RoomDetails';
import ComponanceRecordDeposit from '../RecordDeposit';
import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

const cardContentStyle = { height: 200, width: 500 };
const DormEmployee: FC<{}> = () => {

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);

  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  const [employeeid, setEmployee] = useState(Number);

  
  useEffect(() => { 
        const checkEmployeeLoginData = async () => {
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
        checkEmployeeLoginData();
      
      }, [loading]); 
      console.log("ID พนักงาน")
      console.log(employeeid)
     

  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Dormitory`} subtitle={'ยินดีต้อนรับเข้าสู่หน้า หอพัก สำหรับพนักงานหอพัก'}></Header>
      <Content>
        <TabbedCard title="">
        <CardTab label="เพิ่มข้อมูลห้องพัก">
          <div><ComponanceRoomDetails></ComponanceRoomDetails>
          
          </div>
        </CardTab>
        <CardTab label="รับฝากของ">
          <div><ComponanceRecordDeposit></ComponanceRecordDeposit></div>
        </CardTab>
        <CardTab label="Option 3">
          <div style={cardContentStyle}>Some content 3</div>
        </CardTab>
        <CardTab label="Option 4">
          <div style={cardContentStyle}>Some content 4</div>
        </CardTab>
      </TabbedCard>
      </Content>
    </Page>
  );
};

export default DormEmployee;
