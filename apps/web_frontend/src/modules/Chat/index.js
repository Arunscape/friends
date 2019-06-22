import React from 'react'
import { Redirect } from 'react-router'

import Message from './Message'
import MessageBar from './SendMessage'

import { isTokenValid } from 'services/security/token'

export default class ChatPage extends React.Component {
  render () {
    if (!isTokenValid()) {
      return <Redirect to='/' />
    }
    return (
      <>
        <Message text={'Hello, World'} />
        <MessageBar />
      </>
    )
  }
}
