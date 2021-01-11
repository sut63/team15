import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import BuildIcon from '@material-ui/icons/Build';
import SignOut from '@material-ui/icons/Settings';
import HotelIcon from '@material-ui/icons/Hotel';
import DescriptionIcon from '@material-ui/icons/Description';
import ReceiptIcon from '@material-ui/icons/Receipt';
import CallIcon from '@material-ui/icons/Call';
import AllInboxIcon from '@material-ui/icons/AllInbox';


import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarItem
      icon={HotelIcon}
      to="/Roomdetails"
      text="Room details"
    />
    <SidebarItem
      icon={DescriptionIcon}
      to="/"
      text="ฟอร์ด"
    />
    <SidebarItem
      icon={ReceiptIcon}
      to="/"
      text="พี่อาท"
    />
    <SidebarItem
      icon={BuildIcon}
      to="/"
      text="เอิร์ท"
    />
    <SidebarItem
      icon={CallIcon}
      to="/" //ใส่ตัวแปรของplugin
      text="เทพ"
    />
    <SidebarItem
      icon={AllInboxIcon}
      to="/DepositTable"
      text="เกน"
    />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_out"
      text="Sign Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
