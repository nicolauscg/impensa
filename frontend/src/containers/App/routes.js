import Home from "../Dashboard";
import AuthPage from "../AuthPage";
import NotFoundPage from "../../components/NotFoundPage";
import TransactionsPage from "../TransactionsPage";
import AccountsPage from "../AccountsPage";
import CategoriesPage from "../CategoriesPage";
import ProfilePage from "../ProfilePage";
import GraphPage from "../GraphPage";
import ImportExportPage from "../ImportExportPage";
import VerifyUserPage from "../VerifyUserPage";
import RequestResetPasswordPage from "../RequestResetPasswordPage";
import ResetPasswordPage from "../ResetPasswordPage";

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
  {
    component: AccountsPage,
    exact: true,
    path: "/account",
    protected: true
  },
  {
    component: CategoriesPage,
    exact: true,
    path: "/category",
    protected: true
  },
  {
    component: ProfilePage,
    exact: true,
    path: "/profile",
    protected: true
  },
  {
    component: GraphPage,
    exact: true,
    path: "/graph",
    protected: true
  },
  {
    component: ImportExportPage,
    exact: true,
    path: "/importexport",
    protected: true
  },
  {
    component: VerifyUserPage,
    exact: true,
    path: "/auth/verify",
    protected: false
  },
  {
    component: RequestResetPasswordPage,
    exact: true,
    path: "/auth/requestresetpassword",
    protected: false
  },
  {
    component: ResetPasswordPage,
    exact: true,
    path: "/auth/resetpassword",
    protected: false
  },
  { component: NotFoundPage }
];
