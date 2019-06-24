import React from 'react'
import { connect } from 'react-redux'

import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'

import { upgrade } from 'modules/Auth/actions'

const UPGRADE_TIME = 1000 * 6
class Header extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      timer: null
    }
    this.upgrader()
  }

  componentWillUnmount () {
    clearTimeout(this.state.timer)
  }

  upgrader () {
    this.props.upgrade()
    setTimeout(() => this.upgrader(), UPGRADE_TIME)
  }

  render () {
    return (
      <AppBar position='static'>
        <Toolbar>
          <Typography variant='h6' noWrap>
            { this.props.title }
          </Typography>
        </Toolbar>
      </AppBar>
    )
  }
}

export default connect((state) => ({ }), (dispatch) => ({
  upgrade: (...data) => upgrade(dispatch, ...data)
}))(Header)
