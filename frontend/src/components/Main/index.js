import React from 'react'
import { Switch, Route } from 'react-router-dom'
import "../../App.css"

import ProgramType from '../ProgramType'
import Diagram from '../Diagram'
import Graph from '../Graph'
import CpuReport from '../CpuReport'
import Home from '../Home'
import Program from '../Program' 
import Memory from '../Memory'

const Main = () => (
  <main>
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route exact path='/:programType' component={ProgramType}/>
      <Route exact path='/:programType/:programName' component={Program}/>
      <Route path='/:programType/:programName/graph' component={Graph}/>
      <Route path='/:programType/:programName/diagram' component={Diagram}/>
      <Route path='/:programType/:programName/report' component={CpuReport}/>
      <Route exact path='/memory' component={Memory}/>
    </Switch>
  </main>
)

export default Main
