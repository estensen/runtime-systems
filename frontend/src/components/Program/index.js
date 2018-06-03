import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button'
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

class Program extends Component {
  constructor(props) {
    super(props)
    
    this.state = {
      isProfiled: false,
    }
  }
  
  componentDidMount() {
    const { match: { params: { programName, programType } } } = this.props
    const API= `http://localhost:8080/checkProfiling/${programType}/${programName}`
    fetch(API)
    .then(response => {
      if (response.ok) {
        response.json()
        .then(data => this.setState({ isProfiled: data.profileExists }))
      }
    })
  }
  
  profile = () => {
    const { match: { params: { programName, programType } } } = this.props
    const API = `http://localhost:8080/runprofiling/${programType}/${programName}`  // Hardcoded
    fetch(API)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.setState({ isProfiled: data.isProfiled }))
        }
      })
  }

  render() {
    const { isProfiled } = this.state
    const { match: { params: { programName, programType } } } = this.props
    const { classes } = this.props

    const profileOptions = isProfiled
      ? <Typography component="ul">
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
      : <div></div>

      return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h1">
            {programType} {programName}
          </Typography>
          <Button variant="outlined" size="small" className={classes.button} onClick={this.profile}>
            Run Profiling
          </Button>
          {profileOptions}
        </Paper>
      </div>
    )
  }
  }


Program.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(Program)
