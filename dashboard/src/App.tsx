import { 
  BrowserRouter, 
  Routes as BrowserRoutes,
}  from 'react-router-dom'
import { QueryClientProvider } from '@tanstack/react-query'
import './App.css'

import { queryClient } from './utils/queryClient'
import Routes from './Routes'

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Routes />  
    </QueryClientProvider>
  )
}

export default App
