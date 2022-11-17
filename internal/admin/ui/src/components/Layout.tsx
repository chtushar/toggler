import React from "react";
import Header from "./Header";
import Sidebar from "./Sidebar";

const Layout = () => {
    return (
        <div className="h-full flex flex-col">
            <Header />
            <div className="flex flex-1">
                <Sidebar />
            </div>
        </div>
    )
}

export default Layout;