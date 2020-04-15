import React from "react";
import { getUserObject } from "../../auth";

const Dashboard = () => {
  const username = getUserObject().username;

  return <h2>Welcome, {username}!</h2>;
};

export default Dashboard;
