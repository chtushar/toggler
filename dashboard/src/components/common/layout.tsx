import React from 'react'
import Sidebar from './sidebar'

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="w-full h-full hidden md:flex">
      <Sidebar />
      <main className="flex-1">{children}</main>
    </div>
  )
}

export default Layout
