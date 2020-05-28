import React, { useCallback, useContext, useState } from "react";
import { useFormik } from "formik";
import { useDropzone } from "react-dropzone";
import * as R from "ramda";

import {
  Button,
  TextField,
  Avatar,
  Typography,
  Snackbar
} from "@material-ui/core";
import MuiAlert from "@material-ui/lab/Alert";
import { makeStyles } from "@material-ui/core/styles";

import {
  useAxiosSafely,
  urlUpdateUser,
  urlRequestResetUserPassword
} from "../../api";
import {
  cleanEmptyFromObject,
  transformValuesToUpdateIdPayload,
  pngImagetoBase64
} from "../../ramdaHelpers";
import { UserContext } from "../../containers/App/index";

const useStyles = makeStyles(theme => ({
  pictureField: {
    display: "inline-block",
    position: "relative",
    "&:hover $largePicture": {
      opacity: 0.3
    },
    "&:hover $textBehindPicture": {
      opacity: 1
    }
  },
  largePicture: {
    width: theme.spacing(20),
    height: theme.spacing(20),
    transition: ".5s ease"
  },
  textBehindPicture: {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    textAlign: "center",
    opacity: 0,
    transition: ".5s ease",
    fontWeight: "bold",
    textShadow: "white 0px 0px 10px"
  },
  accountVerifiedStatus: {
    color: "#4caf50"
  },
  accountUnverifiedStatus: {
    color: "#e91e63"
  }
}));

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

export default function ProfilePage() {
  const classes = useStyles();

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

  const { data: userData, refreshUserContext } = useContext(UserContext);
  const username = userData.username;

  const [, updateUser] = useAxiosSafely(urlUpdateUser());
  const [, requestChangeUserPassword] = useAxiosSafely(
    urlRequestResetUserPassword()
  );
  const formikUser = useFormik({
    initialValues: {
      ...userData
    },
    enableReinitialize: true,
    onSubmit: (values, formikBag) => {
      updateUser({
        data: R.pipe(
          cleanEmptyFromObject,
          transformValuesToUpdateIdPayload
        )(values)
      }).then(() => {
        refreshUserContext();
        formikBag.resetForm();
      });
    }
  });

  const onDrop = useCallback(async acceptedFiles => {
    formikUser.setFieldValue(
      "picture",
      (await pngImagetoBase64(acceptedFiles[0])).split(",")[1]
    );
  }, []);
  const { getRootProps, getInputProps } = useDropzone({ onDrop });
  const accountIsVerified = R.propOr(false, "verified", userData);

  return (
    <>
      <h1>{username}&apos;s profile</h1>
      <Snackbar
        open={snackbarOpen}
        autoHideDuration={3000}
        onClose={handleCloseSnackbar}
      >
        <Alert onClose={handleCloseSnackbar} severity="success">
          email to reset password sent!
        </Alert>
      </Snackbar>
      <form onSubmit={formikUser.handleSubmit}>
        <div {...getRootProps()} className={classes.pictureField}>
          <input {...getInputProps()} />
          <Avatar
            alt={username}
            src={
              formikUser.values.picture
                ? `data:image/png;base64,${formikUser.values.picture}`
                : null
            }
            className={classes.largePicture}
          />
          <p className={classes.textBehindPicture}>
            click or drag and drop PNG to edit
          </p>
        </div>
        <Typography
          variant="subtitle2"
          component="p"
          className="my-2"
          classes={{
            root: accountIsVerified
              ? classes.accountVerifiedStatus
              : classes.accountUnverifiedStatus
          }}
        >
          {accountIsVerified ? "account verified" : "account unverified"}
        </Typography>
        <TextField
          id="email"
          label="email"
          name="email"
          type="email"
          onChange={formikUser.handleChange}
          value={formikUser.values.email}
          fullWidth={true}
        />
        <TextField
          id="username"
          label="username"
          name="username"
          type="text"
          onChange={formikUser.handleChange}
          value={formikUser.values.username}
          fullWidth={true}
        />
        <Button
          variant="contained"
          color="primary"
          type="submit"
          fullWidth={true}
          className="mt-5"
        >
          Update
        </Button>
      </form>

      <Button
        variant="contained"
        color="secondary"
        type="submit"
        fullWidth={true}
        className="mt-5"
        onClick={() =>
          requestChangeUserPassword({
            data: {
              email: R.propOr("", "email", userData)
            }
          }).then(handleOpenSnackbar)
        }
      >
        request reset password
      </Button>
    </>
  );
}
