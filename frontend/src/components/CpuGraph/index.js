import React, { Component } from 'react'

var LineChart = require("react-chartjs").Line
const API = 'http://localhost:8080/cpu/graph/sort'

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
  }

  componentDidMount() {
    (async() => {
      try {
        var response = await fetch(API)
        var data = await response.json()
        this.updateState(data)
      } catch (e) {
        console.log(e)
      }
    })()
  }

  updateState(newData) {
    var chartData = {...this.state.chartData}
    chartData.labels = newData.Time
    chartData.datasets[0].data = newData.Percent
    this.setState({chartData})
  }

  render() {
    const { chartData } = this.state

    return (
      <div>
        <LineChart data={chartData} width="600" height="250"/>
      </div>
    )
  }
}

export default CpuGraph