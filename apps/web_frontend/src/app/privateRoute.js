import React from 'react'
import { Route, Redirect } from 'react-router-dom'
import { isTokenValid } from 'services/security/token'

export default class PrivateRoute extends React.PureComponent {
  render () {
    if (isTokenValid()) {
      return <Route {...this.props} />
    } else {
      return <Redirect to='/' />
    }
  }
}
