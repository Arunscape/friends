import React from 'react'
import { connect } from 'react-redux'
import './style.css'

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
      email: ''
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
          className={`email-input ${lookLikeEmail(this.state.email) && 'valid-email'}`}
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
    return <div>signup</div>
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

  submitEmail () {
    // TODO: make http request /isuser
  }
}

function lookLikeEmail (email) {
  return email.match(/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/)
}

export default connect((state) => ({
  user: state.user
}), {
})(IndexPage)
