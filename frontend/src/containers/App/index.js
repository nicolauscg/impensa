/* eslint-disable */
import React from "react";
import { Switch, Route, withRouter, Redirect } from "react-router-dom";
import { connect } from "react-redux";
import { ThemeProvider } from "styled-components";
import { isLoggedIn, getUserObject } from "../../auth";

import { theme } from "./theme";
import { routes } from "./routes";

export class App extends React.Component {
  render() {
    const pages = routes.map((route, i) => (
      <Route
        component={route.component}
        exact={route.exact}
        path={route.path}
        key={i}
      />
    ));
    return (
      <ThemeProvider theme={theme}>
        <ul>
          <li onClick={() => this.props.history.push("/")}>impensa</li>
          {!isLoggedIn() && (
            <li onClick={() => this.props.history.push("/auth")}>auth</li>
          )}
          {isLoggedIn() && (
            <>
              <li>Hello, {getUserObject().email}!</li>
              <li
                onClick={() => {
                  localStorage.removeItem("impensa");
                  this.props.history.push("/auth");
                }}
              >
                logout
              </li>
            </>
          )}
        </ul>
        <Switch>{pages}</Switch>
      </ThemeProvider>
    );
  }
}

export default withRouter(
  connect(
    null,
    null
  )(App)
);
