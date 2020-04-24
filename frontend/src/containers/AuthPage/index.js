import React from "react";
import { useFormik } from "formik";
import useAxios from "axios-hooks";

import LoginRegisterBox from "../../components/LoginRegisterBox";
import { isLoggedIn } from "../../auth";
import { urlLogin, urlRegister } from "../../api";

const AuthPage = ({ history }) => {
  if (isLoggedIn()) {
    history.push("/");
  }

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
            history.push("/");
          })
          .catch(err => {
            const errorMessage = err.response.data.error.message;
            if (errorMessage.indexOf("email") !== -1) {
              formikBag.setFieldError("email", errorMessage);
            } else if (errorMessage.indexOf("password") !== -1) {
              formikBag.setFieldError("password", errorMessage);
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
