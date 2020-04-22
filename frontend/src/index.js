import React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import store, { history } from "./store";
import { ConnectedRouter } from "connected-react-router";
import { configure } from "axios-hooks";
import LRU from "lru-cache";
import Axios from "axios";

import "./index.css";
import "normalize.css";
import App from "./containers/App";
import * as serviceWorker from "./serviceWorker";
import { clearUserObject } from "./auth";

import { MuiPickersUtilsProvider } from "@material-ui/pickers";
import MomentUtils from "@date-io/moment";

const target = document.querySelector("#root");

const axios = Axios.create({ withCredentials: true });
const cache = new LRU({ max: 10 });
axios.interceptors.request.use(
  config => {
    try {
      const userData = JSON.parse(localStorage.impensa);
      if (userData.token) {
        config.headers.Authorization = `Bearer ${userData.token}`;
      }
    } catch (err) {
      localStorage.removeItem("impensa");
    }

    return config;
  },
  error => Promise.reject(error)
);
axios.interceptors.response.use(
  response => response,
  error => {
    const statusCode = error.response.status;

    if (statusCode === 401) {
      clearUserObject();
      Reflect.deleteProperty(axios.defaults.headers.common, "Authorization");
      history.push("/auth");
    }

    return Promise.reject(error);
  }
);
configure({
  axios,
  cache
});

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <MuiPickersUtilsProvider utils={MomentUtils}>
        <App />
      </MuiPickersUtilsProvider>
    </ConnectedRouter>
  </Provider>,
  target
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
