import React, {useState} from 'react'
import { connect } from 'react-redux'

import Header from 'components/atoms/Header'
import { saveSettings } from 'modules/Settings/actions'

import Button from '@material-ui/core/Button'
import Box from '@material-ui/core/Box'
import Radio from '@material-ui/core/Radio'
import Grid from '@material-ui/core/Grid'
import Container from '@material-ui/core/Container'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

import * as Colors from '@material-ui/core/colors'
const colorKeys = Object.keys(Colors)

const SettingsPage = props => {


  const [type, setType] = useState(props.palette.type);
  const [primary, setPrimary] = useState(props.palette.primary)
  const [secondary, setSecondary] = useState(props.palette.secondary)

    return (
      <div>
        <Header title='Settings' />
        <Box pt={3}>
          <Container>
            <Grid container spacing={3}>
              <Grid item xs={12} sm={6} align='center'>
                <ColorPicker title='Primary Color' setColor={color => setPrimary(color)} />
              </Grid>
              <Grid item xs={12} sm={6} align='center'>
                <ColorPicker title='Secondary Color' setColor={color => setSecondary(color)} />
              </Grid>
              <Grid item xs={12} align='center'>
                Dark:
                <Radio
                  checked={type === 'dark'}
                  onChange={ev => setType(ev.target.value)}
                  value='dark'
                />
                Light:
                <Radio
                  checked={type === 'light'}
                  onChange={ev => setType(ev.target.value)}
                  value='light'
                />
              </Grid>
              <Grid item xs={12} align='center'>
                <Button variant='contained' color='primary'
                  onClick={() => props.saveSettings(type, primary, secondary)}> Save </Button>
              </Grid>
            </Grid>
          </Container>
        </Box>
      </div>
    )
  }


const ColorPicker = () => {
 
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


export default connect((state) => ({
  palette: state.settings.palette
}), (dispatch) => ({
  saveSettings: (...data) => saveSettings(dispatch, ...data)
}))(SettingsPage)
