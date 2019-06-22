import React from 'react'

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'

import { Provider } from 'react-redux'
import store from 'app/store'

import ChatPage from 'modules/Chat'
import SettingsPage from 'modules/Settings'
import IndexPage from 'modules/Auth'

import PrivateRoute from './privateRoute'
import DynamicThemeProvider from './DynamicThemeProvider'

import { getTokenData } from 'services/security/token'

function getInitalStateFromToken (tok) {
  console.log(tok)
  return {
    user: {
      name: tok.name,
      email: tok.email,
      pic: tok.picture,
      permissions: tok.permissions
    },
    chat: {
      groups: tok.groups
    }
  }
}

function App () {
  const Redirector = (props) => <Redirect to='/' />
  return (
    <Provider store={store(getInitalStateFromToken(getTokenData()))}>
      <DynamicThemeProvider>
        <BrowserRouter>
          <Switch>
            <PrivateRoute path='/chat/' component={ChatPage} />
            <PrivateRoute path='/settings/' component={SettingsPage} />
            <Route path='/' component={IndexPage} />
            <Route component={Redirector} />
          </Switch>
        </BrowserRouter>
      </DynamicThemeProvider>
    </Provider>
  )
}

export default App
