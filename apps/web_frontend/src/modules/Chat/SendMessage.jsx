import React from 'react'

import Icon from '@material-ui/core/Icon'
import IconButton from '@material-ui/core/IconButton'
import TextField from '@material-ui/core/TextField'

export default class SendMessageBar extends React.Component {
  render () {
    return (
      <div class='hor'>
        <TextField />
        <IconButton>
          <Icon>send</Icon>
        </IconButton>
      </div>
    )
  }
}
