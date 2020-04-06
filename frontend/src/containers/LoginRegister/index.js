import React from "react";
import { useFormik } from "formik";
import useAxios from "axios-hooks";
import { isLoggedIn } from "../../auth";
import { urlLogin } from "../../api";

const LoginRegister = ({ history }) => {
  if (isLoggedIn()) {
    history.push("/");
  }

  return (
    <>
      <h1>auth page</h1>
      <h2>login</h2>
      <LoginForm history={history} />
    </>
  );
};

const LoginForm = ({ history }) => {
  const [{ error }, executepost] = useAxios(urlLogin(), { manual: true });

  const formik = useFormik({
    initialValues: {
      email: "",
      password: ""
    },
    onSubmit: values => {
      executepost({
        data: values
      }).then(result => {
        localStorage.impensa = JSON.stringify(result.data.data);
        history.push("/");
      });
    }
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      {error && <p>{error}</p>}
      <label htmlFor="email">Email Address</label>
      <input
        id="email"
        name="email"
        type="email"
        onChange={formik.handleChange}
        value={formik.values.email}
      />
      <label htmlFor="password">Password</label>
      <input
        id="password"
        name="password"
        type="password"
        onChange={formik.handleChange}
        value={formik.values.password}
      />
      <button type="submit">Submit</button>
    </form>
  );
};

export default LoginRegister;
