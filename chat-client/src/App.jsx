import { useState } from 'react'
import {BrowserRouter, Route, Routes} from 'react-router-dom'


import CreateRoom from "./components/Create"
import Room from "./components/Room"

function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Routes>
        <Route path="/" exact element={<CreateRoom/>}></Route>
        <Route path="/room/:roomID" element={<Room/>}></Route>
      </Routes>
      </BrowserRouter>
      
    </div>
  )
}

export default App
