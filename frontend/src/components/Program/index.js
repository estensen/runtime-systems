import React, { Component } from 'react'
import { Link } from 'react-router-dom'

class CpuProgram extends Component {
  render() {

    const { match: { params: { programName } } } = this.props

    return (
    <div>
      <h1>CPU</h1>
      <h2>{programName}</h2>
      <div>
          <ul>
            <li key="Diagram">
              <Link to={`${programName}/diagram`}>Diagram</Link>
            </li>
            <li key="Graph">
              <Link to={`${programName}/graph`}>Graph</Link>
            </li>
            <li key="Report">
              <Link to={`${programName}/report`}>Report</Link>
            </li>
          </ul>
        </div>
    </div>
    )
  }
}

export default CpuProgram
