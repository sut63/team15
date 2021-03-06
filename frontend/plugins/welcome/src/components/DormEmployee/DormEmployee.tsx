import React, { FC, useState, useEffect } from 'react';
import { Typography, Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  TabbedCard, 
  CardTab ,
  HeaderLabel
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import ComponanceRoomDetails from '../RoomDetails';
import ComponanceRepairinvoice from '../RecordRepairinvoice';
import ComponanceRecordDeposit from '../RecordDeposit';
import ComponanceRecordBill from '../RecordBill';
import { EntRoomdetail } from '../../api/models/EntRoomdetail';
import ComponanceRecordLease from '../RecordLease';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import ComCleaningroom from '../Cleaningroom';
import Button from '@material-ui/core/Button';

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

  const SearchRoom = () => {
    setStatus(false);
    window.location.href ="http://localhost:3000/Search";
    
  };

  
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
      <Header style={HeaderCustom} title={`Dormitory`} subtitle={'ยินดีต้อนรับเข้าสู่หน้า หอพัก สำหรับพนักงานหอพัก'}>
      <Button style={{ backgroundColor: "#ffffff"}} 
               onClick={() => {
                SearchRoom();
              }} > ค้นหาข้อมูล </Button>
      </Header>
      <Content>
        <TabbedCard title="">

        <CardTab label="เพิ่มข้อมูลห้องพัก">
          <div><ComponanceRoomDetails></ComponanceRoomDetails></div>
        </CardTab>

        <CardTab label="สัญญาเช่า">
        <div><ComponanceRecordLease></ComponanceRecordLease></div>
        </CardTab>

        <CardTab label="ออกใบเสร็จ">
        <div><ComponanceRecordBill></ComponanceRecordBill></div>
        </CardTab>

        <CardTab label="แจ้งซ่อมอุปกรณ์ที่ชำรุด">
          <div><ComponanceRepairinvoice></ComponanceRepairinvoice></div>
        </CardTab>

        <CardTab label="แจ้งทำความสะอาด">
          <div><ComCleaningroom></ComCleaningroom></div>
        </CardTab>

        <CardTab label="รับฝากของ">
          <div><ComponanceRecordDeposit></ComponanceRecordDeposit></div>
        </CardTab>
       
      </TabbedCard>
      </Content>
    </Page>
  );
};

export default DormEmployee;
