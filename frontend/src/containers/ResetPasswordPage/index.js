import React, { useState } from "react";
import { useFormik } from "formik";
import * as R from "ramda";

import {
  Button,
  Snackbar,
  FormControl,
  FormHelperText
} from "@material-ui/core";
import MuiAlert from "@material-ui/lab/Alert";

import { useAxiosSafely, urlResetUserPassword } from "../../api";
import { FormFields } from "../../components/CreateOrEditModal";
import { getPasswordDifficulty } from "../../components/LoginRegisterBox";

const queryString = require("query-string");

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const ResetPasswordPage = ({ location }) => {
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const handleOpenSnackbar = () => {
    setSnackbarOpen(true);
  };
  const handleCloseSnackbar = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }
    setSnackbarOpen(false);
  };

  const { email, verifyKey } = queryString.parse(location.search);
  const [, changeUserPassword] = useAxiosSafely(urlResetUserPassword());
  const formikRequestPasswordChange = useFormik({
    initialValues: {
      email,
      newPassword: "",
      verifyKey
    },
    validate: values => {
      const errors = {};
      const { score } = getPasswordDifficulty(values.newPassword);
      if (score === 0) {
        errors.newPassword = `password strength too low`;
      }

      return errors;
    },
    validateOnChange: true,
    onSubmit: values => {
      changeUserPassword({
        data: values
      }).then(handleOpenSnackbar);
    }
  });

  return (
    <div className="d-flex flex-column justify-content-center flex-grow-1">
      <h1>Change password</h1>
      <Snackbar
        open={snackbarOpen}
        autoHideDuration={3000}
        onClose={handleCloseSnackbar}
      >
        <Alert onClose={handleCloseSnackbar} severity="success">
          password changed
        </Alert>
      </Snackbar>
      <form onSubmit={formikRequestPasswordChange.handleSubmit}>
        <FormControl
          fullWidth={true}
          error={R.hasPath(["newPassword"], formikRequestPasswordChange.errors)}
        >
          {FormFields.textField({
            label: "new password",
            type: "password",
            name: "newPassword"
          })(formikRequestPasswordChange)}
          <FormHelperText id="my-helper-text">
            {R.propOr("", "newPassword", formikRequestPasswordChange.errors)}
          </FormHelperText>
        </FormControl>
        <Button
          variant="contained"
          color="primary"
          type="submit"
          fullWidth={true}
          className="mt-5"
          disabled={
            getPasswordDifficulty(
              formikRequestPasswordChange.values.newPassword
            ).score === 0
          }
        >
          send request
        </Button>
      </form>
    </div>
  );
};

export default ResetPasswordPage;
