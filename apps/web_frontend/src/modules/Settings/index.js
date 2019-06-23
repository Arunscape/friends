import React from 'react'
import { connect } from 'react-redux'

import { saveSettings } from 'modules/Settings/actions'

import Button from '@material-ui/core/Button'
import Radio from '@material-ui/core/Radio'
import Grid from '@material-ui/core/Grid'
import Container from '@material-ui/core/Container'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

import * as Colors from '@material-ui/core/colors'
const colorKeys = Object.keys(Colors)

class SettingsPage extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      type: this.props.palette.type,
      primary: this.props.palette.primary,
      secondary: this.props.palette.secondary
    }
  }

  render () {
    return (
      <Container>
        <Grid container spacing={3}>
          <Grid item xs={12} sm={6} align='center'>
            <ColorPicker title='Primary Color' setColor={color => this.setState({ primary: color })} />
          </Grid>
          <Grid item xs={12} sm={6} align='center'>
            <ColorPicker title='Secondary Color' setColor={color => this.setState({ secondary: color })} />
          </Grid>
          <Grid item xs={12} align='center'>
            Dark:
            <Radio
              checked={this.state.type === 'dark'}
              onChange={ev => this.setState({ type: ev.target.value })}
              value='dark'
            />
            Light:
            <Radio
              checked={this.state.type === 'light'}
              onChange={ev => this.setState({ type: ev.target.value })}
              value='light'
            />
          </Grid>
          <Grid item xs={12} align='center'>
            <Button variant='contained' color='primary'
              onClick={() => this.props.saveSettings(this.state.type, this.state.primary, this.state.secondary)}> Save </Button>
          </Grid>
        </Grid>
      </Container>
    )
  }
}

class ColorPicker extends React.Component {
  render () {
    const boxSize = 45
    const columns = 4
    const numToPx = num => `${num}px`
    return (
      <Paper>
        <Typography> { this.props.title } </Typography>
        <Grid container spacing={0} style={{ width: numToPx(columns * boxSize) }}>
          { colorKeys.filter(key =>
            Colors[key]['500'] && Colors[key]['A100']
          ).map(key =>
            <Grid key={key} item xs={12 / columns}>
              <div onClick={() => this.props.setColor(Colors[key])}
                style={{ width: numToPx(boxSize), height: numToPx(boxSize), backgroundColor: Colors[key]['500'] }} />
            </Grid>
          )}
        </Grid>
      </Paper>
    )
  }
}

export default connect((state) => ({
  palette: state.settings.palette
}), (dispatch) => ({
  saveSettings: (...data) => saveSettings(dispatch, ...data)
}))(SettingsPage)
