import React, { Component } from 'react'

var LineChart = require("react-chartjs").Line
const API = 'http://localhost:8080/cpu/live/sort'

const chart1 = {
  labels: ['Team1', 'Team2', 'Team3', 'Team4', 'Team5'],
  datasets: [{
    label: 'Team points',
    data: [503, 385, 270, 133, 65],
    backgroundColor: [
      '#4DB6AC',
      '#E57373',
      '#7986CB',
      '#F06292',
      '#E0E0E0'
    ]
  }]
}

const chart2 = {
  labels: ['Team1', 'Team2', 'Team3', 'Team4', 'Team5'],
  datasets: [{
    label: 'Team points 2',
    data: [303, 185, 470, 313, 65],
    backgroundColor: [
      '#4DB6AC',
      '#E57373',
      '#7986CB',
      '#F06292',
      '#E0E0E0'
    ]
  }]
}

const Button = props => (
  <button id="update-chart" onClick={props.handleOnClick}>Update</button>
)

class CpuGraph extends Component {
  constructor(props) {
    super(props)

    this.handleUpdate = this.handleUpdate.bind(this)
    this.update = false

    this.state = {
      chartData: chart1
    }

    console.log(this.state.chartData)
    console.log(this.updated)
  }

  handleUpdate() {
    const chartData = this.updated ? chart1 : chart2
    this.setState({chartData}, () => {
      this.updated = this.updated ? false : true
      console.log(this.state.chartData)
    })
  }

  render() {
    // const { jsondata } = this.state

    return (
      <div>
        <LineChart data={this.state.chartData} width="600" height="250"/>
        <Button handleOnClick={this.handleUpdate} />
      </div>
    )
  }
}

export default CpuGraph