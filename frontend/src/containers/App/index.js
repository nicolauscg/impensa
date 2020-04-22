import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Switch, Route, withRouter, Redirect } from "react-router-dom";
import { connect } from "react-redux";
import { ThemeProvider } from "styled-components";
import { isLoggedIn } from "../../auth";
import Navbar from "../../components/Navbar";
import Grid from "@material-ui/core/Grid";
import Container from "@material-ui/core/Container";

import { theme } from "./theme";
import { routes } from "./routes";
import ProfileBadge from "../../components/ProfileBadge";

const useStyles = makeStyles(() => ({
  root: {
    minHeight: "100vh"
  },
  container: {
    position: "relative"
  }
}));

export const App = ({ history }) => {
  const pages = routes.map((route, index) =>
    route.protected ? (
      <PrivateRoute key={index} {...route} />
    ) : (
      <Route key={index} {...route} />
    )
  );
  const classes = useStyles();

  return (
    <ThemeProvider theme={theme}>
      <Grid container className={classes.root}>
        {isLoggedIn() ? (
          <>
            <Grid container item xs={2} direction={"column"}>
              <Navbar history={history} />
            </Grid>
            <Grid container item xs={10} className="py-4">
              <Container
                fixed={true}
                className={`d-flex flex-column ${classes.container}`}
              >
                <ProfileBadge history={history} />
                <Switch>{pages}</Switch>
              </Container>
            </Grid>
          </>
        ) : (
          <Container fixed={true} className="d-flex flex-column">
            <Switch>{pages}</Switch>
          </Container>
        )}
      </Grid>
    </ThemeProvider>
  );
};

function PrivateRoute({ component: Component, ...rest }) {
  return (
    <Route
      {...rest}
      render={props =>
        isLoggedIn() ? (
          <Component {...props} />
        ) : (
          <Redirect
            to={{
              pathname: "/auth",
              state: { from: props.location }
            }}
          />
        )
      }
    />
  );
}

export default withRouter(
  connect(
    null,
    null
  )(App)
);
