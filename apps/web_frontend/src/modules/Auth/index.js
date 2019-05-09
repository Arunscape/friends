import React from 'react'
import { connect } from 'react-redux'
import './style.css'

import { checkUser, signup } from './actions'

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
      <div className='welcome-screen drift ver'>
        <div className='ver cent'>
          <h1>Welcome, Friend!</h1>
          <p>I need your email to get started</p>
        </div>
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
      </div>
    )
  }

  renderSignup () {
    return (
      <div className='welcome-screen drift ver'>
        <div className='ver cent'>
          <h1>Welcome, Friend!</h1>
          <p>You seem to be missing an account, let's get you set up</p>
        </div>
        <input className={`name-input`} value={this.state.name}
          onChange={(ev) => this.setState({ name: ev.target.value })}
          type='text' placeholder='name' />
        <input className={`name-input`} value={this.state.pic}
          onChange={(ev) => this.setState({ pic: ev.target.value })}
          type='text' placeholder='profile pic url' />
        <input className={`submit-signup ${this.state.name &&
           this.state.pic && 'show-submit-button'}`}
          type='button' value='Get Started'
          onClick={() => this.submitSignup()} />
      </div>
    )
  }

  renderSignin () {
    return <div>verifiy email</div>
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
  }

  async submitSignup () {
    await this.props.signup(this.state.email)
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
  signup: (...data) => signup(dispatch, ...data)
}))(IndexPage)
