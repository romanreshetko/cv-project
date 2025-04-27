import { BrowserRouter, Route, Routes } from 'react-router'
import './App.css'
import MainPage from './components/MainPage'
import SignUp from './components/SignUp'
import LogIn from './components/LogIn'
import FileUpload from './components/FileUpload'
import Protected from './components/Protected'
import HhUpload from './components/hhUpload'

function App() {
  

  return (
     <BrowserRouter>
        <Routes>
          <Route path="/" element={<MainPage />} />
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="/log-in" element={<LogIn />} />
          <Route path="/upload" element={<Protected><FileUpload /></Protected>} />
          <Route path="/upload-hh" element={<Protected><HhUpload /></Protected>} />
        </Routes>
     </BrowserRouter>
  )
}

export default App
