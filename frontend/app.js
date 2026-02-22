const API_URL = "http://localhost:9097";

let currentUser = JSON.parse(localStorage.getItem('user')) || null;
let token = localStorage.getItem('token') || null;

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    updateUI();
    if (token) {
        loadWatchlist();
        loadRecommendations();
    }
});

function updateUI() {
    const authSection = document.getElementById('auth-section');
    const mainContent = document.getElementById('main-content');
    const topNav = document.getElementById('top-nav');
    const userDisplay = document.getElementById('user-display');

    if (token && currentUser) {
        authSection.classList.add('hidden');
        mainContent.classList.remove('hidden');
        topNav.classList.remove('hidden');
        userDisplay.innerText = `👋 Hi, ${currentUser.username}`;
    } else {
        authSection.classList.remove('hidden');
        mainContent.classList.add('hidden');
        topNav.classList.add('hidden');
    }
}

// Auth functions
function showRegister() {
    document.getElementById('login-card').classList.add('hidden');
    document.getElementById('register-card').classList.remove('hidden');
}

function showLogin() {
    document.getElementById('login-card').classList.remove('hidden');
    document.getElementById('register-card').classList.add('hidden');
}

async function register() {
    const username = document.getElementById('reg-username').value;
    const email = document.getElementById('reg-email').value;
    const password = document.getElementById('reg-password').value;

    try {
        const res = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, email, password })
        });
        if (res.ok) {
            document.getElementById('reg-msg').innerText = "Registration successful! Please login.";
            setTimeout(showLogin, 2000);
        } else {
            const data = await res.text();
            document.getElementById('reg-msg').innerText = data;
        }
    } catch (err) {
        document.getElementById('reg-msg').innerText = "Error connecting to server.";
    }
}

async function login() {
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;

    try {
        const res = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });
        if (res.ok) {
            const data = await res.json();
            token = data.token;
            currentUser = { username: data.username, id: data.user_id };
            localStorage.setItem('token', token);
            localStorage.setItem('user', JSON.stringify(currentUser));
            updateUI();
            loadWatchlist();
            loadRecommendations();
        } else {
            document.getElementById('login-msg').innerText = "Invalid credentials";
        }
    } catch (err) {
        document.getElementById('login-msg').innerText = "Error connecting to server.";
    }
}

function logout() {
    token = null;
    currentUser = null;
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    updateUI();
}

// Movie functions
async function searchMovie() {
    const title = document.getElementById('movie-title').value;
    const resultsDiv = document.getElementById('search-results');
    resultsDiv.innerHTML = '<p class="animate">Searching...</p>';

    try {
        const res = await fetch(`${API_URL}/movies/search?title=${title}`);
        const movies = await res.json();

        resultsDiv.innerHTML = '';
        movies.forEach(movie => {
            const card = createMovieCard(movie, true);
            resultsDiv.appendChild(card);
        });
    } catch (err) {
        resultsDiv.innerHTML = '<p>Search failed.</p>';
    }
}

function createMovieCard(movie, isSearch = false) {
    const div = document.createElement('div');
    div.className = 'movie-card animate';

    div.innerHTML = `
        <img src="${movie.poster_url}" alt="${movie.title}" class="movie-poster">
        <div class="movie-info">
            <h3 class="movie-title">${movie.title}</h3>
            <p class="movie-meta">${movie.year} • ${movie.genre}</p>
            ${isSearch ? `<button style="width: 100%;" onclick='addToWatchlist(${JSON.stringify(movie)})'>Add to Watchlist</button>` : ''}
            ${!isSearch && movie.status === 'WATCHLIST' ? `<button class="btn-secondary" style="width: 100%;" onclick="markWatched(${movie.id})">Mark Watched</button>` : ''}
        </div>
    `;
    return div;
}

async function addToWatchlist(movie) {
    try {
        const res = await fetch(`${API_URL}/api/watchlist`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({
                user_id: parseInt(currentUser.id),
                movie_id: movie.id,
                status: 'WATCHLIST',
                rating: 0
            })
        });
        if (res.ok) {
            alert("Added!");
            loadWatchlist();
        }
    } catch (err) {
        alert("Failed to add.");
    }
}

async function loadWatchlist() {
    const resultsDiv = document.getElementById('watchlist-results');
    try {
        const res = await fetch(`${API_URL}/api/watchlist/${currentUser.id}`, {
            headers: { 'Authorization': `Bearer ${token}` }
        });
        const items = await res.json();
        resultsDiv.innerHTML = items?.length ? '' : '<p style="color: var(--text-dim);">Your watchlist is empty.</p>';

        // Note: For a real app, you'd fetch movie details for each item. 
        // Here we'll simplify and just show IDs or you could expand models.
        if (items) {
            items.forEach(item => {
                const div = document.createElement('div');
                div.className = 'card animate';
                div.innerHTML = `<p>Movie ID: ${item.movie_id}</p><p>Status: ${item.status}</p>`;
                resultsDiv.appendChild(div);
            });
        }
    } catch (err) { }
}

async function loadRecommendations() {
    const resultsDiv = document.getElementById('recommendation-results');
    try {
        const res = await fetch(`${API_URL}/api/recommendations/${currentUser.id}`, {
            headers: { 'Authorization': `Bearer ${token}` }
        });
        const movies = await res.json();
        resultsDiv.innerHTML = movies?.length ? '' : '<p style="color: var(--text-dim);">No recommendations yet.</p>';

        if (movies) {
            movies.forEach(movie => {
                resultsDiv.appendChild(createMovieCard(movie, false));
            });
        }
    } catch (err) { }
}