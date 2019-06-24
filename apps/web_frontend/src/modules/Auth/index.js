import React from 'react'
import { connect } from 'react-redux'

import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography'
import Grid from '@material-ui/core/Grid'
import Container from '@material-ui/core/Container'
import Paper from '@material-ui/core/Paper'
import LinearProgress from '@material-ui/core/LinearProgress'

import { checkUser, signup, signin, upgrade, signout } from './actions'
import { isTokenValid, isTempTokenValid } from 'services/security/token'

const STATE = {
  START: 'start',
  SIGNUP: 'signup',
  SIGNIN: 'signin',
  DONE: 'done'
}

class IndexPage extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      state: STATE.START,
      email: '',
      name: '',
      pic: ''
    }

    // Initial state
    if (isTempTokenValid()) {
      this.state.state = STATE.SIGNIN
      if (isTokenValid()) {
        this.props.upgrade()
        this.goToChatPage()
      }
    }
  }

  componentDidMount () {
    this.poll = setInterval(() => this.pollForUpgrade(), 3000)
    this.pollForUpgrade()
  }

  async pollForUpgrade () {
    if (this.state.state === STATE.SIGNIN) {
      await this.props.upgrade()
      if (isTokenValid()) {
        this.goToChatPage()
      }
    }
  }

  goToChatPage () {
    clearTimeout(this.poll)
    this.props.history.push('/chat')
  }

  render () {
    switch (this.state.state) {
      case STATE.START:
        return this.renderEmailGetter()
      case STATE.SIGNUP:
        return this.renderSignup()
      case STATE.SIGNIN:
        return this.renderSignin()
      default:
        return this.renderError()
    }
  }

  renderEmailGetter () {
    return (
      <Container component='main' maxWidth='md'>
        <Paper style={{ padding: '10%' }}>
          <Grid container spacing={5}>
            <Grid item xs={12} align='center'>
              <Typography component='h2' variant='h3'> Welcome, Friend! </Typography>
            </Grid>
            <Grid item xs={12} align='center'>
              <Typography component='h3' variant='h5'>I need your email to get started</Typography>
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                value={this.state.email}
                onChange={(ev) => this.setState({ email: ev.target.value })}
                fullWidth
                type='email'
                autoComplete='email'
                variant='outlined'
                label='Email' />
            </Grid>
            <Grid item xs={12} md={6} align='center'>
              <Button
                variant='contained'
                color='primary'
                disabled={!lookLikeEmail(this.state.email)}
                onClick={() => this.submitEmail()} >
                Get Started
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    )
  }

  renderSignup () {
    return (
      <Container component='main' maxWidth='md'>
        <Paper style={{ padding: '10%' }}>
          <Grid container spacing={5}>
            <Grid item xs={12} align='center'>
              <Typography component='h2' variant='h3'> Welcome, Friend! </Typography>
            </Grid>
            <Grid item xs={12} align='center'>
              <Typography component='h3' variant='h5'> You seem to be missing an account, let's get you set up </Typography>
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                value={this.state.name}
                onChange={(ev) => this.setState({ name: ev.target.value })}
                fullWidth
                autoComplete='name'
                variant='outlined'
                label='Name' />
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                value={this.state.pic}
                onChange={(ev) => this.setState({ pic: ev.target.value })}
                fullWidth
                variant='outlined'
                label='Profile Picture URL' />
            </Grid>
            <Grid item xs={12} align='center'>
              <Button
                variant='contained'
                color='primary'
                disabled={!(this.state.name && this.state.pic)}
                onClick={() => this.submitSignup()} >
                Sign up
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    )
  }

  renderSignin () {
    return (
      <Container component='main' maxWidth='md'>
        <Paper style={{ padding: '10%' }}>
          <Grid container spacing={5}>
            <Grid item xs={12} align='center'>
              <Typography component='h2' variant='h3'> Waiting for email validation. </Typography>
            </Grid>
            <Grid item xs={12} align='center'>
              <Typography component='h3' variant='h5'> Please check your inbox! </Typography>
            </Grid>
            <Grid item xs={12} align='center'>
              <LinearProgress color='secondary' />
            </Grid>
          </Grid>
        </Paper>
      </Container>
    )
  }

  renderError () {
    return <div>error</div>
  }

  async submitEmail () {
    const res = await this.props.checkUser(this.state.email)
    this.setState({ state: res ? STATE.SIGNIN : STATE.SIGNUP })
    if (res) {
      this.submitSignin()
    }
  }

  async submitSignup () {
    await this.props.signup(this.state.email, this.state.name, this.state.pic)
    this.setState({ state: STATE.SIGNIN })
  }

  async submitSignin () {
    await this.props.signin(this.state.email)
    this.setState({ state: STATE.SIGNIN })
  }
}

function lookLikeEmail (email) {
  return email.match(/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/)
}

export default connect((state) => ({
  user: state.user
}), (dispatch) => ({
  checkUser: (...data) => checkUser(dispatch, ...data),
  upgrade: (...data) => upgrade(dispatch, ...data),
  signin: (...data) => signin(dispatch, ...data),
  signup: (...data) => signup(dispatch, ...data),
  signout: (...data) => signout(dispatch, ...data)
}))(IndexPage)
