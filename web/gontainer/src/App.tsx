import './App.css'
import Header from './components/Header'
import Top from './components/Top'
import Sidebar from './components/Sidebar'
import Footer from './components/Footer'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Containers from './pages/Containers'
import Images from './pages/Images'
import Home from './pages/Home'

function App() {
  // const [count, setCount] = useState(0)

  return (
    <>
      <BrowserRouter>
        <div className="app">
          <Header />
          <Top />
          <Sidebar />
          <div className="content">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/containers" element={<Containers />} />
              <Route path="/images" element={<Images />} />
            </Routes>
          </div>
          <Footer />
        </div>
      </BrowserRouter>
    </>
  )
}

export default App
