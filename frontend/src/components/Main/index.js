import React from 'react'
import { Switch, Route } from 'react-router-dom'

import Cpu from '../Cpu'
import CpuDiagram from '../CpuDiagram'
import Home from '../Home'
import Memory from '../Memory'

const Main = () => (
  <main>
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route exact path='/cpu' component={Cpu}/>
      <Route path='/cpu/diagram/:programName' component={CpuDiagram}/>
      <Route path='/memory' component={Memory}/>
    </Switch>
  </main>
)

export default Main
