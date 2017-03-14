// @flow weak

import React, { Component } from 'react'
import { Redirect } from 'react-router'

import { FKApiClient } from '../../api/api';
import { errorsFor } from '../../common/util';
import type { ErrorMap } from '../../common/util';

import { Unauth } from '../containers/Unauth';

export class Signup extends Component {
  state: {
    errors: ?ErrorMap,
    redirectToApp: boolean
  }
  onSubmit: Function;

  constructor(props) {
    super(props)
    this.state = {
      errors: null,
      redirectToApp: false
    }
    this.onSubmit = this.onSubmit.bind(this)
  }

  async onSubmit(event) {
    event.preventDefault()
    const errors = await FKApiClient.get().signUp(
      this.refs.email.value,
      this.refs.username.value,
      this.refs.password.value,
      this.refs.invite.value
    );
    if (errors) {
      this.setState({ errors });
    } else {
      this.setState({ errors: null, redirectToApp: true });
    }
  }

  render() {
    if (this.state.redirectToApp) {
      return <Redirect to="/app"/>;
    }

    return (
      <Unauth>
        <div className="signup">
          <header>
            <h1>Join us</h1>
          </header>

          <form onSubmit={this.onSubmit}>
            <div className="content">
              <div className="group">
                <label htmlFor="email">Email</label>
                <input ref="email" id="email" name="email" type="email" placeholder="explorer@email.com" />
                { errorsFor(this.state.errors, 'email') }
              </div>

              <div className="group">
                <label htmlFor="username">Username</label>
                <input ref="username" id="username" name="username" type="text" placeholder="explorer123" />
                { errorsFor(this.state.errors, 'username') }
              </div>

              <div className="group">
                <label htmlFor="password">Password</label>
                <input ref="password" id="password" name="password" type="password" placeholder="correct horse battery staple" />
                { errorsFor(this.state.errors, 'password') }
              </div>

              <div className="group">
                <label htmlFor="invite">Invitation token</label>
                <input ref="invite" id="invite" name="invite" type="text"  placeholder="E3NANDTJ3YXM5LNMGNZTF2373LAFTOCC"/>
                { errorsFor(this.state.errors, 'invite') }
              </div>
            </div>

            <footer>
              { this.state.errors &&
                <p className="errors">
                  We found one or multiple errors. Please check your information above or try again later.
                </p> }
              <input type="submit" value="Submit"/>
            </footer>
          </form>
        </div>
      </Unauth>
    )
  }
}
