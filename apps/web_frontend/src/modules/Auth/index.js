import React from 'react'
import { connect } from 'react-redux'
import './style.css'

import { checkUser, signup, signin, upgrade, signout } from './actions'

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
    this.poll = setInterval(() => this.pollForUpgrade(), 3000)
  }

  async pollForUpgrade () {
    if (this.state.state === STATE.SIGNIN) {
      if (await this.props.upgrade()) {
        clearTimeout(this.poll)
        this.props.history.push('/chat')
      }
    }
  }

  render () {
    switch (this.state.state) {
      case STATE.START:
        return this.renderEmailGetter()
      case STATE.SIGNUP:
        return this.renderSignup()
      case STATE.SIGNIN:
        return this.renderSignin()
      case STATE.DONE:
        return this.renderDone()
      default:
        return this.renderError()
    }
  }

  renderEmailGetter () {
    return (
      <WelcomeBox subHeader='I need your email to get started'>
        <input
          className={`email-input ${lookLikeEmail(this.state.email) &&
              'valid-email'}`}
          value={this.state.email}
          onChange={(ev) => this.setState({ email: ev.target.value })}
          type='text' placeholder='email' />
        <input
          className={`submit-email ${lookLikeEmail(this.state.email) &&
            'show-submit-button'}`}
          type='button' value='Get Started'
          onClick={() => this.submitEmail()} />
      </WelcomeBox>
    )
  }

  renderSignup () {
    return (
      <WelcomeBox
        subHeader={'You seem to be missing an account, let\'s get you set up'}>
        <input className={`name-input`} value={this.state.name}
          onChange={(ev) => this.setState({ name: ev.target.value })}
          type='text' placeholder='name' />
        <input className={`name-input`} value={this.state.pic}
          onChange={(ev) => this.setState({ pic: ev.target.value })}
          type='text' placeholder='profile pic url' />
        <input className={`submit-signup ${this.state.name && this.state.pic && 'show-submit-button'}`}
          type='button' value='Get Started'
          onClick={() => this.submitSignup()} />
      </WelcomeBox>
    )
  }

  renderSignin () {
    return <div>Verifying Email</div>
  }

  renderDone () {
    return <div>error</div>
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

class WelcomeBox extends React.PureComponent {
  render () {
    return (
      <div className='drift container ver'>
        <div className='welcome-screen drift ver'>
          <div className='ver cent'>
            <h1>Welcome, Friend!</h1>
            <p> { this.props.subHeader } </p>
          </div>
          { this.props.children }
        </div>
        <div />
      </div>
    )
  }
}
