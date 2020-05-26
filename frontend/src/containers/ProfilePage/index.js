import React, { useCallback, useContext } from "react";
import { useFormik } from "formik";
import { useDropzone } from "react-dropzone";
import * as R from "ramda";

import {
  Button,
  TextField,
  Avatar,
  FormControl,
  FormHelperText,
  Typography
} from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import { useAxiosSafely, urlUpdateUser } from "../../api";
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

export default function ProfilePage() {
  const classes = useStyles();

  const { data: userData, refreshUserContext } = useContext(UserContext);
  const username = userData.username;

  const [, updateUser] = useAxiosSafely(urlUpdateUser());
  const formikUser = useFormik({
    initialValues: {
      ...userData,
      oldPassword: "",
      newPassword: ""
    },
    enableReinitialize: true,
    onSubmit: (values, formikBag) => {
      updateUser({
        data: R.pipe(
          cleanEmptyFromObject,
          transformValuesToUpdateIdPayload
        )(values)
      })
        .then(() => {
          refreshUserContext();
          formikBag.resetForm();
        })
        .catch(err => {
          const errorMessage = err.response.data.error.message;
          if (errorMessage.indexOf("password") !== -1) {
            formikBag.setFieldError("oldPassword", errorMessage);
          }
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
        <FormControl
          fullWidth={true}
          error={R.hasPath(["oldPassword"], formikUser.errors)}
        >
          <TextField
            id="oldPassword"
            label="old password"
            name="oldPassword"
            type="password"
            onChange={formikUser.handleChange}
            value={formikUser.values.oldPassword}
            fullWidth={true}
            error={R.hasPath(["oldPassword"], formikUser.errors)}
          />
          <FormHelperText>
            {R.propOr("", "oldPassword", formikUser.errors)}
          </FormHelperText>
        </FormControl>
        <TextField
          id="newPassword"
          label="new password"
          name="newPassword"
          type="password"
          onChange={formikUser.handleChange}
          value={formikUser.values.newPassword}
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
    </>
  );
}
