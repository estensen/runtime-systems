import React, { Component } from 'react'

const API = 'http://localhost:8080/cpu/diagram/'

class CpuDiagram extends Component {
  render() {
    const { match: { params: { programName } } } = this.props

    return (
      <div>
        <h1>{programName}</h1>
        <img src={API + programName} alt="Program diagram"/>
      </div>
    )
  }
}

export default CpuDiagram
