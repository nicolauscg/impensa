import React, { useContext } from "react";
import { useFormik } from "formik";
import useAxios from "axios-hooks";
import * as R from "ramda";

import LoginRegisterBox from "../../components/LoginRegisterBox";
import { isLoggedIn } from "../../auth";
import { urlLogin, urlRegister } from "../../api";
import { UserContext } from "../../containers/App/index";

const AuthPage = ({ history }) => {
  if (isLoggedIn()) {
    history.push("/");
  }

  const { refreshUserContext } = useContext(UserContext);
  const [, postLogin] = useAxios(urlLogin(), { manual: true });
  const [, postRegister] = useAxios(urlRegister(), { manual: true });

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
      onSubmit: values => {
        postRegister({
          data: values
        }).then(() => {
          refreshUserContext();
          history.push("/");
        });
      }
    })
  };

  return (
    <div className="align-self-center">
      <h1>Impensa</h1>
      <LoginRegisterBox {...passedProp} />
    </div>
  );
};

export default AuthPage;
