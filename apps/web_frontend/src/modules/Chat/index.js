import React from 'react'
import { Redirect } from 'react-router'

import Message from './Message'
import MessageBar from './SendMessage'

import { isTokenValid } from 'services/security/token'
import Header from 'components/atoms/Header'

import Box from '@material-ui/core/Box'
import Container from '@material-ui/core/Container'

export default class ChatPage extends React.Component {
  render () {
    if (!isTokenValid()) {
      return <Redirect to='/' />
    }
    return (
      <Box height='100vh'>
        <Header title='Chat' />
        <Container maxWidth='md' height='100%'>
          <Message text={'Hello, World'} />
          <MessageBar />
        </Container>
      </Box>
    )
  }
}
