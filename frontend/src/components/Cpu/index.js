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
        <li key={program}>
        <Link to={`cpu/diagram/${program}`}>{program}</Link>
        </li>)
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
