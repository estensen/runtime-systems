import React from 'react'
import { Switch, Route } from 'react-router-dom'

import Cpu from '../Cpu'
import Home from '../Home'
import Memory from '../Memory'

const Main = () => (
  <main>
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route path='/cpu' component={Cpu}/>
      <Route path='/memory' component={Memory}/>
    </Switch>
  </main>
)

export default Main
