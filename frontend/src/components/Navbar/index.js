import React from "react";

import AppBar from "@material-ui/core/AppBar";
import { makeStyles } from "@material-ui/core/styles";
import { Toolbar, Typography, Button } from "@material-ui/core";

import { isLoggedIn } from "../../auth";

const useStyles = makeStyles({
  buttonOutlinedPrimary: {
    color: "white"
  }
});

const Navbar = props => {
  const classes = useStyles();
  const links = [
    {
      name: "login",
      action: ({ history }) => () => history.push("/auth"),
      authorized: false
    },
    {
      name: "register",
      action: ({ history }) => () => history.push("/auth"),
      authorized: false
    },
    {
      name: "transactions",
      action: ({ history }) => () => history.push("/transaction"),
      authorized: true
    },
    {
      name: "accounts",
      action: ({ history }) => () => history.push("/account"),
      authorized: true
    },
    {
      name: "categories",
      action: ({ history }) => () => history.push("/category"),
      authorized: true
    },
    {
      name: "graph",
      action: ({ history }) => () => history.push("/graph"),
      authorized: true
    },
    {
      name: "import export",
      action: ({ history }) => () => history.push("/importexport"),
      authorized: true
    }
  ];

  return (
    <AppBar position="static" className={"flex-column flex-grow-1"}>
      <Toolbar className="flex-column flex-grow-1 py-4">
        <Button
          onClick={() => props.history.push("/")}
          className="p-3"
          variant="outlined"
          color="primary"
          classes={{
            outlinedPrimary: classes.buttonOutlinedPrimary
          }}
        >
          <Typography variant="h4">Impensa</Typography>
        </Button>
        {links
          .filter(link => (isLoggedIn() ? link.authorized : !link.authorized))
          .map(link => (
            <Button
              key={link.name}
              color="inherit"
              onClick={link.action(props)}
            >
              {link.name}
            </Button>
          ))}
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
