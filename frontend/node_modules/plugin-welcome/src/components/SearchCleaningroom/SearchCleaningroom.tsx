import React, { useState, useEffect } from 'react';
import { createStyles, fade, makeStyles, Theme } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntCleaningroom } from '../../api/models/EntCleaningroom';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import { InputBase, Link, TextField } from '@material-ui/core';
import moment from 'moment';
import SearchIcon from '@material-ui/icons/Search';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    table: {
      minWidth: 650,
    },
    search: {
      position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.black, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.black, 0.25),
      },
      marginLeft: 0,
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(1),
        width: 'auto',
      },
    },
    searchIcon: {
      padding: theme.spacing(0, 2),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    inputRoot: {
      color: 'inherit',
    },
    inputInput: {
      padding: theme.spacing(1, 1, 1, 0),
      // vertical padding + font size from searchIcon
      paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        width: '30ch',
        '&:focus': {
          width: '40ch',
        },
      },
    },
  }),
);



export default function ComponentsTableUser() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [cleaningroom, setCleaningroom] = React.useState<EntCleaningroom[]>([]);
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = useState('');

  const getCleaningroom = async () => {
    const res = await http.listCleaningroom();
    setLoading(false);
    setCleaningroom(res);
  };
  useEffect(() => {
    getCleaningroom();

  }, [loading]);


  const filtercleaningroom = cleaningroom.filter(cleaningroom => {
    return cleaningroom.edges?.roomdetail?.roomnumber?.includes(search)
  })
  function emptyCleaningroom(): any {
    if (filtercleaningroom.length == 0) {
      const clean = <TableRow> <TableCell align="center" colSpan={9}><p>ไม่มีข้อมูลในระบบ</p></TableCell></TableRow>;
      return clean;
    }
  }

  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      
      <Content>
        <ContentHeader title="ข้อมูลแจ้งทำความสะอาด">
          <div className={classes.search} style={{ marginRight: 10 }}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <TextField
              name="search"
              className={classes.inputInput}
              style={{ marginRight: 100 }}
              placeholder="ค้นหาด้วยเลขห้อง"
              value={search}
              onChange={(event) => { setSearch(event.target.value); }}
              inputProps={{ 'aria-label': 'search' }}
            />
          </div>
          <Link component={RouterLink} to="/DormEmployee">
            <Button variant="contained" color="primary">
              ย้อนกลับ
           </Button>
          </Link>
        </ContentHeader>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">เลขที่</TableCell>
                <TableCell align="center">ชื่อพนักงานที่ทำการบันทึก</TableCell>
                <TableCell align="center">ชื่อพนักงานที่รับผิดชอบ</TableCell>
                <TableCell align="center">จำนวนพนักงานในการทำความสะอาด</TableCell>
                <TableCell align="center">เลขห้อง</TableCell>
                <TableCell align="center">ระยะเวลา</TableCell>
                <TableCell align="center">วันที่และเวลาเริ่มทำความสะอาด</TableCell>
                <TableCell align="center">เบอร์โทรศัพท์ที่สามารถติดต่อได้</TableCell>
                <TableCell align="center">เพิ่มเติม</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {emptyCleaningroom()}
              {filtercleaningroom.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                  <TableCell align="center">{item.edges?.cleanername?.cleanername}</TableCell>
                  <TableCell align="center">{item.numofem}</TableCell>
                  <TableCell align="center">{item.edges?.roomdetail?.roomnumber}</TableCell>
                  <TableCell align="center">{item.edges?.lengthtime?.lengthtime}</TableCell>
                  <TableCell align="center">{item.dateandstarttime}</TableCell>
                  <TableCell align="center">{item.phonenumber}</TableCell>
                  <TableCell align="center">{item.note}</TableCell>
                  <TableCell align="center">

                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </Page>

  );
}
