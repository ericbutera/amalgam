import React, { useState } from "react";
import useAddFeedMutation from "@/app/data/feed-add";

const AddFeedForm = () => {
  const [name, setName] = useState("");
  const [url, setUrl] = useState("");
  const { addFeed, loading, error, validationErrors } = useAddFeedMutation();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await addFeed(name, url);
  };

  return (
    <div>
      {error && (
        <div role="alert" className="alert alert-error">
          <p>{error}</p>
        </div>
      )}

      {validationErrors.length > 0 && (
        <div role="alert" className="alert alert-error">
          {validationErrors.map((msg, idx) => (
            <p key={idx}>{msg}</p>
          ))}
        </div>
      )}

      <form onSubmit={handleSubmit}>
        <div>
          <label className="input input-bordered flex items-center gap-2">
            Name
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              disabled={loading}
              className="grow"
            />
          </label>
        </div>
        <div>
          <label className="input input-bordered flex items-center gap-2">
            URL
            <input
              type="text"
              placeholder="https://example.com/rss"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              disabled={loading}
              className="grow"
            />
          </label>
        </div>

        <button type="submit" disabled={loading} className="btn btn-primary">
          Add Feed
        </button>
      </form>
    </div>
  );
};

export default AddFeedForm;
