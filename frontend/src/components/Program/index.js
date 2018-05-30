import React from 'react'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'
import Button from '@material-ui/core/Button'

const styles = theme => ({
  root: theme.mixins.gutters({
    width: 400,
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})


class CpuProgram extends React.Component {
  constructor(props){
    super(props)
    this.toggle = this.toggle.bind(this);
    this.state = {
      isProfiled: false,
      programName: this.props.programName
    }
  }

  toggle() {
    if (this.state.isProfiled) {
      this.setState({isProfiled: false})
    } else {
      this.runProfiling
    }
  }

  runProfiling(){
    fetch('http://localhost:8080/cpu/runprofiling/wordSearch')
    .then((result) => {
      this.setState({isProfiled: true})
    })
  }
  
  render () {
    const isProfiled = this.state.isProfiled;
    const { match: { params: { programName } } } = this.props
    const { classes } = this.props

    const profilingButton = isProfiled ? (
      <div>
        <Button variant="outlined" size="small" className={classes.button} onClick={this.toggle} disabled="true">
            Run Profiling
        </Button>
        <Typography component="ul" style={{ margin: 10 }}>
          <li key="Diagram" style={{ margin: 5 }}>
            <Link to={`${programName}/diagram`}>Diagram</Link>
          </li>
          <li key="Graph" style={{ margin: 5 }}>
            <Link to={`${programName}/graph`}>Graph</Link>
          </li>
          <li key="Report" style={{ margin: 5 }}>
            <Link to={`${programName}/report`}>Report</Link>
          </li>
        </Typography> 
      </div>
    ) : (
      <Button variant="outlined" size="small" className={classes.button} onClick={this.toggle}>
          Run Profiling
      </Button>
    );
  
  return (
    <div>
      <Paper className={classes.root} elevation={4}>
        <Typography variant="headline" component="h1">
          CPU: {programName}
        </Typography>
        {profilingButton}
      </Paper>
    </div>
  )
}
}


CpuProgram.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(CpuProgram)
