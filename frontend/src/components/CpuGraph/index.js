import React, { Component } from 'react'

const API = 'http://localhost:8080/cpu/live/sort'

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

class CpuGraph extends Component {

  constructor(props) {
    super(props)

    this.state = {
      labels: [],
      datasets: [
        {
          data: [],
          fillColor: "rgba(220,220,220,0)",
          strokeColor: 'rgba(50,80,220,1)',
          pointColor: 'rgba(50,80,220,1)',
          pointHighlightFill: '#fff',
        }
      ],
      jsondata: [],
    }
  }

  
  componentDidMount() {    
    
    fetch(API)
    .then(response => {
      if (response.ok) {
        response.json()
        .then(jsondata => this.updateData(jsondata[0], jsondata[1]))
      }
    })
  }
  
  updateData(label, data){
    //this.setState(this.state.labels.push(label))
    this.setState({labels: label})
    this.setState({datasets: [{ ...this.state.datasets }], data: data})
  }

  render() {
    const { jsondata } = this.state

    return (
      <div>
        {jsondata}
        <LineChart data={this.state} width="600" height="250"/>
      </div>
    )
  }
}

export default CpuGraph