import { Outlet } from 'react-router-dom'

const New = () => {
  return (
    <div className="flex flex-col h-full w-full">
      <div className="border-b h-10 border-solid border-muted-background"></div>
      <div className="flex-1 flex items-center justify-center bg-muted">
        <Outlet />
      </div>
    </div>
  )
}

export default New
