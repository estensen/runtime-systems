import React, { Component } from 'react'
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

class ProgramType extends Component {
  constructor(props) {
    super(props)
    
    this.state = {
      programs: [],
    }
  }
  
  componentDidMount() {
    const { match: { params: { programType } } } = this.props
    const API = `http://localhost:8080/programs/${programType}`

    fetch(API)
    .then(response => {
      if (response.ok) {
        response.json()
        .then(data => this.setState({ programs: [...data.programs] }))
      }
    })
  }
  
  render() {
    const { match: { params: { programType } } } = this.props
    const { classes } = this.props
    const { programs } = this.state
    const listPrograms = (programs.length > 1)
      ? programs.map(program =>
        <div key={program}>
          <Link to={`${programType}/${program}`}>{program}</Link>
        </div>)
      : <li>No files</li>

    return (
    <div>
      <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h1">
            {programType}
          </Typography>
          <Typography component="div">
            {listPrograms}
          </Typography>
      </Paper>
    </div>
    )
  }
}

ProgramType.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(ProgramType)
