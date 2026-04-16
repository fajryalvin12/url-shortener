"use client";

import { useState } from "react";

export default function Home() {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [shortUrl, setShortUrl] = useState("");

  async function onSubmit(event) {
    event.preventDefault();
    setIsLoading(true);
    setError(null);

    const url = event.currentTarget.original_url.value;

    try {
      const response = await fetch("http://localhost:8000/v1/shorten", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ original_url: url }),
      });

      if (!response.ok) {
        setError(response.message);
      }

      const data = await response.json();
      setShortUrl(data.data.short_url);
    } catch (error) {
      console.error(error);
      setError(error.message);
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <>
      <div className="h-screen flex justify-center items-center">
        <form
          onSubmit={onSubmit}
          className="flex flex-col gap-4 bg-black p-6 rounded-xl text-white"
        >
          <h1>URL Shortener App</h1>
          <div className="flex flex-col">
            <label>Insert the URL here : </label>
            <input
              name="original_url"
              className="border rounded-md outline-none p-1"
              type="text"
            />
          </div>
          <button
            type="submit"
            className="border rounded-md bg-white text-black"
            disabled={isLoading}
          >
            {isLoading ? "Loading..." : "Submit!"}
          </button>
          {error && <div style={{ color: "red" }}>{error}</div>}
          {shortUrl && (
            <>
              <label style={{ color: "green" }}>Short URL :</label>
              <a href={shortUrl} style={{ color: "green" }}>
                {shortUrl}
              </a>
            </>
          )}
        </form>
      </div>
    </>
  );
}
