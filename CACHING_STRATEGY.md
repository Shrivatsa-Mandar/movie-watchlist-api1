# Caching Strategy: Movie Watchlist & Recommendation API

To ensure high performance and reduce unnecessary calls to external APIs (TMDB), we implement a multi-layered caching strategy.

## 1. External API (TMDB) Integration
The API search endpoint first queries our local SQLite database before making an external network request.

## 2. SQLite "Persistent Cache"
We use a table called `movies_cache` to store movie metadata fetched from TMDB.
- **Workflow**: 
    1. User searches for a movie.
    2. API checks if a similar title exists in `movies_cache`.
    3. If found, returns local data immediately (Latency ~5ms).
    4. If not found, calls TMDB API (Latency ~200-500ms).
    5. The first result from TMDB is then saved into `movies_cache` for future users.

## 3. Cache Schema
| Field | Purpose |
|-------|---------|
| `tmdb_id` | Unique identifier to prevent duplicate entries of the same movie. |
| `genre` | Cached locally to power the recommendation engine without re-fetching details. |
| `poster_url` | Cached to ensure the UI can render images quickly. |

## 4. Benefits
- **Rate Limit Protection**: Prevents hitting TMDB rate limits.
- **Offline Capability**: Movies once searched are available even if the external API is down.
- **Recommendation Speed**: Recommendations are calculated entirely via SQL joins on the local cache, making them extremely fast.
