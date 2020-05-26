import React, { useEffect, useState } from "react";
import { useAxiosSafely, urlVerifyUser } from "../../api";

import { makeStyles } from "@material-ui/core/styles";
import { Typography, Button } from "@material-ui/core";

const queryString = require("query-string");

const useStyles = makeStyles(() => ({
  messageBox: {
    width: "20rem",
    maxWidth: "95%"
  }
}));

const VerifyUserPage = ({ location, history }) => {
  const classes = useStyles();

  const [success, setSuccess] = useState(null);
  const [, verifyUser] = useAxiosSafely(urlVerifyUser());

  const { userId, verifyKey } = queryString.parse(location.search);

  useEffect(() => {
    if (userId !== undefined && verifyKey !== undefined) {
      verifyUser({
        params: {
          userId,
          verifyKey
        }
      })
        .then(() => {
          setSuccess(true);
        })
        .catch(() => {
          setSuccess(false);
        });
    }
  }, []);

  let pageContent = <></>;
  if (success === true) {
    pageContent = (
      <div
        className={`${
          classes.messageBox
        } d-flex flex-column justify-content-center flex-grow-1 align-self-center`}
      >
        <Typography variant="h4" component="h4" align="center">
          account successfully verified
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={() => history.push("/auth")}
          className="mt-3"
        >
          go to login page
        </Button>
      </div>
    );
  } else if (success === false) {
    pageContent = (
      <div
        className={`${
          classes.messageBox
        } d-flex flex-column justify-content-center flex-grow-1 align-self-center`}
      >
        <Typography variant="h4" component="h4" align="center">
          failed to verify account, check that link is correct
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={() => history.push("/auth")}
          className="mt-3"
        >
          go to login page
        </Button>
      </div>
    );
  }

  return <>{pageContent}</>;
};

export default VerifyUserPage;
