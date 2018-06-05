import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

const styles = theme => ({
  root: theme.mixins.gutters({
    width: 400,
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})

class Home extends Component {
  render() {
    const { classes } = this.props

    return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h3">
            Instructions:
          </Typography>
          <Typography component="p">
            <ol>
              <li>
                Choose between profiling the CPU and the Memory, by following the navigation bar
              </li>
              <li>
                After navigating to one of the types, you get a list of the programs available for profiling
              </li>
              <li>
                Select which program you want to profile
              </li>
              <li>
                Click on the "run profiling" button to profile
              </li>
                If it exists a profile for the selected program, you will have options for visualizing the last profile. If not, 
                you need to run profiling first.
                
              <li>
                Then select how you want to see the profiling. Current options: Diagram, Graph and Report
              </li>
            </ol>
          </Typography>
        </Paper>
      </div>
    )
  }
}

Home.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(Home)
