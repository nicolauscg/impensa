import Home from "../Dashboard";
import AuthPage from "../AuthPage";
import NotFoundPage from "../../components/NotFoundPage";
import TransactionsPage from "../TransactionsPage";

export const routes = [
  {
    component: Home,
    exact: true,
    path: "/",
    protected: true
  },
  {
    component: AuthPage,
    exact: true,
    path: "/auth",
    protected: false
  },
  {
    component: TransactionsPage,
    exact: true,
    path: "/transaction",
    protected: true
  },
  { component: NotFoundPage }
];
