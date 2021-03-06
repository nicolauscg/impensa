import React, { useContext } from "react";
import { useFormik } from "formik";
import useAxios from "axios-hooks";
import * as R from "ramda";

import { Button } from "@material-ui/core";

import LoginRegisterBox, {
  getPasswordDifficulty
} from "../../components/LoginRegisterBox";
import { isLoggedIn } from "../../auth";
import { urlLogin, urlRegister, urlGoogleLogin } from "../../api";
import { UserContext } from "../../containers/App/index";

const AuthPage = ({ history }) => {
  if (isLoggedIn()) {
    history.push("/");
  }

  const { refreshUserContext } = useContext(UserContext);
  const [, postLogin] = useAxios(urlLogin(), { manual: true });
  const [, postRegister] = useAxios(urlRegister(), { manual: true });
  const [, postGoogleLogin] = useAxios(urlGoogleLogin(), { manual: true });

  const passedProp = {
    history,
    formikLogin: useFormik({
      initialValues: {
        email: "",
        password: "",
        rememberMe: false
      },
      onSubmit: (values, formikBag) => {
        postLogin({
          data: values
        })
          .then(() => {
            refreshUserContext();
            history.push("/");
          })
          .catch(err => {
            const errorMessage = R.pathOr(
              "",
              ["response", "data", "error", "message"],
              err
            );
            if (errorMessage.indexOf("email") !== -1) {
              formikBag.setFieldError("email", errorMessage);
            } else if (errorMessage.indexOf("password") !== -1) {
              formikBag.setFieldError("password", errorMessage);
            } else {
              throw new Error(
                "unrecognized error message response after login"
              );
            }
          });
      }
    }),
    formikRegister: useFormik({
      initialValues: {
        email: "",
        username: "",
        password: ""
      },
      validate: values => {
        const errors = {};
        const { score } = getPasswordDifficulty(values.password);
        if (score === 0) {
          errors.password = `password strength too low`;
        }

        return errors;
      },
      validateOnChange: true,
      onSubmit: (values, formikBag) => {
        postRegister({
          data: values
        })
          .then(() => {
            refreshUserContext();
            history.push("/");
          })
          .catch(err => {
            const errorMessage = R.pathOr(
              "",
              ["response", "data", "error", "message"],
              err
            );
            if (errorMessage.indexOf("email") !== -1) {
              formikBag.setFieldError("email", errorMessage);
            } else if (errorMessage.indexOf("username") !== -1) {
              formikBag.setFieldError("username", errorMessage);
            } else {
              throw new Error(
                "unrecognized error message response after register"
              );
            }
          });
      }
    })
  };
  passedProp.disableSubmit =
    getPasswordDifficulty(passedProp.formikRegister.values.password).score ===
    0;

  return (
    <div className="align-self-center">
      <h1>Impensa</h1>
      <LoginRegisterBox {...passedProp} />
      <Button
        variant="contained"
        onClick={() => {
          postGoogleLogin().then(res => {
            const url = R.pathOr("", ["data", "data"], res);
            if (url) {
              window.location.assign(url);
            }
          });
        }}
      >
        sign in with google
      </Button>
    </div>
  );
};

export default AuthPage;
