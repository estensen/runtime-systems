import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

const API = 'http://localhost:8080/cpu/report/sort'

const styles = theme => ({
  root: theme.mixins.gutters({
    margin: '0 auto',
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})

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
    const { classes } = this.props

    return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h1">
            Report
          </Typography>
          <Typography component="pre">
            {this.state.report}
          </Typography>
        </Paper>
      </div>
    )
  }
}

CpuReport.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(CpuReport)
