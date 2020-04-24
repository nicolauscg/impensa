import React, { useEffect, useCallback } from "react";
import { useFormik } from "formik";
import { useDropzone } from "react-dropzone";
import * as R from "ramda";

import {
  Button,
  TextField,
  Avatar,
  FormControl,
  FormHelperText
} from "@material-ui/core";

import { makeStyles } from "@material-ui/core/styles";
import { getUserObject } from "../../auth";
import { useAxiosSafely, urlGetUser, urlUpdateUser } from "../../api";
import {
  cleanEmptyFromObject,
  transformValuesToUpdateIdPayload,
  pngImagetoBase64
} from "../../ramdaHelpers";

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
  }
}));

export default function ProfilePage() {
  const classes = useStyles();
  const username = getUserObject().username;
  const userId = getUserObject().id;

  const [{ data: userData }, fetchUser] = useAxiosSafely(urlGetUser(userId));
  useEffect(() => {
    fetchUser();
  }, []);
  const [, updateUser] = useAxiosSafely(urlUpdateUser());
  const formikUser = useFormik({
    initialValues: {
      email: "",
      username: "",
      picture: null,
      oldPassword: null,
      newPassword: null
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
          formikBag.setFieldValue("oldPassword", null);
          formikBag.setFieldValue("newPassword", null);
          fetchUser();
        })
        .catch(err => {
          const errorMessage = err.response.data.error.message;
          if (errorMessage.indexOf("password") !== -1) {
            formikBag.setFieldError("oldPassword", errorMessage);
          }
        });
    }
  });

  useEffect(() => {
    formikUser.setValues({
      ...userData,
      oldPassword: null,
      newPassword: null
    });
  }, [userData]);

  const onDrop = useCallback(async acceptedFiles => {
    formikUser.setFieldValue(
      "picture",
      (await pngImagetoBase64(acceptedFiles[0])).split(",")[1]
    );
  }, []);
  const { getRootProps, getInputProps } = useDropzone({ onDrop });

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
