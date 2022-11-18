import React from "react";
import Header from "./Header";
import Sidebar from "./Sidebar";
import { Outlet } from 'react-router-dom';

const Layout = () => {
    return (
        <div className="h-full flex flex-col">
            <Header />
            <div className="flex flex-1">
                <Sidebar />
                <div>
                    <Outlet />
                </div>
            </div>
        </div>
    )
}

export default Layout;