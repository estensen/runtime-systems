import React, { Component } from 'react'
import { Link } from 'react-router-dom'

const API = 'http://localhost:8080/cpu'

class Cpu extends Component {
  constructor(props) {
    super(props)

    this.state = {
      programs: []
    }
  }

  componentDidMount() {
    fetch(API)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.setState({ programs: [...data.programs] }))
        }
      })
  }

  render() {
    const { programs } = this.state
    const listPrograms = (programs.length > 1)
      ? programs.map(program =>
        <div>{program}
          <ul>
            <li key="Diagram">
              <Link to={`cpu/${program}/diagram`}>Diagram</Link>
            </li>
            <li key="Graph">
              <Link to={`cpu/${program}/graph`}>Graph</Link>
            </li>
            <li key="Report">
              <Link to={`cpu/${program}/report`}>Report</Link>
            </li>
          </ul>
        </div>)
      : <li>No files</li>

    return (
    <div>
      <h1>CPU</h1>
      <ul>{listPrograms}</ul>
    </div>
    )
  }
}

export default Cpu
