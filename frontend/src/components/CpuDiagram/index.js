import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

const API = 'http://localhost:8080/cpu/diagram/'

const styles = theme => ({
  root: theme.mixins.gutters({
    margin: '0 auto',
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})

class CpuDiagram extends Component {
  render() {
    const { match: { params: { programName } } } = this.props
    const { classes } = this.props

    return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h1">
            {programName} diagram
          </Typography>
          <Typography component="div">
            <img src={API + programName} width={1400} alt="Program diagram"/>
          </Typography>
        </Paper>
      </div>
    )
  }
}


CpuDiagram.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(CpuDiagram)
