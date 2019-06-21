import React from 'react'

import Message from './Message'
import MessageBar from './SendMessage'

export default class ChatPage extends React.Component {
  render () {
    return (
      <>
        <Message text={'Hello, World'} />
        <MessageBar />
      </>
    )
  }
}
