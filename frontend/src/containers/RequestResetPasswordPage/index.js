import React, { useState } from "react";
import { useFormik } from "formik";
import { useAxiosSafely, urlRequestResetUserPassword } from "../../api";

import { Button, Snackbar } from "@material-ui/core";
import MuiAlert from "@material-ui/lab/Alert";

import { FormFields } from "../../components/CreateOrEditModal";

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const RequestResetPasswordPage = ({ history }) => {
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

  const [, requestChangeUserPassword] = useAxiosSafely(
    urlRequestResetUserPassword()
  );
  const formikRequestPasswordChange = useFormik({
    initialValues: {
      email: ""
    },
    onSubmit: values => {
      requestChangeUserPassword({
        data: values
      }).then(handleOpenSnackbar);
    }
  });

  return (
    <div className="d-flex flex-column justify-content-center flex-grow-1">
      <h1 onClick={() => history.push("/")}>Impensa</h1>
      <h3>Request change password</h3>
      <Snackbar
        open={snackbarOpen}
        autoHideDuration={3000}
        onClose={handleCloseSnackbar}
      >
        <Alert onClose={handleCloseSnackbar} severity="success">
          email to reset password sent!
        </Alert>
      </Snackbar>
      <form onSubmit={formikRequestPasswordChange.handleSubmit}>
        {FormFields.textField({
          label: "email",
          type: "text",
          name: "email"
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

export default RequestResetPasswordPage;
