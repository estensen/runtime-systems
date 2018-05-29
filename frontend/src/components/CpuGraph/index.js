import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

var LineChart = require("react-chartjs").Line
const API = 'http://localhost:8080/cpu/graph/sort'

const styles = theme => ({
  root: theme.mixins.gutters({
    margin: '0 auto',
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})

const chart = {
  labels: ['0'],
  datasets: [{
    label: 'CPU Utilization',
    data: [0],
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
    const { match: { params: { programName } } } = this.props
    const { classes } = this.props
    const { chartData } = this.state

    return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h1">
            {programName} % CPU usage
          </Typography>
          <Typography component="div">
            <LineChart data={chartData} width="600" height="250"/>
          </Typography>
        </Paper>
      </div>
    )
  }
}

CpuGraph.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(CpuGraph)
