import React from "react";
import SwipeableViews from "react-swipeable-views";
import * as R from "ramda";
import * as zxcvbn from "zxcvbn";

import { makeStyles, useTheme } from "@material-ui/core/styles";
import {
  Button,
  TextField,
  Box,
  Typography,
  Tab,
  Tabs,
  AppBar,
  FormControlLabel,
  Checkbox,
  FormControl,
  FormHelperText,
  Link
} from "@material-ui/core";

export const getPasswordDifficulty = password => {
  const { score } = zxcvbn(password);
  if (score < 2) {
    return {
      score: 0,
      message: "low"
    };
  } else if (score < 4) {
    return {
      score: 1,
      message: "medium"
    };
  }

  return {
    score: 2,
    message: "high"
  };
};

const TabPanel = props => {
  const { children, value, index, ...other } = props;

  return (
    <Typography
      component="div"
      role="tabpanel"
      hidden={value !== index}
      id={`full-width-tabpanel-${index}`}
      aria-labelledby={`full-width-tab-${index}`}
      {...other}
    >
      {value === index && <Box p={3}>{children}</Box>}
    </Typography>
  );
};

function a11yProps(index) {
  return {
    id: `full-width-tab-${index}`,
    "aria-controls": `full-width-tabpanel-${index}`
  };
}

const useStyles = makeStyles(theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
    width: 500
  }
}));

const LoginRegisterBox = ({
  formikLogin,
  formikRegister,
  history,
  disableSubmit
}) => {
  const classes = useStyles();
  const theme = useTheme();
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  const handleChangeIndex = index => {
    setValue(index);
  };

  return (
    <div className={classes.root}>
      <AppBar position="static" color="default">
        <Tabs
          value={value}
          onChange={handleChange}
          indicatorColor="primary"
          textColor="primary"
          variant="fullWidth"
          aria-label="full width tabs example"
        >
          <Tab label="Login" {...a11yProps(0)} />
          <Tab label="Register" {...a11yProps(1)} />
        </Tabs>
      </AppBar>
      <SwipeableViews
        axis={theme.direction === "rtl" ? "x-reverse" : "x"}
        index={value}
        onChangeIndex={handleChangeIndex}
      >
        <TabPanel value={value} index={0} dir={theme.direction}>
          <form onSubmit={formikLogin.handleSubmit}>
            <FormControl
              fullWidth={true}
              error={R.hasPath(["email"], formikLogin.errors)}
            >
              <TextField
                id="email"
                label="email"
                name="email"
                type="email"
                onChange={formikLogin.handleChange}
                value={formikLogin.values.email}
                fullWidth={true}
                error={R.hasPath(["email"], formikLogin.errors)}
              />
              <FormHelperText id="my-helper-text">
                {R.propOr("", "email", formikLogin.errors)}
              </FormHelperText>
            </FormControl>

            <FormControl
              fullWidth={true}
              error={R.hasPath(["password"], formikLogin.errors)}
            >
              <TextField
                id="password"
                label="password"
                name="password"
                type="password"
                onChange={formikLogin.handleChange}
                value={formikLogin.values.password}
                fullWidth={true}
                error={R.hasPath(["password"], formikLogin.errors)}
              />
              <FormHelperText id="my-helper-text">
                {R.propOr("", "password", formikLogin.errors)}
              </FormHelperText>
            </FormControl>
            <FormControlLabel
              className="d-block"
              control={
                <Checkbox
                  checked={formikLogin.values.rememberMe}
                  onChange={formikLogin.handleChange}
                  name="rememberMe"
                  color="primary"
                />
              }
              label="remember me"
            />
            <Link
              component="button"
              variant="body2"
              onClick={() => {
                history.push("/auth/requestresetpassword");
              }}
            >
              forget password
            </Link>
            <Button
              variant="contained"
              color="primary"
              type="submit"
              fullWidth={true}
              className="mt-5"
            >
              Login
            </Button>
          </form>
        </TabPanel>
        <TabPanel value={value} index={1} dir={theme.direction}>
          <form onSubmit={formikRegister.handleSubmit}>
            <FormControl
              fullWidth={true}
              error={R.hasPath(["email"], formikRegister.errors)}
            >
              <TextField
                id="email"
                label="email"
                name="email"
                type="email"
                onChange={formikRegister.handleChange}
                value={formikRegister.values.email}
                fullWidth={true}
                error={R.hasPath(["email"], formikRegister.errors)}
              />
              <FormHelperText id="my-helper-text">
                {R.propOr("", "email", formikRegister.errors)}
              </FormHelperText>
            </FormControl>
            <FormControl
              fullWidth={true}
              error={R.hasPath(["email"], formikRegister.errors)}
            >
              <TextField
                id="username"
                label="username"
                name="username"
                type="text"
                onChange={formikRegister.handleChange}
                value={formikRegister.values.username}
                fullWidth={true}
                error={R.hasPath(["username"], formikRegister.errors)}
              />
              <FormHelperText id="my-helper-text">
                {R.propOr("", "username", formikRegister.errors)}
              </FormHelperText>
            </FormControl>
            <FormControl
              fullWidth={true}
              error={R.hasPath(["password"], formikRegister.errors)}
            >
              <TextField
                id="password"
                label="password"
                name="password"
                type="password"
                onChange={formikRegister.handleChange}
                value={formikRegister.values.password}
                fullWidth={true}
              />
              <FormHelperText id="my-helper-text">
                {R.propOr("", "password", formikRegister.errors)}
              </FormHelperText>
            </FormControl>
            <Button
              variant="contained"
              color="primary"
              type="submit"
              fullWidth={true}
              className="mt-5"
              disabled={disableSubmit}
            >
              Register
            </Button>
          </form>
        </TabPanel>
      </SwipeableViews>
    </div>
  );
};

export default LoginRegisterBox;
