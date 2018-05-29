import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

const API = 'http://localhost:8080/cpu'

const styles = theme => ({
  root: theme.mixins.gutters({
    width: 400,
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})

class Cpu extends Component {
  constructor(props) {
    super(props)

    this.state = {
      programs: []
    }
  }

  componentDidMount() {
    fetch(API)
      .then(response => {
        if (response.ok) {
          response.json()
          .then(data => this.setState({ programs: [...data.programs] }))
        }
      })
  }

  render() {
    const { classes } = this.props
    const { programs } = this.state
    const listPrograms = (programs.length > 1)
      ? programs.map(program =>
        <div key={program}>
          <Link to={`cpu/${program}`}>{program}</Link>
        </div>)
      : <li>No files</li>

    return (
    <div>
      <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h3">
            <h1>CPU</h1>
            <div>{listPrograms}</div>
          </Typography>
      </Paper>
    </div>
    )
  }
}

Cpu.PropTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(Cpu)
