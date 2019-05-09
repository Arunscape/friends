import React from 'react'
import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import { Provider } from 'react-redux'
import store from 'store'
import './global.css' // Global styles

import ChatPage from 'modules/Chat'
import IndexPage from 'modules/Auth'

function App () {
  const Redirector = (props) => <Redirect to='/' />
  return (
    <Provider store={store}>
      <BrowserRouter>
        <Switch>
          <Route path='/chat/' component={ChatPage} />
          <Route path='/' component={IndexPage} />
          <Route component={Redirector} />
        </Switch>
      </BrowserRouter>
    </Provider>
  )
}

export default App
