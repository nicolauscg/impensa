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
        password: ""
      },
      onSubmit: values => {
        postLogin({
          data: values
        }).then(result => {
          localStorage.impensa = JSON.stringify(result.data.data);
          history.push("/");
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
        }).then(result => {
          localStorage.impensa = JSON.stringify(result.data.data);
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
