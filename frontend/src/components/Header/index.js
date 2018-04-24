import React from 'react'
import PropTypes from 'prop-types'
import { Link } from 'react-router-dom'
import { withStyles } from 'material-ui/styles'
import AppBar from 'material-ui/AppBar'
import Button from 'material-ui/Button'
import Toolbar from 'material-ui/Toolbar'
import Typography from 'material-ui/Typography'

const styles = {
  root: {
    flexGrow: 1,
  },
  logo: {
    marginRight: 25
  }
}

function Header(props) {
  const { classes } = props;
  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="title" color="inherit" className={classes.logo}>
            GoProfile
          </Typography>
          <Button color="inherit" component={Link} to='/' >Home</Button>
          <Button color="inherit" component={Link} to='/cpu'>CPU</Button>
          <Button color="inherit" component={Link} to='/memory'>Memory</Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}

Header.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Header);
