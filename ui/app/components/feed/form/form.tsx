import React, { useState } from "react";
import useAddFeedMutation from "@/app/data/feed-add";
import { LinkIcon } from "@heroicons/react/24/solid";

const AddFeedForm = () => {
  const [url, setUrl] = useState("");
  const { addFeed, loading, error, validationErrors } = useAddFeedMutation();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await addFeed(url);
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
            <div key={idx}>{msg}</div>
          ))}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="items-center ">
          <label className="input input-bordered flex items-center gap-2">
            <LinkIcon className="w-6 h-6 text-default-500" title="Feed URL" />
            <input
              type="url"
              className="grow"
              placeholder="https://example.com/rss"
              title="Field must be a valid secure URL"
              pattern="https://.*"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              disabled={loading}
              required
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
