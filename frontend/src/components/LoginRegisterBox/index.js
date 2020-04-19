import React from "react";
import SwipeableViews from "react-swipeable-views";

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
  Checkbox
} from "@material-ui/core";

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

const LoginRegisterBox = ({ formikLogin, formikRegister }) => {
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
            <TextField
              id="email"
              label="email"
              name="email"
              type="email"
              onChange={formikLogin.handleChange}
              value={formikLogin.values.email}
              fullWidth={true}
            />
            <TextField
              id="password"
              label="password"
              name="password"
              type="password"
              onChange={formikLogin.handleChange}
              value={formikLogin.values.password}
              fullWidth={true}
            />
            <FormControlLabel
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
            <TextField
              id="email"
              label="email"
              name="email"
              type="email"
              onChange={formikRegister.handleChange}
              value={formikRegister.values.email}
              fullWidth={true}
            />
            <TextField
              id="username"
              label="username"
              name="username"
              type="text"
              onChange={formikRegister.handleChange}
              value={formikRegister.values.username}
              fullWidth={true}
            />
            <TextField
              id="password"
              label="password"
              name="password"
              type="password"
              onChange={formikRegister.handleChange}
              value={formikRegister.values.password}
              fullWidth={true}
            />
            <Button
              variant="contained"
              color="primary"
              type="submit"
              fullWidth={true}
              className="mt-5"
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
