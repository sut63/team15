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
import ComponanceSearchRoom from '../SearchRoom';
import ComponanceSearchDeposit from '../SearchDeposit';
import ComponanceSearchBill from '../SearchBill';
import ComponanceSearchLease from '../SearchLease';
import ComponanceSearchCleaningroom from '../SearchCleaningroom';
import ComponanceSearchRepairinvoice from '../SearchRepairinvoice';
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
const Search: FC<{}> = () => {

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);

  const [loading, setLoading] = useState(true);



  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Dormitory`} subtitle={'ยินดีต้อนรับเข้าสู่หน้า ค้นหาข้อมูล'}>
   
      </Header>
      <Content>
        <TabbedCard title="">

        <CardTab label="ค้นหาข้อมูลห้องพัก">
          <div><ComponanceSearchRoom></ComponanceSearchRoom></div>
        </CardTab>

		<CardTab label="ค้นหาสัญญาเช่า">
          <div><ComponanceSearchLease></ComponanceSearchLease></div>
        </CardTab>

		<CardTab label="ค้นหาบิล">
          <div><ComponanceSearchBill></ComponanceSearchBill></div>
        </CardTab>

        <CardTab label="ค้นหาการแจ้งซ่อม">
    <div><ComponanceSearchRepairinvoice></ComponanceSearchRepairinvoice></div>
        </CardTab>

		<CardTab label="ค้นหาแจ้งทำความสะอาด">
          <div><ComponanceSearchCleaningroom></ComponanceSearchCleaningroom></div>
        </CardTab>

        <CardTab label="ค้นหารับฝากของ">
          <div><ComponanceSearchDeposit></ComponanceSearchDeposit></div>
        </CardTab>

      </TabbedCard>
      </Content>
    </Page>
  );
};

export default Search;
