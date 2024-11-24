import './App.css'
import Header from './components/Header'
import Top from './components/Top'
import Sidebar from './components/Sidebar'
import Footer from './components/Footer'

function App() {
  // const [count, setCount] = useState(0)

  return (
    <>
      <div className="app">
        <Header />
        <Top />
        <div className="main-content">
          <Sidebar />
          <div className="content">
            <p> Gontainer. </p>
          </div>
        </div>
        <Footer />
      </div>
    </>
  )
}

export default App
