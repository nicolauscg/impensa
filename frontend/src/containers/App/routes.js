import Home from "../../containers/Home";
import Auth from "../../containers/LoginRegister";
import NotFoundPage from "../../components/NotFoundPage";

export const routes = [
  {
    component: Home,
    exact: true,
    path: "/"
  },
  {
    component: Auth,
    exact: true,
    path: "/auth"
  },
  { component: NotFoundPage }
];
