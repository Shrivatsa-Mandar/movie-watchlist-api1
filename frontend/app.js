const API_URL = "http://localhost:9090";

function createUser() {
    const username = document.getElementById("username").value;
    const email = document.getElementById("email").value;

    fetch(`${API_URL}/users`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ username, email })
    })
    .then(res => res.json())
    .then(data => {
        document.getElementById("user-msg").innerText = `User created with ID: ${data.id}`;
    });
}

function searchMovie() {
    const title = document.getElementById("movie-title").value;

    fetch(`${API_URL}/movies/search?title=${title}`)
    .then(res => res.json())
    .then(data => {
        document.getElementById("movie-result").innerText = JSON.stringify(data, null, 2);
    });
}

function addToWatchlist() {
    const user_id = parseInt(document.getElementById("user-id").value);
    const movie_id = parseInt(document.getElementById("movie-id").value);

    fetch(`${API_URL}/watchlist`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ user_id, movie_id, status: "WATCHLIST", rating: 0 })
    })
    .then(res => res.json())
    .then(data => {
        document.getElementById("watchlist-result").innerText = "Added to watchlist!";
    });
}

function getWatchlist() {
    const user_id = parseInt(document.getElementById("user-id").value);

    fetch(`${API_URL}/watchlist/${user_id}`)
    .then(res => res.json())
    .then(data => {
        document.getElementById("watchlist-result").innerText = JSON.stringify(data, null, 2);
    });
}