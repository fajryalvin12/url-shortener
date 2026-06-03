"use client";

import { useState, useRef, useEffect } from "react";

const MAX_HISTORY = 10;
const STORAGE_KEY = "url_shortener_history";

function validateUrl(str) {
  try {
    const url = new URL(str);
    return url.protocol === "http:" || url.protocol === "https:";
  } catch {
    return false;
  }
}

function loadHistory() {
  if (typeof window === "undefined") return [];
  try {
    return JSON.parse(localStorage.getItem(STORAGE_KEY) || "[]");
  } catch {
    return [];
  }
}

function saveHistory(list) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(list));
}

export default function Home() {
  const [url, setUrl] = useState("");
  const [result, setResult] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [copied, setCopied] = useState(false);
  const [history, setHistory] = useState([]);
  const [copiedIndex, setCopiedIndex] = useState(null);
  const abortRef = useRef(null);

  useEffect(() => {
    setHistory(loadHistory());
  }, []);

  async function handleShorten() {
    setError("");

    if (!url.trim()) {
      setError("URL tidak boleh kosong.");
      return;
    }
    if (!validateUrl(url.trim())) {
      setError("Format URL tidak valid. Pastikan diawali https://");
      return;
    }

    if (abortRef.current) abortRef.current.abort();
    abortRef.current = new AbortController();

    setIsLoading(true);
    setResult(null);
    setCopied(false);

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/v1/shorten`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ original_url: url.trim() }),
        signal: abortRef.current.signal,
      });

      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.message || `Error ${res.status}`);
      }

      const data = await res.json();
      const shortUrl = data.data?.short_url;

      if (!shortUrl || !shortUrl.startsWith("http")) {
        throw new Error("Respons server tidak valid.");
      }

      setResult(shortUrl);
      setUrl("");

      const newEntry = {
        short: shortUrl,
        original: url.trim(),
        createdAt: new Date().toISOString(),
      };
      const updated = [newEntry, ...history].slice(0, MAX_HISTORY);
      setHistory(updated);
      saveHistory(updated);
    } catch (err) {
      if (err.name === "AbortError") return;
      setError(err.message || "Gagal terhubung ke server. Coba lagi.");
    } finally {
      setIsLoading(false);
    }
  }

  function handleKeyDown(e) {
    if (e.key === "Enter") handleShorten();
  }

  async function copyToClipboard(text, index = null) {
    try {
      await navigator.clipboard.writeText(text);
      if (index === null) {
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
      } else {
        setCopiedIndex(index);
        setTimeout(() => setCopiedIndex(null), 2000);
      }
    } catch {
      // fallback for older browsers
      const el = document.createElement("textarea");
      el.value = text;
      document.body.appendChild(el);
      el.select();
      document.execCommand("copy");
      document.body.removeChild(el);
    }
  }

  function clearHistory() {
    setHistory([]);
    saveHistory([]);
  }

  function formatTime(iso) {
    const diff = Math.floor((Date.now() - new Date(iso)) / 1000);
    if (diff < 60) return "Baru saja";
    if (diff < 3600) return `${Math.floor(diff / 60)} menit lalu`;
    if (diff < 86400) return `${Math.floor(diff / 3600)} jam lalu`;
    return `${Math.floor(diff / 86400)} hari lalu`;
  }

  return (
    <main className="min-h-screen bg-[#0f0f0f] flex flex-col items-center justify-start px-4 py-14">
      {/* Brand */}
      <div className="flex items-center gap-3 mb-12">
        <div className="w-9 h-9 bg-[#1a1a1a] border border-[#2a2a2a] rounded-xl flex items-center justify-center text-violet-400 text-lg">
          🔗
        </div>
        <span className="text-[15px] font-medium text-neutral-200 tracking-tight">
          Shortify
        </span>
      </div>

      {/* Main Card */}
      <div className="w-full max-w-[480px] bg-[#1a1a1a] border border-[#2a2a2a] rounded-2xl p-7">
        <h1 className="text-[22px] font-medium text-neutral-100 mb-1">
          Shorten your URL
        </h1>
        <p className="text-[13px] text-neutral-500 mb-6">
          Paste a long link, get a clean short one instantly.
        </p>

        {/* Input */}
        <label
          htmlFor="url-input"
          className="block text-[12px] text-neutral-500 mb-2"
        >
          Paste URL
        </label>
        <div className="relative mb-3">
          <span className="absolute left-3 top-1/2 -translate-y-1/2 text-neutral-600 text-[15px] pointer-events-none">
            🌐
          </span>
          <input
            id="url-input"
            type="url"
            value={url}
            onChange={(e) => {
              setUrl(e.target.value);
              setError("");
            }}
            onKeyDown={handleKeyDown}
            placeholder="https://example.com/very/long/url/here"
            autoComplete="off"
            aria-describedby={error ? "url-error" : undefined}
            className={`w-full bg-[#111] border rounded-xl pl-9 pr-4 py-[10px] text-[14px] text-neutral-200 placeholder-neutral-600 outline-none transition-colors
              ${error ? "border-red-600 focus:border-red-500" : "border-[#2e2e2e] focus:border-violet-600"}`}
          />
        </div>

        {error && (
          <p
            id="url-error"
            role="alert"
            className="text-[12px] text-red-400 mb-3 -mt-1"
          >
            {error}
          </p>
        )}

        <button
          onClick={handleShorten}
          disabled={isLoading}
          aria-busy={isLoading}
          className="w-full bg-violet-700 hover:bg-violet-600 active:scale-[0.99] disabled:opacity-60 disabled:cursor-not-allowed
            rounded-xl py-[10px] text-[14px] font-medium text-white transition-colors flex items-center justify-center gap-2"
        >
          {isLoading ? (
            <>
              <span className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin" />
              Shortening...
            </>
          ) : (
            "Shorten URL"
          )}
        </button>

        {/* Result */}
        {result && (
          <div
            className="mt-5 bg-[#111] border border-[#2a2a2a] rounded-xl px-4 py-3 flex items-center justify-between gap-3"
            aria-live="polite"
          >
            <div className="min-w-0">
              <p className="text-[11px] text-neutral-600 mb-1">Short URL</p>
              <a
                href={result}
                target="_blank"
                rel="noopener noreferrer"
                className="text-[13px] text-violet-400 font-mono break-all hover:underline"
              >
                {result}
              </a>
            </div>
            <button
              onClick={() => copyToClipboard(result)}
              className={`shrink-0 border rounded-lg px-3 py-[6px] text-[12px] flex items-center gap-1.5 transition-colors
                ${
                  copied
                    ? "text-green-400 border-green-600 bg-green-950"
                    : "text-neutral-400 border-[#2e2e2e] bg-[#1f1f1f] hover:bg-[#2a2a2a]"
                }`}
              aria-label="Copy short URL"
            >
              {copied ? "✓ Copied!" : "⎘ Copy"}
            </button>
          </div>
        )}
      </div>

      {/* History */}
      {history.length > 0 && (
        <div className="w-full max-w-[480px] mt-6">
          <div className="flex items-center justify-between mb-3">
            <span className="text-[12px] text-neutral-600">Recent links</span>
            <button
              onClick={clearHistory}
              className="text-[11px] text-neutral-700 hover:text-red-400 transition-colors"
            >
              Clear all
            </button>
          </div>
          <div className="flex flex-col gap-2">
            {history.map((item, i) => (
              <div
                key={i}
                className="bg-[#1a1a1a] border border-[#232323] rounded-xl px-4 py-3 flex items-center justify-between gap-3"
              >
                <div className="min-w-0 flex-1">
                  <p className="text-[13px] text-violet-400 font-mono truncate">
                    {item.short}
                  </p>
                  <p className="text-[11px] text-neutral-600 truncate mt-0.5">
                    {item.original}
                  </p>
                </div>
                <span className="text-[11px] text-neutral-700 shrink-0">
                  {formatTime(item.createdAt)}
                </span>
                <button
                  onClick={() => copyToClipboard(item.short, i)}
                  className={`shrink-0 p-1.5 rounded-lg transition-colors ${
                    copiedIndex === i
                      ? "text-green-400"
                      : "text-neutral-600 hover:text-violet-400"
                  }`}
                  aria-label={`Copy ${item.short}`}
                >
                  {copiedIndex === i ? "✓" : "⎘"}
                </button>
              </div>
            ))}
          </div>
        </div>
      )}
    </main>
  );
}
