import React from 'react'
import { connect } from 'react-redux'

import { saveSettings } from 'modules/Settings/actions'

import Button from '@material-ui/core/Button'

import Red from '@material-ui/core/colors/red'
import Blue from '@material-ui/core/colors/blue'

class SettingsPage extends React.Component {
  render () {
    console.log(this.props)
    return (
      <>
        <button onClick={() => this.props.saveSettings('light', Red, Blue)}>Switch Colors</button>
        <Button variant='contained' color='primary' > Primary </Button>
        <Button variant='contained' color='secondary' > Secondary </Button>
      </>
    )
  }
}

export default connect((state) => ({
  user: state.user
}), (dispatch) => ({
  saveSettings: (...data) => saveSettings(dispatch, ...data)
}))(SettingsPage)
