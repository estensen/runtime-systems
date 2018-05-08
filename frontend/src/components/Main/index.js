import React from 'react'
import { Switch, Route } from 'react-router-dom'
import "../../App.css"

import Cpu from '../Cpu'
import CpuReport from '../CpuReport'
import Home from '../Home'
import Memory from '../Memory'

const Main = () => (
  <main>
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route exact path='/cpu' component={Cpu}/>
      <Route path='/cpu/:filename' component={CpuReport}/>
      <Route path='/memory' component={Memory}/>
    </Switch>
  </main>
)

export default Main
