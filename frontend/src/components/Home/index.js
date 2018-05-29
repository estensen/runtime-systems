import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

const API = 'http://localhost:8080'

const styles = theme => ({
  root: theme.mixins.gutters({
    width: 400,
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
})

class Home extends Component {
  constructor(props) {
    super(props)

    this.state = {
      apiType: null
    }
  }

  componentDidMount() {
    fetch(API)
      .then(response => response.json())
      .then(data => this.setState({ apiType: data.apiType }))
  }

  render() {
    const { classes } = this.props
    const { apiType } = this.state

    return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <Typography variant="headline" component="h3">
            <h1>Home</h1>
            {apiType}
            Add instructions on how to use the app here
          </Typography>
        </Paper>
      </div>
    )
  }
}

Home.PropTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(Home)
