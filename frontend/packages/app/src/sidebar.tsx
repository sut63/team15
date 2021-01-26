import React, { useState, useEffect } from 'react';
import PermIdentityIcon from '@material-ui/icons/PermIdentity';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import MeetingRoomIcon from '@material-ui/icons/MeetingRoom';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';
import HomeOutlinedIcon from '@material-ui/icons/HomeOutlined';
import HotelIcon from '@material-ui/icons/Hotel';
import { EntEmployee } from 'plugin-welcome/src/api/models/EntEmployee';
import { DefaultApi } from 'plugin-welcome/src/api/apis';

export const AppSidebar = () => {

  const api = new DefaultApi();
  const [employeeid, setEmployee] = useState(Number);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getEmployees = async () => {
      const res = await api.listEmployee();
      setLoading(false);
      setEmployees(res);
    };
    getEmployees();
    const data = localStorage.getItem("employeedata");
    if (data) {
      setEmployee(Number(JSON.parse(data)));
      setLoading(false);
    }
    
  }, [loading]);

  return (

    <Sidebar>
      <SidebarDivider />
      {/* Global nav, not org-specific */}
      {(employeeid) ?
        employees.filter((filter:EntEmployee) => filter.id == employeeid).map((item:EntEmployee) => 
          <SidebarItem icon={PermIdentityIcon} text={item.edges?.jobposition?.positionname} />
        )
        :
        null
      }
      {/* End global nav */}
     { /* <SidebarItem icon={HomeOutlinedIcon} to="#" text="Home" />*/} 
      <SidebarDivider />

          <SidebarItem icon={HotelIcon} to="/SearchRoom" text="ค้นหาข้อมูลห้องพัก"/>
       
      {/* End global nav */}
     { /* <SidebarItem icon={HomeOutlinedIcon} to="#" text="Home" />*/} 
      <SidebarDivider />

      <SidebarSpace />

     

      <SidebarDivider />
      <SidebarThemeToggle />
      {(employeeid) ?
        <SidebarItem icon={MeetingRoomIcon} to="./" text="ออกจากระบบ"
          onClick={() => {
            localStorage.setItem("employeedata", JSON.stringify(null));
            localStorage.setItem("employeelogindata", JSON.stringify(null));
            history.pushState("", "", "./");
            window.location.reload(false);
          }} />
        :
       <SidebarItem icon={HomeOutlinedIcon} to="./" text="Home" />
      }

     
      <SidebarPinButton />
    </Sidebar>
  )
};