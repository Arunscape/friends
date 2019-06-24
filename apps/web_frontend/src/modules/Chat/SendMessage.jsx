import React from 'react'

import Grid from '@material-ui/core/Grid'
import Box from '@material-ui/core/Box'
import Icon from '@material-ui/core/Icon'
import IconButton from '@material-ui/core/IconButton'
import TextField from '@material-ui/core/TextField'

export default class SendMessageBar extends React.Component {
  render () {
    return (
      <Box position='absolute' bottom='0' width='80%'>
        <Grid container spacing={2}>
          <Grid item xs={11}>
            <TextField fullWidth label='Send Message' />
          </Grid>
          <Grid item xs={1}>
            <IconButton>
              <Icon>send</Icon>
            </IconButton>
          </Grid>
        </Grid>
      </Box>
    )
  }
}
