import React, { Component } from 'react'

const API = 'http://localhost:8080/cpu/'

class CpuReport extends Component {
  constructor(props) {
    super(props)

    this.state = {
      reportLength: null
    }
  }
  componentDidMount() {    
    const { match: { params } } = this.props

    fetch(API + params.filename)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.setState({ reportLength: data.reportLength }))
        }
      })
  }

  render() {
    const { match: { params: { filename } } } = this.props
    const { reportLength } = this.state

    return (
      <div>
        <h1>Length of report {filename}</h1>
        {reportLength}
      </div>
    )
  }
}

export default CpuReport
