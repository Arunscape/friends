import React from 'react'
import { Redirect } from 'react-router'

import Message from './Message'
import MessageBar from './SendMessage'

import { isTokenValid } from 'services/security/token'
import Header from 'components/atoms/Header'

import Box from '@material-ui/core/Box'
import Button from '@material-ui/core/Button'
import Container from '@material-ui/core/Container'

import { signout } from 'modules/Auth/actions'

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
          <Button variant='contained' color='primary' onClick={() => signout()}> Logout </Button>
          <MessageBar />
        </Container>
      </Box>
    )
  }
}
