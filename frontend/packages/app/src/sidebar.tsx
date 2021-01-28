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
import SearchIcon from '@material-ui/icons/Search';
import { EntEmployee } from 'plugin-welcome/src/api/models/EntEmployee';
import { DefaultApi } from 'plugin-welcome/src/api/apis';

export const AppSidebar = () => {

  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
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

  const Search = () => {
    setStatus(false);
    window.location.href ="http://localhost:3000/Search";
    
  };

  return (

    <Sidebar>
      {/* Global nav, not org-specific */}
      {(employeeid) ?
        employees.filter((filter:EntEmployee) => filter.id == employeeid).map((item:EntEmployee) => 
          <SidebarItem to="/DormEmployee" icon={PermIdentityIcon} text={item.edges?.jobposition?.positionname} />
        )
        :
        
        <SidebarItem icon={SearchIcon} text="ค้นหาข้อมูล"
          onClick={() => {
            Search();
          }}
          
          />
      }
      {/* End global nav */}
     { /* <SidebarItem icon={HomeOutlinedIcon} to="#" text="Home" />*/} 
      

         
       
      {/* End global nav */}
     { /* <SidebarItem icon={HomeOutlinedIcon} to="#" text="Home" />*/} 
      

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