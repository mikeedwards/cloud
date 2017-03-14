// @flow weak

import React, { Component } from 'react'
import { Redirect } from 'react-router'

import { FKApiClient } from '../../api/api';
import { errorsFor } from '../../common/util';
import type { ErrorMap } from '../../common/util';

import { Unauth } from '../containers/Unauth';

type Props = {
  requestSignIn: (u: string, p: string) => Promise<?ErrorMap>;
  location: Object;
};

export class Signin extends Component {
  props: Props;
  state: {
    errors: ?ErrorMap,
    redirectToReferrer: boolean
  }
  onSubmit: Function;

  constructor(props: Props) {
    super(props)
    this.state = {
      errors: null,
      redirectToReferrer: false
    }
    this.onSubmit = this.onSubmit.bind(this)
  }

  async onSubmit(event) {
    event.preventDefault()
    const errors = await FKApiClient.get().signIn(this.refs.username.value, this.refs.password.value);
    if (errors) {
      this.setState({ errors });
    } else {
      this.setState({ errors: null, redirectToReferrer: true });
    }
  }

  render() {
    if (this.state.redirectToReferrer) {
      let from = this.props.location.state || '/';
      if (from === '/signin') { from = '/app' };

      return <Redirect to={from}/>;
    }

    return (
      <Unauth>
        <div className="signin">
          <header>
            <h1>Sign in</h1>
          </header>

          <form onSubmit={this.onSubmit}>
            <div className="content">
              <div className="group">
                <label htmlFor="username">Username</label>
                <input ref="username" id="username" name="username" type="username" placeholder="" />
                { errorsFor(this.state.errors, 'username') }
              </div>
              <div className="group">
                <label htmlFor="password">Password</label>
                <input ref="password" id="password" name="password" type="password" placeholder="" />
                { errorsFor(this.state.errors, 'password') }
              </div>
            </div>
            <footer>
              { this.state.errors &&
                <p className="errors">
                  Username or password invalid. Check your information and try again.
                </p> }
              <input type="submit" value="Submit"/>
            </footer>
          </form>
        </div>
      </Unauth>
    )
  }
}
