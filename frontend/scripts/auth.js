const API_BASE_URL = "https://localhost/api/v1"; // TODO: update

async function getUser() {
  try {
    const response = await fetch(`${API_BASE_URL}/user/me`);

    if (response.ok) {
      data = await response.json();

      updateLoginLink(document.getElementById("login-link"), data.username);

      return data;
    } else {
      throw new Error(response.status);
    }
  } catch (error) {
    console.log("There was an error", error);
  }
}

function updateLoginLink(loginLink, username) {
  loginLink.innerHTML = "Hi " + username + "!";
  loginLink.href = "./profile.html";
}
