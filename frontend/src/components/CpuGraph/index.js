import React, { Component } from 'react'

var LineChart = require("react-chartjs").Line
const API = 'http://localhost:8080/cpu/live/sort'

const chart = {
  labels: ['8:05', '8:10', '8:15', '8:20', '8:25'],
  datasets: [{
    label: 'CPU Utilization',
    data: [303, 185, 470, 313, 65],
    fillColor: "rgba(220,220,220,0)",
    strokeColor: 'rgba(50,80,220,1)',
    pointColor: 'rgba(50,80,220,1)',
    pointHighlightFill: '#fff',
  }]
}


class CpuGraph extends Component {
  constructor(props) {
    super(props)

    this.state = {
      chartData: chart
    }

    console.log(this.state.chartData)
    console.log(this.updated)
  }

  componentDidMount() {
    fetch(API)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.updateState(data))
        }
      })
  }

  updateState(newData) {
    this.setState(prevState => ({
      ...prevState,
      labels: newData.time,
      datasets: {
        ...prevState.datasets,
        data: newData.percent
      }
    }))
  }

  render() {
    return (
      <div>
        <LineChart data={this.state.chartData} width="600" height="250"/>
      </div>
    )
  }
}

export default CpuGraph