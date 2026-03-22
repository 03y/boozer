const API_BASE_URL = "https://192.168.1.3/api/v2"; // TODO: update

async function getUser() {
    const response = await fetch(`${API_BASE_URL}/users/me`, {credentials: 'include'});

    if (response.ok) {
      data = await response.json();

      updateLoginLink(document.getElementById("login-link"), data.username);

      if (data.admin) {
        const nav = document.querySelector('nav');
        const adminLink = document.createElement('a');
        adminLink.href = './admin.html';
        adminLink.textContent = 'Admin';
        nav.appendChild(adminLink);
      }

      return data;
    } else {
        return response.status;
    }
}

function updateLoginLink(loginLink, username) {
    loginLink.innerHTML = "Hi " + username + "!";
    loginLink.href = "./profile.html";
}

async function logout() {
    try {
        const response = await fetch(`${API_BASE_URL}/logout`, {
            method: "POST",
        });
        if (response.ok) {
            window.location.href = "index.html";
        } else {
            console.error("Logout failed");
        }
    } catch (error) {
        console.error("Error logging out:", error);
    }
}
