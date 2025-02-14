async function* fetchAllPages(url, params = {}, headers = {}) {
  while (url) {
    const response = await fetch(url, { headers });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    yield data.results || [];
    url = data.next_url;
    // Clear params for subsequent calls if necessary
    params = {};
  }
}

// Usage with async iteration
(async () => {
  for await (const pageResults of fetchAllPages("https://api.example.com/items")) {
    pageResults.forEach(item => console.log(item));
  }
})();

