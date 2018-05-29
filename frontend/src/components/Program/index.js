import React from 'react'
import { Link } from 'react-router-dom'
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

function CpuProgram(props) {
  const { match: { params: { programName } } } = props
  const { classes } = props

  return (
    <div>
      <Paper className={classes.root} elevation={4}>
        <Typography variant="headline" component="h1">
          CPU {programName}
        </Typography>
        <Typography component="ul">
            <li key="Diagram">
              <Link to={`${programName}/diagram`}>Diagram</Link>
            </li>
            <li key="Graph">
              <Link to={`${programName}/graph`}>Graph</Link>
            </li>
            <li key="Report">
              <Link to={`${programName}/report`}>Report</Link>
            </li>
        </Typography>
      </Paper>
    </div>
  )
}

CpuProgram.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(CpuProgram)
