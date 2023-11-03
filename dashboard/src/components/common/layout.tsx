import React from 'react'

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="w-full h-full hidden md:flex">
      <header></header>
      <main className="flex-1">{children}</main>
    </div>
  )
}

export default Layout
