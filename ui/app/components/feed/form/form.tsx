import React, { useState } from "react";
import useAddFeedMutation from "@/app/data/feed-add";
import {
  BeakerIcon,
  LinkIcon,
  PlusCircleIcon,
} from "@heroicons/react/24/solid";
import { v4 as uuidv4 } from "uuid";

const AddFeedForm = () => {
  const [url, setUrl] = useState("");
  const { addFeed, loading, error, validationErrors } = useAddFeedMutation();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const res = await addFeed(url);
    if (res) {
      setUrl("");
    }
  };

  const generateFakeFeed = () => {
    setUrl(`http://faker:8080/feed/${uuidv4()}`);
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
              pattern="(http|https)://.*"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              disabled={loading}
              required
            />
          </label>
        </div>

        <div className="flex justify-end gap-4">
          <button
            type="button"
            className="btn btn-secondary"
            onClick={generateFakeFeed}
            title="Generate a fake feed URL"
          >
            <BeakerIcon className="w-6 h-6 text-default-500" />
          </button>

          <button type="submit" disabled={loading} className="btn btn-primary">
            <PlusCircleIcon className="w-6 h-6 text-default-500" />
            Add Feed
          </button>
        </div>
      </form>
    </div>
  );
};

export default AddFeedForm;
