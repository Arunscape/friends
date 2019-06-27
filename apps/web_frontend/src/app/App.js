import React from 'react'

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'

import { Provider } from 'react-redux'
import store, { replaceTokenInStore } from 'app/store'
import { getTokenData } from 'services/security/token'

import ChatPage from 'modules/Chat'
import SettingsPage from 'modules/Settings'
import IndexPage from 'modules/Auth'

import PrivateRoute from './privateRoute'
import DynamicThemeProvider from './DynamicThemeProvider'

export default  () => {
  const Redirector = (props) => <Redirect to='/' />
  store.dispatch(replaceTokenInStore(getTokenData()))
  return (
    <Provider store={store}>
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
