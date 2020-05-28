import React, { useEffect, useContext } from "react";

import { useAxiosSafely, urlGoogleCallback } from "../../api";
import { UserContext } from "../../containers/App/index";

const queryString = require("query-string");

const GoogleOauthCallback = ({ location, history }) => {
  const { code, state } = queryString.parse(location.search);
  const [, getGoogleCallback] = useAxiosSafely(urlGoogleCallback());

  const { refreshUserContext } = useContext(UserContext);

  useEffect(() => {
    getGoogleCallback({
      params: {
        code,
        state
      }
    }).then(() => {
      refreshUserContext();
      history.push("/");
    });
  }, []);

  return (
    <>
      <p>logging in with google..</p>
    </>
  );
};

export default GoogleOauthCallback;
