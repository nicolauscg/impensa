import React, { useState } from "react";
import { useFormik } from "formik";
import { useAxiosSafely, urlResetUserPassword } from "../../api";

import { Button, Snackbar } from "@material-ui/core";
import MuiAlert from "@material-ui/lab/Alert";

import { FormFields } from "../../components/CreateOrEditModal";

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
        {FormFields.textField({
          label: "new password",
          type: "password",
          name: "newPassword"
        })(formikRequestPasswordChange)}
        <Button
          variant="contained"
          color="primary"
          type="submit"
          fullWidth={true}
          className="mt-5"
        >
          send request
        </Button>
      </form>
    </div>
  );
};

export default ResetPasswordPage;
