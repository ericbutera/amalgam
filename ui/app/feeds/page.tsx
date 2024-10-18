"use client";

import Feeds from "../components/feed/list";
import FeedForm from "../components/feed/form/form";

export default function Page() {
  return (
    <div>
      <h1>feed list</h1>
      <Feeds />
      <hr />
      <FeedForm />
    </div>
  );
}
