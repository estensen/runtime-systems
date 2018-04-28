import React, { Component } from 'react'

const API = 'http://localhost:8080/cpu'

class Cpu extends Component {
  constructor(props) {
    super(props)

    this.state = {
      filenames: []
    }
  }

  componentDidMount() {
    fetch(API)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.setState({ filenames: [...data.filenames] }))
        }
      })
  }

  render() {
    const { filenames } = this.state
    const listFilenames = (filenames.length > 1)
      ? filenames.map(filename => <li key={filename}>{filename}</li>)
      : <li>No files</li>

    return (
    <div>
      <h1>CPU</h1>
      <ul>{listFilenames}</ul>
    </div>
    )
  }
}

export default Cpu
