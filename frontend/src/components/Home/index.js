import React, { Component } from 'react'

const API = 'http://localhost:8080'

class Home extends Component {
  constructor(props) {
    super(props)

    this.state = {
      apiType: null
    }
  }

  componentDidMount() {
    fetch(API)
      .then(response => response.json())
      .then(data => this.setState({ apiType: data.apiType }))
  }

  render() {
    const { apiType } = this.state

    return (
      <div>
        <h1>Home</h1>
        {apiType}
      </div>
    )
  }
}

export default Home
