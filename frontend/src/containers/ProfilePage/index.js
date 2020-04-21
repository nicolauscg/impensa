import React from "react";
import { getUserObject } from "../../auth";

export default function ProfilePage() {
  const username = getUserObject().username;

  return (
    <>
      <h1>{username}&apos;s profile</h1>
      <p>details here</p>
    </>
  );
}
