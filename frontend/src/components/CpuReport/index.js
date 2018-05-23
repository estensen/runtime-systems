import React, { Component } from 'react'

const API = 'http://localhost:8080/cpu/report/sort'

class CpuReport extends Component {
  constructor(props) {
    super(props)

    this.state = {
      report: []
    }
  }

  componentDidMount() {
    fetch(API)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.setState({ report: data }))
        }
      })
  }

  render() {
    return (
    <div>
      <h1>Report</h1>
      <div>{this.state.report}</div>
    </div>
    )
  }
}

export default CpuReport
