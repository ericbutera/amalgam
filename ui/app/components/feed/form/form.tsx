"use client";

import { useForm } from "react-hook-form";
import { useFormStore } from "./store";

import { getGraph } from "../../../lib/fetch";

export default function FeedForm() {
  // const { mutate } = useSWRConfig();
  const { formData, setFormData, resetForm } = useFormStore();

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm();

  const onSubmit = async (data) => {
    try {
      // const resp = getApi().feedsPost({
      //   request: {
      //     feed: {
      //       url: data.url,
      //     },
      //   },
      // });
      const resp = await getGraph().AddFeed({
        url: data.url,
        name: data.name,
      });

      console.log("response id: %o", resp.addFeed.id);
      alert("Feed created successfully!");
      resetForm();
    } catch (error) {
      console.error(error);
      alert("Failed to create feed");
    }
  };

  const registerWithSync = (field) => {
    // Sync Zustand's state through react-hook-form's onChange
    return {
      ...register(field),
      onChange: (e) => {
        setFormData(field, e.target.value); // Sync Zustand
        setValue(field, e.target.value); // Update react-hook-form
      },
    };
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div>
        <label htmlFor="name">Name:</label>
        <input id="name" value={formData.name} {...registerWithSync("name")} />
      </div>

      <div>
        <label htmlFor="url">URL (required):</label>
        <input
          type="url"
          id="url"
          value={formData.url}
          {...registerWithSync("url", { required: "URL is required" })}
        />
        {errors.url && <p style={{ color: "red" }}>{errors.url.message}</p>}
      </div>

      <button type="submit">Submit</button>
    </form>
  );
}
