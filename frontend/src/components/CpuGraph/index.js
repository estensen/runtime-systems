import React, { Component } from 'react'

const exampleData = {
    labels: ['8:05', '8:10', '8:15', '8:20'],
    datasets: [
      {
        label: 'Example data',
        data: [65, 59, 80, 81],
        fillColor: "rgba(220,220,220,0)",
        strokeColor: 'rgba(50,80,220,1)',
        pointColor: 'rgba(50,80,220,1)',
        pointHighlightFill: '#fff',
      }
    ],
  };

var LineChart = require("react-chartjs").Line;

//const API = 'http://localhost:8080/cpu/livedata/'

class CpuGraph extends Component {
  constructor(props) {
    super(props)
  }

  render() {

    return (
      <div>
        <LineChart data={exampleData} width="600" height="250"/>
      </div>
    )
  }
}

export default CpuGraph