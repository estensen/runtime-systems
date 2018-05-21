import React from 'react'
import { Switch, Route } from 'react-router-dom'
import "../../App.css"

import Cpu from '../Cpu'
import CpuDiagram from '../CpuDiagram'
import CpuGraph from '../CpuGraph'
import CpuReport from '../CpuReport'
import Home from '../Home'
import Memory from '../Memory'

const Main = () => (
  <main>
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route exact path='/cpu' component={Cpu}/>
      <Route path='/cpu/:programName/graph' component={CpuGraph}/>
      <Route path='/cpu/:programName/diagram' component={CpuDiagram}/>
      <Route path='/cpu/:programName/report' component={CpuReport}/>
      <Route exact path='/memory' component={Memory}/>
    </Switch>
  </main>
)

export default Main
