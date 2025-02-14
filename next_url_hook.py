import requests

def fetch_all_pages(url, params=None, headers=None):
    while url:
        response = requests.get(url, params=params, headers=headers)
        response.raise_for_status()
        data = response.json()
        yield data.get("results", [])
        # Update the URL to the next page if it exists
        url = data.get("next_url")
        # Reset params after the first call if needed
        params = None

# Usage
for page_results in fetch_all_pages("https://api.example.com/items"):
    for item in page_results:
        print(item)

