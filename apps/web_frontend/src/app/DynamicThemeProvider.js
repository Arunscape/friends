import React from 'react'
import { connect } from 'react-redux'

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import CssBaseline from '@material-ui/core/CssBaseline'

class DynamicThemeProvider extends React.PureComponent {
  render () {
    const theme = createMuiTheme({ palette: this.props.palette })

    return (
      <MuiThemeProvider theme={theme}>
        <CssBaseline />
        { this.props.children }
      </MuiThemeProvider>
    )
  }
}
export default connect(state => ({
  palette: state.settings.palette
}))(DynamicThemeProvider)
