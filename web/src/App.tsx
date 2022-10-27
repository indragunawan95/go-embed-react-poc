import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'

function App() {
  const [data, setData] = useState("")

  useEffect(() =>{
    const fetchData = async () => {
      const response = await fetch("http://localhost:8080/api")
      const data = await response.text()
      setData(data)
    }

    fetchData().catch(err => console.log(err))
  }, [])

  return (
    <div className="App">
      {data}
      <h1>Rambo</h1>
      <h1>Rambutan</h1>
    </div>
  )
}

export default App
