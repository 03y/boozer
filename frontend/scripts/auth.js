const API_BASE_URL = "https://192.168.0.34:6969"; // TODO: update

async function getUser() {
  let token = localStorage.getItem("token");

  if (token == null) return;

  try {
    let myHeaders = new Headers();
    myHeaders.append("Authorization", "Bearer " + token);
    let requestOptions = {
      method: "GET",
      headers: myHeaders,
      redirect: "follow",
    };

    const response = await fetch(`${API_BASE_URL}/user/me`, requestOptions);

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

async function getConsumptionCount(userId) {
  let token = localStorage.getItem("token");

  if (token == null) {
    console.log("No token found. Cannot fetch consumptions.");
    return;
  }

  try {
    let requestOptions = {
      method: "GET",
      redirect: "follow",
    };

    const response = await fetch(
      `${API_BASE_URL}/consumption_count/${userId}`,
      requestOptions,
    );

    if (response.ok) {
      const data = await response.json();
      return data;
    } else {
      throw new Error(response.status);
    }
  } catch (error) {
    console.error("There was an error fetching consumptions:", error);
  }
}

function updateLoginLink(loginLink, username) {
  loginLink.innerHTML = "Hi " + username + "!";
  loginLink.href = "./profile.html";
}
