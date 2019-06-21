import React from 'react'

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import Purple from '@material-ui/core/colors/purple'
import Orange from '@material-ui/core/colors/orange'

import { Provider } from 'react-redux'
import store from 'app/store'
import 'app/global.css' // Global styles

import ChatPage from 'modules/Chat'
import IndexPage from 'modules/Auth'

const theme = createMuiTheme({
  palette: {
    type: 'dark',
    primary: Purple,
    secondary: Orange
  }
})

function App () {
  const Redirector = (props) => <Redirect to='/' />
  return (
    <MuiThemeProvider theme={theme}>
      <Provider store={store}>
        <BrowserRouter>
          <Switch>
            <Route path='/chat/' component={ChatPage} />
            <Route path='/' component={IndexPage} />
            <Route component={Redirector} />
          </Switch>
        </BrowserRouter>
      </Provider>
    </MuiThemeProvider>
  )
}

export default App
